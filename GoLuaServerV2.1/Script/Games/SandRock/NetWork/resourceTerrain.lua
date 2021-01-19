SandRockResourceTerrainNet = {}

-- 发送道具列表
function SandRockResourceTerrainNet.SendItemList(serverId, player, itemList, nilSend)
    local sendCmd = nil
    if itemList ~= nil then
        sendCmd = ProtoGameSandRock.ItemGet()
        for itemId, num in pairs(itemList) do
            local item = sendCmd.item:add()
            item.itemId = itemId
            item.itemNum = num
        end
        sendCmd.exp = Player.ExpGet(player)
        sendCmd.level = Player.LevelGet(player)
        sendCmd.sp = Player.SpGet(player)
        --print(sendCmd)
        --print("发送客户端采集结果")
    end
    if itemList ~= nil or nilSend then
        NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN_GET, sendCmd, nil)
    end
end


-- 同步资源列表
function SandRockResourceTerrainNet.SendTreeRelive(userId,reliveList)
    local sendCmd = ProtoGameSandRock.ResourceTerrainUpdate()

    for _, element in ipairs(reliveList) do
        local points = sendCmd.points:add()
        points.areaName = element.areaName
        points.areaPoint = element.areaPoint
        points.resourceType = element.resourceType
        points.trunkHealth = element.trunkHealth
        points.stumpHealth = element.stumpHealth
    end

    print("发送树的状态")
    print(sendCmd)
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN, sendCmd, nil)
end

-- 采集资源
function SandRockResourceTerrainNet.GetTerrainResource(serverId, userId, buf)
    print("客户端开始开采资源")
    local msg = ProtoGameSandRock.ResourceTerrainGet()
    msg:ParseFromString(buf)
    print(msg)

    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end
    local player = GameServer.GetPlayerByUID(userId)
    if player == nil then
        return
    end

    local itemList,reliveList = SandRockRoom.GetTerrainResource(room, userId, msg.info.areaName, msg.info.areaPoint, msg.info.resourceType, msg.toolId, msg.damage)

    if msg.toolId == 0 then
        SandRockResourceTerrainNet.SendItemList(serverId, player, itemList, true)        -- 踢树处理
    else
        SandRockResourceTerrainNet.SendItemList(serverId, player, itemList)        -- 砍树
    end


    -- 发送一下这颗树的情况
    if reliveList~=nil then
        SandRockResourceTerrainNet.SendTreeRelive(userId,reliveList)
    end
end