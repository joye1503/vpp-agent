// Code generated by protoc-gen-gogo.
// source: ipsec.proto
// DO NOT EDIT!

/*
Package ipsec is a generated protocol buffer package.

It is generated from these files:
	ipsec.proto

It has these top-level messages:
	SecurityPolicyDatabases
	SecurityAssociations
*/
package ipsec

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// Policy action
type SecurityPolicyDatabases_SPD_PolicyEntry_Action int32

const (
	SecurityPolicyDatabases_SPD_PolicyEntry_BYPASS  SecurityPolicyDatabases_SPD_PolicyEntry_Action = 0
	SecurityPolicyDatabases_SPD_PolicyEntry_DISCARD SecurityPolicyDatabases_SPD_PolicyEntry_Action = 1
	SecurityPolicyDatabases_SPD_PolicyEntry_RESOLVE SecurityPolicyDatabases_SPD_PolicyEntry_Action = 2
	SecurityPolicyDatabases_SPD_PolicyEntry_PROTECT SecurityPolicyDatabases_SPD_PolicyEntry_Action = 3
)

var SecurityPolicyDatabases_SPD_PolicyEntry_Action_name = map[int32]string{
	0: "BYPASS",
	1: "DISCARD",
	2: "RESOLVE",
	3: "PROTECT",
}
var SecurityPolicyDatabases_SPD_PolicyEntry_Action_value = map[string]int32{
	"BYPASS":  0,
	"DISCARD": 1,
	"RESOLVE": 2,
	"PROTECT": 3,
}

func (x SecurityPolicyDatabases_SPD_PolicyEntry_Action) String() string {
	return proto.EnumName(SecurityPolicyDatabases_SPD_PolicyEntry_Action_name, int32(x))
}

// IPSec protocol
type SecurityAssociations_SA_IPSecProtocol int32

const (
	SecurityAssociations_SA_AH  SecurityAssociations_SA_IPSecProtocol = 0
	SecurityAssociations_SA_ESP SecurityAssociations_SA_IPSecProtocol = 1
)

var SecurityAssociations_SA_IPSecProtocol_name = map[int32]string{
	0: "AH",
	1: "ESP",
}
var SecurityAssociations_SA_IPSecProtocol_value = map[string]int32{
	"AH":  0,
	"ESP": 1,
}

func (x SecurityAssociations_SA_IPSecProtocol) String() string {
	return proto.EnumName(SecurityAssociations_SA_IPSecProtocol_name, int32(x))
}

// Cryptographic algorithm for encryption
type SecurityAssociations_SA_CryptoAlgorithm int32

const (
	SecurityAssociations_SA_NONE_CRYPTO SecurityAssociations_SA_CryptoAlgorithm = 0
	SecurityAssociations_SA_AES_CBC_128 SecurityAssociations_SA_CryptoAlgorithm = 1
	SecurityAssociations_SA_AES_CBC_192 SecurityAssociations_SA_CryptoAlgorithm = 2
	SecurityAssociations_SA_AES_CBC_256 SecurityAssociations_SA_CryptoAlgorithm = 3
)

var SecurityAssociations_SA_CryptoAlgorithm_name = map[int32]string{
	0: "NONE_CRYPTO",
	1: "AES_CBC_128",
	2: "AES_CBC_192",
	3: "AES_CBC_256",
}
var SecurityAssociations_SA_CryptoAlgorithm_value = map[string]int32{
	"NONE_CRYPTO": 0,
	"AES_CBC_128": 1,
	"AES_CBC_192": 2,
	"AES_CBC_256": 3,
}

func (x SecurityAssociations_SA_CryptoAlgorithm) String() string {
	return proto.EnumName(SecurityAssociations_SA_CryptoAlgorithm_name, int32(x))
}

// Cryptographic algorithm for authentication
type SecurityAssociations_SA_IntegAlgorithm int32

const (
	SecurityAssociations_SA_NONE_INTEG  SecurityAssociations_SA_IntegAlgorithm = 0
	SecurityAssociations_SA_MD5_96      SecurityAssociations_SA_IntegAlgorithm = 1
	SecurityAssociations_SA_SHA1_96     SecurityAssociations_SA_IntegAlgorithm = 2
	SecurityAssociations_SA_SHA_256_96  SecurityAssociations_SA_IntegAlgorithm = 3
	SecurityAssociations_SA_SHA_256_128 SecurityAssociations_SA_IntegAlgorithm = 4
	SecurityAssociations_SA_SHA_384_192 SecurityAssociations_SA_IntegAlgorithm = 5
	SecurityAssociations_SA_SHA_512_256 SecurityAssociations_SA_IntegAlgorithm = 6
)

