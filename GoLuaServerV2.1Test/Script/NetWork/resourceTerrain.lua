ResourceTerrain ={}

-- 地形资源采集
function ResourceTerrain.Action(serverId, areaName, areaPoint, resourceType ,trunkHealth,stumpHealth)
    local sendCmd = ProtoGameSandRock.ResourceTerrainGet()
    sendCmd.info.areaName  = areaName
    sendCmd.info.areaPoint  = areaPoint
    sendCmd.info.resourceType  = resourceType
    sendCmd.info.trunkHealth  = trunkHealth
    sendCmd.info.stumpHealth  = stumpHealth
    sendCmd.toolId = 11000006    -- 砍树
    sendCmd.damage = 10

    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN_GET,sendCmd,nil)
end

function ResourceTerrain.GetItem(serverId, userId, buf)
    print("获得地形资源采集道具")
    local msg = ProtoGameSandRock.ItemGet()
    msg:ParseFromString(buf)
    print(msg)
end