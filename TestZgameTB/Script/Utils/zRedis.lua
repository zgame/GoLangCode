---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/28 14:27
---

--local ZJson = require("Json")
----------------------------保存-----------------------------
function RedisSaveString(dir,key,value)
    return luaCallGoRedisSaveString(dir,key,value)
end
----------------------------获取-----------------------------
function RedisGetString(dir,key)
    return luaCallGoRedisGetString(dir,key)     -- 这里返回的都是string，其他格式要自己转一下
end
----------------------------删除-----------------------------
function RedisDelKey(dir,key)
    luaCallGoRedisDelKey(dir,key)
end

----------------------------分布式的数字增加， 可广泛用于多个进程高并发同时处理同一个数据， 保证数据的原子性----------------------------
function RedisAddNumber(dir,key,num)
    return luaCallGoAddNumberToRedis(dir,key,num)       -- 返回增加之后的结果
end



