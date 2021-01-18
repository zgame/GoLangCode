

GameSandRock ={}

function GameSandRock.RunGame(serverId)
    --游戏循环
    --print("Run game")
    --Location.SendLocation(serverId)


    --Sleep.Sleep(serverId)     -- 睡觉


    -- 采集地形资源
    ResourceTerrain.Action(serverId,'Terrain_2_2',
            0,4,
            0 , 0 )


end