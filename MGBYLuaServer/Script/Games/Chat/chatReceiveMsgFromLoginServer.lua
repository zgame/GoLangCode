---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.zhengxh
--- DateTime: 2019/11/28 11:30
---

--------------------------------------------------------------------------------------
--- 这里主要是聊天服接受登录服发过来的消息
---（该逻辑在最终整合服务器时可以删除）
--------------------------------------------------------------------------------------


--- 接收登录服网络消息
--- @param serverId 游戏serverId
--- @param subMsgId 消息子协议
--- @param data     消息体
--- @param token    token验证
function ChatReceiveMsgFromLoginServer(serverId, subMsgId, data, token)
    -- 找game
    local game = GetGameByID(GameTypeChat)
    if game == nil then
        Logger("登录服消息聊天服服务game未找到")
        return
    end
    -- 找桌子（聊天只会创建一个桌子）
    local chatTable = game:GetTableByUID(game.TableUUID-1)
    if chatTable == nil then
        Logger("登录服消息聊天服服务chatTable未找到")
        return
    end
    -- 派发消息
    if subMsgId == SUB_SC_REGISTER then
        -- 登录服注册聊天服务
        chatTable:HandleLoginServerRegister(serverId, data, token)
    elseif subMsgId == SUB_SC_PUSH_TOKEN  then
        -- 登录服推送玩家信息
        chatTable:HandleLoginServerPushToken(serverId, data, token)
    elseif subMsgId == SUB_SC_TOKEN_USER_INFO  then
        -- 登录服推送玩家相关数据信息变更(头像和VIP等级)
        chatTable:HandleLoginServerUpdateUserInfo(serverId, data, token)
    end
end