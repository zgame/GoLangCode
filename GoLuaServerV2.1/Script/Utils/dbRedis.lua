---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/28 14:27
---


---------------------------------------------------------
---redis  原则上禁止多个数据库， 如果有特殊情况，仿造mongo db写法即可实现多数据库连接
---------------------------------------------------------

RedisEngine = require('redis')
--RedisEngineConnect = RedisEngine.new()

--- 创建数据库句柄
function RedisNew()
    return RedisEngine.new()
end

----------------------------通用函数-----------------------------
-- 有些命令会返回string 或者int
local function cmdForRedis(cmd, handle, ...)
    if handle == nil then
        handle = RedisEngineConnect
    end
    local arg = { ... }
    local slice = {}
    for i, v in ipairs(arg) do
        slice[i] = v
    end
    local string, number = handle:cmd(cmd, slice)
    return string, number
end

-- 有些命令会返回table
local function getStringListFromRedis(cmd, handle, ...)
    if handle == nil then
        handle = RedisEngineConnect
    end
    local arg = { ... }
    local slice = {}
    for i, v in ipairs(arg) do
        slice[i] = v
    end
    local list = handle:stringList(cmd, slice)
    return list
end

----------------------------初始化-----------------------------
local function test()
    --print("-------string and hash---------")
    ----print(RedisExistKey("test:name"))
    --print(RedisSaveString("test:keyt",nil,"ss"))
    --print(RedisSaveString("test:hash","name","ss1"))
    ----print(RedisExistKey("test:hash","name"))
    --
    --print(RedisGetString("test:keyt"))
    --print(RedisGetString("test:hash","name"))
    --
    --print("-------list---------")
    --print(RedisListPush("test:list200",'ss'))
    --print(RedisListPush("test:list200",20))
    --print(RedisListPush("test:list200",25))
    --print(RedisListLen("test:list200"))
    --print(RedisListIndex("test:list200",0))
    --print(RedisListIndex("test:list200",1))
    --print(RedisListIndex("test:list200",2))
    --print(RedisListIndex("test:list200",3))
    --
    --print(RedisListRem("test:list200","ss",1))
    --printTable(RedisListGetAll("test:list200"))
    --print("-------set---------")
    --print(RedisSetPush("test:set200",'ss'))
    --print(RedisSetPush("test:set200",20))
    --print(RedisSetPush("test:set200",25))
    --print(RedisSetPush("test:set200",25))
    --print(RedisSetLen("test:set200"))
    --
    --printTable(RedisSetGetAll("test:set200"))
    --print(RedisSetRem("test:set200","ss"))
    --print(RedisSetMember("test:set200","20"))
    --print("-------sort set---------")
    --print(RedisSortedSetPush("test:sset200", 'ss', 100))
    --print(RedisSortedSetPush("test:sset200", 20, 200))
    --print(RedisSortedSetPush("test:sset200", 25, 300))
    --print(RedisSortedSetPush("test:sset200", 25, 500))
    --print(RedisSortedSetPush("test:sset200", 35, 50))
    --print(RedisSortedSetLen("test:sset200"))
    --
    --printTable(RedisSortedSetGetAll("test:sset200"))
    --printTable(RedisSortedSetGetEnd("test:sset200"))
    ----print(RedisSortedSetRem("test:sset200","ss"))
    ----print(RedisSortedSetRank("test:sset200","20"))
    ----print(RedisSortedSetRank("test:sset200","25"))
    ----print(RedisSortedSetRank("test:sset200","ss"))
    --print(RedisSortedSetRemEnd("test:sset200"))
    --printTable(RedisSortedSetGetAll("test:sset200"))
end
function RedisInit(handle, RedisAddress, RedisPass)
    local ok, err = handle:connect({ host = RedisAddress, password = RedisPass })
    if ok then
        print(" redis  数据库 ok!")
        --test()
    end
    return ok, err
