SandRockSleep ={}

function SandRockSleep.Sleep(serverId, userId, buf)
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end
    SandRockRoom.ResourcePointUpdate(room)
    SandRockResourcePoint.SendPointList(userId)       -- 发送资源刷新

end