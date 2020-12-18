----------------------------------------------------------------------
---网络总分发， 跳转到各个不同的服务器去处理
----------------------------------------------------------------------

ServerNetwork ={}


---网络连接成功时候的初始化
function ServerNetwork.NetWorkInit(serverId)

    local switch={}
    switch[Const.ServerGame] = GameNetwork.Init                     -- 启动游戏服
    switch[Const.ServerCenter] = CenterServer.Init                 -- 启动主中心服
    -- 运行对应server type的函数
    switch[GlobalVar.ServerTypeName](serverId)

end


----------------------------------------------------------------------
---接收消息 tcp
----------------------------------------------------------------------
-- 网络接收函数
function ServerNetwork.NetWorkReceive(serverId, userId, msgId, subMsgId, data, token)

    local switch={}
    switch[Const.ServerGame] = GameNetwork.Receive                     -- 游戏服
    switch[Const.ServerCenter] = CenterServer.Receive                 -- 主中心服
    -- 运行对应server type的函数
    switch[GlobalVar.ServerTypeName](serverId, userId, msgId, subMsgId, data, token)


end


----------------------------------------------------------------------
---接收消息 upd
----------------------------------------------------------------------
-- 网络接收函数
function ServerNetwork.NetWorkUdpReceive(serverAddr, msgId, subMsgId, data)
    ServerNetwork.NetWorkReceive(serverAddr, nil, msgId, subMsgId, data, nil)
end


----------------------------------------------------------------------
--- 网络中断
----------------------------------------------------------------------
function ServerNetwork.PlayerNetworkBroken(uid, serverId)
    local switch={}
    switch[Const.ServerGame] = GameNetwork.Broken                     -- 游戏服
    switch[Const.ServerCenter] = CenterServer.Broken                 -- 主中心服
    -- 运行对应server type的函数
    switch[GlobalVar.ServerTypeName](uid, serverId)
end