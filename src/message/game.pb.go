// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// point3F
type Point3F struct {
	X                *float32 `protobuf:"fixed32,1,req,name=X" json:"X,omitempty"`
	Y                *float32 `protobuf:"fixed32,2,req,name=Y" json:"Y,omitempty"`
	Z                *float32 `protobuf:"fixed32,3,req,name=Z" json:"Z,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Point3F) Reset()                    { *m = Point3F{} }
func (m *Point3F) String() string            { return proto.CompactTextString(m) }
func (*Point3F) ProtoMessage()               {}
func (*Point3F) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Point3F) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Point3F) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *Point3F) GetZ() float32 {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return 0
}

// 移动包
type C_W_Move struct {
	PacketHead       *Ipacket       `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	Move             *C_W_Move_Move `protobuf:"bytes,2,req,name=move" json:"move,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *C_W_Move) Reset()                    { *m = C_W_Move{} }
func (m *C_W_Move) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move) ProtoMessage()               {}
func (*C_W_Move) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *C_W_Move) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_W_Move) GetMove() *C_W_Move_Move {
	if m != nil {
		return m.Move
	}
	return nil
}

type C_W_Move_Move struct {
	Mode             *int32                `protobuf:"varint,1,req,name=Mode" json:"Mode,omitempty"`
	Normal           *C_W_Move_Move_Normal `protobuf:"bytes,2,opt,name=normal" json:"normal,omitempty"`
	Path             *C_W_Move_Move_Path   `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	Link             *C_W_Move_Move_Blink  `protobuf:"bytes,4,opt,name=link" json:"link,omitempty"`
	Jump             *C_W_Move_Move_Jump   `protobuf:"bytes,5,opt,name=jump" json:"jump,omitempty"`
	Line             *C_W_Move_Move_Line   `protobuf:"bytes,6,opt,name=line" json:"line,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *C_W_Move_Move) Reset()                    { *m = C_W_Move_Move{} }
func (m *C_W_Move_Move) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move) ProtoMessage()               {}
func (*C_W_Move_Move) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

func (m *C_W_Move_Move) GetMode() int32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

func (m *C_W_Move_Move) GetNormal() *C_W_Move_Move_Normal {
	if m != nil {
		return m.Normal
	}
	return nil
}

func (m *C_W_Move_Move) GetPath() *C_W_Move_Move_Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *C_W_Move_Move) GetLink() *C_W_Move_Move_Blink {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *C_W_Move_Move) GetJump() *C_W_Move_Move_Jump {
	if m != nil {
		return m.Jump
	}
	return nil
}

func (m *C_W_Move_Move) GetLine() *C_W_Move_Move_Line {
	if m != nil {
		return m.Line
	}
	return nil
}

type C_W_Move_Move_Normal struct {
	Pos              *Point3F `protobuf:"bytes,1,req,name=Pos" json:"Pos,omitempty"`
	Yaw              *float32 `protobuf:"fixed32,2,req,name=Yaw" json:"Yaw,omitempty"`
	Duration         *float32 `protobuf:"fixed32,3,req,name=Duration" json:"Duration,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *C_W_Move_Move_Normal) Reset()                    { *m = C_W_Move_Move_Normal{} }
func (m *C_W_Move_Move_Normal) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Normal) ProtoMessage()               {}
func (*C_W_Move_Move_Normal) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 0} }

func (m *C_W_Move_Move_Normal) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *C_W_Move_Move_Normal) GetYaw() float32 {
	if m != nil && m.Yaw != nil {
		return *m.Yaw
	}
	return 0
}

func (m *C_W_Move_Move_Normal) GetDuration() float32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

