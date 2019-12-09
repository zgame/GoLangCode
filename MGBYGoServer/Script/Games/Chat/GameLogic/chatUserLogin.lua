---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.zhengxh
--- DateTime: 2019/11/28 15:39
---

--------------------------------------------------------------------------------------
--- 玩家登录聊天服（该逻辑在最终整合服务器时可以删除）
--------------------------------------------------------------------------------------

--- 玩家登录聊天服
--- @param ServerID serverID用于在此绑定UserID
--- @param UserID   玩家UserID这个时候UserID都为0
--- @param data     消息数据
function ChatTable:HandleUserLogin(ServerID, UserID, data)
    -- 解析消息
    local receiveMsg = CMD_GlobalServer_pb.CMD_C_LOGIN()
    receiveMsg:ParseFromString(data)
    -- 获取UserID
    UserID = receiveMsg.user_id
    if UserID == nil or UserID <= 0 then
        Logger("登录上传UserID为nil或者小于0")
        return
    end
    if ServerID == nil or ServerID <= 0 then
        Logger("登录上传ServerID为nil或者小于0")
        return
    end
    -- 消息响应
    local sendMsg = CMD_GlobalServer_pb.CMD_S_LOGIN()
    -- 验证登录服上传的玩家Token
    --[[if self:GetPlayerToken(UserID) ~= receiveMsg.token then
        Logger(string.format("玩家上传的token与登录上传的token不同"))
        -- 判定玩家上传的token是否正确
        local nGameID = receiveMsg.game_id
        local sToken = receiveMsg.token
        local nRandNum = receiveMsg.rand
        local nRandNum2 = receiveMsg.rand2
        -- 做异或
        local nTempData = Bit32Xor(Bit32Or(UserID, nGameID), nRandNum)
        --检查token
        local sSrc = "Im#UserSave!" .. nTempData .. "&(" .. nRandNum .. ")+" .. nRandNum2
        local sRealToken = string.upper(MD5Get(sSrc))
        if sRealToken ~= sToken then
            print("非法连接，token不匹配：userid=" .. UserID .. ",gameid=" .. nGameID .. ",rand=" .. nRandNum .. ",token=" .. sToken)
            sendMsg.result = Enum_ReplyResult.Failed
            LuaNetWorkSendToUser(UserID, MAIN_CHAT_SERVICE_CLIENT, SUB_S_LOGIN, sendMsg)
            return;
        end
        -- 重置Token
        self:AddPlayerToken(UserID, sRealToken)
    end]]--
    -- 从数据库拉取玩家数据信息
    -- 创建一个User对象
    local MyUser
    if RedisExistKey(RedisDirAllPlayers..UserID) == 1 then
        MyUser = RedisGetPlayer(UserID)
    else
        MyUser = User:New()
        -- 加载玩家数据
        LoadUserInfoFromSQLServer(UserID, MyUser)
    end
    if MyUser.UserID ~= UserID then
        MyUser = nil
        Logger("聊天服玩家登录数据库中未找到该登录玩家信息,不处理该玩家登录聊天服"..UserID)
        return
    end
    --绑定serverId，nUserID，修改在线玩家数据
    luaCallGoResisterUID(UserID, ServerID)
    -- 处理玩家Nick<--->UserID对应关系
    if RedisChatFindUserIDByNickName(MyUser.NickName) == nil then
        -- 不存在写入
        RedisChatSaveNickNameToUserID(MyUser.NickName, UserID)
    end
    -- 处理玩家GameID<--->UserID对应关系
    if RedisChatFindUserIDByGameID(MyUser.GameID) == nil then
        -- 不存在写入
        RedisChatSaveGameIDToUserID(MyUser.GameID, UserID)
    end
    -- 保存玩家数据到Redis[目前每次登陆都保存]
    RedisSavePlayer(MyUser)
    -- 加载玩家聊天相关信息
    local chatUser
    -- 优先取Redis玩家聊天信息
    if RedisExistKey(RedisDirChatPlayers, UserID) == 1 then
        chatUser = RedisGetPlayerChatInfo(UserID)
    else
        chatUser = ChatUser:New()
        LoadUserChatInfoFromSQLServer(UserID, chatUser)     -- 通过数据库获取玩家数据
        RedisSavePlayerChatInfo(UserID, chatUser)           -- 玩家聊天信息存Redis
    end
    -- 创建玩家
    local player = Player:New(MyUser)
    player.ChatUser = chatUser
    player.GameType = GameTypeChat
    player.TableID  = 1
    -- 加入到玩家列表
    SetAllPlayerList(UserID, player)
    -- 加入缓存信息中
    self:SetUserCache(player.User)
    -- 返回成功
    sendMsg.result = Enum_ReplyResult.Successful
    LuaNetWorkSendToUser(UserID, MAIN_CHAT_SERVICE_CLIENT, SUB_S_LOGIN, sendMsg)
    -- 处理玩家关系网数据到Redis
    LoadUserRelationPlayerInfoAndSaveToRedis(self, chatUser)
end