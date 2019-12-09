package Const



//////////# 主命令定义
const MAIN_CMD_ID = 1100

//////////# 子命令定义
const SUB_C_MONITOR_REG = 10       // 注册为客户端
const SUB_S_MONITOR_ITEMS = 100    // 下发服务器列表
const SUB_S_MONITOR_STATE = 101    // 更新状态
const SUB_S_MONITOR_LOG = 103      // 添加日志
const SUB_S_NEW_MONITOR_ITEM = 104 // 新增服务器
const SUB_S_DEL_MONITOR_ITEM = 105 // 删除服务器
const SUB_C_MONITOR_KEEPLIVE = 200 // 心跳包
const SUB_S_MONITOR_KEEPLIVE = 201 // 心跳包
const SUB_C_MONITOR_CMD = 2050     // 执行命令
const SUB_S_MONITOR_CMD = 2051     // 执行命令结果

//  日志等级
const LogLevelInfo = 0x1
const LogLevelNormal = 0x2
const LogLevelWarning = 0x4
const LogLevelException = 0x8
const LogLevelDebug = 0x10
const LogLevelCritical = 0x20

////////////# TCPHead定义
//constd cbDataKind = 0   //数据类型
//constd cbCheckCode = 1   //效验字段
//constd wPacketSize = 2   //数据大小
//constd wMainCmdID = 3   // 主命令码
//constd wSubCmdID = 4  // 子命令码
//constd wPacketVer = 5   // 封包版本号



const ToolTypeGameBY = true
const ToolTypeGameDWC = false



const MAX_LOG_COUNT = 60		//# Client 显示的最大消息数量
const MAX_FISH_CNT = 20		//# Client 最多保存10条鱼





//-----------------------------------------------------------------
// 关于桌子的常量定义
//-----------------------------------------------------------------
const(
	//TABLE_ON = 0 	//桌子开放，有人玩
	//TABLE_OFF = 1 	//桌子关闭， 没人玩
	TABLE_CHAIR_NOBODY = -1   // 桌子和椅子没人玩
	BY_TABLE_MAX_PLAYER = 4		// 捕鱼游戏桌子的最大人数
)

//-----------------------------------------------------------------
// 关于子弹的常量定义
//-----------------------------------------------------------------
const(
	MAX_BULLET_NUMBER = 15 		//玩家最大子弹数量
)
//-----------------------------------------------------------------
// 关于鱼的常量定义
//-----------------------------------------------------------------

const(
	FT_SMALL = 1 	//小鱼
	FT_MIDSIZE = 2 	//中型鱼
	FT_BIG = 3 	//大型鱼
	FT_BOSS = 4 	//boos鱼
)



