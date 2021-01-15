SandRockResourceTerrain = {}


-- 同步资源列表
function SandRockResourceTerrain.SendPointList(userId)
    local sendCmd = ProtoGameSandRock.ResourceTerrainUpdate()
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end

    --printTable(room.resourcePoint)
    for areaName, pointList in pairs(room.resourceTerrain) do
        for pointIndex, point in pairs(pointList) do
            local points = sendCmd.points:add()
            points.areaName = areaName
            points.areaPoint = pointIndex
            points.resourceType = point.resourceType
        end
    end
    sendCmd.weather = SandRockRoom.GetWeather(room)
    --print(sendCmd)
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN, sendCmd, nil)
end

-- 采集资源
function SandRockResourceTerrain.GetTerrainResource(serverId, userId, buf)
    --print("客户端开始开采资源")

end