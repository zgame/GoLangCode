---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.zhengxh
--- DateTime: 2019/12/2 16:43
---

--------------------------------------------------------------------------------------
---  主要处理chat中玩家关系网的加载等相关逻辑
--------------------------------------------------------------------------------------


--------------------------------------------------------------------------------------
--- 加载玩家好友列表、被申请列表、申请列表玩家的数据并保存Redis
--- @param ChatUser 用户聊天相关信息
function LoadUserRelationPlayerInfoAndSaveToRedis(tbChatTable, tbChatUser)
    -- 用户判定
    if tbChatUser == nil then
        return
    end
    -- 需要判定的玩家列表
    local tbPlayerList = {}

    -- 1.申请好友列表
    for i, v in ipairs(tbChatUser.ApplyFriendArray) do
        tbPlayerList[tostring(v)] = v
    end
    -- 2.被申请列表
    for i, v in ipairs(tbChatUser.BeApplyFriendArray) do
        tbPlayerList[tostring(v)] = v
    end
    -- 3.好友列表
    for i, v in ipairs(tbChatUser.FriendArray) do
        tbPlayerList[tostring(v)] = v
    end

    local tbUser                = User:New()
    local tbFriendChatUser      = ChatUser:New()
    -- 处理这些玩家
    for i, v in pairs(tbPlayerList) do
        -- 玩家不再线的情况下处理这个
        if GetPlayerByUID(v) == nil then
            -- 取Redis是否存在
            if RedisExistKey(RedisDirAllPlayers..v) == 0 then
                LoadPlayerInfoAndSaveToRedis(v, tbUser, tbFriendChatUser)
                -- 加入缓存信息中
                tbChatTable:SetUserCache(tbUser)
            else
                tbUser = RedisGetPlayer(v)
                -- 加入缓存信息中
                tbChatTable:SetUserCache(tbUser)
            end
        end
    end
end

--------------------------------------------------------------------------------------
--- 加载玩家好友列表、被申请列表、申请列表玩家的数据并保存Redis
--- @param dwUserID     用户UserID
--- @param tbUser       用户对象信息
--- @param tbChatUser   用户聊天信息
function LoadPlayerInfoAndSaveToRedis(dwUserID, tbUser, tbChatUser)
    -- 1.不存在,取Sql server
    LoadUserInfoFromSQLServer(dwUserID, tbUser)
    -- 2.保存玩家基础信息到Redis
    RedisSavePlayer(tbUser)
    -- 3.加载好友聊天信息
    LoadUserChatInfoFromSQLServer(dwUserID, tbChatUser)
    -- 4.保存该玩家好友信息到Redis
    RedisSavePlayerChatInfo(dwUserID, tbChatUser)
    -- 5.处理玩家Nick<--->UserID对应关系
    if RedisChatFindUserIDByNickName(tbUser.NickName) == nil then
        -- 不存在写入
        RedisChatSaveNickNameToUserID(tbUser.NickName, dwUserID)
    end
    -- 6.处理玩家GameID<--->UserID对应关系
    if RedisChatFindUserIDByGameID(tbUser.GameID) == nil then
        -- 不存在写入
        RedisChatSaveGameIDToUserID(tbUser.GameID, dwUserID)
    end
end

--------------------------------------------------------------------------------------
--- 通过玩家搜索key在Redis中查找UserID
--- @param sKey 查找Key
--- @return UserID
function FindUserIDBYSearchKey(sKey)
    local dwUserID
    -- 1.先找GameID
    if tonumber(sKey) ~= nil then
        dwUserID = RedisChatFindUserIDByGameID(tonumber(sKey))
    end
    if dwUserID == nil then
        -- 2.GameID未找到再找NickName
        dwUserID = RedisChatFindUserIDByNickName(sKey)
    end
    return dwUserID
end

--------------------------------------------------------------------------------------
--- 加载屏蔽字库
--- @param tbConfineContent         屏蔽字容器
--- @param tbConfineContentArray    屏蔽字完整字符串容器
function LoadConfineContent(tbConfineContentDict, tbConfineContentArray)
    local sSql = string.format("SELECT String FROM ConfineContent;")
    local ret = SqlServerDataBaseBYQuery(sSql)
    if ret ~= nil and #ret > 0 then
        for i, v in ipairs(ret) do
            local word = v.String
            local t = tbConfineContentDict
            for j = 1, #word do
                local c = tostring(string.byte(word, j))
                if not t[c] then
                    t[c] = {}
                end
                t = t[c]
            end
            tbConfineContentArray[word] = true
        end
    end
end
--------------------------------------------------------------------------------------