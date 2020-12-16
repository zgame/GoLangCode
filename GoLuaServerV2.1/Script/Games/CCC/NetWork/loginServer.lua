---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/12/5 14:05
---

CCCNetWorkLogin = {}

--游客登录申请,获取玩家的数据， 判断是否已经登录，
function CCCNetWorkLogin.SevLoginGSGuest(serverId, buf)

    local msg = Proto_Game_CCC.GameLogin()
    msg:ParseFromString(buf)
    --msg:ParseFromString(sendCmd:SerializeToString())

    print(msg)
    local UserId = 2121
    local sendCmd = Proto_Game_CCC.GameLoginResult()

    sendCmd.success = true

    sendCmd.user.user_id = UserId
    sendCmd.user.open_id = msg.machine_id
    sendCmd.user.nick_name = "test player"
    --print(sendCmd.user.base_user.nick_name)
    sendCmd.room_id = 99099

    --print(sendCmd)


    --    LuaNetWorkSend( MDM_GR_LOGON, SUB_GR_LOGON_SUCCESS, data, " 这是测试错误")

    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_LOGON, sendCmd, nil)
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_LOGON, sendCmd, nil)


    --NetWork.SendToUser(UserId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_LOGON, sendCmd, "message~!$", nil)

end
local function ss()
    --print("gamekind id: ".. msg.kind_id)
    --print("user_id id: ".. msg.user_id)
    --print("machine_id : ".. msg.machine_id)

    local MyUser
    --local openId = msg.machine_id
    local UserId = msg.user_id

    -- 这里以后要进行账号，密码的判断


    if UserId == "" then
        --print("没有账号，创建一个")
        UserId = GetLastUserID()
        MyUser = User:New()
        MyUser.FaceId = 0
        MyUser.Gender = 0
        MyUser.UserId = UserId
        MyUser.GameId = 320395999
        MyUser.Exp = 254
        MyUser.Loveliness = 0
        MyUser.Score = 100000009
        MyUser.NickName = "玩家" .. MyUser.UserId
        MyUser.Level = 1
        MyUser.VipLevel = 0
        MyUser.AccountLevel = 3
        MyUser.CurLevelExp = 0
        MyUser.NextLevelExp = 457
        MyUser.PayTotal = 0
        MyUser.Diamond = 29
        MyUser.OpenId = openId
        --RedisSavePlayerLogin(openId,UserId)
        --print("保存",openId,UserId)
        --RedisSavePlayer(MyUser)           -- redis 数据库 save
        --print("保存玩家信息",UserId)
        --else
        --    --print("有账号，那么取出账号的信息")
        --    UserId = tonumber(UserId)       -- 这里需要转一下到数字
        --    MyUser = RedisGetPlayer(UserId)   -- redis load
    end


    -- 将玩家的uid跟my server进行关联 ，方便以后发送消息
    luaCallGoResisterUID(UserId, serverId)

    -- 发送登录成功
    local sendCmd = CMD_GameServer_pb.CMD_GR_LogonSuccess()
    sendCmd.user_right = UserId          -- 把生成的uid发送给客户端，让客户端以后用这个uid来登录
    sendCmd.server_id = 99099



    --    LuaNetWorkSend( MDM_GR_LOGON, SUB_GR_LOGON_SUCCESS, data, " 这是测试错误")
    LuaNetWorkSendToUser(UserId, MDM_GR_LOGON, SUB_GR_LOGON_SUCCESS, sendCmd, nil, nil)
    LuaNetWorkSendToUser(UserId, MDM_GR_LOGON, SUB_GR_LOGON_FINISH, nil, nil, nil)

end
