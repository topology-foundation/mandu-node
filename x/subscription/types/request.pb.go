// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mandu/subscription/request.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

type Request_Status int32

const (
	Request_UNSPECIFIED Request_Status = 0
)

var Request_Status_name = map[int32]string{
	0: "UNSPECIFIED",
}

var Request_Status_value = map[string]int32{
	"UNSPECIFIED": 0,
}

func (x Request_Status) String() string {
	return proto.EnumName(Request_Status_name, int32(x))
}

func (Request_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a305bc2d08e049, []int{0, 0}
}

type Request struct {
	Id              string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Requester       string         `protobuf:"bytes,2,opt,name=requester,proto3" json:"requester,omitempty"`
	CroId           string         `protobuf:"bytes,3,opt,name=cro_id,json=croId,proto3" json:"cro_id,omitempty"`
	SubscriptionIds []string       `protobuf:"bytes,4,rep,name=subscription_ids,json=subscriptionIds,proto3" json:"subscription_ids,omitempty"`
	Status          Request_Status `protobuf:"varint,5,opt,name=status,proto3,enum=mandu.subscription.Request_Status" json:"status,omitempty"`
	InitialAmount   uint64         `protobuf:"varint,6,opt,name=initial_amount,json=initialAmount,proto3" json:"initial_amount,omitempty"`
	AvailableAmount uint64         `protobuf:"varint,7,opt,name=available_amount,json=availableAmount,proto3" json:"available_amount,omitempty"`
	StartBlock      uint64         `protobuf:"varint,8,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	EndBlock        uint64         `protobuf:"varint,9,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a305bc2d08e049, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetRequester() string {
	if m != nil {
		return m.Requester
	}
	return ""
}

func (m *Request) GetCroId() string {
	if m != nil {
		return m.CroId
	}
	return ""
}

func (m *Request) GetSubscriptionIds() []string {
	if m != nil {
		return m.SubscriptionIds
	}
	return nil
}

func (m *Request) GetStatus() Request_Status {
	if m != nil {
		return m.Status
	}
	return Request_UNSPECIFIED
}

func (m *Request) GetInitialAmount() uint64 {
	if m != nil {
		return m.InitialAmount
	}
	return 0
}

func (m *Request) GetAvailableAmount() uint64 {
	if m != nil {
		return m.AvailableAmount
	}
	return 0
}

func (m *Request) GetStartBlock() uint64 {
	if m != nil {
		return m.StartBlock
	}
	return 0
}

func (m *Request) GetEndBlock() uint64 {
	if m != nil {
		return m.EndBlock
	}
	return 0
}

func init() {
	proto.RegisterEnum("mandu.subscription.Request_Status", Request_Status_name, Request_Status_value)
	proto.RegisterType((*Request)(nil), "mandu.subscription.Request")
}

func init() {
	proto.RegisterFile("mandu/subscription/request.proto", fileDescriptor_77a305bc2d08e049)
}

