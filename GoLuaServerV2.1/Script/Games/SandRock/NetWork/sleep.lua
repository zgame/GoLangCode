SandRockSleepNet ={}

function SandRockSleepNet.Sleep(serverId, userId, buf)
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end
    -- refresh
    ZLog.printTime(1)
    local pickList = SandRockRoom.ResourcePickPointUpdate(room)      -- 刷新采集类资源点列表   3
    ZLog.printTime(2)
    local reliveList = SandRockRoom.ResourceTerrainUpdate(room)      -- 刷新地形树和石头资源列表，获取的是刷新的树  1
    ZLog.printTime(3)
    SandRockRoom.UpdateWeather(room)
    -- send
    SandRockResourcePickNet.SendPickList(userId, true, nil,  pickList)       -- 发送资源刷新  3
    ZLog.printTime(5)
    SandRockResourceTerrainNet.SendTreeRelive(userId, reliveList, true)      -- 发送地形树的重生

end

function SandRockSleepNet.SendItemList(player, itemList)
    local sendCmd
    if itemList ~= nil then
        sendCmd = ProtoGameSandRock.ItemGet()
        for itemId, num in pairs(itemList) do
            local item = sendCmd.item:add()
            item.itemId = tonumber(itemId)
            item.itemNum = num
        end
        sendCmd.exp = Player.ExpGet(player)
        sendCmd.level = Player.LevelGet(player)
        sendCmd.sp = Player.SpGet(player)
        --print(sendCmd)
        --print("发送客户端获得道具结果")
    end
    return sendCmd
end