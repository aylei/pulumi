// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language.proto

package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequiredPluginsRequest struct {
	Project              string   `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Program              string   `protobuf:"bytes,3,opt,name=program" json:"program,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequiredPluginsRequest) Reset()         { *m = GetRequiredPluginsRequest{} }
func (m *GetRequiredPluginsRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequiredPluginsRequest) ProtoMessage()    {}
func (*GetRequiredPluginsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_language_840dd930e81005a9, []int{0}
}
func (m *GetRequiredPluginsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequiredPluginsRequest.Unmarshal(m, b)
}
func (m *GetRequiredPluginsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequiredPluginsRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequiredPluginsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequiredPluginsRequest.Merge(dst, src)
}
func (m *GetRequiredPluginsRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequiredPluginsRequest.Size(m)
}
func (m *GetRequiredPluginsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequiredPluginsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequiredPluginsRequest proto.InternalMessageInfo

func (m *GetRequiredPluginsRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GetRequiredPluginsRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *GetRequiredPluginsRequest) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

type GetRequiredPluginsResponse struct {
	Plugins              []*PluginDependency `protobuf:"bytes,1,rep,name=plugins" json:"plugins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetRequiredPluginsResponse) Reset()         { *m = GetRequiredPluginsResponse{} }
func (m *GetRequiredPluginsResponse) String() string { return proto.CompactTextString(m) }
func (*GetRequiredPluginsResponse) ProtoMessage()    {}
func (*GetRequiredPluginsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_language_840dd930e81005a9, []int{1}
}
func (m *GetRequiredPluginsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequiredPluginsResponse.Unmarshal(m, b)
}
func (m *GetRequiredPluginsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequiredPluginsResponse.Marshal(b, m, deterministic)
}
func (dst *GetRequiredPluginsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequiredPluginsResponse.Merge(dst, src)
}
func (m *GetRequiredPluginsResponse) XXX_Size() int {
	return xxx_messageInfo_GetRequiredPluginsResponse.Size(m)
}
func (m *GetRequiredPluginsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequiredPluginsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequiredPluginsResponse proto.InternalMessageInfo

func (m *GetRequiredPluginsResponse) GetPlugins() []*PluginDependency {
	if m != nil {
		return m.Plugins
	}
	return nil
}

// RunRequest asks the interpreter to execute a program.
type RunRequest struct {
	Project              string            `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	Stack                string            `protobuf:"bytes,2,opt,name=stack" json:"stack,omitempty"`
	Pwd                  string            `protobuf:"bytes,3,opt,name=pwd" json:"pwd,omitempty"`
	Program              string            `protobuf:"bytes,4,opt,name=program" json:"program,omitempty"`
	Args                 []string          `protobuf:"bytes,5,rep,name=args" json:"args,omitempty"`
	Config               map[string]string `protobuf:"bytes,6,rep,name=config" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DryRun               bool              `protobuf:"varint,7,opt,name=dryRun" json:"dryRun,omitempty"`
	Parallel             int32             `protobuf:"varint,8,opt,name=parallel" json:"parallel,omitempty"`
	MonitorAddress       string            `protobuf:"bytes,9,opt,name=monitor_address,json=monitorAddress" json:"monitor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_language_840dd930e81005a9, []int{2}
}
func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (dst *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(dst, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

func (m *RunRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *RunRequest) GetStack() string {
	if m != nil {
		return m.Stack
	}
	return ""
}

func (m *RunRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *RunRequest) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

func (m *RunRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *RunRequest) GetConfig() map[string]string {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *RunRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *RunRequest) GetParallel() int32 {
	if m != nil {
		return m.Parallel
	}
	return 0
}

func (m *RunRequest) GetMonitorAddress() string {
	if m != nil {
		return m.MonitorAddress
	}
	return ""
}

// RunResponse is the response back from the interpreter/source back to the monitor.
type RunResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunResponse) Reset()         { *m = RunResponse{} }
func (m *RunResponse) String() string { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()    {}
func (*RunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_language_840dd930e81005a9, []int{3}
}
func (m *RunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunResponse.Unmarshal(m, b)
}
func (m *RunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunResponse.Marshal(b, m, deterministic)
}
func (dst *RunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunResponse.Merge(dst, src)
}
func (m *RunResponse) XXX_Size() int {
	return xxx_messageInfo_RunResponse.Size(m)
}
func (m *RunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunResponse proto.InternalMessageInfo

func (m *RunResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequiredPluginsRequest)(nil), "pulumirpc.GetRequiredPluginsRequest")
	proto.RegisterType((*GetRequiredPluginsResponse)(nil), "pulumirpc.GetRequiredPluginsResponse")
	proto.RegisterType((*RunRequest)(nil), "pulumirpc.RunRequest")
	proto.RegisterMapType((map[string]string)(nil), "pulumirpc.RunRequest.ConfigEntry")
	proto.RegisterType((*RunResponse)(nil), "pulumirpc.RunResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LanguageRuntime service

type LanguageRuntimeClient interface {
	// GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
	GetRequiredPlugins(ctx context.Context, in *GetRequiredPluginsRequest, opts ...grpc.CallOption) (*GetRequiredPluginsResponse, error)
	// Run executes a program and returns its result.
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PluginInfo, error)
}

type languageRuntimeClient struct {
	cc *grpc.ClientConn
}

func NewLanguageRuntimeClient(cc *grpc.ClientConn) LanguageRuntimeClient {
	return &languageRuntimeClient{cc}
}

func (c *languageRuntimeClient) GetRequiredPlugins(ctx context.Context, in *GetRequiredPluginsRequest, opts ...grpc.CallOption) (*GetRequiredPluginsResponse, error) {
	out := new(GetRequiredPluginsResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.LanguageRuntime/GetRequiredPlugins", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.LanguageRuntime/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) GetPluginInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := grpc.Invoke(ctx, "/pulumirpc.LanguageRuntime/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LanguageRuntime service

type LanguageRuntimeServer interface {
	// GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
	GetRequiredPlugins(context.Context, *GetRequiredPluginsRequest) (*GetRequiredPluginsResponse, error)
	// Run executes a program and returns its result.
	Run(context.Context, *RunRequest) (*RunResponse, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(context.Context, *empty.Empty) (*PluginInfo, error)
}

func RegisterLanguageRuntimeServer(s *grpc.Server, srv LanguageRuntimeServer) {
	s.RegisterService(&_LanguageRuntime_serviceDesc, srv)
}

func _LanguageRuntime_GetRequiredPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequiredPluginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GetRequiredPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.LanguageRuntime/GetRequiredPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GetRequiredPlugins(ctx, req.(*GetRequiredPluginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.LanguageRuntime/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.LanguageRuntime/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GetPluginInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LanguageRuntime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.LanguageRuntime",
	HandlerType: (*LanguageRuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRequiredPlugins",
			Handler:    _LanguageRuntime_GetRequiredPlugins_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _LanguageRuntime_Run_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _LanguageRuntime_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language.proto",
}

func init() { proto.RegisterFile("language.proto", fileDescriptor_language_840dd930e81005a9) }

var fileDescriptor_language_840dd930e81005a9 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x9b, 0xa6, 0xfb, 0x6f, 0x16, 0x5a, 0x64, 0xb5, 0x2b, 0x93, 0x5e, 0x42, 0x00, 0x91,
	0x53, 0x2a, 0x15, 0x81, 0x28, 0x27, 0x10, 0x54, 0x15, 0x12, 0x07, 0x64, 0x1e, 0x00, 0xb9, 0xc9,
	0x6c, 0x14, 0x9a, 0xd8, 0xc6, 0xb1, 0x41, 0x79, 0x4a, 0x5e, 0x85, 0x47, 0x40, 0xb1, 0x93, 0x65,
	0xa1, 0x8b, 0x7a, 0xf3, 0x37, 0xf9, 0x26, 0xfe, 0xf9, 0xb3, 0x07, 0x0e, 0x6b, 0x2e, 0x4a, 0xcb,
	0x4b, 0xcc, 0x94, 0x96, 0x46, 0x92, 0x85, 0xb2, 0xb5, 0x6d, 0x2a, 0xad, 0xf2, 0xe8, 0x9e, 0xaa,
	0x6d, 0x59, 0x09, 0xff, 0x21, 0x3a, 0x2d, 0xa5, 0x2c, 0x6b, 0x3c, 0x73, 0xea, 0xda, 0xae, 0xcf,
	0xb0, 0x51, 0xa6, 0xf3, 0x1f, 0x13, 0x0e, 0x0f, 0xaf, 0xd0, 0x30, 0xfc, 0x66, 0x2b, 0x8d, 0xc5,
	0x27, 0xd7, 0xd7, 0xf6, 0x12, 0x5b, 0x43, 0x28, 0xcc, 0x94, 0x96, 0x5f, 0x31, 0x37, 0x34, 0x88,
	0x83, 0x74, 0xc1, 0x46, 0x49, 0x1e, 0x40, 0xa8, 0x7e, 0x14, 0x74, 0xdf, 0x55, 0xfb, 0xe5, 0xe0,
	0x2d, 0x35, 0x6f, 0x68, 0xb8, 0xf1, 0xf6, 0x32, 0xf9, 0x0c, 0xd1, 0xae, 0x2d, 0x5a, 0x25, 0x45,
	0x8b, 0xe4, 0x05, 0xcc, 0x3c, 0x6d, 0x4b, 0x83, 0x38, 0x4c, 0x97, 0xe7, 0xa7, 0xd9, 0xe6, 0x20,
	0x99, 0x37, 0xbf, 0x47, 0x85, 0xa2, 0x40, 0x91, 0x77, 0x6c, 0xf4, 0x26, 0x3f, 0xf7, 0x01, 0x98,
	0x15, 0x77, 0x93, 0x1e, 0xc3, 0xa4, 0x35, 0x3c, 0xbf, 0x19, 0x58, 0xbd, 0x18, 0xf9, 0xc3, 0x9d,
	0xfc, 0x07, 0x7f, 0xf1, 0x13, 0x02, 0x07, 0x5c, 0x97, 0x2d, 0x9d, 0xc4, 0x61, 0xba, 0x60, 0x6e,
	0x4d, 0x2e, 0x60, 0x9a, 0x4b, 0xb1, 0xae, 0x4a, 0x3a, 0x75, 0xd0, 0x8f, 0xb6, 0xa0, 0xff, 0x60,
	0x65, 0xef, 0x9c, 0xe7, 0x52, 0x18, 0xdd, 0xb1, 0xa1, 0x81, 0xac, 0x60, 0x5a, 0xe8, 0x8e, 0x59,
	0x41, 0x67, 0x71, 0x90, 0xce, 0xd9, 0xa0, 0x48, 0x04, 0x73, 0xc5, 0x35, 0xaf, 0x6b, 0xac, 0xe9,
	0x3c, 0x0e, 0xd2, 0x09, 0xdb, 0x68, 0xf2, 0x0c, 0x8e, 0x1a, 0x29, 0x2a, 0x23, 0xf5, 0x17, 0x5e,
	0x14, 0x1a, 0xdb, 0x96, 0x2e, 0x1c, 0xe4, 0xe1, 0x50, 0x7e, 0xeb, 0xab, 0xd1, 0x05, 0x2c, 0xb7,
	0xf6, 0xec, 0x8f, 0x79, 0x83, 0xdd, 0x10, 0x49, 0xbf, 0xec, 0xe3, 0xf8, 0xce, 0x6b, 0x8b, 0x63,
	0x1c, 0x4e, 0xbc, 0xde, 0x7f, 0x15, 0x24, 0x8f, 0x61, 0xe9, 0xc8, 0x87, 0x7b, 0x39, 0x86, 0x09,
	0x6a, 0x2d, 0xf5, 0xd0, 0xec, 0xc5, 0xf9, 0xaf, 0x00, 0x8e, 0x3e, 0x0e, 0xef, 0x8e, 0x59, 0x61,
	0xaa, 0x06, 0x49, 0x0e, 0xe4, 0xf6, 0xfd, 0x92, 0x27, 0x5b, 0x89, 0xfc, 0xf7, 0x85, 0x45, 0x4f,
	0xef, 0x70, 0x79, 0x98, 0x64, 0x8f, 0xbc, 0x84, 0xb0, 0x0f, 0xe9, 0x64, 0x67, 0xce, 0xd1, 0xea,
	0xdf, 0xf2, 0xa6, 0xef, 0x0d, 0xdc, 0xbf, 0x42, 0xe3, 0xff, 0xf7, 0x41, 0xac, 0x25, 0x59, 0x65,
	0x7e, 0x1c, 0xb2, 0x71, 0x1c, 0xb2, 0xcb, 0x7e, 0x1c, 0xa2, 0x93, 0x5b, 0xcf, 0xae, 0xb7, 0x27,
	0x7b, 0xd7, 0x53, 0x67, 0x7c, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x86, 0xf5, 0x91, 0x70,
	0x03, 0x00, 0x00,
}