var fileDescriptor_77a305bc2d08e049 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xdb, 0x6e, 0xeb, 0xd6, 0x8c, 0x77, 0x1b, 0x81, 0x41, 0x5e, 0xd4, 0x5a, 0x26, 0x83,
	0x7a, 0xe9, 0x40, 0x0f, 0x9e, 0x3c, 0x38, 0x9d, 0xd0, 0x8b, 0x48, 0x87, 0x17, 0x2f, 0x25, 0x6d,
	0x02, 0x06, 0x6b, 0x53, 0x93, 0x54, 0xf4, 0x5b, 0xf8, 0xb1, 0x3c, 0xee, 0xe8, 0x51, 0xd6, 0x2f,
	0x22, 0xcb, 0xba, 0x39, 0xc1, 0xeb, 0xef, 0xf9, 0x25, 0xff, 0x07, 0x1e, 0x70, 0xa4, 0x78, 0x91,
	0x3e, 0x60, 0x96, 0x4f, 0x64, 0x99, 0xc8, 0x54, 0xb0, 0x42, 0x31, 0x9e, 0x4f, 0x04, 0x7d, 0x2e,
	0xa9, 0x54, 0x41, 0x21, 0xb8, 0xe2, 0x70, 0xb8, 0x91, 0x82, 0x5d, 0x69, 0x54, 0x59, 0xa0, 0x1d,
	0xad, 0x45, 0xd8, 0x03, 0x16, 0x23, 0xc8, 0xf4, 0x4c, 0xdf, 0x89, 0x2c, 0x46, 0xe0, 0x3e, 0x70,
	0xea, 0x3f, 0xa8, 0x40, 0x96, 0xc6, 0x3f, 0x00, 0x0e, 0x81, 0x9d, 0x0a, 0x1e, 0x33, 0x82, 0x1a,
	0x3a, 0x6a, 0xa5, 0x82, 0x87, 0x04, 0x1e, 0x83, 0xc1, 0xee, 0x81, 0x98, 0x11, 0x89, 0x9a, 0x5e,
	0xc3, 0x77, 0xa2, 0xfe, 0x2e, 0x0f, 0x89, 0x84, 0xe7, 0xc0, 0x96, 0x0a, 0xab, 0x52, 0xa2, 0x96,
	0x67, 0xfa, 0xbd, 0x93, 0x71, 0xf0, 0x67, 0xc7, 0xa0, 0xee, 0x17, 0xcc, 0xb5, 0x1c, 0xd5, 0x8f,
	0xe0, 0x18, 0xf4, 0x58, 0xce, 0x14, 0xc3, 0x59, 0x8c, 0x9f, 0x78, 0x99, 0x2b, 0x64, 0x7b, 0xa6,
	0xdf, 0x8c, 0xfe, 0xd5, 0xf4, 0x42, 0xc3, 0x55, 0x21, 0xfc, 0x82, 0x59, 0x86, 0x93, 0x8c, 0x6e,
	0xc4, 0xb6, 0x16, 0xfb, 0x5b, 0x5e, 0xab, 0x87, 0xa0, 0x2b, 0x15, 0x16, 0x2a, 0x4e, 0x32, 0x9e,
	0x3e, 0xa2, 0x8e, 0xb6, 0x80, 0x46, 0xd3, 0x15, 0x81, 0x7b, 0xc0, 0xa1, 0x39, 0xa9, 0x63, 0x47,
	0xc7, 0x1d, 0x9a, 0x13, 0x1d, 0x8e, 0xfe, 0x03, 0x7b, 0xdd, 0x10, 0xf6, 0x41, 0xf7, 0xee, 0x66,
	0x7e, 0x3b, 0xbb, 0x0c, 0xaf, 0xc3, 0xd9, 0xd5, 0xc0, 0x98, 0x9e, 0x7d, 0x2c, 0x5d, 0x73, 0xb1,
	0x74, 0xcd, 0xaf, 0xa5, 0x6b, 0xbe, 0x57, 0xae, 0xb1, 0xa8, 0x5c, 0xe3, 0xb3, 0x72, 0x8d, 0xfb,
	0x83, 0xed, 0x76, 0xaf, 0xbf, 0xd7, 0x53, 0x6f, 0x05, 0x95, 0x89, 0xad, 0xc7, 0x3b, 0xfd, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x9a, 0xbd, 0xc9, 0x27, 0xe3, 0x01, 0x00, 0x00,
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndBlock != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.EndBlock))
		i--
		dAtA[i] = 0x48
	}
	if m.StartBlock != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x40
	}
	if m.AvailableAmount != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.AvailableAmount))
		i--
		dAtA[i] = 0x38
	}
	if m.InitialAmount != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.InitialAmount))
		i--
		dAtA[i] = 0x30
	}
	if m.Status != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SubscriptionIds) > 0 {
		for iNdEx := len(m.SubscriptionIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SubscriptionIds[iNdEx])
			copy(dAtA[i:], m.SubscriptionIds[iNdEx])
			i = encodeVarintRequest(dAtA, i, uint64(len(m.SubscriptionIds[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.CroId) > 0 {
		i -= len(m.CroId)
		copy(dAtA[i:], m.CroId)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.CroId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Requester) > 0 {
		i -= len(m.Requester)
		copy(dAtA[i:], m.Requester)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Requester)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.Requester)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.CroId)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if len(m.SubscriptionIds) > 0 {
		for _, s := range m.SubscriptionIds {
			l = len(s)
			n += 1 + l + sovRequest(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovRequest(uint64(m.Status))
	}
	if m.InitialAmount != 0 {
		n += 1 + sovRequest(uint64(m.InitialAmount))
	}
	if m.AvailableAmount != 0 {
		n += 1 + sovRequest(uint64(m.AvailableAmount))
	}
	if m.StartBlock != 0 {
		n += 1 + sovRequest(uint64(m.StartBlock))
	}
	if m.EndBlock != 0 {
		n += 1 + sovRequest(uint64(m.EndBlock))
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requester", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requester = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CroId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CroId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubscriptionIds = append(m.SubscriptionIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Request_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialAmount", wireType)
			}
			m.InitialAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvailableAmount", wireType)
			}
			m.AvailableAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AvailableAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
			}
			m.StartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			m.EndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)
