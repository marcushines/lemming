/*
Package lacp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: v0.11.1: (ygot: v0.29.21)
using the following YANG input files:
  - public/release/models/acl/openconfig-acl.yang
  - public/release/models/acl/openconfig-packet-match.yang
  - public/release/models/aft/openconfig-aft-network-instance.yang
  - public/release/models/aft/openconfig-aft-summary.yang
  - public/release/models/aft/openconfig-aft.yang
  - public/release/models/bfd/openconfig-bfd.yang
  - public/release/models/bgp/openconfig-bgp-policy.yang
  - public/release/models/bgp/openconfig-bgp-types.yang
  - public/release/models/extensions/openconfig-metadata.yang
  - public/release/models/gnsi/openconfig-gnsi-acctz.yang
  - public/release/models/gnsi/openconfig-gnsi-authz.yang
  - public/release/models/gnsi/openconfig-gnsi-certz.yang
  - public/release/models/gnsi/openconfig-gnsi-credentialz.yang
  - public/release/models/gnsi/openconfig-gnsi-pathz.yang
  - public/release/models/gnsi/openconfig-gnsi.yang
  - public/release/models/gribi/openconfig-gribi.yang
  - public/release/models/interfaces/openconfig-if-aggregate.yang
  - public/release/models/interfaces/openconfig-if-ethernet-ext.yang
  - public/release/models/interfaces/openconfig-if-ethernet.yang
  - public/release/models/interfaces/openconfig-if-ip-ext.yang
  - public/release/models/interfaces/openconfig-if-ip.yang
  - public/release/models/interfaces/openconfig-if-sdn-ext.yang
  - public/release/models/interfaces/openconfig-interfaces.yang
  - public/release/models/isis/openconfig-isis-policy.yang
  - public/release/models/isis/openconfig-isis.yang
  - public/release/models/lacp/openconfig-lacp.yang
  - public/release/models/lldp/openconfig-lldp-types.yang
  - public/release/models/lldp/openconfig-lldp.yang
  - public/release/models/local-routing/openconfig-local-routing.yang
  - public/release/models/mpls/openconfig-mpls-types.yang
  - public/release/models/multicast/openconfig-pim.yang
  - public/release/models/network-instance/openconfig-network-instance.yang
  - public/release/models/openconfig-extensions.yang
  - public/release/models/optical-transport/openconfig-transport-types.yang
  - public/release/models/ospf/openconfig-ospf-policy.yang
  - public/release/models/ospf/openconfig-ospfv2.yang
  - public/release/models/p4rt/openconfig-p4rt.yang
  - public/release/models/platform/openconfig-platform-common.yang
  - public/release/models/platform/openconfig-platform-controller-card.yang
  - public/release/models/platform/openconfig-platform-cpu.yang
  - public/release/models/platform/openconfig-platform-ext.yang
  - public/release/models/platform/openconfig-platform-fabric.yang
  - public/release/models/platform/openconfig-platform-fan.yang
  - public/release/models/platform/openconfig-platform-integrated-circuit.yang
  - public/release/models/platform/openconfig-platform-linecard.yang
  - public/release/models/platform/openconfig-platform-pipeline-counters.yang
  - public/release/models/platform/openconfig-platform-psu.yang
  - public/release/models/platform/openconfig-platform-software.yang
  - public/release/models/platform/openconfig-platform-transceiver.yang
  - public/release/models/platform/openconfig-platform.yang
  - public/release/models/policy-forwarding/openconfig-policy-forwarding.yang
  - public/release/models/policy/openconfig-policy-types.yang
  - public/release/models/qos/openconfig-qos-elements.yang
  - public/release/models/qos/openconfig-qos-interfaces.yang
  - public/release/models/qos/openconfig-qos-types.yang
  - public/release/models/qos/openconfig-qos.yang
  - public/release/models/relay-agent/openconfig-relay-agent.yang
  - public/release/models/rib/openconfig-rib-bgp.yang
  - public/release/models/sampling/openconfig-sampling-sflow.yang
  - public/release/models/segment-routing/openconfig-segment-routing-types.yang
  - public/release/models/system/openconfig-system-bootz.yang
  - public/release/models/system/openconfig-system-controlplane.yang
  - public/release/models/system/openconfig-system-utilization.yang
  - public/release/models/system/openconfig-system.yang
  - public/release/models/types/openconfig-inet-types.yang
  - public/release/models/types/openconfig-types.yang
  - public/release/models/types/openconfig-yang-types.yang
  - public/release/models/vlan/openconfig-vlan.yang
  - public/third_party/ietf/iana-if-type.yang
  - public/third_party/ietf/ietf-inet-types.yang
  - public/third_party/ietf/ietf-interfaces.yang
  - public/third_party/ietf/ietf-yang-types.yang
  - yang/openconfig-bgp-gue.yang

Imported modules were sourced from:
  - public/release/models/...
  - public/third_party/ietf/...
*/
package lacp

