---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by zhushiwei.
--- DateTime: 2020/12/8 17:02
---

Main ={}

-----------------------------------服务器启动-------------------------------------
-- 服务器开始创建各个游戏，这里的游戏都是多人的游戏， 如果是单人游戏，玩家自己创建即可
function Main.GoCallLuaStartAllServers()
    -- 初始化
    print("-------------------  启动 mongo db   ---------------------------")
    GlobalVar.MongoMainConnect = MongoDB.new()
    local ok, err = MongoDB.Init(GlobalVar.MongoMainConnect, Setting.MongoAddress,  Setting.MongoDatabase, Setting.MongoUser, Setting.MongoPass)
    if ok == false then
        print("mongo 服务器启动错误: " .. err)
        return
    end
    print("-------------------  启动 redis      ---------------------------")
    GlobalVar.RedisConnect = Redis.new()
    ok, err = Redis.Init(GlobalVar.RedisConnect,Setting.RedisAddress,Setting.RedisPass)
    if ok == false then
        print("redis 服务器启动错误: " .. err)
        return
    end
    print("-------------------  启动 mySql      ---------------------------")
    GlobalVar.MySqlMainConnect = MySql.new()
    ok, err = MySql.Init(GlobalVar.MySqlMainConnect,Setting.MySqlServerIP, Setting.MySqlServerPort, Setting.MySqlDatabase, Setting.MySqlUid, Setting.MySqlPwd)
    if ok == false then
        print("mySql 服务器启动错误: " .. err)
        return
    end

    print("------------------   服务器启动初始化  ---------------------------")
    print("ServerTypeName:".. GlobalVar.ServerTypeName)

    --CreateAllGoRoutineGameTable()     --创建房间使用的goroutine函数列表
    --GetALLUserUUID()                  -- 是一个UUID是不是需要初始化的判断
    --if ServerTypeName1 == "Game1" then
    --    startGamesServers()                     -- 启动游戏服
    ----elseif ServerTypeName1 == "MainCenter" then
    --else
    --    MainCenterServer.Start()                 -- 启动主中心服
    --end
    local switch={}
    switch["Game"] = GameServer.Start               -- 启动游戏服
    switch["MainCenter"] = CenterServer.Start                 -- 启动主中心服
    -- 运行对应server type的函数
    switch[GlobalVar.ServerTypeName]()
end

