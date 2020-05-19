package main

import (
	"./zLog"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type UserList struct {
	id        int
	uid       int // 充值用户uid
	initDate  string
	lastDate  string
	days      int
	uid2      int // 免费用户uid
	initDate2 string
	lastDate2 string
	days2     int
	matchType int
	dayNum    int
}

// 获取玩家持续时间的日期列表
func getTimeList(start string, days int) []string {
	result := make([]string, 0)
	//fmt.Println(" start " ,start)
	startTime, err := time.ParseInLocation("2006-01-02T00:00:00Z", start, time.Local)
	if err != nil {
		fmt.Println("", err.Error())
	}
	//fmt.Println("startTime : ",startTime)
	for i := 0; i < days; i++ {
		time := startTime.AddDate(0, 0, i)
		timeString := time.Format("2006-01-02")
		//fmt.Println("", timeString )
		result = append(result, timeString)
	}
	return result
}

var RecordTimeDict = []string{
	"GameCoinChangeRecord_",
	"GameDiamondChangeRecord_",
	"GameItemChangeRecord_",
	"GameLotteryChangeRecord_",
	"GameScoreChangeRecord_",
	"HDBZExchangeInfo_",
	"HunGameChipRecord_",
	"RecordArenaStarBalanceInfo_",
	"RecordArenaStarSignUpInfo_",
	"RecordGame_MiniGame_",
	"RecordGrantTreasure_",
	"RecordLogon_",
	"RecordUserBombJackpotChange_",
	"RecordUserCommonJackpotChange_",
	"RecordUserRealGoldStockChange_",
	"RecordWorldGodJoinReward_",
	"FishTideUserRecord_",
	"RecordWorldGodReward_",
	"RecordUserInout_",
	"UserLuckySevenDayRecord_"}

// 获取表的字段
func GetTableKeys(name string) string {
	switch name {
	case "GameCoinChangeRecord_":
		return "UserID,KindID,ServerID,ClientKind,ChangeCoin,Coin,Insure,OprAcc,ChangeReson,RecordTime,TableArea,CoinIndb,InsureIndb,IsEmail,Type,SubType,Extend,iDitchId"
	case "GameDiamondChangeRecord_":
		return "UserID,KindID,ServerID,ClientKind,ChangeDiamond,Diamond,OprAcc,ChangeReson,RecordTime,TableArea,DiamondIndb,IsEmail,Type,SubType,Extend,iDitchId"
	case "GameItemChangeRecord_":
		return "UserID,KindID,ServerID,ClientKind,ItemID,ItemNum,OprAcc,ChangeReson,RecordTime,ItemIndbNum,GetScore,MasterID,IsEmail,Type,SubType,Extend,IsBigMG,iDitchId"
	case "GameLotteryChangeRecord_":
		return "UserID,KindID,ServerID,ClientKind,ChangeLottery,Lottery,OprAcc,ChangeReson,RecordTime,TableArea,LotteryIndb,IsEmail,Type,SubType,Extend,iDitchId"
	case "GameScoreChangeRecord_":
		return "UserID,KindID,ServerID,ClientKind,ChangeScore,Score,Insure,OprAcc,ChangeReson,RecordTime,TableArea,ScoreIndb,InsureIndb,IsEmail,Type,SubType,Extend,iDitchId"
	case "HDBZExchangeInfo_":
		return "ServerID,UserID,RecordTime,ExchangeTypeID,ExchangeNum,ExhangeItemType1,ExhangeItemNum1,ExhangeItemType2,ExhangeItemNum2,ExhangeItemType3,ExhangeItemNum3,ExhangeItemType4,ExhangeItemNum4,BaseRewardType,BaseRewardSubType,BaseRewardQuantity,RandomRewardType,RandomRewardSubType,RandomRewardQuantity,GiftRewardType1,GiftRewardSubType1,GiftRewardQuantity1,GiftRewardType2,GiftRewardSubType2,GiftRewardQuantity2"
	case "HunGameChipRecord_":
		return " UserID,KindID,ServerID,ClientKind,GameNum,ChipTotal,WinScore,ChipItem,HitDst,RecordTime,TableArea,UserScore,UserInsure,ControlID,WinLottery,UserLottery,WinDiamond,UserDiamond,GmControlStatus,FishKindID,WinCoin,UserCoin,PumpStoreChange,CurrentStore,WinCoinExtra,iDitchId,status,Mul,FishMul"
	case "RecordArenaStarBalanceInfo_":
		return "UserID,KindID,ServerID,ClientKind,ArenaID,PlayTimes,BaseScore,ExtraScore,TotalScore,RecordTime"
	case "RecordArenaStarSignUpInfo_":
		return "UserID,KindID,ServerID,ClientKind,ArenaID,CostBaseType,CostItemType,CostItemNum,MissileRankScore,SignUpTimes,RecordTime"
	case "RecordGame_MiniGame_":
		return "UserID,KindID,ServerID,ClientKind,TableID,GameNum,AreaChip1,AreaChip2,AreaChip3,AreaChip4,AreaResult1,AreaResult2,AreaResult3,AreaResult4,WinScore,ChipScore,PumpStoreChange,CurrentStore,WinCoinExtra,SystemForceWin,RecordTime"
	case "RecordGrantTreasure_":
		return "MasterID,ClientIP,CollectDate,UserID,CurGold,AddGold,Reason,is_proxy"
	case "RecordLogon_":
		return "UserID,ClientKind,LogonTime,IPAddr,LogonMachine,LogonChannelID,LogonMacID,LogonIMIE,DeviceType,ClientVersion"
	case "RecordUserBombJackpotChange_":
		return "UserID,KindID,ServerID,Type,ChangeValue,CurrentValue,RecordDate,Reason"
	case "RecordUserCommonJackpotChange_":
		return "UserID,KindID,ServerID,Type,ChangeValue,CurrentValue,RecordDate,Reason"
	case "RecordUserRealGoldStockChange_":
		return "UserID,KindID,ServerID,Type,ChangeValue,CurrentValue,RecordDate,Reason"
	case "RecordWorldGodJoinReward_":
		return "ServerID,UserID,GodID,GodTmpID,AwardStatus,AwardType1,AwardNum1,AwardType2,AwardNum2,AwardType3,AwardNum3,JoinAwardScore,UpdateTime"
	case "FishTideUserRecord_":
		return "ServerID,TableID,UserID,RecordTime"
	case "RecordWorldGodReward_":
		return "UserID,GodID,GodTmpID,Cost,ItemID,ItemAmount,CreateTime,CatchRate"
	case "RecordUserInout_":
		return "UserID,KindID,ServerID,EnterTime,EnterScore,EnterGrade,EnterInsure,EnterUserMedal,EnterLoveliness,EnterMachine,EnterClientIP,LeaveTime,LeaveReason,LeaveMachine,LeaveClientIP,Score,Grade,Insure,Revenue,WinCount,LostCount,DrawCount,FleeCount,UserMedal,LoveLiness,Experience,PlayTimeCount,OnLineTimeCount,ClientKind,EnterLottery,Lottery,ChipInCount,EnterDiamond,Diamond,LeaveRoomTime"
	case "UserLuckySevenDayRecord_":
		return "ServerID,UserID,ActionType,ActionNotice,ActionDate,ActionTime,AwardMainType1,AwardSubType1,AwardQuality1,AwardMainType2,AwardSubType2,AwardQuality2,AwardMainType3,AwardSubType3,AwardQuality3"

	default:
		zLog.PrintfLogger("GetTableKeys Error where %s", name)
		return ""

	}
}

// 获取表的字段， 并且经过处理之后的
func GetTableKeysDeal(name string, userInfo UserList) string {
	var allKeysDeal = ""
	allKeys := GetTableKeys(name)
	strTmp := strings.Replace(allKeys, "UserID", strconv.Itoa(userInfo.uid), -1)		// 首先把所有keys 替换uid

	switch name {
	case "GameCoinChangeRecord_", "GameDiamondChangeRecord_","GameItemChangeRecord_","GameLotteryChangeRecord_",
	"GameScoreChangeRecord_","HDBZExchangeInfo_","HunGameChipRecord_","RecordArenaStarBalanceInfo_","RecordArenaStarSignUpInfo_","RecordGame_MiniGame_":
		// 替换 RecordTime
		allKeysDeal = strings.Replace(strTmp, "RecordTime", fmt.Sprintf("dateadd(day,%d,RecordTime) as RecordTime", userInfo.dayNum), -1)
		break
	case "RecordGrantTreasure_":
		allKeysDeal = strings.Replace(strTmp, "CollectDate", fmt.Sprintf("dateadd(day,%d,CollectDate) as CollectDate", userInfo.dayNum), -1)
		break
	case "RecordLogon_":
		allKeysDeal = strings.Replace(strTmp, "LogonTime", fmt.Sprintf("dateadd(day,%d,LogonTime) as LogonTime", userInfo.dayNum), -1)
		break
	case "RecordUserBombJackpotChange_", "RecordUserCommonJackpotChange_", "RecordUserRealGoldStockChange_":
		allKeysDeal = strings.Replace(strTmp, "RecordDate", fmt.Sprintf("dateadd(day,%d,RecordDate) as RecordDate", userInfo.dayNum), -1)
		break
	case "RecordWorldGodJoinReward_":
		allKeysDeal = strings.Replace(strTmp, "UpdateTime", fmt.Sprintf("dateadd(day,%d,UpdateTime) as UpdateTime", userInfo.dayNum), -1)
		break
	case "FishTideUserRecord_":
		allKeysDeal = strings.Replace(strTmp, "RecordTime", fmt.Sprintf("RecordTime + ( %d * 86400 ) AS RecordTime ", userInfo.dayNum), -1)
		break
	case "RecordWorldGodReward_":
		allKeysDeal = strings.Replace(strTmp, "CreateTime", fmt.Sprintf("CreateTime + ( %d * 86400 ) AS CreateTime ", userInfo.dayNum), -1)
		break
	case "RecordUserInout_":
		allKeysDeal = strings.Replace(strTmp, "EnterTime", fmt.Sprintf("dateadd(day,%d,EnterTime) as EnterTime", userInfo.dayNum), -1)
		allKeysDeal = strings.Replace(strTmp, "LeaveTime", fmt.Sprintf("dateadd(day,%d,LeaveTime) as LeaveTime", userInfo.dayNum), -1)
		allKeysDeal = strings.Replace(strTmp, "LeaveRoomTime", fmt.Sprintf("LeaveRoomTime + ( %d * 86400 ) AS LeaveRoomTime ", userInfo.dayNum), -1)
		break
	case "UserLuckySevenDayRecord_":
		allKeysDeal = strings.Replace(strTmp, "ActionDate", fmt.Sprintf("dateadd(day,%d,ActionDate) as ActionDate", userInfo.dayNum), -1)
		allKeysDeal = strings.Replace(strTmp, "ActionTime", fmt.Sprintf("ActionTime + ( %d * 86400 ) AS ActionTime ", userInfo.dayNum), -1)
	default:
		zLog.PrintfLogger("GetTableKeysDeal Error where %s", name)
		allKeysDeal =  ""

	}

	return allKeysDeal

}
