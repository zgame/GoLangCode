

MongoDB = require('mongodb')

-- Mongo数据库连接
function MongoDB.Init(handle, h, d, u, ps)
    local ok, err = handle:connect({ host = h, database = d, user = u, password = ps })

    if ok then
        print(" Mongo  数据库 ok!")
        --local user = {}
        --user["ss"] = 22
        --user["name"] = 889
        --user["ss2"] = '223'
        ----print(MongoInsert("people",user))
        --
        --
        --local user2 = {}
        --user2["userId"] = 1000000004
        --local ss = MongoDB.Find("User", user2 )
        ----local ss = MongoFinds("people",user2,"-key")
        --print("dddddddddddddddddddddddd")
        --printTable(ss)
        --
        ----print(MongoUpdate("people",user,user2))
        --user2["name"] = 1
        --print(MongoUpdate("people",user,user2,"$inc"))

    end

    return ok, err

end


-- 执行find语句
function MongoDB.Find(collection, table, handle)
    --printTable(table)
    local startTime = ZTime.GetOsTimeMillisecond()

    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    local result = handle:find(collection, table)

    local endTime = ZTime.GetOsTimeMillisecond()
    if endTime-startTime > 2 then
        print("mongo db find 花费时间: " .. endTime - startTime .. "毫秒")
        printTable(table)
    end
    return result
end

-- 执行finds语句  查询多条记录  sort里面是列名 负的表示从高到低排序
function MongoDB.Finds(collection, table, sort, handle)
    local startTime = ZTime.GetOsTimeMillisecond()
    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    if sort == nil then
        sort = "-1"
    end

    local result = handle:finds(collection, table, sort)
    local endTime = ZTime.GetOsTimeMillisecond()
    print("mongo db finds cost:".. endTime - startTime)
    return result
end

-- 执行insert语句
function MongoDB.Insert(collection, table, handle)
    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    handle:insert(collection, table)
end


-- 执行del语句
function MongoDB.Del(collection, table, handle)
    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    handle:del(collection, table)
end


-- 执行update语句  selectTable为条件  updateTable为更新的内容
function MongoDB.Update(collection, selectTable, updateTable, cmd, handle )

    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    if cmd == nil then
        cmd = "$set"        -- 默认是更新命令
    end
    handle:update(collection, selectTable, updateTable, cmd)
end


-- 执行ping语句
function MongoDB.Ping(handle)
    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    return handle:ping()
end