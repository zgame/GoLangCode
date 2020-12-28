Location={}

--上报位置
function Location.SendLocation(serverId)
    --print("上传位置信息")
    local player = GameClient.GetPlayer(serverId)
    local uId = player.User.UserId
    local sendCmd = ProtoGameSandRock.PlayerLocation()
    local uu = sendCmd.location:add()
    uu.userId = uId
	uu.x = 10
	uu.y = 672
	uu.z = -20
	uu.faceDir = 5
    uu.action = ZRandom.GetRandom(1,10)
    sendCmd.time = ZTime.GetOsTimeMillisecond()
    --print(sendCmd)
    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_LOCATION,sendCmd,nil)
    
end
-- 其他玩家的同步
function Location.OtherLocation(serverId,userId,buf)
    print("所有玩家同步位置")
    local msg = ProtoGameSandRock.PlayerLocation()
    msg:ParseFromString(buf)
    print(msg)
end