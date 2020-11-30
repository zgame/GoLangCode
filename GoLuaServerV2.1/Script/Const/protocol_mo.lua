----# 主命令定义
 MAIN_CMD_ID = 1100

----# 子命令定义
 SUB_C_MONITOR_REG = 10   -- 注册为客户端
 SUB_S_MONITOR_ITEMS = 100   -- 下发服务器列表
 SUB_S_MONITOR_STATE = 101   -- 更新状态
 SUB_S_MONITOR_LOG = 103   -- 添加日志
 SUB_S_NEW_MONITOR_ITEM = 104   -- 新增服务器
 SUB_S_DEL_MONITOR_ITEM = 105   -- 删除服务器
 SUB_C_MONITOR_KEEPLIVE = 200   -- 心跳包
 SUB_S_MONITOR_KEEPLIVE = 201   -- 心跳包
 SUB_C_MONITOR_CMD = 2050   -- 执行命令
 SUB_S_MONITOR_CMD = 2051   -- 执行命令结果
