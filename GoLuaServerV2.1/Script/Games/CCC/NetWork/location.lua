CCCNetworkLocation = {}

-- 上报位置
function CCCNetworkLocation.Location(serverId, userId, buf)
    local msg = ProtoGameCCC.PlayerLocation()
    msg:ParseFromString(buf)
    --print("上报位置")
    local room = GameServer.GetRoomByUserId(userId)
    CCCRoom.SetPlayerLocation(room,userId, msg.location[1])

end

--
function CCCNetworkLocation.Copy(source,dec)
    dec.userId = source.userId
    dec.x = source.x
    dec.y = source.y
    dec.z = source.z
    dec.faceDir = source.faceDir
    dec.action = source.action
    return dec
end