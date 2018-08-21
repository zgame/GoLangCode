// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CMD_DDZ_Game.CMD

/*
Package CMD_DDZ is a generated protocol buffer package.

Namespace: MESSAGE

It is generated from these files:
	CMD_DDZ_Game.CMD

It has these top-level messages:
	CMD_GameInfo_S
	CMD_SUB_C_ReqStart
	CMD_GameStart_S
	CMD_SUB_C_CallScore
	CMD_CallScore_S
	CMD_BankerInfo_S
	CMD_SUB_C_OutCard
	CMD_OutCard_S
	CMD_SUB_C_PassCard
	CMD_PassCard_S
	CMD_StatusPlay_S
	CMD_GameConclude_S
	CMD_SysMessage_S
*/
package CMD_DDZ

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

type CMD_GameInfo_S struct {
	UserScore         *int64   `protobuf:"varint,1,opt,name=user_score,json=userScore" json:"user_score,omitempty"`
	DoublePrice       *int64   `protobuf:"varint,2,opt,name=double_price,json=doublePrice" json:"double_price,omitempty"`
	DoubleRate        *int32   `protobuf:"varint,3,opt,name=double_rate,json=doubleRate" json:"double_rate,omitempty"`
	NoteCardPrice     *int64   `protobuf:"varint,4,opt,name=note_card_price,json=noteCardPrice" json:"note_card_price,omitempty"`
	ChipValue         []uint32 `protobuf:"varint,5,rep,name=chip_value,json=chipValue" json:"chip_value,omitempty"`
	LimitRate         *int32   `protobuf:"varint,6,opt,name=limit_rate,json=limitRate" json:"limit_rate,omitempty"`
	BombRate          *int32   `protobuf:"varint,7,opt,name=bomb_rate,json=bombRate" json:"bomb_rate,omitempty"`
	RocketRate        *int32   `protobuf:"varint,8,opt,name=rocket_rate,json=rocketRate" json:"rocket_rate,omitempty"`
	SpringRate        *int32   `protobuf:"varint,9,opt,name=spring_rate,json=springRate" json:"spring_rate,omitempty"`
	AdverseSpringRate *int32   `protobuf:"varint,10,opt,name=adverse_spring_rate,json=adverseSpringRate" json:"adverse_spring_rate,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *CMD_GameInfo_S) Reset()                    { *m = CMD_GameInfo_S{} }
func (m *CMD_GameInfo_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_GameInfo_S) ProtoMessage()               {}
func (*CMD_GameInfo_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CMD_GameInfo_S) GetUserScore() int64 {
	if m != nil && m.UserScore != nil {
		return *m.UserScore
	}
	return 0
}

func (m *CMD_GameInfo_S) GetDoublePrice() int64 {
	if m != nil && m.DoublePrice != nil {
		return *m.DoublePrice
	}
	return 0
}

func (m *CMD_GameInfo_S) GetDoubleRate() int32 {
	if m != nil && m.DoubleRate != nil {
		return *m.DoubleRate
	}
	return 0
}

func (m *CMD_GameInfo_S) GetNoteCardPrice() int64 {
	if m != nil && m.NoteCardPrice != nil {
		return *m.NoteCardPrice
	}
	return 0
}

func (m *CMD_GameInfo_S) GetChipValue() []uint32 {
	if m != nil {
		return m.ChipValue
	}
	return nil
}

func (m *CMD_GameInfo_S) GetLimitRate() int32 {
	if m != nil && m.LimitRate != nil {
		return *m.LimitRate
	}
	return 0
}

func (m *CMD_GameInfo_S) GetBombRate() int32 {
	if m != nil && m.BombRate != nil {
		return *m.BombRate
	}
	return 0
}

func (m *CMD_GameInfo_S) GetRocketRate() int32 {
	if m != nil && m.RocketRate != nil {
		return *m.RocketRate
	}
	return 0
}

func (m *CMD_GameInfo_S) GetSpringRate() int32 {
	if m != nil && m.SpringRate != nil {
		return *m.SpringRate
	}
	return 0
}

func (m *CMD_GameInfo_S) GetAdverseSpringRate() int32 {
	if m != nil && m.AdverseSpringRate != nil {
		return *m.AdverseSpringRate
	}
	return 0
}

// 请求开始
type CMD_SUB_C_ReqStart struct {
	ChipValue        *uint32 `protobuf:"varint,1,opt,name=chip_value,json=chipValue" json:"chip_value,omitempty"`
	UseNoteCard      *bool   `protobuf:"varint,2,opt,name=use_note_card,json=useNoteCard" json:"use_note_card,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMD_SUB_C_ReqStart) Reset()                    { *m = CMD_SUB_C_ReqStart{} }
func (m *CMD_SUB_C_ReqStart) String() string            { return proto.CompactTextString(m) }
func (*CMD_SUB_C_ReqStart) ProtoMessage()               {}
func (*CMD_SUB_C_ReqStart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CMD_SUB_C_ReqStart) GetChipValue() uint32 {
	if m != nil && m.ChipValue != nil {
		return *m.ChipValue
	}
	return 0
}

func (m *CMD_SUB_C_ReqStart) GetUseNoteCard() bool {
	if m != nil && m.UseNoteCard != nil {
		return *m.UseNoteCard
	}
	return false
}

// 开始发送扑克
type CMD_GameStart_S struct {
	CardData         []uint32 `protobuf:"varint,1,rep,name=card_data,json=cardData" json:"card_data,omitempty"`
	UserDiamond      *int64   `protobuf:"varint,2,opt,name=user_diamond,json=userDiamond" json:"user_diamond,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CMD_GameStart_S) Reset()                    { *m = CMD_GameStart_S{} }
func (m *CMD_GameStart_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_GameStart_S) ProtoMessage()               {}
func (*CMD_GameStart_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CMD_GameStart_S) GetCardData() []uint32 {
	if m != nil {
		return m.CardData
	}
	return nil
}

func (m *CMD_GameStart_S) GetUserDiamond() int64 {
	if m != nil && m.UserDiamond != nil {
		return *m.UserDiamond
	}
	return 0
}

// 用户叫分
type CMD_SUB_C_CallScore struct {
	CallScore        *uint32 `protobuf:"varint,1,opt,name=call_score,json=callScore" json:"call_score,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMD_SUB_C_CallScore) Reset()                    { *m = CMD_SUB_C_CallScore{} }
func (m *CMD_SUB_C_CallScore) String() string            { return proto.CompactTextString(m) }
func (*CMD_SUB_C_CallScore) ProtoMessage()               {}
func (*CMD_SUB_C_CallScore) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CMD_SUB_C_CallScore) GetCallScore() uint32 {
	if m != nil && m.CallScore != nil {
		return *m.CallScore
	}
	return 0
}

// 用户叫分
type CMD_CallScore_S struct {
	CurUser          *uint32 `protobuf:"varint,1,opt,name=cur_user,json=curUser" json:"cur_user,omitempty"`
	CallScore        *uint32 `protobuf:"varint,2,opt,name=call_score,json=callScore" json:"call_score,omitempty"`
	ErrorType        *uint32 `protobuf:"varint,3,opt,name=error_type,json=errorType" json:"error_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMD_CallScore_S) Reset()                    { *m = CMD_CallScore_S{} }
func (m *CMD_CallScore_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_CallScore_S) ProtoMessage()               {}
func (*CMD_CallScore_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CMD_CallScore_S) GetCurUser() uint32 {
	if m != nil && m.CurUser != nil {
		return *m.CurUser
	}
	return 0
}

func (m *CMD_CallScore_S) GetCallScore() uint32 {
	if m != nil && m.CallScore != nil {
		return *m.CallScore
	}
	return 0
}

func (m *CMD_CallScore_S) GetErrorType() uint32 {
	if m != nil && m.ErrorType != nil {
		return *m.ErrorType
	}
	return 0
}

// 庄家信息
type CMD_BankerInfo_S struct {
	BankerUser       *uint32  `protobuf:"varint,1,opt,name=banker_user,json=bankerUser" json:"banker_user,omitempty"`
	CurUser          *uint32  `protobuf:"varint,2,opt,name=cur_user,json=curUser" json:"cur_user,omitempty"`
	BankerScore      *uint32  `protobuf:"varint,3,opt,name=banker_score,json=bankerScore" json:"banker_score,omitempty"`
	ValidCardData    []uint32 `protobuf:"varint,4,rep,name=valid_card_data,json=validCardData" json:"valid_card_data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CMD_BankerInfo_S) Reset()                    { *m = CMD_BankerInfo_S{} }
func (m *CMD_BankerInfo_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_BankerInfo_S) ProtoMessage()               {}
func (*CMD_BankerInfo_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CMD_BankerInfo_S) GetBankerUser() uint32 {
	if m != nil && m.BankerUser != nil {
		return *m.BankerUser
	}
	return 0
}

func (m *CMD_BankerInfo_S) GetCurUser() uint32 {
	if m != nil && m.CurUser != nil {
		return *m.CurUser
	}
	return 0
}

func (m *CMD_BankerInfo_S) GetBankerScore() uint32 {
	if m != nil && m.BankerScore != nil {
		return *m.BankerScore
	}
	return 0
}

func (m *CMD_BankerInfo_S) GetValidCardData() []uint32 {
	if m != nil {
		return m.ValidCardData
	}
	return nil
}

// 用户出牌
type CMD_SUB_C_OutCard struct {
	CardCount        *uint32  `protobuf:"varint,1,opt,name=card_count,json=cardCount" json:"card_count,omitempty"`
	CardData         []uint32 `protobuf:"varint,2,rep,name=card_data,json=cardData" json:"card_data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CMD_SUB_C_OutCard) Reset()                    { *m = CMD_SUB_C_OutCard{} }
func (m *CMD_SUB_C_OutCard) String() string            { return proto.CompactTextString(m) }
func (*CMD_SUB_C_OutCard) ProtoMessage()               {}
func (*CMD_SUB_C_OutCard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CMD_SUB_C_OutCard) GetCardCount() uint32 {
	if m != nil && m.CardCount != nil {
		return *m.CardCount
	}
	return 0
}

func (m *CMD_SUB_C_OutCard) GetCardData() []uint32 {
	if m != nil {
		return m.CardData
	}
	return nil
}

// 用户出牌
type CMD_OutCard_S struct {
	CardCount        *uint32  `protobuf:"varint,1,opt,name=card_count,json=cardCount" json:"card_count,omitempty"`
	CurUser          *uint32  `protobuf:"varint,2,opt,name=cur_user,json=curUser" json:"cur_user,omitempty"`
	OutUser          *uint32  `protobuf:"varint,3,opt,name=out_user,json=outUser" json:"out_user,omitempty"`
	ValidCardData    []uint32 `protobuf:"varint,4,rep,name=valid_card_data,json=validCardData" json:"valid_card_data,omitempty"`
	CardType         *uint32  `protobuf:"varint,5,opt,name=card_type,json=cardType" json:"card_type,omitempty"`
	ErrorType        *uint32  `protobuf:"varint,6,opt,name=error_type,json=errorType" json:"error_type,omitempty"`
	AutoCardCount    *uint32  `protobuf:"varint,7,opt,name=auto_card_count,json=autoCardCount" json:"auto_card_count,omitempty"`
	AutoCardData     []uint32 `protobuf:"varint,8,rep,name=auto_card_data,json=autoCardData" json:"auto_card_data,omitempty"`
	TurnOver         *bool    `protobuf:"varint,9,opt,name=turn_over,json=turnOver" json:"turn_over,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CMD_OutCard_S) Reset()                    { *m = CMD_OutCard_S{} }
func (m *CMD_OutCard_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_OutCard_S) ProtoMessage()               {}
func (*CMD_OutCard_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CMD_OutCard_S) GetCardCount() uint32 {
	if m != nil && m.CardCount != nil {
		return *m.CardCount
	}
	return 0
}

func (m *CMD_OutCard_S) GetCurUser() uint32 {
	if m != nil && m.CurUser != nil {
		return *m.CurUser
	}
	return 0
}

func (m *CMD_OutCard_S) GetOutUser() uint32 {
	if m != nil && m.OutUser != nil {
		return *m.OutUser
	}
	return 0
}

func (m *CMD_OutCard_S) GetValidCardData() []uint32 {
	if m != nil {
		return m.ValidCardData
	}
	return nil
}

func (m *CMD_OutCard_S) GetCardType() uint32 {
	if m != nil && m.CardType != nil {
		return *m.CardType
	}
	return 0
}

func (m *CMD_OutCard_S) GetErrorType() uint32 {
	if m != nil && m.ErrorType != nil {
		return *m.ErrorType
	}
	return 0
}

func (m *CMD_OutCard_S) GetAutoCardCount() uint32 {
	if m != nil && m.AutoCardCount != nil {
		return *m.AutoCardCount
	}
	return 0
}

func (m *CMD_OutCard_S) GetAutoCardData() []uint32 {
	if m != nil {
		return m.AutoCardData
	}
	return nil
}

func (m *CMD_OutCard_S) GetTurnOver() bool {
	if m != nil && m.TurnOver != nil {
		return *m.TurnOver
	}
	return false
}

// 放弃出牌
type CMD_SUB_C_PassCard struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMD_SUB_C_PassCard) Reset()                    { *m = CMD_SUB_C_PassCard{} }
func (m *CMD_SUB_C_PassCard) String() string            { return proto.CompactTextString(m) }
func (*CMD_SUB_C_PassCard) ProtoMessage()               {}
func (*CMD_SUB_C_PassCard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// 放弃出牌
type CMD_PassCard_S struct {
	TurnOver         *bool   `protobuf:"varint,1,opt,name=turn_over,json=turnOver" json:"turn_over,omitempty"`
	CurUser          *uint32 `protobuf:"varint,2,opt,name=cur_user,json=curUser" json:"cur_user,omitempty"`
	PassUser         *uint32 `protobuf:"varint,3,opt,name=pass_user,json=passUser" json:"pass_user,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMD_PassCard_S) Reset()                    { *m = CMD_PassCard_S{} }
func (m *CMD_PassCard_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_PassCard_S) ProtoMessage()               {}
func (*CMD_PassCard_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CMD_PassCard_S) GetTurnOver() bool {
	if m != nil && m.TurnOver != nil {
		return *m.TurnOver
	}
	return false
}

func (m *CMD_PassCard_S) GetCurUser() uint32 {
	if m != nil && m.CurUser != nil {
		return *m.CurUser
	}
	return 0
}

func (m *CMD_PassCard_S) GetPassUser() uint32 {
	if m != nil && m.PassUser != nil {
		return *m.PassUser
	}
	return 0
}

// 游戏状态
type CMD_StatusPlay_S struct {
	LeftCardCount    *uint32           `protobuf:"varint,1,opt,name=left_card_count,json=leftCardCount" json:"left_card_count,omitempty"`
	RightCardCount   *uint32           `protobuf:"varint,2,opt,name=right_card_count,json=rightCardCount" json:"right_card_count,omitempty"`
	CurUser          *uint32           `protobuf:"varint,3,opt,name=cur_user,json=curUser" json:"cur_user,omitempty"`
	ValidCardData    []uint32          `protobuf:"varint,4,rep,name=valid_card_data,json=validCardData" json:"valid_card_data,omitempty"`
	MyCardData       []uint32          `protobuf:"varint,5,rep,name=my_card_data,json=myCardData" json:"my_card_data,omitempty"`
	LeftCardData     []uint32          `protobuf:"varint,6,rep,name=left_card_data,json=leftCardData" json:"left_card_data,omitempty"`
	RightCardData    []uint32          `protobuf:"varint,7,rep,name=right_card_data,json=rightCardData" json:"right_card_data,omitempty"`
	OutCardData      []uint32          `protobuf:"varint,8,rep,name=out_card_data,json=outCardData" json:"out_card_data,omitempty"`
	AutoCardData     []uint32          `protobuf:"varint,9,rep,name=auto_card_data,json=autoCardData" json:"auto_card_data,omitempty"`
	TurnOver         *bool             `protobuf:"varint,10,opt,name=turn_over,json=turnOver" json:"turn_over,omitempty"`
	CurRate          *uint32           `protobuf:"varint,11,opt,name=cur_rate,json=curRate" json:"cur_rate,omitempty"`
	Bankerinfo       *CMD_BankerInfo_S `protobuf:"bytes,12,opt,name=bankerinfo" json:"bankerinfo,omitempty"`
	GameInfo         *CMD_GameInfo_S   `protobuf:"bytes,13,opt,name=GameInfo" json:"GameInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *CMD_StatusPlay_S) Reset()                    { *m = CMD_StatusPlay_S{} }
func (m *CMD_StatusPlay_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_StatusPlay_S) ProtoMessage()               {}
func (*CMD_StatusPlay_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CMD_StatusPlay_S) GetLeftCardCount() uint32 {
	if m != nil && m.LeftCardCount != nil {
		return *m.LeftCardCount
	}
	return 0
}

func (m *CMD_StatusPlay_S) GetRightCardCount() uint32 {
	if m != nil && m.RightCardCount != nil {
		return *m.RightCardCount
	}
	return 0
}

func (m *CMD_StatusPlay_S) GetCurUser() uint32 {
	if m != nil && m.CurUser != nil {
		return *m.CurUser
	}
	return 0
}

func (m *CMD_StatusPlay_S) GetValidCardData() []uint32 {
	if m != nil {
		return m.ValidCardData
	}
	return nil
}

func (m *CMD_StatusPlay_S) GetMyCardData() []uint32 {
	if m != nil {
		return m.MyCardData
	}
	return nil
}

func (m *CMD_StatusPlay_S) GetLeftCardData() []uint32 {
	if m != nil {
		return m.LeftCardData
	}
	return nil
}

func (m *CMD_StatusPlay_S) GetRightCardData() []uint32 {
	if m != nil {
		return m.RightCardData
	}
	return nil
}

func (m *CMD_StatusPlay_S) GetOutCardData() []uint32 {
	if m != nil {
		return m.OutCardData
	}
	return nil
}

func (m *CMD_StatusPlay_S) GetAutoCardData() []uint32 {
	if m != nil {
		return m.AutoCardData
	}
	return nil
}

func (m *CMD_StatusPlay_S) GetTurnOver() bool {
	if m != nil && m.TurnOver != nil {
		return *m.TurnOver
	}
	return false
}

func (m *CMD_StatusPlay_S) GetCurRate() uint32 {
	if m != nil && m.CurRate != nil {
		return *m.CurRate
	}
	return 0
}

func (m *CMD_StatusPlay_S) GetBankerinfo() *CMD_BankerInfo_S {
	if m != nil {
		return m.Bankerinfo
	}
	return nil
}

func (m *CMD_StatusPlay_S) GetGameInfo() *CMD_GameInfo_S {
	if m != nil {
		return m.GameInfo
	}
	return nil
}

// 游戏结束
type CMD_GameConclude_S struct {
	CellScore          *int64   `protobuf:"varint,1,opt,name=cell_score,json=cellScore" json:"cell_score,omitempty"`
	ChunTian           *int32   `protobuf:"varint,2,opt,name=chun_tian,json=chunTian" json:"chun_tian,omitempty"`
	FanChunTian        *int32   `protobuf:"varint,3,opt,name=fan_chun_tian,json=fanChunTian" json:"fan_chun_tian,omitempty"`
	BombCount          *int32   `protobuf:"varint,4,opt,name=bomb_count,json=bombCount" json:"bomb_count,omitempty"`
	RocketCount        *int32   `protobuf:"varint,5,opt,name=rocket_count,json=rocketCount" json:"rocket_count,omitempty"`
	LeftValidCardData  []uint32 `protobuf:"varint,6,rep,name=left_valid_card_data,json=leftValidCardData" json:"left_valid_card_data,omitempty"`
	RightValidCardData []uint32 `protobuf:"varint,7,rep,name=right_valid_card_data,json=rightValidCardData" json:"right_valid_card_data,omitempty"`
	LimitScore         *bool    `protobuf:"varint,8,opt,name=limit_score,json=limitScore" json:"limit_score,omitempty"`
	XXX_unrecognized   []byte   `json:"-"`
}

func (m *CMD_GameConclude_S) Reset()                    { *m = CMD_GameConclude_S{} }
func (m *CMD_GameConclude_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_GameConclude_S) ProtoMessage()               {}
func (*CMD_GameConclude_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CMD_GameConclude_S) GetCellScore() int64 {
	if m != nil && m.CellScore != nil {
		return *m.CellScore
	}
	return 0
}

func (m *CMD_GameConclude_S) GetChunTian() int32 {
	if m != nil && m.ChunTian != nil {
		return *m.ChunTian
	}
	return 0
}

func (m *CMD_GameConclude_S) GetFanChunTian() int32 {
	if m != nil && m.FanChunTian != nil {
		return *m.FanChunTian
	}
	return 0
}

func (m *CMD_GameConclude_S) GetBombCount() int32 {
	if m != nil && m.BombCount != nil {
		return *m.BombCount
	}
	return 0
}

func (m *CMD_GameConclude_S) GetRocketCount() int32 {
	if m != nil && m.RocketCount != nil {
		return *m.RocketCount
	}
	return 0
}

func (m *CMD_GameConclude_S) GetLeftValidCardData() []uint32 {
	if m != nil {
		return m.LeftValidCardData
	}
	return nil
}

func (m *CMD_GameConclude_S) GetRightValidCardData() []uint32 {
	if m != nil {
		return m.RightValidCardData
	}
	return nil
}

func (m *CMD_GameConclude_S) GetLimitScore() bool {
	if m != nil && m.LimitScore != nil {
		return *m.LimitScore
	}
	return false
}

// 消息类型
type CMD_SysMessage_S struct {
	SysType          *uint32 `protobuf:"varint,1,opt,name=sys_type,json=sysType" json:"sys_type,omitempty"`
	SysMessage       []byte  `protobuf:"bytes,2,opt,name=sys_message,json=sysMessage" json:"sys_message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMD_SysMessage_S) Reset()                    { *m = CMD_SysMessage_S{} }
func (m *CMD_SysMessage_S) String() string            { return proto.CompactTextString(m) }
func (*CMD_SysMessage_S) ProtoMessage()               {}
func (*CMD_SysMessage_S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CMD_SysMessage_S) GetSysType() uint32 {
	if m != nil && m.SysType != nil {
		return *m.SysType
	}
	return 0
}

func (m *CMD_SysMessage_S) GetSysMessage() []byte {
	if m != nil {
		return m.SysMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*CMD_GameInfo_S)(nil), "CMD_DDZ.CMD_GameInfo_S")
	proto.RegisterType((*CMD_SUB_C_ReqStart)(nil), "CMD_DDZ.CMD_SUB_C_ReqStart")
	proto.RegisterType((*CMD_GameStart_S)(nil), "CMD_DDZ.CMD_GameStart_S")
	proto.RegisterType((*CMD_SUB_C_CallScore)(nil), "CMD_DDZ.CMD_SUB_C_CallScore")
	proto.RegisterType((*CMD_CallScore_S)(nil), "CMD_DDZ.CMD_CallScore_S")
	proto.RegisterType((*CMD_BankerInfo_S)(nil), "CMD_DDZ.CMD_BankerInfo_S")
	proto.RegisterType((*CMD_SUB_C_OutCard)(nil), "CMD_DDZ.CMD_SUB_C_OutCard")
	proto.RegisterType((*CMD_OutCard_S)(nil), "CMD_DDZ.CMD_OutCard_S")
	proto.RegisterType((*CMD_SUB_C_PassCard)(nil), "CMD_DDZ.CMD_SUB_C_PassCard")
	proto.RegisterType((*CMD_PassCard_S)(nil), "CMD_DDZ.CMD_PassCard_S")
	proto.RegisterType((*CMD_StatusPlay_S)(nil), "CMD_DDZ.CMD_StatusPlay_S")
	proto.RegisterType((*CMD_GameConclude_S)(nil), "CMD_DDZ.CMD_GameConclude_S")
	proto.RegisterType((*CMD_SysMessage_S)(nil), "CMD_DDZ.CMD_SysMessage_S")
}

func init() { proto.RegisterFile("CMD_DDZ_Game.CMD", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 949 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x86, 0x24, 0x3b, 0xa2, 0x46, 0xa2, 0x1d, 0x2b, 0x29, 0x2a, 0x23, 0x30, 0x6c, 0x13, 0x81,
	0xa1, 0x93, 0x8b, 0xfe, 0x5c, 0x7a, 0x35, 0x05, 0xb4, 0x3d, 0x24, 0x76, 0xa9, 0x24, 0x05, 0x7a,
	0x59, 0xac, 0xc5, 0x95, 0xcd, 0x86, 0x22, 0xd5, 0xdd, 0xa5, 0x00, 0x3d, 0x42, 0x1f, 0xa1, 0xcf,
	0xd1, 0x63, 0x5f, 0xa8, 0x8f, 0x51, 0xcc, 0xcc, 0xf2, 0xb7, 0x81, 0x9b, 0x1b, 0xf9, 0xcd, 0xb7,
	0x33, 0xb3, 0xf3, 0xcd, 0xce, 0xc0, 0x34, 0x7c, 0xb3, 0x10, 0x8b, 0xc5, 0xaf, 0xe2, 0x07, 0xb9,
	0x51, 0xd7, 0x5b, 0x9d, 0xdb, 0x7c, 0x3a, 0x74, 0x58, 0xf0, 0x4f, 0x1f, 0x8e, 0xf0, 0x1b, 0x6d,
	0x3f, 0x65, 0xeb, 0x5c, 0x2c, 0xa7, 0x67, 0x00, 0x85, 0x51, 0x5a, 0x98, 0x55, 0xae, 0xd5, 0xac,
	0x77, 0xd1, 0x9b, 0x0f, 0xa2, 0x11, 0x22, 0x4b, 0x04, 0xa6, 0x97, 0x30, 0x89, 0xf3, 0xe2, 0x3e,
	0x55, 0x62, 0xab, 0x93, 0x95, 0x9a, 0xf5, 0x89, 0x30, 0x66, 0xec, 0x0e, 0xa1, 0xe9, 0x39, 0xb8,
	0x5f, 0xa1, 0xa5, 0x55, 0xb3, 0xc1, 0x45, 0x6f, 0x7e, 0x18, 0x01, 0x43, 0x91, 0xb4, 0x6a, 0x7a,
	0x05, 0xc7, 0x59, 0x6e, 0x95, 0x58, 0x49, 0x1d, 0x3b, 0x37, 0x07, 0xe4, 0xc6, 0x47, 0x38, 0x94,
	0x3a, 0x66, 0x47, 0x67, 0x00, 0xab, 0xc7, 0x64, 0x2b, 0x76, 0x32, 0x2d, 0xd4, 0xec, 0xf0, 0x62,
	0x30, 0xf7, 0xa3, 0x11, 0x22, 0x1f, 0x10, 0x40, 0x73, 0x9a, 0x6c, 0x12, 0xcb, 0x61, 0x9e, 0x51,
	0x98, 0x11, 0x21, 0x14, 0xe5, 0x15, 0x8c, 0xee, 0xf3, 0xcd, 0x3d, 0x5b, 0x87, 0x64, 0xf5, 0x10,
	0x20, 0xe3, 0x39, 0x8c, 0x75, 0xbe, 0xfa, 0xa8, 0xdc, 0x61, 0x8f, 0x73, 0x64, 0xa8, 0x24, 0x98,
	0xad, 0x4e, 0xb2, 0x07, 0x26, 0x8c, 0x98, 0xc0, 0x10, 0x11, 0xae, 0xe1, 0x85, 0x8c, 0x77, 0x4a,
	0x1b, 0x25, 0x9a, 0x44, 0x20, 0xe2, 0x89, 0x33, 0x2d, 0x2b, 0x7e, 0xf0, 0x0b, 0x2b, 0xb1, 0x7c,
	0x7f, 0x23, 0x42, 0x11, 0xa9, 0xdf, 0x97, 0x56, 0x6a, 0xdb, 0xb9, 0x22, 0x56, 0xbb, 0x75, 0xc5,
	0x00, 0xfc, 0xc2, 0x28, 0x51, 0x55, 0x8b, 0xca, 0xed, 0x45, 0xe3, 0xc2, 0xa8, 0xb7, 0xae, 0x54,
	0xc1, 0xcf, 0x70, 0x5c, 0x4a, 0x48, 0x3e, 0xc5, 0x12, 0xaf, 0x4e, 0xb5, 0x8d, 0xa5, 0x95, 0xb3,
	0x1e, 0xd5, 0xcd, 0x43, 0x60, 0x21, 0xad, 0x44, 0x05, 0x49, 0xe0, 0x38, 0x91, 0x9b, 0x3c, 0x8b,
	0x4b, 0x05, 0x11, 0x5b, 0x30, 0x14, 0x7c, 0x07, 0x2f, 0xea, 0x5c, 0x43, 0x99, 0xa6, 0xac, 0x3d,
	0x26, 0x2b, 0xd3, 0xb4, 0xd1, 0x1a, 0x98, 0x6c, 0x69, 0x0e, 0x7e, 0xe3, 0x44, 0x2a, 0xbe, 0x58,
	0x4e, 0x4f, 0xc1, 0x5b, 0x15, 0x5a, 0xa0, 0x6f, 0xc7, 0x1f, 0xae, 0x0a, 0xfd, 0xde, 0x28, 0xdd,
	0x71, 0xd6, 0xef, 0x38, 0x43, 0xb3, 0xd2, 0x3a, 0xd7, 0xc2, 0xee, 0xb7, 0xdc, 0x43, 0x7e, 0x34,
	0x22, 0xe4, 0xdd, 0x7e, 0xab, 0x82, 0x3f, 0x7b, 0xf0, 0x1c, 0x83, 0xdd, 0xc8, 0xec, 0xa3, 0xd2,
	0xae, 0x75, 0xcf, 0x61, 0x7c, 0x4f, 0xff, 0xcd, 0x80, 0xc0, 0x10, 0xc5, 0x6c, 0xa6, 0xd3, 0x6f,
	0xa7, 0x73, 0x09, 0x13, 0x77, 0x96, 0x13, 0xe2, 0x88, 0xce, 0x1f, 0xa7, 0x74, 0x05, 0xc7, 0x3b,
	0x99, 0x26, 0xb1, 0xa8, 0x6b, 0x7b, 0x40, 0xb5, 0xf5, 0x09, 0x0e, 0x5d, 0x81, 0x83, 0x5b, 0x38,
	0xa9, 0xab, 0x77, 0x5b, 0x58, 0x34, 0xf0, 0x75, 0x75, 0x2c, 0x56, 0x79, 0x91, 0xd9, 0xba, 0x76,
	0x3a, 0x0e, 0x11, 0x68, 0x2b, 0xd6, 0x6f, 0x2b, 0x16, 0xfc, 0xd5, 0x07, 0x1f, 0x3d, 0x3a, 0x5f,
	0xfc, 0x48, 0x9f, 0xf2, 0xf6, 0xc4, 0x3d, 0x4f, 0xc1, 0xcb, 0x0b, 0xcb, 0x26, 0xbe, 0xe3, 0x30,
	0x2f, 0x2c, 0x99, 0x3e, 0xf3, 0x7e, 0x55, 0xae, 0xa4, 0xcc, 0x21, 0xf9, 0xa0, 0x5c, 0x51, 0x98,
	0x8e, 0x6e, 0xcf, 0x3a, 0xba, 0x61, 0x0c, 0x59, 0xd8, 0x5c, 0x34, 0xb2, 0x1f, 0x12, 0xc7, 0x47,
	0x38, 0xac, 0x6e, 0xf0, 0x1a, 0x8e, 0x6a, 0x1e, 0xa5, 0xe2, 0x51, 0x2a, 0x93, 0x92, 0x56, 0x66,
	0x62, 0x0b, 0x9d, 0x89, 0x7c, 0xa7, 0x34, 0x3d, 0x51, 0x2f, 0xf2, 0x10, 0xb8, 0xdd, 0x29, 0x1d,
	0xbc, 0x6c, 0x3e, 0xb8, 0x3b, 0x69, 0x0c, 0xbd, 0x16, 0xc5, 0x03, 0xaf, 0xfc, 0xe7, 0xc7, 0x52,
	0x3b, 0xe9, 0xb5, 0x9d, 0x3c, 0x55, 0xc9, 0x57, 0x30, 0xda, 0x4a, 0x63, 0x9a, 0xa5, 0xf4, 0x10,
	0x40, 0x63, 0xf0, 0xc7, 0x01, 0xf7, 0xe7, 0xd2, 0x4a, 0x5b, 0x98, 0xbb, 0x54, 0xee, 0xc5, 0x12,
	0x2f, 0x9f, 0xaa, 0xb5, 0x15, 0xff, 0x91, 0xce, 0x47, 0xb8, 0xbe, 0xfc, 0x1c, 0x9e, 0xeb, 0xe4,
	0xe1, 0xb1, 0x45, 0xe4, 0xe0, 0x47, 0x84, 0x87, 0x9f, 0x14, 0x7a, 0xd0, 0x4e, 0xef, 0x73, 0xd5,
	0xbc, 0x80, 0xc9, 0x66, 0xdf, 0x20, 0xf1, 0x98, 0x85, 0xcd, 0xbe, 0x62, 0xbc, 0x86, 0xa3, 0x3a,
	0x6d, 0xe2, 0x3c, 0x63, 0x2d, 0xca, 0xac, 0x89, 0x75, 0x05, 0xc7, 0x8d, 0xa4, 0x89, 0x36, 0xe4,
	0x78, 0x55, 0xce, 0xc4, 0x0b, 0xc0, 0xc7, 0x06, 0xec, 0x0a, 0x3b, 0xce, 0x0b, 0xdb, 0x8c, 0xd8,
	0x51, 0x7f, 0xf4, 0x7f, 0xea, 0xc3, 0xa7, 0x85, 0xa3, 0x99, 0x3c, 0xae, 0x2a, 0x43, 0x93, 0xfb,
	0x7b, 0x70, 0x33, 0x21, 0xc9, 0xd6, 0xf9, 0x6c, 0x72, 0xd1, 0x9b, 0x8f, 0xbf, 0x39, 0xbd, 0x76,
	0x2b, 0xf1, 0xba, 0x3b, 0x55, 0xa2, 0x06, 0x79, 0xfa, 0x2d, 0x78, 0xe5, 0xaa, 0x9c, 0xf9, 0x74,
	0xf0, 0xcb, 0xd6, 0xc1, 0x7a, 0x8f, 0x46, 0x15, 0x31, 0xf8, 0xbb, 0xcf, 0x9d, 0x88, 0x40, 0x98,
	0x67, 0xab, 0xb4, 0x88, 0x95, 0x7b, 0xc3, 0xaa, 0x35, 0x4d, 0x07, 0xd1, 0x08, 0x11, 0x9e, 0x36,
	0xf8, 0xca, 0x1e, 0x8b, 0x4c, 0xd8, 0x44, 0x66, 0xa4, 0xfe, 0x61, 0xe4, 0x21, 0xf0, 0x2e, 0x91,
	0x19, 0x16, 0x71, 0x2d, 0x33, 0x51, 0x13, 0x78, 0xc9, 0x8e, 0xd7, 0x32, 0x0b, 0x4b, 0xce, 0x19,
	0x00, 0xed, 0x3f, 0xee, 0x9f, 0x03, 0x5e, 0x8f, 0x88, 0x70, 0xeb, 0x5c, 0xc2, 0xc4, 0x6d, 0x40,
	0x26, 0x1c, 0xb2, 0x07, 0xc6, 0x98, 0xf2, 0x15, 0xbc, 0x24, 0xe1, 0xbb, 0x7d, 0xc4, 0xf2, 0x9f,
	0xa0, 0xed, 0x43, 0xab, 0x97, 0xbe, 0x86, 0x2f, 0xb8, 0x07, 0xba, 0x27, 0xb8, 0x13, 0xa6, 0x64,
	0x6c, 0x1f, 0x39, 0x87, 0x31, 0x2f, 0x71, 0x2e, 0x83, 0x47, 0x32, 0xf2, 0x5e, 0xe7, 0xad, 0xf2,
	0xd6, 0x3d, 0xa4, 0xbd, 0x79, 0xa3, 0x8c, 0x91, 0x0f, 0x6e, 0xad, 0x98, 0xbd, 0xe1, 0x11, 0xe3,
	0xd6, 0x8a, 0xd9, 0x1b, 0x1a, 0x30, 0xb8, 0xb7, 0xf7, 0x46, 0x6c, 0x98, 0x4b, 0x85, 0x9b, 0x44,
	0x60, 0xaa, 0xd3, 0x37, 0xfd, 0x1f, 0x07, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x46, 0xc8, 0x64,
	0xd0, 0x14, 0x09, 0x00, 0x00,
}
