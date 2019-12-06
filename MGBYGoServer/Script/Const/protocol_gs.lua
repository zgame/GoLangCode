

	--#  消息类别
	MDM_GR_LOGON             = 1  --# 登陆命令
	MDM_GR_CONFIG            = 2  --# 配置命令
	MDM_GR_USER              = 3  --# 用户命令
	MDM_GR_STATUS            = 4  --# 状态命令
	MDM_GR_MANAGE            = 5  --# 管理命令
	MDM_GF_FRAME             = 6  --# 框架命令
	MDM_GF_GAME              = 7  --# 框架游戏命令
	MDM_CR_SYSTEM            = 8  --# 系统命令
	MDM_GR_ONLINE_CHECK      = 9  --# 在线检测
	MDM_GR_PURCHASE          = 10 --# 充值命令
	MDM_GR_BIND              = 11 --# 手机绑定
	MDM_GR_HEARTBEAT         = 12 --# 心跳
	MDM_GR_VIP               = 13 --# VIP
	MDM_GR_REWARD            = 14 --# 奖励命令
	MDM_GR_MATCH             = 15 --# 比赛
	MDM_GR_MISSION           = 16 --# 任务
	MDM_GR_ACHIEVEMENT_TITLE = 17 --# 称号
	MDM_GR_GUILD             = 18 --# 公会
	MDM_GF_GMCMD             = 19 --# GM指令

	--# 登录相关子协议
	SUB_GR_LOGON_USERID   = 1 --# I D 登录
	SUB_GR_LOGON_SUCCESS  = 2 --# 登录成功
	SUB_GR_LOGON_FAILURE  = 3 --# 登录失败
	SUB_GR_LOGON_FINISH   = 4 --# 登录完成
	SUB_GR_LOGON_TIMER    = 5 --# 登陆计时
	SUB_GR_UPDATE_NOTIFY  = 6 --# 版本更新
	SUB_MB_C2G_SDKBINDING = 7 --# SDK绑定
	SUB_MB_G2C_SDKBINDING = 8 --# SDK绑定返回

	--# 未知
	GLFT_NO_USER_RIGHT       = 300 --# 无用户权限
	GLFT_MANAGER_KICK        = 301 --# 被管理员踢出
	GLFT_SYSTEM_MAINTAIN     = 302 --# 系统维护
	GLFT_CANNON_UNLOCK       = 303 --# 炮台未解锁
	GLFT_SCORE_LESS_ROOM     = 304 --# 金币分数低于房间要求
	GLFT_SCORE_MORE_ROOM     = 305 --# 金币分数高于房间要求
	GLFT_ROOM_MORE_PLAYER    = 306 --# 房间已满员
	GLFT_SRL_SERVER_OPENED   = 307 --# 房间开启
	GLFT_SRL_SERVER_CLOSED   = 308 --# 房间关闭
	GLFT_SERVER_INVALIDATION = 309 --# 房间不可用
	GLFT_LOGON_MODE_ERROR    = 310 --# 登录方式错误
	GLFT_LOGON_FREQUENTLY    = 311 --# 登陆过于频繁
	GLFT_DIAMOND_LESS_ROOM   = 312 --# 钻石不足
	GLFT_LOTTERY_LESS_ROOM   = 313 --# 奖券不足
	GLFT_VIP_LESS_ROOM       = 314 --# VIP
	GLFT_UNKNOW_ERROR        = 315 --# 未知错误

	--# 配置类子协议
	SUB_GR_CONFIG_COLUMN     = 1 --# 列表配置
	SUB_GR_CONFIG_SERVER     = 2 --# 房间配置
	SUB_GR_CONFIG_PROPERTY   = 3 --# 道具配置
	SUB_GR_CONFIG_FINISH     = 4 --# 配置完成
	SUB_GR_CONFIG_USER_RIGHT = 5 --# 玩家权限

	--# 用户动作
	SUB_GR_USER_RULE               = 1  --# 用户规则
	SUB_GR_USER_LOOKON             = 2  --# 旁观请求
	SUB_GR_USER_SITDOWN            = 3  --# 坐下请求
	SUB_GR_USER_STANDUP            = 4  --# 起立请求
	SUB_GR_USER_INVITE             = 5  --# 用户邀请
	SUB_GR_USER_INVITE_REQ         = 6  --# 邀请请求
	SUB_GR_USER_REPULSE_SIT        = 7  --# 拒绝玩家坐下
	SUB_GR_USER_KICK_USER          = 8  --# 踢出用户
	SUB_GR_USER_INFO_REQ           = 9  --# 请求用户信息
	SUB_GR_USER_CHAIR_REQ          = 10 --# 请求更换位置
	SUB_GR_USER_CHAIR_INFO_REQ     = 11 --# 请求椅子用户信息
	SUB_GR_USER_ENTER              = 12 --# 用户进入
	SUB_GR_USER_SCORE              = 13 --# 用户分数
	SUB_GR_USER_STATUS             = 14 --# 用户状态
	SUB_GR_REQUEST_FAILURE         = 15 --# 请求失败
	SUB_GR_USER_LEVUP              = 16 --# 用户升级
	SUB_GR_OTHER_SCORE             = 17 --# 玩家分数
	SUB_GR_USER_CHAT               = 18 --# 聊天消息
	SUB_GR_USER_EXPRESSION         = 19 --# 表情消息
	SUB_GR_WISPER_CHAT             = 20 --# 私聊消息
	SUB_GR_WISPER_EXPRESSION       = 21 --# 私聊表情
	SUB_GR_COLLOQUY_CHAT           = 22 --# 会话消息
	SUB_GR_COLLOQUY_EXPRESSION     = 23 --# 会话表情
	SUB_GR_PROPERTY_MESSAGE        = 24 --# 道具消息
	SUB_GR_PROPERTY_EFFECT         = 25 --# 道具效应
	SUB_GF_USER_GIFT_PACK_INFO     = 26 --# 用户礼包(游戏服)
	SUB_GF_C2L_MY_GIFT_PACK_INFO   = 27 --# 请求我的礼包(游戏服)
	SUB_GF_L2C_MY_GIFT_PACK_INFO   = 28 --# 请求我的礼包返回(游戏服)
	SUB_GR_S_RED_ENVELOPE          = 29 --# 产生红包
	SUB_GR_S_USER_MESSAGE          = 30 --# 用户消息
	SUB_GR_S_USER_STANDUP          = 31 --# 用户起立返回消息
	SUB_GR_S_GM_USER_LIST_GET      = 32 --# 获取杀分抬分用户列表
	SUB_GR_S_GM_USER_LIST_BACK     = 33 --# 获取杀分抬分用户列表返回
	SUB_GR_C_REQUEST_ARENA         = 34 --# 请求竞技场数据
	SUB_GR_S_REQUEST_ARENA         = 35 --# 请求竞技场数据
	SUB_GR_C_START_ARENA           = 36 --# 重新开始大师赛
	SUB_GR_S_START_ARENA           = 37 --# 重新开始大师赛
	SUB_GR_S_GM_TOOL_OPTION        = 38 --# GM工具配置信息
	SUB_GR_C_MATCH_INVITE          = 39 --# 客户端邀战
	SUB_GR_S_MATCH_INVITE          = 40 --# 服务器广播邀战
	SUB_GR_S_SITDOWN_SUCCESS       = 41 --# 坐下成功
	SUB_GR_C_TOMORROW_PACKAGE      = 42 --# 领取明日礼包
	SUB_GR_S_TOMORROW_PACKAGE_BACK = 43 --# 领取明日礼包返回
	SUB_GR_C_CREATE_TABLE          = 44 --# 创建桌子
	SUB_GR_S_CREATE_TABLE          = 45 --# 创建桌子
	SUB_GR_ROOM_USER_CHAT          = 46 --# 用户聊天
	SUB_GR_ROOM_USER_EXPRESSION    = 47 --# 用户表情
	SUB_GR_SPEAKER_SEND_MSG        = 48 --# 小喇叭请求
	SUB_GR_SPEAKER_SEND_MSG_BACK   = 49 --# 小喇叭广播
	SUB_GR_SPEAKER_SEND_MSG_FAIL   = 50 --# 小喇叭失败
	SUB_GR_S_SPEAKER_RECORD        = 51 --# 留言板信息
	SUB_GR_C_SPEAKER_RECORD        = 52 --# 请求留言板信息

	--# 桌子信息
	SUB_GR_TABLE_INFO   = 1 --# 桌子信息
	SUB_GR_TABLE_STATUS = 2 --# 桌子状态

	--# 管理
	SUB_GR_SEND_WARNING     = 1    --# 发送警告
	SUB_GR_SEND_MESSAGE     = 2    --# 发送消息
	SUB_GR_LOOK_USER_IP     = 3    --# 查看地址
	SUB_GR_KILL_USER        = 4    --# 踢出用户
	SUB_GR_LIMIT_ACCOUNS    = 5    --# 禁用帐户
	SUB_GR_SET_USER_RIGHT   = 6    --# 权限设置
	SUB_GR_QUERY_OPTION     = 7    --# 查询设置
	SUB_GR_OPTION_SERVER    = 8    --# 房间设置
	SUB_GR_OPTION_CURRENT   = 9    --# 当前设置
	SUB_GR_LIMIT_USER_CHAT  = 10   --# 限制聊天
	SUB_GR_KICK_ALL_USER    = 11   --# 踢出用户
	SUB_GR_DISMISSGAME      = 12   --# 解散游戏
	SUB_GR_GM_RESULT        = 1000 --# GM操作结果
	SUB_GR_GM_STOP          = 1001 --# 停止服务
	SUB_GR_GM_APPLY_CONFIG  = 1002 --# 应用配置
	SUB_GR_GM_MAINTENANCE   = 1003 --# 维护
	SUB_GR_GM_KICK_ALL_USER = 1004 --# 踢出所有用户
	SUB_GR_GM_FORBID        = 1005 --# 禁止进入
	SUB_GR_GM_GET_ONLINE    = 1006 --# 获取在线人数
	SUB_GR_GM_KICK_USER     = 1007 --# 踢出用户

	--# 用户命令
	SUB_GF_GAME_OPTION           = 1  --# 游戏配置
	SUB_GF_USER_READY            = 2  --# 用户准备
	SUB_GF_LOOKON_CONFIG         = 3  --# 旁观配置
	SUB_GF_USER_CHAT             = 4  --# 用户聊天 --#  1.9之后用MDM_GR_USER 那个
	SUB_GF_USER_EXPRESSION       = 5  --# 用户表情 --#  1.9之后用MDM_GR_USER 那个
	SUB_GF_GAME_STATUS           = 6  --# 游戏状态
	SUB_GF_GAME_SCENE            = 7  --# 游戏场景
	SUB_GF_LOOKON_STATUS         = 8  --# 旁观状态
	SUB_GF_SYSTEM_MESSAGE        = 9  --# 系统消息
	SUB_GF_ACTION_MESSAGE        = 10 --# 动作消息
	SUB_GF_SPEAKER_SEND_MSG      = 11 --# 小喇叭请求 --#  1.9之后用MDM_GR_USER 那个
	SUB_GF_SPEAKER_SEND_MSG_BACK = 12 --# 小喇叭广播 --#  1.9之后用MDM_GR_USER 那个
	SUB_GF_SERVER_USER_INFO      = 13 --# 房间用户信息
	SUB_GF_USER_SKILL            = 14 --# 用户技能
	SUB_GF_GAME_SKILL            = 15 --# 游戏技能
	SUB_GF_ACTIVITY_GIFT_PRIZE   = 16 --# 夺宝活动广播
	SUB_GF_USER_SCORE            = 17 --# 用户信息
	SUB_GF_SPEAKER_SEND_MSG_FAIL = 18 --# 小喇叭失败 --#  1.9之后用MDM_GR_USER 那个
	SUB_GF_UPDATE_USER_SCORE     = 19 --# 框架改变玩家货币信息
	SUB_GF_MASTER_GUIDE_INFO     = 20 --# 用户主线引导任务信息
	SUB_GF_COMPLETED_GUIDE_INFO  = 21 --# 用户主线引导任务完成信息

	SUB_CR_SYSTEM_MESSAGE   = 1 --# 系统消息
	SUB_CR_ACTION_MESSAGE   = 2 --# 动作消息
	SUB_CR_DOWN_LOAD_MODULE = 3 --# 下载消息

	SUB_ONLINE_GET        = 1 --# 获取在线验证数据包		(在游戏服)
	SUB_ONLINE_BACK       = 2 --# 服务器发送消息包给客户端	(在游戏服)
	SUB_REFRESH_ONE_DAY   = 3 --# 刷新一天新的数据（重置签到在线奖励等）
	SUB_REFRESH_INTER_DAY = 4 --# 跨天刷新消息

	SUB_C2G_PURCHASE_QUERY_RESULT            = 1  --# 查询充值结果这个将转发给登录服务器
	SUB_G2C_PURCHASE_QUERY_RESULT            = 2  --# 返回client
	SUB_C2G_PURCHASE_TRADE_VIEW_STATUS       = 3  --# 获取支付界面的状态
	SUB_G2C_PURCHASE_TRADE_VIEW_STATUS       = 4  --# 获取支付界面的状态返回
	SUB_GR_S2C_DIAMOND_EXCHANGE              = 5  --# 钻石兑换
	SUB_GR_C2S_DIAMOND_EXCHANGE              = 6  --# 钻石兑换结果
	SUB_GR_C2G_PURCHASE_QUERY_UNDEALED_TRADE = 7  --# 查询漏单
	SUB_GR_G2C_USER_PURCHASE_BONUS_INFO      = 8  --# 用户加赠信息(暂时再用登录服那个消息)
	SUB_GR_C2G_USER_PURCHASE_BONUS_INFO      = 9  --# 请求用户加赠信息游戏服用这个
	SUB_GR_G2C_PRODUCT_INFO                  = 10 --# 返回产品信息兑换比例
	SUB_GR_C2G_PRODUCT_INFO                  = 11
	SUB_GR_C2G_FIRST_PAY_INFO                = 12 --# 请求首冲信息
	SUB_GR_G2C_FIRST_PAY_INFO                = 13 --# 请求首冲信息

	SUB_C2G_GET_PHONE_CONFIRMCODE  = 0 --# 获取手机验证码
	SUB_G2C_GET_PHONE_CONFIRMCODE  = 1 --# 获取回复
	SUB_C2G_BIND_PHONE             = 2 --# 绑定手机
	SUB_G2C_BIND_PHONE             = 3 --# 绑定回复
	SUB_C2G_GUEST_REGISTER_ACCOUNT = 4 --# 绑定帐号
	SUB_G2C_GUEST_REGISTER_ACCOUNT = 5 --# 绑定帐号回复

	--#  心跳
	SUB_GR_C_HEARTBEAT = 0
	SUB_GR_S_HEARTBEAT = 1

	--# Vip 相关
	SUB_C2G_REFRESH_VIP          = 0 --# 刷新VIP
	SUB_G2C_REFRESH_VIP          = 1 --# 刷新VIP
	SUB_C2G_VIP_USER_INFO_UPLOAD = 2 --# 上传vip用户信息
	SUB_G2C_VIP_USER_INFO_UPLOAD = 3 --# 上传vip用户信息结果

	--# 奖励
	SUB_C2G_REWARD_GET = 0
	SUB_G2C_REWARD_GET = 1

	--# 比赛消息
	SUB_GR_S_MATCH_FEE      = 0 --# 报名费用
	SUB_GR_C_MATCH_FEE      = 1 --# 报名费用
	SUB_GR_S_MATCH_FEE_FAIL = 2 --# 报名失败
	SUB_GR_C_LEAVE_MATCH    = 3 --# 退出比赛
	SUB_GR_S_MATCH_INFO     = 4 --# 比赛信息
	SUB_GR_S_MATCH_WAIT_TIP = 5 --# 等待提示
	SUB_GR_S_MATCH_RESULT   = 6 --# 比赛结果
	SUB_GR_S_MATCH_STATUS   = 7 --# 比赛状态
	SUB_GR_S_MATCH_RANK     = 8 --# 比赛排行

	SUB_GR_C_MISSION_UPDATE      = 0
	SUB_GR_S_MISSION_UPDATE      = 1
	SUB_GR_C_LOAD_MISSION        = 2 --# 加载任务
	SUB_GR_S_LOAD_MISSION        = 3 --# 加载任务
	SUB_GR_C_GET_MISSION_REWARD  = 4 --# 领取任务奖励
	SUB_GR_S_GET_MISSION_REWARD  = 5 --# 领取任务奖励
	SUB_GR_C_GET_LIVENESS_REWARD = 6 --# 获取活跃度奖励
	SUB_GR_S_GET_LIVENESS_REWARD = 7 --# 获取活跃度奖励

	--# 称号相关
	SUB_GR_C_LOAD_ACHIEVEMENT_TITLE  = 0 --# 加载角色称号
	SUB_GR_S_ACHIEVEMENT_TITLE       = 1 --# 返回角色称号
	SUB_GR_C_PUTON_ACHIEVEMENT_TITLE = 2 --# 佩戴称号C->S
	SUB_GR_S_PUTON_ACHIEVEMENT_TITLE = 3 --# 佩戴称号S->C

	--# 公会相关
	SUB_GR_C_GUILD_CREATE               = 0
	SUB_GR_S_GUILD_CREATE               = 1
	SUB_GR_C_GUILD_DISBAND              = 2
	SUB_GR_S_GUILD_DISBAND              = 3
	SUB_GR_C_GUILD_APPROVE_USER         = 4
	SUB_GR_S_GUILD_APPROVE_USER         = 5
	SUB_GR_C_GUILD_REJECT_USER          = 6
	SUB_GR_S_GUILD_REJECT_USER          = 7
	SUB_GR_C_GUILD_KICK_USER            = 8
	SUB_GR_S_GUILD_KICK_USER            = 9
	SUB_GR_C_GUILD_LOAD_ONLINE_USER     = 10
	SUB_GR_S_GUILD_LOAD_ONLINE_USER     = 11
	SUB_GR_C_GUILD_LOAD_RECOMMEND       = 12
	SUB_GR_S_GUILD_LOAD_RECOMMEND       = 13
	SUB_GR_C_GUILD_LOAD_USER_GUILD_INFO = 14
	SUB_GR_S_GUILD_LOAD_USER_GUILD_INFO = 15
	SUB_GR_C_GUILD_LOAD_USER_LIST       = 16
	SUB_GR_S_GUILD_LOAD_USER_LIST       = 17
	SUB_GR_C_GUILD_MODIFY_NOTICE        = 18
	SUB_GR_S_GUILD_MODIFY_NOTICE        = 19
	SUB_GR_C_GUILD_MODIFY_RULE          = 20
	SUB_GR_S_GUILD_MODIFY_RULE          = 21
	SUB_GR_C_GUILD_REQ_JOIN             = 22
	SUB_GR_S_GUILD_REQ_JOIN             = 23
	SUB_GR_C_GUILD_REQ_QUIT             = 24
	SUB_GR_S_GUILD_REQ_QUIT             = 25
	SUB_GR_C_GUILD_SEARCH               = 26
	SUB_GR_S_GUILD_SEARCH               = 27
	SUB_GR_C_GUILD_CHAT                 = 28
	SUB_GR_S_GUILD_CHAT                 = 29
	SUB_GR_C_GUILD_LOAD_REQ_LIST        = 30
	SUB_GR_S_GUILD_LOAD_REQ_LIST        = 31
	SUB_GR_C_GUILD_LOAD_REQ_STATUS      = 32
	SUB_GR_S_GUILD_LOAD_REQ_STATUS      = 33
	SUB_GR_S_GUILD_BC_USER_CHANGE       = 34 --# 公会成员变动
	SUB_GR_S_GUILD_BC_DISBAND           = 35 --# 公会解散
	SUB_GR_S_GUILD_BC_NEW_NOTICE        = 36 --# 公会新公告
	SUB_GR_S_GUILD_BC_CHAT              = 37 --# 公会聊天群发

	--# 游戏状态
	GAME_STATUS_FREE = 0   --# 空闲状态
	GAME_STATUS_PLAY = 100 --# 游戏状态
	GAME_STATUS_WAIT = 200 --# 等待状态

	--# 比赛状态
	MS_NULL        = 0 --# 没有状态
	MS_SIGNUP      = 1 --# 报名状态
	MS_MATCH_READY = 2 --# 准备开始
	MS_MATCHING    = 3 --# 比赛状态
	MS_OUT         = 4 --# 淘汰状态

	--# 邮件类型
	AMT_SYSTEM              = 0  --# 系统
	AMT_SYSTEM_COMPENSATION = 1  --# 系统补偿
	AMT_RANK_REWARD         = 2  --# 排行榜奖励
	AMT_MOBILE_BIND         = 3  --# 手机绑定
	AMT_FIRST_PAY           = 4  --# 首充赠送
	AMT_BUG                 = 5  --# 你提我改
	AMT_TOTAL_PAY           = 6  --# 累计充值
	AMT_GIFT                = 7  --# 礼包赠送
	AMT_EXCHANGE            = 8  --# 兑换返还
	AMT_RAIDERS             = 9  --# 夺宝奇兵
	AMT_PLAYER_GIVE         = 10 --# 玩家赠送
	AMT_MATCH               = 11 --# 竞技场8人和40人
	AMT_GM                  = 12 --# 后台赠送
	AMT_ARENA               = 13 --# 大师赛
	AMT_SIGN                = 14 --# 微信签到
	AMT_ARENA_ACTIVITY      = 15 --# 活动赛
	AMT_ARENA_BIG_PRIZE     = 16 --# 大奖赛

	GUUST_CHAT = 0 --# 聊天

	GF_ROOM_CHAT_TXT   = 0 --# 聊天
	GF_ROOM_CHAT_VOICE = 1

	--# 请求错误
	REQUEST_FAILURE_NORMAL           = 0 --# 常规原因
	REQUEST_FAILURE_NOGOLD           = 2 --# 金币不足
	REQUEST_FAILURE_NOSCORE          = 3 --# 积分不足
	REQUEST_FAILURE_PASSWORD         = 4 --# 密码错误
	REQUEST_FAILURE_NOTFIND_CHAIR    = 5 --# 没有找到座位
	REQUEST_FAILURE_NOTFIND_TABLE    = 6 --# 没有找到桌子
	REQUEST_FAILURE_NO_PLAY_COUNT    = 7 --# 参赛次数已用完
	REQUEST_FAILURE_NO_MASTER_USERID = 8 --# 没有创建者

	--# 游戏类子协议
	SUB_S_ENTER_SCENE            = 100 --# 进入场景
	SUB_S_OTHER_ENTER_SCENE      = 101 --# 其他进入场景
	SUB_S_SCENE_FISH             = 102 --# 场景鱼
	SUB_S_SCENE_BULLET           = 103 --# 场景子弹
	SUB_S_SWITCH_SCENE           = 104 --# 转换场景
	SUB_S_DISTRIBUTE_FISH        = 105 --# 服务器生成鱼
	SUB_C_USER_FIRE              = 106 --# 开火
	SUB_S_USER_FIRE              = 107 --# 开火
	SUB_C_USE_SKILL              = 108 --# 使用技能
	SUB_S_USE_SKILL              = 109 --# 服务器返回使用技能
	SUB_C_SKILL_FINISH           = 110 --# 技能结束
	SUB_C_BUY_SKILL              = 111 --# 购买技能
	SUB_S_BUY_SKILL              = 112 --# 购买技能
	SUB_C_CATCH_FISH             = 113 --# 捕获鱼儿
	SUB_S_CATCH_FISH             = 114 --# 服务器返回捕获鱼儿
	SUB_S_EXIT_SCENE             = 115 --# 离开场景
	SUB_S_UPDATE_USER_SCORE      = 116 --# 更新用户分数
	SUB_S_START_ALMS             = 117 --# 开始领取救济金
	SUB_S_STOP_ALMS              = 118 --# 停止领取救济金
	SUB_C_GET_ALMS               = 119 --# 获取救济金
	SUB_S_GET_ALMS               = 120 --# 获取救济金
	SUB_S_BOSS_COME              = 121 --# BOSS来了消息
	SUB_C_CHANGE_FIRE            = 122 --# 玩家换炮的外形
	SUB_S_CHANGE_FIRE            = 123 --# 玩家换炮的外形
	SUB_S_TRRIGER_EGG            = 124 --# 触发砸蛋
	SUB_S_TRRIGER_NUM            = 125 --# 触发猜大小
	SUB_S_TRRIGER_TREE           = 126 --# 触发摇钱树
	SUB_C_SMASH_EGG              = 127 --# 客户端请求 砸蛋
	SUB_S_SMASH_EGG              = 128 --# 返回结果 砸蛋
	SUB_C_NUM_VAL                = 129 --# 客户端请求 猜大小
	SUB_S_NUM_VAL                = 130 --# 返回结果 猜大小
	SUB_C_SMASH_TREE             = 131 --# 客户端请求 摇钱树
	SUB_S_SMASH_TREE             = 132 --# 返回结果 摇钱树
	SUB_C_VIP_GET_ALMS           = 133 --# VIP获取救济金
	SUB_C_GET_RED_ENVELOPE       = 134 --# 抢红包
	SUB_S_GET_RED_ENVELOPE       = 135 --# 抢红包结果
	SUB_S_SUMMON_FISH_INFO       = 136 --# 同步召唤鱼聚宝盆的信息
	SUB_S_FISH_CHANGE_PATH       = 137 --# 鱼改变路径
	SUB_S_READY_REWARDPOOL       = 138 --# 准备奖池
	SUB_S_OPEN_REWARDPOOL        = 139 --# 打开奖池任务
	SUB_S_CLOSE_REWARDPOOL       = 140 --# 关闭奖池任务
	SUB_S_UPDATE_REWARDPOOL_RANK = 141 --# 更新奖池排行
	SUB_S_UPDATE_REWARDPOOL_TASK = 142 --# 更新奖池任务
	SUB_C_UNLOCK_CANNON_MULTIPLE = 143 --# 解锁炮倍数
	SUB_S_UNLOCK_CANNON_MULTIPLE = 144 --# 解锁炮倍数
	SUB_S_USER_UP_LEVEL          = 145 --# 用户升级
	SUB_S_USER_ADD_EXP           = 146 --# 用户增加经验
	SUB_C_CONFIRM_NEWGUIDE_STATE = 147 --# 更新新手引导记录
	SUB_S_GAME_START             = 148 --# 游戏开始
	SUB_S_GAME_RESULT            = 149 --# 游戏结果
	SUB_S_GAME_END               = 150 --# 游戏结束
	SUB_S_GAME_RANK              = 151 --# 游戏排行
	--# 神灯召唤鱼相关
	SUB_S_SHENDENG_FISH              = 152 --# 神灯召唤鱼服务器包
	SUB_C_SHENDENG_FISH              = 153 --# 神灯召唤鱼客户端包
	SUB_C_SKILL_CATCH_FISH           = 154 --# 技能捕获鱼儿
	SUB_S_UPLEVEL_REWARD             = 155 --# 升级奖励
	SUB_C_CANNON_ANGLE               = 156 --# 炮的朝向
	SUB_S_CANNON_ANGLE               = 157 --# 炮的朝向
	SUB_S_ALL_CANNON_LEVEL           = 158 --# 广播炮的改变倍数
	SUB_C_BACK_GAME                  = 159 --# 返回前台
	SUB_S_MATCH_UP_CANNON            = 160 --# 升级炮
	SUB_C_ENTER_BACKSTAGE            = 161 --# 进入后台
	SUB_S_FISH_FORM_END              = 162 --# 鱼潮结束
	SUB_C_START_EGG                  = 163 --# 开始砸蛋
	SUB_S_START_EGG                  = 164 --# 开始砸蛋
	SUB_C_START_TREE                 = 165 --# 开始摇钱树
	SUB_S_START_TREE                 = 166 --# 开始摇钱树
	SUB_C_REFRESH_ALMS               = 167 --# 刷新救济金
	SUB_S_TREASURE_INFO              = 168 --# 同步抽奖数据
	SUB_C_TREASURE_START             = 169 --# 开始抽奖
	SUB_S_TREASURE_RESULT            = 170 --# 抽奖结果
	SUB_C_START_SKILL                = 171 --# 准备使用技能
	SUB_S_START_SKILL                = 172 --# 服务器返回准备使用技能
	SUB_S_START_MYSTERIOUS_TASK      = 173 --# 开始神秘任务
	SUB_S_CLOSE_MYSTERIOUS_TASK      = 174 --# 结束神秘任务
	SUB_S_UPDATE_MYSTERIOUS_TASK     = 175 --# 更新神秘任务
	SUB_S_ARENA_GAME_END             = 176 --# 竞技场结束
	SUB_S_ARENA_CLOSE                = 177 --# 竞技场关闭
	SUB_S_FISH_STATE_INFO            = 178 --# 鱼的状态信息
	SUB_S_PIRANHA_ATTACK_INFO        = 179 --# 食人鱼的攻击信息
	SUB_S_SCENE_STATE                = 180 --# 场景状态信息
	SUB_C_CANCEL_SKILL               = 181 --# 取消技能
	SUB_S_CANCEL_SKILL               = 182 --# 服务器返回取消技能
	SUB_C_REQ_REWARD_POOL_DETAIL     = 183 --# 请求查看奖池详细信息
	SUB_S_REQ_REWARD_POOL_DETAIL     = 184 --# 服务器返回
	SUB_C_LASER_CANNON_CATCH_FISH    = 185 --# 激光炮捕获鱼
	SUB_S_LASER_CANNON_CATCH_FISH    = 186 --# 激光炮捕获鱼
	SUB_S_CHAGE_ACHIEVEMENT_TITLE    = 187 --# 更换成就称号
	SUB_S_GUESS_SINDOU               = 188 --# 触发猜单双
	SUB_C_SINDOU_REQ                 = 189 --# 客户端请求 猜单双
	SUB_S_SINDOU_REQ                 = 190 --# 返回结果 猜单双
	SUB_S_BOMB_FISH                  = 191 --# 触发炸弹鱼
	SUB_S_ARENA_ACTIVITY_END         = 192 --# 活动赛结束
	SUB_C_BUY_BULLET_SKIN            = 193 --# 购买皮肤
	SUB_S_BUY_BULLET_SKIN            = 194 --# 返回
	SUB_C_CHANGE_BULLET_SKIN         = 195 --# 切换皮肤
	SUB_S_CHANGE_BULLET_SKIN         = 196
	SUB_S_BC_USER_CHANGE_BULLET_SKIN = 197 --# 切换皮肤广播
	SUB_S_USER_BULLET_SKIN_INFO      = 198 --# 用户皮肤信息
	SUB_C_RESIGNUP_ARENA_ACTIVITY    = 199 --# 重新报名
	SUB_S_RESIGNUP_ARENA_ACTIVITY    = 200 --# 重新报名

	GMCMD_CMD      = 1
	GMCMD_CMD_RESP = 2