end

----------------------------key是否存在判定-----------------------------
function RedisExistKey(dir, hashKey)
    local number
    if hashKey == nil then
        _, number = cmdForRedis("exists", nil, dir)        -- 判断key
    else
        _, number = cmdForRedis("hexists", nil, dir, hashKey)        -- 判断hash里面的key
    end
    return number   -- 存在返回 1
end


----------------------------string  and  hash---------------------------------------
--保存
function RedisSaveString(dir, hashKey, value)
    local string, number
    if hashKey == nil then
        string = cmdForRedis("set", nil, dir, value)   -- 成功返回OK
    else
        string, number = cmdForRedis("hset", nil, dir, hashKey, value)      -- 新增返回1， 更新返回0
        if number == 1 or number == 0 then
            string = "OK"
        end
    end
    return string       -- 成功返回OK
end
--获取
function RedisGetString(dir, hashKey)
    local string
    if hashKey == nil then
        string = cmdForRedis("get", nil, dir)
    else
        string = cmdForRedis("hget", nil, dir, hashKey)
    end
    return string
end
--删除
function RedisDelKey(dir, hashKey)
    if hashKey == nil then
        cmdForRedis("del", nil, dir)
    else
        cmdForRedis("hdel", nil, dir, hashKey)
    end
end

----------------------------script-----------------------------
--分布式的数字协调，避免加分布式锁， 可广泛用于多个进程高并发同时处理同一个数据， 保证数据的原子性----------------------------
function RedisRunLuaScript(redis_lua_str, name, handle)
    -- name 是用来报错用的，定位问题用
    if handle == nil then
        handle = RedisEngineConnect
    end
    local number, err = RedisEngineConnect:script(redis_lua_str, name)
    if err ~= nil then
        Logger(err)
    end
    return number       -- 一般情况都是针对数字的处理，返回结果值
end


----------------------------list-----------------------------
-- 添加内容
function RedisListPush(dir, value)
    local _, num = cmdForRedis("rpush", nil, dir, value)
    return num      -- 返回list长度
end
-- 获取长度
function RedisListLen(dir)
    local _, num = cmdForRedis("llen", nil, dir)
    return num      -- 返回list长度
end
-- 获取固定位置的元素
function RedisListIndex(dir, index)
    -- 第一个从0开始
    local str = cmdForRedis("lindex", nil, dir, index)
    return str      -- 返回 内容 都是字符串，如果是数字自己转换
end
-- 获取所有元素
function RedisListGetAll(dir, listStart, listEnd)
    if listStart == nil then
        listStart = 0
    end
    if listEnd == nil then
        listEnd = -1
    end
    local list = getStringListFromRedis("lrange", nil, dir, listStart, listEnd)
    return list
end
-- 移除跟内容一样的元素， count > 0 从头计数 ， < 0 从尾部计数 , 0 是所有等于 ,  输入的value是数字还是字符串都会匹配删除
function RedisListRem(dir, value, count)
    if count == nil then
        count = 0
    end
    local _, num = cmdForRedis("lrem", nil, dir, count, value)
    return num
end


----------------------------set-----------------------------
-- 添加内容
function RedisSetPush(dir, value)
    local _, num = cmdForRedis("sadd", nil, dir, value)
    return num      -- 返回set长度
end
-- 获取长度
function RedisSetLen(dir)
    local _, num = cmdForRedis("scard", nil, dir)
    return num      -- 返回set长度
end

-- 获取所有元素
function RedisSetGetAll(dir)
    local set = getStringListFromRedis("smembers", nil, dir)
    return set
end
-- 移除元素
function RedisSetRem(dir, value)
    local _, num = cmdForRedis("srem", nil, dir, value)
    return num
end
-- 是否是成员
function RedisSetMember(dir, value)
    local _, num = cmdForRedis("sismember", nil, dir, value)
    return num
end


