

------------------------------保存服务器状态信息，是按照时间保存的，可以看历史记录-----------------------------
--function RedisSaveServerState(state)
--    local day = GetOsDayNow()
--    local time= GetOsDateTimeNow()
--    RedisSaveString(RedisDirServerState..ServerIP_Port..":"..day..":"..time, time, ZJson.encode(state))
--end
------------------------------保存房间状态信息，当前的运行信息，房间销毁就删掉-----------------------------
--function RedisSaveGameState(gameType,roomId, state)
--    RedisSaveString(RedisDirGameState..ServerIP_Port..":GameID_"..gameType..":roomId"..roomId, roomId, ZJson.encode(state))
--end
------------------------------删掉房间状态信息-----------------------------
--function RedisDelGameState(gameType,roomId)         -- 清理掉房间的运行状态
--    RedisDelKey(RedisDirGameState..ServerIP_Port..":GameID_"..gameType..":roomId"..roomId ,roomId)
--end





