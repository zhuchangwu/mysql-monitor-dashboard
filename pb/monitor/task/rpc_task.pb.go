// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_task.proto

package task

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// 监控项的类型
type Type int32

const (
	Type_Sys   Type = 0
	Type_App   Type = 1
	Type_MySQL Type = 2
)

var Type_name = map[int32]string{
	0: "Sys",
	1: "App",
	2: "MySQL",
}

var Type_value = map[string]int32{
	"Sys":   0,
	"App":   1,
	"MySQL": 2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7aaef3364e691aab, []int{0}
}

// 定义消息，消息为service中发送或者是接受的内容
type RpcSysTask struct {
	ItemType             Type      `protobuf:"varint,1,opt,name=itemType,proto3,enum=proto.Type" json:"itemType,omitempty"`
	Sysmsg               *SysMsg   `protobuf:"bytes,2,opt,name=sysmsg,proto3" json:"sysmsg,omitempty"`
	Appmsg               *AppMsg   `protobuf:"bytes,3,opt,name=appmsg,proto3" json:"appmsg,omitempty"`
	Mysqlmsg             *MySQLMsg `protobuf:"bytes,4,opt,name=mysqlmsg,proto3" json:"mysqlmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RpcSysTask) Reset()         { *m = RpcSysTask{} }
func (m *RpcSysTask) String() string { return proto.CompactTextString(m) }
func (*RpcSysTask) ProtoMessage()    {}
func (*RpcSysTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aaef3364e691aab, []int{0}
}

func (m *RpcSysTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcSysTask.Unmarshal(m, b)
}
func (m *RpcSysTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcSysTask.Marshal(b, m, deterministic)
}
func (m *RpcSysTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcSysTask.Merge(m, src)
}
func (m *RpcSysTask) XXX_Size() int {
	return xxx_messageInfo_RpcSysTask.Size(m)
}
func (m *RpcSysTask) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcSysTask.DiscardUnknown(m)
}

var xxx_messageInfo_RpcSysTask proto.InternalMessageInfo

func (m *RpcSysTask) GetItemType() Type {
	if m != nil {
		return m.ItemType
	}
	return Type_Sys
}

func (m *RpcSysTask) GetSysmsg() *SysMsg {
	if m != nil {
		return m.Sysmsg
	}
	return nil
}

func (m *RpcSysTask) GetAppmsg() *AppMsg {
	if m != nil {
		return m.Appmsg
	}
	return nil
}

func (m *RpcSysTask) GetMysqlmsg() *MySQLMsg {
	if m != nil {
		return m.Mysqlmsg
	}
	return nil
}

// 操作系统监控项相关
type SysMsg struct {
	ItemName             string   `protobuf:"bytes,1,opt,name=itemName,proto3" json:"itemName,omitempty"`
	Cycle                string   `protobuf:"bytes,2,opt,name=cycle,proto3" json:"cycle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SysMsg) Reset()         { *m = SysMsg{} }
func (m *SysMsg) String() string { return proto.CompactTextString(m) }
func (*SysMsg) ProtoMessage()    {}
func (*SysMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aaef3364e691aab, []int{1}
}

func (m *SysMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SysMsg.Unmarshal(m, b)
}
func (m *SysMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SysMsg.Marshal(b, m, deterministic)
}
func (m *SysMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SysMsg.Merge(m, src)
}
func (m *SysMsg) XXX_Size() int {
	return xxx_messageInfo_SysMsg.Size(m)
}
func (m *SysMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SysMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SysMsg proto.InternalMessageInfo

func (m *SysMsg) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *SysMsg) GetCycle() string {
	if m != nil {
		return m.Cycle
	}
	return ""
}

// 应用监控项相关
type AppMsg struct {
	ItemName             string   `protobuf:"bytes,1,opt,name=itemName,proto3" json:"itemName,omitempty"`
	Cycle                string   `protobuf:"bytes,2,opt,name=cycle,proto3" json:"cycle,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	IsNew                bool     `protobuf:"varint,4,opt,name=isNew,proto3" json:"isNew,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppMsg) Reset()         { *m = AppMsg{} }
func (m *AppMsg) String() string { return proto.CompactTextString(m) }
func (*AppMsg) ProtoMessage()    {}
func (*AppMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aaef3364e691aab, []int{2}
}

func (m *AppMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppMsg.Unmarshal(m, b)
}
func (m *AppMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppMsg.Marshal(b, m, deterministic)
}
func (m *AppMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppMsg.Merge(m, src)
}
func (m *AppMsg) XXX_Size() int {
	return xxx_messageInfo_AppMsg.Size(m)
}
func (m *AppMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_AppMsg.DiscardUnknown(m)
}

var xxx_messageInfo_AppMsg proto.InternalMessageInfo

func (m *AppMsg) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *AppMsg) GetCycle() string {
	if m != nil {
		return m.Cycle
	}
	return ""
}

func (m *AppMsg) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *AppMsg) GetIsNew() bool {
	if m != nil {
		return m.IsNew
	}
	return false
}

