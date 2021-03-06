// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package deltav_model

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

type StorageType int32

const (
	StorageType_UNKNOWN_STORAGE_TYPE StorageType = 0
	StorageType_HELIUM3_KG           StorageType = 1
	StorageType_DEUTERIUM_KG         StorageType = 2
	StorageType_WATER_KG             StorageType = 3
	StorageType_ELECTRIC_JOULES      StorageType = 4
	StorageType_HEAT_JOULES          StorageType = 5
)

var StorageType_name = map[int32]string{
	0: "UNKNOWN_STORAGE_TYPE",
	1: "HELIUM3_KG",
	2: "DEUTERIUM_KG",
	3: "WATER_KG",
	4: "ELECTRIC_JOULES",
	5: "HEAT_JOULES",
}

var StorageType_value = map[string]int32{
	"UNKNOWN_STORAGE_TYPE": 0,
	"HELIUM3_KG":           1,
	"DEUTERIUM_KG":         2,
	"WATER_KG":             3,
	"ELECTRIC_JOULES":      4,
	"HEAT_JOULES":          5,
}

func (x StorageType) String() string {
	return proto.EnumName(StorageType_name, int32(x))
}

func (StorageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

type Storage struct {
	Amount               float64     `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Type                 StorageType `protobuf:"varint,2,opt,name=type,proto3,enum=deltav.model.StorageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Storage) GetType() StorageType {
	if m != nil {
		return m.Type
	}
	return StorageType_UNKNOWN_STORAGE_TYPE
}

type AddStorageRequest struct {
	Storage              *Storage `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStorageRequest) Reset()         { *m = AddStorageRequest{} }
func (m *AddStorageRequest) String() string { return proto.CompactTextString(m) }
func (*AddStorageRequest) ProtoMessage()    {}
func (*AddStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *AddStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStorageRequest.Unmarshal(m, b)
}
func (m *AddStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStorageRequest.Marshal(b, m, deterministic)
}
func (m *AddStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStorageRequest.Merge(m, src)
}
func (m *AddStorageRequest) XXX_Size() int {
	return xxx_messageInfo_AddStorageRequest.Size(m)
}
func (m *AddStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddStorageRequest proto.InternalMessageInfo

func (m *AddStorageRequest) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

type AddStorageResponse struct {
	Storage              *Storage `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStorageResponse) Reset()         { *m = AddStorageResponse{} }
func (m *AddStorageResponse) String() string { return proto.CompactTextString(m) }
func (*AddStorageResponse) ProtoMessage()    {}
func (*AddStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

func (m *AddStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStorageResponse.Unmarshal(m, b)
}
func (m *AddStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStorageResponse.Marshal(b, m, deterministic)
}
func (m *AddStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStorageResponse.Merge(m, src)
}
func (m *AddStorageResponse) XXX_Size() int {
	return xxx_messageInfo_AddStorageResponse.Size(m)
}
func (m *AddStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddStorageResponse proto.InternalMessageInfo

func (m *AddStorageResponse) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *AddStorageResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type WithdrawStorageRequest struct {
	Storage              *Storage `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawStorageRequest) Reset()         { *m = WithdrawStorageRequest{} }
func (m *WithdrawStorageRequest) String() string { return proto.CompactTextString(m) }
func (*WithdrawStorageRequest) ProtoMessage()    {}
func (*WithdrawStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{3}
}

func (m *WithdrawStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawStorageRequest.Unmarshal(m, b)
}
func (m *WithdrawStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawStorageRequest.Marshal(b, m, deterministic)
}
func (m *WithdrawStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawStorageRequest.Merge(m, src)
}
func (m *WithdrawStorageRequest) XXX_Size() int {
	return xxx_messageInfo_WithdrawStorageRequest.Size(m)
}
func (m *WithdrawStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawStorageRequest proto.InternalMessageInfo

func (m *WithdrawStorageRequest) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

type WithdrawStorageResponse struct {
	Storage              *Storage `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawStorageResponse) Reset()         { *m = WithdrawStorageResponse{} }
func (m *WithdrawStorageResponse) String() string { return proto.CompactTextString(m) }
func (*WithdrawStorageResponse) ProtoMessage()    {}
func (*WithdrawStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{4}
}

func (m *WithdrawStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawStorageResponse.Unmarshal(m, b)
}
func (m *WithdrawStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawStorageResponse.Marshal(b, m, deterministic)
}
func (m *WithdrawStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawStorageResponse.Merge(m, src)
}
func (m *WithdrawStorageResponse) XXX_Size() int {
	return xxx_messageInfo_WithdrawStorageResponse.Size(m)
}
func (m *WithdrawStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawStorageResponse proto.InternalMessageInfo

func (m *WithdrawStorageResponse) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *WithdrawStorageResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type StorageStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StorageStatusRequest) Reset()         { *m = StorageStatusRequest{} }
func (m *StorageStatusRequest) String() string { return proto.CompactTextString(m) }
func (*StorageStatusRequest) ProtoMessage()    {}
func (*StorageStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{5}
}

func (m *StorageStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageStatusRequest.Unmarshal(m, b)
}
func (m *StorageStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageStatusRequest.Marshal(b, m, deterministic)
}
func (m *StorageStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageStatusRequest.Merge(m, src)
}
func (m *StorageStatusRequest) XXX_Size() int {
	return xxx_messageInfo_StorageStatusRequest.Size(m)
}
func (m *StorageStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StorageStatusRequest proto.InternalMessageInfo

type StorageStatusResponse struct {
	Capacity             float64     `protobuf:"fixed64,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Current              float64     `protobuf:"fixed64,2,opt,name=current,proto3" json:"current,omitempty"`
	Type                 StorageType `protobuf:"varint,3,opt,name=type,proto3,enum=deltav.model.StorageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StorageStatusResponse) Reset()         { *m = StorageStatusResponse{} }
func (m *StorageStatusResponse) String() string { return proto.CompactTextString(m) }
func (*StorageStatusResponse) ProtoMessage()    {}
func (*StorageStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{6}
}

func (m *StorageStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageStatusResponse.Unmarshal(m, b)
}
func (m *StorageStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageStatusResponse.Marshal(b, m, deterministic)
}
func (m *StorageStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageStatusResponse.Merge(m, src)
}
func (m *StorageStatusResponse) XXX_Size() int {
	return xxx_messageInfo_StorageStatusResponse.Size(m)
}
func (m *StorageStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StorageStatusResponse proto.InternalMessageInfo

func (m *StorageStatusResponse) GetCapacity() float64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *StorageStatusResponse) GetCurrent() float64 {
	if m != nil {
		return m.Current
	}
	return 0
}

func (m *StorageStatusResponse) GetType() StorageType {
	if m != nil {
		return m.Type
	}
	return StorageType_UNKNOWN_STORAGE_TYPE
}

func init() {
	proto.RegisterEnum("deltav.model.StorageType", StorageType_name, StorageType_value)
	proto.RegisterType((*Storage)(nil), "deltav.model.Storage")
	proto.RegisterType((*AddStorageRequest)(nil), "deltav.model.AddStorageRequest")
	proto.RegisterType((*AddStorageResponse)(nil), "deltav.model.AddStorageResponse")
	proto.RegisterType((*WithdrawStorageRequest)(nil), "deltav.model.WithdrawStorageRequest")
	proto.RegisterType((*WithdrawStorageResponse)(nil), "deltav.model.WithdrawStorageResponse")
	proto.RegisterType((*StorageStatusRequest)(nil), "deltav.model.StorageStatusRequest")
	proto.RegisterType((*StorageStatusResponse)(nil), "deltav.model.StorageStatusResponse")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x7f, 0x6f, 0x93, 0x50,
	0x14, 0x1d, 0xed, 0x6c, 0xe7, 0x6d, 0x5d, 0xf1, 0xba, 0x55, 0xe4, 0x1f, 0x1b, 0xd4, 0x64, 0x31,
	0x11, 0x93, 0xee, 0x13, 0x90, 0xed, 0x65, 0xc3, 0x75, 0xed, 0x7c, 0x40, 0x88, 0x89, 0x09, 0x3e,
	0xcb, 0x8b, 0x36, 0xb6, 0x80, 0xf0, 0xd0, 0x34, 0xd1, 0xcf, 0xe1, 0xd7, 0x35, 0xa5, 0x8f, 0xfe,
	0xf6, 0x57, 0xdc, 0x7f, 0x9c, 0xc3, 0xb9, 0xe7, 0x1e, 0xde, 0x79, 0xc0, 0xbd, 0x4c, 0xc4, 0x29,
	0xfb, 0xc0, 0xcd, 0x24, 0x8d, 0x45, 0x8c, 0xcd, 0x90, 0x8f, 0x05, 0xfb, 0x62, 0x4e, 0xe2, 0x90,
	0x8f, 0x8d, 0x1b, 0xa8, 0x3b, 0xf3, 0xd7, 0xd8, 0x86, 0x1a, 0x9b, 0xc4, 0x79, 0x24, 0x34, 0xa5,
	0xa3, 0x9c, 0x28, 0x54, 0x22, 0x7c, 0x01, 0xfb, 0x62, 0x9a, 0x70, 0xad, 0xd2, 0x51, 0x4e, 0x0e,
	0xbb, 0x8f, 0xcc, 0xd5, 0x79, 0x53, 0x0e, 0xbb, 0xd3, 0x84, 0xd3, 0x42, 0x66, 0x9c, 0xc3, 0x7d,
	0x2b, 0x0c, 0x25, 0x4f, 0xf9, 0xe7, 0x9c, 0x67, 0x02, 0x5f, 0x42, 0x5d, 0xa6, 0x28, 0xcc, 0x1b,
	0xdd, 0xe3, 0x9d, 0x36, 0xb4, 0x54, 0x19, 0x3e, 0xe0, 0xaa, 0x4b, 0x96, 0xc4, 0x51, 0xc6, 0xff,
	0xd9, 0x06, 0x55, 0xa8, 0xf2, 0x34, 0x2d, 0xa2, 0xdf, 0xa5, 0xb3, 0x47, 0xc3, 0x86, 0xb6, 0x3f,
	0x12, 0x1f, 0xc3, 0x94, 0x7d, 0xfd, 0xdf, 0x8c, 0x6f, 0xe1, 0xe1, 0x96, 0xd5, 0xed, 0x05, 0x6d,
	0xc3, 0x91, 0x54, 0x39, 0x82, 0x89, 0x3c, 0x93, 0x31, 0x8d, 0x6f, 0x70, 0xbc, 0xc1, 0xcb, 0x9d,
	0x3a, 0x1c, 0x0c, 0x59, 0xc2, 0x86, 0x23, 0x31, 0x95, 0x0d, 0x2e, 0x30, 0x6a, 0x50, 0x1f, 0xe6,
	0x69, 0xca, 0x23, 0x51, 0xac, 0x50, 0x68, 0x09, 0x17, 0xed, 0x56, 0xff, 0xaa, 0xdd, 0xe7, 0xdf,
	0xa1, 0xb1, 0x42, 0xa2, 0x06, 0x47, 0x5e, 0xff, 0xaa, 0x3f, 0xf0, 0xfb, 0x81, 0xe3, 0x0e, 0xa8,
	0x75, 0x41, 0x02, 0xf7, 0xcd, 0x0d, 0x51, 0xf7, 0xf0, 0x10, 0xe0, 0x92, 0xf4, 0x6c, 0xef, 0xfa,
	0x34, 0xb8, 0xba, 0x50, 0x15, 0x54, 0xa1, 0x79, 0x4e, 0x3c, 0x97, 0x50, 0xdb, 0xbb, 0x9e, 0x31,
	0x15, 0x6c, 0xc2, 0x81, 0x6f, 0xb9, 0x84, 0xce, 0x50, 0x15, 0x1f, 0x40, 0x8b, 0xf4, 0xc8, 0x99,
	0x4b, 0xed, 0xb3, 0xe0, 0xd5, 0xc0, 0xeb, 0x11, 0x47, 0xdd, 0xc7, 0x16, 0x34, 0x2e, 0x89, 0xe5,
	0x96, 0xc4, 0x9d, 0xee, 0x8f, 0xca, 0x72, 0x3f, 0x8b, 0x3e, 0xe1, 0x3b, 0x68, 0x6d, 0x54, 0x80,
	0x4f, 0xd7, 0x3f, 0x61, 0x77, 0xd9, 0xfa, 0xb3, 0x3f, 0xa8, 0xe6, 0x67, 0x6a, 0xec, 0xe1, 0x6b,
	0x80, 0xe5, 0x45, 0xc4, 0xc7, 0xeb, 0x63, 0x5b, 0x17, 0x5d, 0xef, 0xfc, 0x5a, 0xb0, 0xb0, 0x74,
	0xa0, 0x36, 0xaf, 0x0e, 0x8d, 0x9d, 0xc7, 0xbd, 0xd6, 0xb7, 0xfe, 0xe4, 0xb7, 0x9a, 0xd2, 0xf4,
	0x7d, 0xad, 0xf8, 0xbb, 0x4f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xda, 0xb4, 0x5d, 0xee,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageTankClient is the client API for StorageTank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageTankClient interface {
	WithdrawStorage(ctx context.Context, in *WithdrawStorageRequest, opts ...grpc.CallOption) (*WithdrawStorageResponse, error)
	AddStorage(ctx context.Context, in *AddStorageRequest, opts ...grpc.CallOption) (*AddStorageResponse, error)
	Status(ctx context.Context, in *StorageStatusRequest, opts ...grpc.CallOption) (*StorageStatusResponse, error)
}

type storageTankClient struct {
	cc *grpc.ClientConn
}

func NewStorageTankClient(cc *grpc.ClientConn) StorageTankClient {
	return &storageTankClient{cc}
}

func (c *storageTankClient) WithdrawStorage(ctx context.Context, in *WithdrawStorageRequest, opts ...grpc.CallOption) (*WithdrawStorageResponse, error) {
	out := new(WithdrawStorageResponse)
	err := c.cc.Invoke(ctx, "/deltav.model.StorageTank/WithdrawStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageTankClient) AddStorage(ctx context.Context, in *AddStorageRequest, opts ...grpc.CallOption) (*AddStorageResponse, error) {
	out := new(AddStorageResponse)
	err := c.cc.Invoke(ctx, "/deltav.model.StorageTank/AddStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageTankClient) Status(ctx context.Context, in *StorageStatusRequest, opts ...grpc.CallOption) (*StorageStatusResponse, error) {
	out := new(StorageStatusResponse)
	err := c.cc.Invoke(ctx, "/deltav.model.StorageTank/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageTankServer is the server API for StorageTank service.
type StorageTankServer interface {
	WithdrawStorage(context.Context, *WithdrawStorageRequest) (*WithdrawStorageResponse, error)
	AddStorage(context.Context, *AddStorageRequest) (*AddStorageResponse, error)
	Status(context.Context, *StorageStatusRequest) (*StorageStatusResponse, error)
}

func RegisterStorageTankServer(s *grpc.Server, srv StorageTankServer) {
	s.RegisterService(&_StorageTank_serviceDesc, srv)
}

func _StorageTank_WithdrawStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageTankServer).WithdrawStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deltav.model.StorageTank/WithdrawStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageTankServer).WithdrawStorage(ctx, req.(*WithdrawStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageTank_AddStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageTankServer).AddStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deltav.model.StorageTank/AddStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageTankServer).AddStorage(ctx, req.(*AddStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageTank_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageTankServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deltav.model.StorageTank/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageTankServer).Status(ctx, req.(*StorageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageTank_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deltav.model.StorageTank",
	HandlerType: (*StorageTankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WithdrawStorage",
			Handler:    _StorageTank_WithdrawStorage_Handler,
		},
		{
			MethodName: "AddStorage",
			Handler:    _StorageTank_AddStorage_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _StorageTank_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}
