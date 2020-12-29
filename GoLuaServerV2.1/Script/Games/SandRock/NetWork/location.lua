SandRockLocation = {}

-- 上报位置
function SandRockLocation.Location(serverId, userId, buf)
    local msg = ProtoGameSandRock.PlayerLocation()
    msg:ParseFromString(buf)
    --print("上报位置---------------------------------------")
    --print(msg)
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        ZLog.Logger("没有获取到房间".. tostring(userId))
        return
    end
    local location = {}
    location = SandRockLocation.Copy(msg.location[1],location)
    SandRockRoom.SetPlayerLocation(room, userId, location)

end

--
function SandRockLocation.Copy(source, dec)
    dec.userId = source.userId
    dec.x = source.x
    dec.y = source.y
    dec.z = source.z
    dec.faceDir = source.faceDir
    dec.action = source.action
    dec.param = source.param
    return dec
end