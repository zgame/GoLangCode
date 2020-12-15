---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/12/5 14:05
---

--local CMD_GameServer_pb = require("CMD_GameServer_pb")
----游客登录申请,获取玩家的数据， 判断是否已经登录，
--function SevLoginGSGuest(serverId,buf)
--    local msg = CMD_GameServer_pb.CMD_GR_LogonUserID()
--    msg:ParseFromString(buf)
--
--    --print("gamekind id: ".. msg.kind_id)
--    --print("user_id id: ".. msg.user_id)
--    --print("machine_id : ".. msg.machine_id)
--
--    local MyUser
--    local openId = msg.machine_id
--    local UserId = RedisGetPlayerLogin(openId)
--    if UserId == "" then
--        --print("没有账号，创建一个")
--        UserId = GetLastUserID()
--        MyUser = User:New()
--        MyUser.FaceId = 0
--        MyUser.Gender = 0
--        MyUser.UserId = UserId
--        MyUser.GameId = 320395999
--        MyUser.Exp = 254
--        MyUser.Loveliness = 0
--        MyUser.Score = 100000009
--        MyUser.NickName = "玩家"..MyUser.UserId
--        MyUser.Level = 1
--        MyUser.VipLevel = 0
--        MyUser.AccountLevel = 3
--        MyUser.SiteLevel = 0
--        MyUser.CurLevelExp = 0
--        MyUser.NextLevelExp = 457
--        MyUser.PayTotal = 0
--        MyUser.Diamond = 29
--        MyUser.OpenId = openId
--        RedisSavePlayerLogin(openId,UserId)
--        RedisSavePlayer(MyUser)           -- redis 数据库 save
--    else
--        --print("有账号，那么取出账号的信息")
--        UserId = tonumber(UserId)       -- 这里需要转一下到数字
--        MyUser = RedisGetPlayer(UserId)   -- redis load
--    end
--
--    --printTable(MyUser)
--    --MyGameType = msg.kind_id
--    --if MyGame == nil then
--    --    Logger("请求登录游戏类型不正确"..msg.kind_id)
--    --    LuaNetWorkSend( MDM_GR_LOGON, SUB_GR_LOGON_FAILURE, nil, "请求登录游戏类型不正确")
--    --    return
--    --end
--
--    -- 先判断该玩家是否已经登录了， 如果登录了，返回错误消息给客户端告知已经登录了，可以踢掉之前登录的账号。
--    -- 为了统一，可以用数据库来判断
--
--    --result.UserId = GetLastUserID()       -- 桌子会发送消息给玩家
--    -- 去游戏管理器申请一个玩家uid
--    --local result = MultiThreadChannelGameManagerToPlayer("GetLastUserID",nil)    -- 申请分配一个桌子， 返回的数据中带有桌子和椅子的id了
--    --print("分配uid", result.UserId)
--
--
--
--    --RedisSavePlayer(MyUser)           -- redis 数据库 save
--    --local user = RedisGetPlayer(UserId)   -- redis load
--    --printTable(user)
--
--
--    --local player = Player:New(MyUser)
--    --player.GameType = msg.kind_id
--
--    -- 将玩家的uid跟my server进行关联 ，方便以后发送消息
--    luaCallGoResisterUID(UserId,serverId)
--
--    -- 发送登录成功
--    local sendCmd = CMD_GameServer_pb.CMD_GR_LogonSuccess()
--    sendCmd.user_right = UserId          -- 把生成的uid发送给客户端，让客户端以后用这个uid来登录
--    sendCmd.server_id = 99099
--
--    --    LuaNetWorkSend( MDM_GR_LOGON, SUB_GR_LOGON_SUCCESS, data, " 这是测试错误")
--    LuaNetWorkSendToUser(UserId, MDM_GR_LOGON, SUB_GR_LOGON_SUCCESS, sendCmd, nil)
--    LuaNetWorkSendToUser(UserId, MDM_GR_LOGON, SUB_GR_LOGON_FINISH, nil, nil)
--
--end



--------------------------------------------客户端----------------------------------------
local function build_mac_addr(index)
    local str = "74-D4-36-AD-"
    local i1 = math.floor(index / 256)
    local i2 = math.floor(index % 256)

    local ss1 = string.format("%x",i1)
    if #ss1 == 1 then
        ss1 = "0" .. ss1
    end

    str = str .. ss1
    str = str .. "-"
    local ss2 = string.format("%x",i2)
    if #ss2 == 1 then
        ss2 = "0" .. ss2
    end
    str = str .. ss2

    return str
    --return "74-D4-36-AD-1E-EE"
end

LoginServer={}

-- 发送登录服务器请求
function LoginServer.SendLogin(serverId)
    local sendCmd = Proto_Game_CCC.GameLogin()

    sendCmd.machine_id = build_mac_addr(serverId)
    LuaNetWorkSendUdp(serverId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_LOGON,sendCmd,nil)
    print("-----申请登录serverId:----------",serverId)
end

-- 服务器返回登录服务器成功，下发uid
function LoginServer.LoginGameServer(serverId, buf)
    local msg = Proto_Game_CCC.GameLoginResult()
    msg:ParseFromString(buf)
    local uid = msg.user_right

    --print( msg)

    --luaCallGoResisterUID(uid,serverId)
    print("--------登录成功---UID:",uid)
end