Location={}

--上报位置
function Location.SendLocation(serverId)
    print("上传位置信息")
    local player = GameClient.GetPlayer(serverId)
    local uId = player.User.UserId
    local sendCmd = ProtoGameCCC.PlayerLocation()
    local uu = sendCmd.location:add()
    uu.userId = uId
	uu.x = 2
	uu.y = 3
	uu.z = 4
	uu.faceDir = 5
    uu.action = 6
    sendCmd.time = ZTime.GetOsTimeMillisecond()
    print(sendCmd)
    Network.Send(serverId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_LOCATION,sendCmd,nil)
    
end
-- 其他玩家的同步
function Location.OtherLocation(serverId,userId,buf)
    print("所有玩家同步位置")
    local msg = ProtoGameCCC.PlayerLocation()
    msg:ParseFromString(buf)
    print(msg)
end