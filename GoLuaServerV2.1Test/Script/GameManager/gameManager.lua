

GameClient ={
    playerList = {}     -- serverId   player
}

-----------------------------------客户端开启-------------------------------------
-- 服务器开始创建各个游戏，这里的游戏都是多人的游戏， 如果是单人游戏，玩家自己创建即可
function GameClient.Start(serverId)
    print("客户端开启",serverId)
    -- 开始连接服务器
    LoginServer.SendLogin(serverId)
end

function GameClient.AddPlayer(serverId , player)
    GameClient.playerList[serverId] = player
end

function GameClient.GetPlayer(serverId)
    return GameClient.playerList[serverId]
end