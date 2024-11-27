// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mandu/subscription/deal.proto

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

type Deal_Status int32

const (
	Deal_UNDEFINED   Deal_Status = 0
	Deal_SCHEDULED   Deal_Status = 1
	Deal_INITIALIZED Deal_Status = 2
	Deal_ACTIVE      Deal_Status = 3
	Deal_INACTIVE    Deal_Status = 4
	Deal_CANCELLED   Deal_Status = 5
	Deal_EXPIRED     Deal_Status = 6
)

var Deal_Status_name = map[int32]string{
	0: "UNDEFINED",
	1: "SCHEDULED",
	2: "INITIALIZED",
	3: "ACTIVE",
	4: "INACTIVE",
	5: "CANCELLED",
	6: "EXPIRED",
}

var Deal_Status_value = map[string]int32{
	"UNDEFINED":   0,
	"SCHEDULED":   1,
	"INITIALIZED": 2,
	"ACTIVE":      3,
	"INACTIVE":    4,
	"CANCELLED":   5,
	"EXPIRED":     6,
}

func (x Deal_Status) String() string {
	return proto.EnumName(Deal_Status_name, int32(x))
}

func (Deal_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9de14b4f09b5e255, []int{0, 0}
}

type Deal struct {
	Id              string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Requester       string      `protobuf:"bytes,2,opt,name=requester,proto3" json:"requester,omitempty"`
	CroId           string      `protobuf:"bytes,3,opt,name=cro_id,json=croId,proto3" json:"cro_id,omitempty"`
	SubscriptionIds []string    `protobuf:"bytes,4,rep,name=subscription_ids,json=subscriptionIds,proto3" json:"subscription_ids,omitempty"`
	Status          Deal_Status `protobuf:"varint,5,opt,name=status,proto3,enum=mandu.subscription.Deal_Status" json:"status,omitempty"`
	TotalAmount     uint64      `protobuf:"varint,6,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	AvailableAmount uint64      `protobuf:"varint,7,opt,name=available_amount,json=availableAmount,proto3" json:"available_amount,omitempty"`
	StartBlock      uint64      `protobuf:"varint,8,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	EndBlock        uint64      `protobuf:"varint,9,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
	InitialFrontier []string    `protobuf:"bytes,10,rep,name=initial_frontier,json=initialFrontier,proto3" json:"initial_frontier,omitempty"`
}

func (m *Deal) Reset()         { *m = Deal{} }
func (m *Deal) String() string { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()    {}
func (*Deal) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de14b4f09b5e255, []int{0}
}
func (m *Deal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deal.Merge(m, src)
}
func (m *Deal) XXX_Size() int {
	return m.Size()
}
func (m *Deal) XXX_DiscardUnknown() {
	xxx_messageInfo_Deal.DiscardUnknown(m)
}

var xxx_messageInfo_Deal proto.InternalMessageInfo

func (m *Deal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Deal) GetRequester() string {
	if m != nil {
		return m.Requester
	}
	return ""
}

func (m *Deal) GetCroId() string {
	if m != nil {
		return m.CroId
	}
	return ""
}

func (m *Deal) GetSubscriptionIds() []string {
	if m != nil {
		return m.SubscriptionIds
	}
	return nil
}

func (m *Deal) GetStatus() Deal_Status {
	if m != nil {
		return m.Status
	}
	return Deal_UNDEFINED
}

func (m *Deal) GetTotalAmount() uint64 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *Deal) GetAvailableAmount() uint64 {
	if m != nil {
		return m.AvailableAmount
	}
	return 0
}

func (m *Deal) GetStartBlock() uint64 {
	if m != nil {
		return m.StartBlock
	}
	return 0
}

func (m *Deal) GetEndBlock() uint64 {
	if m != nil {
		return m.EndBlock
	}
	return 0
}

func (m *Deal) GetInitialFrontier() []string {
	if m != nil {
		return m.InitialFrontier
	}
	return nil
}

func init() {
	proto.RegisterEnum("mandu.subscription.Deal_Status", Deal_Status_name, Deal_Status_value)
	proto.RegisterType((*Deal)(nil), "mandu.subscription.Deal")
}

func init() { proto.RegisterFile("mandu/subscription/deal.proto", fileDescriptor_9de14b4f09b5e255) }

