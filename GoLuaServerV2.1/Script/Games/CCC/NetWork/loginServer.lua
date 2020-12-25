

CCCNetworkLogin = {}


local function newPlayer(openId,machineId)
    local userId = GameServer.GetLastUserID()
    local user = User:New(userId,openId,machineId)
end

-- 根据客户端发来的信息，进行登录
local function findUserInfo(msg)
    print(msg)
    local userId = msg.userId
    local openId = msg.openId
    local machineId = msg.machineId

    if userId ~= nil then

    end

    if openId ~= nil then

    end

    if machineId ~= nil then
       
    end

    return nil
end


--游客登录申请,获取玩家的数据， 判断是否已经登录，
function CCCNetworkLogin.SevLoginGSGuest(serverId, buf)

    local msg = ProtoGameCCC.GameLogin()
    msg:ParseFromString(buf)

    local player = findUserInfo(msg)
    if player == nil then
        ZLog.Logger("登录错误"..msg)
        return
    end
    -- 将玩家的uid跟my server进行关联 ，方便以后发送消息
    luaCallGoResisterUID(player:UId(), serverId)


    --玩家登录游戏
    local gameId = msg.gameId
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