var SecurityAssociations_SA_IntegAlgorithm_name = map[int32]string{
	0: "NONE_INTEG",
	1: "MD5_96",
	2: "SHA1_96",
	3: "SHA_256_96",
	4: "SHA_256_128",
	5: "SHA_384_192",
	6: "SHA_512_256",
}
var SecurityAssociations_SA_IntegAlgorithm_value = map[string]int32{
	"NONE_INTEG":  0,
	"MD5_96":      1,
	"SHA1_96":     2,
	"SHA_256_96":  3,
	"SHA_256_128": 4,
	"SHA_384_192": 5,
	"SHA_512_256": 6,
}

func (x SecurityAssociations_SA_IntegAlgorithm) String() string {
	return proto.EnumName(SecurityAssociations_SA_IntegAlgorithm_name, int32(x))
}

// Security Policy Database (SPD)
type SecurityPolicyDatabases struct {
	Spds []*SecurityPolicyDatabases_SPD `protobuf:"bytes,1,rep,name=spds" json:"spds,omitempty"`
}

func (m *SecurityPolicyDatabases) Reset()         { *m = SecurityPolicyDatabases{} }
func (m *SecurityPolicyDatabases) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabases) ProtoMessage()    {}

func (m *SecurityPolicyDatabases) GetSpds() []*SecurityPolicyDatabases_SPD {
	if m != nil {
		return m.Spds
	}
	return nil
}

type SecurityPolicyDatabases_SPD struct {
	// uint32 id = 1;  /* SPD ID */
	Name          string                                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interfaces    []*SecurityPolicyDatabases_SPD_Interface   `protobuf:"bytes,2,rep,name=interfaces" json:"interfaces,omitempty"`
	PolicyEntries []*SecurityPolicyDatabases_SPD_PolicyEntry `protobuf:"bytes,3,rep,name=policy_entries" json:"policy_entries,omitempty"`
}

func (m *SecurityPolicyDatabases_SPD) Reset()         { *m = SecurityPolicyDatabases_SPD{} }
func (m *SecurityPolicyDatabases_SPD) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabases_SPD) ProtoMessage()    {}

func (m *SecurityPolicyDatabases_SPD) GetInterfaces() []*SecurityPolicyDatabases_SPD_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *SecurityPolicyDatabases_SPD) GetPolicyEntries() []*SecurityPolicyDatabases_SPD_PolicyEntry {
	if m != nil {
		return m.PolicyEntries
	}
	return nil
}

// Interface
type SecurityPolicyDatabases_SPD_Interface struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *SecurityPolicyDatabases_SPD_Interface) Reset()         { *m = SecurityPolicyDatabases_SPD_Interface{} }
func (m *SecurityPolicyDatabases_SPD_Interface) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabases_SPD_Interface) ProtoMessage()    {}