var fileDescriptor_9de14b4f09b5e255 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xe3, 0xfc, 0x71, 0xe3, 0x37, 0x5d, 0x63, 0x04, 0x05, 0xc1, 0x36, 0xcf, 0xcb, 0x29,
	0xbd, 0xb8, 0xb0, 0x1d, 0x06, 0xbb, 0xa5, 0xb1, 0xca, 0x04, 0xc1, 0x0c, 0xb7, 0x1d, 0xa3, 0x17,
	0xa3, 0x58, 0x1a, 0x13, 0x73, 0xad, 0x4c, 0x52, 0xc6, 0xf6, 0x2d, 0xf6, 0xb1, 0x76, 0xec, 0xb1,
	0xc7, 0x91, 0x7c, 0x91, 0x61, 0xc5, 0xcd, 0x32, 0xd8, 0x51, 0xbf, 0xe7, 0xf7, 0x8a, 0xf7, 0x85,
	0x07, 0x62, 0xab, 0x56, 0xe5, 0x67, 0x26, 0xeb, 0x73, 0xb3, 0x5e, 0x9a, 0x52, 0xcb, 0x95, 0x95,
	0xaa, 0x3e, 0xe7, 0x82, 0x55, 0xc9, 0x4a, 0x2b, 0xab, 0xd0, 0xe9, 0xa3, 0x91, 0x1c, 0x1a, 0x93,
	0x87, 0x1e, 0xf4, 0x53, 0xc1, 0x2a, 0x74, 0x02, 0x5d, 0xc9, 0xb1, 0x17, 0x7b, 0xd3, 0x20, 0xef,
	0x4a, 0x8e, 0x9e, 0x41, 0xa0, 0xc5, 0xd7, 0xb5, 0x30, 0x56, 0x68, 0xdc, 0x75, 0xf8, 0x2f, 0x40,
	0xa7, 0xe0, 0x97, 0x5a, 0x15, 0x92, 0xe3, 0x9e, 0x8b, 0x06, 0xa5, 0x56, 0x94, 0xa3, 0x33, 0x08,
	0x0f, 0x7f, 0x2f, 0x24, 0x37, 0xb8, 0x1f, 0xf7, 0xa6, 0x41, 0x3e, 0x3e, 0xe4, 0x94, 0x1b, 0xf4,
	0x16, 0x7c, 0x63, 0x99, 0x5d, 0x1b, 0x3c, 0x88, 0xbd, 0xe9, 0xc9, 0xab, 0x49, 0xf2, 0xdf, 0x05,
	0x93, 0x66, 0xb9, 0xe4, 0xca, 0x99, 0x79, 0x3b, 0x81, 0x5e, 0xc2, 0xb1, 0x55, 0x96, 0x55, 0x05,
	0xbb, 0x53, 0xeb, 0xda, 0x62, 0x3f, 0xf6, 0xa6, 0xfd, 0x7c, 0xe4, 0xd8, 0xcc, 0xa1, 0x66, 0x13,
	0xf6, 0x8d, 0xc9, 0x8a, 0x2d, 0x2b, 0xf1, 0xa8, 0x1d, 0x39, 0x6d, 0xbc, 0xe7, 0xad, 0xfa, 0x02,
	0x46, 0xc6, 0x32, 0x6d, 0x8b, 0x65, 0xa5, 0xca, 0x2f, 0x78, 0xe8, 0x2c, 0x70, 0xe8, 0xa2, 0x21,
	0xe8, 0x29, 0x04, 0xa2, 0xe6, 0x6d, 0x1c, 0xb8, 0x78, 0x28, 0x6a, 0xbe, 0x0b, 0xcf, 0x20, 0x94,
	0xb5, 0xb4, 0x92, 0x55, 0xc5, 0x27, 0xad, 0x6a, 0x2b, 0x85, 0xc6, 0xb0, 0x3b, 0xb9, 0xe5, 0x97,
	0x2d, 0x9e, 0xdc, 0x81, 0xbf, 0x3b, 0x04, 0x3d, 0x81, 0xe0, 0x26, 0x4b, 0xc9, 0x25, 0xcd, 0x48,
	0x1a, 0x76, 0x9a, 0xe7, 0xd5, 0xfc, 0x1d, 0x49, 0x6f, 0x16, 0x24, 0x0d, 0x3d, 0x34, 0x86, 0x11,
	0xcd, 0xe8, 0x35, 0x9d, 0x2d, 0xe8, 0x2d, 0x49, 0xc3, 0x2e, 0x02, 0xf0, 0x67, 0xf3, 0x6b, 0xfa,
	0x81, 0x84, 0x3d, 0x74, 0x0c, 0x43, 0x9a, 0xb5, 0xaf, 0x7e, 0x33, 0x39, 0x9f, 0x65, 0x73, 0xb2,
	0x68, 0x26, 0x07, 0x68, 0x04, 0x47, 0xe4, 0xe3, 0x7b, 0x9a, 0x93, 0x34, 0xf4, 0x2f, 0xde, 0xfc,
	0xda, 0x44, 0xde, 0xfd, 0x26, 0xf2, 0x7e, 0x6f, 0x22, 0xef, 0xe7, 0x36, 0xea, 0xdc, 0x6f, 0xa3,
	0xce, 0xc3, 0x36, 0xea, 0xdc, 0x3e, 0xdf, 0xb7, 0xe5, 0xfb, 0xbf, 0x7d, 0xb1, 0x3f, 0x56, 0xc2,
	0x2c, 0x7d, 0xd7, 0x98, 0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x14, 0x66, 0x42, 0x55,
	0x02, 0x00, 0x00,
}

