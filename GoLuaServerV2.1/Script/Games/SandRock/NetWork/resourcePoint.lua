SandRockResourcePoint = {}


-- 同步资源列表
function SandRockResourcePoint.SendPointList(userId)
    local sendCmd = ProtoGameSandRock.ResourceUpdate()
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end

    --printTable(room.resourcePoint)
    for areaName, pointList in pairs(room.resourcePoint) do
            for pointIndex, point in pairs(pointList) do
                local points = sendCmd.points:add()
                points.areaName = areaName
                points.areaPoint = pointIndex
                points.resourceType = point.resourceType
            end
    end
    --print(sendCmd)
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_POINT, sendCmd, nil)


end

-- 采集资源
function SandRockResourcePoint.GetResource(serverId, userId, buf)
    local msg = ProtoGameSandRock.PlayerLocation()
    msg:ParseFromString(buf)
    print(msg)

    local room = GameServer.GetRoomByUserId(userId)
    if room ==nil then
        return
    end

    SandRockRoom.GetResource(userId, areaName, pointIndex, resourceType)
end
