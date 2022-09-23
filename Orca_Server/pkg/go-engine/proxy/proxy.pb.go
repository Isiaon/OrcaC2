// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy.proto

package proxy

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PROXY_PROTO int32

const (
	PROXY_PROTO_TCP   PROXY_PROTO = 0
	PROXY_PROTO_UDP   PROXY_PROTO = 1
	PROXY_PROTO_RUDP  PROXY_PROTO = 2
	PROXY_PROTO_RICMP PROXY_PROTO = 3
	PROXY_PROTO_KCP   PROXY_PROTO = 4
)

var PROXY_PROTO_name = map[int32]string{
	0: "TCP",
	1: "UDP",
	2: "RUDP",
	3: "RICMP",
	4: "KCP",
}

var PROXY_PROTO_value = map[string]int32{
	"TCP":   0,
	"UDP":   1,
	"RUDP":  2,
	"RICMP": 3,
	"KCP":   4,
}

func (x PROXY_PROTO) String() string {
	return proto.EnumName(PROXY_PROTO_name, int32(x))
}

func (PROXY_PROTO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{0}
}

type CLIENT_TYPE int32

const (
	// client fromaddr -> server toaddr
	CLIENT_TYPE_PROXY CLIENT_TYPE = 0
	// server fromaddr -> client toaddr
	CLIENT_TYPE_REVERSE_PROXY CLIENT_TYPE = 1
	// client fromaddr -> server
	CLIENT_TYPE_SOCKS5 CLIENT_TYPE = 2
	// server fromaddr -> client
	CLIENT_TYPE_REVERSE_SOCKS5 CLIENT_TYPE = 3
	// client fromaddr -> server shadowsocks address [SS_LOCAL_HOST:SS_LOCAL_PORT]
	CLIENT_TYPE_SS_PROXY CLIENT_TYPE = 4
)

var CLIENT_TYPE_name = map[int32]string{
	0: "PROXY",
	1: "REVERSE_PROXY",
	2: "SOCKS5",
	3: "REVERSE_SOCKS5",
	4: "SS_PROXY",
}

var CLIENT_TYPE_value = map[string]int32{
	"PROXY":          0,
	"REVERSE_PROXY":  1,
	"SOCKS5":         2,
	"REVERSE_SOCKS5": 3,
	"SS_PROXY":       4,
}

func (x CLIENT_TYPE) String() string {
	return proto.EnumName(CLIENT_TYPE_name, int32(x))
}

func (CLIENT_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{1}
}

type FRAME_TYPE int32

const (
	FRAME_TYPE_LOGIN    FRAME_TYPE = 0
	FRAME_TYPE_LOGINRSP FRAME_TYPE = 1
	FRAME_TYPE_DATA     FRAME_TYPE = 2
	FRAME_TYPE_PING     FRAME_TYPE = 3
	FRAME_TYPE_PONG     FRAME_TYPE = 4
	FRAME_TYPE_OPEN     FRAME_TYPE = 5
	FRAME_TYPE_OPENRSP  FRAME_TYPE = 6
	FRAME_TYPE_CLOSE    FRAME_TYPE = 7
)

var FRAME_TYPE_name = map[int32]string{
	0: "LOGIN",
	1: "LOGINRSP",
	2: "DATA",
	3: "PING",
	4: "PONG",
	5: "OPEN",
	6: "OPENRSP",
	7: "CLOSE",
}

var FRAME_TYPE_value = map[string]int32{
	"LOGIN":    0,
	"LOGINRSP": 1,
	"DATA":     2,
	"PING":     3,
	"PONG":     4,
	"OPEN":     5,
	"OPENRSP":  6,
	"CLOSE":    7,
}

func (x FRAME_TYPE) String() string {
	return proto.EnumName(FRAME_TYPE_name, int32(x))
}

func (FRAME_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{2}
}

type LoginFrame struct {
	Proxyproto           PROXY_PROTO `protobuf:"varint,1,opt,name=proxyproto,proto3,enum=PROXY_PROTO" json:"proxyproto,omitempty"`
	Clienttype           CLIENT_TYPE `protobuf:"varint,2,opt,name=clienttype,proto3,enum=CLIENT_TYPE" json:"clienttype,omitempty"`
	Fromaddr             string      `protobuf:"bytes,3,opt,name=fromaddr,proto3" json:"fromaddr,omitempty"`
	Toaddr               string      `protobuf:"bytes,4,opt,name=toaddr,proto3" json:"toaddr,omitempty"`
	Name                 string      `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Key                  string      `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LoginFrame) Reset()         { *m = LoginFrame{} }