type C_W_Move_Move_Path struct {
	PathId           *int32 `protobuf:"varint,1,req,name=PathId" json:"PathId,omitempty"`
	TimePos          *int32 `protobuf:"varint,2,req,name=TimePos" json:"TimePos,omitempty"`
	MountId          *int32 `protobuf:"varint,3,req,name=MountId" json:"MountId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *C_W_Move_Move_Path) Reset()                    { *m = C_W_Move_Move_Path{} }
func (m *C_W_Move_Move_Path) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Path) ProtoMessage()               {}
func (*C_W_Move_Move_Path) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 1} }

func (m *C_W_Move_Move_Path) GetPathId() int32 {
	if m != nil && m.PathId != nil {
		return *m.PathId
	}
	return 0
}

func (m *C_W_Move_Move_Path) GetTimePos() int32 {
	if m != nil && m.TimePos != nil {
		return *m.TimePos
	}
	return 0
}

func (m *C_W_Move_Move_Path) GetMountId() int32 {
	if m != nil && m.MountId != nil {
		return *m.MountId
	}
	return 0
}

type C_W_Move_Move_Blink struct {
	Pos              *Point3F `protobuf:"bytes,1,req,name=Pos" json:"Pos,omitempty"`
	RPos             *Point3F `protobuf:"bytes,2,req,name=RPos" json:"RPos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *C_W_Move_Move_Blink) Reset()                    { *m = C_W_Move_Move_Blink{} }
func (m *C_W_Move_Move_Blink) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Blink) ProtoMessage()               {}
func (*C_W_Move_Move_Blink) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 2} }

func (m *C_W_Move_Move_Blink) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *C_W_Move_Move_Blink) GetRPos() *Point3F {
	if m != nil {
		return m.RPos
	}
	return nil
}

type C_W_Move_Move_Jump struct {
	BPos             *Point3F `protobuf:"bytes,1,req,name=BPos" json:"BPos,omitempty"`
	EPos             *Point3F `protobuf:"bytes,2,req,name=EPos" json:"EPos,omitempty"`
	Duration         *int32   `protobuf:"varint,3,req,name=Duration" json:"Duration,omitempty"`
	TimePos          *int32   `protobuf:"varint,4,req,name=TimePos" json:"TimePos,omitempty"`
	UpExDur          *int32   `protobuf:"varint,5,req,name=UpExDur" json:"UpExDur,omitempty"`
	DownExDur        *int32   `protobuf:"varint,6,req,name=DownExDur" json:"DownExDur,omitempty"`
	A                *int32   `protobuf:"varint,7,req,name=A" json:"A,omitempty"`
	B                *int32   `protobuf:"varint,8,req,name=B" json:"B,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *C_W_Move_Move_Jump) Reset()                    { *m = C_W_Move_Move_Jump{} }
func (m *C_W_Move_Move_Jump) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Jump) ProtoMessage()               {}
func (*C_W_Move_Move_Jump) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 3} }

func (m *C_W_Move_Move_Jump) GetBPos() *Point3F {
	if m != nil {
		return m.BPos
	}
	return nil
}

func (m *C_W_Move_Move_Jump) GetEPos() *Point3F {
	if m != nil {
		return m.EPos
	}
	return nil
}

func (m *C_W_Move_Move_Jump) GetDuration() int32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetTimePos() int32 {
	if m != nil && m.TimePos != nil {
		return *m.TimePos
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetUpExDur() int32 {
	if m != nil && m.UpExDur != nil {
		return *m.UpExDur
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetDownExDur() int32 {
	if m != nil && m.DownExDur != nil {
		return *m.DownExDur
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetA() int32 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *C_W_Move_Move_Jump) GetB() int32 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

type C_W_Move_Move_Line struct {
	BPos             *Point3F `protobuf:"bytes,1,req,name=BPos" json:"BPos,omitempty"`
	EPos             *Point3F `protobuf:"bytes,2,req,name=EPos" json:"EPos,omitempty"`
	Duration         *int32   `protobuf:"varint,3,req,name=Duration" json:"Duration,omitempty"`
	TimePos          *int32   `protobuf:"varint,4,req,name=TimePos" json:"TimePos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *C_W_Move_Move_Line) Reset()                    { *m = C_W_Move_Move_Line{} }
func (m *C_W_Move_Move_Line) String() string            { return proto.CompactTextString(m) }
func (*C_W_Move_Move_Line) ProtoMessage()               {}
func (*C_W_Move_Move_Line) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0, 4} }

func (m *C_W_Move_Move_Line) GetBPos() *Point3F {
	if m != nil {
		return m.BPos
	}
	return nil
}

func (m *C_W_Move_Move_Line) GetEPos() *Point3F {
	if m != nil {
		return m.EPos
	}
	return nil
}

