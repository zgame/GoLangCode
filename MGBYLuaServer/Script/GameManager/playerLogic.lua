---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.zhengxh
--- DateTime: 2019/10/28 15:27
---

--------------------------------------------------------------------------------------
--- 这个文件主要是做player对象的成员函数处理
--------------------------------------------------------------------------------------

--- 增加玩家Score值
--- @param addScore 增加值(必须为正)
function Player:AddScore(addScore)
    if addScore < 0 then return end
    self.User.Score = self.User.Score + addScore
end
--------------------------------------------------------------------------------------
--- 扣除玩家Score值
--- @param addScore 减少值(必须为正)
function Player:DecScore(decScore)
    if addScore < 0 then return end
    if self.User.Score <= decScore then
        self.User.Score = 0
    else
        self.User.Score = self.User.Score - decScore
    end
end
--------------------------------------------------------------------------------------
--- 获取道具，通过ItemID获取道具
--- @param dwItemID 道具ID
--- @return Item对象
function Player:GetItemByItemID(dwItemID)
    -- 道具ID判定
    if dwItemID == 0 then return nil end
    -- 道具配置判定
    if GetExcelItemValue(dwItemID, "item_type") == nil then return nil end
    for i, v in pairs(self.User.SkillInfoArray) do
        if v.ItemID == dwItemID then
            return v
        end
    end
    return nil
end
--------------------------------------------------------------------------------------
--- 获取道具的数量
--- @param dwItemID 道具ID
--- @return 道具数量
function Player:GetItemNumByItemID(dwItemID)
    -- 道具ID判定
    if dwItemID == 0 then return 0 end
    -- 道具配置判定
    if GetExcelItemValue(dwItemID, "item_type") == nil then return 0 end
    for i, v in pairs(self.User.SkillInfoArray) do
        if v.ItemID == dwItemID then
            return (v.Used - v.Total)
        end
    end
    return 0
end
--------------------------------------------------------------------------------------
--- 增加道具
--- @param dwItemID 增加的道具ID
--- @param addNum   增加的道具数量
--- @return bool成功与否
function Player:AddItem(dwItemID, addNum)
    -- 道具ID判定
    if dwItemID == 0 or addNum == 0 then return false end
    -- 道具配置判定
    if GetExcelItemValue(dwItemID, "item_type") == nil then return false end
    -- 获取玩家身上该道具
    local userItem = self:GetItemByItemID(dwItemID)
    if userItem then
        userItem.Total = userItem.Total + addNum
        -- 保存数据库
        SaveUserItemAdd(self.User.UserID, dwItemID, addNum, false)
    else
        self.User.SkillInfoArray[#self.User.SkillInfoArray+1] = {
            ItemID  = dwItemID,
            Used    = 0,
            Total   = addNum,
            UsedTime= 0
        }
        -- 保存数据库
        SaveUserItemAdd(self.User.UserID, dwItemID, addNum, true)
    end
    return true
end
--------------------------------------------------------------------------------------
--- 扣除道具
--- @param dwItemID 扣除的道具ID
--- @param decNum   扣除的道具数量
--- @return bool成功与否
function Player:DecItem(dwItemID, decNum)
    -- 道具ID判定
    if dwItemID == 0 or decNum == 0 then return false end
    -- 道具配置判定
    if GetExcelItemValue(dwItemID, "item_type") == nil then return false end
    -- 获取玩家身上该道具
    local userItem = self:GetItemByItemID(dwItemID)
    if userItem and (userItem.Total - userItem.Used) >= decNum then
        userItem.Used = userItem.Used + decNum
        -- 保存数据库
        SaveUserItemDec(self.User.UserID, dwItemID, decNum)
        return true
    end
    return false
end
--------------------------------------------------------------------------------------
--- 获取UserID是否在玩家申请列表
--- @param UserID 需要查询的UserID
--- @param 列表索引
function Player:IsInApplyFriendArray(dwUserID)
    for i, v in ipairs(self.ChatUser.ApplyFriendArray) do
        if v == dwUserID then
            return i
        end
    end
    return 0
end

--- 获取玩家申请列表数量
--- @return 当前申请列表数量
function Player:GetApplyFriendNum()
    return #self.ChatUser.ApplyFriendArray
