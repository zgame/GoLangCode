package BY

import (
	"time"
	"../../CSV"
	"strconv"
	"../../Core/Utils/zRandom"
	"../../Const"
	"fmt"
)

//----------------------------------------------------------------------------------------------------------------------------------
// 生成鱼池， 这里其实是桌子的拓展， 因为只有桌子才会管理如何生成鱼
//----------------------------------------------------------------------------------------------------------------------------------

// ------------------------生成鱼池的结构， 用于将鱼分类创建----------------------------
type FishDistributeInfo struct {
	FishKindID             int       // 鱼的类型
	CreateTime             time.Time //创建时间
	DistributeIntervalTime int       // 下次生成鱼的时间是一个间隔，保存这个随机间隔

	// 同时生成多条鱼，并且路径相同，路上有时间差
	BuildNumber        int       // 生成鱼的数量
	FirstPathID        int       // 鱼的路径
	CreateType         int       // 生成下条鱼的方式， 1表示一条路径，不做位置偏移， 2表示要做位置偏移

	NextCreateTime     time.Time // 下条鱼的时间起始时间
	NextInterBuildTime int // 下条鱼的生成时间间隔
}

// ------------------------初始化鱼池的生成组----------------------------
func (table *BYTable) InitDistributeInfo(roomScore int) {
	table.DistributeArray = make([]FishDistributeInfo,0)
	table.BossDistributeArray = make([]FishDistributeInfo,0)

	startID := roomScore * 100
	endID := roomScore*100 + 100
	for fishKind, v := range CSV.FishServerExcel {
		if v[CSV.FishServerExcelIsUse] == "1" && fishKind > startID && fishKind < endID { // 用分数去匹配鱼的分类
			var Distribute FishDistributeInfo
			Distribute.FishKindID, _ = strconv.Atoi(v[CSV.FishServerExcelKindid]) // 鱼类型
			Distribute.CreateTime = time.Now()                                    // 生成时间
			Distribute.DistributeIntervalTime = GetIntervalTime(fishKind)			// 获取时间间隔
			//fmt.Println("生成间隔" ,Distribute.DistributeIntervalTime )

			fishType, _ := strconv.Atoi(v[CSV.FishServerExcelType])
			if fishType == Const.FT_BOSS{
				table.BossDistributeArray = append(table.BossDistributeArray, Distribute) // 加入到Boss鱼生成列表中
			}else {
				table.DistributeArray = append(table.DistributeArray, Distribute) // 加入到普通鱼生成列表中
			}
		}
	}
	fmt.Println("桌子初始化鱼池结束")
}

// --------------------------循环鱼池的生成组--------------------------------
func (table *BYTable) RunDistributeInfo(roomScore int) {
	//fmt.Println("循环鱼池的生成")
	for _, Distribute := range table.DistributeArray {
		kindId := Distribute.FishKindID
		// 到下一个生成时间了, 那么我们来生成鱼吧
		if time.Now().After( Distribute.CreateTime.Add(time.Millisecond *time.Duration( Distribute.DistributeIntervalTime))){
			//fmt.Println("生成时间到了")
			createType := 0		// 鱼怎么走
			buildNum := 1		//鱼生成数量
			max,_ := CSV.GetFishServerExcel(kindId,CSV.FishServerExcelCountMax,true)		//获取最大生成数量
			//fmt.Println(kindId,"生成",max)
			if max > 1 {
				//  多生成几条鱼
				buildNum = GetCount(kindId)			// 随机出来生成的数量
				//fmt.Println("随机生成",buildNum,"条鱼")
				if buildNum < 1{
					buildNum = 1
				}else{
					createType = 1			//生成一条路径的
					if buildNum >= 5 || zRandom.ZRandomPercentRate(50){
						createType = 2		// 位置要做偏移
					}
				}
				Distribute.NextCreateTime = time.Now()				//？？
				Distribute.NextInterBuildTime = GetCountFishTime(kindId)
				//fmt.Println("Distribute.NextInterBuildTime",Distribute.NextInterBuildTime)
			}
			Distribute.BuildNumber = buildNum
			Distribute.CreateTime = time.Now()		//用于下次判断
			Distribute.DistributeIntervalTime = GetIntervalTime(kindId)
			Distribute.CreateType = createType			// 创建类型
			Distribute.FirstPathID = GetPathType()		//获取路径

			// 创建鱼
			table.DistributeNewFish(Distribute,0,0)
			//fmt.Println("生成一条鱼",kindId)
		}

		// 多条鱼的判断
		if Distribute.BuildNumber > 1 {
			if time.Now().After(Distribute.NextCreateTime.Add(time.Millisecond * time.Duration(Distribute.NextInterBuildTime))) {
				//fmt.Println("生成多条鱼",kindId)
				offsetX := float32(0)
				offsetY := float32(0)
				Distribute.NextCreateTime = time.Now()
				Distribute.BuildNumber --
				Distribute.NextInterBuildTime = GetCountFishTime(kindId)
				if Distribute.CreateType == 2 { // 位置偏移
					offsetX = GetOffsetXY()[0]
					offsetY = GetOffsetXY()[1]
				}
				// 创建鱼
				table.DistributeNewFish(Distribute,offsetX,offsetY)
			}
		}
	}

}


