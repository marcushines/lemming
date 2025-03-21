/*
Package sampling is a generated package which contains definitions
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
package sampling

import (
	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Sampling_Sflow_CollectorPath represents the /openconfig-sampling/sampling/sflow/collectors/collector YANG schema element.
type Sampling_Sflow_CollectorPath struct {
	*ygnmi.NodePath
}

// Sampling_Sflow_CollectorPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/collectors/collector YANG schema element.
type Sampling_Sflow_CollectorPathAny struct {
	*ygnmi.NodePath
}

// Sampling_Sflow_CollectorPathMap represents the /openconfig-sampling/sampling/sflow/collectors/collector YANG schema element.
type Sampling_Sflow_CollectorPathMap struct {
	*ygnmi.NodePath
}

// Sampling_Sflow_CollectorPathMapAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/collectors/collector YANG schema element.
type Sampling_Sflow_CollectorPathMapAny struct {
	*ygnmi.NodePath
}

// Address (leaf): IPv4/IPv6 address of the sFlow collector.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "*/address"
//	Path from root:       "/sampling/sflow/collectors/collector/*/address"
func (n *Sampling_Sflow_CollectorPath) Address() *Sampling_Sflow_Collector_AddressPath {
	ps := &Sampling_Sflow_Collector_AddressPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Address (leaf): IPv4/IPv6 address of the sFlow collector.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "*/address"
//	Path from root:       "/sampling/sflow/collectors/collector/*/address"
func (n *Sampling_Sflow_CollectorPathAny) Address() *Sampling_Sflow_Collector_AddressPathAny {
	ps := &Sampling_Sflow_Collector_AddressPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// NetworkInstance (leaf): Reference to the network instance used to reach the
// sFlow collector.  If uspecified, the collector destination
// is reachable in the default network instance.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "*/network-instance"
//	Path from root:       "/sampling/sflow/collectors/collector/*/network-instance"
func (n *Sampling_Sflow_CollectorPath) NetworkInstance() *Sampling_Sflow_Collector_NetworkInstancePath {
	ps := &Sampling_Sflow_Collector_NetworkInstancePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "network-instance"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// NetworkInstance (leaf): Reference to the network instance used to reach the
// sFlow collector.  If uspecified, the collector destination
// is reachable in the default network instance.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "*/network-instance"
//	Path from root:       "/sampling/sflow/collectors/collector/*/network-instance"
func (n *Sampling_Sflow_CollectorPathAny) NetworkInstance() *Sampling_Sflow_Collector_NetworkInstancePathAny {
	ps := &Sampling_Sflow_Collector_NetworkInstancePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "network-instance"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// PacketsSent (leaf): The total number of packets sampled and sent to the
// collector.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "state/packets-sent"
//	Path from root:       "/sampling/sflow/collectors/collector/state/packets-sent"
func (n *Sampling_Sflow_CollectorPath) PacketsSent() *Sampling_Sflow_Collector_PacketsSentPath {
	ps := &Sampling_Sflow_Collector_PacketsSentPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"state", "packets-sent"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// PacketsSent (leaf): The total number of packets sampled and sent to the
// collector.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "state/packets-sent"
//	Path from root:       "/sampling/sflow/collectors/collector/state/packets-sent"
func (n *Sampling_Sflow_CollectorPathAny) PacketsSent() *Sampling_Sflow_Collector_PacketsSentPathAny {
	ps := &Sampling_Sflow_Collector_PacketsSentPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"state", "packets-sent"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Port (leaf): UDP port number for the sFlow collector.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "*/port"
//	Path from root:       "/sampling/sflow/collectors/collector/*/port"
func (n *Sampling_Sflow_CollectorPath) Port() *Sampling_Sflow_Collector_PortPath {
	ps := &Sampling_Sflow_Collector_PortPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "port"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Port (leaf): UDP port number for the sFlow collector.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "*/port"
//	Path from root:       "/sampling/sflow/collectors/collector/*/port"
func (n *Sampling_Sflow_CollectorPathAny) Port() *Sampling_Sflow_Collector_PortPathAny {
	ps := &Sampling_Sflow_Collector_PortPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "port"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// SourceAddress (leaf): Sets the source IPv4/IPv6 address for sFlow datagrams sent
// to sFlow collectors.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "*/source-address"
//	Path from root:       "/sampling/sflow/collectors/collector/*/source-address"
func (n *Sampling_Sflow_CollectorPath) SourceAddress() *Sampling_Sflow_Collector_SourceAddressPath {
	ps := &Sampling_Sflow_Collector_SourceAddressPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "source-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// SourceAddress (leaf): Sets the source IPv4/IPv6 address for sFlow datagrams sent
// to sFlow collectors.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling"
//	Path from parent:     "*/source-address"
//	Path from root:       "/sampling/sflow/collectors/collector/*/source-address"
func (n *Sampling_Sflow_CollectorPathAny) SourceAddress() *Sampling_Sflow_Collector_SourceAddressPathAny {
	ps := &Sampling_Sflow_Collector_SourceAddressPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "source-address"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// State returns a Query that can be used in gNMI operations.
func (n *Sampling_Sflow_CollectorPath) State() ygnmi.SingletonQuery[*oc.Sampling_Sflow_Collector] {
	return ygnmi.NewSingletonQuery[*oc.Sampling_Sflow_Collector](
		"Sampling_Sflow_Collector",
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
func (n *Sampling_Sflow_CollectorPathAny) State() ygnmi.WildcardQuery[*oc.Sampling_Sflow_Collector] {
	return ygnmi.NewWildcardQuery[*oc.Sampling_Sflow_Collector](
		"Sampling_Sflow_Collector",
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
func (n *Sampling_Sflow_CollectorPath) Config() ygnmi.ConfigQuery[*oc.Sampling_Sflow_Collector] {
	return ygnmi.NewConfigQuery[*oc.Sampling_Sflow_Collector](
		"Sampling_Sflow_Collector",
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
func (n *Sampling_Sflow_CollectorPathAny) Config() ygnmi.WildcardQuery[*oc.Sampling_Sflow_Collector] {
	return ygnmi.NewWildcardQuery[*oc.Sampling_Sflow_Collector](
		"Sampling_Sflow_Collector",
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
func (n *Sampling_Sflow_CollectorPathMap) State() ygnmi.SingletonQuery[map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector] {
	return ygnmi.NewSingletonQuery[map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector](
		"Sampling_Sflow",
		true,
		false,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector, bool) {
			ret := gs.(*oc.Sampling_Sflow).Collector
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-sampling-sflow:collectors"},
			PostRelPath: []string{"openconfig-sampling-sflow:collector"},
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *Sampling_Sflow_CollectorPathMapAny) State() ygnmi.WildcardQuery[map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector] {
	return ygnmi.NewWildcardQuery[map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector](
		"Sampling_Sflow",
		true,
		false,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector, bool) {
			ret := gs.(*oc.Sampling_Sflow).Collector
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-sampling-sflow:collectors"},
			PostRelPath: []string{"openconfig-sampling-sflow:collector"},
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Sampling_Sflow_CollectorPathMap) Config() ygnmi.ConfigQuery[map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector] {
	return ygnmi.NewConfigQuery[map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector](
		"Sampling_Sflow",
		false,
		true,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector, bool) {
			ret := gs.(*oc.Sampling_Sflow).Collector
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-sampling-sflow:collectors"},
			PostRelPath: []string{"openconfig-sampling-sflow:collector"},
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *Sampling_Sflow_CollectorPathMapAny) Config() ygnmi.WildcardQuery[map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector] {
	return ygnmi.NewWildcardQuery[map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector](
		"Sampling_Sflow",
		false,
		true,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[oc.Sampling_Sflow_Collector_Key]*oc.Sampling_Sflow_Collector, bool) {
			ret := gs.(*oc.Sampling_Sflow).Collector
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-sampling-sflow:collectors"},
			PostRelPath: []string{"openconfig-sampling-sflow:collector"},
		},
	)
}

// Sampling_Sflow_Interface_EgressSamplingRatePath represents the /openconfig-sampling/sampling/sflow/interfaces/interface/state/egress-sampling-rate YANG schema element.
type Sampling_Sflow_Interface_EgressSamplingRatePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Sampling_Sflow_Interface_EgressSamplingRatePathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/interfaces/interface/state/egress-sampling-rate YANG schema element.
type Sampling_Sflow_Interface_EgressSamplingRatePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling-sflow"
//	Path from parent:     "state/egress-sampling-rate"
//	Path from root:       "/sampling/sflow/interfaces/interface/state/egress-sampling-rate"
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) State() ygnmi.SingletonQuery[uint32] {
	return ygnmi.NewSingletonQuery[uint32](
		"Sampling_Sflow_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "egress-sampling-rate"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint32, bool) {
			ret := gs.(*oc.Sampling_Sflow_Interface).EgressSamplingRate
			if ret == nil {
				var zero uint32
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow_Interface) },
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
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling-sflow"
//	Path from parent:     "state/egress-sampling-rate"
//	Path from root:       "/sampling/sflow/interfaces/interface/state/egress-sampling-rate"
func (n *Sampling_Sflow_Interface_EgressSamplingRatePathAny) State() ygnmi.WildcardQuery[uint32] {
	return ygnmi.NewWildcardQuery[uint32](
		"Sampling_Sflow_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "egress-sampling-rate"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint32, bool) {
			ret := gs.(*oc.Sampling_Sflow_Interface).EgressSamplingRate
			if ret == nil {
				var zero uint32
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow_Interface) },
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
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling-sflow"
//	Path from parent:     "config/egress-sampling-rate"
//	Path from root:       "/sampling/sflow/interfaces/interface/config/egress-sampling-rate"
func (n *Sampling_Sflow_Interface_EgressSamplingRatePath) Config() ygnmi.ConfigQuery[uint32] {
	return ygnmi.NewConfigQuery[uint32](
		"Sampling_Sflow_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "egress-sampling-rate"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint32, bool) {
			ret := gs.(*oc.Sampling_Sflow_Interface).EgressSamplingRate
			if ret == nil {
				var zero uint32
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow_Interface) },
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
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling-sflow"
//	Path from parent:     "config/egress-sampling-rate"
//	Path from root:       "/sampling/sflow/interfaces/interface/config/egress-sampling-rate"
func (n *Sampling_Sflow_Interface_EgressSamplingRatePathAny) Config() ygnmi.WildcardQuery[uint32] {
	return ygnmi.NewWildcardQuery[uint32](
		"Sampling_Sflow_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "egress-sampling-rate"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint32, bool) {
			ret := gs.(*oc.Sampling_Sflow_Interface).EgressSamplingRate
			if ret == nil {
				var zero uint32
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow_Interface) },
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

// Sampling_Sflow_Interface_EnabledPath represents the /openconfig-sampling/sampling/sflow/interfaces/interface/state/enabled YANG schema element.
type Sampling_Sflow_Interface_EnabledPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Sampling_Sflow_Interface_EnabledPathAny represents the wildcard version of the /openconfig-sampling/sampling/sflow/interfaces/interface/state/enabled YANG schema element.
type Sampling_Sflow_Interface_EnabledPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling-sflow"
//	Path from parent:     "state/enabled"
//	Path from root:       "/sampling/sflow/interfaces/interface/state/enabled"
func (n *Sampling_Sflow_Interface_EnabledPath) State() ygnmi.SingletonQuery[bool] {
	return ygnmi.NewSingletonQuery[bool](
		"Sampling_Sflow_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "enabled"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Sampling_Sflow_Interface).Enabled
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow_Interface) },
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
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling-sflow"
//	Path from parent:     "state/enabled"
//	Path from root:       "/sampling/sflow/interfaces/interface/state/enabled"
func (n *Sampling_Sflow_Interface_EnabledPathAny) State() ygnmi.WildcardQuery[bool] {
	return ygnmi.NewWildcardQuery[bool](
		"Sampling_Sflow_Interface",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "enabled"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Sampling_Sflow_Interface).Enabled
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow_Interface) },
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
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling-sflow"
//	Path from parent:     "config/enabled"
//	Path from root:       "/sampling/sflow/interfaces/interface/config/enabled"
func (n *Sampling_Sflow_Interface_EnabledPath) Config() ygnmi.ConfigQuery[bool] {
	return ygnmi.NewConfigQuery[bool](
		"Sampling_Sflow_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "enabled"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Sampling_Sflow_Interface).Enabled
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow_Interface) },
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
//	Defining module:      "openconfig-sampling-sflow"
//	Instantiating module: "openconfig-sampling-sflow"
//	Path from parent:     "config/enabled"
//	Path from root:       "/sampling/sflow/interfaces/interface/config/enabled"
func (n *Sampling_Sflow_Interface_EnabledPathAny) Config() ygnmi.WildcardQuery[bool] {
	return ygnmi.NewWildcardQuery[bool](
		"Sampling_Sflow_Interface",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "enabled"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Sampling_Sflow_Interface).Enabled
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Sampling_Sflow_Interface) },
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
