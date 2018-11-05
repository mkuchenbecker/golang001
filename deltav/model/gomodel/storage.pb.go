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
	Amount               int64       `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
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

func (m *Storage) GetAmount() int64 {
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

type StorageRequest struct {
	Storage              []*Storage `protobuf:"bytes,1,rep,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StorageRequest) Reset()         { *m = StorageRequest{} }
func (m *StorageRequest) String() string { return proto.CompactTextString(m) }
func (*StorageRequest) ProtoMessage()    {}
func (*StorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *StorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageRequest.Unmarshal(m, b)
}
func (m *StorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageRequest.Marshal(b, m, deterministic)
}
func (m *StorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageRequest.Merge(m, src)
}
func (m *StorageRequest) XXX_Size() int {
	return xxx_messageInfo_StorageRequest.Size(m)
}
func (m *StorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StorageRequest proto.InternalMessageInfo

func (m *StorageRequest) GetStorage() []*Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

type StorageResponse struct {
	Storage              []*Storage `protobuf:"bytes,1,rep,name=storage,proto3" json:"storage,omitempty"`
	Err                  string     `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StorageResponse) Reset()         { *m = StorageResponse{} }
func (m *StorageResponse) String() string { return proto.CompactTextString(m) }
func (*StorageResponse) ProtoMessage()    {}
func (*StorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

func (m *StorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageResponse.Unmarshal(m, b)
}
func (m *StorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageResponse.Marshal(b, m, deterministic)
}
func (m *StorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageResponse.Merge(m, src)
}
func (m *StorageResponse) XXX_Size() int {
	return xxx_messageInfo_StorageResponse.Size(m)
}
func (m *StorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StorageResponse proto.InternalMessageInfo

func (m *StorageResponse) GetStorage() []*Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *StorageResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterEnum("deltav.model.StorageType", StorageType_name, StorageType_value)
	proto.RegisterType((*Storage)(nil), "deltav.model.Storage")
	proto.RegisterType((*StorageRequest)(nil), "deltav.model.StorageRequest")
	proto.RegisterType((*StorageResponse)(nil), "deltav.model.StorageResponse")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xc7, 0x59, 0x8a, 0xa0, 0x03, 0xd2, 0x66, 0xfc, 0x48, 0x35, 0x9a, 0x90, 0x9e, 0x88, 0x89,
	0x35, 0x81, 0x27, 0x68, 0x70, 0x03, 0xc8, 0x67, 0x96, 0x6d, 0x88, 0xa7, 0xa6, 0xa6, 0x1b, 0x35,
	0x02, 0x5b, 0xdb, 0x45, 0xc3, 0xc1, 0x37, 0xf2, 0x21, 0x0d, 0xb4, 0x55, 0x0e, 0x78, 0x30, 0xf1,
	0xb6, 0xf3, 0x9f, 0x99, 0xdf, 0xfe, 0x77, 0x76, 0xe0, 0x30, 0x56, 0x32, 0xf2, 0x1f, 0x85, 0x1d,
	0x46, 0x52, 0x49, 0xac, 0x04, 0x62, 0xa6, 0xfc, 0x37, 0x7b, 0x2e, 0x03, 0x31, 0xb3, 0xc6, 0x50,
	0x9a, 0x24, 0x69, 0x3c, 0x85, 0xa2, 0x3f, 0x97, 0xcb, 0x85, 0x32, 0x49, 0x8d, 0xd4, 0x35, 0x96,
	0x46, 0x78, 0x0d, 0x05, 0xb5, 0x0a, 0x85, 0x99, 0xaf, 0x91, 0x7a, 0xb5, 0x71, 0x66, 0x6f, 0xf7,
	0xdb, 0x69, 0x33, 0x5f, 0x85, 0x82, 0x6d, 0xca, 0x2c, 0x07, 0xaa, 0xa9, 0xc8, 0xc4, 0xeb, 0x52,
	0xc4, 0x0a, 0x6f, 0xa0, 0x94, 0x5a, 0x30, 0x49, 0x4d, 0xab, 0x97, 0x1b, 0x27, 0x3b, 0x19, 0x2c,
	0xab, 0xb2, 0x38, 0xe8, 0xdf, 0x88, 0x38, 0x94, 0x8b, 0x58, 0xfc, 0x99, 0x81, 0x06, 0x68, 0x22,
	0x8a, 0x36, 0xa6, 0x0f, 0xd8, 0xfa, 0x78, 0xf5, 0x01, 0xe5, 0x2d, 0xb7, 0x68, 0xc2, 0xb1, 0x3b,
	0xec, 0x0d, 0x47, 0xd3, 0xa1, 0x37, 0xe1, 0x23, 0xe6, 0xb4, 0xa9, 0xc7, 0xef, 0xc7, 0xd4, 0xc8,
	0x61, 0x15, 0xa0, 0x43, 0xfb, 0x5d, 0x77, 0xd0, 0xf4, 0x7a, 0x6d, 0x83, 0xa0, 0x01, 0x95, 0x5b,
	0xea, 0x72, 0xca, 0xba, 0xee, 0x60, 0xad, 0xe4, 0xb1, 0x02, 0xfb, 0x53, 0x87, 0x53, 0xb6, 0x8e,
	0x34, 0x3c, 0x02, 0x9d, 0xf6, 0x69, 0x8b, 0xb3, 0x6e, 0xcb, 0xbb, 0x1b, 0xb9, 0x7d, 0x3a, 0x31,
	0x0a, 0xa8, 0x43, 0xb9, 0x43, 0x1d, 0x9e, 0x09, 0x7b, 0x8d, 0x4f, 0xf2, 0x73, 0xbf, 0xbf, 0x78,
	0xc1, 0x31, 0xe8, 0xd3, 0x67, 0xf5, 0x14, 0x44, 0xfe, 0x7b, 0xf6, 0x03, 0x17, 0xbb, 0xdf, 0x94,
	0x8c, 0xf1, 0xfc, 0xf2, 0x97, 0x6c, 0x32, 0x21, 0x2b, 0x87, 0x3d, 0x00, 0x27, 0x08, 0xfe, 0x07,
	0xf6, 0x50, 0xdc, 0x6c, 0x4b, 0xf3, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x16, 0x2c, 0xc2, 0xb2, 0x3e,
	0x02, 0x00, 0x00,
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
	WithdrawStorage(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageResponse, error)
	AddStorage(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageResponse, error)
}

type storageTankClient struct {
	cc *grpc.ClientConn
}

func NewStorageTankClient(cc *grpc.ClientConn) StorageTankClient {
	return &storageTankClient{cc}
}

func (c *storageTankClient) WithdrawStorage(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageResponse, error) {
	out := new(StorageResponse)
	err := c.cc.Invoke(ctx, "/deltav.model.StorageTank/WithdrawStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageTankClient) AddStorage(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageResponse, error) {
	out := new(StorageResponse)
	err := c.cc.Invoke(ctx, "/deltav.model.StorageTank/AddStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageTankServer is the server API for StorageTank service.
type StorageTankServer interface {
	WithdrawStorage(context.Context, *StorageRequest) (*StorageResponse, error)
	AddStorage(context.Context, *StorageRequest) (*StorageResponse, error)
}

func RegisterStorageTankServer(s *grpc.Server, srv StorageTankServer) {
	s.RegisterService(&_StorageTank_serviceDesc, srv)
}

func _StorageTank_WithdrawStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
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
		return srv.(StorageTankServer).WithdrawStorage(ctx, req.(*StorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageTank_AddStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
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
		return srv.(StorageTankServer).AddStorage(ctx, req.(*StorageRequest))
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}
