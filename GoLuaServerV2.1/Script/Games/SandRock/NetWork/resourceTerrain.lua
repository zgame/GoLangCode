SandRockResourceTerrainNet = {}

-- 发送道具列表
function SandRockResourceTerrainNet.SendItemList(serverId, player, itemList, nilSend)
    local sendCmd = SandRockSleepNet.SendItemList(player, itemList)
    if itemList ~= nil or nilSend then
        --print("发送踢树发道具".. serverId .. " uid  ".. Player.UId(player))
        --NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN_GET, sendCmd, nil)
        NetWork.SendToUser(Player.UId(player),CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN_GET, sendCmd,nil,nil)
    end
end


-- 同步地形资源列表
function SandRockResourceTerrainNet.SendTreeRelive(userId, reliveList,allPlayer)
    local sendCmd = ProtoGameSandRock.ResourceTerrainUpdate()

    for _, element in ipairs(reliveList) do
        local points = sendCmd.points:add()
        points.areaName = element.areaName
        points.areaPoint = element.areaPoint
        points.resourceType = element.resourceType
        points.trunkHealth = element.trunkHealth
        points.stumpHealth = element.stumpHealth
    end

    --print("发送地形树的状态"..os.time())
    --print(sendCmd)

    if allPlayer == nil then
        NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN, sendCmd, nil)
    else
        -- 给所有玩家同步
        local room = GameServer.GetRoomByUserId(userId)
        SandRockRoom.SendMsgToAllUsers(room, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN, sendCmd, userId)
    end
end

-- 采集资源
function SandRockResourceTerrainNet.GetTerrainResource(serverId, userId, buf)
    --print("资源砍树开始"..ZTime.GetOsTimeMillisecond())
    local msg = ProtoGameSandRock.ResourceTerrainGet()
    msg:ParseFromString(buf)


    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end
    local player = GameServer.GetPlayerByUID(userId)
    if player == nil then
        return
    end

    if msg.toolId == 0 then
        --print("客户端开始 踢树".. userId)
        --print(msg)
        -- 踢树处理
        local itemList = SandRockRoom.GetTerrainKick(room, userId, msg.info.areaName, msg.info.areaPoint, msg.info.resourceType)
        SandRockResourceTerrainNet.SendItemList(serverId, player, itemList, true)        -- 踢树处理
    else
        --print("客户端开始 砍树".. userId)
        --print(msg)
        -- 砍树，砍石头处理
        local itemList, reliveList = SandRockRoom.GetTerrainResource(room, userId, msg.info.areaName, msg.info.areaPoint, msg.info.resourceType, msg.toolId, msg.damage)
        SandRockResourceTerrainNet.SendItemList(serverId, player, itemList)        -- 砍树
        -- 发送一下这颗树的情况
        if reliveList ~= nil then
            --print("同步其他玩家")
            SandRockResourceTerrainNet.SendTreeRelive(userId, reliveList, true)
        end
    end

    --print("资源砍树结束"..ZTime.GetOsTimeMillisecond())
end