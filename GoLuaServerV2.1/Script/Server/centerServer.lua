---
--- 主中心服务器
---

local AllServerList = {}   --- key 是ip+port  value是游戏类型名字
local AllPlayersList = {}   --- 全部所有玩家列表   key  userId(type string)    value   ip+port
local AllPlayersListNumber = 0   -- 所有玩家的人数，记住不要遍历上面的map，性能太慢

CenterServer = {}
-- 启动主中心服务器
function CenterServer.Start()
    print("-------------------  启动主中心服  ------------------------------")
    -- 维护玩家总列表

    -- 维护游戏服务器列表

    -- 创建定时器处理公共数据
    ZTimer.SetNewTimer("MainCenterServer", "TimerMainCenter", 2 * 1000, CenterServer.TimerMainCenter)      -- lua 自己设定计时器
    ZTimer.SetNewClockTimer("MainCenterServer", "ClockMainCenter12", 0, CenterServer.ClockMainCenter12)    -- lua 自己设定的固定时间定时器，  24:00
end


-- 自己设定的新的计时器
function CenterServer.TimerMainCenter()
    print("这是lua自己设定的定时器")
end


-- 自己设定的固定时间计时器
function CenterServer.ClockMainCenter12()
    print("这是lua自己设定的夜里12点的定时器")
end




------------------------ 网络 ----------------------------

-- 连接成功
function CenterServer.Init(serverId)

end

-- 接受网络消息
local function Login(serverId, data)
    print("**************其他服务器注册申请******************* ")

end


-- 接受网络消息
function CenterServer.Receive(serverId, userId, msgId, subMsgId, data, token)
    if msgId == CMD_MAIN.MDM_SERVER then
        if subMsgId == CMD_SERVER.SUB_LOGON then
            Login(serverId, data)      -- 返回给客户端，玩家的数据，用来显示的
        end
    end
end

-- 链接我的连接中断
function CenterServer.Broken(uid, serverId)

end