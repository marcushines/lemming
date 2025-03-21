// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gnmi contains a reference on-device gNMI implementation.
package gnmi

import (
	"context"
	"fmt"
	"runtime/debug"
	"slices"
	"strconv"
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/openconfig/gnmi/subscribe"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	"github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/lemming/gnmi/reconciler"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

const (
	OpenConfigOrigin = "openconfig"
)

const (
	enableDebugLog = true
)

// Server is a reference gNMI implementation.
type Server struct {
	// The subscribe Server implements only Subscribe for gNMI.
	*subscribe.Server
	c *Collector

	// TODO(wenbli): Implement gnmi.Get and remove this.
	GetResponses []interface{}

	configMu     sync.Mutex
	configSchema *ytypes.Schema

	stateMu     sync.Mutex
	stateSchema *ytypes.Schema

	validators  []func(*oc.Root) error
	reconcilers []reconciler.Reconciler

	pathAuth PathAuth
}

// New creates and registers a reference gNMI server on the given gRPC server.
//
// To avoid having to specify the Target field in gNMI requests in order to comply
// with https://www.openconfig.net/docs/gnmi/gnmi-specification/#2221-path-target,
// register the gRPC interceptor
// grpc.StreamInterceptor(gnmi.NewSubscribeTargetUpdateInterceptor(targetName))
// when creating your grpc.Server object.
//
// - targetName is the gNMI target name of the datastore.
// - pa is an optional PathAuth instance used for authorization gNMI requests, set to nil for no authorization.
func New(srv *grpc.Server, targetName string, pa PathAuth, recs ...reconciler.Reconciler) (*Server, error) {
	gnmiServer, err := newServer(context.Background(), targetName, true, recs...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gNMI server: %v", err)
	}
	gnmiServer.pathAuth = pa
	gpb.RegisterGNMIServer(srv, gnmiServer)
	return gnmiServer, nil
}

// newServer returns a new reference gNMI server implementation that can be
// registered on an existing gRPC server.
//
// - configSchema is the specification of the schema if gnmi.Set on config paths is used.
// - stateSchema is the specification of the schema if gnmi.Set on state paths is used.
// - targetName is the name of the target.
func newServer(ctx context.Context, targetName string, enableSet bool, recs ...reconciler.Reconciler) (*Server, error) {
	c := NewCollector(targetName)
	subscribeSrv, err := c.Start(ctx, false)
	if err != nil {
		return nil, err
	}

	gnmiServer := &Server{
		Server:      subscribeSrv, // use the 'subscribe' implementation.
		c:           c,
		reconcilers: recs,
	}

	if !enableSet {
		return gnmiServer, nil
	}

	emptySchema, err := oc.Schema()
	if err != nil {
		return nil, fmt.Errorf("cannot create ygot schema object: %v", err)
	}

	configSchema, err := oc.Schema()
	if err != nil {
		return nil, fmt.Errorf("cannot create ygot schema object: %v", err)
	}
	// Initialize the cache with the input schema root.
	if configSchema != nil {
		if err := setupSchema(configSchema, true); err != nil {
			return nil, err
		}
		if err := ygot.PruneConfigFalse(configSchema.RootSchema(), configSchema.Root); err != nil {
			return nil, fmt.Errorf("gnmi: %v", err)
		}
		if err := updateCache(c, configSchema.Root, emptySchema.Root, OpenConfigOrigin, true, 0, "", nil); err != nil {
			return nil, fmt.Errorf("gnmi newServer: %v", err)
		}
	}

	stateSchema, err := oc.Schema()
	if err != nil {
		return nil, fmt.Errorf("cannot create ygot schema object: %v", err)
	}
	if stateSchema != nil {
		if err := setupSchema(stateSchema, false); err != nil {
			return nil, err
		}
		if err := updateCache(c, stateSchema.Root, emptySchema.Root, OpenConfigOrigin, true, 0, "", nil); err != nil {
			return nil, fmt.Errorf("gnmi newServer: %v", err)
		}
	}

	for _, rec := range recs {
		if len(rec.ValidationPaths()) > 0 {
			gnmiServer.validators = append(gnmiServer.validators, rec.Validate)
		}
	}

	gnmiServer.configSchema = configSchema
	gnmiServer.stateSchema = stateSchema

	return gnmiServer, nil
}

type populateDefaultser interface {
	PopulateDefaults()
}

// setupSchema takes in a ygot schema object which it assumes to be
// uninitialized. It initializes and validates it, returning any errors
// encountered.
//
// If config is set, then default values are automatically populated.
// State paths are not automatically populated since internal goroutines should
// populate them instead.
func setupSchema(schema *ytypes.Schema, config bool) error {
	if !schema.IsValid() {
		return fmt.Errorf("cannot obtain valid schema for GoStructs: %v", schema)
	}

	if config {
		schema.Root.(populateDefaultser).PopulateDefaults()
	}
	if err := schema.Validate(); err != nil {
		return fmt.Errorf("default root of input schema fails validation: %v", err)
	}

	return nil
}

// updateCache updates the cache with the difference between the root ->
// dirtyRoot such that if the root represents the cache, then the dirtyRoot
// will represent the cache afterwards.
//
// If root is nil, then it is assumed the cache is empty, and the entirety of
// the dirtyRoot is put into the cache.
// - auth adds authorization to before writing vals to the cache, if set to nil, not authorization is checked.
func updateCache(collector *Collector, dirtyRoot, root ygot.GoStruct, origin string, preferShadowPath bool, timestamp int64, user string, auth PathAuth) error {
	var nos []*gpb.Notification
	if root == nil {
		var err error
		if nos, err = ygot.TogNMINotifications(dirtyRoot, timestamp, ygot.GNMINotificationsConfig{
			UsePathElem: true,
		}); err != nil {
			return fmt.Errorf("gnmi: %v", err)
		}
	} else {
		var err error
		if nos, err = ygot.DiffWithAtomic(root, dirtyRoot, &ygot.DiffPathOpt{PreferShadowPath: preferShadowPath}); err != nil {
			return fmt.Errorf("gnmi: error while creating update notification for Set: %v", err)
		}
		for _, n := range nos {
			n.Timestamp = timestamp
		}
	}

	if auth != nil && auth.IsInitialized() {
		// Check authorization of the diff to check if implicit deletes (caused by replaces) are allowed.
		allowed, err := checkWritePermission(auth, user, nos...)
		if err != nil {
			return err
		}
		if !allowed {
			return status.Errorf(codes.PermissionDenied, "cannot set all paths in request")
		}
	}

	return updateCacheNotifs(collector, nos, origin)
}

func checkWritePermission(auth PathAuth, user string, nos ...*gpb.Notification) (bool, error) {
	for _, no := range nos {
		for _, del := range no.Delete {
			p, err := util.JoinPaths(no.GetPrefix(), del)
			if err != nil {
				return false, err
			}
			if !auth.CheckPermit(p, user, true) {
				log.V(1).Infof("user %q not allowed to set path %s", user, p)
				return false, nil
			}
		}
		for _, upd := range no.Update {
			p, err := util.JoinPaths(no.GetPrefix(), upd.GetPath())
			if err != nil {
				return false, err
			}
			if !auth.CheckPermit(p, user, true) {
				log.V(1).Infof("user %q not allowed to set path %s", user, p)
				return false, nil
			}
		}
	}
	return true, nil
}

// updateCacheNotifs updates the cache with the given notifications.
func updateCacheNotifs(c *Collector, nos []*gpb.Notification, origin string) error {
	for _, n := range nos {
		if n.Prefix == nil {
			n.Prefix = &gpb.Path{}
		}
		n.Prefix.Origin = origin
		if n.Prefix.Origin == "" {
			n.Prefix.Origin = OpenConfigOrigin
		}

		var pathsForDelete []string
		for _, path := range n.Delete {
			p, err := ygot.PathToString(path)
			if err != nil {
				return fmt.Errorf("cannot convert deleted path to string: %v", err)
			}
			pathsForDelete = append(pathsForDelete, p)
		}
		if len(n.Update) > 0 {
			log.V(3).Infof("datastore: updating the following values: %+v", n.Update)
		}
		if len(pathsForDelete) > 0 {
			log.V(3).Infof("datastore: deleting the following paths: %+v", pathsForDelete)
		}
		log.V(3).Infof("datastore: calling GnmiUpdate with the following notification:\n%s", prototext.Format(n))
		if err := c.GnmiUpdate(n); err != nil {
			return fmt.Errorf("%w: notification:\n%s\n%s", err, prototext.Format(n), string(debug.Stack()))
		}
		if enableDebugLog && (len(n.Delete) != 0 || len(n.Update) != 0) {
			logUpdate(n)
		}
	}
	return nil
}

func logUpdate(n *gpb.Notification) {
	prefix, err := ygot.PathToString(n.Prefix)
	if err != nil {
		return
	}
	for _, d := range n.Delete {
		path, err := ygot.PathToString(d)
		if err != nil {
			return
		}
		log.V(1).Infof("ts: %d, pre: %s, delete: %s", n.GetTimestamp(), prefix, path)
	}
	for _, u := range n.Update {
		path, err := ygot.PathToString(u.GetPath())
		if err != nil {
			return
		}
		log.V(1).Infof("ts: %d, pre: %s, update %s: %v\n", n.GetTimestamp(), prefix, path, u.GetVal())
	}
}

// unmarshalSetRequest unmarshals the setrequest into the schema.
// Where preferShadowPath=true, this means that any default configuration will
// be automatically populated in the schema.
func unmarshalSetRequest(schema *ytypes.Schema, req *gpb.SetRequest, preferShadowPath bool) error {
	var unmarshalOpts []ytypes.UnmarshalOpt
	if preferShadowPath {
		unmarshalOpts = append(unmarshalOpts, &ytypes.PreferShadowPath{})
	}
	if err := ytypes.UnmarshalSetRequest(schema, req, unmarshalOpts...); err != nil {
		return status.Errorf(codes.InvalidArgument, "failed to unmarshal set request %v", err)
	}
	if preferShadowPath {
		// Populate and defaults after any possible replace/delete operations.
		// NOTE: This statement can introduce significant slowdowns.
		schema.Root.(populateDefaultser).PopulateDefaults()
	}
	return nil
}

// set updates the datastore and intended configuration with the SetRequest,
// allowing read-only values to be updated.
//
// set returns a gRPC error with the correct code and shouldn't be wrapped again.
//
// - timestamp specifies the timestamp of the values that are to be updated in
// the gNMI cache. If zero, then time.Now().UnixNano() is used.
// - auth adds authorization to before writing vals to the cache, if set to nil, not authorization is checked.
func set(schema *ytypes.Schema, c *Collector, req *gpb.SetRequest, preferShadowPath bool, validators []func(*oc.Root) error, timestamp int64, user string, auth PathAuth) error {
	// skip diffing and deepcopy for performance when handling state update paths.
	// Currently this is not possible for replace/delete paths, since
	// without doing a diff, it is not possible to compute what was
	// deleted. Furthermore, since we're using a single cache, we would be
	// affecting both config/state leafs at the same time.
	//
	// TODO: Once unmarshalSetRequest can return a diff itself, then this
	// block can be deleted.
	if !preferShadowPath && len(req.Delete)+len(req.Replace) == 0 {
		tempSchema := &ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: schema.SchemaTree,
			Unmarshal:  schema.Unmarshal,
		}
		if err := unmarshalSetRequest(tempSchema, req, preferShadowPath); err != nil {
			return fmt.Errorf("error while unmarshalling SetRequest: %v", err)
		}
		notifs, err := ygot.TogNMINotifications(tempSchema.Root, timestamp, ygot.GNMINotificationsConfig{UsePathElem: true})
		if err != nil {
			return err
		}

		if err := updateCacheNotifs(c, notifs, req.Prefix.Origin); err != nil {
			return err
		}

		return unmarshalSetRequest(schema, req, preferShadowPath)
	}

	prevRoot, err := ygot.DeepCopy(schema.Root)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to ygot.DeepCopy the cached root object: %v", err)
	}

	success := false

	// Rollback function
	defer func() {
		if !success {
			log.V(1).Infof("Rolling back set request: %v", req)
			schema.Root = prevRoot
		}
	}()

	if err := unmarshalSetRequest(schema, req, preferShadowPath); err != nil {
		return err
	}

	// Only validate for configuration changes since state changes are
	// internal and assumed to be correct. State changes in general also
	// occur much more frequently than config changes, so also want to
	// avoid the performance hit.
	if preferShadowPath {
		if err := schema.Validate(); err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid SetRequest: %v", err)
		}
		for _, validator := range validators {
			if err := validator(schema.Root.(*oc.Root)); err != nil {
				return status.Errorf(codes.InvalidArgument, "validation error: %v", err)
			}
		}
	}

	if err := updateCache(c, schema.Root, prevRoot, req.Prefix.Origin, preferShadowPath, timestamp, user, auth); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	success = true

	return nil
}

