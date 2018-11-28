---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/10/12 14:56
---


require("gameFire")
require("gameLogin")

----------------------------------------------------------------------
---发送消息
----------------------------------------------------------------------
-- 玩家自己的网络发送函数
function LuaNetWorkSend(msgId,subMsgId,sendCmd,err)
    return LuaNetWorkSendToUser(0,msgId,subMsgId,sendCmd,err)      -- userId 如果是0的话， 就是给玩家自己回消息 ，这是在go那边定义的
end

-- 发送消息给其他玩家
function LuaNetWorkSendToUser(userId,msgId,subMsgId,sendCmd,err)
    local buffer = ""
    if sendCmd ~= nil then
        buffer = sendCmd:SerializeToString()
    end

    if err == nil then
        err = ""
    end
    return luaCallGoNetWorkSend(userId,msgId,subMsgId,buffer,err)       -- 返回结果 true 发送成功  false 发送失败
end


----------------------------------------------------------------------
---接收消息
----------------------------------------------------------------------
-- 网络接收函数
function GoCallLuaNetWorkReceive(msgId,subMsgId,data)
    --Logger("lua收到了消息："..msgId)
    --Logger("lua收到了消息："..subMsgId)
    --Logger("lua收到了消息："..data)
    ReceiveMsg(msgId,subMsgId,data)

--    LuaNetWorkSend(msgId,subMsgId,"lua想发送消息", "")
end


-- 根据命令进行分支处理
function ReceiveMsg(msgId,subMsgId,data)
    --print("msgId",msgId, "subMsgId",subMsgId)

    if msgId == MDM_MB_LOGON  then
        if subMsgId == SUB_MB_GUESTLOGIN  then
            print("**************游客登录服申请******************* ")
        end
    elseif msgId == MDM_GR_LOGON  then
        if subMsgId == SUB_GR_LOGON_USERID  then
            print("**************游客登录游戏服申请******************* ")     ----这里是原来的登录， 主要是返回客户端玩家的一些数据
            SevLoginGSGuest(data)

        end
    elseif msgId == MDM_GF_FRAME  then
        if subMsgId == SUB_GF_GAME_OPTION  then
            print("**************游游客进入大厅申请***************** ")      ---- 这里是玩家申请登录游戏的类型，进入游戏房间， 分配桌子坐下开始玩 , 客户端需要申请房间的类型
            SevEnterScence(data)
        end
    elseif msgId == MDM_GF_GAME  then
        if subMsgId == SUB_C_USER_FIRE  then
--            print("**************客户端开火***************** ")
            HandleUserFire(data)
        elseif subMsgId == SUB_C_CATCH_FISH  then
--            print("*************客户端抓鱼***************** ")
            HandleCatchFish(data)
        end
    end
end