// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpc/product/product.proto

package product

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProductInfo struct {
	// 小駝峰＿
	ProductID  string  `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Brand      string  `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name       string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Old        string  `protobuf:"bytes,4,opt,name=old,proto3" json:"old,omitempty"`
	Label      string  `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Color      string  `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	Sale_Price float64 `protobuf:"fixed64,7,opt,name=sale_Price,json=salePrice,proto3" json:"sale_Price,omitempty"`
	Price      float64 `protobuf:"fixed64,8,opt,name=price,proto3" json:"price,omitempty"`
	Size_      string  `protobuf:"bytes,9,opt,name=size,proto3" json:"size,omitempty"`
	Sum        int32   `protobuf:"varint,10,opt,name=sum,proto3" json:"sum,omitempty"`
	UpdateId   string  `protobuf:"bytes,11,opt,name=updateId,proto3" json:"updateId,omitempty"`
	DeleteId   string  `protobuf:"bytes,12,opt,name=deleteId,proto3" json:"deleteId,omitempty"`
}

func (m *ProductInfo) Reset()         { *m = ProductInfo{} }
func (m *ProductInfo) String() string { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()    {}
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b01c3841a6bb5008, []int{0}
}
func (m *ProductInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProductInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProductInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProductInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductInfo.Merge(m, src)
}
func (m *ProductInfo) XXX_Size() int {
	return m.Size()
}
func (m *ProductInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProductInfo proto.InternalMessageInfo

func (m *ProductInfo) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *ProductInfo) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *ProductInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductInfo) GetOld() string {
	if m != nil {
		return m.Old
	}
	return ""
}

func (m *ProductInfo) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ProductInfo) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *ProductInfo) GetSale_Price() float64 {
	if m != nil {
		return m.Sale_Price
	}
	return 0
}

func (m *ProductInfo) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductInfo) GetSize_() string {
	if m != nil {
		return m.Size_
	}
	return ""
}

func (m *ProductInfo) GetSum() int32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *ProductInfo) GetUpdateId() string {
	if m != nil {
		return m.UpdateId
	}
	return ""
}

func (m *ProductInfo) GetDeleteId() string {
	if m != nil {
		return m.DeleteId
	}
	return ""
}

//
type StatusReply struct {
	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *StatusReply) Reset()         { *m = StatusReply{} }
func (m *StatusReply) String() string { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()    {}
func (*StatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b01c3841a6bb5008, []int{1}
}
func (m *StatusReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReply.Merge(m, src)
}
func (m *StatusReply) XXX_Size() int {
	return m.Size()
}
func (m *StatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReply proto.InternalMessageInfo

func (m *StatusReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StatusReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductInfo)(nil), "product.ProductInfo")
	proto.RegisterType((*StatusReply)(nil), "product.StatusReply")
}

func init() { proto.RegisterFile("grpc/product/product.proto", fileDescriptor_b01c3841a6bb5008) }

var fileDescriptor_b01c3841a6bb5008 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3d, 0x6e, 0xea, 0x40,
	0x10, 0xf6, 0x02, 0x06, 0x3c, 0xbc, 0xe2, 0x69, 0x45, 0xb1, 0x42, 0xef, 0x59, 0x88, 0x8a, 0x8a,
	0x48, 0x41, 0xca, 0x01, 0x22, 0x1a, 0x77, 0xc8, 0x28, 0x75, 0xb4, 0x78, 0x37, 0x04, 0x69, 0xf1,
	0x5a, 0x6b, 0xbb, 0x48, 0x4e, 0x91, 0x53, 0xe4, 0x2c, 0x29, 0x29, 0x53, 0x46, 0x50, 0xe5, 0x16,
	0xd1, 0xcc, 0xe2, 0x40, 0x4b, 0xe5, 0xef, 0xc7, 0x9f, 0x66, 0xe6, 0xb3, 0x61, 0xb4, 0x71, 0x45,
	0x76, 0x53, 0x38, 0xab, 0xea, 0xac, 0x6a, 0x9e, 0xb3, 0xc2, 0xd9, 0xca, 0xf2, 0xde, 0x89, 0x4e,
	0xde, 0x5b, 0x30, 0x58, 0x7a, 0x9c, 0xe4, 0x4f, 0x96, 0xff, 0x83, 0xe8, 0x64, 0x25, 0x0b, 0xc1,
	0xc6, 0x6c, 0x1a, 0xa5, 0x67, 0x81, 0x0f, 0x21, 0x5c, 0x3b, 0x99, 0x2b, 0xd1, 0x22, 0xc7, 0x13,
	0xce, 0xa1, 0x93, 0xcb, 0x9d, 0x16, 0x6d, 0x12, 0x09, 0xf3, 0xbf, 0xd0, 0xb6, 0x46, 0x89, 0x0e,
	0x49, 0x08, 0x31, 0x6b, 0xe4, 0x5a, 0x1b, 0x11, 0xfa, 0x2c, 0x11, 0x54, 0x33, 0x6b, 0xac, 0x13,
	0x5d, 0xaf, 0x12, 0xe1, 0xff, 0x01, 0x4a, 0x69, 0xf4, 0xe3, 0xd2, 0x6d, 0x33, 0x2d, 0x7a, 0x63,
	0x36, 0x65, 0x69, 0x84, 0x0a, 0x09, 0x18, 0x2a, 0xc8, 0xe9, 0x93, 0xe3, 0x09, 0xae, 0x51, 0x6e,
	0x5f, 0xb5, 0x88, 0xfc, 0x1a, 0x88, 0x71, 0x8d, 0xb2, 0xde, 0x09, 0x18, 0xb3, 0x69, 0x98, 0x22,
	0xe4, 0x23, 0xe8, 0xd7, 0x85, 0x92, 0x95, 0x4e, 0x94, 0x18, 0xd0, 0x9b, 0xbf, 0x1c, 0x3d, 0xa5,
	0x8d, 0x26, 0xef, 0x8f, 0xf7, 0x1a, 0x3e, 0x99, 0xc3, 0x60, 0x55, 0xc9, 0xaa, 0x2e, 0x53, 0x5d,
	0x98, 0x17, 0x1c, 0x96, 0x59, 0xa5, 0xa9, 0xa2, 0x30, 0x25, 0x8c, 0xc3, 0x76, 0xe5, 0xe6, 0xd4,
	0x0d, 0xc2, 0xdb, 0x6f, 0x06, 0x4d, 0xd3, 0xfc, 0x0e, 0xba, 0x49, 0x5e, 0x6a, 0x57, 0xf1, 0xe1,
	0xac, 0xf9, 0x18, 0x17, 0xcd, 0x8f, 0xce, 0xea, 0xc5, 0x9c, 0x49, 0x80, 0xb9, 0x95, 0x96, 0x2e,
	0x7b, 0xbe, 0x3e, 0xf7, 0x40, 0x87, 0x5d, 0x9f, 0x5b, 0xd0, 0xd1, 0xd7, 0xe5, 0xee, 0xc5, 0xc7,
	0x21, 0x66, 0xfb, 0x43, 0xcc, 0xbe, 0x0e, 0x31, 0x7b, 0x3b, 0xc6, 0xc1, 0xfe, 0x18, 0x07, 0x9f,
	0xc7, 0x38, 0x58, 0x77, 0xe9, 0x9f, 0x9b, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x40, 0x08, 0x78,
	0x57, 0x91, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductClient interface {
	// 插入接口
	Insert(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*StatusReply, error)
	Search(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*StatusReply, error)
	Update(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*StatusReply, error)
	Delete(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*StatusReply, error)
}