const (
	// InternalOrigin is a special gNMI path origin used to store schemaless values.
	InternalOrigin = "lemming-internal"
)

// handleInternalOrigin handles SetRequests whose paths contain schemaless
// values.
//
// An error is returned if the request doesn't purely contain schemaless
// values.
func (s *Server) handleInternalOrigin(req *gpb.SetRequest) (bool, error) {
	if req.Prefix == nil {
		req.Prefix = &gpb.Path{}
	}
	notif := &gpb.Notification{
		Prefix: &gpb.Path{
			Origin: InternalOrigin,
			Elem:   req.Prefix.Elem,
			Target: req.Prefix.Target,
		},
		Timestamp: time.Now().UnixNano(),
	}
	var hasInternal bool
	var hasExternal bool

	for _, del := range req.Delete {
		if del.Origin == InternalOrigin {
			hasInternal = true
			notif.Delete = append(notif.Delete, del)
		} else {
			hasExternal = true
		}
	}
	if hasInternal {
		if err := s.c.GnmiUpdate(notif); err != nil {
			return true, err
		}
	}

	notif.Delete = nil

	for _, replace := range req.Replace {
		if replace.Path.Origin == InternalOrigin {
			hasInternal = true
			notif.Update = append(notif.Update, replace)
		} else {
			hasExternal = true
		}
	}
	for _, update := range req.Update {
		if update.Path.Origin == InternalOrigin {
			hasInternal = true
			notif.Update = append(notif.Update, update)
		} else {
			hasExternal = true
		}
	}
	log.V(2).Infof("internal origin notification: %v", notif)
	if hasInternal {
		if err := s.c.GnmiUpdate(notif); err != nil {
			return true, err
		}
	}

	if hasInternal && hasExternal {
		return true, fmt.Errorf("error: SetRequest contained both internal and external origins: %v", prototext.Format(req))
	}

	return hasInternal, nil
}

