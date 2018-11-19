package GlobalVar

import (
	"github.com/yuin/gopher-lua"
)

//--------------------------------------------------------------------------
// 全部变量的定义
//--------------------------------------------------------------------------

var (
	LuaCodeToShare *lua.FunctionProto
	LuaReloadTime  int //lua脚本当前最新版本的时间戳，后台设置的，保存在服务器中，定期去更新一次

)
