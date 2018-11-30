package GlobalVar

import (
	"sync"
)

//--------------------------------------------------------------------------
// 全部变量的定义
//--------------------------------------------------------------------------

var (
	//LuaCodeToShare *lua.FunctionProto
	LuaReloadTime  int //lua脚本当前最新版本的时间戳，后台设置的，保存在服务器中，定期去更新一次
	Mutex sync.Mutex		// 主要用于lua逻辑调用时候的加锁
	RWMutex sync.RWMutex	// 主要用于针对map进行读写时候的锁
)
