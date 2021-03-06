// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadatosDisplay.proto

package proto

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

type Mensaje struct {
	Mensaje              string   `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mensaje) Reset()         { *m = Mensaje{} }
func (m *Mensaje) String() string { return proto.CompactTextString(m) }
func (*Mensaje) ProtoMessage()    {}
func (*Mensaje) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7124339a030319b, []int{0}
}

func (m *Mensaje) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mensaje.Unmarshal(m, b)
}
func (m *Mensaje) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mensaje.Marshal(b, m, deterministic)
}
func (m *Mensaje) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mensaje.Merge(m, src)
}
func (m *Mensaje) XXX_Size() int {
	return xxx_messageInfo_Mensaje.Size(m)
}
func (m *Mensaje) XXX_DiscardUnknown() {
	xxx_messageInfo_Mensaje.DiscardUnknown(m)
}

var xxx_messageInfo_Mensaje proto.InternalMessageInfo

func (m *Mensaje) GetMensaje() string {
	if m != nil {
		return m.Mensaje
	}
	return ""
}

type ResponseDepartures struct {
	Departures           []*ObjectDepartures `protobuf:"bytes,1,rep,name=departures,proto3" json:"departures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ResponseDepartures) Reset()         { *m = ResponseDepartures{} }
func (m *ResponseDepartures) String() string { return proto.CompactTextString(m) }
func (*ResponseDepartures) ProtoMessage()    {}
func (*ResponseDepartures) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7124339a030319b, []int{1}
}

func (m *ResponseDepartures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseDepartures.Unmarshal(m, b)
}
func (m *ResponseDepartures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseDepartures.Marshal(b, m, deterministic)
}
func (m *ResponseDepartures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseDepartures.Merge(m, src)
}
func (m *ResponseDepartures) XXX_Size() int {
	return xxx_messageInfo_ResponseDepartures.Size(m)
}
func (m *ResponseDepartures) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseDepartures.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseDepartures proto.InternalMessageInfo

func (m *ResponseDepartures) GetDepartures() []*ObjectDepartures {
	if m != nil {
		return m.Departures
	}
	return nil
}

type ResponseArrivals struct {
	Arrivals             []*ObjectArrivals `protobuf:"bytes,1,rep,name=arrivals,proto3" json:"arrivals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResponseArrivals) Reset()         { *m = ResponseArrivals{} }
func (m *ResponseArrivals) String() string { return proto.CompactTextString(m) }
func (*ResponseArrivals) ProtoMessage()    {}
func (*ResponseArrivals) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7124339a030319b, []int{2}
}

func (m *ResponseArrivals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseArrivals.Unmarshal(m, b)
}
func (m *ResponseArrivals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseArrivals.Marshal(b, m, deterministic)
}
func (m *ResponseArrivals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseArrivals.Merge(m, src)
}
func (m *ResponseArrivals) XXX_Size() int {
	return xxx_messageInfo_ResponseArrivals.Size(m)
}
func (m *ResponseArrivals) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseArrivals.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseArrivals proto.InternalMessageInfo

func (m *ResponseArrivals) GetArrivals() []*ObjectArrivals {
	if m != nil {
		return m.Arrivals
	}
	return nil
}

type ObjectDepartures struct {
	CodigoAvion          string   `protobuf:"bytes,1,opt,name=codigoAvion,proto3" json:"codigoAvion,omitempty"`
	Destino              string   `protobuf:"bytes,2,opt,name=destino,proto3" json:"destino,omitempty"`
	PistaDespegue        string   `protobuf:"bytes,3,opt,name=pistaDespegue,proto3" json:"pistaDespegue,omitempty"`
	HoraDespegue         string   `protobuf:"bytes,4,opt,name=horaDespegue,proto3" json:"horaDespegue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectDepartures) Reset()         { *m = ObjectDepartures{} }
func (m *ObjectDepartures) String() string { return proto.CompactTextString(m) }
func (*ObjectDepartures) ProtoMessage()    {}
func (*ObjectDepartures) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7124339a030319b, []int{3}
}

