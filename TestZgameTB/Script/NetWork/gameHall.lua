---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.
--- DateTime: 2019/6/1 11:35
---




-- 整点领取金币
function SendClockScore(userId)
    local sendCmd = CMD_GameServer_pb.CMD_C_Clock_Gold()
    sendCmd.index = 1
    LuaNetWorkSendToUser(userId, MDM_GF_GAME_TB, SUB_C_CLOCK_GOLD ,sendCmd,nil)
end

function HandleGetClockScore(userId,buf)
    local msg = CMD_GameServer_pb.CMD_S_Clock_Gold()
    msg:ParseFromString(buf)
    print(msg)
end

-- 签到
function SendSignInScore(userId)
    local sendCmd = CMD_GameServer_pb.CMD_C_Sign_In()
    sendCmd.index = 1
    LuaNetWorkSendToUser(userId, MDM_GF_GAME_TB, SUB_C_SIGN_IN ,sendCmd,nil)
end
function  HandleGetSignInScore(userId,buf)
    local msg = CMD_GameServer_pb.CMD_S_Sign_In()
    msg:ParseFromString(buf)
    print(msg)
end


-- 申请房间信息
function SendGameRoomInfo(userId)
    print("发送房间")
    local sendCmd = CMD_GameServer_pb.CMD_C_GameRoomInfo()
    LuaNetWorkSendToUser(userId, MDM_GF_GAME_TB, SUB_C_GAME_ROOM_INFO ,sendCmd,nil)
end

function HandleGetGameRoomInfo(userId,buf)
    local msg = CMD_GameServer_pb.CMD_S_GameRoomInfo()
    msg:ParseFromString(buf)
    print(msg)

end