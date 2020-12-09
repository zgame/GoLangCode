---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by zhushiwei.
--- DateTime: 2020/12/8 17:02
---



-----------------------------------服务器启动-------------------------------------
-- 服务器开始创建各个游戏，这里的游戏都是多人的游戏， 如果是单人游戏，玩家自己创建即可
function GoCallLuaStartAllServers()
    -- 初始化
    print("-------------------  启动 mongo db   ---------------------------")
    MongoMainEngineConnect = MongoEngine.new()
    local ok, err = MyMongoConnect(MongoMainEngineConnect,ConstMongoAddress,  ConstMongoDatabase, ConstMongoUser, ConstMongoPass)
    if ok == false then
        print("mongo 服务器启动错误: " .. err)
        return
    end
    print("-------------------  启动 redis      ---------------------------")
    RedisEngineConnect = RedisEngine.new()
    ok, err = RedisInit(RedisEngineConnect,ConstRedisAddress,ConstRedisPass)
    if ok == false then
        print("redis 服务器启动错误: " .. err)
        return
    end
    print("-------------------  启动 mySql      ---------------------------")
    MySqlMainEngineConnect = MySqlEngine.new()
    ok, err = MysqlConnect(MySqlMainEngineConnect,ConstMySqlServerIP, ConstMySqlServerPort, ConstMySqlDatabase, ConstMySqlUid, ConstMySqlPwd)
    if ok == false then
        print("mySql 服务器启动错误: " .. err)
        return
    end

    print("------------------   服务器启动初始化  ---------------------------")
    print("ServerTypeName:"..ServerTypeName)

    --CreateAllGoRoutineGameTable()     --创建桌子使用的goroutine函数列表
    --GetALLUserUUID()                  -- 是一个UUID是不是需要初始化的判断
    --if ServerTypeName == "Game1" then
    --    startGamesServers()                     -- 启动游戏服
    ----elseif ServerTypeName == "MainCenter" then
    --else
    --    MainCenterServer.Start()                 -- 启动主中心服
    --end
    local switch={}
    switch["Game"] = StartGamesServers                     -- 启动游戏服
    switch["MainCenter"] = MainCenterServer.Start                 -- 启动主中心服
    -- 运行对应server type的函数
    switch[ServerTypeName]()
end

