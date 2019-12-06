---创建者:zjy
---时间: 2019/10/25 15:17
--- 申请控制

---@class ChatApplyController
ChatApplyController = {

}

---@return ChatApplyController
function ChatApplyController:New()
    local t = {
        dwUserID = 0, -- 用户ID
        btVipMin = 0, --  最小VIP
        btVipMax = 0, --  最大VIP
        wGunLvMin = 0, --  最小炮等级
        wGunLvMax = 0, --  最大炮等级
        btUserLvMin = 0, --  最小玩家等级
        btUserLvMax = 0, --  最大玩家等级
        dwOnlineMin = 0, --  最小在线时长
        dwOnlineMax = 0, --  最大在线时长
        dwContextID = 0, -- 玩家ContextID
        bIsAddApply = true  --  是否推送申请
    }
    setmetatable(t, self)
    self.__index = self
    return t
end