func (m *C_W_Move_Move_Line) GetDuration() int32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *C_W_Move_Move_Line) GetTimePos() int32 {
	if m != nil && m.TimePos != nil {
		return *m.TimePos
	}
	return 0
}

type W_C_LoginMap struct {
	PacketHead       *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	Id               *int64   `protobuf:"varint,2,req,name=Id" json:"Id,omitempty"`
	Pos              *Point3F `protobuf:"bytes,3,req,name=Pos" json:"Pos,omitempty"`
	Rotation         *float32 `protobuf:"fixed32,4,req,name=Rotation" json:"Rotation,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *W_C_LoginMap) Reset()                    { *m = W_C_LoginMap{} }
func (m *W_C_LoginMap) String() string            { return proto.CompactTextString(m) }
func (*W_C_LoginMap) ProtoMessage()               {}
func (*W_C_LoginMap) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *W_C_LoginMap) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *W_C_LoginMap) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *W_C_LoginMap) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *W_C_LoginMap) GetRotation() float32 {
	if m != nil && m.Rotation != nil {
		return *m.Rotation
	}
	return 0
}

type W_C_Move struct {
	PacketHead       *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	Id               *int64   `protobuf:"varint,2,req,name=Id" json:"Id,omitempty"`
	Pos              *Point3F `protobuf:"bytes,3,req,name=Pos" json:"Pos,omitempty"`
	Rotation         *float32 `protobuf:"fixed32,4,req,name=Rotation" json:"Rotation,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *W_C_Move) Reset()                    { *m = W_C_Move{} }
func (m *W_C_Move) String() string            { return proto.CompactTextString(m) }
func (*W_C_Move) ProtoMessage()               {}
func (*W_C_Move) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *W_C_Move) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *W_C_Move) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *W_C_Move) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *W_C_Move) GetRotation() float32 {
	if m != nil && m.Rotation != nil {
		return *m.Rotation
	}
	return 0
}

type W_C_ADD_SIMOBJ struct {
	PacketHead       *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	Id               *int64   `protobuf:"varint,2,req,name=Id" json:"Id,omitempty"`
	Pos              *Point3F `protobuf:"bytes,3,req,name=Pos" json:"Pos,omitempty"`
	Rotation         *float32 `protobuf:"fixed32,4,req,name=Rotation" json:"Rotation,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *W_C_ADD_SIMOBJ) Reset()                    { *m = W_C_ADD_SIMOBJ{} }
func (m *W_C_ADD_SIMOBJ) String() string            { return proto.CompactTextString(m) }
func (*W_C_ADD_SIMOBJ) ProtoMessage()               {}
func (*W_C_ADD_SIMOBJ) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *W_C_ADD_SIMOBJ) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *W_C_ADD_SIMOBJ) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *W_C_ADD_SIMOBJ) GetPos() *Point3F {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *W_C_ADD_SIMOBJ) GetRotation() float32 {
	if m != nil && m.Rotation != nil {
		return *m.Rotation
	}
	return 0
}

type C_W_LoginCopyMap struct {
	PacketHead       *Ipacket `protobuf:"bytes,1,req,name=PacketHead" json:"PacketHead,omitempty"`
	DataId           *int32   `protobuf:"varint,2,req,name=DataId" json:"DataId,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *C_W_LoginCopyMap) Reset()                    { *m = C_W_LoginCopyMap{} }
func (m *C_W_LoginCopyMap) String() string            { return proto.CompactTextString(m) }
func (*C_W_LoginCopyMap) ProtoMessage()               {}
func (*C_W_LoginCopyMap) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *C_W_LoginCopyMap) GetPacketHead() *Ipacket {
	if m != nil {
		return m.PacketHead
	}
	return nil
}

func (m *C_W_LoginCopyMap) GetDataId() int32 {
	if m != nil && m.DataId != nil {
		return *m.DataId
	}
	return 0
}

