SandRockResourceTerrain = {}

-- 同步资源列表
function SandRockResourceTerrain.SendTreeRelive(userId,reliveList)
    local sendCmd = ProtoGameSandRock.ResourceTerrainUpdate()

    for _, element in ipairs(reliveList) do
        local points = sendCmd.points:add()
        points.areaName = element.areaName
        points.areaPoint = element.areaPoint
        points.resourceType = element.resourceType
        points.trunkHealth = element.trunkHealth
        points.stumpHealth = element.stumpHealth
    end

    --print(sendCmd)
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN, sendCmd, nil)
end

-- 采集资源
function SandRockResourceTerrain.GetTerrainResource(serverId, userId, buf)
    --print("客户端开始开采资源")
    local msg = ProtoGameSandRock.ResourceTerrainGet()
    msg:ParseFromString(buf)
    print(msg)

    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end

    local areaName = msg.info.areaName
    local areaPoint = msg.info.areaPoint
    local resourceType = msg.info.resourceType
    local toolId = msg.toolId

    local itemList,reliveList = SandRockRoom.GetTerrainResource(room, userId, areaName, areaPoint, resourceType, toolId)
    if itemList ~= nil then
        -- 获得了道具
        local sendCmd = ProtoGameSandRock.ItemGet()
        for itemId, num in pairs(itemList) do
            local item = sendCmd.item:add()
            item.itemId = itemId
            item.itemNum = num
        end
        --print(sendCmd)
        --print("发送客户端采集结果")
        NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN_GET, sendCmd, nil)
    end
    -- 发送一下这颗树的情况
    if reliveList~=nil then
        SandRockResourceTerrain.SendTreeRelive(userId,reliveList)
    end
end