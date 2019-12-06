---创建者:zjy
---时间: 2019/10/22 19:23
--- 玩家管理类

---@class ChatUserManager
ChatUserManager = {

}
---@return ChatUserManager
function ChatUserManager:New(t)
    t = t or {
        m_mapUserInfo = {},   -- 用户列表 [userid] = ChatServerUserInfo
        m_mapUserInfoPtr ={}, -- 在线用户列表 [userid] = ChatServerUserInfo
        --- m_mapSearchCache = {}, -- 搜索列表
    }
    setmetatable(t, self)
    self.__index = self
    return t
end
--- 添加用户
---@param info ChatServerUserInfo 类的结构体，玩家信息
function ChatUserManager:AddUser(info)
    if info == nil or info.UserId == nil then
        return nil;
    end
    self.m_mapUserInfo[tostring(info.UserId)] = info;
    --[[local pUser = self:FindUserInfoByUserID(info.UserId);
    if pUser then
        if info.sToken and info.sToken ~= "" then
            pUser:SetToken(info.sToken);
        end
        if pUser.sNick ~= info.sNick then
            --// 更新search缓存
            local tCache = self.m_mapSearchCache[pUser.sNick];
            if tCache ~= nil then
                tCache[ChatServerSearchType.Nick] = nil;
                if tCache[ChatServerSearchType.GameID] == nil then
                    self.m_mapSearchCache[pUser.sNick] = nil;
                end
            end
        end
        --self.m_mapUserInfo[tostring(info.UserId)] = info;
    else
        self.m_mapUserInfo[tostring(info.UserId)] = info;
    end]]--
    return  info;
end
--- 获取在线用户信息
---@return ChatServerUserInfo
---@param  userid number
function ChatUserManager:FindUserInfoPtrByUserID(userid)
    local info = self.m_mapUserInfoPtr[tostring(userid)];
    return info;
end

---获取玩家用户信息
---@param dwUserID 玩家id
---@return ChatServerUserInfo
function ChatUserManager:FindUserInfoByUserID(dwUserID)
    local info = self.m_mapUserInfo[tostring(dwUserID)];
    return info;
end
---玩家登录事件
---@param userid 玩家的UID
---@param pUser UserInfo类的对象，玩家数据
function ChatUserManager:OnEventPlayerLogin(userid,pUser)
    if pUser == nil then
        return;
    end
    self.m_mapUserInfoPtr[tostring(userid)] = pUser;
end
---玩家退出事件
---@param dwUserID 退出玩家的UID
function ChatUserManager:OnEventPlayerLogout(dwUserID)
    local tCache = self:FindUserInfoPtrByUserID(dwUserID);
    Logger(string.format("OnEventPlayerLogout:玩家[%d]",dwUserID))
    local userStrid = tostring(dwUserID)
    if tCache ~= nil then
        ChatServerUserInfo.Rest(tCache);
        --tCache:Rest();
        self.m_mapUserInfoPtr[userStrid] = nil;
    end
    --- 离线清理用户信息
    self.m_mapUserInfo[userStrid] = nil
end
---获取在线用户列表
function ChatUserManager:GetOnlineUserList()
    return self.m_mapUserInfoPtr;
end
----将玩家的信息写入Redis。chatServer需要按照GameID、和NickName遍历的两UserI三张表。
---NickName/UserID表关联对应的GameID。
---@param tUserInfo 玩家的信息数据
function ChatUserManager:SaveUserInfo2Redis(tUserInfo)
    --printTable(tUserInfo,0,"SaveUserInfo2Redis.tUserInfo");
    RedisSaveString(ByChatServerRedisDir.UserGameIDInfo..tUserInfo.GameId,tUserInfo.GameId, ZJson.encode(tUserInfo));
    RedisSaveString(ByChatServerRedisDir.UserNickNameInfo..tUserInfo.NickName,tUserInfo.NickName, tUserInfo.GameId);
    RedisSaveString(ByChatServerRedisDir.UserUIDInfo..tUserInfo.UserId,tUserInfo.UserId,tUserInfo.GameId);
    --- 添加user id存储信息
    RedisSaveString(ByChatServerRedisDir.UserIdInfo..tUserInfo.UserId,tUserInfo.UserId, ZJson.encode(tUserInfo))
    --RedisSaveString(ByChatServerRedisDir.UserGameIDInfo,tUserInfo.GameId, ZJson.encode(tUserInfo));
    --RedisSaveString(ByChatServerRedisDir.UserNickNameInfo,tUserInfo.NickName, tUserInfo.GameId);
