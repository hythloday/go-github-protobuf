// Code generated by protoc-gen-gogo.
// source: push_event.proto
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

type PushCommit struct {
	Sha       string      `protobuf:"bytes,1,opt,name=sha,proto3" json:"sha,omitempty"`
	TreeId    string      `protobuf:"bytes,2,opt,name=tree_id,json=treeId,proto3" json:"tree_id,omitempty"`
	Distinct  bool        `protobuf:"varint,3,opt,name=distinct,proto3" json:"distinct,omitempty"`
	Message   string      `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp string      `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Url       string      `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Author    *CommitUser `protobuf:"bytes,7,opt,name=author" json:"author,omitempty"`
	Committer *CommitUser `protobuf:"bytes,8,opt,name=committer" json:"committer,omitempty"`
	Added     []string    `protobuf:"bytes,9,rep,name=added" json:"added,omitempty"`
	Removed   []string    `protobuf:"bytes,10,rep,name=removed" json:"removed,omitempty"`
	Modified  []string    `protobuf:"bytes,11,rep,name=modified" json:"modified,omitempty"`
}

func (m *PushCommit) Reset()                    { *m = PushCommit{} }
func (m *PushCommit) String() string            { return proto.CompactTextString(m) }
func (*PushCommit) ProtoMessage()               {}
func (*PushCommit) Descriptor() ([]byte, []int) { return fileDescriptorPushEvent, []int{0} }

type PushEvent struct {
	Ref        string        `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Before     string        `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	After      string        `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	Created    bool          `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Deleted    bool          `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Forced     bool          `protobuf:"varint,6,opt,name=forced,proto3" json:"forced,omitempty"`
	BaseRef    string        `protobuf:"bytes,7,opt,name=base_ref,json=baseRef,proto3" json:"base_ref,omitempty"`
	Compare    string        `protobuf:"bytes,8,opt,name=compare,proto3" json:"compare,omitempty"`
	Commits    []*PushCommit `protobuf:"bytes,9,rep,name=commits" json:"commits,omitempty"`
	HeadCommit *PushCommit   `protobuf:"bytes,10,opt,name=head_commit,json=headCommit" json:"head_commit,omitempty"`
	Repository *Repository   `protobuf:"bytes,11,opt,name=repository" json:"repository,omitempty"`
	Pusher     *CommitUser   `protobuf:"bytes,12,opt,name=pusher" json:"pusher,omitempty"`
	Sender     *User         `protobuf:"bytes,13,opt,name=sender" json:"sender,omitempty"`
}

func (m *PushEvent) Reset()                    { *m = PushEvent{} }
func (m *PushEvent) String() string            { return proto.CompactTextString(m) }
func (*PushEvent) ProtoMessage()               {}
func (*PushEvent) Descriptor() ([]byte, []int) { return fileDescriptorPushEvent, []int{1} }