const (
	usernameKey = "username"
)

// Set implements lemming's gNMI Set operation.
//
// If the given SetRequest is schema compliant AND passes higher-level
// validators, then the gNMI datastore is updated. Otherwise, the gNMI
// datastore is untouched and an error is returned.
//
// Context metadata modifies the behaviour of this API.
// - GNMIModeMetadataKey specifies whether config or state leaves are expected
// in the SetRequest. This is an exclusive expectation: either the update
// consists solely of config values (a user request), or state values (an
// internal goroutine update).
// - TimestampMetadataKey specifies the timestamp for state leaf updates. This
// is to support cases where the data comes from an externally-timestamped
// source.
func (s *Server) Set(ctx context.Context, req *gpb.SetRequest) (*gpb.SetResponse, error) {
	var timestamp int64
	// Use ConfigMode by default so that external users don't need to set metadata.
	gnmiMode := ConfigMode
	md, ok := metadata.FromIncomingContext(ctx)
	var user string
	if ok {
		switch {
		case slices.Contains(md.Get(GNMIModeMetadataKey), string(ConfigMode)):
			gnmiMode = ConfigMode
		case slices.Contains(md.Get(GNMIModeMetadataKey), string(StateMode)):
			gnmiMode = StateMode

			timestampMD := md.Get(TimestampMetadataKey)
			if len(timestampMD) > 0 {
				var err error
				if timestamp, err = strconv.ParseInt(timestampMD[0], 10, 64); err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "timestamp metadata specified in SetRequest cannot be parsed: %v", err)
				}
			}
		}
		p, ok := peer.FromContext(ctx)
		if s.pathAuth != nil && s.pathAuth.IsInitialized() && ok && p.Addr != nil {
			if len(md[usernameKey]) != 1 || md[usernameKey][0] == "" {
				return nil, status.Errorf(codes.Unauthenticated, "no username set in metadata %v", user)
			}
			user = md[usernameKey][0]
		}
	}

	if found, err := s.handleInternalOrigin(req); found {
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error handling set request with internal origin: %v", err)
		}
		return &gpb.SetResponse{}, nil
	}

	switch gnmiMode {
	case ConfigMode:
		s.configMu.Lock()
		defer s.configMu.Unlock()

		log.V(2).Infof("config datastore service received SetRequest: %v", prototext.Format(req))
		if s.configSchema == nil {
			return s.UnimplementedGNMIServer.Set(ctx, req)
		}

		// TODO(wenbli): Reject paths that try to modify read-only values.
		// TODO(wenbli): Question: what to do if there are operational-state values in a container that is specified to be replaced or deleted?
		err := set(s.configSchema, s.c, req, true, s.validators, timestamp, user, s.pathAuth)

		// TODO(wenbli): Currently the SetResponse is not filled.
		return &gpb.SetResponse{
			Timestamp: time.Now().UnixNano(),
		}, err
	case StateMode:
		s.stateMu.Lock()
		defer s.stateMu.Unlock()

		log.V(3).Infof("operational state datastore service received SetRequest: %v", req)
		if s.stateSchema == nil {
			return s.UnimplementedGNMIServer.Set(ctx, req)
		}
		// TODO(wenbli): Reject values that modify config values. We only allow modifying state in this mode.
		// Don't authorize setting state since only internal reconcilers do that.
		if err := set(s.stateSchema, s.c, req, false, nil, timestamp, user, nil); err != nil {
			return &gpb.SetResponse{}, err
		}

		// This mode is intended to be used internally, and the SetResponse doesn't matter.
		return &gpb.SetResponse{}, nil
	default:
		return nil, status.Errorf(codes.Internal, "incoming gNMI SetRequest must specify a valid gnmi.Mode via context metadata: %v", md)
	}
}

