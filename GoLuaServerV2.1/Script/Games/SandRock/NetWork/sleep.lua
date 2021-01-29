SandRockSleepNet ={}

function SandRockSleepNet.Sleep(serverId, userId, buf)
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end
    -- refresh
    SandRockRoom.ResourcePointUpdate(room)      -- 刷新资源列表
    local reliveList = SandRockRoom.ResourceTerrainUpdate(room)      -- 刷新地形树和石头资源列表，获取的是刷新的树
    SandRockRoom.UpdateWeather(room)

    -- send
    SandRockResourcePickNet.SendSleepPickList(userId)       -- 发送资源刷新
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