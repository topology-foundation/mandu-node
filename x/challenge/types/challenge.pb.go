// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mandu/challenge/challenge.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Challenge struct {
	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Challenger       string `protobuf:"bytes,2,opt,name=challenger,proto3" json:"challenger,omitempty"`
	Subscriber       string `protobuf:"bytes,3,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
	Amount           uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	LastActive       uint64 `protobuf:"varint,5,opt,name=last_active,json=lastActive,proto3" json:"last_active,omitempty"`
	ChallengedHashes []byte `protobuf:"bytes,6,opt,name=challenged_hashes,json=challengedHashes,proto3" json:"challenged_hashes,omitempty"`
}

func (m *Challenge) Reset()         { *m = Challenge{} }
func (m *Challenge) String() string { return proto.CompactTextString(m) }
func (*Challenge) ProtoMessage()    {}
func (*Challenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_b69a026757d72cb5, []int{0}
}
func (m *Challenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Challenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Challenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Challenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Challenge.Merge(m, src)
}
func (m *Challenge) XXX_Size() int {
	return m.Size()
}
func (m *Challenge) XXX_DiscardUnknown() {
	xxx_messageInfo_Challenge.DiscardUnknown(m)
}

var xxx_messageInfo_Challenge proto.InternalMessageInfo

func (m *Challenge) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Challenge) GetChallenger() string {
	if m != nil {
		return m.Challenger
	}
	return ""
}

func (m *Challenge) GetSubscriber() string {
	if m != nil {
		return m.Subscriber
	}
	return ""
}

func (m *Challenge) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Challenge) GetLastActive() uint64 {
	if m != nil {
		return m.LastActive
	}
	return 0
}

func (m *Challenge) GetChallengedHashes() []byte {
	if m != nil {
		return m.ChallengedHashes
	}
	return nil
}

func init() {
	proto.RegisterType((*Challenge)(nil), "mandu.challenge.Challenge")
}

func init() { proto.RegisterFile("mandu/challenge/challenge.proto", fileDescriptor_b69a026757d72cb5) }

var fileDescriptor_b69a026757d72cb5 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x4d, 0xcc, 0x4b,
	0x29, 0xd5, 0x4f, 0xce, 0x48, 0xcc, 0xc9, 0x49, 0xcd, 0x4b, 0x4f, 0x45, 0xb0, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0xc1, 0x0a, 0xf4, 0xe0, 0xc2, 0x4a, 0x07, 0x19, 0xb9, 0x38, 0x9d,
	0x61, 0x3c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6,
	0xcc, 0x14, 0x21, 0x39, 0x2e, 0x2e, 0xb8, 0xd2, 0x22, 0x09, 0x26, 0xb0, 0x38, 0x92, 0x08, 0x48,
	0xbe, 0xb8, 0x34, 0xa9, 0x38, 0xb9, 0x28, 0x33, 0x29, 0xb5, 0x48, 0x82, 0x19, 0x22, 0x8f, 0x10,
	0x11, 0x12, 0xe3, 0x62, 0x4b, 0xcc, 0xcd, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x51, 0x60, 0xd4, 0x60,
	0x09, 0x82, 0xf2, 0x84, 0xe4, 0xb9, 0xb8, 0x73, 0x12, 0x8b, 0x4b, 0xe2, 0x13, 0x93, 0x4b, 0x32,
	0xcb, 0x52, 0x25, 0x58, 0xc1, 0x92, 0x5c, 0x20, 0x21, 0x47, 0xb0, 0x88, 0x90, 0x36, 0x97, 0x20,
	0xdc, 0x9a, 0x94, 0xf8, 0x8c, 0xc4, 0xe2, 0x8c, 0xd4, 0x62, 0x09, 0x36, 0x05, 0x46, 0x0d, 0x9e,
	0x20, 0x01, 0x84, 0x84, 0x07, 0x58, 0xdc, 0xc9, 0xf0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f,
	0xe5, 0x18, 0xa2, 0xc4, 0x21, 0xe1, 0x51, 0x81, 0x14, 0x22, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49,
	0x6c, 0xe0, 0xe0, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x46, 0x2a, 0x46, 0x29, 0x31, 0x01,
	0x00, 0x00,
}

func (m *Challenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Challenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Challenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChallengedHashes) > 0 {
		i -= len(m.ChallengedHashes)
		copy(dAtA[i:], m.ChallengedHashes)
		i = encodeVarintChallenge(dAtA, i, uint64(len(m.ChallengedHashes)))
		i--
		dAtA[i] = 0x32
	}
	if m.LastActive != 0 {
		i = encodeVarintChallenge(dAtA, i, uint64(m.LastActive))
		i--
		dAtA[i] = 0x28
	}
	if m.Amount != 0 {
		i = encodeVarintChallenge(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Subscriber) > 0 {
		i -= len(m.Subscriber)
		copy(dAtA[i:], m.Subscriber)
		i = encodeVarintChallenge(dAtA, i, uint64(len(m.Subscriber)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Challenger) > 0 {
		i -= len(m.Challenger)
		copy(dAtA[i:], m.Challenger)
		i = encodeVarintChallenge(dAtA, i, uint64(len(m.Challenger)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintChallenge(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChallenge(dAtA []byte, offset int, v uint64) int {
	offset -= sovChallenge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Challenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovChallenge(uint64(l))
	}
	l = len(m.Challenger)
	if l > 0 {
		n += 1 + l + sovChallenge(uint64(l))
	}
	l = len(m.Subscriber)
	if l > 0 {
		n += 1 + l + sovChallenge(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovChallenge(uint64(m.Amount))
	}
	if m.LastActive != 0 {
		n += 1 + sovChallenge(uint64(m.LastActive))
	}
	l = len(m.ChallengedHashes)
	if l > 0 {
		n += 1 + l + sovChallenge(uint64(l))
	}
	return n
}

func sovChallenge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChallenge(x uint64) (n int) {
	return sovChallenge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Challenge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChallenge
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
			return fmt.Errorf("proto: Challenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Challenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenge
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
				return ErrInvalidLengthChallenge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Challenger", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenge
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
				return ErrInvalidLengthChallenge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Challenger = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenge
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
				return ErrInvalidLengthChallenge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastActive", wireType)
			}
			m.LastActive = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastActive |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengedHashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChallenge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChallenge
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChallenge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengedHashes = append(m.ChallengedHashes[:0], dAtA[iNdEx:postIndex]...)
			if m.ChallengedHashes == nil {
				m.ChallengedHashes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChallenge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChallenge
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
func skipChallenge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChallenge
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
					return 0, ErrIntOverflowChallenge
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
					return 0, ErrIntOverflowChallenge
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
				return 0, ErrInvalidLengthChallenge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChallenge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChallenge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChallenge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChallenge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChallenge = fmt.Errorf("proto: unexpected end of group")
)
