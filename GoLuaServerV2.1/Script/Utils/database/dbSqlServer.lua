------------------------------------------------------------------------
--- 这是lua直接调用sql server
------------------------------------------------------------------------


SqlServer = require('sqlServer')
--SqlServerMainEngineConnect = SqlServerEngine.new()



--- mysql数据库连接
function SqlServer.Init(handle,h,p,d,u,ps)
    local ok, err = handle:connect({ host = h, port = p, database = d, user = u, password = ps })

    if ok then
        print("sql server is ok")
        --local res, err = SqlServerEngineConnect:query('SELECT * FROM dbo.WisdomInfo')
        --if err then
        --    print(err)
        --else
        --    printTable(res)
        --end
    end

    return ok,err

end

--- 执行sql select语句
function SqlServer.Query(sql,handle)
    if handle == nil then
        handle = GlobalVar.SqlServerMainEngineConnect
    end
    local re,err = handle:query(sql)
    if err ~= nil then
        ZLog.Logger(err)
    end
    return re
end


--- 执行sql exec语句
function SqlServer.Exec(sql,handle)
    if handle == nil then
        handle = GlobalVar.SqlServerMainEngineConnect
    end
    local err = handle:exec(sql)
    if err ~= nil then
        ZLog.Logger(err)
    end
end