----------------------------sorted set-----------------------------
-- 添加内容
function RedisSortedSetPush(dir, value, score)
    local _, num = cmdForRedis("zadd", nil, dir, score, value)
    return num      -- 返回set长度
end
-- 获取长度
function RedisSortedSetLen(dir)
    local _, num = cmdForRedis("zcard", nil, dir)
    return num      -- 返回set长度
end

-- 获取所有元素
function RedisSortedSetGetAll(dir)
    local set = getStringListFromRedis("zrevrange", nil, dir, 0, -1, "withscores")   -- 按照从高到低排序，带分数
    return set
end
-- 获取排名最后的人和分数
function RedisSortedSetGetEnd(dir)
    local set = getStringListFromRedis("zrevrange", nil, dir, -1, -1, "withscores")   -- 分数最低的人
    return set
end
-- 移除排名最低的元素
function RedisSortedSetRemEnd(dir)
    local _, num = cmdForRedis("zremRangeByRank", nil, dir, 0, 0)
    return num
end
-- 获取成员排名
function RedisSortedSetRank(dir, value)
    local _, num = cmdForRedis("zrevrank", nil, dir, value)       -- 从高到低， 排行位次
    return num   -- 排名第一是0
end


--/*
--//----------------------------------------------------------------------------------------
--// test  这是测试用的例子 ， 可以作为参考
--//----------------------------------------------------------------------------------------
--
--fmt.Println("//-------------------------------------------- string --------------------------------------------------")
--CmdForRedis("set" ,"zsw", "zsw_value1")			// 成功返回  "OK"
--CmdForRedis("get" ,"zsw")							// 存在返回 (value,0)   如果不存在返回("",0)
--CmdForRedis("setnx" ,"zsw", "zsw_value11") 		// 如果不存在，那么set   成功返回1 ，失败返回0
--fmt.Println("//-------------------------------------------- key --------------------------------------------------")
--CmdForRedis("exists" ,"zsw")						// 存在返回1 ，不存在返回0
--CmdForRedis("setex" ,"zsw", "50","zsw_value11") 	// set  value ，并且设定过期时间 *秒   成功返回 "OK"
--CmdForRedis("expire" ,"zsw", "100")              	// 设定过期时间   成功返回1 ，失败返回0
--CmdForRedis("expireat" ,"zsw", "2293840000")   	// 设定过期时间 在某个系统时间戳定义的时间过期    成功返回1 ，失败返回0
--CmdForRedis("ttl" ,"zsw")                          // 返回过期时间
--CmdForRedis("del" ,"zsw")							// 返回删除key的数量
--fmt.Println("//-------------------------------------------- hash --------------------------------------------------")
--CmdForRedis("hset" ,"zsw_hash", "zsw_key1","zsw_value11")	// 新增返回1 ，更新返回0
--CmdForRedis("hget" ,"zsw_hash", "zsw_key1")		// 成功返回value ，失败返回 ""
--CmdForRedis("hexists" ,"zsw_hash", "zsw_key1")		// 存在返回1 ，不存在返回0
--CmdForRedis("hdel" ,"zsw_hash", "zsw_key1")		// 成功返回1 ，失败返回0
--fmt.Println("//-------------------------------------------- list --------------------------------------------------")
--CmdForRedis("lpush" ,"zsw_list", "zsw_key1")                          // 从头部插入 返回list 长度
--CmdForRedis("rpush" ,"zsw_list", "zsw_key1")                          // 从尾部插入 返回list 长度
--CmdForRedis("llen" ,"zsw_list")                                       // 返回list 长度
--CmdForRedis("lindex" ,"zsw_list",0)                                   // 获取 index 元素
--CmdForRedis("lset" ,"zsw_list",0, "zsw_key2")                         // 成功返回  "OK"
--GetStringListFromRedis("lrange", "zsw_list", 0, -1)				   // 获取所有元素
--CmdForRedis("lrem" ,"zsw_list",0 ,"zsw_key1")                         // 移除 一定数量的 value , > 0 从头计数 ， < 0 从尾部计数 , 0 是所有等于 , 返回值是移除元素数量
--CmdForRedis("rpop" ,"zsw_list")                                       // 删尾部一个，返回删除的值， 不存在的话 ( "",0)
--CmdForRedis("lpop" ,"zsw_list")                                       // 删头部一个，返回删除的值， 不存在的话 ( "",0)
--fmt.Println("//-------------------------------------------- set --------------------------------------------------")
--CmdForRedis("sadd" ,"zsw_set","zsw_key1")			// 添加成员 , 返回成员数量
--CmdForRedis("scard" ,"zsw_set")					// 成员数量
--CmdForRedis("sismember" ,"zsw_set", "zsw_key1")	// 是否是成员 存在返回1 ，不存在返回0
--GetStringListFromRedis("smembers", "zsw_set")		// 获取所有元素
--CmdForRedis("srem" ,"zsw_set", "zsw_key1")			// 移除
--fmt.Println("//-------------------------------------------- sorted set --------------------------------------------------")
--CmdForRedis("zadd" ,"zsw_zset", 100, "zsw_key1")			// 添加成员 , 返回成功添加成员数量
--CmdForRedis("zadd" ,"zsw_zset", 1100, "zsw_key2")			// 添加成员 , 返回成功添加成员数量
--CmdForRedis("zadd" ,"zsw_zset", 3100, "zsw_key3")			// 添加成员 , 返回成功添加成员数量
--CmdForRedis("zadd" ,"zsw_zset", 4100, "zsw_key4")			// 添加成员 , 返回成功添加成员数量
--CmdForRedis("zadd" ,"zsw_zset", 110, "zsw_key1")			// 更新成员 , 返回成功添加成员数量 , 更新返回 0
--CmdForRedis("zscore" ,"zsw_zset", "zsw_key1")				// 返回成员分数
--CmdForRedis("zcount" ,"zsw_zset", 0, 1000)					// 返回成员数量在分数范围内
--CmdForRedis("zcard" ,"zsw_zset")							// 成员数量
--CmdForRedis("zrevrank", "zsw_zset","zsw_key1")				// 获取成员排名， 从大到小, 排名第一是0
--CmdForRedis("zrank", "zsw_zset","zsw_key1")					// 获取成员排名， 从小到大,排名第一是0
--GetStringListFromRedis("zrange", "zsw_zset",0,1, "withscores")	// 获取 all 成员排名 从小到大   0,-1 是所有  WITHSCORES则同时返回对应的分数
--GetStringListFromRedis("zrange", "zsw_zset",0,-1)				// 获取 all 成员排名 从小到大   0,-1 是所有
--GetStringListFromRedis("zrevrange", "zsw_zset",0,2, "withscores")// 获取 all 成员排名 从大到小   0,-1 是所有
--GetStringListFromRedis("zrangebyscore", "zsw_zset",0,3000)		// 获取  成员区间排名	score
--GetStringListFromRedis("zrevrangebyscore", "zsw_zset",3000,0)		// 获取  成员区间排名 score
--GetStringListFromRedis("zrangebylex", "zsw_zset","(zsw_key1","[zsw_key4")// 获取  成员区间排名
--GetStringListFromRedis("zrevrangebylex", "zsw_zset","(zsw_key1","[zsw_key4")// 获取  成员区间排名
--CmdForRedis("zremRangeByLex", "zsw_zset","[zsw_key1","[zsw_key2")		// 移除按成员区间
--CmdForRedis("zremRangeByRank", "zsw_zset",0,0)				// 移除按排名 按排名， 返回移除成员数量
--CmdForRedis("zremRangeByScore", "zsw_zset",0,10000)			// 移除按排名 按分数， 返回移除成员数量
--
--//return false		// test open
--*/
