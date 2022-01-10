// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

/*
Package proto_csmsg is a generated protocol buffer package.

It is generated from these files:
	demo.proto

It has these top-level messages:
	CS_Login
	SC_LogoutNtf
	PlayerColor
	Position
	Vec3
	PlayerInfo
	SC_Login
	SC_NtfLogin
	InputState
	CS_SyncMove
	SC_SyncMove
*/
package proto_csmsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MSG_ID int32

const (
	MSG_ID_MSG_ID_CS_None      MSG_ID = 0
	MSG_ID_MSG_ID_CS_Login     MSG_ID = 10001
	MSG_ID_MSG_ID_SC_LogIn     MSG_ID = 10002
	MSG_ID_MSG_ID_SC_NtfLogin  MSG_ID = 10003
	MSG_ID_MSG_ID_SC_LogoutNtf MSG_ID = 10004
	MSG_ID_MSG_ID_CS_SyncMove  MSG_ID = 20002
	MSG_ID_MSG_ID_SC_SyncMove  MSG_ID = 20003
)

var MSG_ID_name = map[int32]string{
	0:     "MSG_ID_CS_None",
	10001: "MSG_ID_CS_Login",
	10002: "MSG_ID_SC_LogIn",
	10003: "MSG_ID_SC_NtfLogin",
	10004: "MSG_ID_SC_LogoutNtf",
	20002: "MSG_ID_CS_SyncMove",
	20003: "MSG_ID_SC_SyncMove",
}
var MSG_ID_value = map[string]int32{
	"MSG_ID_CS_None":      0,
	"MSG_ID_CS_Login":     10001,
	"MSG_ID_SC_LogIn":     10002,
	"MSG_ID_SC_NtfLogin":  10003,
	"MSG_ID_SC_LogoutNtf": 10004,
	"MSG_ID_CS_SyncMove":  20002,
	"MSG_ID_SC_SyncMove":  20003,
}

func (x MSG_ID) String() string {
	return proto.EnumName(MSG_ID_name, int32(x))
}
func (MSG_ID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// MSG_ID_CS_Login = 10001;
type CS_Login struct {
	Account string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *CS_Login) Reset()                    { *m = CS_Login{} }
func (m *CS_Login) String() string            { return proto.CompactTextString(m) }
func (*CS_Login) ProtoMessage()               {}
func (*CS_Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CS_Login) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

// MSG_ID_SC_LogoutNtf = 10004;
type SC_LogoutNtf struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
}

func (m *SC_LogoutNtf) Reset()                    { *m = SC_LogoutNtf{} }
func (m *SC_LogoutNtf) String() string            { return proto.CompactTextString(m) }
func (*SC_LogoutNtf) ProtoMessage()               {}
func (*SC_LogoutNtf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SC_LogoutNtf) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type PlayerColor struct {
	A float32 `protobuf:"fixed32,1,opt,name=a" json:"a,omitempty"`
	B float32 `protobuf:"fixed32,2,opt,name=b" json:"b,omitempty"`
	C float32 `protobuf:"fixed32,3,opt,name=c" json:"c,omitempty"`
	D float32 `protobuf:"fixed32,4,opt,name=d" json:"d,omitempty"`
}

func (m *PlayerColor) Reset()                    { *m = PlayerColor{} }
func (m *PlayerColor) String() string            { return proto.CompactTextString(m) }
func (*PlayerColor) ProtoMessage()               {}
func (*PlayerColor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PlayerColor) GetA() float32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *PlayerColor) GetB() float32 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *PlayerColor) GetC() float32 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *PlayerColor) GetD() float32 {
	if m != nil {
		return m.D
	}
	return 0
}

type Position struct {
	X float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z" json:"z,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

type Vec3 struct {
	X float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z" json:"z,omitempty"`
}

func (m *Vec3) Reset()                    { *m = Vec3{} }
func (m *Vec3) String() string            { return proto.CompactTextString(m) }
func (*Vec3) ProtoMessage()               {}
func (*Vec3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Vec3) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vec3) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vec3) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

type PlayerInfo struct {
	Addr        string       `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Position    *Position    `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	PlayerColor *PlayerColor `protobuf:"bytes,3,opt,name=playerColor" json:"playerColor,omitempty"`
	Vel         *Position    `protobuf:"bytes,4,opt,name=vel" json:"vel,omitempty"`
}