// mysql监控相关
type MySQLMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLMsg) Reset()         { *m = MySQLMsg{} }
func (m *MySQLMsg) String() string { return proto.CompactTextString(m) }
func (*MySQLMsg) ProtoMessage()    {}
func (*MySQLMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aaef3364e691aab, []int{3}
}

func (m *MySQLMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MySQLMsg.Unmarshal(m, b)
}
func (m *MySQLMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MySQLMsg.Marshal(b, m, deterministic)
}
func (m *MySQLMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySQLMsg.Merge(m, src)
}
func (m *MySQLMsg) XXX_Size() int {
	return xxx_messageInfo_MySQLMsg.Size(m)
}
func (m *MySQLMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MySQLMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MySQLMsg proto.InternalMessageInfo

// 响应信息
type Response struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ResponseMsg          string   `protobuf:"bytes,2,opt,name=responseMsg,proto3" json:"responseMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aaef3364e691aab, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetResponseMsg() string {
	if m != nil {
		return m.ResponseMsg
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.Type", Type_name, Type_value)
	proto.RegisterType((*RpcSysTask)(nil), "proto.RpcSysTask")
	proto.RegisterType((*SysMsg)(nil), "proto.SysMsg")
	proto.RegisterType((*AppMsg)(nil), "proto.AppMsg")
	proto.RegisterType((*MySQLMsg)(nil), "proto.MySQLMsg")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() { proto.RegisterFile("rpc_task.proto", fileDescriptor_7aaef3364e691aab) }

var fileDescriptor_7aaef3364e691aab = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xd7, 0x6d, 0xad, 0xed, 0x5b, 0x9d, 0x33, 0xc8, 0x18, 0x3b, 0x8d, 0x8a, 0x38, 0x14,
	0x26, 0xcc, 0x9b, 0x07, 0x61, 0x22, 0x9e, 0xec, 0xc0, 0x74, 0x27, 0x2f, 0x52, 0xbb, 0x30, 0xca,
	0xd6, 0x25, 0xf6, 0x8d, 0x8e, 0x7c, 0x2b, 0x3f, 0xa2, 0xe4, 0xcf, 0xa6, 0x1e, 0x3d, 0x35, 0x4f,
	0xde, 0x5f, 0x78, 0x9e, 0x3e, 0x09, 0x74, 0x6a, 0x51, 0xbc, 0xca, 0x1c, 0x57, 0x63, 0x51, 0x73,
	0xc9, 0x89, 0x6f, 0x3e, 0xc9, 0x97, 0x07, 0x40, 0x45, 0x91, 0x29, 0x9c, 0xe7, 0xb8, 0x22, 0x17,
	0x10, 0x96, 0x92, 0x55, 0x73, 0x25, 0x58, 0xdf, 0x1b, 0x7a, 0xa3, 0xce, 0x24, 0xb6, 0xfc, 0x58,
	0x6f, 0xd1, 0xfd, 0x90, 0x9c, 0x43, 0x80, 0x0a, 0x2b, 0x5c, 0xf6, 0x9b, 0x43, 0x6f, 0x14, 0x4f,
	0x8e, 0x1c, 0x96, 0x29, 0x4c, 0x71, 0x49, 0xdd, 0x50, 0x63, 0xb9, 0x10, 0x1a, 0x6b, 0xfd, 0xc1,
	0xa6, 0x42, 0x18, 0xcc, 0x0e, 0xc9, 0x15, 0x84, 0x95, 0xc2, 0xf7, 0xb5, 0x06, 0xdb, 0x06, 0x3c,
	0x76, 0x60, 0xaa, 0xb2, 0xe7, 0x27, 0x8d, 0xee, 0x81, 0xe4, 0x16, 0x02, 0xeb, 0x42, 0x06, 0x36,
	0xed, 0x2c, 0xaf, 0x6c, 0xda, 0x88, 0xee, 0x35, 0x39, 0x05, 0xbf, 0x50, 0xc5, 0x9a, 0x99, 0x7c,
	0x11, 0xb5, 0x22, 0x59, 0x40, 0x60, 0xad, 0xff, 0x7f, 0x96, 0x10, 0x68, 0x0b, 0x5e, 0x4b, 0xf3,
	0x27, 0x3e, 0x35, 0x6b, 0x4d, 0x96, 0x38, 0x63, 0x5b, 0x93, 0x3a, 0xa4, 0x56, 0x24, 0x00, 0xe1,
	0x2e, 0x77, 0xf2, 0x00, 0x21, 0x65, 0x28, 0xf8, 0x06, 0x19, 0xe9, 0x41, 0x80, 0x32, 0x97, 0x1f,
	0x68, 0x1c, 0x7d, 0xea, 0x14, 0x19, 0x42, 0x5c, 0x3b, 0x26, 0x75, 0x8d, 0x46, 0xf4, 0xf7, 0xd6,
	0xe5, 0x19, 0xb4, 0x4d, 0xed, 0x07, 0xd0, 0xca, 0x14, 0x76, 0x1b, 0x7a, 0x31, 0x15, 0xa2, 0xeb,
	0x91, 0x08, 0x7c, 0xe3, 0xd5, 0x6d, 0x4e, 0x52, 0x88, 0x1f, 0xd7, 0x7c, 0x9b, 0xb1, 0xfa, 0xb3,
	0x2c, 0x18, 0xb9, 0x83, 0x5e, 0xc6, 0x36, 0x0b, 0x7d, 0xaf, 0x73, 0x9e, 0xea, 0xf6, 0x52, 0xbe,
	0x29, 0x25, 0xaf, 0xc9, 0x89, 0x2b, 0xf7, 0xe7, 0xe2, 0x07, 0xbb, 0xbe, 0x77, 0x59, 0x93, 0xc6,
	0x7d, 0xe7, 0xe5, 0xb0, 0xb2, 0x07, 0xae, 0xf5, 0xbb, 0x79, 0x0b, 0x0c, 0x71, 0xf3, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xb3, 0xd6, 0xab, 0xee, 0x4a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FlowServiceClient is the client API for FlowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlowServiceClient interface {
	SendTaskToMysqlMonitor(ctx context.Context, in *RpcSysTask, opts ...grpc.CallOption) (*Response, error)
}

type flowServiceClient struct {
	cc *grpc.ClientConn
}

func NewFlowServiceClient(cc *grpc.ClientConn) FlowServiceClient {
	return &flowServiceClient{cc}
}

func (c *flowServiceClient) SendTaskToMysqlMonitor(ctx context.Context, in *RpcSysTask, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.FlowService/SendTaskToMysqlMonitor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlowServiceServer is the server API for FlowService service.
type FlowServiceServer interface {
	SendTaskToMysqlMonitor(context.Context, *RpcSysTask) (*Response, error)
}

func RegisterFlowServiceServer(s *grpc.Server, srv FlowServiceServer) {
	s.RegisterService(&_FlowService_serviceDesc, srv)
}

func _FlowService_SendTaskToMysqlMonitor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcSysTask)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).SendTaskToMysqlMonitor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FlowService/SendTaskToMysqlMonitor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).SendTaskToMysqlMonitor(ctx, req.(*RpcSysTask))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FlowService",
	HandlerType: (*FlowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTaskToMysqlMonitor",
			Handler:    _FlowService_SendTaskToMysqlMonitor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_task.proto",
}