func init() {
	proto.RegisterType((*PushCommit)(nil), "github.PushCommit")
	proto.RegisterType((*PushEvent)(nil), "github.PushEvent")
}
func (m *PushCommit) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PushCommit) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sha) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.Sha)))
		i += copy(data[i:], m.Sha)
	}
	if len(m.TreeId) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.TreeId)))
		i += copy(data[i:], m.TreeId)
	}
	if m.Distinct {
		data[i] = 0x18
		i++
		if m.Distinct {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if len(m.Message) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.Message)))
		i += copy(data[i:], m.Message)
	}
	if len(m.Timestamp) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.Timestamp)))
		i += copy(data[i:], m.Timestamp)
	}
	if len(m.Url) > 0 {
		data[i] = 0x32
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.Url)))
		i += copy(data[i:], m.Url)
	}
	if m.Author != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintPushEvent(data, i, uint64(m.Author.Size()))
		n1, err := m.Author.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Committer != nil {
		data[i] = 0x42
		i++
		i = encodeVarintPushEvent(data, i, uint64(m.Committer.Size()))
		n2, err := m.Committer.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Added) > 0 {
		for _, s := range m.Added {
			data[i] = 0x4a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.Removed) > 0 {
		for _, s := range m.Removed {
			data[i] = 0x52
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.Modified) > 0 {
		for _, s := range m.Modified {
			data[i] = 0x5a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *PushEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PushEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ref) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.Ref)))
		i += copy(data[i:], m.Ref)
	}
	if len(m.Before) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.Before)))
		i += copy(data[i:], m.Before)
	}
	if len(m.After) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.After)))
		i += copy(data[i:], m.After)
	}
	if m.Created {
		data[i] = 0x20
		i++
		if m.Created {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Deleted {
		data[i] = 0x28
		i++
		if m.Deleted {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Forced {
		data[i] = 0x30
		i++
		if m.Forced {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if len(m.BaseRef) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.BaseRef)))
		i += copy(data[i:], m.BaseRef)
	}
	if len(m.Compare) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintPushEvent(data, i, uint64(len(m.Compare)))
		i += copy(data[i:], m.Compare)
	}
	if len(m.Commits) > 0 {
		for _, msg := range m.Commits {
			data[i] = 0x4a
			i++
			i = encodeVarintPushEvent(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.HeadCommit != nil {
		data[i] = 0x52
		i++
		i = encodeVarintPushEvent(data, i, uint64(m.HeadCommit.Size()))
		n3, err := m.HeadCommit.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Repository != nil {
		data[i] = 0x5a
		i++
		i = encodeVarintPushEvent(data, i, uint64(m.Repository.Size()))
		n4, err := m.Repository.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Pusher != nil {
		data[i] = 0x62
		i++
		i = encodeVarintPushEvent(data, i, uint64(m.Pusher.Size()))
		n5, err := m.Pusher.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Sender != nil {
		data[i] = 0x6a
		i++
		i = encodeVarintPushEvent(data, i, uint64(m.Sender.Size()))
		n6, err := m.Sender.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func encodeFixed64PushEvent(data []byte, offset int, v uint64) int {
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
func encodeFixed32PushEvent(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPushEvent(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *PushCommit) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sha)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.TreeId)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Distinct {
		n += 2
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Author != nil {
		l = m.Author.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Committer != nil {
		l = m.Committer.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if len(m.Added) > 0 {
		for _, s := range m.Added {
			l = len(s)
			n += 1 + l + sovPushEvent(uint64(l))
		}
	}
	if len(m.Removed) > 0 {
		for _, s := range m.Removed {
			l = len(s)
			n += 1 + l + sovPushEvent(uint64(l))
		}
	}
	if len(m.Modified) > 0 {
		for _, s := range m.Modified {
			l = len(s)
			n += 1 + l + sovPushEvent(uint64(l))
		}
	}
	return n
}

func (m *PushEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Ref)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.Before)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.After)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Created {
		n += 2
	}
	if m.Deleted {
		n += 2
	}
	if m.Forced {
		n += 2
	}
	l = len(m.BaseRef)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	l = len(m.Compare)
	if l > 0 {
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if len(m.Commits) > 0 {
		for _, e := range m.Commits {
			l = e.Size()
			n += 1 + l + sovPushEvent(uint64(l))
		}
	}
	if m.HeadCommit != nil {
		l = m.HeadCommit.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Repository != nil {
		l = m.Repository.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Pusher != nil {
		l = m.Pusher.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovPushEvent(uint64(l))
	}
	return n
}

func sovPushEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPushEvent(x uint64) (n int) {
	return sovPushEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PushCommit) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPushEvent
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
			return fmt.Errorf("proto: PushCommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushCommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TreeId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distinct", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Distinct = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Author == nil {
				m.Author = &CommitUser{}
			}
			if err := m.Author.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Committer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Committer == nil {
				m.Committer = &CommitUser{}
			}
			if err := m.Committer.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Added", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Added = append(m.Added, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Removed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Removed = append(m.Removed, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modified", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Modified = append(m.Modified, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPushEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPushEvent
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
func (m *PushEvent) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPushEvent
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
			return fmt.Errorf("proto: PushEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ref = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Before", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Before = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field After", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.After = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Created = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Forced", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Forced = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseRef", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseRef = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Compare = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commits = append(m.Commits, &PushCommit{})
			if err := m.Commits[len(m.Commits)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeadCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HeadCommit == nil {
				m.HeadCommit = &PushCommit{}
			}
			if err := m.HeadCommit.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repository", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
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
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pusher", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pusher == nil {
				m.Pusher = &CommitUser{}
			}
			if err := m.Pusher.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPushEvent
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
				return ErrInvalidLengthPushEvent
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
			skippy, err := skipPushEvent(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPushEvent
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
func skipPushEvent(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPushEvent
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
					return 0, ErrIntOverflowPushEvent
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
					return 0, ErrIntOverflowPushEvent
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
				return 0, ErrInvalidLengthPushEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPushEvent
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
				next, err := skipPushEvent(data[start:])
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
	ErrInvalidLengthPushEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPushEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("push_event.proto", fileDescriptorPushEvent) }

var fileDescriptorPushEvent = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x6e, 0x9a, 0xc6, 0xb1, 0x27, 0x45, 0x0a, 0xab, 0x0a, 0x96, 0x08, 0x45, 0x55, 0xc5, 0x01,
	0x21, 0x48, 0x51, 0xfb, 0x06, 0x20, 0x0e, 0xdc, 0x90, 0x25, 0xce, 0x96, 0x9d, 0x1d, 0xc7, 0x96,
	0xea, 0xac, 0xb5, 0xbb, 0xae, 0xc4, 0x9b, 0xf0, 0x22, 0x3c, 0x03, 0x3d, 0xf2, 0x08, 0xfc, 0xbc,
	0x08, 0x33, 0xbb, 0x76, 0x92, 0x1e, 0x72, 0xb0, 0x32, 0xdf, 0x7c, 0xdf, 0x64, 0x66, 0x3e, 0x8f,
	0x61, 0xde, 0x76, 0xb6, 0xca, 0xf0, 0x1e, 0xb7, 0x6e, 0xd5, 0x1a, 0xed, 0xb4, 0x88, 0x36, 0xb5,
	0xab, 0xba, 0x62, 0xf1, 0x2e, 0xfc, 0xae, 0xd6, 0xba, 0xb9, 0xde, 0xe8, 0x8d, 0xbe, 0xf6, 0x74,
	0xd1, 0x95, 0x1e, 0x79, 0xe0, 0xa3, 0x50, 0xb6, 0x80, 0xce, 0xa2, 0xe9, 0xe3, 0xa7, 0x54, 0xd3,
	0xd4, 0x2e, 0x3b, 0x48, 0xcd, 0x0d, 0xb6, 0xda, 0xd6, 0x4e, 0x9b, 0x6f, 0x21, 0x73, 0xf5, 0xf3,
	0x14, 0xe0, 0x0b, 0x35, 0xff, 0xe8, 0xb5, 0x62, 0x0e, 0x63, 0x5b, 0xe5, 0x72, 0x74, 0x39, 0x7a,
	0x9d, 0xa4, 0x1c, 0x8a, 0xe7, 0x30, 0x75, 0x06, 0x31, 0xab, 0x95, 0x3c, 0xf5, 0xd9, 0x88, 0xe1,
	0x67, 0x25, 0x16, 0x10, 0xab, 0xda, 0xba, 0x7a, 0xbb, 0x76, 0x72, 0x4c, 0x4c, 0x9c, 0xee, 0xb0,
	0x90, 0x30, 0x6d, 0xd0, 0xda, 0x7c, 0x83, 0xf2, 0xcc, 0x17, 0x0d, 0x50, 0xbc, 0x84, 0xc4, 0xd5,
	0x04, 0x5c, 0xde, 0xb4, 0x72, 0xe2, 0xb9, 0x7d, 0x82, 0xdb, 0x77, 0xe6, 0x4e, 0x46, 0xa1, 0x3d,
	0x85, 0xe2, 0x0d, 0x44, 0x79, 0xe7, 0x2a, 0x6d, 0xe4, 0x94, 0x92, 0xb3, 0x1b, 0xb1, 0xea, 0x0d,
	0x09, 0x03, 0x7f, 0xa5, 0xdd, 0xd2, 0x5e, 0x21, 0xde, 0x43, 0x12, 0x56, 0x76, 0x68, 0x64, 0x7c,
	0x54, 0xbe, 0x17, 0x89, 0x0b, 0x98, 0xe4, 0x4a, 0xa1, 0x92, 0xc9, 0xe5, 0x98, 0x3a, 0x06, 0xc0,
	0xd3, 0x1b, 0x6c, 0xf4, 0x3d, 0xe5, 0xc1, 0xe7, 0x07, 0xc8, 0x3b, 0x37, 0x5a, 0xd5, 0x65, 0x4d,
	0xd4, 0xcc, 0x53, 0x3b, 0x7c, 0xf5, 0x63, 0x0c, 0x09, 0x3b, 0xf9, 0x89, 0xdf, 0x22, 0x6f, 0x62,
	0xb0, 0x1c, 0x8c, 0xa4, 0x50, 0x3c, 0x83, 0xa8, 0xc0, 0x52, 0x1b, 0x1c, 0x7c, 0x0c, 0xc8, 0xcf,
	0x50, 0xf2, 0xc4, 0x63, 0x9f, 0x0e, 0x80, 0x67, 0x58, 0x1b, 0xcc, 0x1d, 0x35, 0x3a, 0xf3, 0xe6,
	0x0e, 0x90, 0x19, 0x85, 0x77, 0xc8, 0xcc, 0x24, 0x30, 0x3d, 0xe4, 0x0e, 0xf4, 0x8f, 0x6b, 0x22,
	0x22, 0x4f, 0xf4, 0x48, 0xbc, 0x80, 0xb8, 0xc8, 0x2d, 0x66, 0x3c, 0xd0, 0x34, 0xbc, 0x0e, 0xc6,
	0x29, 0x0d, 0xc5, 0x6d, 0x74, 0xd3, 0xe6, 0x34, 0x55, 0x1c, 0x98, 0x1e, 0x8a, 0xb7, 0x9e, 0x21,
	0x9f, 0xac, 0x37, 0xe7, 0xc0, 0xca, 0xfd, 0xb9, 0xa4, 0x83, 0x44, 0xdc, 0xc2, 0xac, 0xc2, 0x5c,
	0x65, 0x01, 0x93, 0x6d, 0xa3, 0x23, 0x15, 0xc0, 0xb2, 0xfe, 0xd8, 0x6e, 0x00, 0xf6, 0xf7, 0x48,
	0x7e, 0x3e, 0xaa, 0x49, 0x77, 0x4c, 0x7a, 0xa0, 0xe2, 0x7b, 0xe0, 0x6f, 0x85, 0xec, 0x3a, 0x3f,
	0x7e, 0x0f, 0x41, 0x21, 0x5e, 0x41, 0x64, 0x71, 0xab, 0x48, 0xfb, 0xc4, 0x6b, 0xcf, 0x07, 0x6d,
	0x50, 0x05, 0xee, 0xc3, 0xc5, 0xc3, 0x9f, 0xe5, 0xc9, 0xc3, 0xdf, 0xe5, 0xe8, 0x17, 0x3d, 0xbf,
	0xe9, 0xf9, 0xfe, 0x6f, 0x79, 0x52, 0x44, 0xfe, 0xf3, 0xb8, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0x57, 0xb5, 0x88, 0xa3, 0x9a, 0x03, 0x00, 0x00,
}
