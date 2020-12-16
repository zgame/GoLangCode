----------------------------------------------------------------------
---网络总分发， 跳转到各个不同的服务器去处理
----------------------------------------------------------------------

---网络连接成功时候的初始化
function GoCallLuaNetWorkInit(serverId)

    local switch={}
    switch["Game"] = GameNetwork.Init                     -- 启动游戏服
    switch["Center"] = CenterServer.Init                 -- 启动主中心服
    -- 运行对应server type的函数
    switch[GlobalVar.ServerTypeName](serverId)

end


----------------------------------------------------------------------
---接收消息 tcp
----------------------------------------------------------------------
-- 网络接收函数
function GoCallLuaNetWorkReceive(serverId, userId, msgId, subMsgId, data, token)

    local switch={}
    switch["Game"] = GameNetwork.Receive                     -- 游戏服
    switch["Center"] = CenterServer.Receive                 -- 主中心服
    -- 运行对应server type的函数
    switch[GlobalVar.ServerTypeName](serverId, userId, msgId, subMsgId, data, token)


end


----------------------------------------------------------------------
---接收消息 upd
----------------------------------------------------------------------
-- 网络接收函数
function GoCallLuaNetWorkUdpReceive(serverAddr,  msgId, subMsgId, data)

    GoCallLuaNetWorkReceive(serverAddr, nil, msgId, subMsgId, data, nil)

end


----------------------------------------------------------------------
--- 网络中断
----------------------------------------------------------------------
function GoCallLuaPlayerNetworkBroken(uid, serverId)
    local switch={}
    switch["Game"] = GameNetwork.Broken                     -- 游戏服
    switch["Center"] = CenterServer.Broken                 -- 主中心服
    -- 运行对应server type的函数
    switch[GlobalVar.ServerTypeName](uid, serverId)
end