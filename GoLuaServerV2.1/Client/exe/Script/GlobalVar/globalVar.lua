---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/21 16:20
---


------------------------统计信息用-----------------------
ServerIP_Port = ""      -- 当前服务器的地址和端口号
ServerStateSendNum = 0      -- 平均网络发送包数量
ServerStateReceiveNum = 0   -- 平均网络接收包数量
ServerSendWriteChannelNum = 0  -- 用来发送数据包的缓存，如果太大就会中断网络
ServerDataHeadErrorNum = 0  -- 数据包的头尾有校验码， 这是校验码的错误数量


ZswLogShowSendMsgNum = 0        -- 发送数量
ZswLogShowSendMsgLastTime = 0   -- 发送时间
ZswLogShowReceiveMsgNum = 0     -- 接收数量
ZswLogShowReceiveLastTime = 0      -- 接收时间




------------------------游戏和玩家列表-----------------------
AllGamesList = {}   -- gameType , game
AllPlayerList = {}   -- 所有玩家列表   key  userId , value player

ALLUserUUID = 0   -- 玩家uid的自增

