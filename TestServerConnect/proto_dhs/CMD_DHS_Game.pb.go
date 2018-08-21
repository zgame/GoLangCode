// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CMD_DHS_Game.CMD

/*
Package CMD_DHS is a generated protocol buffer package.

Namespace: MESSAGE

It is generated from these files:
	CMD_DHS_Game.CMD

It has these top-level messages:
	TagSeaMonster
	TagUserMonsterItem
	CMD_User_SeaMonsterInfo_S
	CMD_S_USE_SUMMON_GEM
	CMD_C_USE_SUMMON_GEM
	CMD_C_DHS_USER_FIRE
	CMD_S_DHS_USER_FIRE
	CMD_S_TRRIGETER_BUFF
	CMD_C_PLAY_PUZZLE_REQ
	CMD_S_PLAY_PUZZLE_RESULT
	CMD_C_PUZZLE_REWARD_REQ
	TagPuzzleItem
	CMD_S_PUZZLE_REWARD_RESULT
*/
package CMD_DHS

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the CMD package it is being compiled against.
// A compilation error at this line likely means your copy of the
// CMD package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the CMD package

// SUB_S_USER_DHS_INFO
type TagSeaMonster struct {
	MonsterId        *int32 `protobuf:"varint,1,opt,name=monster_id,json=monsterId" json:"monster_id,omitempty"`
	MonsterHp        *int32 `protobuf:"varint,2,opt,name=monster_hp,json=monsterHp" json:"monster_hp,omitempty"`
	BulletNum        *int32 `protobuf:"varint,3,opt,name=bullet_num,json=bulletNum" json:"bullet_num,omitempty"`
	LeftTimes        *int64 `protobuf:"varint,4,opt,name=left_times,json=leftTimes" json:"left_times,omitempty"`
	MonsterMaxHp     *int32 `protobuf:"varint,5,opt,name=monster_max_hp,json=monsterMaxHp" json:"monster_max_hp,omitempty"`
	SummonTimes      *int64 `protobuf:"varint,6,opt,name=summon_times,json=summonTimes" json:"summon_times,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TagSeaMonster) Reset()                    { *m = TagSeaMonster{} }
func (m *TagSeaMonster) String() string            { return proto.CompactTextString(m) }
func (*TagSeaMonster) ProtoMessage()               {}
func (*TagSeaMonster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TagSeaMonster) GetMonsterId() int32 {
	if m != nil && m.MonsterId != nil {
		return *m.MonsterId
	}
	return 0
}

func (m *TagSeaMonster) GetMonsterHp() int32 {
	if m != nil && m.MonsterHp != nil {
		return *m.MonsterHp
	}
	return 0
}

func (m *TagSeaMonster) GetBulletNum() int32 {
	if m != nil && m.BulletNum != nil {
		return *m.BulletNum
	}
	return 0
}

func (m *TagSeaMonster) GetLeftTimes() int64 {
	if m != nil && m.LeftTimes != nil {
		return *m.LeftTimes
	}
	return 0
}

func (m *TagSeaMonster) GetMonsterMaxHp() int32 {
	if m != nil && m.MonsterMaxHp != nil {
		return *m.MonsterMaxHp
	}
	return 0
}

func (m *TagSeaMonster) GetSummonTimes() int64 {
	if m != nil && m.SummonTimes != nil {
		return *m.SummonTimes
	}
	return 0
}

type TagUserMonsterItem struct {
	ItemId           *uint32 `protobuf:"varint,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	Used             *int32  `protobuf:"varint,2,opt,name=used" json:"used,omitempty"`
	Total            *int32  `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TagUserMonsterItem) Reset()                    { *m = TagUserMonsterItem{} }
func (m *TagUserMonsterItem) String() string            { return proto.CompactTextString(m) }
func (*TagUserMonsterItem) ProtoMessage()               {}
func (*TagUserMonsterItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TagUserMonsterItem) GetItemId() uint32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *TagUserMonsterItem) GetUsed() int32 {
	if m != nil && m.Used != nil {
		return *m.Used
	}
	return 0
}

func (m *TagUserMonsterItem) GetTotal() int32 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

// SUB_S_USER_DHS_GAME_INFO
type CMD_User_SeaMonsterInfo_S struct {
	UserMonster      *TagSeaMonster        `protobuf:"bytes,1,opt,name=user_monster,json=userMonster" json:"user_monster,omitempty"`
	UserItems        []*TagUserMonsterItem `protobuf:"bytes,2,rep,name=user_items,json=userItems" json:"user_items,omitempty"`
	UserPuzzle       []int32               `protobuf:"varint,3,rep,name=user_puzzle,json=userPuzzle" json:"user_puzzle,omitempty"`
	PuzzidStart      *int64                `protobuf:"varint,4,opt,name=puzzid_start,json=puzzidStart" json:"puzzid_start,omitempty"`
	PuzzidEnd        *int64                `protobuf:"varint,5,opt,name=puzzid_end,json=puzzidEnd" json:"puzzid_end,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *CMD_User_SeaMonsterInfo_S) Reset()                    { *m = CMD_User_SeaMonsterInfo_S{} }
func (m *CMD_User_SeaMonsterInfo_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_User_SeaMonsterInfo_S) ProtoMessage()               {}
func (*CMD_User_SeaMonsterInfo_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CMD_User_SeaMonsterInfo_S) GetUserMonster() *TagSeaMonster {
	if m != nil {
		return m.UserMonster
	}
	return nil
}

func (m *CMD_User_SeaMonsterInfo_S) GetUserItems() []*TagUserMonsterItem {
	if m != nil {
		return m.UserItems
	}
	return nil
}

func (m *CMD_User_SeaMonsterInfo_S) GetUserPuzzle() []int32 {
	if m != nil {
		return m.UserPuzzle
	}
	return nil
}

func (m *CMD_User_SeaMonsterInfo_S) GetPuzzidStart() int64 {
	if m != nil && m.PuzzidStart != nil {
		return *m.PuzzidStart
	}
	return 0
}

func (m *CMD_User_SeaMonsterInfo_S) GetPuzzidEnd() int64 {
	if m != nil && m.PuzzidEnd != nil {
		return *m.PuzzidEnd
	}
	return 0
}

// 使用召唤石
type CMD_S_USE_SUMMON_GEM struct {
	ItemId           *int32 `protobuf:"varint,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMD_S_USE_SUMMON_GEM) Reset()                    { *m = CMD_S_USE_SUMMON_GEM{} }
func (m *CMD_S_USE_SUMMON_GEM) String() string            { return proto.CompactTextString(m) }
func (*CMD_S_USE_SUMMON_GEM) ProtoMessage()               {}
func (*CMD_S_USE_SUMMON_GEM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CMD_S_USE_SUMMON_GEM) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

// 使用召唤石
type CMD_C_USE_SUMMON_GEM struct {
	Result           *int32  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	BulletNum        *int32  `protobuf:"varint,2,opt,name=bullet_num,json=bulletNum" json:"bullet_num,omitempty"`
	ItemId           *uint32 `protobuf:"varint,3,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMD_C_USE_SUMMON_GEM) Reset()                    { *m = CMD_C_USE_SUMMON_GEM{} }
func (m *CMD_C_USE_SUMMON_GEM) String() string            { return proto.CompactTextString(m) }
func (*CMD_C_USE_SUMMON_GEM) ProtoMessage()               {}
func (*CMD_C_USE_SUMMON_GEM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CMD_C_USE_SUMMON_GEM) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *CMD_C_USE_SUMMON_GEM) GetBulletNum() int32 {
	if m != nil && m.BulletNum != nil {
		return *m.BulletNum
	}
	return 0
}

func (m *CMD_C_USE_SUMMON_GEM) GetItemId() uint32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

// 开火
type CMD_C_DHS_USER_FIRE struct {
	IsCrit           *int32 `protobuf:"varint,1,opt,name=is_crit,json=isCrit" json:"is_crit,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMD_C_DHS_USER_FIRE) Reset()                    { *m = CMD_C_DHS_USER_FIRE{} }
func (m *CMD_C_DHS_USER_FIRE) String() string            { return proto.CompactTextString(m) }
func (*CMD_C_DHS_USER_FIRE) ProtoMessage()               {}
func (*CMD_C_DHS_USER_FIRE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CMD_C_DHS_USER_FIRE) GetIsCrit() int32 {
	if m != nil && m.IsCrit != nil {
		return *m.IsCrit
	}
	return 0
}

// 开火
type CMD_S_DHS_USER_FIRE struct {
	Result           *int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	MonsterHp        *int64 `protobuf:"varint,2,opt,name=monster_hp,json=monsterHp" json:"monster_hp,omitempty"`
	Money            *int32 `protobuf:"varint,3,opt,name=money" json:"money,omitempty"`
	ItemId           *int32 `protobuf:"varint,4,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	ItemNum          *int32 `protobuf:"varint,5,opt,name=item_num,json=itemNum" json:"item_num,omitempty"`
	PuzzleId         *int32 `protobuf:"varint,6,opt,name=puzzle_id,json=puzzleId" json:"puzzle_id,omitempty"`
	PuzzleNum        *int32 `protobuf:"varint,7,opt,name=puzzle_num,json=puzzleNum" json:"puzzle_num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMD_S_DHS_USER_FIRE) Reset()                    { *m = CMD_S_DHS_USER_FIRE{} }
func (m *CMD_S_DHS_USER_FIRE) String() string            { return proto.CompactTextString(m) }
func (*CMD_S_DHS_USER_FIRE) ProtoMessage()               {}
func (*CMD_S_DHS_USER_FIRE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CMD_S_DHS_USER_FIRE) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *CMD_S_DHS_USER_FIRE) GetMonsterHp() int64 {
	if m != nil && m.MonsterHp != nil {
		return *m.MonsterHp
	}
	return 0
}

func (m *CMD_S_DHS_USER_FIRE) GetMoney() int32 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *CMD_S_DHS_USER_FIRE) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *CMD_S_DHS_USER_FIRE) GetItemNum() int32 {
	if m != nil && m.ItemNum != nil {
		return *m.ItemNum
	}
	return 0
}

func (m *CMD_S_DHS_USER_FIRE) GetPuzzleId() int32 {
	if m != nil && m.PuzzleId != nil {
		return *m.PuzzleId
	}
	return 0
}

func (m *CMD_S_DHS_USER_FIRE) GetPuzzleNum() int32 {
	if m != nil && m.PuzzleNum != nil {
		return *m.PuzzleNum
	}
	return 0
}

// 狂暴状态
type CMD_S_TRRIGETER_BUFF struct {
	Time             *uint32 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMD_S_TRRIGETER_BUFF) Reset()                    { *m = CMD_S_TRRIGETER_BUFF{} }
func (m *CMD_S_TRRIGETER_BUFF) String() string            { return proto.CompactTextString(m) }
func (*CMD_S_TRRIGETER_BUFF) ProtoMessage()               {}
func (*CMD_S_TRRIGETER_BUFF) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CMD_S_TRRIGETER_BUFF) GetTime() uint32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

// 请求拼图
type CMD_C_PLAY_PUZZLE_REQ struct {
	PuzzleId         *int32 `protobuf:"varint,1,opt,name=puzzle_id,json=puzzleId" json:"puzzle_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMD_C_PLAY_PUZZLE_REQ) Reset()                    { *m = CMD_C_PLAY_PUZZLE_REQ{} }
func (m *CMD_C_PLAY_PUZZLE_REQ) String() string            { return proto.CompactTextString(m) }
func (*CMD_C_PLAY_PUZZLE_REQ) ProtoMessage()               {}
func (*CMD_C_PLAY_PUZZLE_REQ) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CMD_C_PLAY_PUZZLE_REQ) GetPuzzleId() int32 {
	if m != nil && m.PuzzleId != nil {
		return *m.PuzzleId
	}
	return 0
}

type CMD_S_PLAY_PUZZLE_RESULT struct {
	Result           *int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	PuzzleId         *int32 `protobuf:"varint,2,opt,name=puzzle_id,json=puzzleId" json:"puzzle_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMD_S_PLAY_PUZZLE_RESULT) Reset()                    { *m = CMD_S_PLAY_PUZZLE_RESULT{} }
func (m *CMD_S_PLAY_PUZZLE_RESULT) String() string            { return proto.CompactTextString(m) }
func (*CMD_S_PLAY_PUZZLE_RESULT) ProtoMessage()               {}
func (*CMD_S_PLAY_PUZZLE_RESULT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CMD_S_PLAY_PUZZLE_RESULT) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *CMD_S_PLAY_PUZZLE_RESULT) GetPuzzleId() int32 {
	if m != nil && m.PuzzleId != nil {
		return *m.PuzzleId
	}
	return 0
}

// 请求拼图奖励
type CMD_C_PUZZLE_REWARD_REQ struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMD_C_PUZZLE_REWARD_REQ) Reset()                    { *m = CMD_C_PUZZLE_REWARD_REQ{} }
func (m *CMD_C_PUZZLE_REWARD_REQ) String() string            { return proto.CompactTextString(m) }
func (*CMD_C_PUZZLE_REWARD_REQ) ProtoMessage()               {}
func (*CMD_C_PUZZLE_REWARD_REQ) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type TagPuzzleItem struct {
	ItemId           *int32 `protobuf:"varint,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	ItemNum          *int32 `protobuf:"varint,2,opt,name=item_num,json=itemNum" json:"item_num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TagPuzzleItem) Reset()                    { *m = TagPuzzleItem{} }
func (m *TagPuzzleItem) String() string            { return proto.CompactTextString(m) }
func (*TagPuzzleItem) ProtoMessage()               {}
func (*TagPuzzleItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TagPuzzleItem) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *TagPuzzleItem) GetItemNum() int32 {
	if m != nil && m.ItemNum != nil {
		return *m.ItemNum
	}
	return 0
}

