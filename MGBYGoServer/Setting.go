package main

import (
	//"github.com/go-ini/ini"
	//"fmt"
	"./Core/Utils/zLog"
	"./Core/Utils/zIP"
	"github.com/go-ini/ini"
	"fmt"
)

// 数据库变量
var RedisAddress string		// redis 服务器地址
var RedisPass string		// redis pwd

var MySqlServerIP string		// zMysqlForLua
var MySqlServerPort string		// zMysqlForLua port
var MySqlDatabase string
var MySqlUid string
var MySqlPwd string

var SqlServerIP string		// sql server
var SqlServerDatabase string
var SqlServerUid string
var SqlServerPwd string
var SqlServerIPLog string		// sql server Log
var SqlServerDatabaseLog string
var SqlServerUidLog string
var SqlServerPwdLog string
var SqlServerIPFriend string		// sql server Friend
var SqlServerDatabaseFriend string
var SqlServerUidFriend string
var SqlServerPwdFriend string



// -WebSocketPort=8089 -SocketPort=8124
//-----------------------------本地配置文件---------------------------------------------------
func initSetting()  bool{
	f, err := ini.Load("Setting.ini")
	if err != nil{
		fmt.Println("读取配置文件出错")
		return false
	}

	//-------------------------------------------------------------------
	//if WebSocketPort == 0 {
	//	WebSocketPort, err = f.Section("Server").Key("WebSocketPort").Int()
	//}
	//if SocketPort == 0 {
	//	fmt.Println("Warning!!!! You sould write arguments like : -WebSocketPort=8089 -SocketPort=8124")
	//	SocketPort, err = f.Section("Server").Key("SocketPort").Int()
	//}

	zLog.ShowLog,err  = f.Section("Server").Key("ShowLog").Bool()
	//WebSocketServer,err  = f.Section("Server").Key("WebSocketServer").Bool()
	//SocketServer,err  = f.Section("Server").Key("SocketServer").Bool()
	RedisAddress = f.Section("Server").Key("RedisAddress").String()
	RedisPass = f.Section("Server").Key("RedisPass").String()
	ServerAddress = f.Section("Server").Key("ServerAddress").String()
	//WebSocketAddress = f.Section("Server").Key("WebSocketAddress").String()
	//GoroutineMax ,err  = f.Section("Server").Key("GoroutineMax").Int()

	MySqlServerIP = f.Section("Server").Key("MySqlServerIP").Value()
	MySqlServerPort = f.Section("Server").Key("MySqlServerPort").Value()
	MySqlDatabase = f.Section("Server").Key("MySqlDatabase").Value()
	MySqlUid = f.Section("Server").Key("MySqlUid").Value()
	MySqlPwd = f.Section("Server").Key("MySqlPwd").Value()

	SqlServerIP = f.Section("Server").Key("SqlServerIP").Value()
	SqlServerDatabase = f.Section("Server").Key("SqlServerDatabase").Value()
	SqlServerUid = f.Section("Server").Key("SqlServerUid").Value()
	SqlServerPwd = f.Section("Server").Key("SqlServerPwd").Value()

	SqlServerIPLog = f.Section("Server").Key("ConstSqlServerIP_Log").Value()
	SqlServerDatabaseLog = f.Section("Server").Key("ConstSqlServerDatabase_Log").Value()
	SqlServerUidLog = f.Section("Server").Key("ConstSqlServerUid_Log").Value()
	SqlServerPwdLog = f.Section("Server").Key("ConstSqlServerPwd_Log").Value()

	SqlServerIPFriend = f.Section("Server").Key("ConstSqlServerIP_Friend").Value()
	SqlServerDatabaseFriend = f.Section("Server").Key("ConstSqlServerDatabase_Friend").Value()
	SqlServerUidFriend = f.Section("Server").Key("ConstSqlServerUid_Friend").Value()
	SqlServerPwdFriend = f.Section("Server").Key("ConstSqlServerPwd_Friend").Value()






	zLog.CheckError(err)

	ServerAddress = string(zIP.GetInternal(-1))		// 获取本机内网ip
	fmt.Println("本机内网ip :",ServerAddress)

	return  true

}
