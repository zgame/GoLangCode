SandRockResourcePickNet = {}


-- 同步采集资源列表
function SandRockResourcePickNet.SendPickList(userId, allPlayer, all, list)
    local sendCmd = ProtoGameSandRock.ResourceUpdate()
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end

    --printTable(list)
    if all ~= nil then          -- 发送登录全同步
        for areaName, pointList in pairs(all) do
            for pointIndex, point in pairs(pointList) do
                local points = sendCmd.points:add()
                points.areaName = areaName
                points.areaPoint = pointIndex
                points.resourceType = point.resourceType
            end
        end
    end

    if list ~= nil then             -- 发送增量同步
            for _, point in pairs(list) do
                local points = sendCmd.points:add()
                points.areaName = point.areaName
                points.areaPoint = point.areaPoint
                points.resourceType = point.resourceType
            end
    end
    sendCmd.weather = SandRockRoom.GetWeather(room)
    --print("发送 采集资源列表")
    --print(sendCmd)
    if allPlayer == nil then
        NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_POINT, sendCmd, nil)
    else
        SandRockRoom.SendMsgToAllUsers(room,CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_POINT, sendCmd)
    end




end

-- 采集资源
function SandRockResourcePickNet.GetPickResource(serverId, userId, buf)
    --print("客户端开始采集资源")
    local msg = ProtoGameSandRock.ResourceGet()
    msg:ParseFromString(buf)
    --print(msg)

    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end
    local player = GameServer.GetPlayerByUID(userId)
    if player == nil then
        return
    end

    local areaName = msg.info.areaName
    local areaPoint = msg.info.areaPoint
    local resourceType = msg.info.resourceType

    local itemList = SandRockRoom.GetPickResource(room, userId, areaName, areaPoint, resourceType)
    if itemList == nil then
        ZLog.Logger("资源采集失败")
        return
    end
    local sendCmd = SandRockSleepNet.SendItemList(player, itemList)
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_GET, sendCmd, nil)

    -- 给其他玩家同步 采集资源点的刷新
    local sendCmd2 = ProtoGameSandRock.ResourceUpdate()
    local points = sendCmd2.points:add()
    points.areaName = areaName
    points.areaPoint = areaPoint
    points.resourceType = 0  -- 清理掉

    SandRockRoom.SendMsgToAllUsers(room,CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_POINT, sendCmd2)


end

