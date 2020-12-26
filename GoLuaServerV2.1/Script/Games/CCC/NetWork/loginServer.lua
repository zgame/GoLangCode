

CCCNetworkLogin = {}


local function newUser(openId,machineId)
    local userId = GameServer.GetLastUserID()
    local user = User.New(userId,openId,machineId)
    CCCLoginDB.UserInsert(user)
    CCCLoginDB.OpenIdInsert(openId,user.userId)
    return user
end

-- 根据客户端发来的信息，进行登录， 优先级是 mac地址， openid, uid，
local function getUserDB(msg)
    print(msg)
    local userId = msg.userId
    local openId = msg.openId
    local machineId = msg.machineId

    local user
    if userId ~= nil then
        user =  CCCLoginDB.User(userId)
    end

    if machineId ~= nil then
        openId = 'cccmac'..machineId
    end

    -- 注意如果客户端， 同时发mac 和 openid， 会使用mac，忽略openid
    if openId ~= nil then
        userId = CCCLoginDB.UId(openId)
        user = CCCLoginDB.User(userId)
    end

    if user == nil then
        return newUser(openId,machineId)
    end

    return user
end


--游客登录申请,获取玩家的数据， 判断是否已经登录，
function CCCNetworkLogin.SevLoginGSGuest(serverId, uId, buf)

    local msg = ProtoGameCCC.GameLogin()
    msg:ParseFromString(buf)

    -- 加载玩家数据
    local user = getUserDB(msg)
    local player = Player(user)
    -- 将玩家的uid跟my server进行关联 ，方便以后发送消息
    local userId = Player.UId(player)
    luaCallGoResisterUID(userId, serverId)

    --玩家登录游戏
    if GameServer.Login(msg.gameId,player) == false then
        return
    end

    -- 发送消息
    local sendCmd = ProtoGameCCC.GameLoginResult()
    sendCmd.success = true
    Player.Copy(player,sendCmd.user)
    print(sendCmd)

    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_LOGON, sendCmd, nil)

    -- 给该玩家下发其他玩家信息
    CCCNetworkLogin.SendPlayersInfo(userId)
    -- 同步场景信息给登录的玩家
    CCCNetworkLogin.SendEnterSceneInfo(userId)
end


-- 给该玩家下发其他玩家信息
function CCCNetworkLogin.SendPlayersInfo(userId)
    local sendCmd = ProtoGameCCC.OtherEnterRoom()
    local room = GameServer.GetRoomByUserId(userId)
    for i, player in pairs(room.userSeatArray) do
        if player ~= nil then
            local uu = sendCmd.user:add()
            Player.Copy(player,uu)
        end
    end
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_OTHER_LOGON, sendCmd, nil, nil)
end


-- 同步场景信息给登录的玩家
function CCCNetworkLogin.SendEnterSceneInfo(userId)
    local sendCmd = ProtoGameCCC.GameInfo()
    sendCmd.npcList = 1
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_GAME_INFO, sendCmd, nil, nil)
end




--登出申请
function CCCNetworkLogin.SevLogout(serverId, uId, buf)
    local msg = ProtoGameCCC.GameLogout()
    msg:ParseFromString(buf)

    GameNetwork.Broken(uId, serverId)
    -- 发送消息
    local sendCmd = ProtoGameCCC.GameLogoutResult()
    sendCmd.success = true
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_LOGOUT, sendCmd, nil)

end