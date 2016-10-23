// Code generated by protoc-gen-gogo.
// source: create_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateEvent struct {
	Ref          string      `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	RefType      string      `protobuf:"bytes,2,opt,name=ref_type,json=refType,proto3" json:"ref_type,omitempty"`
	MasterBranch string      `protobuf:"bytes,3,opt,name=master_branch,json=masterBranch,proto3" json:"master_branch,omitempty"`
	Description  string      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	PusherType   string      `protobuf:"bytes,5,opt,name=pusher_type,json=pusherType,proto3" json:"pusher_type,omitempty"`
	Repository   *Repository `protobuf:"bytes,6,opt,name=repository" json:"repository,omitempty"`
	Sender       *User       `protobuf:"bytes,7,opt,name=sender" json:"sender,omitempty"`
}

func (m *CreateEvent) Reset()                    { *m = CreateEvent{} }
func (m *CreateEvent) String() string            { return proto.CompactTextString(m) }
func (*CreateEvent) ProtoMessage()               {}
func (*CreateEvent) Descriptor() ([]byte, []int) { return fileDescriptorCreateEvent, []int{0} }

func init() {
	proto.RegisterType((*CreateEvent)(nil), "github.CreateEvent")
}
func (m *CreateEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CreateEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ref) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCreateEvent(data, i, uint64(len(m.Ref)))
		i += copy(data[i:], m.Ref)
	}
	if len(m.RefType) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintCreateEvent(data, i, uint64(len(m.RefType)))
		i += copy(data[i:], m.RefType)
	}
	if len(m.MasterBranch) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintCreateEvent(data, i, uint64(len(m.MasterBranch)))
		i += copy(data[i:], m.MasterBranch)
	}
	if len(m.Description) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintCreateEvent(data, i, uint64(len(m.Description)))
		i += copy(data[i:], m.Description)
	}
	if len(m.PusherType) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintCreateEvent(data, i, uint64(len(m.PusherType)))
		i += copy(data[i:], m.PusherType)
	}
	if m.Repository != nil {
		data[i] = 0x32
		i++
		i = encodeVarintCreateEvent(data, i, uint64(m.Repository.Size()))
		n1, err := m.Repository.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Sender != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintCreateEvent(data, i, uint64(m.Sender.Size()))
		n2, err := m.Sender.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64CreateEvent(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32CreateEvent(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCreateEvent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *CreateEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Ref)
	if l > 0 {
		n += 1 + l + sovCreateEvent(uint64(l))
	}
	l = len(m.RefType)
	if l > 0 {
		n += 1 + l + sovCreateEvent(uint64(l))
	}
	l = len(m.MasterBranch)
	if l > 0 {
		n += 1 + l + sovCreateEvent(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCreateEvent(uint64(l))
	}
	l = len(m.PusherType)
	if l > 0 {
		n += 1 + l + sovCreateEvent(uint64(l))
	}
	if m.Repository != nil {
		l = m.Repository.Size()
		n += 1 + l + sovCreateEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovCreateEvent(uint64(l))
	}
	return n
}

func sovCreateEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCreateEvent(x uint64) (n int) {
	return sovCreateEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCreateEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreateEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCreateEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ref = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreateEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCreateEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefType = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterBranch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreateEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCreateEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MasterBranch = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreateEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCreateEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PusherType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreateEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCreateEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PusherType = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repository", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreateEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCreateEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repository == nil {
				m.Repository = &Repository{}
			}
			if err := m.Repository.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreateEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCreateEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &User{}
			}
			if err := m.Sender.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCreateEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCreateEvent
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
func skipCreateEvent(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCreateEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowCreateEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowCreateEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCreateEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCreateEvent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipCreateEvent(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthCreateEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCreateEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("create_event.proto", fileDescriptorCreateEvent) }

var fileDescriptorCreateEvent = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x50, 0x3d, 0x4e, 0xf3, 0x40,
	0x10, 0x8d, 0xbf, 0x7c, 0x38, 0x30, 0x0e, 0x52, 0xb4, 0xa2, 0x58, 0x52, 0x98, 0x08, 0x28, 0x68,
	0x70, 0xa4, 0x70, 0x83, 0x20, 0x2e, 0x60, 0x41, 0x6d, 0xd9, 0xce, 0xc4, 0x76, 0x11, 0xef, 0x6a,
	0x76, 0x8d, 0x94, 0x9b, 0x70, 0xa4, 0x94, 0x1c, 0x81, 0x9f, 0x6b, 0x50, 0xb0, 0x99, 0x35, 0x90,
	0x62, 0x34, 0xf3, 0x7e, 0xec, 0xb7, 0x7a, 0x20, 0x4a, 0xc2, 0xdc, 0x62, 0x86, 0xcf, 0xd8, 0xda,
	0x44, 0x93, 0xb2, 0x4a, 0x84, 0x55, 0x63, 0xeb, 0xae, 0x98, 0xde, 0xfa, 0x9d, 0x94, 0x6a, 0x33,
	0xaf, 0x54, 0xa5, 0xe6, 0x2c, 0x17, 0xdd, 0x9a, 0x11, 0x03, 0xbe, 0xfc, 0x67, 0x53, 0xe8, 0x0c,
	0x52, 0x7f, 0x4f, 0x08, 0xb5, 0x32, 0x8d, 0x55, 0xb4, 0xf5, 0xcc, 0xe5, 0x57, 0x00, 0xd1, 0x3d,
	0x67, 0x3d, 0xec, 0xa3, 0xc4, 0x04, 0x86, 0x84, 0x6b, 0x19, 0xcc, 0x82, 0x9b, 0x93, 0x74, 0x7f,
	0x8a, 0x73, 0x38, 0x76, 0x2b, 0xb3, 0x5b, 0x8d, 0xf2, 0x1f, 0xd3, 0x23, 0x87, 0x1f, 0x1d, 0x14,
	0x57, 0x70, 0xba, 0xc9, 0x8d, 0x45, 0xca, 0x0a, 0xca, 0xdb, 0xb2, 0x96, 0x43, 0xd6, 0xc7, 0x9e,
	0x5c, 0x32, 0x27, 0x66, 0x10, 0xad, 0xd0, 0x94, 0xd4, 0x68, 0xdb, 0xa8, 0x56, 0xfe, 0x67, 0xcb,
	0x21, 0x25, 0x2e, 0x20, 0xd2, 0x9d, 0xa9, 0xdd, 0x6f, 0x38, 0xe4, 0x88, 0x1d, 0xe0, 0x29, 0xce,
	0x59, 0x00, 0xfc, 0x3d, 0x5c, 0x86, 0x4e, 0x8f, 0x16, 0x22, 0xe9, 0x6b, 0x48, 0x7f, 0x95, 0xf4,
	0xc0, 0x25, 0xae, 0x21, 0x34, 0xd8, 0xae, 0x90, 0xe4, 0x88, 0xfd, 0xe3, 0x1f, 0xff, 0x93, 0xab,
	0x23, 0xed, 0xb5, 0xe5, 0xd9, 0xee, 0x3d, 0x1e, 0xec, 0x3e, 0xe2, 0xe0, 0xd5, 0xcd, 0x9b, 0x9b,
	0x97, 0xcf, 0x78, 0x50, 0x84, 0xdc, 0xcd, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x91,
	0xc0, 0xce, 0x86, 0x01, 0x00, 0x00,
}
