Resource ={}

-- 资源采集
function Resource.Action(serverId, areaName,areaPoint,resourceType)
    local sendCmd = ProtoGameSandRock.ResourceGet()
    sendCmd.info.areaName  = areaName
    sendCmd.info.areaPoint  = areaPoint
    sendCmd.info.resourceType  = resourceType

    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_GET,sendCmd,nil)
end

function Resource.Get(serverId,userId, buf)
    print("获得资源采集道具")
    local msg = ProtoGameSandRock.ItemGet()
    msg:ParseFromString(buf)
    print(msg)
end