end
----从Redis中依据GameID或者NickName查询玩家信息，如果没有找到，则到数据库中查询
---@param dwGameID 玩家的GameID
---@param sNickName 玩家的昵称
function ChatUserManager:GetUserInfo(dwGameID,sNickName,dwUserID)
    local tUserInfo = {};
    local sUserInfo = "";
    if dwGameID ~= nil then
        sUserInfo = RedisGetString(ByChatServerRedisDir.UserGameIDInfo..dwGameID,dwGameID);
    elseif sNickName ~= nil and sNickName ~= "" then
        local sCurGameID = RedisGetString(ByChatServerRedisDir.UserNickNameInfo..sNickName,sNickName);
        local dwCurGameID = tonumber(sCurGameID);
        if dwCurGameID ~= nil and dwCurGameID > 0 then
            sUserInfo = RedisGetString(ByChatServerRedisDir.UserGameIDInfo..sCurGameID,sCurGameID);
        end
    elseif dwUserID ~= nil then
        local sCurGameID = RedisGetString(ByChatServerRedisDir.UserUIDInfo..dwUserID,dwUserID);
        local dwCurGameID = tonumber(sCurGameID);
        if dwCurGameID ~= nil and dwCurGameID > 0 then
            sUserInfo = RedisGetString(ByChatServerRedisDir.UserGameIDInfo..sCurGameID,sCurGameID);
        end
    end
    if string.len(sUserInfo) > 0 then
        tUserInfo = ZJson.decode(sUserInfo);
    end
    --printTable(tUserInfo,0,"tUserInfo")
    if tUserInfo and next(tUserInfo) then
        -- print(string.format("ChatUserManager:GetUserInfo 在redis查询到玩家[gameid=%d,sNickName=%s,dwUserID=%d]信息",dwGameID or 0,sNickName or "",dwUserID or 0));
        return tUserInfo;
    elseif dwGameID ~= nil then
        ---这个时候需要到数据库查询数据，如果有结果需写入redis
        --print(string.format("UserManager:GetUserInfo 到数据库查询玩家[gameid=%d]信息",dwGameID));
        local tUserCache = ChatServerUserInfo:New();
        local tDboData = ChatServerLogic.OnLoadUserInfo(nil,Enum_DBSearchType.GameID,dwGameID)
        --printTable(tDboData,0,"ChatUserManager:GetUserInfo.tDboData")
        if tDboData ~= nil  and next(tDboData)  then
            tUserCache.UserId = tDboData[1].UserID;
            tUserCache.GameId = tDboData[1].GameID;
            tUserCache.FaceId = tDboData[1].FaceID;
            tUserCache.VipLevel = tDboData[1].VipLev;
            tUserCache.NickName = tDboData[1].NickName;
            tUserCache.GuildID = tDboData[1].GuildID;
            tUserCache.GuildName = tDboData[1].GuildName;
            self:SaveUserInfo2Redis(tUserCache);
            return tUserCache;
        end
    elseif sNickName ~= nil and sNickName ~= "" then
        ---这个时候需要到数据库查询数据，如果有结果需写入redis
       -- print(string.format("ChatUserManager:GetUserInfo 到数据库查询玩家[NickName=%s]信息",sNickName));
        local tUserCache = ChatServerUserInfo:New();
        local tDboData = ChatServerLogic.OnLoadUserInfo(nil,Enum_DBSearchType.Nick,sNickName)
        --printTable(tDboData,0,"ChatUserManager:GetUserInfo.tDboData")
        if tDboData ~= nil  and next(tDboData)  then
            tUserCache.UserId = tDboData[1].UserID;
            tUserCache.GameId = tDboData[1].GameID;
            tUserCache.FaceId = tDboData[1].FaceID;
            tUserCache.VipLevel = tDboData[1].VipLev;
            tUserCache.NickName = tDboData[1].NickName;
            tUserCache.GuildID = tDboData[1].GuildID;
            tUserCache.GuildName = tDboData[1].GuildName;
            self:SaveUserInfo2Redis(tUserCache);
            return tUserCache;
        end
    elseif dwUserID ~= nil and dwUserID > 0 then
        ---这个时候需要到数据库查询数据，如果有结果需写入redis
       -- print(string.format("ChatUserManager:GetUserInfo 到数据库查询玩家[userid=%d]信息",dwUserID));
        local tUserCache = ChatServerUserInfo:New();
        local tDboData = ChatServerLogic.OnLoadUserInfo(nil,Enum_DBSearchType.UserID,dwUserID)
        --printTable(tDboData,0,"ChatUserManager:GetUserInfo.tDboData")
        if tDboData ~= nil  and next(tDboData)  then
            tUserCache.UserId = tDboData[1].UserID;
            tUserCache.GameId = tDboData[1].GameID;
            tUserCache.FaceId = tDboData[1].FaceID;
            tUserCache.VipLevel = tDboData[1].VipLev;
            tUserCache.NickName = tDboData[1].NickName;
            tUserCache.GuildID = tDboData[1].GuildID;
            tUserCache.GuildName = tDboData[1].GuildName;
            self:SaveUserInfo2Redis(tUserCache);
            return tUserCache;
        end
    end
    print(string.format("ChatUserManager:GetUserInfo 未查询玩家[gameid=%d,nick=%s,userid=%d]信息",dwGameID or 0,sNickName or "",dwUserID or 0));
    return nil