// --------------------------循环Boss鱼池的生成组--------------------------------
func (table *BYTable) RunBossDistributeInfo(roomScore int) {
	return
	for _, Distribute := range table.BossDistributeArray {
		// 到下一个生成时间了, 那么我们来生成鱼吧
		if time.Now().After( Distribute.CreateTime.Add(time.Millisecond *time.Duration( Distribute.DistributeIntervalTime))){
			kindId := Distribute.FishKindID
			buildNum := 1		//鱼生成数量

			Distribute.BuildNumber = buildNum
			Distribute.CreateTime = time.Now()		//用于下次判断
			Distribute.DistributeIntervalTime = GetIntervalTime(kindId)
			Distribute.FirstPathID = GetPathType()		//获取路径

			// 创建鱼
			table.DistributeNewFish(Distribute,0,0)
		}
	}
}


// ------------- 具体生成鱼 --------------------
func (table *BYTable)DistributeNewFish(Distribute FishDistributeInfo, offsetX float32, offsetY float32)  {
	kindId := Distribute.FishKindID
	// 创建鱼
	fish := table.CreateFish()
	fishType, _ := CSV.GetFishServerExcel(kindId, CSV.FishServerExcelType, true)
	speed, _ := CSV.GetFishServerExcel(kindId, CSV.FishServerExcelSpeed, true)
	Mulriple, _ := CSV.GetFishServerExcel(kindId, CSV.FishServerExcelMulriple, true)
	CapturePro, _ := CSV.GetFishServerExcel(kindId, CSV.FishServerExcelCapture, true)
	totalAliveTime := 120

	//fmt.Println("创建鱼")
	fish.CreateFish(kindId, fishType, offsetX, offsetY, totalAliveTime, speed, Distribute.FirstPathID, Mulriple, CapturePro)

	// 发送给所有客户端生成鱼的消息
	table.SendNewFishes(fish)
}





//-----------------------------------------工具----------------------------------------------------------

// 获取生成时间间隔
func GetIntervalTime(kindId int) int {
	// 获取时间间隔
	DistributeIntervalMin, _ := CSV.GetFishServerExcel(kindId,CSV.FishServerExcelDistributeIntervalMin,true)
	DistributeIntervalMax, _ := CSV.GetFishServerExcel(kindId,CSV.FishServerExcelDistributeIntervalMax,true)
	return  zRandom.ZRandomTo(DistributeIntervalMin, DistributeIntervalMax)
}
// 获取生成数量间隔
func GetCount(kindId int) int {
	// 获取数量间隔
	min,_ := CSV.GetFishServerExcel(kindId,CSV.FishServerExcelCountMin,true)
	max,_ := CSV.GetFishServerExcel(kindId,CSV.FishServerExcelCountMax,true)		//获取最大生成数量
	return  zRandom.ZRandomTo(min, max)
}
// 获取多条鱼生成时间间隔
func GetCountFishTime(kindId int) int {
	// 获取数量间隔
	min,_ := CSV.GetFishServerExcel(kindId,CSV.FishServerExcelTimeMin,true)
	max,_ := CSV.GetFishServerExcel(kindId,CSV.FishServerExcelTimeMax,true)
	return  zRandom.ZRandomTo(min, max)
}

// 获取路径的类型
func GetPathType() int {
	return 1
}

// 获得路径位置偏移
func GetOffsetXY() []float32 {
	offset := [][]float32{{0,0}, {-1,1}, {-0.5,1},{0,1},{0.5,1},{1,1},
		{-1,0.5}, {-0.5,0.5},{0,0.5},{0.5,0.5},{1,0.5},{-1,1},{1,0},
		{-1,-0.5}, {-0.5,-0.5},{0,-0.5},{0.5,-0.5},{1,-0.5},
		{-1,-1}, {-0.5,-1},{0,-1},{0.5,-1},{1,-1}}

	RandOff := zRandom.ZRandom(23)
	return offset[RandOff]
}



//---------------------------------------------------------------------------------------------------