import (
	"reflect"

	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Lacp_Interface_NamePath represents the /openconfig-lacp/lacp/interfaces/interface/state/name YANG schema element.
type Lacp_Interface_NamePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_NamePathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/state/name YANG schema element.
type Lacp_Interface_NamePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/name"
//	Path from root:       "/lacp/interfaces/interface/state/name"
func (n *Lacp_Interface_NamePath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewSingletonQuery[string](
		"Lacp_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lacp_Interface).Name
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/name"
//	Path from root:       "/lacp/interfaces/interface/state/name"
func (n *Lacp_Interface_NamePathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lacp_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lacp_Interface).Name
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/name"
//	Path from root:       "/lacp/interfaces/interface/config/name"
func (n *Lacp_Interface_NamePath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewConfigQuery[string](
		"Lacp_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lacp_Interface).Name
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/name"
//	Path from root:       "/lacp/interfaces/interface/config/name"
func (n *Lacp_Interface_NamePathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lacp_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lacp_Interface).Name
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Lacp_Interface_SystemIdMacPath represents the /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac YANG schema element.
type Lacp_Interface_SystemIdMacPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_SystemIdMacPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/state/system-id-mac YANG schema element.
type Lacp_Interface_SystemIdMacPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/system-id-mac"
//	Path from root:       "/lacp/interfaces/interface/state/system-id-mac"
func (n *Lacp_Interface_SystemIdMacPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewSingletonQuery[string](
		"Lacp_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-id-mac"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lacp_Interface).SystemIdMac
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/system-id-mac"
//	Path from root:       "/lacp/interfaces/interface/state/system-id-mac"
func (n *Lacp_Interface_SystemIdMacPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lacp_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-id-mac"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lacp_Interface).SystemIdMac
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/system-id-mac"
//	Path from root:       "/lacp/interfaces/interface/config/system-id-mac"
func (n *Lacp_Interface_SystemIdMacPath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewConfigQuery[string](
		"Lacp_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-id-mac"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lacp_Interface).SystemIdMac
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/system-id-mac"
//	Path from root:       "/lacp/interfaces/interface/config/system-id-mac"
func (n *Lacp_Interface_SystemIdMacPathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lacp_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-id-mac"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lacp_Interface).SystemIdMac
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Lacp_Interface_SystemPriorityPath represents the /openconfig-lacp/lacp/interfaces/interface/state/system-priority YANG schema element.
type Lacp_Interface_SystemPriorityPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_SystemPriorityPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/state/system-priority YANG schema element.
type Lacp_Interface_SystemPriorityPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/system-priority"
//	Path from root:       "/lacp/interfaces/interface/state/system-priority"
func (n *Lacp_Interface_SystemPriorityPath) State() ygnmi.SingletonQuery[uint16] {
	return ygnmi.NewSingletonQuery[uint16](
		"Lacp_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-priority"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface).SystemPriority
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/system-priority"
//	Path from root:       "/lacp/interfaces/interface/state/system-priority"
func (n *Lacp_Interface_SystemPriorityPathAny) State() ygnmi.WildcardQuery[uint16] {
	return ygnmi.NewWildcardQuery[uint16](
		"Lacp_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-priority"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface).SystemPriority
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/system-priority"
//	Path from root:       "/lacp/interfaces/interface/config/system-priority"
func (n *Lacp_Interface_SystemPriorityPath) Config() ygnmi.ConfigQuery[uint16] {
	return ygnmi.NewConfigQuery[uint16](
		"Lacp_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-priority"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface).SystemPriority
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/system-priority"
//	Path from root:       "/lacp/interfaces/interface/config/system-priority"
func (n *Lacp_Interface_SystemPriorityPathAny) Config() ygnmi.WildcardQuery[uint16] {
	return ygnmi.NewWildcardQuery[uint16](
		"Lacp_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-priority"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface).SystemPriority
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Lacp_InterfacePath represents the /openconfig-lacp/lacp/interfaces/interface YANG schema element.
type Lacp_InterfacePath struct {
	*ygnmi.NodePath
}

// Lacp_InterfacePathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface YANG schema element.
type Lacp_InterfacePathAny struct {
	*ygnmi.NodePath
}

// Lacp_InterfacePathMap represents the /openconfig-lacp/lacp/interfaces/interface YANG schema element.
type Lacp_InterfacePathMap struct {
	*ygnmi.NodePath
}

// Lacp_InterfacePathMapAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface YANG schema element.
type Lacp_InterfacePathMapAny struct {
	*ygnmi.NodePath
}

// Fallback (leaf): If the fallback is set to true, current LACP interface is
// able to establish a Link Aggregation (LAG) before it receives
// LACP PDUs from its peer, and fallback to a single port active
// after the expiry of the timeout period.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/fallback"
//	Path from root:       "/lacp/interfaces/interface/*/fallback"
func (n *Lacp_InterfacePath) Fallback() *Lacp_Interface_FallbackPath {
	ps := &Lacp_Interface_FallbackPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "fallback"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Fallback (leaf): If the fallback is set to true, current LACP interface is
// able to establish a Link Aggregation (LAG) before it receives
// LACP PDUs from its peer, and fallback to a single port active
// after the expiry of the timeout period.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/fallback"
//	Path from root:       "/lacp/interfaces/interface/*/fallback"
func (n *Lacp_InterfacePathAny) Fallback() *Lacp_Interface_FallbackPathAny {
	ps := &Lacp_Interface_FallbackPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "fallback"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Interval (leaf): Set the period between LACP messages -- uses
// the lacp-period-type enumeration.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/interval"
//	Path from root:       "/lacp/interfaces/interface/*/interval"
func (n *Lacp_InterfacePath) Interval() *Lacp_Interface_IntervalPath {
	ps := &Lacp_Interface_IntervalPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "interval"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Interval (leaf): Set the period between LACP messages -- uses
// the lacp-period-type enumeration.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/interval"
//	Path from root:       "/lacp/interfaces/interface/*/interval"
func (n *Lacp_InterfacePathAny) Interval() *Lacp_Interface_IntervalPathAny {
	ps := &Lacp_Interface_IntervalPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "interval"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpMode (leaf): ACTIVE is to initiate the transmission of LACP packets.
// PASSIVE is to wait for peer to initiate the transmission of
// LACP packets.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/lacp-mode"
//	Path from root:       "/lacp/interfaces/interface/*/lacp-mode"
func (n *Lacp_InterfacePath) LacpMode() *Lacp_Interface_LacpModePath {
	ps := &Lacp_Interface_LacpModePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "lacp-mode"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpMode (leaf): ACTIVE is to initiate the transmission of LACP packets.
// PASSIVE is to wait for peer to initiate the transmission of
// LACP packets.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/lacp-mode"
//	Path from root:       "/lacp/interfaces/interface/*/lacp-mode"
func (n *Lacp_InterfacePathAny) LacpMode() *Lacp_Interface_LacpModePathAny {
	ps := &Lacp_Interface_LacpModePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "lacp-mode"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// MemberAny (list): List of member interfaces and their associated status for
// a LACP-controlled aggregate interface
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "members/member"
//	Path from root:       "/lacp/interfaces/interface/members/member"
func (n *Lacp_InterfacePath) MemberAny() *Lacp_Interface_MemberPathAny {
	ps := &Lacp_Interface_MemberPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"members", "member"},
			map[string]interface{}{"interface": "*"},
			n,
		),
	}
	return ps
}

// MemberAny (list): List of member interfaces and their associated status for
// a LACP-controlled aggregate interface
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "members/member"
//	Path from root:       "/lacp/interfaces/interface/members/member"
func (n *Lacp_InterfacePathAny) MemberAny() *Lacp_Interface_MemberPathAny {
	ps := &Lacp_Interface_MemberPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"members", "member"},
			map[string]interface{}{"interface": "*"},
			n,
		),
	}
	return ps
}

// Member (list): List of member interfaces and their associated status for
// a LACP-controlled aggregate interface
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "members/member"
//	Path from root:       "/lacp/interfaces/interface/members/member"
//
//	Interface: string
func (n *Lacp_InterfacePath) Member(Interface string) *Lacp_Interface_MemberPath {
	ps := &Lacp_Interface_MemberPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"members", "member"},
			map[string]interface{}{"interface": Interface},
			n,
		),
	}
	return ps
}

// Member (list): List of member interfaces and their associated status for
// a LACP-controlled aggregate interface
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "members/member"
//	Path from root:       "/lacp/interfaces/interface/members/member"
//
//	Interface: string
func (n *Lacp_InterfacePathAny) Member(Interface string) *Lacp_Interface_MemberPathAny {
	ps := &Lacp_Interface_MemberPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"members", "member"},
			map[string]interface{}{"interface": Interface},
			n,
		),
	}
	return ps
}

// MemberMap (list): List of member interfaces and their associated status for
// a LACP-controlled aggregate interface
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "members/member"
//	Path from root:       "/lacp/interfaces/interface/members/member"
func (n *Lacp_InterfacePath) MemberMap() *Lacp_Interface_MemberPathMap {
	ps := &Lacp_Interface_MemberPathMap{
		NodePath: ygnmi.NewNodePath(
			[]string{"members"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// MemberMap (list): List of member interfaces and their associated status for
// a LACP-controlled aggregate interface
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "members/member"
//	Path from root:       "/lacp/interfaces/interface/members/member"
func (n *Lacp_InterfacePathAny) MemberMap() *Lacp_Interface_MemberPathMapAny {
	ps := &Lacp_Interface_MemberPathMapAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"members"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// Name (leaf): Reference to the interface on which LACP should be
// configured.   The type of the target interface must be
// ieee8023adLag
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/name"
//	Path from root:       "/lacp/interfaces/interface/*/name"
func (n *Lacp_InterfacePath) Name() *Lacp_Interface_NamePath {
	ps := &Lacp_Interface_NamePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "name"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Name (leaf): Reference to the interface on which LACP should be
// configured.   The type of the target interface must be
// ieee8023adLag
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/name"
//	Path from root:       "/lacp/interfaces/interface/*/name"
func (n *Lacp_InterfacePathAny) Name() *Lacp_Interface_NamePathAny {
	ps := &Lacp_Interface_NamePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "name"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// SystemIdMac (leaf): The MAC address portion of the node's System ID. This is
// combined with the system priority to construct the 8-octet
// system-id
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/system-id-mac"
//	Path from root:       "/lacp/interfaces/interface/*/system-id-mac"
func (n *Lacp_InterfacePath) SystemIdMac() *Lacp_Interface_SystemIdMacPath {
	ps := &Lacp_Interface_SystemIdMacPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "system-id-mac"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// SystemIdMac (leaf): The MAC address portion of the node's System ID. This is
// combined with the system priority to construct the 8-octet
// system-id
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/system-id-mac"
//	Path from root:       "/lacp/interfaces/interface/*/system-id-mac"
func (n *Lacp_InterfacePathAny) SystemIdMac() *Lacp_Interface_SystemIdMacPathAny {
	ps := &Lacp_Interface_SystemIdMacPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "system-id-mac"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// SystemPriority (leaf): Sytem priority used by the node on this LAG interface.
// Lower value is higher priority for determining which node
// is the controlling system.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/system-priority"
//	Path from root:       "/lacp/interfaces/interface/*/system-priority"
func (n *Lacp_InterfacePath) SystemPriority() *Lacp_Interface_SystemPriorityPath {
	ps := &Lacp_Interface_SystemPriorityPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "system-priority"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// SystemPriority (leaf): Sytem priority used by the node on this LAG interface.
// Lower value is higher priority for determining which node
// is the controlling system.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/system-priority"
//	Path from root:       "/lacp/interfaces/interface/*/system-priority"
func (n *Lacp_InterfacePathAny) SystemPriority() *Lacp_Interface_SystemPriorityPathAny {
	ps := &Lacp_Interface_SystemPriorityPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "system-priority"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// State returns a Query that can be used in gNMI operations.
func (n *Lacp_InterfacePath) State() ygnmi.SingletonQuery[*oc.Lacp_Interface] {
	return ygnmi.NewSingletonQuery[*oc.Lacp_Interface](
		"Lacp_Interface",
		true,
		false,
		false,
		false,
		true,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *Lacp_InterfacePathAny) State() ygnmi.WildcardQuery[*oc.Lacp_Interface] {
	return ygnmi.NewWildcardQuery[*oc.Lacp_Interface](
		"Lacp_Interface",
		true,
		false,
		false,
		false,
		true,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Lacp_InterfacePath) Config() ygnmi.ConfigQuery[*oc.Lacp_Interface] {
	return ygnmi.NewConfigQuery[*oc.Lacp_Interface](
		"Lacp_Interface",
		false,
		true,
		false,
		false,
		true,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Lacp_InterfacePathAny) Config() ygnmi.WildcardQuery[*oc.Lacp_Interface] {
	return ygnmi.NewWildcardQuery[*oc.Lacp_Interface](
		"Lacp_Interface",
		false,
		true,
		false,
		false,
		true,
		false,
		n,
		nil,
		nil,
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *Lacp_InterfacePathMap) State() ygnmi.SingletonQuery[map[string]*oc.Lacp_Interface] {
	return ygnmi.NewSingletonQuery[map[string]*oc.Lacp_Interface](
		"Lacp",
		true,
		false,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.Lacp_Interface, bool) {
			ret := gs.(*oc.Lacp).Interface
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-lacp:interfaces"},
			PostRelPath: []string{"openconfig-lacp:interface"},
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *Lacp_InterfacePathMapAny) State() ygnmi.WildcardQuery[map[string]*oc.Lacp_Interface] {
	return ygnmi.NewWildcardQuery[map[string]*oc.Lacp_Interface](
		"Lacp",
		true,
		false,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.Lacp_Interface, bool) {
			ret := gs.(*oc.Lacp).Interface
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-lacp:interfaces"},
			PostRelPath: []string{"openconfig-lacp:interface"},
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Lacp_InterfacePathMap) Config() ygnmi.ConfigQuery[map[string]*oc.Lacp_Interface] {
	return ygnmi.NewConfigQuery[map[string]*oc.Lacp_Interface](
		"Lacp",
		false,
		true,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.Lacp_Interface, bool) {
			ret := gs.(*oc.Lacp).Interface
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-lacp:interfaces"},
			PostRelPath: []string{"openconfig-lacp:interface"},
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Lacp_InterfacePathMapAny) Config() ygnmi.WildcardQuery[map[string]*oc.Lacp_Interface] {
	return ygnmi.NewWildcardQuery[map[string]*oc.Lacp_Interface](
		"Lacp",
		false,
		true,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.Lacp_Interface, bool) {
			ret := gs.(*oc.Lacp).Interface
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-lacp:interfaces"},
			PostRelPath: []string{"openconfig-lacp:interface"},
		},
	)
}

// Lacp_Interface_Member_ActivityPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity YANG schema element.
type Lacp_Interface_Member_ActivityPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_ActivityPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity YANG schema element.
type Lacp_Interface_Member_ActivityPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/activity"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/activity"
func (n *Lacp_Interface_Member_ActivityPath) State() ygnmi.SingletonQuery[oc.E_Lacp_LacpActivityType] {
	return ygnmi.NewSingletonQuery[oc.E_Lacp_LacpActivityType](
		"Lacp_Interface_Member",
		true,
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "activity"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpActivityType, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).Activity
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/activity"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/activity"
func (n *Lacp_Interface_Member_ActivityPathAny) State() ygnmi.WildcardQuery[oc.E_Lacp_LacpActivityType] {
	return ygnmi.NewWildcardQuery[oc.E_Lacp_LacpActivityType](
		"Lacp_Interface_Member",
		true,
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "activity"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpActivityType, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).Activity
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}
