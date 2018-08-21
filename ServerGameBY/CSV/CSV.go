package CSV

import (
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
)
//------------------------------------------------------读CSV文件--------------------------------------------------------------------------------

// 读取Excel文件的数据
func ReadExcelFile(fileName string)  [][]string{
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("读取csv错误",fileName)
	}
	rows, err := csv.NewReader(f).ReadAll()		// rows是每一行
	f.Close()
	if err != nil {
		fmt.Println("读取行错误")
	}
	return rows
}

//------------------------------------------------------FishServerExcel--------------------------------------------------------------------------------

const(
	FishServerExcelKindid = 0			// kind
	FishServerExcelType   = 1			// type
	FishServerExcelName   = 2			// 名字
	FishServerExcelBuff   = 3			// buff

	FishServerExcelVip   = 5			// vip影响
	FishServerExcelGlory   = 6			// 荣誉值影响
	FishServerExcelNoble   = 7			// 贵族值影响
	FishServerExcelSport   = 8			// 竞技值影响
	FishServerExcelJackpot   = 9			// 奖池累计影响
	FishServerExcelIsUse   = 10			// isUse

	FishServerExcelMulriple   = 12			//  鱼的倍数
	FishServerExcelExp   = 13			//  经验值
	FishServerExceldrop   = 14			//  弹头掉落
	FishServerExcelNobleEnable   = 15			//  贵族值生效
	FishServerExcelSpeed   = 16			//  速度
	FishServerExcelCapture   = 17			//  捕获概率
	FishServerExcelMaxCapture   = 18			//  最大捕获概率

	FishServerExcelProduceSkill   = 20			//  掉落技能
	FishServerExcelProduceSkillRate   = 21			//  掉落技能概率

	FishServerExcelDiamondRate   = 23			//  钻石概率
	FishServerExcelDiamondNum   = 24			//  钻石数量
	FishServerExcelLotteryRate   = 25			//  奖券概率
	FishServerExcelLotteryNum   = 26			//  奖券数量

	FishServerExcelDistributeIntervalMin = 30
	FishServerExcelDistributeIntervalMax = 31		//生成鱼时间间隔max
	FishServerExcelCountMin   = 32
	FishServerExcelCountMax   = 33		//生成鱼数量max
	FishServerExcelBirth   = 34		//出生方式
	FishServerExcelTimeMin   = 35		//多条鱼的cd
	FishServerExcelTimeMax   = 36		//多条鱼的cd

	FishServerExcelGoldRate   = 39		//金币掉率
	FishServerExcelGoldMultiples   = 40		//金币倍数
	FishServerExcelGoldWeight   = 41		//金币权重
	FishServerExcelRandomMin   = 42		//最小随机，同类炸弹？？
	FishServerExcelRandomMax   = 43		//最小随机，同类炸弹？？
	FishServerExcelCaptureProbaility   = 44		//大厅同类炸弹鱼概率
	FishServerExcelCaptureProbaility2   = 45		//鱼潮同类炸弹鱼概率
	FishServerExcelTask   = 46		//完成任务
	FishServerExcelMiss   = 47		//miss库积累
	FishServerExcelMissMiniCannon   = 48		// 倍数，低于该值积累
	FishServerExcelMissreturn   = 49		// miss库返
	FishServerExcelForce   = 50				// 强发库积累
	FishServerExcelForceReturn   = 51				// 强发库返还
	FishServerExcelForceStop   = 52				// 强发库终止
	FishServerExcelProcess   = 53				// 流程类型
	FishServerExcelCleanValue   = 54				// 净分值回补
	FishServerExcelRechargeValue   = 55				// 充值回补
	FishServerExcelEnergyInfcuence   = 56				// 储能
	FishServerExcelBossSkill   = 57				// Boss携带技能
	FishServerExcelStrikeRange   = 58				// 暴击范围
	FishServerExcelStrikeRate   = 59				// 暴击概率
	FishServerExcelNewNoble   = 60				// 新贵族值加成

	FishServerExcelStrike   = 62				// 是否暴击
	FishServerExcelStrikeRateNew   = 63			// 新暴击概率
	FishServerExcelStrikeMulti   = 64			// 暴击倍数
	FishServerExcelStrikeWeight   = 65			// 暴击权重
	FishServerExcelForceKilled   = 66			// 保底
	FishServerExcelForceRate   = 67			// 保底系数
	FishServerExcelForceConversion   = 68			// 保底转换系数
	FishServerExcelForceMinGold   = 69			// 最低强杀金币数
	FishServerExcelForceMinBullet   = 70			// 最低强杀子弹数
	FishServerExcelForceTimes   = 71			// 强杀次数限制
	FishServerExcelCurrentCannon   = 72			// 炮的倍数限制
	FishServerExcelMaxForceCannon   = 73			// 最大强杀炮倍数
	FishServerExcelBulletDecrease   = 74			// 子弹减少系数
	FishServerExcelPriorityStock   = 75			// 新贵族值影响
	FishServerExcelStockInfluence   = 76			// 劣后隐藏库影响
	FishServerExcelGoldStockInfluence   = 77			// 弹头金币库
	FishServerExcelSameNumberMax   = 78			// 鱼上限
	FishServerExcelFishGrouping   = 79			// 鱼分组
	FishServerExcelCreateWay   = 80			// 鱼刷新方式
	FishServerExcelCreateMinCD   = 81			// 循环刷新
	FishServerExcelCreateMaxCD   = 82			// 循环刷新max
	FishServerExcelIntegral   = 83			// 是否积分鱼
	FishServerExcelIntegralMin   = 84			// 最小智慧分
	FishServerExcelIntegralMax   = 85			// 最大智慧分
	FishServerExcelEnergy   = 86			// 能量值
	FishServerExcelAnger   = 87			// 愤怒值
	FishServerExcelPersonal   = 88			// 最小满贯积分
	FishServerExcelFishserverexcel   = 89			// 最大满贯积分

)

var FishServerExcel map[int][]string		// 鱼的服务器用表读到内存中
// 加载表格
func LoadFishServerExcel() {
	fileName := "./CSV/mgby_fish_sever.csv"
	rows := ReadExcelFile(fileName)
	FishServerExcel = make(map[int][]string)

	// 把表格数据处理成哈希表，方便查询
	for i,v := range rows{
		if i>= 4 {			//去掉头部没用的信息
			kindId, _ := strconv.Atoi(v[FishServerExcelKindid])
			FishServerExcel[kindId] = v
		}
	}
}
// 获取表格的数据
func GetFishServerExcel(kindId int, column int,isInt bool) (int,string){
	str := FishServerExcel[kindId][column]
	if isInt {
		// 返回数字
		//fmt.Println("",kindId,"-", column,"表",str)
		re,_ := strconv.Atoi(str)
		//fmt.Println("re",re)
		return  re, ""
	}else{
		// 返回string
		return 0,str
	}
}



//------------------------------------------------------FishServerExcel--------------------------------------------------------------------------------
