// Code generated by protoc-gen-go. DO NOT EDIT.
// source: booking/booking.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateReq struct {
	Booking              *BookingModel `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateReq) Reset()         { *m = CreateReq{} }
func (m *CreateReq) String() string { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()    {}
func (*CreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ffdf80b5823d65, []int{0}
}

func (m *CreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReq.Unmarshal(m, b)
}
func (m *CreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReq.Marshal(b, m, deterministic)
}
func (m *CreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReq.Merge(m, src)
}
func (m *CreateReq) XXX_Size() int {
	return xxx_messageInfo_CreateReq.Size(m)
}
func (m *CreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReq proto.InternalMessageInfo

func (m *CreateReq) GetBooking() *BookingModel {
	if m != nil {
		return m.Booking
	}
	return nil
}

type CreateRes struct {
	Booking              *BookingModel `protobuf:"bytes,1,opt,name=booking,proto3" json:"booking,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateRes) Reset()         { *m = CreateRes{} }
func (m *CreateRes) String() string { return proto.CompactTextString(m) }
func (*CreateRes) ProtoMessage()    {}
func (*CreateRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ffdf80b5823d65, []int{1}
}

func (m *CreateRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRes.Unmarshal(m, b)
}
func (m *CreateRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRes.Marshal(b, m, deterministic)
}
func (m *CreateRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRes.Merge(m, src)
}
func (m *CreateRes) XXX_Size() int {
	return xxx_messageInfo_CreateRes.Size(m)
}
func (m *CreateRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRes proto.InternalMessageInfo

func (m *CreateRes) GetBooking() *BookingModel {
	if m != nil {
		return m.Booking
	}
	return nil
}

type BookingModel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StartedAt            string   `protobuf:"bytes,3,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	EndedAt              string   `protobuf:"bytes,4,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
	RoomId               string   `protobuf:"bytes,5,opt,name=roomId,proto3" json:"roomId,omitempty"`
	UserId               string   `protobuf:"bytes,6,opt,name=userId,proto3" json:"userId,omitempty"`
	ParticipantsId       []string `protobuf:"bytes,7,rep,name=participantsId,proto3" json:"participantsId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookingModel) Reset()         { *m = BookingModel{} }
func (m *BookingModel) String() string { return proto.CompactTextString(m) }
func (*BookingModel) ProtoMessage()    {}
func (*BookingModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ffdf80b5823d65, []int{2}
}

func (m *BookingModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookingModel.Unmarshal(m, b)
}
func (m *BookingModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookingModel.Marshal(b, m, deterministic)
}
func (m *BookingModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookingModel.Merge(m, src)
}
func (m *BookingModel) XXX_Size() int {
	return xxx_messageInfo_BookingModel.Size(m)
}
func (m *BookingModel) XXX_DiscardUnknown() {
	xxx_messageInfo_BookingModel.DiscardUnknown(m)
}

var xxx_messageInfo_BookingModel proto.InternalMessageInfo

func (m *BookingModel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BookingModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BookingModel) GetStartedAt() string {
	if m != nil {
		return m.StartedAt
	}
	return ""
}

func (m *BookingModel) GetEndedAt() string {
	if m != nil {
		return m.EndedAt
	}
	return ""
}

func (m *BookingModel) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *BookingModel) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *BookingModel) GetParticipantsId() []string {
	if m != nil {
		return m.ParticipantsId
	}
	return nil
}

type FindReq struct {
	StartedAt            string   `protobuf:"bytes,1,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	EndedAt              string   `protobuf:"bytes,2,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindReq) Reset()         { *m = FindReq{} }
func (m *FindReq) String() string { return proto.CompactTextString(m) }
func (*FindReq) ProtoMessage()    {}
func (*FindReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ffdf80b5823d65, []int{3}
}