end

--- 给玩家申请列表增加一条数据
--- @param dwApplyUserID 申请UserID
function Player:AddApplyFriendInfo(dwApplyUserID)
    table.insert(self.ChatUser.ApplyFriendArray, dwApplyUserID)
end

--- 从玩家申请列表删除一条数据
--- @param dwApplyUserID
function Player:DelApplyFriendInfo(dwApplyUserID)
    table.remove(self.ChatUser.ApplyFriendArray, self:IsInApplyFriendArray(dwApplyUserID))
end
--------------------------------------------------------------------------------------
--- 获取UserID是否在玩家被申请列表
--- @param UserID 需要查询的UserID
--- @param 列表索引
function Player:IsInBeApplyFriendArray(dwUserID)
    for i, v in ipairs(self.ChatUser.BeApplyFriendArray) do
        if v == dwUserID then
            return i
        end
    end
    return 0
end

--- 获取玩家被申请列表数量
--- @return 当前被申请列表数量
function Player:GetBeApplyFriendNum()
    return #self.ChatUser.BeApplyFriendArray
end

--- 给玩家被申请列表增加一条数据
--- @param dwBeApplyUserID 被申请UserID
function Player:AddBeApplyFriendInfo(dwBeApplyUserID)
    table.insert(self.ChatUser.BeApplyFriendArray, dwBeApplyUserID)
end

--- 从玩家被申请列表删除一条数据
--- @param dwBeApplyUserID 被申请UserID
function Player:DelBeApplyFriendInfo(dwBeApplyUserID)
    table.remove(self.ChatUser.BeApplyFriendArray, self:IsInBeApplyFriendArray(dwBeApplyUserID))
end
--------------------------------------------------------------------------------------
--- 获取UserID是否在玩家好友列表
--- @param UserID 需要查询的UserID
--- @param 列表索引
function Player:IsInFriendArray(dwUserID)
    for i, v in ipairs(self.ChatUser.FriendArray) do
        if v == dwUserID then
            return i
        end
    end
    return 0
end

--- 获取玩家好友列表数量
--- @return 当前好友列表数量
function Player:GetFriendNum()
    return #self.ChatUser.FriendArray
end

--- 给玩家好友列表增加一条数据
--- @param dwFriendUserID 好友UserID
function Player:AddFriendInfo(dwFriendUserID)
    table.insert(self.ChatUser.FriendArray, dwFriendUserID)
end

--- 从玩家好友列表删除一条数据
--- @param dwFriendUserID 好友UserID
function Player:DelFriendInfo(dwFriendUserID)
    table.remove(self.ChatUser.FriendArray, self:IsInFriendArray(dwFriendUserID))
end
--------------------------------------------------------------------------------------

--- 添加一条好友的未读消息
--- @param dwFriendID 好友ID
--- @param dwEmotion  表情ID
--- @param strMessage 消息内容
function Player:AddFriendUnReadMessage(dwFriendID, dwEmotion, strMessage)
    if self.ChatUser.UnReadMessageArray[tostring(dwFriendID)] == nil then
        self.ChatUser.UnReadMessageArray[tostring(dwFriendID)] = {}
    end
    local tbMessage = {
        dwEmotion,
        os.time(),
        strMessage
    }
    table.insert(self.ChatUser.UnReadMessageArray[tostring(dwFriendID)], tbMessage)
end

--- 清空某个玩家发来的未读消息
--- @param dwFriendID 好友ID
function Player:DelFriendUnReadMessage(dwFriendID)
    if self.ChatUser.UnReadMessageArray[tostring(dwFriendID)] ~= nil then
        self.ChatUser.UnReadMessageArray[tostring(dwFriendID)] = nil
    end
end

--- 获取某个玩家发来的未读消息
--- @param dwFriendID 好友ID
--- @return 未读消息列表
function Player:GetFriendUnReadMessage(dwFriendID)
    if self.ChatUser.UnReadMessageArray[tostring(dwFriendID)] ~= nil then
        return self.ChatUser.UnReadMessageArray[tostring(dwFriendID)]
    end
    return nil
end