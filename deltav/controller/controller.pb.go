// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controller.proto

package controller

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DetectableProperty_Type int32

const (
	DetectableProperty_UNKNOWN DetectableProperty_Type = 0
	DetectableProperty_RF      DetectableProperty_Type = 1
	DetectableProperty_GAMMA   DetectableProperty_Type = 2
	DetectableProperty_THERM   DetectableProperty_Type = 3
)

var DetectableProperty_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "RF",
	2: "GAMMA",
	3: "THERM",
}

var DetectableProperty_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"RF":      1,
	"GAMMA":   2,
	"THERM":   3,
}

func (x DetectableProperty_Type) String() string {
	return proto.EnumName(DetectableProperty_Type_name, int32(x))
}

func (DetectableProperty_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{1, 0}
}

type Position struct {
	X                    float64  `protobuf:"fixed64,1,opt,name=X,json=x,proto3" json:"X,omitempty"`
	Y                    float64  `protobuf:"fixed64,2,opt,name=Y,json=y,proto3" json:"Y,omitempty"`
	Z                    float64  `protobuf:"fixed64,3,opt,name=Z,json=z,proto3" json:"Z,omitempty"`
	T                    int64    `protobuf:"varint,4,opt,name=T,json=t,proto3" json:"T,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{0}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetT() int64 {
	if m != nil {
		return m.T
	}
	return 0
}

type DetectableProperty struct {
	Type                 DetectableProperty_Type `protobuf:"varint,1,opt,name=type,proto3,enum=controller.DetectableProperty_Type" json:"type,omitempty"`
	Intensity            float32                 `protobuf:"fixed32,2,opt,name=intensity,proto3" json:"intensity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DetectableProperty) Reset()         { *m = DetectableProperty{} }
func (m *DetectableProperty) String() string { return proto.CompactTextString(m) }
func (*DetectableProperty) ProtoMessage()    {}
func (*DetectableProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{1}
}

func (m *DetectableProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectableProperty.Unmarshal(m, b)
}
func (m *DetectableProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectableProperty.Marshal(b, m, deterministic)
}
func (m *DetectableProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectableProperty.Merge(m, src)
}
func (m *DetectableProperty) XXX_Size() int {
	return xxx_messageInfo_DetectableProperty.Size(m)
}
func (m *DetectableProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectableProperty.DiscardUnknown(m)
}

var xxx_messageInfo_DetectableProperty proto.InternalMessageInfo

func (m *DetectableProperty) GetType() DetectableProperty_Type {
	if m != nil {
		return m.Type
	}
	return DetectableProperty_UNKNOWN
}

func (m *DetectableProperty) GetIntensity() float32 {
	if m != nil {
		return m.Intensity
	}
	return 0
}

type RegisterRequest struct {
	Position             *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{2}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

type RegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{3}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

type DetectRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetectRequest) Reset()         { *m = DetectRequest{} }
func (m *DetectRequest) String() string { return proto.CompactTextString(m) }
func (*DetectRequest) ProtoMessage()    {}
func (*DetectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{4}
}

func (m *DetectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectRequest.Unmarshal(m, b)
}
func (m *DetectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectRequest.Marshal(b, m, deterministic)
}
func (m *DetectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectRequest.Merge(m, src)
}
func (m *DetectRequest) XXX_Size() int {
	return xxx_messageInfo_DetectRequest.Size(m)
}
func (m *DetectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetectRequest proto.InternalMessageInfo

type DetectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetectResponse) Reset()         { *m = DetectResponse{} }
func (m *DetectResponse) String() string { return proto.CompactTextString(m) }
func (*DetectResponse) ProtoMessage()    {}
func (*DetectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{5}
}

func (m *DetectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectResponse.Unmarshal(m, b)
}
func (m *DetectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectResponse.Marshal(b, m, deterministic)
}
func (m *DetectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectResponse.Merge(m, src)
}
func (m *DetectResponse) XXX_Size() int {
	return xxx_messageInfo_DetectResponse.Size(m)
}
func (m *DetectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetectResponse proto.InternalMessageInfo

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{6}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type GetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{7}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("controller.DetectableProperty_Type", DetectableProperty_Type_name, DetectableProperty_Type_value)
	proto.RegisterType((*Position)(nil), "controller.Position")
	proto.RegisterType((*DetectableProperty)(nil), "controller.DetectableProperty")
	proto.RegisterType((*RegisterRequest)(nil), "controller.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "controller.RegisterResponse")
	proto.RegisterType((*DetectRequest)(nil), "controller.DetectRequest")
	proto.RegisterType((*DetectResponse)(nil), "controller.DetectResponse")
	proto.RegisterType((*GetRequest)(nil), "controller.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "controller.GetResponse")
}

func init() { proto.RegisterFile("controller.proto", fileDescriptor_ed7f10298fa1d90f) }

var fileDescriptor_ed7f10298fa1d90f = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x65, 0x5a, 0x3e, 0x3e, 0xb8, 0xfc, 0x35, 0x37, 0x46, 0x2b, 0x12, 0x43, 0xea, 0x86, 0x15,
	0x51, 0x5c, 0xe8, 0x16, 0x51, 0x6b, 0x62, 0x8a, 0x64, 0x82, 0x41, 0xdd, 0x09, 0xdc, 0x98, 0x26,
	0x4d, 0xa7, 0xb6, 0x63, 0x62, 0x7d, 0x18, 0x1f, 0xcf, 0xe7, 0x30, 0xfd, 0x6f, 0x84, 0x5d, 0xcf,
	0x3d, 0x67, 0xce, 0x9c, 0x73, 0x3b, 0xa0, 0xad, 0x85, 0x2b, 0x7d, 0xe1, 0x38, 0xe4, 0x8f, 0x3c,
	0x5f, 0x48, 0x81, 0x50, 0x4c, 0x8c, 0x2b, 0xa8, 0xcf, 0x45, 0x60, 0x4b, 0x5b, 0xb8, 0xd8, 0x02,
	0xf6, 0xa4, 0xb3, 0x01, 0x1b, 0x32, 0xce, 0x3e, 0x23, 0xf4, 0xac, 0x2b, 0x09, 0x0a, 0x23, 0xf4,
	0xa2, 0xab, 0x09, 0xfa, 0x8a, 0xd0, 0x42, 0xaf, 0x0e, 0xd8, 0x50, 0xe5, 0x4c, 0x1a, 0xdf, 0x0c,
	0xf0, 0x9a, 0x24, 0xad, 0xe5, 0xeb, 0xca, 0xa1, 0xb9, 0x2f, 0x3c, 0xf2, 0x65, 0x88, 0x17, 0x50,
	0x95, 0xa1, 0x47, 0xb1, 0x63, 0x67, 0x7c, 0x32, 0x2a, 0xe5, 0xd8, 0x56, 0x8f, 0x16, 0xa1, 0x47,
	0x3c, 0x3e, 0x80, 0x7d, 0x68, 0xd8, 0xae, 0x24, 0x37, 0xb0, 0x65, 0x18, 0x27, 0x50, 0x78, 0x31,
	0x30, 0xce, 0xa0, 0x1a, 0x69, 0xb1, 0x09, 0xff, 0x1f, 0x67, 0xf7, 0xb3, 0x87, 0xe5, 0x4c, 0xab,
	0x60, 0x0d, 0x14, 0x7e, 0xab, 0x31, 0x6c, 0xc0, 0x3f, 0x73, 0x62, 0x59, 0x13, 0x4d, 0x89, 0x3e,
	0x17, 0x77, 0x37, 0xdc, 0xd2, 0x54, 0x63, 0x0a, 0x5d, 0x4e, 0x6f, 0x76, 0x20, 0xc9, 0xe7, 0xf4,
	0xfe, 0x41, 0x81, 0xc4, 0x53, 0xa8, 0x7b, 0x69, 0xef, 0x38, 0x60, 0x73, 0xbc, 0x57, 0x0e, 0x98,
	0xed, 0x84, 0xe7, 0x2a, 0x03, 0x41, 0x2b, 0x4c, 0x02, 0x4f, 0xb8, 0x01, 0x19, 0x5d, 0x68, 0x27,
	0x55, 0x52, 0x5b, 0x43, 0x83, 0x4e, 0x36, 0x48, 0x25, 0x2d, 0x00, 0x93, 0x72, 0xbe, 0x0d, 0xcd,
	0x18, 0x25, 0xe4, 0xf8, 0x87, 0xc1, 0x71, 0xb1, 0x8b, 0xa5, 0xf0, 0x9d, 0x8d, 0x25, 0x36, 0xe4,
	0x4c, 0xf3, 0x24, 0x68, 0x42, 0x3d, 0xbb, 0x16, 0x8f, 0xca, 0x11, 0xff, 0x34, 0xea, 0xf5, 0x77,
	0x93, 0x69, 0x8c, 0x0a, 0x4e, 0xa0, 0x96, 0x5c, 0x85, 0x87, 0xdb, 0xbf, 0x22, 0x33, 0xe9, 0xed,
	0xa2, 0x72, 0x8b, 0x4b, 0x50, 0x4d, 0x92, 0xb8, 0x5f, 0x16, 0x15, 0xe5, 0x7a, 0x07, 0x5b, 0xf3,
	0xec, 0xe4, 0xaa, 0x16, 0xbf, 0xbc, 0xf3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x3e, 0x40,
	0xf3, 0x8d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DetectableWorldModelControllerClient is the client API for DetectableWorldModelController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetectableWorldModelControllerClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Detect(ctx context.Context, in *DetectRequest, opts ...grpc.CallOption) (*DetectResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type detectableWorldModelControllerClient struct {
	cc *grpc.ClientConn
}

func NewDetectableWorldModelControllerClient(cc *grpc.ClientConn) DetectableWorldModelControllerClient {
	return &detectableWorldModelControllerClient{cc}
}

func (c *detectableWorldModelControllerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/controller.DetectableWorldModelController/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detectableWorldModelControllerClient) Detect(ctx context.Context, in *DetectRequest, opts ...grpc.CallOption) (*DetectResponse, error) {
	out := new(DetectResponse)
	err := c.cc.Invoke(ctx, "/controller.DetectableWorldModelController/Detect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detectableWorldModelControllerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/controller.DetectableWorldModelController/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetectableWorldModelControllerServer is the server API for DetectableWorldModelController service.
type DetectableWorldModelControllerServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Detect(context.Context, *DetectRequest) (*DetectResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterDetectableWorldModelControllerServer(s *grpc.Server, srv DetectableWorldModelControllerServer) {
	s.RegisterService(&_DetectableWorldModelController_serviceDesc, srv)
}

func _DetectableWorldModelController_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectableWorldModelControllerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.DetectableWorldModelController/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectableWorldModelControllerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DetectableWorldModelController_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectableWorldModelControllerServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.DetectableWorldModelController/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectableWorldModelControllerServer).Detect(ctx, req.(*DetectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DetectableWorldModelController_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectableWorldModelControllerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.DetectableWorldModelController/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectableWorldModelControllerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DetectableWorldModelController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "controller.DetectableWorldModelController",
	HandlerType: (*DetectableWorldModelControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _DetectableWorldModelController_Register_Handler,
		},
		{
			MethodName: "Detect",
			Handler:    _DetectableWorldModelController_Detect_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DetectableWorldModelController_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller.proto",
}
