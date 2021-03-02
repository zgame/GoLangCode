SandRockLocationNet = {}

-- 上报位置
function SandRockLocationNet.Location(serverId, userId, buf)
    local msg = ProtoGameSandRock.PlayerLocation()
    msg:ParseFromString(buf)
    --print("上报位置和动作---------------------------------------"..os.time())
    --print(msg)
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        ZLog.Logger("没有获取到房间".. tostring(userId))
        return
    end
    local player = GameServer.GetPlayerByUID(userId)
    player.scene = msg.location[1].scene        -- 记录一下该玩家所在场景

    local location = {}
    location = SandRockLocationNet.Copy(msg.location[1],location)
    SandRockRoom.LocationPlayerSet(room, userId, location)

end

--
-- 用来copy数据， 必须要有dec，因为格式需要提前初始化
function SandRockLocationNet.Copy(source, dec)
    dec.userId = source.userId
    dec.x = source.x
    dec.y = source.y
    dec.z = source.z
    dec.faceDir = source.faceDir
    dec.action = source.action
    dec.param = source.param
    dec.item = source.item
    dec.scene = source.scene
    return dec
end

--
---- 手持道具同步
--function SandRockLocationNet.PlayerHold(serverId, userId, buf)
--    local msg = ProtoGameSandRock.PlayerHold()
--    msg:ParseFromString(buf)
--    print("手持道具同步")
--    print(msg)
--    local item = msg.item
--    local room = GameServer.GetRoomByUserId(userId)
--    if room == nil then
--        ZLog.Logger("没有获取到房间".. tostring(userId))
--        return
--    end
--    local sendCmd = ProtoGameSandRock.PlayerHold()
--    sendCmd.item = item
--    sendCmd.userId = userId
--    -- 给该玩家下发
--    SandRockRoom.SendMsgToOtherUsers(room,CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_PLAYER_HOLD,sendCmd,userId)
--
--end