type productClient struct {
	cc *grpc.ClientConn
}

func NewProductClient(cc *grpc.ClientConn) ProductClient {
	return &productClient{cc}
}

func (c *productClient) Insert(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/product.product/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Search(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/product.product/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Update(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/product.product/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Delete(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/product.product/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
type ProductServer interface {
	// 插入接口
	Insert(context.Context, *ProductInfo) (*StatusReply, error)
	Search(context.Context, *ProductInfo) (*StatusReply, error)
	Update(context.Context, *ProductInfo) (*StatusReply, error)
	Delete(context.Context, *ProductInfo) (*StatusReply, error)
}

// UnimplementedProductServer can be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (*UnimplementedProductServer) Insert(ctx context.Context, req *ProductInfo) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (*UnimplementedProductServer) Search(ctx context.Context, req *ProductInfo) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedProductServer) Update(ctx context.Context, req *ProductInfo) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedProductServer) Delete(ctx context.Context, req *ProductInfo) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterProductServer(s *grpc.Server, srv ProductServer) {
	s.RegisterService(&_Product_serviceDesc, srv)
}

func _Product_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.product/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Insert(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.product/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Search(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.product/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Update(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.product/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Delete(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Product_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Product_Insert_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Product_Search_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Product_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Product_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/product/product.proto",
}

func (m *ProductInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProductInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProductInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeleteId) > 0 {
		i -= len(m.DeleteId)
		copy(dAtA[i:], m.DeleteId)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.DeleteId)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.UpdateId) > 0 {
		i -= len(m.UpdateId)
		copy(dAtA[i:], m.UpdateId)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.UpdateId)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Sum != 0 {
		i = encodeVarintProduct(dAtA, i, uint64(m.Sum))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Size_) > 0 {
		i -= len(m.Size_)
		copy(dAtA[i:], m.Size_)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.Size_)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Price != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Price))))
		i--
		dAtA[i] = 0x41
	}
	if m.Sale_Price != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Sale_Price))))
		i--
		dAtA[i] = 0x39
	}
	if len(m.Color) > 0 {
		i -= len(m.Color)
		copy(dAtA[i:], m.Color)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.Color)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Old) > 0 {
		i -= len(m.Old)
		copy(dAtA[i:], m.Old)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.Old)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Brand) > 0 {
		i -= len(m.Brand)
		copy(dAtA[i:], m.Brand)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.Brand)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ProductID) > 0 {
		i -= len(m.ProductID)
		copy(dAtA[i:], m.ProductID)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.ProductID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StatusReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintProduct(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintProduct(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProduct(dAtA []byte, offset int, v uint64) int {
	offset -= sovProduct(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProductInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProductID)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	l = len(m.Brand)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	l = len(m.Old)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	l = len(m.Color)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	if m.Sale_Price != 0 {
		n += 9
	}
	if m.Price != 0 {
		n += 9
	}
	l = len(m.Size_)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	if m.Sum != 0 {
		n += 1 + sovProduct(uint64(m.Sum))
	}
	l = len(m.UpdateId)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	l = len(m.DeleteId)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	return n
}

func (m *StatusReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovProduct(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovProduct(uint64(l))
	}
	return n
}

func sovProduct(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProduct(x uint64) (n int) {
	return sovProduct(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProductInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProduct
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProductInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProductInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProductID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Brand", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Brand = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Old", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Old = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Color", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Color = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sale_Price", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Sale_Price = float64(math.Float64frombits(v))
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Price = float64(math.Float64frombits(v))
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Size_ = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			m.Sum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sum |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdateId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeleteId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProduct(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProduct
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProduct
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProduct
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProduct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProduct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProduct(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProduct
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProduct
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProduct(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProduct
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProduct
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProduct
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProduct
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProduct
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProduct        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProduct          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProduct = fmt.Errorf("proto: unexpected end of group")
)
