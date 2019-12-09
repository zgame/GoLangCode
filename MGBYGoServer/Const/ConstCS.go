
package Const

const (
	//#  消息类别
	MAIN_CHAT_CMD = 1003

	SUB_C_LOGIN                   = 0  //# 登录
	SUB_S_LOGIN                   = 1  //# 登录返回
	SUB_C_FRIEND_LIST             = 2  //# 获取好友列表
	SUB_S_FRIEND_LIST             = 3  //# 获取好友列表
	SUB_C_SEARCH_USER             = 4  //# 查找玩家
	SUB_S_SEARCH_USER             = 5  //# 查找玩家
	SUB_C_APPLY_FRIEND            = 6  //# 申请好友
	SUB_S_APPLY_FRIEND_RESULT     = 7  //# 申请好友结果
	SUB_C_DELETE_FRIEND           = 8  //# 删除好友
	SUB_S_DELETE_FRIEND           = 9  //# 删除好友
	SUB_C_GET_CHAT_MESSAGE        = 10 //# 获取未读聊天信息
	SUB_S_GET_CHAT_MESSAGE        = 11 //# 获取未读聊天信息
	SUB_C_CHAT_MESSAGE            = 12 //# 聊天信息
	SUB_S_CHAT_MESSAGE            = 13 //# 聊天信息
	SUB_S_NOTIFY_NEW_CHAT_MESSAGE = 14 //# 新的聊天信息
	SUB_C_HEARTBEAT_CHAT          = 15 //# 心跳包
	SUB_S_HEARTBEAT_CHAT          = 16 //# 心跳包
)
