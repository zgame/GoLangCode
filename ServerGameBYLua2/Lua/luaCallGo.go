package Lua

import (
	"github.com/yuin/gopher-lua"
	"../Utils/log"
	"../Utils/ztimer"
	"../Utils/zRedis"
	"../Utils/zMySql"
	"time"
	mysql "../Utils/mysql"
	//mysql "github.com/tengattack/gluasql/mysql"
)

//--------------------------------------------------------------------------------
// Lua调用的go函数
// 需要像下面一样，start的时候先注册进去，才可以正常调用
// L.SetGlobal("double", L.NewFunction(Double))
//--------------------------------------------------------------------------------

// 统一的go给lua调用的函数注册点
func (m *MyLua)InitResister() {
	// Lua调用go函数声明
	//m.L.SetGlobal("double", m.L.NewFunction(Double))
	m.L.SetGlobal("luaCallGoNetWorkSend", m.L.NewFunction(luaCallGoNetWorkSend))		//注册到lua 网络发送函数
	m.L.SetGlobal("luaCallGoPrintLogger", m.L.NewFunction(luaCallGoPrintLogger))		//注册到lua 日志打印
	m.L.SetGlobal("luaCallGoGetOsTimeMillisecond", m.L.NewFunction(luaCallGoGetOsTimeMillisecond))		//注册到lua 获取毫秒时间
	m.L.SetGlobal("luaCallGoCreateNewTimer", m.L.NewFunction(luaCallGoCreateNewTimer))		//注册到lua 设置定时器
	m.L.SetGlobal("luaCallGoCreateNewClockTimer", m.L.NewFunction(luaCallGoCreateNewClockTimer))		//注册到lua 设置定时器，固定时间
	m.L.SetGlobal("luaCallGoResisterUID", m.L.NewFunction(luaCallGoResisterUID))		//注册到lua 将uid注册到列表中
	m.L.SetGlobal("luaCallGoRedisSaveString", m.L.NewFunction(luaCallGoRedisSaveString))		//注册到lua redis save
	m.L.SetGlobal("luaCallGoRedisGetString", m.L.NewFunction(luaCallGoRedisGetString))		//注册到lua redis load
	m.L.SetGlobal("luaCallGoRedisDelKey", m.L.NewFunction(luaCallGoRedisDelKey))		//注册到lua redis del key
	m.L.SetGlobal("luaCallGoAddNumberToRedis", m.L.NewFunction(luaCallGoAddNumberToRedis))		//注册到lua redis add number
	m.L.SetGlobal("luaCallGoSqlSaveGameState", m.L.NewFunction(luaCallGoSqlSaveGameState))		//lua要保存房间的信息到mysql
	m.L.SetGlobal("luaCallGoSqlExec", m.L.NewFunction(luaCallGoSqlExec))		//lua 执行sql语句， 不带返回， 需要select用lua自己的mysql

	//m.L.SetGlobal("luaCallGoCreateGoroutine", m.L.NewFunction(luaCallGoCreateGoroutine))		//注册到lua 创建go协程


	luaopen_pb(m.L)		//加载protobuf的lua调用
	m.L.PreloadModule("mysql", mysql.Loader)		//加载mysql的lua调用 ，性能一般，写起来方便
}



//------------------------------------------------------------------------------------------------------------------------
// 下面是lua 和 go 的交互函数
//------------------------------------------------------------------------------------------------------------------------

//// test
//func Double(L *lua.LState) int {
//	lv := L.ToInt(1)             //第一个参数
//	lv2 :=  L.ToInt(2)			 //第二个参数
//	str := L.ToString(3)
//
//	L.Push(lua.LString(str+"  call "+strconv.Itoa(lv * lv2))) /* push result */
//
//	return 1                     /* number of results */
//}