end



---获取存在的玩家用户信息  先在用户信息里面找，没找到到redis,redis没找到到数据库查找
--- 最后返回 存在的用户数据 如果用户不存在 将返回nil
---@param dwUserID 玩家id
---@return ChatServerUserInfo
function ChatUserManager:GetExistUserInfo(dwUserID)
    local info = self:FindUserInfoByUserID(dwUserID)
    if info == nil then  --- 在内存中未找到
        --- 先在redis 查询
        local sUserInfo = RedisGetString(ByChatServerRedisDir.UserIdInfo..dwUserID,dwUserID);
        if string.len(sUserInfo) > 0 then
            info = ZJson.decode(sUserInfo);
        end
        --- redis 找到就返回　
        if info  and next(info)  then
            --print(string.format("ChatUserManager:GetExistUserInfo 在redis查询到玩家[dwUserID=%d]信息",dwUserID or 0));
            return info
        end

        if ChatServerLogicInstance then
            --- 从数据库查找
            local resT = ChatServerLogicInstance:OnLoadUserInfo(Enum_DBSearchType.UserID, dwUserID)
            if resT ~= nil and next(resT) ~= nil then
                ---取第一个下标用数据库数据创建用户信息
                info = ChatServerLogicInstance:NewUserInfoByDBData(resT[1])
                if info  then
                    --- 刷新redis 数据
                    self:SaveUserInfo2Redis(info)
                end
                return info;
            end
        end
    end

    return info
end


---玩家离线时设置玩家离线时间
---@param userid number用户id
---@param times number离线时间戳
function ChatUserManager:SetUserOfflineTime(userid,times)
    local userInfo =self.m_mapUserInfo[tostring(userid)]
    if userInfo == nil then
        return
    end
    userInfo.llOffLineTime = times
end

--- 重新加载
function ChatUserManager:Reload(o)
    G_SetMetaTable(o,self)
end

--- 定时删除没有在线的用户
function ChatUserManager:RemoveOfflineUser()
    --- 遍历用户信息
    for userid,_ in pairs(self.m_mapUserInfo) do
        --- 没有在线的用户
        if self.m_mapUserInfoPtr[userid] == nil  then
            --- 清除 用户信息
            self.m_mapUserInfo[userid]  = nil
        end
    end

end