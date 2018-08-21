package main

//# 主命令定义
const MAIN_CMD_ID = 1100

//# 子命令定义
const SUB_C_MONITOR_REG = 10   // 注册为客户端
const SUB_S_MONITOR_ITEMS = 100   // 下发服务器列表
const SUB_S_MONITOR_STATE = 101   // 更新状态
const SUB_S_MONITOR_LOG = 103   // 添加日志
const SUB_S_NEW_MONITOR_ITEM = 104   // 新增服务器
const SUB_S_DEL_MONITOR_ITEM = 105   // 删除服务器
const SUB_C_MONITOR_KEEPLIVE = 200   // 心跳包
const SUB_S_MONITOR_KEEPLIVE = 201   // 心跳包
const SUB_C_MONITOR_CMD = 2050   // 执行命令
const SUB_S_MONITOR_CMD = 2051   // 执行命令结果

//  日志等级
const LogLevelInfo = 0x1
const LogLevelNormal = 0x2
const LogLevelWarning = 0x4
const LogLevelException = 0x8
const LogLevelDebug = 0x10
const LogLevelCritical = 0x20


////# TCPHead定义
//const cbDataKind = 0   //数据类型
//const cbCheckCode = 1   //效验字段
//const wPacketSize = 2   //数据大小
//const wMainCmdID = 3   // 主命令码
//const wSubCmdID = 4  // 子命令码
//const wPacketVer = 5   // 封包版本号
type TCPHeader struct{
	DateKind uint8	 //数据类型
	CheckCode uint8   //效验字段
	PackSize uint16  //数据大小
	MainCMDID uint16   // 主命令码
	SubCMDID uint16  // 子命令码
	PackerVer uint16   // 封包版本号
}