func init() {
	proto.RegisterType((*Point3F)(nil), "message.Point3F")
	proto.RegisterType((*C_W_Move)(nil), "message.C_W_Move")
	proto.RegisterType((*C_W_Move_Move)(nil), "message.C_W_Move.Move")
	proto.RegisterType((*C_W_Move_Move_Normal)(nil), "message.C_W_Move.Move.Normal")
	proto.RegisterType((*C_W_Move_Move_Path)(nil), "message.C_W_Move.Move.Path")
	proto.RegisterType((*C_W_Move_Move_Blink)(nil), "message.C_W_Move.Move.Blink")
	proto.RegisterType((*C_W_Move_Move_Jump)(nil), "message.C_W_Move.Move.Jump")
	proto.RegisterType((*C_W_Move_Move_Line)(nil), "message.C_W_Move.Move.Line")
	proto.RegisterType((*W_C_LoginMap)(nil), "message.W_C_LoginMap")
	proto.RegisterType((*W_C_Move)(nil), "message.W_C_Move")
	proto.RegisterType((*W_C_ADD_SIMOBJ)(nil), "message.W_C_ADD_SIMOBJ")
	proto.RegisterType((*C_W_LoginCopyMap)(nil), "message.C_W_LoginCopyMap")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4d, 0x6b, 0x13, 0x41,
	0x18, 0x66, 0xbf, 0xe3, 0x9b, 0x5a, 0xc2, 0x1c, 0xc2, 0xb0, 0x56, 0x28, 0xc1, 0x43, 0xf1, 0x10,
	0x4b, 0x8a, 0x3f, 0x20, 0xc9, 0x56, 0xdc, 0xd2, 0x68, 0x1c, 0x95, 0x36, 0x41, 0x08, 0x83, 0x3b,
	0xa4, 0x6b, 0xb3, 0x3b, 0x4b, 0xb2, 0x69, 0xf5, 0xec, 0xc5, 0x83, 0x07, 0xf1, 0xf7, 0xf9, 0x1b,
	0xfc, 0x0d, 0xf2, 0xbe, 0xbb, 0xdb, 0x16, 0xd2, 0x06, 0x09, 0x82, 0xbd, 0xec, 0xce, 0x33, 0xcf,
	0xf3, 0x7e, 0xcf, 0x0c, 0xc0, 0x54, 0x26, 0xaa, 0x9d, 0xcd, 0x75, 0xae, 0x99, 0x97, 0xa8, 0xc5,
	0x42, 0x4e, 0x95, 0xff, 0xb0, 0x5c, 0x14, 0xfb, 0xad, 0x03, 0xf0, 0x86, 0x3a, 0x4e, 0xf3, 0x83,
	0x17, 0x6c, 0x0b, 0x8c, 0x53, 0x6e, 0xec, 0x9a, 0x7b, 0xa6, 0x30, 0x4e, 0x11, 0x8d, 0xb8, 0x59,
	0xa0, 0x11, 0xa2, 0x31, 0xb7, 0x0a, 0x34, 0x6e, 0xfd, 0xf6, 0xa0, 0xd6, 0x9f, 0x9c, 0x4c, 0x06,
	0xfa, 0x42, 0xb1, 0x7d, 0x80, 0xa1, 0xfc, 0x78, 0xae, 0xf2, 0x97, 0x4a, 0x46, 0x64, 0x5f, 0xef,
	0x34, 0xda, 0x55, 0x94, 0x30, 0x23, 0x4e, 0xdc, 0xd0, 0xb0, 0xa7, 0x60, 0x27, 0xfa, 0x42, 0x91,
	0xf7, 0x7a, 0xa7, 0x79, 0xa5, 0xad, 0x5c, 0xb6, 0xf1, 0x23, 0x48, 0xe3, 0x7f, 0xf5, 0xc0, 0xa6,
	0x30, 0x0c, 0xff, 0x91, 0xa2, 0x00, 0x8e, 0xa0, 0x35, 0x7b, 0x0e, 0x6e, 0xaa, 0xe7, 0x89, 0x9c,
	0x71, 0x73, 0xd7, 0xd8, 0xab, 0x77, 0x1e, 0xdf, 0xee, 0xaa, 0xfd, 0x8a, 0x44, 0xa2, 0x14, 0xb3,
	0x67, 0x60, 0x67, 0x32, 0x3f, 0xe3, 0x16, 0x19, 0x3d, 0xba, 0xc3, 0x68, 0x28, 0xf3, 0x33, 0x41,
	0x42, 0xb6, 0x0f, 0xf6, 0x2c, 0x4e, 0xcf, 0xb9, 0x4d, 0x06, 0x3b, 0x77, 0x18, 0xf4, 0x50, 0x23,
	0x48, 0x89, 0x21, 0x3e, 0x2d, 0x93, 0x8c, 0x3b, 0x6b, 0x43, 0x1c, 0x2d, 0x93, 0x4c, 0x90, 0x10,
	0x0d, 0x66, 0x71, 0xaa, 0xb8, 0xbb, 0xd6, 0xe0, 0x38, 0x4e, 0x15, 0x45, 0x50, 0xfe, 0x18, 0xdc,
	0xa2, 0x2c, 0xd6, 0x02, 0x6b, 0xa8, 0x17, 0x2b, 0x9d, 0x2f, 0xc7, 0x2a, 0x90, 0x64, 0x0d, 0xb0,
	0x46, 0xf2, 0xb2, 0x9c, 0x27, 0x2e, 0x99, 0x0f, 0xb5, 0x60, 0x39, 0x97, 0x79, 0xac, 0xd3, 0x72,
	0xb0, 0x57, 0xd8, 0x17, 0x60, 0x63, 0xf5, 0xac, 0x09, 0x2e, 0xfe, 0xc3, 0xa8, 0xec, 0x7a, 0x89,
	0x18, 0x07, 0xef, 0x5d, 0x9c, 0x28, 0x8c, 0x6a, 0x12, 0x51, 0x41, 0x64, 0x06, 0x7a, 0x99, 0xe6,
	0x61, 0x44, 0x4e, 0x1d, 0x51, 0x41, 0xff, 0x0d, 0x38, 0xd4, 0xa0, 0xbf, 0x4a, 0xf7, 0x09, 0xd8,
	0xa2, 0xf2, 0x7e, 0x9b, 0x88, 0x58, 0xff, 0x97, 0x01, 0x36, 0xb6, 0x10, 0xe5, 0xbd, 0x75, 0x3e,
	0x89, 0x45, 0xd5, 0xe1, 0x5a, 0xa7, 0xc8, 0xae, 0xf4, 0xc5, 0xb9, 0xee, 0xcb, 0xcd, 0xba, 0xed,
	0x95, 0xba, 0xdf, 0x67, 0x87, 0x9f, 0x83, 0xe5, 0x9c, 0x3b, 0x05, 0x53, 0x42, 0xb6, 0x03, 0x0f,
	0x02, 0x7d, 0x99, 0x16, 0x9c, 0x4b, 0xdc, 0xf5, 0x06, 0xde, 0xab, 0x2e, 0xf7, 0x68, 0xd7, 0xe8,
	0x22, 0xea, 0xf1, 0x5a, 0x81, 0x7a, 0xfe, 0x37, 0x03, 0x6c, 0x1c, 0xf8, 0xff, 0x2f, 0xaf, 0xf5,
	0xc3, 0x80, 0xad, 0x93, 0x49, 0x7f, 0x72, 0xac, 0xa7, 0x71, 0x3a, 0x90, 0xd9, 0x06, 0x97, 0x7e,
	0x1b, 0xcc, 0x30, 0xa2, 0xe4, 0x2c, 0x61, 0x86, 0x51, 0x75, 0x0c, 0xac, 0x75, 0xc7, 0xc0, 0x87,
	0x9a, 0xd0, 0x79, 0x91, 0xac, 0x5d, 0x9c, 0xd1, 0x0a, 0xb7, 0xbe, 0x1b, 0x50, 0xc3, 0x94, 0x36,
	0x7c, 0x83, 0xfe, 0x75, 0x3a, 0x3f, 0x0d, 0xd8, 0xc6, 0x74, 0xba, 0x41, 0x30, 0x79, 0x1b, 0x0e,
	0x5e, 0xf7, 0x8e, 0xee, 0x41, 0x52, 0x1f, 0xa0, 0x81, 0xef, 0x07, 0x4d, 0xad, 0xaf, 0xb3, 0x2f,
	0x9b, 0x4d, 0xae, 0x09, 0x6e, 0x20, 0x73, 0x59, 0x66, 0xe6, 0x88, 0x12, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x13, 0xdf, 0x4d, 0xab, 0x5f, 0x06, 0x00, 0x00,
}
