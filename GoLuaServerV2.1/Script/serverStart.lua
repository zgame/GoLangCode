
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
    SandRockResourceGenerator.Init()
    SandRockGeneratorItem.Init()
    print("------------------   数据初始化完成  ---------------------------")
    print("ServerTypeName:".. GlobalVar.ServerTypeName)


    local switch={}
    switch[Const.ServerGame] = GameServer.Start               -- 启动游戏服
    switch[Const.ServerCenter] = CenterServer.Start                 -- 启动主中心服
    -- 运行对应server type的函数
    switch[GlobalVar.ServerTypeName]()


    print("ServerIP_Port:"..GlobalVar.ServerIP_Port.."用来做日志")



    -----------------------------------------------------------------



    --
    --local table = Json.Decode(str)
    --    for i,v in ipairs(table.resourceAreaConfigs) do
    --        --print(v.sceneAreaName)
    --           --print(v.minCount..","..v.maxCount)
    --            local idList={}
    --            local weightList={}
    --            for i,vv in ipairs(v.weightConfigs) do
    --                idList[i] = tostring(vv.id)
    --                weightList[i] = tostring(vv.weight)
    --            end
    --            --print(ZString.Join(idList,","))
    --            --print(ZString.Join(weightList,","))
    --            --print('--------------------')
    --            print(string.format("%s,\"%s\",\"%s\",\"%s\"",    v.sceneAreaName,  v.minCount..","..v.maxCount, ZString.Join(idList," , "),ZString.Join(weightList," , ")))
    --
    --        --end
    --    end

    --end

    --print(GetExcelItemValue(1,"GeneratorGroup"))
    --print(CSV_resourceType.GetValue(1,"GeneratorGroup"))
    --printTable(CSV_resourceType.GetAllKeys())



end

