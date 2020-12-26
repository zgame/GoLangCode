
-----------------------------------------------------------------
--- 游戏服务器的网络分发, 主要功能是根据游戏不同， 跳转到不同的游戏处理
-----------------------------------------------------------------
GameNetwork ={}

function GameNetwork.Init(serverId)

end


-- 根据命令进行分支处理
function GameNetwork.Receive(serverId, userId, mainSgId, subMsgId, data, token)
    --print("msgId",msgId, "subMsgId",subMsgId)
    --UserToken = token           -- 保存到全局里面，发送的时候取出来GameMessage

    local switch={}
    switch[CMD_MAIN.MDM_GAME_CCC] = CCCNetwork.Receive         -- 跳转到ccc处理

    switch[mainSgId](serverId, userId, mainSgId, subMsgId, data, token)



end



--- go通知lua 所有掉线的连接都要走这里
function GameNetwork.Broken(uId, serverId)
    ZLog.Logger("通知：" .. uId .. "  掉线了")
    local player = GameServer.GetPlayerByUID(uId)
    if player ~= nil then
        local game = GameServer.GetGameByID(player.gameId)
        if game ~= nil then
            Game.PlayerLogOutGame(game,player)
            --player.NetWorkState = false
            --player.NetWorkCloseTimer = GetOsTimeMillisecond()
        else
            ZLog.Logger("游戏为空"..player.gameId)
        end
    else
        ZLog.Logger("玩家为空".. uId)
    end
end