// lua发送网络数据
func luaCallGoNetWorkSend(L *lua.LState) int {
	userId := L.ToInt(1)
	serverId := L.ToInt(2)
	mainCmd := L.ToInt(3)
	subCmd := L.ToInt(4)
	data := L.ToString(5)
	msg := L.ToString(6)
	token := L.ToInt(7)

	// lua传递过来之后， 立即开启一个新的协程去专门做发送工作
	//go func() {
	//bufferEnd := NetWork.DealSendData(data, msg, mainCmd, subCmd, 0) // token始终是0，服务器不用发token
	//_, err := Conn.Write(bufferEnd)
	//log.CheckError(err)


	var result bool
	// 发送出去
	if userId == 0 {
		// 给玩家自己回复消息
		result = GetMyServerByLSate(serverId).SendMsg(data, msg, mainCmd, subCmd, token)		// 把客户端发来的token返回给客户端，标记出这是哪个消息的返回
		//result = GetMyServerByLSate(serverId).WriteMsg(bufferEnd)
	} else {
		// 给其他玩家发送消息
		result = GetMyServerByUID(userId).SendMsg(data, msg, mainCmd, subCmd, token)	// 把客户端发来的token返回给客户端，标记出这是哪个消息的返回
		//result = GetMyServerByUID(userId).WriteMsg(bufferEnd)
	}
	//}()

	L.Push(lua.LBool(result))		 /* push result */
	//fmt.Println("lua send :" + str)
	return 1 // 返回1个参数 ， 设定2就是返回2个参数，0就是不返回
}


// user id 要注册，方便以后查询
func luaCallGoResisterUID(L * lua.LState) int  {
	uid := L.ToNumber(1)                        // 玩家uid
	serverId := L.ToNumber(2)                   //
	server := GetMyServerByLSate(int(serverId)) // my server
	luaUIDConnectMyServer[int(uid)] = server    // 进行关联 ,  因为lua是单线程跑， 所以不存在线程安全问题， 如果是go，需要加锁
	server.UserId = int(uid)                    // 保存uid
	return 0
}


//--------------------------------Utils-------------------------------------
// lua的日志处理
func luaCallGoPrintLogger(L * lua.LState) int  {
	str := L.ToString(1)
	log.PrintLogger(str)
	return 0
}

// lua 创建一个计时器
func luaCallGoCreateNewTimer(L * lua.LState) int  {
	funcName := L.ToString(1)	// 定期调用函数名字
	time1 := L.ToInt(2) 			// 时间，秒

	ztimer.TimerCheckUpdate(func() {
		GameManagerLuaHandle.GoCallLuaLogic(funcName) //定时调用函数
	},  time.Duration(time1) )


	return 0
}

// lua 创建一个到固定时间触发器
func luaCallGoCreateNewClockTimer(L * lua.LState) int  {
	funcName := L.ToString(1)	// 定期调用函数名字
	clock := L.ToInt(2) 			// 时间，几点，24小时制

	ztimer.TimerClock(func() {
		GameManagerLuaHandle.GoCallLuaLogic(funcName) //定时调用函数
	},  clock )

	return 0
}

// 获取毫秒级系统时间
func luaCallGoGetOsTimeMillisecond(L *lua.LState) int {
	L.Push(lua.LNumber(ztimer.GetOsTimeMillisecond()))
	return 1
}

//--------------------------------Redis-------------------------------------
// redis set value
func  luaCallGoRedisSaveString(L * lua.LState) int  {
	dir := L.ToString(1)
	key := L.ToString(2)
	value := L.ToString(3)
	zRedis.SaveStringToRedis(dir , key ,value )
	return 0
}
// redis get value
func  luaCallGoRedisGetString(L * lua.LState) int  {
	dir := L.ToString(1)
	key := L.ToString(2)

	value := zRedis.GetStringFromRedis(dir , key  )
	//fmt.Println("value",value)
	L.Push(lua.LString(value))
	return 1
}

// redis del key
func  luaCallGoRedisDelKey(L * lua.LState) int  {
	dir := L.ToString(1)
	key := L.ToString(2)
	zRedis.DelKeyToRedis(dir , key)
	return 0
}
// redis add number
func  luaCallGoAddNumberToRedis(L * lua.LState) int  {
	dir := L.ToString(1)
	key := L.ToString(2)
	num := L.ToInt(3)
	value := zRedis.AddNumberToRedis(dir , key, num)
	L.Push(lua.LNumber(value))
	return 1
}

//----------------------------------mysql-------------------------------------------
// lua要保存服务器的房间信息到mysql， 因为性能问题，所以用go
func luaCallGoSqlSaveGameState(L * lua.LState) int  {
	ServerIP_Port := L.ToString(1)
	gameType:= L.ToInt(2)
	tableId:= L.ToInt(3)
	FishNum:= L.ToInt(4)
	BulletNum:= L.ToInt(5)
	SeatArray:= L.ToInt(6)
	go zMySql.SqlSaveGameState(ServerIP_Port ,gameType ,tableId  ,FishNum ,BulletNum ,SeatArray )
	return 0
}

func luaCallGoSqlExec(L * lua.LState) int  {
	sql := L.ToString(1)
	go zMySql.SqlExec(sql )
	return 0
}