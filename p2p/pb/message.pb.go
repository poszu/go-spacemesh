// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p/pb/message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// data common to all messages - Top level msg format
type CommonMessageData struct {
	SessionId            []byte   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonMessageData) Reset()         { *m = CommonMessageData{} }
func (m *CommonMessageData) String() string { return proto.CompactTextString(m) }
func (*CommonMessageData) ProtoMessage()    {}
func (*CommonMessageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_0bea1fdc4ee728e4, []int{0}
}
func (m *CommonMessageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonMessageData.Unmarshal(m, b)
}
func (m *CommonMessageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonMessageData.Marshal(b, m, deterministic)
}
func (dst *CommonMessageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonMessageData.Merge(dst, src)
}
func (m *CommonMessageData) XXX_Size() int {
	return xxx_messageInfo_CommonMessageData.Size(m)
}
func (m *CommonMessageData) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonMessageData.DiscardUnknown(m)
}

var xxx_messageInfo_CommonMessageData proto.InternalMessageInfo

func (m *CommonMessageData) GetSessionId() []byte {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *CommonMessageData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CommonMessageData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// Handshake protocol data used for both request and response - sent unencrypted over the wire
type HandshakeData struct {
	SessionId            []byte   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ClientVersion        string   `protobuf:"bytes,4,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	NetworkID            int32    `protobuf:"varint,5,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Protocol             string   `protobuf:"bytes,6,opt,name=protocol,proto3" json:"protocol,omitempty"`
	NodePubKey           []byte   `protobuf:"bytes,7,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`
	Iv                   []byte   `protobuf:"bytes,8,opt,name=iv,proto3" json:"iv,omitempty"`
	PubKey               []byte   `protobuf:"bytes,9,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Hmac                 []byte   `protobuf:"bytes,10,opt,name=hmac,proto3" json:"hmac,omitempty"`
	Sign                 string   `protobuf:"bytes,11,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeData) Reset()         { *m = HandshakeData{} }
func (m *HandshakeData) String() string { return proto.CompactTextString(m) }
func (*HandshakeData) ProtoMessage()    {}
func (*HandshakeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_0bea1fdc4ee728e4, []int{1}
}
func (m *HandshakeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeData.Unmarshal(m, b)
}
func (m *HandshakeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeData.Marshal(b, m, deterministic)
}
func (dst *HandshakeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeData.Merge(dst, src)
}
func (m *HandshakeData) XXX_Size() int {
	return xxx_messageInfo_HandshakeData.Size(m)
}
func (m *HandshakeData) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeData.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeData proto.InternalMessageInfo

func (m *HandshakeData) GetSessionId() []byte {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *HandshakeData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HandshakeData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *HandshakeData) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *HandshakeData) GetNetworkID() int32 {
	if m != nil {
		return m.NetworkID
	}
	return 0
}

func (m *HandshakeData) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *HandshakeData) GetNodePubKey() []byte {
	if m != nil {
		return m.NodePubKey
	}
	return nil
}

func (m *HandshakeData) GetIv() []byte {
	if m != nil {
		return m.Iv
	}
	return nil
}

func (m *HandshakeData) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *HandshakeData) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

func (m *HandshakeData) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

