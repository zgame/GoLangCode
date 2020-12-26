

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
        --user2["name"] = 889
        --local ss = MongoFind("people",user)
        ----local ss = MongoFinds("people",user2,"-key")
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
    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    local result = handle:find(collection, table)
    return result
end

-- 执行finds语句  查询多条记录  sort里面是列名 负的表示从高到低排序
function MongoDB.Finds(collection, table, sort, handle)
    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    if sort == nil then
        sort = "-1"
    end

    local result = handle:finds(collection, table, sort)
    return result
end

-- 执行insert语句
function MongoDB.Insert(collection, table, handle)
    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    local err = handle:insert(collection, table)
    if err ~= nil then
        ZLog.Logger(err)
    end
    return err
end


-- 执行del语句
function MongoDB.Del(collection, table, handle)
    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    local err = handle:del(collection, table)
    if err ~= nil then
        ZLog.Logger(err)
    end
    return err
end


-- 执行update语句  selectTable为条件  updateTable为更新的内容
function MongoDB.Update(collection, selectTable, updateTable, cmd, handle )
    if handle == nil then
        handle = GlobalVar.MongoMainConnect
    end
    if cmd == nil then
        cmd = "$set"        -- 默认是更新命令
    end
    local err = handle:update(collection, selectTable, updateTable, cmd)
    if err ~= nil then
        ZLog.Logger(err)
    end
    return err
end