func (m *Deal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InitialFrontier) > 0 {
		for iNdEx := len(m.InitialFrontier) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.InitialFrontier[iNdEx])
			copy(dAtA[i:], m.InitialFrontier[iNdEx])
			i = encodeVarintDeal(dAtA, i, uint64(len(m.InitialFrontier[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.EndBlock != 0 {
		i = encodeVarintDeal(dAtA, i, uint64(m.EndBlock))
		i--
		dAtA[i] = 0x48
	}
	if m.StartBlock != 0 {
		i = encodeVarintDeal(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x40
	}
	if m.AvailableAmount != 0 {
		i = encodeVarintDeal(dAtA, i, uint64(m.AvailableAmount))
		i--
		dAtA[i] = 0x38
	}
	if m.TotalAmount != 0 {
		i = encodeVarintDeal(dAtA, i, uint64(m.TotalAmount))
		i--
		dAtA[i] = 0x30
	}
	if m.Status != 0 {
		i = encodeVarintDeal(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SubscriptionIds) > 0 {
		for iNdEx := len(m.SubscriptionIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SubscriptionIds[iNdEx])
			copy(dAtA[i:], m.SubscriptionIds[iNdEx])
			i = encodeVarintDeal(dAtA, i, uint64(len(m.SubscriptionIds[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.CroId) > 0 {
		i -= len(m.CroId)
		copy(dAtA[i:], m.CroId)
		i = encodeVarintDeal(dAtA, i, uint64(len(m.CroId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Requester) > 0 {
		i -= len(m.Requester)
		copy(dAtA[i:], m.Requester)
		i = encodeVarintDeal(dAtA, i, uint64(len(m.Requester)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDeal(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeal(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Deal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDeal(uint64(l))
	}
	l = len(m.Requester)
	if l > 0 {
		n += 1 + l + sovDeal(uint64(l))
	}
	l = len(m.CroId)
	if l > 0 {
		n += 1 + l + sovDeal(uint64(l))
	}
	if len(m.SubscriptionIds) > 0 {
		for _, s := range m.SubscriptionIds {
			l = len(s)
			n += 1 + l + sovDeal(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovDeal(uint64(m.Status))
	}
	if m.TotalAmount != 0 {
		n += 1 + sovDeal(uint64(m.TotalAmount))
	}
	if m.AvailableAmount != 0 {
		n += 1 + sovDeal(uint64(m.AvailableAmount))
	}
	if m.StartBlock != 0 {
		n += 1 + sovDeal(uint64(m.StartBlock))
	}
	if m.EndBlock != 0 {
		n += 1 + sovDeal(uint64(m.EndBlock))
	}
	if len(m.InitialFrontier) > 0 {
		for _, s := range m.InitialFrontier {
			l = len(s)
			n += 1 + l + sovDeal(uint64(l))
		}
	}
	return n
}

func sovDeal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeal(x uint64) (n int) {
	return sovDeal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Deal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeal
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
			return fmt.Errorf("proto: Deal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeal
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
				return ErrInvalidLengthDeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeal
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
					return ErrIntOverflowDeal
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
				return ErrInvalidLengthDeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeal
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
					return ErrIntOverflowDeal
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
				return ErrInvalidLengthDeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeal
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
					return ErrIntOverflowDeal
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
				return ErrInvalidLengthDeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeal
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
					return ErrIntOverflowDeal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Deal_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
			}
			m.TotalAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalAmount |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowDeal
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
					return ErrIntOverflowDeal
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
					return ErrIntOverflowDeal
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
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialFrontier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeal
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
				return ErrInvalidLengthDeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialFrontier = append(m.InitialFrontier, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeal
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
func skipDeal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeal
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
					return 0, ErrIntOverflowDeal
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
					return 0, ErrIntOverflowDeal
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
				return 0, ErrInvalidLengthDeal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeal = fmt.Errorf("proto: unexpected end of group")
)