type CMD_S_PUZZLE_REWARD_RESULT struct {
	Result           *int32           `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Money            *int32           `protobuf:"varint,2,opt,name=money" json:"money,omitempty"`
	UserItems        []*TagPuzzleItem `protobuf:"bytes,3,rep,name=user_items,json=userItems" json:"user_items,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CMD_S_PUZZLE_REWARD_RESULT) Reset()                    { *m = CMD_S_PUZZLE_REWARD_RESULT{} }
func (m *CMD_S_PUZZLE_REWARD_RESULT) String() string            { return proto.CompactTextString(m) }
func (*CMD_S_PUZZLE_REWARD_RESULT) ProtoMessage()               {}
func (*CMD_S_PUZZLE_REWARD_RESULT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CMD_S_PUZZLE_REWARD_RESULT) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *CMD_S_PUZZLE_REWARD_RESULT) GetMoney() int32 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *CMD_S_PUZZLE_REWARD_RESULT) GetUserItems() []*TagPuzzleItem {
	if m != nil {
		return m.UserItems
	}
	return nil
}

func init() {
	proto.RegisterType((*TagSeaMonster)(nil), "CMD_DHS.tagSeaMonster")
	proto.RegisterType((*TagUserMonsterItem)(nil), "CMD_DHS.tagUserMonsterItem")
	proto.RegisterType((*CMD_User_SeaMonsterInfo_S)(nil), "CMD_DHS.CMD_User_SeaMonsterInfo_S")
	proto.RegisterType((*CMD_S_USE_SUMMON_GEM)(nil), "CMD_DHS.CMD_S_USE_SUMMON_GEM")
	proto.RegisterType((*CMD_C_USE_SUMMON_GEM)(nil), "CMD_DHS.CMD_C_USE_SUMMON_GEM")
	proto.RegisterType((*CMD_C_DHS_USER_FIRE)(nil), "CMD_DHS.CMD_C_DHS_USER_FIRE")
	proto.RegisterType((*CMD_S_DHS_USER_FIRE)(nil), "CMD_DHS.CMD_S_DHS_USER_FIRE")
	proto.RegisterType((*CMD_S_TRRIGETER_BUFF)(nil), "CMD_DHS.CMD_S_TRRIGETER_BUFF")
	proto.RegisterType((*CMD_C_PLAY_PUZZLE_REQ)(nil), "CMD_DHS.CMD_C_PLAY_PUZZLE_REQ")
	proto.RegisterType((*CMD_S_PLAY_PUZZLE_RESULT)(nil), "CMD_DHS.CMD_S_PLAY_PUZZLE_RESULT")
	proto.RegisterType((*CMD_C_PUZZLE_REWARD_REQ)(nil), "CMD_DHS.CMD_C_PUZZLE_REWARD_REQ")
	proto.RegisterType((*TagPuzzleItem)(nil), "CMD_DHS.tagPuzzleItem")
	proto.RegisterType((*CMD_S_PUZZLE_REWARD_RESULT)(nil), "CMD_DHS.CMD_S_PUZZLE_REWARD_RESULT")
}