func (s *Server) Capabilities(context.Context, *gpb.CapabilityRequest) (*gpb.CapabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Reference Implementation Unimplemented")
}

func (s *Server) Get(context.Context, *gpb.GetRequest) (*gpb.GetResponse, error) {
	if len(s.GetResponses) == 0 {
		return nil, status.Errorf(codes.Unimplemented, "Reference Implementation Unimplemented")
	}
	resp := s.GetResponses[0]
	s.GetResponses = s.GetResponses[1:]
	switch v := resp.(type) {
	case error:
		return nil, v
	case *gpb.GetResponse:
		return v, nil
	default:
		return nil, status.Errorf(codes.DataLoss, "Unknown message type: %T", resp)
	}
}

// PathAuth is an interface for checking authorization for gNMI paths.
type PathAuth interface {
	// CheckPermit returns if the user is allowed to read from or write from in the input path.
	CheckPermit(path *gpb.Path, user string, write bool) bool
	// IsInitialized returns if the authorized has been initialized, if not authorization is not checked.
	IsInitialized() bool
}

// subscribeTargetUpdateStream wraps around the embedded grpc.ServerStream, and
// intercepts the RecvMsg and SendMsg method call to populate gNMI Subscribe's
// to correct Lemming's behavior. The gnmi implementation is borrowed from a collector
// which has different semantics than a device.
type subscribeTargetUpdateStream struct {
	grpc.ServerStream
	// target is the target of the gNMI collector used by the lemming
	// instance.
	target string
	// subscribeTarget stores the target field value of the last
	// SubscribeRequest in the stream.
	subscribeTarget string
}