func (m *LoginFrame) String() string { return proto.CompactTextString(m) }
func (*LoginFrame) ProtoMessage()    {}
func (*LoginFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{0}
}

func (m *LoginFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginFrame.Unmarshal(m, b)
}
func (m *LoginFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginFrame.Marshal(b, m, deterministic)
}
func (m *LoginFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginFrame.Merge(m, src)
}
func (m *LoginFrame) XXX_Size() int {
	return xxx_messageInfo_LoginFrame.Size(m)
}
func (m *LoginFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginFrame.DiscardUnknown(m)
}

var xxx_messageInfo_LoginFrame proto.InternalMessageInfo

func (m *LoginFrame) GetProxyproto() PROXY_PROTO {
	if m != nil {
		return m.Proxyproto
	}
	return PROXY_PROTO_TCP
}

func (m *LoginFrame) GetClienttype() CLIENT_TYPE {
	if m != nil {
		return m.Clienttype
	}
	return CLIENT_TYPE_PROXY
}

func (m *LoginFrame) GetFromaddr() string {
	if m != nil {
		return m.Fromaddr
	}
	return ""
}

func (m *LoginFrame) GetToaddr() string {
	if m != nil {
		return m.Toaddr
	}
	return ""
}

func (m *LoginFrame) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginFrame) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type LoginRspFrame struct {
	Ret                  bool     `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRspFrame) Reset()         { *m = LoginRspFrame{} }
func (m *LoginRspFrame) String() string { return proto.CompactTextString(m) }
func (*LoginRspFrame) ProtoMessage()    {}
func (*LoginRspFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{1}
}

func (m *LoginRspFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRspFrame.Unmarshal(m, b)
}
func (m *LoginRspFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRspFrame.Marshal(b, m, deterministic)
}
func (m *LoginRspFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRspFrame.Merge(m, src)
}
func (m *LoginRspFrame) XXX_Size() int {
	return xxx_messageInfo_LoginRspFrame.Size(m)
}
func (m *LoginRspFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRspFrame.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRspFrame proto.InternalMessageInfo

func (m *LoginRspFrame) GetRet() bool {
	if m != nil {
		return m.Ret
	}
	return false
}

func (m *LoginRspFrame) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PingFrame struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingFrame) Reset()         { *m = PingFrame{} }
func (m *PingFrame) String() string { return proto.CompactTextString(m) }
func (*PingFrame) ProtoMessage()    {}
func (*PingFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{2}
}

func (m *PingFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingFrame.Unmarshal(m, b)
}
func (m *PingFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingFrame.Marshal(b, m, deterministic)
}
func (m *PingFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingFrame.Merge(m, src)
}
func (m *PingFrame) XXX_Size() int {
	return xxx_messageInfo_PingFrame.Size(m)
}
func (m *PingFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_PingFrame.DiscardUnknown(m)
}

var xxx_messageInfo_PingFrame proto.InternalMessageInfo

func (m *PingFrame) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type PongFrame struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongFrame) Reset()         { *m = PongFrame{} }
func (m *PongFrame) String() string { return proto.CompactTextString(m) }
func (*PongFrame) ProtoMessage()    {}
func (*PongFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{3}
}

func (m *PongFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongFrame.Unmarshal(m, b)
}
func (m *PongFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongFrame.Marshal(b, m, deterministic)
}
func (m *PongFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongFrame.Merge(m, src)
}
func (m *PongFrame) XXX_Size() int {
	return xxx_messageInfo_PongFrame.Size(m)
}
func (m *PongFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_PongFrame.DiscardUnknown(m)
}

var xxx_messageInfo_PongFrame proto.InternalMessageInfo

func (m *PongFrame) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type OpenConnFrame struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Toaddr               string   `protobuf:"bytes,2,opt,name=toaddr,proto3" json:"toaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenConnFrame) Reset()         { *m = OpenConnFrame{} }
func (m *OpenConnFrame) String() string { return proto.CompactTextString(m) }
func (*OpenConnFrame) ProtoMessage()    {}
func (*OpenConnFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{4}
}

func (m *OpenConnFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenConnFrame.Unmarshal(m, b)
}
func (m *OpenConnFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenConnFrame.Marshal(b, m, deterministic)
}
func (m *OpenConnFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenConnFrame.Merge(m, src)
}
func (m *OpenConnFrame) XXX_Size() int {
	return xxx_messageInfo_OpenConnFrame.Size(m)
}
func (m *OpenConnFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenConnFrame.DiscardUnknown(m)
}

var xxx_messageInfo_OpenConnFrame proto.InternalMessageInfo

func (m *OpenConnFrame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OpenConnFrame) GetToaddr() string {
	if m != nil {
		return m.Toaddr
	}
	return ""
}

type OpenConnRspFrame struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ret                  bool     `protobuf:"varint,2,opt,name=ret,proto3" json:"ret,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenConnRspFrame) Reset()         { *m = OpenConnRspFrame{} }
func (m *OpenConnRspFrame) String() string { return proto.CompactTextString(m) }
func (*OpenConnRspFrame) ProtoMessage()    {}
func (*OpenConnRspFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{5}
}

func (m *OpenConnRspFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenConnRspFrame.Unmarshal(m, b)
}
func (m *OpenConnRspFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenConnRspFrame.Marshal(b, m, deterministic)
}
func (m *OpenConnRspFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenConnRspFrame.Merge(m, src)
}
func (m *OpenConnRspFrame) XXX_Size() int {
	return xxx_messageInfo_OpenConnRspFrame.Size(m)
}
func (m *OpenConnRspFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenConnRspFrame.DiscardUnknown(m)
}

var xxx_messageInfo_OpenConnRspFrame proto.InternalMessageInfo

func (m *OpenConnRspFrame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OpenConnRspFrame) GetRet() bool {
	if m != nil {
		return m.Ret
	}
	return false
}

func (m *OpenConnRspFrame) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type CloseFrame struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseFrame) Reset()         { *m = CloseFrame{} }
func (m *CloseFrame) String() string { return proto.CompactTextString(m) }
func (*CloseFrame) ProtoMessage()    {}
func (*CloseFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{6}
}

func (m *CloseFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseFrame.Unmarshal(m, b)
}
func (m *CloseFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseFrame.Marshal(b, m, deterministic)
}
func (m *CloseFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseFrame.Merge(m, src)
}
func (m *CloseFrame) XXX_Size() int {
	return xxx_messageInfo_CloseFrame.Size(m)
}
func (m *CloseFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseFrame.DiscardUnknown(m)
}

var xxx_messageInfo_CloseFrame proto.InternalMessageInfo

func (m *CloseFrame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DataFrame struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Compress             bool     `protobuf:"varint,2,opt,name=compress,proto3" json:"compress,omitempty"`
	Crc                  string   `protobuf:"bytes,3,opt,name=crc,proto3" json:"crc,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Index                int32    `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataFrame) Reset()         { *m = DataFrame{} }
func (m *DataFrame) String() string { return proto.CompactTextString(m) }
func (*DataFrame) ProtoMessage()    {}
func (*DataFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{7}
}

func (m *DataFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataFrame.Unmarshal(m, b)
}
func (m *DataFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataFrame.Marshal(b, m, deterministic)
}
func (m *DataFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataFrame.Merge(m, src)
}
func (m *DataFrame) XXX_Size() int {
	return xxx_messageInfo_DataFrame.Size(m)
}
func (m *DataFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_DataFrame.DiscardUnknown(m)
}

var xxx_messageInfo_DataFrame proto.InternalMessageInfo

func (m *DataFrame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataFrame) GetCompress() bool {
	if m != nil {
		return m.Compress
	}
	return false
}

func (m *DataFrame) GetCrc() string {
	if m != nil {
		return m.Crc
	}
	return ""
}

func (m *DataFrame) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DataFrame) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ProxyFrame struct {
	Type                 FRAME_TYPE        `protobuf:"varint,1,opt,name=type,proto3,enum=FRAME_TYPE" json:"type,omitempty"`
	LoginFrame           *LoginFrame       `protobuf:"bytes,2,opt,name=loginFrame,proto3" json:"loginFrame,omitempty"`
	LoginRspFrame        *LoginRspFrame    `protobuf:"bytes,3,opt,name=loginRspFrame,proto3" json:"loginRspFrame,omitempty"`
	DataFrame            *DataFrame        `protobuf:"bytes,4,opt,name=dataFrame,proto3" json:"dataFrame,omitempty"`
	PingFrame            *PingFrame        `protobuf:"bytes,5,opt,name=pingFrame,proto3" json:"pingFrame,omitempty"`
	PongFrame            *PongFrame        `protobuf:"bytes,6,opt,name=pongFrame,proto3" json:"pongFrame,omitempty"`
	OpenFrame            *OpenConnFrame    `protobuf:"bytes,7,opt,name=openFrame,proto3" json:"openFrame,omitempty"`
	OpenRspFrame         *OpenConnRspFrame `protobuf:"bytes,8,opt,name=openRspFrame,proto3" json:"openRspFrame,omitempty"`
	CloseFrame           *CloseFrame       `protobuf:"bytes,9,opt,name=closeFrame,proto3" json:"closeFrame,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProxyFrame) Reset()         { *m = ProxyFrame{} }
func (m *ProxyFrame) String() string { return proto.CompactTextString(m) }
func (*ProxyFrame) ProtoMessage()    {}
func (*ProxyFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{8}
}

func (m *ProxyFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyFrame.Unmarshal(m, b)
}
func (m *ProxyFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyFrame.Marshal(b, m, deterministic)
}
func (m *ProxyFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyFrame.Merge(m, src)
}
func (m *ProxyFrame) XXX_Size() int {
	return xxx_messageInfo_ProxyFrame.Size(m)
}
func (m *ProxyFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyFrame.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyFrame proto.InternalMessageInfo

func (m *ProxyFrame) GetType() FRAME_TYPE {
	if m != nil {
		return m.Type
	}
	return FRAME_TYPE_LOGIN
}

func (m *ProxyFrame) GetLoginFrame() *LoginFrame {
	if m != nil {
		return m.LoginFrame
	}
	return nil
}

func (m *ProxyFrame) GetLoginRspFrame() *LoginRspFrame {
	if m != nil {
		return m.LoginRspFrame
	}
	return nil
}

func (m *ProxyFrame) GetDataFrame() *DataFrame {
	if m != nil {
		return m.DataFrame
	}
	return nil
}

func (m *ProxyFrame) GetPingFrame() *PingFrame {
	if m != nil {
		return m.PingFrame
	}
	return nil
}

func (m *ProxyFrame) GetPongFrame() *PongFrame {
	if m != nil {
		return m.PongFrame
	}
	return nil
}

func (m *ProxyFrame) GetOpenFrame() *OpenConnFrame {
	if m != nil {
		return m.OpenFrame
	}
	return nil
}

func (m *ProxyFrame) GetOpenRspFrame() *OpenConnRspFrame {
	if m != nil {
		return m.OpenRspFrame
	}
	return nil
}

func (m *ProxyFrame) GetCloseFrame() *CloseFrame {
	if m != nil {
		return m.CloseFrame
	}
	return nil
}

func init() {
	proto.RegisterEnum("PROXY_PROTO", PROXY_PROTO_name, PROXY_PROTO_value)
	proto.RegisterEnum("CLIENT_TYPE", CLIENT_TYPE_name, CLIENT_TYPE_value)
	proto.RegisterEnum("FRAME_TYPE", FRAME_TYPE_name, FRAME_TYPE_value)
	proto.RegisterType((*LoginFrame)(nil), "LoginFrame")
	proto.RegisterType((*LoginRspFrame)(nil), "LoginRspFrame")
	proto.RegisterType((*PingFrame)(nil), "PingFrame")
	proto.RegisterType((*PongFrame)(nil), "PongFrame")
	proto.RegisterType((*OpenConnFrame)(nil), "OpenConnFrame")
	proto.RegisterType((*OpenConnRspFrame)(nil), "OpenConnRspFrame")
	proto.RegisterType((*CloseFrame)(nil), "CloseFrame")
	proto.RegisterType((*DataFrame)(nil), "DataFrame")
	proto.RegisterType((*ProxyFrame)(nil), "ProxyFrame")
}

func init() { proto.RegisterFile("proxy.proto", fileDescriptor_700b50b08ed8dbaf) }

var fileDescriptor_700b50b08ed8dbaf = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x6f, 0x9b, 0x40,
	0x10, 0x8d, 0x01, 0x7f, 0x30, 0xd8, 0xd6, 0x66, 0x55, 0x55, 0x56, 0x54, 0x29, 0x91, 0x4f, 0x51,
	0x1a, 0x71, 0x70, 0x1a, 0xf5, 0xd4, 0x43, 0x4a, 0x48, 0x64, 0xc5, 0x31, 0x68, 0x70, 0xab, 0xa6,
	0x97, 0x88, 0xc2, 0x36, 0x42, 0xb5, 0x01, 0x61, 0x0e, 0xf1, 0xbf, 0xe9, 0x6f, 0xe9, 0x2f, 0xab,
	0x76, 0xf8, 0x74, 0x1a, 0xf5, 0xf6, 0x76, 0xdf, 0x9b, 0xdd, 0x79, 0xb3, 0x0f, 0xc0, 0x48, 0xb3,
	0xe4, 0x79, 0x67, 0xa6, 0x59, 0x92, 0x27, 0xd3, 0x3f, 0x1d, 0x80, 0x45, 0xf2, 0x14, 0xc5, 0x37,
	0x99, 0xbf, 0x11, 0xfc, 0x1c, 0x80, 0x58, 0x22, 0x27, 0x9d, 0x93, 0xce, 0xe9, 0x78, 0x36, 0x34,
	0x5d, 0x74, 0xbe, 0x3d, 0x3c, 0xba, 0xe8, 0xac, 0x1c, 0x6c, 0xf1, 0x52, 0x1d, 0xac, 0x23, 0x11,
	0xe7, 0xf9, 0x2e, 0x15, 0x13, 0xa5, 0x54, 0x5b, 0x8b, 0xb9, 0xbd, 0x5c, 0x3d, 0xae, 0x1e, 0x5c,
	0x1b, 0x5b, 0x3c, 0x3f, 0x82, 0xc1, 0xcf, 0x2c, 0xd9, 0xf8, 0x61, 0x98, 0x4d, 0xd4, 0x93, 0xce,
	0xa9, 0x8e, 0xf5, 0x9a, 0xbf, 0x85, 0x5e, 0x9e, 0x10, 0xa3, 0x11, 0x53, 0xae, 0x38, 0x07, 0x2d,
	0xf6, 0x37, 0x62, 0xd2, 0xa5, 0x5d, 0xc2, 0x9c, 0x81, 0xfa, 0x4b, 0xec, 0x26, 0x3d, 0xda, 0x92,
	0x70, 0x7a, 0x01, 0x23, 0xf2, 0x80, 0xdb, 0xb4, 0xb0, 0xc1, 0x40, 0xcd, 0x44, 0x4e, 0xfd, 0x0f,
	0x50, 0x42, 0xb9, 0xb3, 0xd9, 0x3e, 0x51, 0x8f, 0x3a, 0x4a, 0x38, 0x3d, 0x06, 0xdd, 0x8d, 0xe2,
	0xa7, 0xa2, 0x80, 0x83, 0x96, 0x47, 0x1b, 0x41, 0x15, 0x2a, 0x12, 0x26, 0x41, 0xf2, 0x3f, 0xc1,
	0x47, 0x18, 0x39, 0xa9, 0x88, 0xad, 0x24, 0x2e, 0xa7, 0x37, 0x06, 0x25, 0x0a, 0x49, 0xa2, 0xa3,
	0x12, 0x85, 0x2d, 0x57, 0x4a, 0xdb, 0xd5, 0xf4, 0x06, 0x58, 0x55, 0x58, 0xb7, 0xfc, 0xb2, 0xb6,
	0xb4, 0xa0, 0xfc, 0x63, 0x41, 0x6d, 0x2c, 0xbc, 0x03, 0xb0, 0xd6, 0xc9, 0x56, 0xbc, 0x7a, 0xc2,
	0x74, 0x0b, 0xfa, 0xb5, 0x9f, 0xfb, 0xaf, 0x1f, 0x7f, 0x04, 0x83, 0x20, 0xd9, 0xa4, 0x99, 0xd8,
	0x6e, 0xcb, 0x3b, 0xea, 0xb5, 0xbc, 0x28, 0xc8, 0x82, 0xea, 0xa2, 0x20, 0x0b, 0xa4, 0xfb, 0xd0,
	0xcf, 0x7d, 0x7a, 0x9c, 0x21, 0x12, 0xe6, 0x6f, 0xa0, 0x1b, 0xc5, 0xa1, 0x78, 0xa6, 0xb7, 0xe9,
	0x62, 0xb1, 0x98, 0xfe, 0x56, 0x01, 0x5c, 0x99, 0x90, 0xe2, 0xda, 0x63, 0xd0, 0x28, 0x1b, 0x45,
	0x92, 0x0c, 0xf3, 0x06, 0xaf, 0xee, 0xed, 0x22, 0x1a, 0x44, 0xf0, 0xf7, 0x00, 0xeb, 0x3a, 0x7e,
	0xd4, 0x89, 0x31, 0x33, 0xcc, 0x26, 0x91, 0xd8, 0xa2, 0xf9, 0x07, 0x18, 0xad, 0xdb, 0xef, 0x4c,
	0x2d, 0x1a, 0xb3, 0xb1, 0xb9, 0xf7, 0xfa, 0xb8, 0x2f, 0xe2, 0xa7, 0xa0, 0x87, 0xd5, 0x1c, 0xc8,
	0x81, 0x31, 0x03, 0xb3, 0x9e, 0x0c, 0x36, 0xa4, 0x54, 0xa6, 0x55, 0x24, 0xc8, 0x96, 0x54, 0xd6,
	0x21, 0xc1, 0x86, 0x24, 0x65, 0x95, 0x0d, 0x4a, 0x22, 0x29, 0x93, 0x46, 0x59, 0x07, 0xe7, 0x1c,
	0xf4, 0x24, 0x15, 0xa5, 0xbf, 0x7e, 0xd9, 0xef, 0x5e, 0x6c, 0xb0, 0x11, 0xf0, 0x4b, 0x18, 0xca,
	0x45, 0x6d, 0x70, 0x40, 0x05, 0x87, 0xe6, 0xcb, 0xb8, 0xe0, 0x9e, 0x4c, 0x4e, 0x31, 0xa8, 0x83,
	0x30, 0xd1, 0xcb, 0x29, 0x36, 0xd9, 0xc0, 0x16, 0x7d, 0xf6, 0x09, 0x8c, 0xd6, 0x07, 0xcd, 0xfb,
	0xa0, 0xae, 0x2c, 0x97, 0x1d, 0x48, 0xf0, 0xe5, 0xda, 0x65, 0x1d, 0x3e, 0x00, 0x0d, 0x25, 0x52,
	0xb8, 0x0e, 0x5d, 0x9c, 0x5b, 0xf7, 0x2e, 0x53, 0x25, 0x7b, 0x67, 0xb9, 0x4c, 0x3b, 0x7b, 0x00,
	0xa3, 0xf5, 0x85, 0x4b, 0x09, 0x9d, 0xc6, 0x0e, 0xf8, 0x21, 0x8c, 0xd0, 0xfe, 0x6a, 0xa3, 0x67,
	0x3f, 0x16, 0x5b, 0x1d, 0x0e, 0xd0, 0xf3, 0x1c, 0xeb, 0xce, 0xbb, 0x64, 0x0a, 0xe7, 0x30, 0xae,
	0xe8, 0x72, 0x4f, 0xe5, 0x43, 0x18, 0x78, 0x5e, 0xa9, 0xd6, 0xce, 0x04, 0x40, 0x13, 0x10, 0x79,
	0xf2, 0xc2, 0xb9, 0x9d, 0x2f, 0xd9, 0x81, 0x94, 0x11, 0x44, 0xaf, 0xec, 0xef, 0xfa, 0x6a, 0x75,
	0xc5, 0x14, 0x89, 0xdc, 0xf9, 0xf2, 0x96, 0xa9, 0x84, 0x9c, 0xe5, 0x2d, 0xd3, 0x24, 0x72, 0x5c,
	0x7b, 0xc9, 0xba, 0xdc, 0x80, 0xbe, 0x44, 0xb2, 0xa8, 0x27, 0x4f, 0xb3, 0x16, 0x8e, 0x67, 0xb3,
	0xfe, 0xe7, 0xfe, 0xf7, 0x2e, 0xfd, 0xc4, 0x7e, 0xf4, 0xe8, 0x37, 0x76, 0xf1, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xd7, 0x9c, 0x13, 0xae, 0x12, 0x05, 0x00, 0x00,
}
