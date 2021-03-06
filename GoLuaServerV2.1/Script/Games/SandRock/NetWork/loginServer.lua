SandRockLoginNet = {}

----------------------------------------登录----------------------------------
local function newUser(openId, machineId)
    --print("创建用户")
    local userId = GameServer.GetLastUserID()
    if userId == nil then
        ZLog.Logger("redis 出问题 没有获取到userId")
        return nil
    end
    local user = User.New(userId, openId, machineId)
    SandRockLoginDB.OpenIdInsert(openId, user.userId)
    SandRockLoginDB.UserInsert(user)
    return user
end

-- 根据客户端发来的信息，进行登录， 优先级是 mac地址， openid, uid，
local function getUserDB(msg)
    --print(msg)
    local userId = msg.userId
    local openId = msg.openId
    local machineId = msg.machineId
    local user

    if userId ~= 0 then
        -- 说明输入了userId，优先处理userId
        user = SandRockLoginDB.User(userId)
        if user ~= nil then
            return user
        end
    end

    if machineId ~= "" then
        -- 输入了machineId ，忽略openid
        openId = 'sandRockMac' .. machineId
    end

    -- 注意如果客户端， 同时发mac 和 openid， 会使用mac，忽略openid
    if openId ~= nil then
        userId = SandRockLoginDB.UId(openId)
        user = SandRockLoginDB.User(userId)
    end

    -- 最后看一下如果数据库没有数据，那么新建
    if user == nil then
        return newUser(openId, machineId)
    end

    return user
end


--游客登录申请,获取玩家的数据， 判断是否已经登录，
function SandRockLoginNet.Login(serverId, uId, buf)
    local msg = ProtoGameSandRock.GameLogin()
    msg:ParseFromString(buf)

    -- 加载玩家数据
    local user = getUserDB(msg)
    --print("-------------获取玩家数据-----------")
    --printTable(user)
    --print("---------------------------------")
    if user == nil then
        ZLog.Logger("数据库获取玩家数据出错")
        return
    end
    local player = Player(user)
    -- 将玩家的uid跟my server进行关联 ，方便以后发送消息
    local userId = Player.UId(player)
    luaCallGoResisterUID(userId, serverId)

    --玩家登录游戏
    if GameServer.Login(msg.gameId, player) == false then
        return
    end

    -- 发送消息
    local sendCmd = ProtoGameSandRock.GameLoginResult()
    sendCmd.success = true
    Player.Copy(player, sendCmd.user)
    --print(sendCmd)

    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_LOGON, sendCmd, nil)

    -- 给该玩家下发其他玩家信息
    SandRockLoginNet.SendPlayersInfo(userId)

    -- 同步场景信息给登录的玩家
    SandRockLoginNet.SendEnterSceneInfo(userId)

    -- 同步场景树信息给玩家
    SandRockLoginNet.SendTerrainInfo(userId)

    --下发该玩家道具信息
    SandRockLoginNet.SendItemInfo(userId)

    -- 同步采集资源点信息给玩家
    local room = GameServer.GetRoomByUserId(userId)
    SandRockResourcePickNet.SendPickList(userId, nil, room.resourcePoint, nil)

end

----------------------------------------同步----------------------------------
-- 给该玩家下发其他玩家信息
function SandRockLoginNet.SendPlayersInfo(userId)
    local sendCmd = ProtoGameSandRock.UserList()
    local room = GameServer.GetRoomByUserId(userId)
    for _, player in pairs(room.userSeatArray) do
        if player ~= nil and Player.UId(player) ~= userId then
            local uu = sendCmd.user:add()
            Player.Copy(player, uu)
        end
    end
    --print("下发其他玩家数据")
    --print(sendCmd)
    --print(sendCmd == nil)
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_ROOM_LIST, sendCmd, nil, nil)
end


-- 同步场景信息给登录的玩家
function SandRockLoginNet.SendEnterSceneInfo(userId)
    local sendCmd = ProtoGameSandRock.GameInfo()
    sendCmd.npcList = 1
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_ROOM_INFO, sendCmd, nil, nil)
end

-- 同步场景树信息给登录的玩家
function SandRockLoginNet.SendTerrainInfo(userId)
    local room = GameServer.GetRoomByUserId(userId)
    local reliveList = SandRockRoom.ResourceTerrainSync(room)
    SandRockResourceTerrainNet.SendTreeRelive(userId, reliveList)
end

-- 下发该玩家道具信息
function SandRockLoginNet.SendItemInfo(userId)
    local player = GameServer.GetPlayerByUID(userId)
    local sendCmd = ProtoGameSandRock.ItemGet()
    for itemId, info in pairs(player.user.package) do
        if type(info) == "number" then
            -- 可堆叠道具
            local item = sendCmd.item:add()
            item.itemId = tonumber(itemId)
            item.itemNum = info
            item.itemNumTotal = info
        else
            for key, value in pairs(info) do
                local item = sendCmd.item:add()
                item.itemId = tonumber(itemId)
                item.itemNum = 1
                item.itemNumTotal = 1
                item.itemUId = tonumber(key)
            end
        end
    end
    --print("-----下发道具信息-----")
    --print(sendCmd)
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_PLAYER_ITEM_INFO, sendCmd, nil, nil)
end

----------------------------------------登出----------------------------------
--登出申请
function SandRockLoginNet.Logout(serverId, uId, buf)
    local msg = ProtoGameSandRock.GameLogout()
    msg:ParseFromString(buf)

    GameNetwork.Broken(uId, serverId)
    -- 发送消息
    local sendCmd = ProtoGameSandRock.GameLogoutResult()
    sendCmd.success = true
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_LOGOUT, sendCmd, nil)

end