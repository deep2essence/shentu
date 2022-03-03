// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/interview/v1alpha1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
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

// MsgCreatePool defines the attributes of a create-pool transaction.
type MsgLockUser struct {
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	Id   uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
}

func (m *MsgLockUser) Reset()         { *m = MsgLockUser{} }
func (m *MsgLockUser) String() string { return proto.CompactTextString(m) }
func (*MsgLockUser) ProtoMessage()    {}
func (*MsgLockUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ed4aaa105dd2890, []int{0}
}
func (m *MsgLockUser) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLockUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLockUser.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLockUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLockUser.Merge(m, src)
}
func (m *MsgLockUser) XXX_Size() int {
	return m.Size()
}
func (m *MsgLockUser) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLockUser.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLockUser proto.InternalMessageInfo

type MsgLockUserResponse struct {
}

func (m *MsgLockUserResponse) Reset()         { *m = MsgLockUserResponse{} }
func (m *MsgLockUserResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLockUserResponse) ProtoMessage()    {}
func (*MsgLockUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ed4aaa105dd2890, []int{1}
}
func (m *MsgLockUserResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLockUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLockUserResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLockUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLockUserResponse.Merge(m, src)
}
func (m *MsgLockUserResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLockUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLockUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLockUserResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgLockUser)(nil), "shentu.interview.v1alpha1.MsgLockUser")
	proto.RegisterType((*MsgLockUserResponse)(nil), "shentu.interview.v1alpha1.MsgLockUserResponse")
}

func init() {
	proto.RegisterFile("shentu/interview/v1alpha1/tx.proto", fileDescriptor_8ed4aaa105dd2890)
}

var fileDescriptor_8ed4aaa105dd2890 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0xce, 0x48, 0xcd,
	0x2b, 0x29, 0xd5, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x2a, 0xcb, 0x4c, 0x2d, 0xd7, 0x2f, 0x33, 0x4c,
	0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x84, 0xa8, 0xd1, 0x83, 0xab, 0xd1, 0x83, 0xa9, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab,
	0xd2, 0x07, 0xb1, 0x20, 0x1a, 0x94, 0xa2, 0xb9, 0xb8, 0x7d, 0x8b, 0xd3, 0x7d, 0xf2, 0x93, 0xb3,
	0x43, 0x8b, 0x53, 0x8b, 0x84, 0x94, 0xb9, 0x58, 0xd2, 0x8a, 0xf2, 0x73, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x9d, 0xf8, 0x3f, 0xdd, 0x93, 0xe7, 0xae, 0x4c, 0xcc, 0xcd, 0xb1, 0x52, 0x02, 0x89,
	0x2a, 0x05, 0x81, 0x25, 0x85, 0x64, 0xb9, 0x98, 0x32, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58,
	0x9c, 0x78, 0x3f, 0xdd, 0x93, 0xe7, 0x84, 0x28, 0xc9, 0x4c, 0x51, 0x0a, 0x62, 0xca, 0x4c, 0xb1,
	0xe2, 0xe8, 0x58, 0x20, 0xcf, 0xf0, 0x62, 0x81, 0x3c, 0x83, 0x92, 0x28, 0x97, 0x30, 0x92, 0xe1,
	0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x46, 0x99, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42,
	0x49, 0x5c, 0x1c, 0x70, 0x7b, 0xd5, 0xf4, 0x70, 0x3a, 0x5c, 0x0f, 0xc9, 0x08, 0x29, 0x3d, 0xe2,
	0xd4, 0xc1, 0xac, 0x72, 0xf2, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0x93, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe4, 0xd4, 0xa2, 0x92,
	0xcc, 0xec, 0xb4, 0xfc, 0xd2, 0xbc, 0x94, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x7d, 0x68, 0x48, 0x57,
	0x20, 0x85, 0x75, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xd4, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x23, 0x58, 0x18, 0x28, 0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	LockUser(ctx context.Context, in *MsgLockUser, opts ...grpc.CallOption) (*MsgLockUserResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) LockUser(ctx context.Context, in *MsgLockUser, opts ...grpc.CallOption) (*MsgLockUserResponse, error) {
	out := new(MsgLockUserResponse)
	err := c.cc.Invoke(ctx, "/shentu.interview.v1alpha1.Msg/LockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	LockUser(context.Context, *MsgLockUser) (*MsgLockUserResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) LockUser(ctx context.Context, req *MsgLockUser) (*MsgLockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockUser not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_LockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLockUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.interview.v1alpha1.Msg/LockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LockUser(ctx, req.(*MsgLockUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shentu.interview.v1alpha1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LockUser",
			Handler:    _Msg_LockUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shentu/interview/v1alpha1/tx.proto",
}

func (m *MsgLockUser) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLockUser) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLockUser) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLockUserResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLockUserResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLockUserResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLockUser) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgLockUserResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLockUser) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLockUser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLockUser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgLockUserResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLockUserResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLockUserResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
