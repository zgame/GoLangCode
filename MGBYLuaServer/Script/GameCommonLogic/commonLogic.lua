---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/10/12 14:57
---

---------------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------------------------
--- go 来创建和调用的通用代码处理模块， 用来处理世界boss，排行榜，各种公开活动
---------------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------------------------


--公共逻辑循环处理
function GoCallLuaCommonLogicRun()
--    print("公共逻辑循环处理")
    ShowAllGameStates()

end

-- 每60秒记录一下，服务器的状态到数据库中
function GoCallLuaSaveServerState()
    local allGamesTablesNum = 0
    for _, game in pairs(AllGamesList) do
        allGamesTablesNum = allGamesTablesNum + game.AllTableListNumber
    end
    local state = {}
    state["TableNum"] = allGamesTablesNum
    state["PlayerNum"] = AllPlayerListNumber
    state["SendNum"] = ServerStateSendNum
    state["ReceiveNum"] = ServerStateReceiveNum
    state["WriteChannelNum"] = ServerSendWriteChannelNum
    state["HeadErrorNum"] = ServerDataHeadErrorNum
    state["HeapInUse"] = ServerHeapInUse
    state["NetWorkDelay"] = ServerNetWorkDelay
    --RedisSaveServerState(state)

    -- 把记录数据保存到数据库中
    SqlSaveServerState(state)

end


--夜里12点触发公共逻辑变动，因为新的一天开始了
function GoCallLuaCommonLogic12clock()
    Logger("夜里12点触发公共逻辑变动，因为新的一天开始了")
    -- 记录活动期间的排行榜结果
    -- 发放各种奖励
    -- 各种活动的新一天的初始化
    -- 联盟工会的新一天的初始化
    --

end


-- 自己设定的新的计时器
function NewTimerBy2Second()
    --print("这是lua自己设定的定时器")
end

-- 自己设定的固定时间计时器
function NewTimerByAfternoon4()
    print("这是lua自己设定的下午4点的定时器")
end

