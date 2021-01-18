SandRockSleepNet ={}

function SandRockSleepNet.Sleep(serverId, userId, buf)
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end
    -- refresh
    SandRockRoom.ResourcePointUpdate(room)      -- 刷新资源列表
    local reliveList = SandRockRoom.ResourceTerrainUpdate(room)      -- 刷新地形树和石头资源列表
    SandRockRoom.UpdateWeather(room)

    -- send
    SandRockResourcePickNet.SendPickList(userId)       -- 发送资源刷新
    SandRockResourceTerrainNet.SendTreeRelive(userId, reliveList)      -- 发送地形树的重生

end