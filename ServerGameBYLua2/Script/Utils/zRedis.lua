---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/28 14:27
---

local zJson = require("Json")

function RedisSaveString(dir,key,value)
    return luaCallGoRedisSaveString(dir,key,value)
end

function RedisGetString(dir,key)
    return luaCallGoRedisGetString(dir,key)
end

function RedisDelKey(dir,key)
    luaCallGoRedisDelKey(dir,key)
end

-----------------------DIR-------------------------
local RedisDirAllPlayers = "BY_AllPlayers"           -- 所有玩家列表
local RedisDirServerState = "BY_ServerState_"         -- 各个服务器状态 多少个桌子，多少玩家在线， 网络情况，1分钟记录一次，永久记忆
local RedisDirGameState = "BY_GameState_"                -- 当前各个服务器，各个游戏的状态，多少鱼，多少子弹，多少椅子有人

-----------------------KEY-------------------------
local RedisKeyPlayer = "UID_"
local RedisKeyServerState = "Time_"
local RedisKeyGameState = "Table_ID_"

----------------------------玩家信息-----------------------------
function RedisSavePlayer(User)
    RedisSaveString(RedisDirAllPlayers,RedisKeyPlayer..User.UserId, zJson.encode(User))
end

function RedisGetPlayer(uid)
    return  zJson.decode(RedisGetString(RedisDirAllPlayers, RedisKeyPlayer..uid))
end

----------------------------保存服务器状态信息-----------------------------
function RedisSaveServerState(state)
    RedisSaveString(RedisDirServerState..ServerIP_Port,RedisKeyServerState..GetOsDateNow(), zJson.encode(state))
end
----------------------------保存桌子状态信息-----------------------------
function RedisSaveGameState(gameType,tableId, state)
    RedisSaveString(RedisDirGameState..ServerIP_Port.."_ID_"..gameType,RedisKeyGameState..tableId, zJson.encode(state))
end

function RedisDelGameState(gameType,tableId)         -- 清理掉桌子的运行状态
    RedisDelKey(RedisDirGameState..ServerIP_Port.."_ID_"..gameType,RedisKeyGameState..tableId)
end


----------------------------保存桌子状态信息-----------------------------