func (m *PlayerInfo) Reset()                    { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()               {}
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PlayerInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *PlayerInfo) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *PlayerInfo) GetPlayerColor() *PlayerColor {
	if m != nil {
		return m.PlayerColor
	}
	return nil
}

func (m *PlayerInfo) GetVel() *Position {
	if m != nil {
		return m.Vel
	}
	return nil
}

// MSG_ID_SC_LogIn = 10002;
type SC_Login struct {
	Mine   *PlayerInfo   `protobuf:"bytes,1,opt,name=mine" json:"mine,omitempty"`
	Theirs []*PlayerInfo `protobuf:"bytes,2,rep,name=theirs" json:"theirs,omitempty"`
}

func (m *SC_Login) Reset()                    { *m = SC_Login{} }
func (m *SC_Login) String() string            { return proto.CompactTextString(m) }
func (*SC_Login) ProtoMessage()               {}
func (*SC_Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SC_Login) GetMine() *PlayerInfo {
	if m != nil {
		return m.Mine
	}
	return nil
}

func (m *SC_Login) GetTheirs() []*PlayerInfo {
	if m != nil {
		return m.Theirs
	}
	return nil
}

// MSG_ID_SC_NtfLogin = 10003;
type SC_NtfLogin struct {
	Player *PlayerInfo `protobuf:"bytes,1,opt,name=player" json:"player,omitempty"`
}

func (m *SC_NtfLogin) Reset()                    { *m = SC_NtfLogin{} }
func (m *SC_NtfLogin) String() string            { return proto.CompactTextString(m) }
func (*SC_NtfLogin) ProtoMessage()               {}
func (*SC_NtfLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SC_NtfLogin) GetPlayer() *PlayerInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

type InputState struct {
	Up    bool `protobuf:"varint,1,opt,name=up" json:"up,omitempty"`
	Down  bool `protobuf:"varint,2,opt,name=down" json:"down,omitempty"`
	Left  bool `protobuf:"varint,3,opt,name=left" json:"left,omitempty"`
	Right bool `protobuf:"varint,4,opt,name=right" json:"right,omitempty"`
}

func (m *InputState) Reset()                    { *m = InputState{} }
func (m *InputState) String() string            { return proto.CompactTextString(m) }
func (*InputState) ProtoMessage()               {}
func (*InputState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *InputState) GetUp() bool {
	if m != nil {
		return m.Up
	}
	return false
}

func (m *InputState) GetDown() bool {
	if m != nil {
		return m.Down
	}
	return false
}

func (m *InputState) GetLeft() bool {
	if m != nil {
		return m.Left
	}
	return false
}

func (m *InputState) GetRight() bool {
	if m != nil {
		return m.Right
	}
	return false
}

type CS_SyncMove struct {
	Moves []*CS_SyncMove_Move `protobuf:"bytes,1,rep,name=moves" json:"moves,omitempty"`
}

func (m *CS_SyncMove) Reset()                    { *m = CS_SyncMove{} }
func (m *CS_SyncMove) String() string            { return proto.CompactTextString(m) }
func (*CS_SyncMove) ProtoMessage()               {}
func (*CS_SyncMove) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CS_SyncMove) GetMoves() []*CS_SyncMove_Move {
	if m != nil {
		return m.Moves
	}
	return nil
}

type CS_SyncMove_Move struct {
	InputState *InputState `protobuf:"bytes,1,opt,name=inputState" json:"inputState,omitempty"`
	// 用于服务器模拟以及返回通知客户端
	Tm    float32 `protobuf:"fixed32,2,opt,name=tm" json:"tm,omitempty"`
	Delta float32 `protobuf:"fixed32,3,opt,name=delta" json:"delta,omitempty"`
}

func (m *CS_SyncMove_Move) Reset()                    { *m = CS_SyncMove_Move{} }
func (m *CS_SyncMove_Move) String() string            { return proto.CompactTextString(m) }
func (*CS_SyncMove_Move) ProtoMessage()               {}
func (*CS_SyncMove_Move) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

func (m *CS_SyncMove_Move) GetInputState() *InputState {
	if m != nil {
		return m.InputState
	}
	return nil
}

func (m *CS_SyncMove_Move) GetTm() float32 {
	if m != nil {
		return m.Tm
	}
	return 0
}

