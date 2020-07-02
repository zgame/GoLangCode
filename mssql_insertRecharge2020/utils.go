package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)


// 随机[ min, max)
func ZRandomTo(min int, max int) int {
	if min >= max || max == 0 {
		//fmt.Println("随机数格式不正确")
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

type RechargeList struct {
	id        int			//自增编号
	//orderNo       string
	UserId		int			// 用户
	//name  string
	//payType      int
	payStatus      int
	kindId	int
	Money int
	coin     int		//金币
	giftOnceCoin int    // 赠送金币
	//giftTotalCoin    int

	giftOnePayCoin    int	// 邮件发送金币
	//createTime    int
	SuccessTime    int
	//operateUserId    int
	//remark    string
	//IP    string
	//sendStatus    int

	//OnePay    int
	//payGiftMoneyLimit    int
	//crontabPayCount    int
	//CrontabPayDate    string
	ClientKind    int
	//AppstoreEnvironment    int
	//ditchNumber    int
	//coinType    int

	//actionId    int
	//userGameId    int
	gitPackageId    int			// 礼包id
	//appStoreProductId    string
	Diamond    int			// 钻石
	giftOnceDiamond    int		// 赠送钻石
	//giftTotalDiamond    int
	giftOnePayDiamond    int		// 邮件赠送钻石
	//payDiamondMoneyLimit    int

	//orderDitch    int
	channelId    int
	//registerMachine    string
	//otherMoney    int
	//registerDate    string
	//logonMachine    string
	//vipLev	int
	//receiptUserName string
	//itemId	int
	//discountMoney int
	//payCount int

}

//// 获取玩家持续时间的日期列表
//func getTimeList(start string, days int) []string {
//	result := make([]string, 0)
//	//fmt.Println(" start " ,start)
//	startTime, err := time.ParseInLocation("2006-01-02T00:00:00Z", start, time.Local)
//	if err != nil {
//		fmt.Println("", err.Error())
//	}
//	//fmt.Println("startTime : ",startTime)
//	for i := 0; i < days; i++ {
//		time := startTime.AddDate(0, 0, i)
//		timeString := time.Format("2006-01-02")
//		//fmt.Println("", timeString )
//		result = append(result, timeString)
//	}
//	return result
//}
//
//var RecordTimeDict = []string{
//	"GameCoinChangeRecord_",		//灵力
//	"GameDiamondChangeRecord_",		// 钻石
//	"GameItemChangeRecord_",		// 道具
//	//"GameLotteryChangeRecord_",
//	"GameScoreChangeRecord_",		// 金币
//	//"HDBZExchangeInfo_",
//	//"HunGameChipRecord_",
//	//"RecordArenaStarBalanceInfo_",
//	//"RecordArenaStarSignUpInfo_",
//	//"RecordGame_MiniGame_",
//	//"RecordGrantTreasure_",
//	//"RecordLogon_",
//	//"RecordUserBombJackpotChange_",
//	//"RecordUserCommonJackpotChange_",
//	//"RecordUserRealGoldStockChange_",
//	//"RecordWorldGodJoinReward_",
//	//"FishTideUserRecord_",
//	//"RecordWorldGodReward_",
//	//"RecordUserInout_",
//	//"UserLuckySevenDayRecord_",
//}
//
//// 获取表的字段
//func GetTableKeys(name string) string {
//	switch name {
//	case "GameCoinChangeRecord_":
//		return "UserID,KindID,ServerID,ClientKind,ChangeCoin,Coin,Insure,OprAcc,ChangeReson,RecordTime,TableArea,CoinIndb,InsureIndb,IsEmail,Type,SubType,Extend,iDitchId"
//	case "GameDiamondChangeRecord_":
//		return "UserID,KindID,ServerID,ClientKind,ChangeDiamond,Diamond,OprAcc,ChangeReson,RecordTime,TableArea,DiamondIndb,IsEmail,Type,SubType,Extend,iDitchId"
//	case "GameItemChangeRecord_":
//		return "UserID,KindID,ServerID,ClientKind,ItemID,ItemNum,OprAcc,ChangeReson,RecordTime,ItemIndbNum,GetScore,MasterID,IsEmail,Type,SubType,Extend,IsBigMG,iDitchId"
//	case "GameLotteryChangeRecord_":
//		return "UserID,KindID,ServerID,ClientKind,ChangeLottery,Lottery,OprAcc,ChangeReson,RecordTime,TableArea,LotteryIndb,IsEmail,Type,SubType,Extend,iDitchId"
//	case "GameScoreChangeRecord_":
//		return "UserID,KindID,ServerID,ClientKind,ChangeScore,Score,Insure,OprAcc,ChangeReson,RecordTime,TableArea,ScoreIndb,InsureIndb,IsEmail,Type,SubType,Extend,iDitchId"
//	case "HDBZExchangeInfo_":
//		return "ServerID,UserID,RecordTime,ExchangeTypeID,ExchangeNum,ExhangeItemType1,ExhangeItemNum1,ExhangeItemType2,ExhangeItemNum2,ExhangeItemType3,ExhangeItemNum3,ExhangeItemType4,ExhangeItemNum4,BaseRewardType,BaseRewardSubType,BaseRewardQuantity,RandomRewardType,RandomRewardSubType,RandomRewardQuantity,GiftRewardType1,GiftRewardSubType1,GiftRewardQuantity1,GiftRewardType2,GiftRewardSubType2,GiftRewardQuantity2"
//	case "HunGameChipRecord_":
//		return " UserID,KindID,ServerID,ClientKind,GameNum,ChipTotal,WinScore,ChipItem,HitDst,RecordTime,TableArea,UserScore,UserInsure,ControlID,WinLottery,UserLottery,WinDiamond,UserDiamond,GmControlStatus,FishKindID,WinCoin,UserCoin,PumpStoreChange,CurrentStore,WinCoinExtra,iDitchId,status,Mul,FishMul"
//	case "RecordArenaStarBalanceInfo_":
//		return "UserID,KindID,ServerID,ClientKind,ArenaID,PlayTimes,BaseScore,ExtraScore,TotalScore,RecordTime"
//	case "RecordArenaStarSignUpInfo_":
//		return "UserID,KindID,ServerID,ClientKind,ArenaID,CostBaseType,CostItemType,CostItemNum,MissileRankScore,SignUpTimes,RecordTime"
//	case "RecordGame_MiniGame_":
//		return "UserID,KindID,ServerID,ClientKind,TableID,GameNum,AreaChip1,AreaChip2,AreaChip3,AreaChip4,AreaResult1,AreaResult2,AreaResult3,AreaResult4,WinScore,ChipScore,PumpStoreChange,CurrentStore,WinCoinExtra,SystemForceWin,RecordTime"
//	case "RecordGrantTreasure_":
//		return "MasterID,ClientIP,CollectDate,UserID,CurGold,AddGold,Reason,is_proxy"
//	case "RecordLogon_":
//		return "UserID,ClientKind,LogonTime,IPAddr,LogonMachine,LogonChannelID,LogonMacID,LogonIMIE,DeviceType,ClientVersion"
//	case "RecordUserBombJackpotChange_":
//		return "UserID,KindID,ServerID,Type,ChangeValue,CurrentValue,RecordDate,Reason"
//	case "RecordUserCommonJackpotChange_":
//		return "UserID,KindID,ServerID,Type,ChangeValue,CurrentValue,RecordDate,Reason"
//	case "RecordUserRealGoldStockChange_":
//		return "UserID,KindID,ServerID,Type,ChangeValue,CurrentValue,RecordDate,Reason"
//	case "RecordWorldGodJoinReward_":
//		return "ServerID,UserID,GodID,GodTmpID,AwardStatus,AwardType1,AwardNum1,AwardType2,AwardNum2,AwardType3,AwardNum3,JoinAwardScore,UpdateTime"
//	case "FishTideUserRecord_":
//		return "ServerID,TableID,UserID,RecordTime"
//	case "RecordWorldGodReward_":
//		return "UserID,GodID,GodTmpID,Cost,ItemID,ItemAmount,CreateTime,CatchRate"
//	case "RecordUserInout_":
//		return "UserID,KindID,ServerID,EnterTime,EnterScore,EnterGrade,EnterInsure,EnterUserMedal,EnterLoveliness,EnterMachine,EnterClientIP,LeaveTime,LeaveReason,LeaveMachine,LeaveClientIP,Score,Grade,Insure,Revenue,WinCount,LostCount,DrawCount,FleeCount,UserMedal,LoveLiness,Experience,PlayTimeCount,OnLineTimeCount,ClientKind,EnterLottery,Lottery,ChipInCount,EnterDiamond,Diamond,LeaveRoomTime"
//	case "UserLuckySevenDayRecord_":
//		return "ServerID,UserID,ActionType,ActionNotice,ActionDate,ActionTime,AwardMainType1,AwardSubType1,AwardQuality1,AwardMainType2,AwardSubType2,AwardQuality2,AwardMainType3,AwardSubType3,AwardQuality3"
//
//	default:
//		zLog.PrintfLogger("GetTableKeys Error where %s", name)
//		return ""
//
//	}
//}
//
//// 获取表的字段， 并且经过处理之后的
//func GetTableKeysDeal(name string, userInfo RechargeList) string {
//	var allKeysDeal = ""
//	//allKeys := GetTableKeys(name)
//	//strTmp := strings.Replace(allKeys, "UserID", strconv.Itoa(userInfo.uid), -1)		// 首先把所有keys 替换uid
//
//
//
//	return allKeysDeal
//
//}


// 时间戳换成字符串
func GetTimeFromInt(TimeInt int)  string{
	return time.Unix(int64(TimeInt), 0).Format("2006-01-02 15:04:05")

}




// 获取游戏库资源
func GetDataBaseBY(gameDB03 *sql.DB, userId int )  (int,int,int){

	sqlStr := fmt.Sprintf("select top(1)Score,Diamond,Coin from dbo.GameScoreInfo where UserID = %d ",  userId)
	//zLog.PrintfLogger("获取uid:%d  游戏库资源sql: %s ", userId, sqlStr)

	_, rows, _ := mssql.Query(gameDB03, sqlStr)
	for rows.Next() { // 循环遍历
		var Score int
		var Diamond int
		var Coin int
		err := rows.Scan(&Score,&Diamond,&Coin)
		if err != nil {
			zLog.PrintfLogger(" %d 游戏库资源 , %s \n", userId,  err)
			continue
		}
		//if Score >= 0 {
		//	zLog.PrintfLogger("GetDataBaseBY userId : %d,     获取数量： %d", userId,    Score)
		mssql.CloseQuery(rows)
		return Score, Diamond,Coin
		//}
	}
	//mssql.CloseQuery(rows)
	//return Score, Diamond,Coin
	return 0, 0, 0
}
// 获取游戏库资源
func GetDataBaseBYItem(gameDB03 *sql.DB, userId int ,itemId int)  int{
	sqlStr := fmt.Sprintf("select top(1)Total,Used from dbo.UserSkillInfo where UserID = %d and ItemID = %d",  userId,itemId)
	//zLog.PrintfLogger("获取uid:%d  游戏库资源sql: %s ", userId, sqlStr)

	_, rows, _ := mssql.Query(gameDB03, sqlStr)
	for rows.Next() { // 循环遍历
		var Num int
		var total int
		var used int
		err := rows.Scan(&total,&used)
		if err != nil {
			zLog.PrintfLogger(" %d 游戏库资源    , %s \n", userId,  err)
			continue
		}
		Num = total - used
		//if Num >= 0 {
		//zLog.PrintfLogger("GetDataBaseBYItem  userId : %d,    id:%d 获取数量： %d", userId,  itemId, Num)
		mssql.CloseQuery(rows)
		return Num
		//}
	}
	return 0

}