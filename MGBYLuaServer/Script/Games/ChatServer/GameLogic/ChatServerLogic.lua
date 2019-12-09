--- 聊天服总接口 类似c++ Attemper
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.
--- DateTime: 2019/10/22 14:13


-------------------- 处理自动申请好友  与获取聊天消息--------


---@class ChatServerLogic 聊天服类对外总访问 ，处理网络相关的消息
ChatServerLogic = {
}

--- 构造函数
---@return ChatServerLogic
function ChatServerLogic:New()
    local t = {
        ---@type ChatUserManager
        UserMgr = {}, -- 用戶管理
        ---@type ChatFriendManager
        FriendMgr = {}, -- 好友管理
        --ConfineContent = {}, -- 屏蔽字 ["sdsad"] = {[1] = ""}
        m_WillSaveMessagelist = {}, -- 聊天消息缓存  ChatWillSaveMessage 列表
        m_iPlayerCount = 0, --在线玩家数
        m_pApplyControllerList = {}, ---后台申请好友列表,这个是数组 为优化算法修改为hash表，userid为key
        m_bApplyController = false, ---是否触发推送
        m_LinkInfo = {}, --连接信息
        m_iLoadMessageInDays = 5, --从数据库加载N天之内的未读消息
        m_tConfineContent = {}, --敏感词库
        g_iServerCount = 0, -- 服务器连接数
    }
    G_SetMetaTable(t, self)
    return t
end

--- 初始化类
function ChatServerLogic:OnInitLogic()
    -- 初始化相关消息
    self.UserMgr = ChatUserManager:New()
    self.FriendMgr = ChatFriendManager:New()
    self.m_iLoadMessageInDays = ConstChatServerLoadMessageInDays;
    --初始化敏感词
    self:OnLoadConfineContent();
    Logger("加载敏感词完成");
    -- printTable(self.m_tConfineContent);
end




--region   过滤屏蔽字 屏蔽字算法====================
---　过滤屏蔽字 在屏蔽字库里面查找
---@param Content 待检查敏感词的原始字符串
---@return string
function ChatServerLogic:FliterConfineContent(Content)
    local strlen = string.len(Content)
    if strlen == 0 or next(self.m_tConfineContent) == nil then
        return
    end
    local nTcharLen = string.len("*");
    ---这里缓存已经检测过的敏感字分类
    local tExeConfineContent = {}
    local tchContentKey = 0
    local tSubConfineContent = {}
    local nStartConfinePos = 1
    local nFindStarPos = 1
    local nFindEndPos = 1
    for i = 1, strlen do
        tchContentKey = string.sub(Content, i, i);
        if tchContentKey ~= nil and not tExeConfineContent[tchContentKey] then
            tExeConfineContent[tchContentKey] = true;
            tSubConfineContent = self.m_tConfineContent[tchContentKey] or {};
            for _, strConfine in ipairs(tSubConfineContent) do
                --print(Content,strConfine)
                ---因为一些敏感词包含特殊符号（乙酸[含量＞80%]之类），string.gsub不能用
                --Content = string.gsub(Content,strConfine,"*");
                while nFindStarPos ~= nil do
                    nFindStarPos, nFindEndPos = string.find(Content, strConfine, nStartConfinePos, true);
                    if nFindEndPos ~= nil and nFindStarPos ~= nil then
                        --找到敏感词，替换掉
                        nStartConfinePos = nFindStarPos + nTcharLen;
                        Content = string.format("%s*%s", string.sub(Content, 1, nFindStarPos - 1), string.sub(Content, nFindEndPos + 1));
                    end
                end
                nStartConfinePos = 1;
                nFindStarPos = 1;
                nFindEndPos = 1;
            end
        end
    end
    return Content
end
--endregion --------------------



--- 检查聊天服大部分table 长度
function ChatServerLogic:CheckTableLen()
    local checkStr = "------------------lua table 长度统计------------------------------ \n"
    checkStr = checkStr .. string.format('m_LinkInfo len = %d\n', self:GetLinkNum())
    checkStr = checkStr .. string.format('m_pApplyControllerList len = %d\n', GetHashTableLen(self.m_pApplyControllerList))
    checkStr = checkStr .. string.format('WillSaveMessagelist len = %d\n', GetHashTableLen(self.m_WillSaveMessagelist))
    checkStr = checkStr .. string.format('UserMgr.m_mapUserInfo len = %d\n', GetHashTableLen(self.UserMgr.m_mapUserInfo))
    checkStr = checkStr .. string.format('UserMgr.m_mapUserInfoPtr len = %d\n', GetHashTableLen(self.UserMgr.m_mapUserInfoPtr))
    --checkStr =checkStr .. string.format('m_tConfineContent len = %d\n',GetHashTableLen(self.m_tConfineContent))
    --for key,_ in pairs(self.m_tConfineContent) do
    --    checkStr =checkStr .. string.format('m_tConfineContent[%s] len = %d\n',key,GetHashTableLen(self.m_tConfineContent[key]))
    --end
    local friendNum = 0
    local applyFriendListNum = 0
    local beAppliednum = 0
    checkStr = checkStr .. string.format('FriendMgr.m_mapFriendRelation len = %d\n', GetHashTableLen(self.FriendMgr.m_mapFriendRelation))
    for i, friendRelation in pairs(self.FriendMgr.m_mapFriendRelation) do
        if friendRelation then
            friendNum = friendNum + friendRelation:GetFriendNum()
            applyFriendListNum = applyFriendListNum + friendRelation:GetApplyFriendListNum()
            beAppliednum = beAppliednum + GetHashTableLen(friendRelation:GetBeAppliedForFriendList())
        end
    end
    checkStr = checkStr .. string.format('friendNum = %d ,applyFriendListNum =%d,beAppliednum = %d \n', friendNum, applyFriendListNum, beAppliednum)
    checkStr = checkStr .. string.format('FriendMgr.m_mapUnreadMessageCache len = %d\n', GetHashTableLen(self.FriendMgr.m_mapUnreadMessageCache))
    local cacheLstlen = 0
    for _, messageCache in pairs(self.FriendMgr.m_mapUnreadMessageCache) do
        for _, mapFriendCache in pairs(messageCache.mapFriendCache) do
            cacheLstlen = cacheLstlen + GetHashTableLen(mapFriendCache.lstMsg)
        end
    end
    checkStr = checkStr .. string.format(' 总共有多少条未读消息lstMsg len = %d\n', cacheLstlen)

    Logger(checkStr)
end