func (w *subscribeTargetUpdateStream) RecvMsg(m any) error {
	if err := w.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	// Store the target from the subscribe request, then set the target to lemming's target name.
	if req, ok := m.(*gpb.SubscribeRequest); ok {
		sub := req.GetSubscribe()
		if sub != nil {
			w.subscribeTarget = sub.GetPrefix().GetTarget()
			if sub.Prefix == nil {
				sub.Prefix = &gpb.Path{}
			}
			sub.Prefix.Target = w.target
		}
	}
	return nil
}

func (w *subscribeTargetUpdateStream) SendMsg(m any) error {
	// The returned target from the server must match the target sent by the client,
	// override the target in the request.
	// https://www.openconfig.net/docs/gnmi/gnmi-specification/#2221-path-target
	if resp, ok := m.(*gpb.SubscribeResponse); ok {
		if notif := resp.GetUpdate(); notif != nil {
			// A clone of the entire notification is
			// required; otherwise the notification in the
			// collector cache is also altered.
			//
			// TODO(wenbli): We should use a shallow-copy of the
			// proto here for better performance:
			// https://github.com/golang/protobuf/issues/1155
			resp.Response = &gpb.SubscribeResponse_Update{Update: proto.Clone(notif).(*gpb.Notification)}
			resp.GetUpdate().Prefix.Target = w.subscribeTarget
		}
	}
	return w.ServerStream.SendMsg(m)
}