func (m *CS_SyncMove_Move) GetDelta() float32 {
	if m != nil {
		return m.Delta
	}
	return 0
}

// MSG_ID_SC_SyncMove = 20003;
type SC_SyncMove struct {
	Diffs []*SC_SyncMove_PlayerDiff `protobuf:"bytes,1,rep,name=diffs" json:"diffs,omitempty"`
}

func (m *SC_SyncMove) Reset()                    { *m = SC_SyncMove{} }
func (m *SC_SyncMove) String() string            { return proto.CompactTextString(m) }
func (*SC_SyncMove) ProtoMessage()               {}
func (*SC_SyncMove) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SC_SyncMove) GetDiffs() []*SC_SyncMove_PlayerDiff {
	if m != nil {
		return m.Diffs
	}
	return nil
}

type SC_SyncMove_PlayerDiff struct {
	Addr    string                               `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	PosDiff *SC_SyncMove_PlayerDiff_PositionDiff `protobuf:"bytes,2,opt,name=posDiff" json:"posDiff,omitempty"`
}

func (m *SC_SyncMove_PlayerDiff) Reset()                    { *m = SC_SyncMove_PlayerDiff{} }
func (m *SC_SyncMove_PlayerDiff) String() string            { return proto.CompactTextString(m) }
func (*SC_SyncMove_PlayerDiff) ProtoMessage()               {}
func (*SC_SyncMove_PlayerDiff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 0} }

func (m *SC_SyncMove_PlayerDiff) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SC_SyncMove_PlayerDiff) GetPosDiff() *SC_SyncMove_PlayerDiff_PositionDiff {
	if m != nil {
		return m.PosDiff
	}
	return nil
}

type SC_SyncMove_PlayerDiff_PositionDiff struct {
	Pos *Position `protobuf:"bytes,1,opt,name=pos" json:"pos,omitempty"`
	Tm  float32   `protobuf:"fixed32,2,opt,name=tm" json:"tm,omitempty"`
	Vel *Position `protobuf:"bytes,3,opt,name=vel" json:"vel,omitempty"`
}

func (m *SC_SyncMove_PlayerDiff_PositionDiff) Reset()         { *m = SC_SyncMove_PlayerDiff_PositionDiff{} }
func (m *SC_SyncMove_PlayerDiff_PositionDiff) String() string { return proto.CompactTextString(m) }
func (*SC_SyncMove_PlayerDiff_PositionDiff) ProtoMessage()    {}
func (*SC_SyncMove_PlayerDiff_PositionDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0, 0}
}

func (m *SC_SyncMove_PlayerDiff_PositionDiff) GetPos() *Position {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *SC_SyncMove_PlayerDiff_PositionDiff) GetTm() float32 {
	if m != nil {
		return m.Tm
	}
	return 0
}

func (m *SC_SyncMove_PlayerDiff_PositionDiff) GetVel() *Position {
	if m != nil {
		return m.Vel
	}
	return nil
}

func init() {
	proto.RegisterType((*CS_Login)(nil), "proto_csmsg.CS_Login")
	proto.RegisterType((*SC_LogoutNtf)(nil), "proto_csmsg.SC_LogoutNtf")
	proto.RegisterType((*PlayerColor)(nil), "proto_csmsg.PlayerColor")
	proto.RegisterType((*Position)(nil), "proto_csmsg.Position")
	proto.RegisterType((*Vec3)(nil), "proto_csmsg.Vec3")
	proto.RegisterType((*PlayerInfo)(nil), "proto_csmsg.PlayerInfo")
	proto.RegisterType((*SC_Login)(nil), "proto_csmsg.SC_Login")
	proto.RegisterType((*SC_NtfLogin)(nil), "proto_csmsg.SC_NtfLogin")
	proto.RegisterType((*InputState)(nil), "proto_csmsg.InputState")
	proto.RegisterType((*CS_SyncMove)(nil), "proto_csmsg.CS_SyncMove")
	proto.RegisterType((*CS_SyncMove_Move)(nil), "proto_csmsg.CS_SyncMove.Move")
	proto.RegisterType((*SC_SyncMove)(nil), "proto_csmsg.SC_SyncMove")
	proto.RegisterType((*SC_SyncMove_PlayerDiff)(nil), "proto_csmsg.SC_SyncMove.PlayerDiff")
	proto.RegisterType((*SC_SyncMove_PlayerDiff_PositionDiff)(nil), "proto_csmsg.SC_SyncMove.PlayerDiff.PositionDiff")
	proto.RegisterEnum("proto_csmsg.MSG_ID", MSG_ID_name, MSG_ID_value)
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xec, 0xa4, 0xa9, 0xbf, 0xeb, 0xaa, 0x98, 0x69, 0x51, 0xad, 0x48, 0x48, 0x95, 0x41,
	0x02, 0x81, 0x94, 0x94, 0x04, 0x09, 0xc1, 0x82, 0x05, 0x2e, 0x42, 0x41, 0xd4, 0xaa, 0xc6, 0x12,
	0x8b, 0x6e, 0x2c, 0xc7, 0x1e, 0x27, 0x96, 0x62, 0x8f, 0x65, 0x4f, 0x42, 0xd3, 0xb7, 0xe0, 0x67,
	0xd9, 0x15, 0xac, 0xba, 0x62, 0xcb, 0xd3, 0xf0, 0x2c, 0x68, 0x66, 0xec, 0xd8, 0x11, 0x29, 0x85,
	0x4d, 0x3b, 0x67, 0xee, 0x99, 0x7b, 0xcf, 0x3d, 0x39, 0x09, 0x40, 0x48, 0x12, 0xda, 0xcb, 0x72,
	0xca, 0x28, 0xd2, 0xc5, 0x3f, 0x2f, 0x28, 0x92, 0x62, 0x62, 0xdd, 0x07, 0xcd, 0x76, 0xbd, 0x77,
	0x74, 0x12, 0xa7, 0xc8, 0x84, 0x6d, 0x3f, 0x08, 0xe8, 0x3c, 0x65, 0xa6, 0x72, 0xa8, 0x3c, 0xfc,
	0x1f, 0x57, 0xd0, 0xb2, 0x60, 0xc7, 0xb5, 0x39, 0x8b, 0xce, 0x99, 0xc3, 0x22, 0x84, 0xa0, 0xed,
	0x87, 0x61, 0x5e, 0xd2, 0xc4, 0xd9, 0x7a, 0x0d, 0xfa, 0xe9, 0xcc, 0x5f, 0x92, 0xdc, 0xa6, 0x33,
	0x9a, 0xa3, 0x1d, 0x50, 0x7c, 0x51, 0x57, 0xb1, 0xe2, 0x73, 0x34, 0x36, 0x55, 0x89, 0xc6, 0x1c,
	0x05, 0x66, 0x4b, 0xa2, 0x80, 0xa3, 0xd0, 0x6c, 0x4b, 0x14, 0x5a, 0x4f, 0x41, 0x3b, 0xa5, 0x45,
	0xcc, 0x62, 0x9a, 0xf2, 0xca, 0x79, 0xd5, 0xe3, 0x9c, 0xa3, 0x65, 0xd5, 0x63, 0xc9, 0xd1, 0x45,
	0xd5, 0xe3, 0xc2, 0x3a, 0x82, 0xf6, 0x7b, 0x12, 0x0c, 0xff, 0xe1, 0xc5, 0x0f, 0x05, 0x40, 0xea,
	0x1d, 0xa5, 0x11, 0xdd, 0xb4, 0x11, 0x7a, 0x02, 0x5a, 0x56, 0x4a, 0x11, 0x5d, 0xf4, 0xc1, 0x9d,
	0x5e, 0xc3, 0xbb, 0x5e, 0xa5, 0x13, 0xaf, 0x68, 0xe8, 0x05, 0xe8, 0x59, 0x6d, 0x82, 0x98, 0xa6,
	0x0f, 0xcc, 0xf5, 0x57, 0x75, 0x1d, 0x37, 0xc9, 0xe8, 0x01, 0xb4, 0x16, 0x64, 0x26, 0x9c, 0xb8,
	0x76, 0x12, 0x67, 0x58, 0x53, 0xd0, 0xe4, 0xa7, 0x11, 0xa7, 0xe8, 0x31, 0xb4, 0x93, 0x38, 0x25,
	0x42, 0xb7, 0x3e, 0x38, 0xd8, 0x30, 0x89, 0xaf, 0x87, 0x05, 0x09, 0xf5, 0xa1, 0xc3, 0xa6, 0x24,
	0xce, 0x0b, 0x53, 0x3d, 0x6c, 0xfd, 0x89, 0x5e, 0xd2, 0xac, 0x97, 0xa0, 0xbb, 0xb6, 0xe7, 0xb0,
	0x48, 0x0e, 0xeb, 0x43, 0x47, 0x0a, 0xbe, 0x69, 0x5c, 0x49, 0xb3, 0xce, 0x00, 0x46, 0x69, 0x36,
	0x67, 0x2e, 0xf3, 0x19, 0x41, 0xbb, 0xa0, 0xce, 0x33, 0xf1, 0x54, 0xc3, 0xea, 0x3c, 0xe3, 0x9e,
	0x87, 0xf4, 0x83, 0xf4, 0x56, 0xc3, 0xe2, 0xcc, 0xef, 0x66, 0x24, 0x62, 0xc2, 0x39, 0x0d, 0x8b,
	0x33, 0xda, 0x87, 0xad, 0x3c, 0x9e, 0x4c, 0x99, 0xb0, 0x46, 0xc3, 0x12, 0x58, 0x57, 0x0a, 0xe8,
	0xb6, 0xeb, 0xb9, 0xcb, 0x34, 0x38, 0xa1, 0x0b, 0x82, 0x86, 0xb0, 0x95, 0xd0, 0x05, 0x29, 0x4c,
	0x45, 0xec, 0x76, 0x77, 0x4d, 0x5b, 0x83, 0xd8, 0xe3, 0x7f, 0xb0, 0xe4, 0x76, 0x09, 0xb4, 0xc5,
	0xe3, 0x67, 0x00, 0xf1, 0x4a, 0xe8, 0xc6, 0xed, 0xea, 0x3d, 0x70, 0x83, 0xca, 0x77, 0x62, 0x49,
	0x99, 0x31, 0x95, 0x25, 0x5c, 0x6b, 0x48, 0x66, 0xcc, 0x2f, 0x83, 0x26, 0x81, 0x75, 0xa5, 0x0a,
	0x23, 0x57, 0x5a, 0x9f, 0xc3, 0x56, 0x18, 0x47, 0x51, 0xa5, 0xf5, 0xde, 0xda, 0xa4, 0x06, 0xb1,
	0xf4, 0xf4, 0x38, 0x8e, 0x22, 0x2c, 0x5f, 0x74, 0x7f, 0xae, 0x72, 0xcb, 0x6f, 0x37, 0xe6, 0xf6,
	0x2d, 0x6c, 0x67, 0xb4, 0xe0, 0xe5, 0x32, 0xb6, 0x47, 0x7f, 0xd1, 0x7f, 0x95, 0x31, 0x31, 0xac,
	0x6a, 0xd0, 0xcd, 0x60, 0xa7, 0x59, 0xe0, 0x21, 0xcd, 0x68, 0x51, 0x3a, 0x74, 0x5d, 0x48, 0x33,
	0x5a, 0xfc, 0x66, 0x4c, 0x99, 0xee, 0xd6, 0x4d, 0xe9, 0x7e, 0xf4, 0x5d, 0x81, 0xce, 0x89, 0xfb,
	0xc6, 0x1b, 0x1d, 0x23, 0x04, 0xbb, 0xf2, 0xe4, 0xd9, 0xae, 0xe7, 0xd0, 0x94, 0x18, 0xff, 0xa1,
	0x7d, 0xb8, 0x55, 0xdf, 0x89, 0x58, 0x1a, 0x1f, 0x9d, 0xc6, 0xad, 0xfc, 0x66, 0x8c, 0x52, 0xe3,
	0x93, 0x83, 0x0e, 0x00, 0xd5, 0xb7, 0x55, 0x8a, 0x8d, 0xcf, 0x0e, 0x32, 0x61, 0x6f, 0x8d, 0x2e,
	0x7f, 0xd6, 0x8c, 0x2f, 0xbc, 0x82, 0xea, 0xf6, 0x95, 0x4d, 0xc6, 0xd7, 0x4b, 0xa5, 0x51, 0x69,
	0x18, 0x68, 0x7c, 0xbb, 0x54, 0x5e, 0xed, 0x9d, 0xdd, 0x16, 0xeb, 0xf4, 0x1b, 0x4b, 0x8d, 0x3b,
	0x02, 0x0c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x55, 0xb3, 0xab, 0x62, 0x7a, 0x05, 0x00, 0x00,
}
