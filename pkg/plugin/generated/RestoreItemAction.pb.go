// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RestoreItemAction.proto

package generated

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RestoreItemActionExecuteRequest struct {
	Plugin         string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	Item           []byte `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Restore        []byte `protobuf:"bytes,3,opt,name=restore,proto3" json:"restore,omitempty"`
	ItemFromBackup []byte `protobuf:"bytes,4,opt,name=itemFromBackup,proto3" json:"itemFromBackup,omitempty"`
}

func (m *RestoreItemActionExecuteRequest) Reset()                    { *m = RestoreItemActionExecuteRequest{} }
func (m *RestoreItemActionExecuteRequest) String() string            { return proto.CompactTextString(m) }
func (*RestoreItemActionExecuteRequest) ProtoMessage()               {}
func (*RestoreItemActionExecuteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *RestoreItemActionExecuteRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *RestoreItemActionExecuteRequest) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *RestoreItemActionExecuteRequest) GetRestore() []byte {
	if m != nil {
		return m.Restore
	}
	return nil
}

func (m *RestoreItemActionExecuteRequest) GetItemFromBackup() []byte {
	if m != nil {
		return m.ItemFromBackup
	}
	return nil
}

type RestoreItemActionExecuteResponse struct {
	Item            []byte                `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	AdditionalItems []*ResourceIdentifier `protobuf:"bytes,2,rep,name=additionalItems" json:"additionalItems,omitempty"`
	SkipRestore     bool                  `protobuf:"varint,3,opt,name=skipRestore" json:"skipRestore,omitempty"`
}

func (m *RestoreItemActionExecuteResponse) Reset()         { *m = RestoreItemActionExecuteResponse{} }
func (m *RestoreItemActionExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*RestoreItemActionExecuteResponse) ProtoMessage()    {}
func (*RestoreItemActionExecuteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{1}
}

func (m *RestoreItemActionExecuteResponse) GetItem() []byte {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *RestoreItemActionExecuteResponse) GetAdditionalItems() []*ResourceIdentifier {
	if m != nil {
		return m.AdditionalItems
	}
	return nil
}

func (m *RestoreItemActionExecuteResponse) GetSkipRestore() bool {
	if m != nil {
		return m.SkipRestore
	}
	return false
}

