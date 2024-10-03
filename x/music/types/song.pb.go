// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: music/music/song.proto

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

type Song struct {
	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Creator  string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Playlist uint64 `protobuf:"varint,3,opt,name=playlist,proto3" json:"playlist,omitempty"`
	Id       uint64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}
func (*Song) Descriptor() ([]byte, []int) {
	return fileDescriptor_663363df4cf3a846, []int{0}
}
func (m *Song) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Song) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Song.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Song) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Song.Merge(m, src)
}
func (m *Song) XXX_Size() int {
	return m.Size()
}
func (m *Song) XXX_DiscardUnknown() {
	xxx_messageInfo_Song.DiscardUnknown(m)
}

var xxx_messageInfo_Song proto.InternalMessageInfo

func (m *Song) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Song) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Song) GetPlaylist() uint64 {
	if m != nil {
		return m.Playlist
	}
	return 0
}

func (m *Song) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Song)(nil), "music.music.Song")
}

func init() { proto.RegisterFile("music/music/song.proto", fileDescriptor_663363df4cf3a846) }

var fileDescriptor_663363df4cf3a846 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x2d, 0x2d, 0xce,
	0x4c, 0xd6, 0x87, 0x90, 0xc5, 0xf9, 0x79, 0xe9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc,
	0x60, 0x11, 0x3d, 0x30, 0xa9, 0x94, 0xc4, 0xc5, 0x12, 0x9c, 0x9f, 0x97, 0x2e, 0x24, 0xc2, 0xc5,
	0x5a, 0x92, 0x59, 0x92, 0x93, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x49,
	0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4, 0x17, 0x49, 0x30, 0x81, 0xc5, 0x61, 0x5c, 0x21,
	0x29, 0x2e, 0x8e, 0x82, 0x9c, 0xc4, 0xca, 0x9c, 0xcc, 0xe2, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0x96, 0x20, 0x38, 0x5f, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x05, 0x2c, 0xca, 0x94, 0x99,
	0xe2, 0xa4, 0x7b, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xc2, 0x10, 0xc7,
	0x55, 0x40, 0x1d, 0x59, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x76, 0xa6, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xe1, 0x8f, 0x4f, 0xb3, 0xc0, 0x00, 0x00, 0x00,
}

func (m *Song) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Song) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Song) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintSong(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x20
	}
	if m.Playlist != 0 {
		i = encodeVarintSong(dAtA, i, uint64(m.Playlist))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSong(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintSong(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSong(dAtA []byte, offset int, v uint64) int {
	offset -= sovSong(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Song) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovSong(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSong(uint64(l))
	}
	if m.Playlist != 0 {
		n += 1 + sovSong(uint64(m.Playlist))
	}
	if m.Id != 0 {
		n += 1 + sovSong(uint64(m.Id))
	}
	return n
}

func sovSong(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSong(x uint64) (n int) {
	return sovSong(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Song) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSong
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
			return fmt.Errorf("proto: Song: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Song: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSong
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
				return ErrInvalidLengthSong
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSong
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
				return ErrInvalidLengthSong
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Playlist", wireType)
			}
			m.Playlist = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSong
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Playlist |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSong
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
			skippy, err := skipSong(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSong
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
func skipSong(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSong
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
					return 0, ErrIntOverflowSong
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
					return 0, ErrIntOverflowSong
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
				return 0, ErrInvalidLengthSong
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSong
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSong
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSong        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSong          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSong = fmt.Errorf("proto: unexpected end of group")
)