// used for protocol messages (non-handshake) - this is the decrypted CommonMessageData.payload
// it allows multi310.445plexing back to higher level protocols
// data is here and not in CommonMessageData to avoid leaked data on unencrypted connections
type ProtocolMessage struct {
	Metadata             *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Payload              []byte    `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ProtocolMessage) Reset()         { *m = ProtocolMessage{} }
func (m *ProtocolMessage) String() string { return proto.CompactTextString(m) }
func (*ProtocolMessage) ProtoMessage()    {}
func (*ProtocolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_0bea1fdc4ee728e4, []int{2}
}
func (m *ProtocolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolMessage.Unmarshal(m, b)
}
func (m *ProtocolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolMessage.Marshal(b, m, deterministic)
}
func (dst *ProtocolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolMessage.Merge(dst, src)
}
func (m *ProtocolMessage) XXX_Size() int {
	return xxx_messageInfo_ProtocolMessage.Size(m)
}
func (m *ProtocolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolMessage proto.InternalMessageInfo

func (m *ProtocolMessage) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ProtocolMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Metadata struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ClientVersion        string   `protobuf:"bytes,2,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Gossip               bool     `protobuf:"varint,4,opt,name=gossip,proto3" json:"gossip,omitempty"`
	AuthPubKey           []byte   `protobuf:"bytes,5,opt,name=authPubKey,proto3" json:"authPubKey,omitempty"`
	AuthorSign           string   `protobuf:"bytes,6,opt,name=authorSign,proto3" json:"authorSign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_0bea1fdc4ee728e4, []int{3}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Metadata) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *Metadata) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Metadata) GetGossip() bool {
	if m != nil {
		return m.Gossip
	}
	return false
}

func (m *Metadata) GetAuthPubKey() []byte {
	if m != nil {
		return m.AuthPubKey
	}
	return nil
}

func (m *Metadata) GetAuthorSign() string {
	if m != nil {
		return m.AuthorSign
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonMessageData)(nil), "pb.CommonMessageData")
	proto.RegisterType((*HandshakeData)(nil), "pb.HandshakeData")
	proto.RegisterType((*ProtocolMessage)(nil), "pb.ProtocolMessage")
	proto.RegisterType((*Metadata)(nil), "pb.Metadata")
}

func init() { proto.RegisterFile("p2p/pb/message.proto", fileDescriptor_message_0bea1fdc4ee728e4) }

var fileDescriptor_message_0bea1fdc4ee728e4 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xcd, 0x6e, 0xe2, 0x30,
	0x10, 0x80, 0x15, 0x03, 0x21, 0x0c, 0xb0, 0xab, 0xb5, 0x56, 0xc8, 0x5a, 0xad, 0xaa, 0x28, 0xea,
	0x21, 0x27, 0x90, 0xe8, 0x1b, 0xb4, 0x1c, 0x8a, 0x2a, 0x24, 0x94, 0xaa, 0x3d, 0xf4, 0xe6, 0x10,
	0x0b, 0x2c, 0xf0, 0x8f, 0x62, 0x43, 0xc5, 0x83, 0xf5, 0x05, 0xfa, 0x64, 0x55, 0x9c, 0x40, 0xa0,
	0x15, 0xbd, 0xf5, 0xe6, 0xf9, 0xc6, 0xe3, 0xf1, 0x7c, 0x03, 0x7f, 0xf5, 0x58, 0x8f, 0x74, 0x3a,
	0x12, 0xcc, 0x18, 0xba, 0x64, 0x43, 0x9d, 0x2b, 0xab, 0x30, 0xd2, 0x69, 0xc4, 0xe1, 0xcf, 0x9d,
	0x12, 0x42, 0xc9, 0x59, 0x99, 0x9a, 0x50, 0x4b, 0xf1, 0x7f, 0xe8, 0x18, 0x66, 0x0c, 0x57, 0x72,
	0x9a, 0x11, 0x2f, 0xf4, 0xe2, 0x5e, 0x52, 0x03, 0x4c, 0xa0, 0xad, 0xe9, 0x7e, 0xa3, 0x68, 0x46,
	0x90, 0xcb, 0x1d, 0xc2, 0xa2, 0xce, 0x72, 0xc1, 0x8c, 0xa5, 0x42, 0x93, 0x46, 0xe8, 0xc5, 0x8d,
	0xa4, 0x06, 0xd1, 0x1b, 0x82, 0xfe, 0x3d, 0x95, 0x99, 0x59, 0xd1, 0xf5, 0x0f, 0xf6, 0xc1, 0xd7,
	0xd0, 0x5f, 0x6c, 0x38, 0x93, 0xf6, 0x99, 0xe5, 0xc5, 0x53, 0xa4, 0x19, 0x7a, 0x71, 0x27, 0x39,
	0x87, 0xc5, 0x1b, 0x92, 0xd9, 0x57, 0x95, 0xaf, 0xa7, 0x13, 0xd2, 0x0a, 0xbd, 0xb8, 0x95, 0xd4,
	0x00, 0xff, 0x83, 0xc0, 0x39, 0x5a, 0xa8, 0x0d, 0xf1, 0x5d, 0xf9, 0x31, 0xc6, 0x57, 0x00, 0x52,
	0x65, 0x6c, 0xbe, 0x4d, 0x1f, 0xd8, 0x9e, 0xb4, 0xdd, 0xd7, 0x4e, 0x08, 0xfe, 0x05, 0x88, 0xef,
	0x48, 0xe0, 0x38, 0xe2, 0x3b, 0x3c, 0x00, 0x5f, 0x97, 0x77, 0x3b, 0x8e, 0x55, 0x11, 0xc6, 0xd0,
	0x5c, 0x09, 0xba, 0x20, 0xe0, 0xa8, 0x3b, 0x17, 0xcc, 0xf0, 0xa5, 0x24, 0x5d, 0xd7, 0xd3, 0x9d,
	0xa3, 0x27, 0xf8, 0x3d, 0xaf, 0x7a, 0x57, 0x4b, 0xc2, 0x31, 0x04, 0x82, 0x59, 0x9a, 0x51, 0x4b,
	0x9d, 0xb7, 0xee, 0xb8, 0x37, 0xd4, 0xe9, 0x70, 0x56, 0xb1, 0xe4, 0x98, 0xbd, 0x2c, 0x31, 0x7a,
	0xf7, 0x20, 0x38, 0x14, 0x9c, 0xcd, 0xeb, 0x7d, 0x9a, 0xf7, 0x8b, 0x4f, 0x74, 0xc1, 0xe7, 0x37,
	0x3b, 0x19, 0x80, 0xbf, 0x54, 0xc6, 0x70, 0xed, 0x96, 0x11, 0x24, 0x55, 0x54, 0xb8, 0xa4, 0x5b,
	0xbb, 0xaa, 0x5c, 0xb6, 0x4a, 0x97, 0x35, 0x39, 0xe4, 0x55, 0xfe, 0x58, 0x58, 0x29, 0x37, 0x71,
	0x42, 0x6e, 0x9b, 0x2f, 0x48, 0xa7, 0xa9, 0xef, 0xfe, 0x7a, 0xf3, 0x11, 0x00, 0x00, 0xff, 0xff,
	0x0c, 0xb2, 0xd6, 0x77, 0xe7, 0x02, 0x00, 0x00,
}
