package BYModel

import (
	"time"
	"../../../CSV"
)

//----------------------------------------------------------------------------------
// CommonLogin 定义的结构（类） 是基类， 具体游戏有共同的逻辑部分可以归纳到Common里面， 具体游戏可以继承，也可以重载
//----------------------------------------------------------------------------------




//// -------------------------鱼的统一定义接口-------------------------------------
//type FishInterface interface{
//	NewFish(uid int) FishInterface
//	GetFishUID() int
//	FishRun(table BYTableInterface)
//
//	CreateFish(KindId int, fishType int, offsetX float32, offsetY float32, TotalAliveTime int, speed int, PathID int, Mulriple int, CapturePro int) // 生成一条鱼
//
//	GetFishKindID() int
//	GetFishType() int
//	GetFishPathID() int
//	GetFishOffsetX() float32
//	GetFishOffsetY() float32
//	GetFishSpeed() int
//	GetFishScore() int
//}



// ------------------------鱼的结构----------------------------
type CommonFish struct {
	FishUID int			//UID
	FishKindID int		// 鱼的表中的类型， 是小丑鱼，娃娃鱼，还是锤头鲨
	FishType int		// 鱼大的分类类型， 是大鱼，小鱼，中鱼，还是boss

	DeadTime       time.Time // 过期时间
	PathID	int			// 路径
	//pathIndex int		// 路径索引
	OffsetX  float32
	OffsetY  float32
	Speed int
	//
	Mulriple int		// 倍率
	CapturePro int		// 捕获率
	//
	CreateTime time.Time		// 创建时间
	TotalAliveTime int		// 存活时间，秒
	//MovePerPointTime time.Time	// 移动一个坐标点需要的时间

}



// -------------------------构造函数-------------------------
func NewFish(uid int) *CommonFish{
	return &CommonFish{FishUID:uid}
}
// 获取鱼的uid
func (fish *CommonFish) GetFishUID() int {
	return fish.FishUID
}

func (fish *CommonFish) GetFishKindID() int {
	return fish.FishKindID
}
func (fish *CommonFish) GetFishType() int {
	return fish.FishType
}
func (fish *CommonFish) GetFishPathID() int {
	return fish.PathID
}
func (fish *CommonFish) GetFishOffsetX() float32 {
	return fish.OffsetX
}
func (fish *CommonFish) GetFishOffsetY() float32 {
	return fish.OffsetY
}
func (fish *CommonFish) GetFishSpeed() int {
	return fish.Speed
}

// 获取鱼的分数
func (fish * CommonFish)GetFishScore() int {
	re,_ := CSV.GetFishServerExcel(fish.FishKindID, CSV.FishServerExcelGoldMultiples, true)
	return re
}



// 创建一条鱼
func (fish *CommonFish)CreateFish(KindId int,fishType int,offsetX float32,offsetY float32, TotalAliveTime int, speed int ,PathID	int	 , Mulriple int , CapturePro int)  {
	fish.FishKindID = KindId
	fish.FishType = fishType
	fish.CreateTime = time.Now()
	fish.TotalAliveTime = TotalAliveTime
	fish.DeadTime = time.Now().Add(time.Second * time.Duration(fish.TotalAliveTime))
	fish.Speed = speed
	fish.PathID = PathID
	fish.Mulriple = Mulriple
	fish.CapturePro = CapturePro

	//fmt.Println("创建了一条鱼", KindId)
}





