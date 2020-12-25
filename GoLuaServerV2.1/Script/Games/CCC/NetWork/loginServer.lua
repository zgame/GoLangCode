

CCCNetworkLogin = {}


local function newUser(openId,machineId)
    local userId = GameServer.GetLastUserID()
    local user = User.New(userId,openId,machineId)
    CCCLoginDB.UserInsert(user)
    CCCLoginDB.OpenIdInsert(openId,user.userId)
    return user
end

-- 根据客户端发来的信息，进行登录， 优先级是 mac地址， openid, uid，
local function findUserInfo(msg)
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
        print(userId)
        user = CCCLoginDB.User(userId)
    end

    if user == nil then
        return newUser(openId,machineId)
    end

    return user
end


--游客登录申请,获取玩家的数据， 判断是否已经登录，
function CCCNetworkLogin.SevLoginGSGuest(serverId, buf)

    local msg = ProtoGameCCC.GameLogin()
    msg:ParseFromString(buf)

    -- 加载玩家数据
    local user = findUserInfo(msg)
    local player = Player(user)
    printTable(player)
    -- 将玩家的uid跟my server进行关联 ，方便以后发送消息
    luaCallGoResisterUID(player:UId(), serverId)


    --玩家登录游戏
    local game = GameServer.GetGameByID(msg.gameId)
    if game == nil then
        ZLog.Logger("没有找到游戏类型"..msg.gameId)
        return
    end
    Game.PlayerLoginGame(game,player)


    -- 发送消息
    local sendCmd = ProtoGameCCC.GameLoginResult()
    sendCmd.success = true
    Player.Copy(player,sendCmd.user)
    print(sendCmd)

    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_LOGON, sendCmd, nil)

    -- 给该玩家下发其他玩家信息
    -- sendCmd = ProtoGameCCC.OtherEnterRoom()
    -- local cmd
    -- for k,v in pairs(self.FishArray) do
    --     cmd = sendCmd.user:add()
    -- end
    -- NetWork.Send(serverId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_OTHER_LOGON, sendCmd, nil)

    -- 给其他玩家下发该玩家信息





end
