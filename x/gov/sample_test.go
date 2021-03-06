// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/gov/sample_test.proto

package gov

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// ProposalOptions is a sum type of all possible messages that
// may be dispatched via a governance proposal.
//
// For the test case, we only refer to package-internal messages
// and handlers, but an application can reference messages from any package.
type ProposalOptions struct {
	// Types that are valid to be assigned to Option:
	//	*ProposalOptions_Text
	//	*ProposalOptions_Electorate
	//	*ProposalOptions_Rule
	Option isProposalOptions_Option `protobuf_oneof:"option"`
}

func (m *ProposalOptions) Reset()         { *m = ProposalOptions{} }
func (m *ProposalOptions) String() string { return proto.CompactTextString(m) }
func (*ProposalOptions) ProtoMessage()    {}
func (*ProposalOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ad181f69ab09f1, []int{0}
}
func (m *ProposalOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalOptions.Merge(m, src)
}
func (m *ProposalOptions) XXX_Size() int {
	return m.Size()
}
func (m *ProposalOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalOptions proto.InternalMessageInfo

type isProposalOptions_Option interface {
	isProposalOptions_Option()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ProposalOptions_Text struct {
	Text *CreateTextResolutionMsg `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}
type ProposalOptions_Electorate struct {
	Electorate *UpdateElectorateMsg `protobuf:"bytes,2,opt,name=electorate,proto3,oneof"`
}
type ProposalOptions_Rule struct {
	Rule *UpdateElectionRuleMsg `protobuf:"bytes,3,opt,name=rule,proto3,oneof"`
}

func (*ProposalOptions_Text) isProposalOptions_Option()       {}
func (*ProposalOptions_Electorate) isProposalOptions_Option() {}
func (*ProposalOptions_Rule) isProposalOptions_Option()       {}

func (m *ProposalOptions) GetOption() isProposalOptions_Option {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *ProposalOptions) GetText() *CreateTextResolutionMsg {
	if x, ok := m.GetOption().(*ProposalOptions_Text); ok {
		return x.Text
	}
	return nil
}

func (m *ProposalOptions) GetElectorate() *UpdateElectorateMsg {
	if x, ok := m.GetOption().(*ProposalOptions_Electorate); ok {
		return x.Electorate
	}
	return nil
}

func (m *ProposalOptions) GetRule() *UpdateElectionRuleMsg {
	if x, ok := m.GetOption().(*ProposalOptions_Rule); ok {
		return x.Rule
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ProposalOptions) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ProposalOptions_OneofMarshaler, _ProposalOptions_OneofUnmarshaler, _ProposalOptions_OneofSizer, []interface{}{
		(*ProposalOptions_Text)(nil),
		(*ProposalOptions_Electorate)(nil),
		(*ProposalOptions_Rule)(nil),
	}
}

func _ProposalOptions_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ProposalOptions)
	// option
	switch x := m.Option.(type) {
	case *ProposalOptions_Text:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Text); err != nil {
			return err
		}
	case *ProposalOptions_Electorate:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Electorate); err != nil {
			return err
		}
	case *ProposalOptions_Rule:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rule); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ProposalOptions.Option has unexpected type %T", x)
	}
	return nil
}

func _ProposalOptions_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ProposalOptions)
	switch tag {
	case 1: // option.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateTextResolutionMsg)
		err := b.DecodeMessage(msg)
		m.Option = &ProposalOptions_Text{msg}
		return true, err
	case 2: // option.electorate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateElectorateMsg)
		err := b.DecodeMessage(msg)
		m.Option = &ProposalOptions_Electorate{msg}
		return true, err
	case 3: // option.rule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateElectionRuleMsg)
		err := b.DecodeMessage(msg)
		m.Option = &ProposalOptions_Rule{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ProposalOptions_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ProposalOptions)
	// option
	switch x := m.Option.(type) {
	case *ProposalOptions_Text:
		s := proto.Size(x.Text)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProposalOptions_Electorate:
		s := proto.Size(x.Electorate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProposalOptions_Rule:
		s := proto.Size(x.Rule)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ProposalOptions)(nil), "gov.ProposalOptions")
}

func init() { proto.RegisterFile("x/gov/sample_test.proto", fileDescriptor_a3ad181f69ab09f1) }

var fileDescriptor_a3ad181f69ab09f1 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xcf, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc0, 0xf1, 0xc6, 0x3b, 0x0e, 0x89, 0x83, 0x58, 0x04, 0x43, 0x91, 0x20, 0x4e, 0x4e, 0xad,
	0x9c, 0x9b, 0xe3, 0x89, 0xe0, 0x22, 0x4a, 0xd1, 0x59, 0x62, 0xef, 0x11, 0x0e, 0xe2, 0xbd, 0x90,
	0xbc, 0x96, 0x7e, 0x0c, 0xbf, 0x90, 0xbb, 0xe3, 0x8d, 0x8e, 0xd2, 0x7e, 0x11, 0xb9, 0x17, 0x11,
	0xb9, 0x2d, 0xc9, 0xfb, 0xff, 0x20, 0x4f, 0x9e, 0xf4, 0x95, 0xc5, 0xae, 0x8a, 0xe6, 0xcd, 0x3b,
	0x78, 0x21, 0x88, 0x54, 0xfa, 0x80, 0x84, 0xf9, 0xc4, 0x62, 0x57, 0x1c, 0x5b, 0xb4, 0xc8, 0xf7,
	0x6a, 0x7b, 0x4a, 0xa3, 0xe2, 0x28, 0x99, 0x06, 0x97, 0xd0, 0xa4, 0xa7, 0xf3, 0x0f, 0x21, 0x0f,
	0x1f, 0x03, 0x7a, 0x8c, 0xc6, 0x3d, 0x78, 0x5a, 0xe1, 0x3a, 0xe6, 0x73, 0x39, 0x25, 0xe8, 0x49,
	0x89, 0x33, 0x71, 0x71, 0x30, 0x3f, 0x2d, 0x2d, 0x76, 0xe5, 0x4d, 0x00, 0x43, 0xf0, 0x04, 0x3d,
	0xd5, 0x10, 0xd1, 0xb5, 0xdb, 0xf2, 0x3e, 0xda, 0xbb, 0xac, 0xe6, 0x36, 0xbf, 0x96, 0x12, 0x1c,
	0x34, 0x84, 0xc1, 0x10, 0xa8, 0x3d, 0x96, 0x8a, 0xe5, 0xb3, 0x5f, 0x1a, 0x82, 0xdb, 0xbf, 0x61,
	0x52, 0xff, 0xea, 0xfc, 0x52, 0x4e, 0x43, 0xeb, 0x40, 0x4d, 0x58, 0x15, 0xbb, 0x6a, 0x85, 0xeb,
	0xba, 0x75, 0xbf, 0x8e, 0xcb, 0xc5, 0xbe, 0x9c, 0x21, 0x7f, 0x76, 0xa1, 0x3e, 0x07, 0x2d, 0x36,
	0x83, 0x16, 0xdf, 0x83, 0x16, 0xef, 0xa3, 0xce, 0x36, 0xa3, 0xce, 0xbe, 0x46, 0x9d, 0xbd, 0xce,
	0x78, 0xc1, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x40, 0xb5, 0xf6, 0x29, 0x01, 0x00,
	0x00,
}

func (m *ProposalOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Option != nil {
		nn1, err := m.Option.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *ProposalOptions_Text) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Text != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Text.Size()))
		n2, err := m.Text.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *ProposalOptions_Electorate) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Electorate != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Electorate.Size()))
		n3, err := m.Electorate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *ProposalOptions_Rule) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Rule != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Rule.Size()))
		n4, err := m.Rule.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func encodeVarintSampleTest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ProposalOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Option != nil {
		n += m.Option.Size()
	}
	return n
}

func (m *ProposalOptions_Text) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Text != nil {
		l = m.Text.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	return n
}
func (m *ProposalOptions_Electorate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Electorate != nil {
		l = m.Electorate.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	return n
}
func (m *ProposalOptions_Rule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rule != nil {
		l = m.Rule.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	return n
}

func sovSampleTest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSampleTest(x uint64) (n int) {
	return sovSampleTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSampleTest
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
			return fmt.Errorf("proto: ProposalOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CreateTextResolutionMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Option = &ProposalOptions_Text{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Electorate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &UpdateElectorateMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Option = &ProposalOptions_Electorate{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rule", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &UpdateElectionRuleMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Option = &ProposalOptions_Rule{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSampleTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSampleTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSampleTest
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
func skipSampleTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSampleTest
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
					return 0, ErrIntOverflowSampleTest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSampleTest
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
				return 0, ErrInvalidLengthSampleTest
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSampleTest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSampleTest
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSampleTest(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSampleTest
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSampleTest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSampleTest   = fmt.Errorf("proto: integer overflow")
)