func init() {
	proto.RegisterType((*RestoreItemActionExecuteRequest)(nil), "generated.RestoreItemActionExecuteRequest")
	proto.RegisterType((*RestoreItemActionExecuteResponse)(nil), "generated.RestoreItemActionExecuteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RestoreItemAction service

type RestoreItemActionClient interface {
	AppliesTo(ctx context.Context, in *AppliesToRequest, opts ...grpc.CallOption) (*AppliesToResponse, error)
	Execute(ctx context.Context, in *RestoreItemActionExecuteRequest, opts ...grpc.CallOption) (*RestoreItemActionExecuteResponse, error)
}

type restoreItemActionClient struct {
	cc *grpc.ClientConn
}

func NewRestoreItemActionClient(cc *grpc.ClientConn) RestoreItemActionClient {
	return &restoreItemActionClient{cc}
}

func (c *restoreItemActionClient) AppliesTo(ctx context.Context, in *AppliesToRequest, opts ...grpc.CallOption) (*AppliesToResponse, error) {
	out := new(AppliesToResponse)
	err := grpc.Invoke(ctx, "/generated.RestoreItemAction/AppliesTo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restoreItemActionClient) Execute(ctx context.Context, in *RestoreItemActionExecuteRequest, opts ...grpc.CallOption) (*RestoreItemActionExecuteResponse, error) {
	out := new(RestoreItemActionExecuteResponse)
	err := grpc.Invoke(ctx, "/generated.RestoreItemAction/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RestoreItemAction service

type RestoreItemActionServer interface {
	AppliesTo(context.Context, *AppliesToRequest) (*AppliesToResponse, error)
	Execute(context.Context, *RestoreItemActionExecuteRequest) (*RestoreItemActionExecuteResponse, error)
}

func RegisterRestoreItemActionServer(s *grpc.Server, srv RestoreItemActionServer) {
	s.RegisterService(&_RestoreItemAction_serviceDesc, srv)
}

func _RestoreItemAction_AppliesTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppliesToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestoreItemActionServer).AppliesTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.RestoreItemAction/AppliesTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestoreItemActionServer).AppliesTo(ctx, req.(*AppliesToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestoreItemAction_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreItemActionExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestoreItemActionServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.RestoreItemAction/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestoreItemActionServer).Execute(ctx, req.(*RestoreItemActionExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestoreItemAction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.RestoreItemAction",
	HandlerType: (*RestoreItemActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppliesTo",
			Handler:    _RestoreItemAction_AppliesTo_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _RestoreItemAction_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RestoreItemAction.proto",
}

func init() { proto.RegisterFile("RestoreItemAction.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x5b, 0x4a, 0xc3, 0x40,
	0x14, 0x65, 0xda, 0xd2, 0x9a, 0xdb, 0xa2, 0x38, 0x1f, 0x3a, 0x44, 0xc5, 0xd0, 0x0f, 0x09, 0x0a,
	0xf9, 0xa8, 0x2b, 0xa8, 0x60, 0xa5, 0xbf, 0xa3, 0x1b, 0x48, 0x93, 0x6b, 0x1d, 0x9a, 0x64, 0xc6,
	0x79, 0x80, 0xbb, 0x70, 0x0d, 0xae, 0xc4, 0xad, 0x49, 0x1e, 0x86, 0xd8, 0xaa, 0xf4, 0x2f, 0xf7,
	0xdc, 0x73, 0x72, 0xce, 0x99, 0x0b, 0xa7, 0x1c, 0x8d, 0x95, 0x1a, 0x97, 0x16, 0xf3, 0x79, 0x62,
	0x85, 0x2c, 0x22, 0xa5, 0xa5, 0x95, 0xd4, 0x5b, 0x63, 0x81, 0x3a, 0xb6, 0x98, 0xfa, 0x93, 0xc7,
	0x97, 0x58, 0x63, 0x5a, 0x2f, 0xa6, 0xef, 0x04, 0x2e, 0x77, 0x44, 0xf7, 0x6f, 0x98, 0x38, 0x8b,
	0x1c, 0x5f, 0x1d, 0x1a, 0x4b, 0x4f, 0x60, 0xa8, 0x32, 0xb7, 0x16, 0x05, 0x23, 0x01, 0x09, 0x3d,
	0xde, 0x4c, 0x94, 0xc2, 0x40, 0x58, 0xcc, 0x59, 0x2f, 0x20, 0xe1, 0x84, 0x57, 0xdf, 0x94, 0xc1,
	0x48, 0xd7, 0xbf, 0x63, 0xfd, 0x0a, 0xfe, 0x1e, 0xe9, 0x15, 0x1c, 0x96, 0x8c, 0x85, 0x96, 0xf9,
	0x5d, 0x9c, 0x6c, 0x9c, 0x62, 0x83, 0x8a, 0xb0, 0x85, 0x4e, 0x3f, 0x08, 0x04, 0x7f, 0x27, 0x32,
	0x4a, 0x16, 0x06, 0x5b, 0x6b, 0xd2, 0xb1, 0x7e, 0x80, 0xa3, 0x38, 0x4d, 0x45, 0x49, 0x8f, 0xb3,
	0x52, 0x6a, 0x58, 0x2f, 0xe8, 0x87, 0xe3, 0xd9, 0x45, 0xd4, 0xb6, 0x8f, 0x38, 0x1a, 0xe9, 0x74,
	0x82, 0xcb, 0x14, 0x0b, 0x2b, 0x9e, 0x05, 0x6a, 0xbe, 0xad, 0xa2, 0x01, 0x8c, 0xcd, 0x46, 0x28,
	0xde, 0xe9, 0x71, 0xc0, 0xbb, 0xd0, 0xec, 0x93, 0xc0, 0xf1, 0x4e, 0x46, 0xba, 0x00, 0x6f, 0xae,
	0x54, 0x26, 0xd0, 0x3c, 0x49, 0x7a, 0xd6, 0x31, 0x6d, 0xd1, 0xe6, 0x45, 0xfd, 0xf3, 0xdf, 0x97,
	0x4d, 0xb9, 0x15, 0x8c, 0x9a, 0xbe, 0xf4, 0xfa, 0x67, 0xf4, 0xff, 0xce, 0xe4, 0xdf, 0xec, 0xc5,
	0xad, 0x3d, 0x56, 0xc3, 0xea, 0xfc, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xe2, 0xba,
	0x19, 0x32, 0x02, 0x00, 0x00,
}
