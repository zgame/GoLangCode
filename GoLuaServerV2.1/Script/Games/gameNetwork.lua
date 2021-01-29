-----------------------------------------------------------------
--- 游戏服务器的网络分发, 主要功能是根据游戏不同， 跳转到不同的游戏处理
-----------------------------------------------------------------
GameNetwork = {}

function GameNetwork.Init(serverId)
    print("创建了链接, serverId:"..tostring(serverId))
end


-- 根据命令进行分支处理
function GameNetwork.Receive(serverId, userId, mainSgId, subMsgId, data, token)
    --print("serverId",serverId, "userId",userId,"msgId",mainSgId, "subMsgId",subMsgId)
    --print("msgId",mainSgId, "subMsgId",subMsgId)
    --UserToken = token           -- 保存到全局里面，发送的时候取出来GameMessage

    local switch = {}
    switch[CMD_MAIN.MDM_GAME_SAND_ROCK] = SandRockNetwork.Receive         -- 跳转到ccc处理

    switch[mainSgId](serverId, userId, mainSgId, subMsgId, data, token)


end

--- go通知lua 所有掉线的连接都要走这里
function GameNetwork.Broken(uId, serverId)
    --ZLog.Logger("通知：" .. uId .. "  掉线了")
    local room = GameServer.GetRoomByUserId(uId)
    if room ~= nil then
        room:PlayerStandUp(uId)
    end
end