SandRockSleep ={}

function SandRockSleep.Sleep(serverId, userId, buf)
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end
    SandRockRoom.ResourcePointUpdate(room)      -- 刷新资源列表
    SandRockRoom.UpdateWeather(room)
    SandRockResourcePick.SendPickList(userId)       -- 发送资源刷新

end