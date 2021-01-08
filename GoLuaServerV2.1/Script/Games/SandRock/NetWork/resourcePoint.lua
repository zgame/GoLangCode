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
    sendCmd.weather = SandRockRoom.GetWeather(room)
    --print(sendCmd)
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_POINT, sendCmd, nil)


end

-- 采集资源
function SandRockResourcePoint.GetResource(serverId, userId, buf)
    --print("客户端开始采集资源")
    local msg = ProtoGameSandRock.ResourceGet()
    msg:ParseFromString(buf)
    --print(msg)

    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end

    local areaName = msg.info.areaName
    local areaPoint = msg.info.areaPoint
    local resourceType = msg.info.resourceType

    local itemGet = SandRockRoom.GetResource(room, userId, areaName, areaPoint, resourceType)
    if itemGet == nil then
        ZLog.Logger("资源采集失败")
        return
    end
    local sendCmd = ProtoGameSandRock.ItemGet()
    local item = sendCmd.item:add()
    item.itemId = itemGet
    item.itemNum = 1
    --print(sendCmd)
    --print("发送客户端采集结果")
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_GET, sendCmd, nil)

end
