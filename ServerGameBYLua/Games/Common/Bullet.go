package Common

//----------------------------------------------------------------------------------
// Common 定义的结构（类） 是基类， 具体游戏有共同的逻辑部分可以归纳到Common里面， 具体游戏可以继承，也可以重载
//----------------------------------------------------------------------------------


import (
	"time"
	//"fmt"
)

// -------------------------子弹的统一定义接口-------------------------------------
type BulletInterface interface{

	NewBullet(uid int) BulletInterface
	GetBulletUID() int
	BulletRun(table TableInterface)

	GetBulletPlayerUID() int	// 获取玩家uid
	SetBulletPlayerUID(uid int)		// 设定玩家uid

	GetLockFishID() int		//  获取锁定鱼的uid
	SetLockFishID(lockFishID int)		//  设定锁定鱼的uid
}



// ------------------------子弹的结构----------------------------
type CommonBullet struct {
	BulletUID      int			//UID
	TempID         int
	UserID         int       // 所属玩家的id
	FireAngle      int       // 发射的角度
	BulletMulriple int       // 倍率
	lockFishID     int       // 锁定鱼的ID
	DeadTime       time.Time // 过期时间

}


// -------------------------构造函数-------------------------
func (bullet *CommonBullet) NewBullet(uid int) BulletInterface{
	//创建子弹的时候，设定子弹的生存时间是20秒
	return &CommonBullet{BulletUID:uid,	DeadTime:time.Now().Add(time.Second *20 )	}
}

func (bullet *CommonBullet) GetBulletUID() int {
	return bullet.BulletUID
}

// 获取玩家uid
func (bullet *CommonBullet) GetBulletPlayerUID() int {
	return bullet.UserID
}
// 设定玩家uid
func (bullet *CommonBullet) SetBulletPlayerUID(uid int)  {
	 bullet.UserID = uid
}

//  获取锁定鱼的uid
func (bullet *CommonBullet) GetLockFishID() int {
	return bullet.lockFishID
}
//  设定锁定鱼的uid
func (bullet *CommonBullet) SetLockFishID(lockFishID int) {
	bullet.lockFishID = lockFishID
}

// -------------------------子弹的主循环------------------------
func (bullet *CommonBullet) BulletRun(table TableInterface)   {
	//fmt.Println("子弹的run", bullet.GetBulletUID())
	// 子弹的生存时间到了
	if time.Now().After(bullet.DeadTime){
		table.DelBullet(bullet.BulletUID)		// 生存时间已经到了，销毁
		return
	}
	// 子弹的移动

	// 子弹是否击中的判断

}

//--------------------------del--------------------------------






