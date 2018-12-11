---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/12/11 11:18
---


MySqlEngine = require('mysql')
MySqlEngineConnect = MySqlEngine.new()

-- mysql数据库连接
function MysqlConnect(h,d,u,ps)
    ok, err = MySqlEngineConnect:connect({ host = h, port = 3306, database = d, user = u, password = ps })

    if ok then
        print("Mysql is ok")
        --res, err = MySqlEngineConnect:query('SELECT * FROM user')
        --printTable(res)
        end

    return ok

end

-- 执行sql select语句
function MysqlQuery(sql)
    local re,err
    re,err = MySqlEngineConnect:query(sql)
    return re
end


-- 执行sql exec语句
function MysqlExec(sql)
    MySqlEngineConnect:exec(sql)
end
