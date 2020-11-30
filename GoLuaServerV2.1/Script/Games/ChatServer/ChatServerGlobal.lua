--- 聊天服全局方法网络消息分发，初始化游戏服
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.
--- DateTime: 2019/10/22 16:49
---

---@type ChatServerLogic 聊天服接口
ChatServerLogicInstance = nil




-----接收登录服的消息 可以再内部区分 这里避免多次
function ChatLoginServerToChatServerReceive(serverId, userId, subMsgId, data, token)
    if ChatServerLogicInstance ~= nil then
        ChatServerLogicInstance:HandlerLoginServerAction(serverId, userId,subMsgId, data, token)
    else
        Logger("ChatServerLogicInstance == nil")
    end
end

---接收客户端消息接口
function ChatClientToChatServerReceive(serverId, userId, subMsgId, data, token)
    if ChatServerLogicInstance ~= nil then
        ChatServerLogicInstance:HandlerClientAction(serverId, userId, subMsgId, data, token)
    else
        Logger("ChatServerLogicInstance == nil")
    end
end


--- 根据vip等级获得好友最大数量
function GetFriendMaxNumByVipLv(VipLv)
    ---好友数量判断
    return GetExcelMgbyVipValue(VipLv,"friend_max_num") or  99999
end


--- 检测lua相关
function ChatCheckLuaMemery_G()
    if ChatServerLogicInstance then
        ChatServerLogicInstance:CheckTableLen()
    end
end


--- 设置元表 并把__index元方法绑定到元表
function G_SetMetaTable(oldT, metaT)
    if oldT == nil or  metaT == nil then
        return
    end
    setmetatable(oldT,metaT)
    metaT.__index = metaT
end

--- 存储消息
function G_ChatSaveUnReadMessage()
    if ChatServerLogicInstance ~= nil and ChatServerLogicInstance.FriendMgr ~= nil then
        local willMsg  = ChatServerLogicInstance.FriendMgr:GetWillSaveMessageList()
        --- 向数据库保存未读消息
        ChatServerLogicInstance:OnSaveUnreadMessage(willMsg)
    end
end