func newSubscribeTargetUpdateStream(s grpc.ServerStream, target string) grpc.ServerStream {
	return &subscribeTargetUpdateStream{
		ServerStream: s,
		target:       target,
	}
}

// NewSubscribeTargetUpdateInterceptor returns a stream interceptor to populate
// gNMI Subscribe's target if it's not populated in the SubscribeRequest
// message in order to retrieve the device's data.
//
// ref: https://www.openconfig.net/docs/gnmi/gnmi-specification/#2221-path-target
func NewSubscribeTargetUpdateInterceptor(target string) func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, newSubscribeTargetUpdateStream(ss, target))
		if err != nil {
			fmt.Printf("ERROR in interceptor stream: %v\n", err)
		}
		return err
	}
}

type subscribeWithAuth struct {
	gpb.GNMI_SubscribeServer
	auth PathAuth
	user string
}

// Send implements gNMI subscribe send with authorization.
func (s *subscribeWithAuth) Send(resp *gpb.SubscribeResponse) error {
	if resp.GetSyncResponse() {
		return s.GNMI_SubscribeServer.Send(resp)
	}
	// Create a copy of the resp so that we don't modify the notification stored in the cache.
	authResp := &gpb.SubscribeResponse{
		Response: &gpb.SubscribeResponse_Update{
			Update: &gpb.Notification{
				Prefix:    resp.GetUpdate().GetPrefix(),
				Timestamp: resp.GetUpdate().GetTimestamp(),
				Atomic:    resp.GetUpdate().GetAtomic(),
			},
		},
	}
	respUpd := resp.Response.(*gpb.SubscribeResponse_Update).Update
	authUpd := authResp.Response.(*gpb.SubscribeResponse_Update).Update

	for _, del := range respUpd.Delete {
		p, err := util.JoinPaths(authUpd.GetPrefix(), del)
		if err != nil {
			return err
		}
		if s.auth.CheckPermit(p, s.user, false) {
			authUpd.Delete = append(authUpd.Delete, del)
		}
	}

	for _, upd := range respUpd.Update {
		p, err := util.JoinPaths(authUpd.GetPrefix(), upd.GetPath())
		if err != nil {
			return err
		}
		if s.auth.CheckPermit(p, s.user, false) {
			authUpd.Update = append(authUpd.Update, upd)
		}
	}
	if len(authUpd.Update) == 0 && len(authUpd.Delete) == 0 {
		return nil
	}

	return s.GNMI_SubscribeServer.Send(authResp)
}

// Subscribe wraps the internal subscribe with optional authorization.
func (s *Server) Subscribe(srv gpb.GNMI_SubscribeServer) error {
	p, ok := peer.FromContext(srv.Context())

	if s.pathAuth == nil || !s.pathAuth.IsInitialized() || !ok || p.Addr == nil { // Addr is nil for calls from the reconcilers.
		return s.Server.Subscribe(srv)
	}
	md, _ := metadata.FromIncomingContext(srv.Context()) // Metadata exists even if not explicitly set by client.
	// TODO: Authentication, for now just looking at the username field.
	user := md["username"]
	if len(user) != 1 || user[0] == "" {
		return status.Errorf(codes.Unauthenticated, "no username set in metadata %v", user)
	}
	sa := &subscribeWithAuth{
		GNMI_SubscribeServer: srv,
		auth:                 s.pathAuth,
		user:                 user[0],
	}

	return s.Server.Subscribe(sa)
}

// LocalClient returns a gNMI client for the server.
func (s *Server) LocalClient() gpb.GNMIClient {
	return newLocalClient(s)
}

// StartReconcilers starts all the reconcilers.
func (s *Server) StartReconcilers(ctx context.Context) error {
	c := s.LocalClient()
	for _, rec := range s.reconcilers {
		if err := rec.Start(ctx, c, s.c.name); err != nil {
			return err
		}
	}
	return nil
}

// StopReconcilers stops all the reconcilers.
func (s *Server) StopReconcilers(ctx context.Context) error {
	for _, rec := range s.reconcilers {
		if err := rec.Stop(ctx); err != nil {
			return err
		}
	}
	return nil
}