func (m *FindReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindReq.Unmarshal(m, b)
}
func (m *FindReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindReq.Marshal(b, m, deterministic)
}
func (m *FindReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindReq.Merge(m, src)
}
func (m *FindReq) XXX_Size() int {
	return xxx_messageInfo_FindReq.Size(m)
}
func (m *FindReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindReq proto.InternalMessageInfo

func (m *FindReq) GetStartedAt() string {
	if m != nil {
		return m.StartedAt
	}
	return ""
}

func (m *FindReq) GetEndedAt() string {
	if m != nil {
		return m.EndedAt
	}
	return ""
}

type FindRes struct {
	Booking              []*BookingModel `protobuf:"bytes,1,rep,name=booking,proto3" json:"booking,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FindRes) Reset()         { *m = FindRes{} }
func (m *FindRes) String() string { return proto.CompactTextString(m) }
func (*FindRes) ProtoMessage()    {}
func (*FindRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ffdf80b5823d65, []int{4}
}

func (m *FindRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRes.Unmarshal(m, b)
}
func (m *FindRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRes.Marshal(b, m, deterministic)
}
func (m *FindRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRes.Merge(m, src)
}
func (m *FindRes) XXX_Size() int {
	return xxx_messageInfo_FindRes.Size(m)
}
func (m *FindRes) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRes.DiscardUnknown(m)
}

var xxx_messageInfo_FindRes proto.InternalMessageInfo

func (m *FindRes) GetBooking() []*BookingModel {
	if m != nil {
		return m.Booking
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateReq)(nil), "api.CreateReq")
	proto.RegisterType((*CreateRes)(nil), "api.CreateRes")
	proto.RegisterType((*BookingModel)(nil), "api.BookingModel")
	proto.RegisterType((*FindReq)(nil), "api.FindReq")
	proto.RegisterType((*FindRes)(nil), "api.FindRes")
}

func init() { proto.RegisterFile("booking/booking.proto", fileDescriptor_00ffdf80b5823d65) }

var fileDescriptor_00ffdf80b5823d65 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4b, 0xf4, 0x30,
	0x10, 0xc6, 0xe9, 0x9f, 0xb7, 0xa5, 0xf3, 0x2e, 0x05, 0x07, 0x94, 0xb2, 0x78, 0x58, 0x7a, 0x90,
	0x05, 0xa1, 0xc2, 0x0a, 0xe2, 0x75, 0x15, 0x04, 0x0f, 0x5e, 0x7a, 0xf1, 0x9c, 0x35, 0x41, 0x82,
	0x6e, 0x52, 0x93, 0xf8, 0x05, 0xfd, 0x64, 0xa6, 0xd3, 0x54, 0xdb, 0x3d, 0x2c, 0x78, 0x6a, 0x9e,
	0xe7, 0x37, 0x93, 0x3e, 0xcc, 0x04, 0x4e, 0x77, 0x5a, 0xbf, 0x49, 0xf5, 0x7a, 0x15, 0xbe, 0x4d,
	0x67, 0xb4, 0xd3, 0x98, 0xb0, 0x4e, 0xd6, 0xb7, 0x50, 0xdc, 0x1b, 0xc1, 0x9c, 0x68, 0xc5, 0x07,
	0x5e, 0x42, 0x1e, 0x4a, 0xaa, 0x68, 0x15, 0xad, 0xff, 0x6f, 0x4e, 0x1a, 0x5f, 0xd3, 0xdc, 0x0d,
	0xde, 0x93, 0xe6, 0xe2, 0xbd, 0x1d, 0x2b, 0xa6, 0x9d, 0xf6, 0x6f, 0x9d, 0x5f, 0x11, 0x2c, 0xa6,
	0x04, 0x4b, 0x88, 0x25, 0xa7, 0xc6, 0xa2, 0xf5, 0x27, 0x44, 0x48, 0x15, 0xdb, 0x8b, 0x2a, 0x26,
	0x87, 0xce, 0x78, 0x0e, 0x85, 0x75, 0xcc, 0x38, 0xc1, 0xb7, 0xae, 0x4a, 0x08, 0xfc, 0x1a, 0x58,
	0x41, 0x2e, 0x14, 0x27, 0x96, 0x12, 0x1b, 0x25, 0x9e, 0x41, 0x66, 0xb4, 0xde, 0x3f, 0xf2, 0xea,
	0x1f, 0x81, 0xa0, 0x7a, 0xff, 0xd3, 0x0a, 0xe3, 0xfd, 0x6c, 0xf0, 0x07, 0x85, 0x17, 0x50, 0x76,
	0xfe, 0x56, 0xf9, 0x22, 0x3b, 0xa6, 0x9c, 0xf5, 0x3c, 0x5f, 0x25, 0x9e, 0x1f, 0xb8, 0xf5, 0x16,
	0xf2, 0x07, 0xa9, 0x78, 0x3f, 0xb6, 0x59, 0xb4, 0xe8, 0x48, 0xb4, 0x78, 0x16, 0xad, 0xbe, 0x19,
	0xaf, 0x38, 0x98, 0x5f, 0x72, 0x7c, 0x7e, 0x9b, 0x67, 0xc8, 0x03, 0xc0, 0x35, 0x64, 0xc3, 0x12,
	0xb0, 0xa4, 0x86, 0x9f, 0x5d, 0x2e, 0xe7, 0xda, 0x62, 0x0d, 0x69, 0xff, 0x33, 0x5c, 0x90, 0x1f,
	0xa2, 0x2f, 0xa7, 0xca, 0xee, 0x32, 0x7a, 0x18, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5,
	0x70, 0x47, 0x94, 0x31, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookingClient is the client API for Booking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookingClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error)
	Find(ctx context.Context, in *FindReq, opts ...grpc.CallOption) (*FindRes, error)
}

type bookingClient struct {
	cc *grpc.ClientConn
}

func NewBookingClient(cc *grpc.ClientConn) BookingClient {
	return &bookingClient{cc}
}

func (c *bookingClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error) {
	out := new(CreateRes)
	err := c.cc.Invoke(ctx, "/api.Booking/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingClient) Find(ctx context.Context, in *FindReq, opts ...grpc.CallOption) (*FindRes, error) {
	out := new(FindRes)
	err := c.cc.Invoke(ctx, "/api.Booking/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServer is the server API for Booking service.
type BookingServer interface {
	Create(context.Context, *CreateReq) (*CreateRes, error)
	Find(context.Context, *FindReq) (*FindRes, error)
}

// UnimplementedBookingServer can be embedded to have forward compatible implementations.
type UnimplementedBookingServer struct {
}

func (*UnimplementedBookingServer) Create(ctx context.Context, req *CreateReq) (*CreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedBookingServer) Find(ctx context.Context, req *FindReq) (*FindRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterBookingServer(s *grpc.Server, srv BookingServer) {
	s.RegisterService(&_Booking_serviceDesc, srv)
}

func _Booking_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Booking/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Booking_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Booking/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).Find(ctx, req.(*FindReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Booking_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Booking",
	HandlerType: (*BookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Booking_Create_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _Booking_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking/booking.proto",
}
