
-----------------------------------------------------------------------------
--- 这里保存了全局变量， 切记本文件不能reload , 否则热更新之后， 数据都重置了
-----------------------------------------------------------------------------

GlobalVar = {
    ------------------------统计信息用-----------------------
    ServerTypeName = "", -- 定义了当前服务器的类型名字   例子 ： Game  Center
    ServerIP_Port = "", -- 当前服务器的地址和端口号
    GameRoomServerID = 0, -- 游戏房间的ServerID(用于夹在游戏房间相关信息)


    ------------------------数据库句柄-----------------------
    RedisConnect = nil, --redis 数据库句柄
    MongoMainConnect = nil, -- mongo db 主库
    MySqlMainConnect = nil, -- 主mysql数据库
    SqlServerMainEngineConnect = nil, -- sql server主数据库

    ------------------------各个服务器的serverId-----------------------
    ServerMainServer = 0, -- 主中心服务器


    ------------------------游戏和玩家列表-----------------------

    AllGamesList = {}  , -- gameType (type string) , game          --- 要注意， key 不能用数字，因为占用内存太大， goperlua的问题          AllGamesList[tostring(gameType)]
    AllPlayerList = {}   ,-- 所有玩家列表   key  userId(type string) , value player     --- 要注意， key 不能用数字，因为占用内存太大， goperlua的问题
    AllPlayerListNumber = 0 ,  -- 所有玩家的人数，记住不要遍历上面的map，性能太慢
}


--ServerStateSendNum = 0      -- 平均网络发送包数量
--ServerStateReceiveNum = 0   -- 平均网络接收包数量
--ServerSendWriteChannelNum = 0  -- 用来发送数据包的缓存，如果太大就会中断网络
--ServerDataHeadErrorNum = 0  -- 数据包的头尾有校验码， 这是校验码的错误数量
--ServerHeapInUse = 0       -- 程序员申请的堆内存
--ServerNetWorkDelay = 0    -- 平均网络延迟时间


--ZswLogShowSendMsgNum = 0        -- 发送数量
--ZswLogShowSendMsgLastTime = 0   -- 发送时间
--ZswLogShowReceiveMsgNum = 0     -- 接收数量
--ZswLogShowReceiveLastTime = 0      -- 接收时间



--AllGamesListRunCurrentTableIndex = {}                     -- 所有游戏中的房间已经run过的map ，  run过之后，就会加入到这里，每次循环的时候判断是否run过了



--ALLUserUUID = 0   -- 玩家uid的自增

-----------------------临时全局变量---------------------------------------
UserToken = 0