// Policy Entry
type SecurityPolicyDatabases_SPD_PolicyEntry struct {
	// uint32 spd_id = 1;
	Priority        int32                                          `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	IsOutbound      bool                                           `protobuf:"varint,2,opt,name=is_outbound,proto3" json:"is_outbound,omitempty"`
	IsIpv6          bool                                           `protobuf:"varint,3,opt,name=is_ipv6,proto3" json:"is_ipv6,omitempty"`
	RemoteAddrStart string                                         `protobuf:"bytes,4,opt,name=remote_addr_start,proto3" json:"remote_addr_start,omitempty"`
	RemoteAddrStop  string                                         `protobuf:"bytes,5,opt,name=remote_addr_stop,proto3" json:"remote_addr_stop,omitempty"`
	LocalAddrStart  string                                         `protobuf:"bytes,6,opt,name=local_addr_start,proto3" json:"local_addr_start,omitempty"`
	LocalAddrStop   string                                         `protobuf:"bytes,7,opt,name=local_addr_stop,proto3" json:"local_addr_stop,omitempty"`
	Protocol        bool                                           `protobuf:"varint,8,opt,name=protocol,proto3" json:"protocol,omitempty"`
	RemotePortStart uint32                                         `protobuf:"varint,9,opt,name=remote_port_start,proto3" json:"remote_port_start,omitempty"`
	RemotePortStop  uint32                                         `protobuf:"varint,10,opt,name=remote_port_stop,proto3" json:"remote_port_stop,omitempty"`
	LocalPortStart  uint32                                         `protobuf:"varint,11,opt,name=local_port_start,proto3" json:"local_port_start,omitempty"`
	LocalPortStop   uint32                                         `protobuf:"varint,12,opt,name=local_port_stop,proto3" json:"local_port_stop,omitempty"`
	Action          SecurityPolicyDatabases_SPD_PolicyEntry_Action `protobuf:"varint,13,opt,name=action,proto3,enum=ipsec.SecurityPolicyDatabases_SPD_PolicyEntry_Action" json:"action,omitempty"`
	// uint32 sa_id = 14;
	Sa string `protobuf:"bytes,14,opt,name=sa,proto3" json:"sa,omitempty"`
}

func (m *SecurityPolicyDatabases_SPD_PolicyEntry) Reset() {
	*m = SecurityPolicyDatabases_SPD_PolicyEntry{}
}
func (m *SecurityPolicyDatabases_SPD_PolicyEntry) String() string { return proto.CompactTextString(m) }
func (*SecurityPolicyDatabases_SPD_PolicyEntry) ProtoMessage()    {}

// Security Association (SA)
type SecurityAssociations struct {
	Sas []*SecurityAssociations_SA `protobuf:"bytes,1,rep,name=sas" json:"sas,omitempty"`
}

func (m *SecurityAssociations) Reset()         { *m = SecurityAssociations{} }
func (m *SecurityAssociations) String() string { return proto.CompactTextString(m) }
func (*SecurityAssociations) ProtoMessage()    {}

func (m *SecurityAssociations) GetSas() []*SecurityAssociations_SA {
	if m != nil {
		return m.Sas
	}
	return nil
}

type SecurityAssociations_SA struct {
	// uint32 id = 1; /* SA ID */
	Name          string                                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Spi           uint32                                  `protobuf:"varint,2,opt,name=spi,proto3" json:"spi,omitempty"`
	Protocol      SecurityAssociations_SA_IPSecProtocol   `protobuf:"varint,3,opt,name=protocol,proto3,enum=ipsec.SecurityAssociations_SA_IPSecProtocol" json:"protocol,omitempty"`
	CryptoAlg     SecurityAssociations_SA_CryptoAlgorithm `protobuf:"varint,4,opt,name=crypto_alg,proto3,enum=ipsec.SecurityAssociations_SA_CryptoAlgorithm" json:"crypto_alg,omitempty"`
	CryptoKey     string                                  `protobuf:"bytes,5,opt,name=crypto_key,proto3" json:"crypto_key,omitempty"`
	IntegAlg      SecurityAssociations_SA_IntegAlgorithm  `protobuf:"varint,6,opt,name=integ_alg,proto3,enum=ipsec.SecurityAssociations_SA_IntegAlgorithm" json:"integ_alg,omitempty"`
	IntegKey      string                                  `protobuf:"bytes,7,opt,name=integ_key,proto3" json:"integ_key,omitempty"`
	UseEsn        bool                                    `protobuf:"varint,8,opt,name=use_esn,proto3" json:"use_esn,omitempty"`
	UseAntiReplay bool                                    `protobuf:"varint,9,opt,name=use_anti_replay,proto3" json:"use_anti_replay,omitempty"`
	IsTunnel      bool                                    `protobuf:"varint,10,opt,name=is_tunnel,proto3" json:"is_tunnel,omitempty"`
	IsTunnelIpv6  bool                                    `protobuf:"varint,11,opt,name=is_tunnel_ipv6,proto3" json:"is_tunnel_ipv6,omitempty"`
	TunnelSrcAddr string                                  `protobuf:"bytes,12,opt,name=tunnel_src_addr,proto3" json:"tunnel_src_addr,omitempty"`
	TunnelDstAddr string                                  `protobuf:"bytes,13,opt,name=tunnel_dst_addr,proto3" json:"tunnel_dst_addr,omitempty"`
}

func (m *SecurityAssociations_SA) Reset()         { *m = SecurityAssociations_SA{} }
func (m *SecurityAssociations_SA) String() string { return proto.CompactTextString(m) }
func (*SecurityAssociations_SA) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("ipsec.SecurityPolicyDatabases_SPD_PolicyEntry_Action", SecurityPolicyDatabases_SPD_PolicyEntry_Action_name, SecurityPolicyDatabases_SPD_PolicyEntry_Action_value)
	proto.RegisterEnum("ipsec.SecurityAssociations_SA_IPSecProtocol", SecurityAssociations_SA_IPSecProtocol_name, SecurityAssociations_SA_IPSecProtocol_value)
	proto.RegisterEnum("ipsec.SecurityAssociations_SA_CryptoAlgorithm", SecurityAssociations_SA_CryptoAlgorithm_name, SecurityAssociations_SA_CryptoAlgorithm_value)
	proto.RegisterEnum("ipsec.SecurityAssociations_SA_IntegAlgorithm", SecurityAssociations_SA_IntegAlgorithm_name, SecurityAssociations_SA_IntegAlgorithm_value)
}
