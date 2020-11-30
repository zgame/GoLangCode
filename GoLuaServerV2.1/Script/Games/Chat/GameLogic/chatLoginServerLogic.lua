---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.zhengxh
--- DateTime: 2019/11/28 11:28
---

--------------------------------------------------------------------------------------
--- 处理登录服与聊天服之间的逻辑
---（该逻辑在最终整合服务器时可以删除）
--------------------------------------------------------------------------------------

--- 登录服注册聊天服务
--- @param serverId 服务器id
--- @param data     数据
--- @param token
function ChatTable:HandleLoginServerRegister(serverId, data, token)
    -- 解析消息
    local receiveMsg =  CMD_GlobalServer_Inner_pb.CMD_SC_REGISTER()
    receiveMsg:ParseFromString(data)

    -- 处理MD5解密
    local str = string.format("InnerServerType%d", Bit32LShift(16,receiveMsg.server_type))
    local md5Str = string.upper(MD5Get(str))
    if md5Str ~= receiveMsg.server_token then
        print(string.format("md5Str = %s,resMsg.server_token == %s",md5Str,receiveMsg.server_token));
        print("登录服发送的 聊天服token出错")
        return false
    end

    -- 添加登录服连接信息
    self:AddLoginServerLink(serverId)

    -- 返回登录服消息响应
    local sendMsg =  CMD_GlobalServer_Inner_pb.CMD_SS_REGISTER()
    sendMsg.result = CMD_GlobalServer_pb.REPLYRESULT_SUCCESSFUL_ENUM.number
    LuaNetWorkSend( serverId, MAIN_CHAT_SERVICE_INNER, SUB_SS_REGISTER , sendMsg,nil)
end

--------------------------------------------------------------------------------------
--- 登录服推送玩家信息
--- @param serverId 服务器id
--- @param data     数据
--- @param token
function ChatTable:HandleLoginServerPushToken(serverId, data, token)
    -- 解析消息
    local receiveMsg =  CMD_GlobalServer_Inner_pb.CMD_SC_PUSH_TOKEN()
    receiveMsg:ParseFromString(data)

    local tMsgUserInfo = receiveMsg.user_info or {}
    if tMsgUserInfo.user_id == nil or tMsgUserInfo.user_id == 0 then
        return
    end

    -- 添加玩家Token到列表
    self:AddPlayerToken(tMsgUserInfo.user_id, receiveMsg.token)
end

--------------------------------------------------------------------------------------
--- 登录服推送玩家相关数据信息变更(头像和VIP等级)
--- @param serverId  服务器id
--- @param data      数据
--- @param token
function ChatTable:HandleLoginServerUpdateUserInfo(serverId, data, token)
    -- 解析消息
    local resMsg =  CMD_GlobalServer_Inner_pb.CMD_SC_TOKEN_USER_INFO()
    resMsg:ParseFromString(data)
    if resMsg == nil or resMsg.userid == nil then
        return false
    end
    -- 找到player
    local player = GetPlayerByUID(resMsg.user_id)
    -- 更改玩家相关信息
    local nInfoType = resMsg.info
    if nInfoType == nil or nInfoType == Enum_UserInfoType.UIT_FACEID then
        player.User.FaceID = resMsg.date
    elseif nInfoType == nil or nInfoType == Enum_UserInfoType.UIT_VIPLV then
        player.User.VipLev = resMsg.date
    end
    return true
end