---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/28 14:27
---

zJson = require("Json")

function RedisSaveString(dir,key,value)
    return luaCallGoRedisSaveString(dir,key,value)
end

function RedisGetString(dir,key)
    return luaCallGoRedisGetString(dir,key)
end

RedisDirAllPlayers = "AllPlayers"
RedisDirServerState = "ServerState_IP_"
RedisDirGameState = "GameState_"


RedisKeyPlayer = "UID_"
RedisKeyServerState = "Time_"
RedisKeyGameState = "Table_"

----------------------------玩家信息-----------------------------
function RedisSavePlayer(User)
    local value = zJson.encode(User)
    --print("保存玩家信息")
    --print(value)
    RedisSaveString(RedisDirAllPlayers,RedisKeyPlayer..User.UserId, value)
end

function RedisGetPlayer(uid)
    local userStr =  RedisGetString(RedisDirAllPlayers, RedisKeyPlayer..uid)
    --print("获取玩家信息",userStr)

    return  zJson.decode(userStr)
    --printTable(user)
end

----------------------------保存服务器状态信息-----------------------------
function RedisSaveServerState(state)
    local value = zJson.encode(state)
    RedisSaveString(RedisDirServerState..ServerIP_Port,RedisKeyServerState..GetOsDateNow(), value)
end
----------------------------保存桌子状态信息-----------------------------
function RedisSaveGameState(gameType,tableId, state)
    local value = zJson.encode(state)
    RedisSaveString(RedisDirGameState..gameType,RedisKeyGameState..tableId, value)
end

