// Code generated by protoc-gen-gogo.
// source: page_build_event.proto
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

type BuildError struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *BuildError) Reset()                    { *m = BuildError{} }
func (m *BuildError) String() string            { return proto.CompactTextString(m) }
func (*BuildError) ProtoMessage()               {}
func (*BuildError) Descriptor() ([]byte, []int) { return fileDescriptorPageBuildEvent, []int{0} }

type Build struct {
	Url       string      `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Status    string      `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error     *BuildError `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	Pusher    *User       `protobuf:"bytes,4,opt,name=pusher" json:"pusher,omitempty"`
	Commit    string      `protobuf:"bytes,5,opt,name=commit,proto3" json:"commit,omitempty"`
	Duration  int32       `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	CreatedAt string      `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string      `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptorPageBuildEvent, []int{1} }

type PageBuildEvent struct {
	Id         int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Build      *Build      `protobuf:"bytes,2,opt,name=build" json:"build,omitempty"`
	Repository *Repository `protobuf:"bytes,3,opt,name=repository" json:"repository,omitempty"`
	Sender     *User       `protobuf:"bytes,4,opt,name=sender" json:"sender,omitempty"`
}

func (m *PageBuildEvent) Reset()                    { *m = PageBuildEvent{} }
func (m *PageBuildEvent) String() string            { return proto.CompactTextString(m) }
func (*PageBuildEvent) ProtoMessage()               {}
func (*PageBuildEvent) Descriptor() ([]byte, []int) { return fileDescriptorPageBuildEvent, []int{2} }

func init() {
	proto.RegisterType((*BuildError)(nil), "github.BuildError")
	proto.RegisterType((*Build)(nil), "github.Build")
	proto.RegisterType((*PageBuildEvent)(nil), "github.PageBuildEvent")
}
func (m *BuildError) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *BuildError) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.Message)))
		i += copy(data[i:], m.Message)
	}
	return i, nil
}

func (m *Build) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Build) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.Url)))
		i += copy(data[i:], m.Url)
	}
	if len(m.Status) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.Status)))
		i += copy(data[i:], m.Status)
	}
	if m.Error != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Error.Size()))
		n1, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Pusher != nil {
		data[i] = 0x22
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Pusher.Size()))
		n2, err := m.Pusher.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Commit) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.Commit)))
		i += copy(data[i:], m.Commit)
	}
	if m.Duration != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Duration))
	}
	if len(m.CreatedAt) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.CreatedAt)))
		i += copy(data[i:], m.CreatedAt)
	}
	if len(m.UpdatedAt) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(len(m.UpdatedAt)))
		i += copy(data[i:], m.UpdatedAt)
	}
	return i, nil
}

