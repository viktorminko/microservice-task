// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client-service.proto

package v1

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Client struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Mobile               string   `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_108c2df084c60df0, []int{0}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Client) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Client) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type CreateRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Client               *Client  `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_108c2df084c60df0, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

type CreateResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_108c2df084c60df0, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Client)(nil), "v1.Client")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
}

func init() { proto.RegisterFile("client-service.proto", fileDescriptor_108c2df084c60df0) }

var fileDescriptor_108c2df084c60df0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xd9, 0x74, 0x0d, 0x38, 0xb2, 0x8b, 0x0e, 0x8b, 0x04, 0x4f, 0x92, 0x93, 0x17, 0x2b,
	0xad, 0x2f, 0x20, 0x14, 0x5f, 0xa0, 0xde, 0xbc, 0xa5, 0xed, 0x1c, 0x02, 0x6d, 0x53, 0x9b, 0xd8,
	0xe7, 0x97, 0xce, 0x54, 0xe8, 0x61, 0x6f, 0x33, 0xff, 0xc0, 0xf7, 0x7f, 0x09, 0x5c, 0xda, 0xde,
	0xd3, 0x98, 0x5e, 0x23, 0xcd, 0x8b, 0x6f, 0x29, 0x9f, 0xe6, 0x90, 0x02, 0xaa, 0xa5, 0xb0, 0xdf,
	0xa0, 0x2b, 0xbe, 0xe1, 0x19, 0x94, 0xef, 0xcc, 0xe1, 0xf9, 0xf0, 0x72, 0x5b, 0x2b, 0xdf, 0x21,
	0xc2, 0x71, 0x74, 0x03, 0x19, 0xc5, 0x09, 0xcf, 0x78, 0x81, 0x1b, 0x1a, 0x9c, 0xef, 0x4d, 0xc6,
	0xa1, 0x2c, 0xf8, 0x08, 0x7a, 0x08, 0x8d, 0xef, 0xc9, 0x1c, 0x39, 0xde, 0x36, 0xfb, 0x09, 0xa7,
	0x6a, 0x26, 0x97, 0xa8, 0xa6, 0x9f, 0x5f, 0x8a, 0x09, 0xef, 0x21, 0x73, 0x93, 0xdf, 0x3a, 0xd6,
	0x11, 0x2d, 0x68, 0x51, 0xe3, 0x9a, 0xbb, 0x12, 0xf2, 0xa5, 0xc8, 0x45, 0xa8, 0xde, 0x2e, 0xb6,
	0x84, 0xf3, 0x3f, 0x26, 0x4e, 0x61, 0x8c, 0x74, 0x85, 0x23, 0xf2, 0x2b, 0x23, 0x5b, 0xe5, 0xcb,
	0x0f, 0x38, 0x09, 0xe5, 0x4b, 0x5e, 0x8c, 0x6f, 0xa0, 0x05, 0x82, 0x0f, 0x5c, 0xb1, 0xf7, 0x7a,
	0xc2, 0x7d, 0x24, 0x1d, 0x8d, 0xe6, 0x3f, 0x7a, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x02, 0xf4,
	0x2b, 0x4e, 0x3b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type clientServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientServiceClient(cc *grpc.ClientConn) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.ClientService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ClientService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client-service.proto",
}