func (m *ObjectDepartures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectDepartures.Unmarshal(m, b)
}
func (m *ObjectDepartures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectDepartures.Marshal(b, m, deterministic)
}
func (m *ObjectDepartures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectDepartures.Merge(m, src)
}
func (m *ObjectDepartures) XXX_Size() int {
	return xxx_messageInfo_ObjectDepartures.Size(m)
}
func (m *ObjectDepartures) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectDepartures.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectDepartures proto.InternalMessageInfo

func (m *ObjectDepartures) GetCodigoAvion() string {
	if m != nil {
		return m.CodigoAvion
	}
	return ""
}

func (m *ObjectDepartures) GetDestino() string {
	if m != nil {
		return m.Destino
	}
	return ""
}

func (m *ObjectDepartures) GetPistaDespegue() string {
	if m != nil {
		return m.PistaDespegue
	}
	return ""
}

func (m *ObjectDepartures) GetHoraDespegue() string {
	if m != nil {
		return m.HoraDespegue
	}
	return ""
}

type ObjectArrivals struct {
	CodigoAvion          string   `protobuf:"bytes,1,opt,name=codigoAvion,proto3" json:"codigoAvion,omitempty"`
	Proveniente          string   `protobuf:"bytes,2,opt,name=proveniente,proto3" json:"proveniente,omitempty"`
	PistaLlegada         string   `protobuf:"bytes,3,opt,name=pistaLlegada,proto3" json:"pistaLlegada,omitempty"`
	HoraLlegada          string   `protobuf:"bytes,4,opt,name=horaLlegada,proto3" json:"horaLlegada,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectArrivals) Reset()         { *m = ObjectArrivals{} }
func (m *ObjectArrivals) String() string { return proto.CompactTextString(m) }
func (*ObjectArrivals) ProtoMessage()    {}
func (*ObjectArrivals) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7124339a030319b, []int{4}
}

func (m *ObjectArrivals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectArrivals.Unmarshal(m, b)
}
func (m *ObjectArrivals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectArrivals.Marshal(b, m, deterministic)
}
func (m *ObjectArrivals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectArrivals.Merge(m, src)
}
func (m *ObjectArrivals) XXX_Size() int {
	return xxx_messageInfo_ObjectArrivals.Size(m)
}
func (m *ObjectArrivals) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectArrivals.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectArrivals proto.InternalMessageInfo

func (m *ObjectArrivals) GetCodigoAvion() string {
	if m != nil {
		return m.CodigoAvion
	}
	return ""
}

func (m *ObjectArrivals) GetProveniente() string {
	if m != nil {
		return m.Proveniente
	}
	return ""
}

func (m *ObjectArrivals) GetPistaLlegada() string {
	if m != nil {
		return m.PistaLlegada
	}
	return ""
}

func (m *ObjectArrivals) GetHoraLlegada() string {
	if m != nil {
		return m.HoraLlegada
	}
	return ""
}

func init() {
	proto.RegisterType((*Mensaje)(nil), "proto.mensaje")
	proto.RegisterType((*ResponseDepartures)(nil), "proto.responseDepartures")
	proto.RegisterType((*ResponseArrivals)(nil), "proto.responseArrivals")
	proto.RegisterType((*ObjectDepartures)(nil), "proto.ObjectDepartures")
	proto.RegisterType((*ObjectArrivals)(nil), "proto.ObjectArrivals")
}

func init() { proto.RegisterFile("metadatosDisplay.proto", fileDescriptor_c7124339a030319b) }

var fileDescriptor_c7124339a030319b = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xd1, 0x4a, 0xfb, 0x30,
	0x14, 0xc6, 0xd7, 0xff, 0xfe, 0xea, 0x3c, 0xd3, 0x31, 0x02, 0xba, 0xea, 0x55, 0x89, 0x5e, 0xec,
	0x6a, 0xe0, 0x44, 0xc4, 0xcb, 0x42, 0xbd, 0x10, 0x14, 0xa5, 0x6f, 0x90, 0xad, 0x87, 0x9a, 0xd1,
	0x25, 0x21, 0xc9, 0x2a, 0xbe, 0x89, 0xf8, 0xb4, 0xd2, 0x36, 0xed, 0xda, 0x7a, 0xe3, 0x55, 0x72,
	0xbe, 0xfc, 0xf2, 0xe5, 0x3b, 0x39, 0x70, 0xbe, 0x45, 0xcb, 0x12, 0x66, 0xa5, 0x89, 0xb8, 0x51,
	0x19, 0xfb, 0x5c, 0x28, 0x2d, 0xad, 0x24, 0x07, 0xe5, 0x42, 0xaf, 0xe0, 0x68, 0x8b, 0xc2, 0xb0,
	0x0d, 0x12, 0xbf, 0xd9, 0xfa, 0x5e, 0xe0, 0xcd, 0x8f, 0xe3, 0xba, 0xa4, 0x2f, 0x40, 0x34, 0x1a,
	0x25, 0x85, 0xc1, 0x08, 0x15, 0xd3, 0x76, 0xa7, 0xd1, 0x90, 0x7b, 0x80, 0xa4, 0xa9, 0x7c, 0x2f,
	0x18, 0xce, 0xc7, 0xcb, 0x59, 0xe5, 0xbe, 0x78, 0x5d, 0x6d, 0x70, 0x6d, 0xf7, 0x70, 0xdc, 0x42,
	0xe9, 0x23, 0x4c, 0x6b, 0xbb, 0x50, 0x6b, 0x9e, 0xb3, 0xcc, 0x90, 0x1b, 0x18, 0x31, 0xb7, 0x77,
	0x56, 0x67, 0x1d, 0xab, 0x1a, 0x8c, 0x1b, 0x8c, 0x7e, 0x79, 0x30, 0xed, 0xbf, 0x43, 0x02, 0x18,
	0xaf, 0x65, 0xc2, 0x53, 0x19, 0xe6, 0x5c, 0x0a, 0xd7, 0x48, 0x5b, 0x2a, 0xda, 0x4c, 0xd0, 0x58,
	0x2e, 0xa4, 0xff, 0xaf, 0x6a, 0xd3, 0x95, 0xe4, 0x1a, 0x4e, 0x15, 0x37, 0x96, 0x45, 0x68, 0x14,
	0xa6, 0x3b, 0xf4, 0x87, 0xe5, 0x79, 0x57, 0x24, 0x14, 0x4e, 0xde, 0xa5, 0xde, 0x43, 0xff, 0x4b,
	0xa8, 0xa3, 0xd1, 0x6f, 0x0f, 0x26, 0xdd, 0xdc, 0x7f, 0x08, 0x16, 0xc0, 0x58, 0x69, 0x99, 0xa3,
	0xe0, 0x28, 0x2c, 0xba, 0x70, 0x6d, 0xa9, 0x78, 0xba, 0xcc, 0xf2, 0x9c, 0x61, 0xca, 0x12, 0xe6,
	0xf2, 0x75, 0xb4, 0xc2, 0xa5, 0x88, 0x52, 0x23, 0x55, 0xba, 0xb6, 0xb4, 0xfc, 0x80, 0x61, 0xf8,
	0xf6, 0x44, 0x1e, 0x00, 0x5a, 0xff, 0x36, 0x71, 0xbf, 0xed, 0x46, 0x7e, 0x79, 0xe1, 0xea, 0xdf,
	0x73, 0xa7, 0x03, 0x72, 0x07, 0xa3, 0xa6, 0xaf, 0xfe, 0xc5, 0x59, 0xef, 0x62, 0x0d, 0xd2, 0xc1,
	0xea, 0xb0, 0x3c, 0xb9, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x1e, 0x77, 0x61, 0x93, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	Departures(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*ResponseDepartures, error)
	Arrivals(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*ResponseArrivals, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Departures(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*ResponseDepartures, error) {
	out := new(ResponseDepartures)
	err := c.cc.Invoke(ctx, "/proto.API/Departures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Arrivals(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*ResponseArrivals, error) {
	out := new(ResponseArrivals)
	err := c.cc.Invoke(ctx, "/proto.API/Arrivals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	Departures(context.Context, *Mensaje) (*ResponseDepartures, error)
	Arrivals(context.Context, *Mensaje) (*ResponseArrivals, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Departures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Departures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/Departures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Departures(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Arrivals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Arrivals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/Arrivals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Arrivals(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Departures",
			Handler:    _API_Departures_Handler,
		},
		{
			MethodName: "Arrivals",
			Handler:    _API_Arrivals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadatosDisplay.proto",
}
