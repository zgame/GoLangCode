
-----------------------------------------------------------------
--- 游戏服务器的网络分发, 主要功能是根据游戏不同， 跳转到不同的游戏处理
-----------------------------------------------------------------
GameNetwork ={}

function GameNetwork.Init(serverId)

end


-- 根据命令进行分支处理
function GameNetwork.Receive(serverId, userId, mainSgId, subMsgId, data, token)
    --print("msgId",msgId, "subMsgId",subMsgId)
    UserToken = token           -- 保存到全局里面，发送的时候取出来GameMessage

    if mainSgId == CMD_MAIN.MDM_GAME_CCC then
        -- 跳转到ccc处理
        CCCNetWork.Receive(serverId, userId, mainSgId, subMsgId, data, token)
    end

end

-- 根据命令进行分支处理
function GameNetwork.ReceiveUdp(serverAddr,  mainSgId, subMsgId, data)
    --print("msgId",msgId, "subMsgId",subMsgId)
    --UserToken = token           -- 保存到全局里面，发送的时候取出来GameMessage

    if mainSgId == CMD_MAIN.MDM_GAME_CCC then
        -- 跳转到ccc处理
        CCCNetWork.Receive(serverId, userId, mainSgId, subMsgId, data, token)
    end

end



--- go通知lua 所有掉线的连接都要走这里
function GameNetwork.Broken(uid, serverId)
    ZLog.Logger("通知：" .. uid .. "  掉线了")

    local player = GameServer.GetPlayerByUID(uid)
    if player ~= nil then
        --printTable(player,0,"LeavePlayer")
        --print("LeavePlayer.UID="..player.User.UserId)
        local game = GameServer.GetGameByID(player.GameType)
        --printTable(game)
        if game ~= nil then
            Game.PlayerLogOutGame(game,player)
            --player.NetWorkState = false
            --player.NetWorkCloseTimer = GetOsTimeMillisecond()
        end
    end
end