func init() { proto.RegisterFile("CMD_DHS_Game.CMD", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x4e, 0xdb, 0x40,
	0x10, 0x95, 0x6d, 0x12, 0xc8, 0x04, 0x7a, 0xd8, 0x52, 0x30, 0x45, 0x55, 0x53, 0xab, 0x87, 0xa8,
	0x87, 0x54, 0x42, 0xed, 0xa1, 0xbd, 0x41, 0x30, 0x24, 0x12, 0x01, 0xba, 0x8e, 0x85, 0xca, 0x65,
	0xe5, 0xd6, 0x0b, 0xb5, 0x94, 0xb5, 0x23, 0x7b, 0x2d, 0x51, 0x6e, 0xfd, 0xc5, 0xaa, 0xbf, 0xd2,
	0x7b, 0x35, 0xbb, 0x6b, 0x62, 0x5b, 0x94, 0x9b, 0xf7, 0xcd, 0xce, 0x9b, 0x37, 0x6f, 0x67, 0x0c,
	0x64, 0x3c, 0x3b, 0x66, 0xc7, 0x93, 0x80, 0x9d, 0x46, 0x82, 0x8f, 0x96, 0x79, 0x26, 0x33, 0xb2,
	0x6e, 0x30, 0xef, 0x8f, 0x05, 0x5b, 0x32, 0xba, 0x0d, 0x78, 0x34, 0xcb, 0xd2, 0x42, 0xf2, 0x9c,
	0xbc, 0x02, 0x10, 0xfa, 0x93, 0x25, 0xb1, 0x6b, 0x0d, 0xac, 0x61, 0x87, 0xf6, 0x0c, 0x32, 0x8d,
	0xeb, 0xe1, 0x1f, 0x4b, 0xd7, 0x6e, 0x84, 0x27, 0x4b, 0x0c, 0x7f, 0x2b, 0x17, 0x0b, 0x2e, 0x59,
	0x5a, 0x0a, 0xd7, 0xd1, 0x61, 0x8d, 0x9c, 0x97, 0x02, 0xc3, 0x0b, 0x7e, 0x23, 0x99, 0x4c, 0x04,
	0x2f, 0xdc, 0xb5, 0x81, 0x35, 0x74, 0x68, 0x0f, 0x91, 0x39, 0x02, 0xe4, 0x2d, 0x3c, 0xab, 0xc8,
	0x45, 0x74, 0x87, 0x05, 0x3a, 0x8a, 0x61, 0xd3, 0xa0, 0xb3, 0xe8, 0x6e, 0xb2, 0x24, 0x6f, 0x60,
	0xb3, 0x28, 0x85, 0xc8, 0x52, 0x43, 0xd3, 0x55, 0x34, 0x7d, 0x8d, 0x29, 0x22, 0xef, 0x0a, 0x88,
	0x8c, 0x6e, 0xc3, 0x82, 0xe7, 0xa6, 0xad, 0xa9, 0xe4, 0x82, 0xec, 0xc2, 0x7a, 0x22, 0xb9, 0xa8,
	0xfa, 0xda, 0xa2, 0x5d, 0x3c, 0x4e, 0x63, 0x42, 0x60, 0xad, 0x2c, 0x78, 0x6c, 0xda, 0x51, 0xdf,
	0x64, 0x1b, 0x3a, 0x32, 0x93, 0xd1, 0xc2, 0x34, 0xa1, 0x0f, 0xde, 0x5f, 0x0b, 0xf6, 0xd0, 0x3b,
	0xa4, 0x66, 0x2b, 0xd7, 0xa6, 0xe9, 0x4d, 0xc6, 0x02, 0xf2, 0x09, 0x36, 0x4b, 0x0c, 0x18, 0xb9,
	0xaa, 0x4a, 0xff, 0x60, 0x67, 0x64, 0xdc, 0x1e, 0x35, 0x9c, 0xa6, 0xfd, 0x72, 0xa5, 0x8f, 0x7c,
	0x06, 0x50, 0xa9, 0xa8, 0xa8, 0x70, 0xed, 0x81, 0x33, 0xec, 0x1f, 0xec, 0xd7, 0x13, 0x5b, 0xcd,
	0xd0, 0x1e, 0x5e, 0xc7, 0xaf, 0x82, 0xbc, 0x06, 0x45, 0xc5, 0x96, 0xe5, 0xfd, 0xfd, 0x82, 0xbb,
	0xce, 0xc0, 0x19, 0x76, 0xa8, 0xa2, 0xbb, 0x54, 0x08, 0x3a, 0x86, 0xb1, 0x24, 0x66, 0x85, 0x8c,
	0x72, 0x69, 0x8c, 0xef, 0x6b, 0x2c, 0x40, 0x08, 0x5f, 0xc6, 0x5c, 0xe1, 0x69, 0xac, 0x6c, 0x77,
	0x68, 0x4f, 0x23, 0x7e, 0x1a, 0x7b, 0xef, 0x61, 0x1b, 0xb5, 0x04, 0x2c, 0x0c, 0x7c, 0x16, 0x84,
	0xb3, 0xd9, 0xc5, 0x39, 0x3b, 0xf5, 0x67, 0x6d, 0x4b, 0x3b, 0x95, 0xa5, 0xde, 0x8d, 0x4e, 0x18,
	0xb7, 0x13, 0x76, 0xa0, 0x9b, 0xf3, 0xa2, 0x5c, 0xc8, 0xea, 0xbe, 0x3e, 0xb5, 0x06, 0xc7, 0x6e,
	0x0f, 0x4e, 0xad, 0x8e, 0x53, 0x7f, 0x3a, 0x6f, 0x04, 0xcf, 0x75, 0x1d, 0x9c, 0xf0, 0x30, 0xf0,
	0x29, 0x3b, 0x99, 0x52, 0x5f, 0xdd, 0x2f, 0xd8, 0xf7, 0x3c, 0x79, 0xa8, 0x93, 0x14, 0xe3, 0x3c,
	0x91, 0xde, 0x6f, 0x4b, 0x27, 0x04, 0xad, 0x84, 0x27, 0x74, 0xb5, 0xe6, 0xdd, 0xa9, 0xcf, 0xfb,
	0x36, 0x74, 0x44, 0x96, 0xf2, 0x9f, 0xd5, 0x94, 0xa8, 0x43, 0x5d, 0xed, 0x5a, 0xdd, 0x15, 0xb2,
	0x07, 0x1b, 0x2a, 0x80, 0x3d, 0xea, 0xd1, 0x56, 0x17, 0xb1, 0xc3, 0x7d, 0xe8, 0xe9, 0xf7, 0xc3,
	0xac, 0xae, 0x8a, 0x6d, 0x68, 0x40, 0x6f, 0x9d, 0x09, 0x62, 0xe6, 0xba, 0x76, 0x47, 0x23, 0xe7,
	0xa5, 0xf0, 0xde, 0x55, 0xaf, 0x33, 0xa7, 0x74, 0x7a, 0xea, 0xcf, 0x7d, 0xca, 0x8e, 0xc2, 0x93,
	0x13, 0x9c, 0x6b, 0x5c, 0x11, 0x33, 0xed, 0xea, 0xdb, 0xfb, 0x00, 0x2f, 0xb4, 0x61, 0x97, 0x67,
	0x87, 0x5f, 0xd9, 0x65, 0x78, 0x7d, 0x7d, 0xe6, 0x33, 0xea, 0x7f, 0x69, 0x0a, 0xb0, 0x9a, 0x02,
	0xbc, 0x0b, 0x70, 0x75, 0x85, 0x66, 0x56, 0x10, 0x9e, 0xcd, 0xff, 0x6b, 0x5d, 0x83, 0xd0, 0x6e,
	0x11, 0xee, 0xc1, 0xae, 0x91, 0x51, 0x71, 0x5d, 0x1d, 0xd2, 0x63, 0x14, 0xe2, 0x8d, 0xd5, 0x2f,
	0x49, 0x8f, 0xee, 0x63, 0x7b, 0xfb, 0xb8, 0x9d, 0x76, 0xc3, 0x4e, 0xef, 0x97, 0x05, 0x2f, 0x8d,
	0xe2, 0x56, 0x81, 0x27, 0x35, 0x3f, 0xbc, 0xa7, 0x5d, 0x7f, 0xcf, 0x8f, 0x8d, 0xe5, 0x74, 0xd4,
	0x72, 0x36, 0xb6, 0x7a, 0x25, 0xb6, 0xb6, 0x97, 0x47, 0xf6, 0xc4, 0xf9, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x40, 0x4c, 0xeb, 0xb1, 0x7e, 0x05, 0x00, 0x00,
}
