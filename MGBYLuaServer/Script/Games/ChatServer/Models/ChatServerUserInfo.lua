---创建者:zjy
---时间: 2019/10/22 19:31
---玩家信息类

---@class  ChatServerUserInfo
ChatServerUserInfo = User:New()

---@return ChatServerUserInfo
function ChatServerUserInfo:New()
    local t ={
        --[[UserId = 0,            -- 玩家ID
        GameId = 0,            -- 游戏ID
        FaceId = 0,            -- 头像ID
        VipLevel = 0,             -- vip等级
        Level = 0,            -- 用户等级
        OnlineTime = 0,        -- 在线时长
        sNick = '',            -- 昵称]]--
        GuildID= 0,            -- 公会id
        GuildName = '',        -- 公会名称
        llOffLineTime = 0,
        sToken = '',
        tStartTime = 0,
        dwBindIndex = 0,
        TFriendChatTime = {}, -- 好友聊天时间
        LastSearchTime = 0, -- 最近收缩时间
        ActiveTime = 0,  -- 连接时激活时间
    }
    setmetatable(t, self)
    self.__index = self
    return t
end

--- 设置token信息
---@param tokenstr 令牌号
---@return boolean 是否设置成功 ture = 是
function ChatServerUserInfo:SetToken(tokenstr)
    if string.len(tokenstr) ~= 32 then
        return false
    end
    self.sToken = tokenstr
    self.tStartTime = os.time()
    return true
end

--- 返回token字符串
function ChatServerUserInfo:GetToken()
    if self.tStartTime > 0 and self.tStartTime + TOKEN_LIVE_SEC <= os.time() then
        self.sToken = ''
        self.tStartTime = 0
    end

    return self.sToken
end


--- 检查是否可以和好友说话
---@param friendid 好友UID
---@param btinterval 说话的间隔
function ChatServerUserInfo:CanChatWithFriend(friendid,btinterval)
    local tCurTime = os.time()
    local friendt  = self.TFriendChatTime[tostring(friendid)]
    if friendt == nil then
         self.TFriendChatTime[tostring(friendid)] = tCurTime
    else
        -- 没有到时间
        if tCurTime - friendt <  btinterval then
            print(string.format("CanChatWithFriend:[speaker=%d,friendid=%d,friendt=%d,tCurTime=%d,btinterval=%d]间隔不够",self.UserId,friendid,friendt,tCurTime,btinterval))
            return false
        end
        -- 到了时间可以说话 重置时间
        self.TFriendChatTime[tostring(friendid)] = tCurTime
    end

    return true
end

-- 是否可以查找
---@param btInterval 时间间隔
function ChatServerUserInfo:CanSearchUser(btInterval)
    local tCurTime = os.time();
    if (tCurTime - self.LastSearchTime < btInterval) then
        return false
    end
    self.LastSearchTime = tCurTime;
    return true;
end


---重置信息
function ChatServerUserInfo:Rest()
    self.TFriendChatTime = {}
    self.LastSearchTime = 0
end

--- 重新加载类
function ChatServerUserInfo:Reload(o)
    --- 重新刷一次元表,以便调用新定义方法，更新老方法
    G_SetMetaTable(o,self)
    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end