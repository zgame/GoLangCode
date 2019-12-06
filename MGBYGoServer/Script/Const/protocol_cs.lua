--- 聊天服主命令
MAIN_CHAT_SERVICE_INNER = 1002		-- 聊天服务
MAIN_CHAT_SERVICE_CLIENT = 1003		-- 聊天服务（客户端）

--- 聊天服与登录服通信
SUB_SC_REGISTER    = 0						-- 注册连接
SUB_SS_REGISTER    = 1						-- 注册连接
SUB_SC_PUSH_TOKEN  = 2						-- 推送token
SUB_SC_UPDATE_USER_INFO = 3				-- 更新玩家信息
SUB_SC_TOKEN_USER_INFO  = 4                  -- 推送玩家数据

--- 聊天服与客户端通信
SUB_C_LOGIN                   = 0  --# 登录
SUB_S_LOGIN                   = 1  --# 登录返回
SUB_C_FRIEND_LIST             = 2  --# 获取好友列表
SUB_S_FRIEND_LIST             = 3  --# 获取好友列表
SUB_C_SEARCH_USER             = 4  --# 查找玩家
SUB_S_SEARCH_USER             = 5  --# 查找玩家
SUB_C_APPLY_FRIEND            = 6  --# 申请好友
SUB_S_APPLY_FRIEND_RESULT     = 7  --# 申请好友结果
SUB_C_DELETE_FRIEND           = 8  --# 删除好友
SUB_S_DELETE_FRIEND           = 9  --# 删除好友
SUB_C_GET_CHAT_MESSAGE        = 10 --# 获取未读聊天信息
SUB_S_GET_CHAT_MESSAGE        = 11 --# 获取未读聊天信息
SUB_C_CHAT_MESSAGE            = 12 --# 聊天信息
SUB_S_CHAT_MESSAGE            = 13 --# 聊天信息
SUB_S_NOTIFY_NEW_CHAT_MESSAGE = 14 --# 新的聊天信息
SUB_C_HEARTBEAT_CHAT          = 15 --# 心跳包
SUB_S_HEARTBEAT_CHAT          = 16 --# 心跳包