func (m *PageBuildEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PageBuildEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Id))
	}
	if m.Build != nil {
		data[i] = 0x12
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Build.Size()))
		n3, err := m.Build.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Repository != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Repository.Size()))
		n4, err := m.Repository.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Sender != nil {
		data[i] = 0x22
		i++
		i = encodeVarintPageBuildEvent(data, i, uint64(m.Sender.Size()))
		n5, err := m.Sender.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func encodeFixed64PageBuildEvent(data []byte, offset int, v uint64) int {
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
func encodeFixed32PageBuildEvent(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPageBuildEvent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *BuildError) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	return n
}

func (m *Build) Size() (n int) {
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Pusher != nil {
		l = m.Pusher.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	l = len(m.Commit)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovPageBuildEvent(uint64(m.Duration))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	return n
}

func (m *PageBuildEvent) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPageBuildEvent(uint64(m.Id))
	}
	if m.Build != nil {
		l = m.Build.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Repository != nil {
		l = m.Repository.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovPageBuildEvent(uint64(l))
	}
	return n
}

func sovPageBuildEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPageBuildEvent(x uint64) (n int) {
	return sovPageBuildEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BuildError) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPageBuildEvent
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
			return fmt.Errorf("proto: BuildError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BuildError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPageBuildEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPageBuildEvent
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
func (m *Build) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPageBuildEvent
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
			return fmt.Errorf("proto: Build: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Build: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &BuildError{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pusher", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pusher == nil {
				m.Pusher = &User{}
			}
			if err := m.Pusher.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commit = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Duration |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPageBuildEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPageBuildEvent
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
func (m *PageBuildEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPageBuildEvent
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
			return fmt.Errorf("proto: PageBuildEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PageBuildEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Build", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Build == nil {
				m.Build = &Build{}
			}
			if err := m.Build.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repository", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPageBuildEvent
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
				return ErrInvalidLengthPageBuildEvent
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
			skippy, err := skipPageBuildEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPageBuildEvent
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
func skipPageBuildEvent(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPageBuildEvent
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
					return 0, ErrIntOverflowPageBuildEvent
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
					return 0, ErrIntOverflowPageBuildEvent
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
				return 0, ErrInvalidLengthPageBuildEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPageBuildEvent
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
				next, err := skipPageBuildEvent(data[start:])
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
	ErrInvalidLengthPageBuildEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPageBuildEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("page_build_event.proto", fileDescriptorPageBuildEvent) }

var fileDescriptorPageBuildEvent = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xdd, 0x4a, 0x02, 0x41,
	0x14, 0xc7, 0x5d, 0x6d, 0x57, 0x3d, 0x96, 0xc8, 0x10, 0x32, 0x08, 0x49, 0x58, 0x84, 0x37, 0x29,
	0xd8, 0x13, 0x28, 0x74, 0x1f, 0x03, 0x5d, 0xcb, 0xae, 0x3b, 0xad, 0x03, 0xea, 0x2c, 0xf3, 0x11,
	0xf4, 0x26, 0xdd, 0xf5, 0x3a, 0x5e, 0xf6, 0x08, 0x7d, 0x3c, 0x42, 0x2f, 0xd0, 0x7c, 0xa9, 0x15,
	0x74, 0x31, 0xec, 0xf9, 0x9f, 0xdf, 0x99, 0x3d, 0xf3, 0x3f, 0x07, 0xba, 0x65, 0x5a, 0xd0, 0x79,
	0xa6, 0xd9, 0x2a, 0x9f, 0xd3, 0x47, 0xba, 0x51, 0xa3, 0x52, 0x70, 0xc5, 0x51, 0x52, 0x30, 0xb5,
	0xd4, 0x59, 0xef, 0xda, 0x7f, 0x47, 0x0b, 0xbe, 0x1e, 0x17, 0xbc, 0xe0, 0x63, 0x87, 0x33, 0xfd,
	0xe0, 0x94, 0x13, 0x2e, 0xf2, 0xd7, 0x7a, 0xa0, 0x25, 0x15, 0x21, 0xee, 0x08, 0x5a, 0x72, 0xc9,
	0x14, 0x17, 0x4f, 0x3e, 0x33, 0xb8, 0x02, 0x98, 0xd9, 0x4e, 0xb7, 0x42, 0x70, 0x81, 0x30, 0xd4,
	0xd7, 0x54, 0x4a, 0xd3, 0x1f, 0x47, 0xe7, 0xd1, 0xb0, 0x49, 0x76, 0x72, 0xf0, 0x15, 0x41, 0xec,
	0x0a, 0x51, 0x07, 0x6a, 0x5a, 0xac, 0x02, 0xb7, 0x21, 0xea, 0x42, 0x22, 0x55, 0xaa, 0xb4, 0xc4,
	0x55, 0x97, 0x0c, 0x0a, 0x0d, 0x21, 0xa6, 0xf6, 0xb7, 0xb8, 0x66, 0xd2, 0xad, 0x09, 0x1a, 0x85,
	0x87, 0x1f, 0x1a, 0x12, 0x5f, 0x80, 0x2e, 0x21, 0x29, 0xb5, 0x5c, 0x52, 0x81, 0x8f, 0x5c, 0xe9,
	0xf1, 0xae, 0xf4, 0xde, 0xbc, 0x9d, 0x04, 0x66, 0xfb, 0x18, 0xcf, 0x6b, 0xa6, 0x70, 0xec, 0xfb,
	0x78, 0x85, 0x7a, 0xd0, 0xc8, 0xb5, 0x48, 0x15, 0xe3, 0x1b, 0x9c, 0x18, 0x12, 0x93, 0xbd, 0x46,
	0x67, 0x00, 0x0b, 0x41, 0x53, 0x45, 0xf3, 0x79, 0xaa, 0x70, 0xdd, 0xdd, 0x6b, 0x86, 0xcc, 0x54,
	0x59, 0xac, 0xcb, 0x7c, 0x87, 0x1b, 0x1e, 0x87, 0xcc, 0x54, 0x0d, 0x5e, 0x22, 0x68, 0xdf, 0x19,
	0xfb, 0xfe, 0xc5, 0x76, 0x17, 0xa8, 0x0d, 0x55, 0x96, 0x3b, 0xf7, 0x31, 0x31, 0x11, 0xba, 0x80,
	0xd8, 0xad, 0xca, 0x79, 0x6f, 0x4d, 0x4e, 0x7e, 0x99, 0x24, 0x9e, 0xa1, 0x09, 0xc0, 0x61, 0xf2,
	0x7f, 0xc7, 0x41, 0xf6, 0x84, 0xfc, 0xa8, 0xb2, 0x33, 0x91, 0x74, 0x93, 0xff, 0x37, 0x13, 0xcf,
	0x66, 0xa7, 0xdb, 0xf7, 0x7e, 0x65, 0xfb, 0xd1, 0x8f, 0x5e, 0xcd, 0x79, 0x33, 0xe7, 0xf9, 0xb3,
	0x5f, 0xc9, 0x12, 0xb7, 0xdc, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xe0, 0x12, 0x04,
	0x4b, 0x02, 0x00, 0x00,
}
