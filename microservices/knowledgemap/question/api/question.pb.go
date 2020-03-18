// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/question.proto

package api

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CRqQueryMyQuestionInfoBySubject struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Subject              string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CRqQueryMyQuestionInfoBySubject) Reset()         { *m = CRqQueryMyQuestionInfoBySubject{} }
func (m *CRqQueryMyQuestionInfoBySubject) String() string { return proto.CompactTextString(m) }
func (*CRqQueryMyQuestionInfoBySubject) ProtoMessage()    {}
func (*CRqQueryMyQuestionInfoBySubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_500b4a171d5aefd1, []int{0}
}
func (m *CRqQueryMyQuestionInfoBySubject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CRqQueryMyQuestionInfoBySubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CRqQueryMyQuestionInfoBySubject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CRqQueryMyQuestionInfoBySubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CRqQueryMyQuestionInfoBySubject.Merge(m, src)
}
func (m *CRqQueryMyQuestionInfoBySubject) XXX_Size() int {
	return m.Size()
}
func (m *CRqQueryMyQuestionInfoBySubject) XXX_DiscardUnknown() {
	xxx_messageInfo_CRqQueryMyQuestionInfoBySubject.DiscardUnknown(m)
}

var xxx_messageInfo_CRqQueryMyQuestionInfoBySubject proto.InternalMessageInfo

func (m *CRqQueryMyQuestionInfoBySubject) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CRqQueryMyQuestionInfoBySubject) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type CRpMyQuestionInfoBySubject struct {
	Knowledgenodes       []string `protobuf:"bytes,1,rep,name=knowledgenodes,proto3" json:"knowledgenodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CRpMyQuestionInfoBySubject) Reset()         { *m = CRpMyQuestionInfoBySubject{} }
func (m *CRpMyQuestionInfoBySubject) String() string { return proto.CompactTextString(m) }
func (*CRpMyQuestionInfoBySubject) ProtoMessage()    {}
func (*CRpMyQuestionInfoBySubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_500b4a171d5aefd1, []int{1}
}
func (m *CRpMyQuestionInfoBySubject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CRpMyQuestionInfoBySubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CRpMyQuestionInfoBySubject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CRpMyQuestionInfoBySubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CRpMyQuestionInfoBySubject.Merge(m, src)
}
func (m *CRpMyQuestionInfoBySubject) XXX_Size() int {
	return m.Size()
}
func (m *CRpMyQuestionInfoBySubject) XXX_DiscardUnknown() {
	xxx_messageInfo_CRpMyQuestionInfoBySubject.DiscardUnknown(m)
}

var xxx_messageInfo_CRpMyQuestionInfoBySubject proto.InternalMessageInfo

func (m *CRpMyQuestionInfoBySubject) GetKnowledgenodes() []string {
	if m != nil {
		return m.Knowledgenodes
	}
	return nil
}

func init() {
	proto.RegisterType((*CRqQueryMyQuestionInfoBySubject)(nil), "api.CRqQueryMyQuestionInfoBySubject")
	proto.RegisterType((*CRpMyQuestionInfoBySubject)(nil), "api.CRpMyQuestionInfoBySubject")
}

func init() { proto.RegisterFile("api/question.proto", fileDescriptor_500b4a171d5aefd1) }

var fileDescriptor_500b4a171d5aefd1 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2c, 0xc8, 0xd4,
	0x2f, 0x2c, 0x4d, 0x2d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4e, 0x2c, 0xc8, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x25, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05,
	0xd1, 0xa3, 0x14, 0xc5, 0x25, 0xef, 0x1c, 0x54, 0x18, 0x58, 0x9a, 0x5a, 0x54, 0xe9, 0x5b, 0x19,
	0x08, 0x35, 0xcf, 0x33, 0x2f, 0x2d, 0xdf, 0xa9, 0x32, 0xb8, 0x34, 0x29, 0x2b, 0x35, 0xb9, 0x44,
	0x48, 0x80, 0x8b, 0xb9, 0x34, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14,
	0x52, 0xe5, 0x62, 0x2f, 0x86, 0x48, 0x4a, 0x30, 0x81, 0x44, 0x9d, 0xb8, 0x5f, 0xdd, 0x93, 0x87,
	0x09, 0x05, 0xc1, 0x18, 0x4a, 0x2e, 0x5c, 0x52, 0xce, 0x41, 0x05, 0xb8, 0x8c, 0x55, 0xe3, 0xe2,
	0xcb, 0xce, 0xcb, 0x2f, 0xcf, 0x49, 0x4d, 0x49, 0x4f, 0xcd, 0xcb, 0x4f, 0x49, 0x2d, 0x96, 0x60,
	0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x42, 0x13, 0x35, 0xca, 0xe0, 0xe2, 0x80, 0x19, 0x20, 0x14, 0xc3,
	0x25, 0xe8, 0x9e, 0x5a, 0x82, 0x6a, 0xa2, 0x90, 0x8a, 0x5e, 0x62, 0x41, 0xa6, 0x1e, 0x01, 0x5f,
	0x48, 0xc9, 0x43, 0x55, 0xe1, 0x72, 0x8f, 0x12, 0x83, 0x93, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0x38, 0x90,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x38, 0x9f, 0xd3, 0x6e, 0x01, 0x00, 0x00,
}

func (m *CRqQueryMyQuestionInfoBySubject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CRqQueryMyQuestionInfoBySubject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQuestion(dAtA, i, uint64(len(m.Uid)))
		i += copy(dAtA[i:], m.Uid)
	}
	if len(m.Subject) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintQuestion(dAtA, i, uint64(len(m.Subject)))
		i += copy(dAtA[i:], m.Subject)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CRpMyQuestionInfoBySubject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CRpMyQuestionInfoBySubject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Knowledgenodes) > 0 {
		for _, s := range m.Knowledgenodes {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintQuestion(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CRqQueryMyQuestionInfoBySubject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovQuestion(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovQuestion(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CRpMyQuestionInfoBySubject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Knowledgenodes) > 0 {
		for _, s := range m.Knowledgenodes {
			l = len(s)
			n += 1 + l + sovQuestion(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovQuestion(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozQuestion(x uint64) (n int) {
	return sovQuestion(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CRqQueryMyQuestionInfoBySubject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuestion
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
			return fmt.Errorf("proto: CRqQueryMyQuestionInfoBySubject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CRqQueryMyQuestionInfoBySubject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuestion
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
				return ErrInvalidLengthQuestion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuestion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuestion
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
				return ErrInvalidLengthQuestion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuestion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuestion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuestion
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuestion
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CRpMyQuestionInfoBySubject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuestion
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
			return fmt.Errorf("proto: CRpMyQuestionInfoBySubject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CRpMyQuestionInfoBySubject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Knowledgenodes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuestion
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
				return ErrInvalidLengthQuestion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuestion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Knowledgenodes = append(m.Knowledgenodes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuestion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuestion
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuestion
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuestion(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuestion
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
					return 0, ErrIntOverflowQuestion
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
					return 0, ErrIntOverflowQuestion
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
				return 0, ErrInvalidLengthQuestion
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthQuestion
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQuestion
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
				next, err := skipQuestion(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthQuestion
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
	ErrInvalidLengthQuestion = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuestion   = fmt.Errorf("proto: integer overflow")
)