---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/12/7 15:25
---


------------------------------保存服务器状态信息，是按照时间保存的，可以看历史记录-----------------------------
--function RedisSaveServerState(state)
--    local day = GetOsDayNow()
--    local time= GetOsDateTimeNow()
--    RedisSaveString(RedisDirServerState..ServerIP_Port..":"..day..":"..time, time, ZJson.encode(state))
--end
------------------------------保存桌子状态信息，当前的运行信息，桌子销毁就删掉-----------------------------
--function RedisSaveGameState(gameType,tableId, state)
--    RedisSaveString(RedisDirGameState..ServerIP_Port..":GameID_"..gameType..":TableId"..tableId, tableId, ZJson.encode(state))
--end
------------------------------删掉桌子状态信息-----------------------------
--function RedisDelGameState(gameType,tableId)         -- 清理掉桌子的运行状态
--    RedisDelKey(RedisDirGameState..ServerIP_Port..":GameID_"..gameType..":TableId"..tableId ,tableId)
--end




----------------------------获取UUID，只能用于服务器初始化-----------------------------
function RedisGetAllPlayersUUID()         -- 清理掉桌子的运行状态
    return RedisGetString(RedisDirAllPlayersUUID.."BY_UUID" ,"BY_UUID")
end
----------------------------设置UUID，只能用于服务器初始化-----------------------------
function RedisSaveAllPlayersUUID(num)         -- 清理掉桌子的运行状态
    RedisSaveString(RedisDirAllPlayersUUID.."BY_UUID" ,"BY_UUID",num)
end
----------------------------多进程申请UUID信息， 会执行脚本先增加，然后把最新的数字返回-----------------------------
function RedisMultiProcessGetAllPlayersUUID(num)
    return RedisAddNumber(RedisDirAllPlayersUUID.."BY_UUID" ,"BY_UUID",num)
end


