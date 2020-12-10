---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by zhushiwei.
--- DateTime: 2020/12/8 17:13
---

GameNetwork ={}

function GameNetwork.Init(serverId)

end


-- 根据命令进行分支处理
function GameNetwork.Receive(serverId, userId, msgId, subMsgId, data, token)

    --print("msgId",msgId, "subMsgId",subMsgId)
    UserToken = token           -- 保存到全局里面，发送的时候取出来GameMessage
    if msgId == MDM_MB_LOGON then
        --if subMsgId == SUB_MB_GUESTLOGIN  then
        --    --print("**************游客登录服申请******************* ")
        --end
    elseif msgId == MDM_GR_LOGON then
        if subMsgId == SUB_GR_LOGON_USERID then
            print("**************游客登录游戏服申请******************* ")
            ----这里是原来的登录， 主要是返回客户端玩家的一些数据
            SevLoginGSGuest(serverId, data)      -- 返回给客户端，玩家的数据，用来显示的
        end
    elseif msgId == MDM_GF_FRAME then
        if subMsgId == SUB_GF_GAME_OPTION then
            print("**************游游客进入游戏房间申请***************** ", userId)
            ---- 这里是玩家申请登录游戏的类型，进入游戏房间， 分配房间坐下开始玩 , 客户端需要申请房间的类型
            SevEnterScene(userId, data)
        end
    elseif msgId == MDM_GR_USER then
        if subMsgId == SUB_GR_USER_STANDUP then
            -- 数据验证
            local player, game, gameTable = GetPlayer_Game_Table(userId)
            if player == nil or game == nil or gameTable == nil then
                Logger("玩家数据：" .. player .. ";game:" .. game .. ";table:" .. gameTable)
                return
            end
            gameTable:PlayerStandUp(player.ChairID, player)
            local sendResult = CMD_GameServer_pb.CMD_GR_S_UserStandUp()
            sendResult.result_code = 0
            LuaNetWorkSendToUser(userId, MDM_GR_USER, SUB_GR_S_USER_STANDUP, sendResult)
        end
    elseif msgId == MDM_GF_GAME then
        if subMsgId == SUB_C_USER_FIRE then
            print("**************客户端开火***************** ", userId)
            HandleUserFire(userId, data)
        elseif subMsgId == SUB_C_CATCH_FISH then
            print("*************客户端抓鱼***************** ", userId)
            HandleCatchFish(userId, data)

        elseif subMsgId == SUB_S_BOSS_COME then
            --print("*************暂时用来统计消息的返回时间***************** ",userId)
            --HandleStaticsNetWorkTime(userId)
        elseif subMsgId >= SUB_C_USE_SUMMON_GEM and subMsgId <= SUB_S_USER_DHS_INFO then
            --print("*************小海兽消息***************** ",userId)
            XhsReceiveMsg(serverId, userId, msgId, subMsgId, data, token)
        end
    elseif msgId == MAIN_CMD_ID then
        --print("*************日志服***************** ",serverId)
        if subMsgId == SUB_S_MONITOR_ITEMS then
            print("*************下发服务器列表***************** ", serverId)
        elseif subMsgId == SUB_S_MONITOR_STATE then
            --print("*************刷新服务器状态***************** ",serverId)
        end
    elseif msgId == MDM_GR_HEARTBEAT then
        if subMsgId == SUB_GR_C_HEARTBEAT then
            --print("*************处理玩家心跳***************** ",serverId)
            local receiveMsg = CMD_GameServer_pb.CMD_C_GAME_HEART_C2G()
            receiveMsg:ParseFromString(data)
            local sendMsg = CMD_GameServer_pb.CMD_C_GAME_HEART_C2G()
            sendMsg.target_user_id = userId
            LuaNetWorkSendToUser(userId, MDM_GR_HEARTBEAT, SUB_GR_S_HEARTBEAT, sendMsg)
        end
    elseif msgId == MAIN_CHAT_SERVICE_INNER then
        -- 登录服 到 聊天服消息
        ChatReceiveMsgFromLoginServer(serverId,subMsgId,data, token)
        -- ChatLoginServerToChatServerReceive(serverId, userId, subMsgId, data, token)
    elseif msgId == MAIN_CHAT_SERVICE_CLIENT then
        -- 客户端 到 聊天服消息
        ChatReceiveMsgFromClient(serverId, userId, subMsgId, data, token)
        --ChatClientToChatServerReceive(serverId, userId, subMsgId, data, token)
    else

    end
end



--- go通知lua 所有掉线的连接都要走这里
function GameNetwork.Broken(uid, serverId)
    ZLog.Logger("通知：" .. uid .. "  掉线了")

    local player = GameServer.GetPlayerByUID(uid)
    if player ~= nil then
        --printTable(player,0,"LeavePlayer")
        --print("LeavePlayer.UID="..player.User.UserId)
        local game = GameServer.GetGameByID(player.GameType)
        --printTable(game)
        if game ~= nil then
            Game.PlayerLogOutGame(game,player)
            --player.NetWorkState = false
            --player.NetWorkCloseTimer = GetOsTimeMillisecond()
        end
    end
end