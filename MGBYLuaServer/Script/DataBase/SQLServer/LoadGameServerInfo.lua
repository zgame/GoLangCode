---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.zhengxh
--- DateTime: 2019/11/4 13:21
---

------------------------------------------------------------------------------------------
--- 处理加载游戏相关的信息(迎合C++版本需求，需要加入该游戏房间信息加载)
------------------------------------------------------------------------------------------

function LoadGameServerInfoFromSQLServer()
    local gameRoomInfo = {}
    -- 获取GameServerID
    local gameServerID = GameRoomServerID
    local sSql = string.format("Select ServerID,NodeID,SortID,GameID,MinEnterScore,MaxEnterScore,MinEnterVip,MaxEnterVip,"..
            "MinEnterCannonLev,MaxEnterCannonLev,ServerType,ServerRule FROM GameRoomInfo Where ServerID = %d;", gameServerID)
    local ret = SqlServerDataBaseBYQuery(sSql)
    if #ret == 0 then
        -- 没有找到该游戏数据
        Logger("LoadGameServerInfoFromSQLServer GameRoomInfo 返回 nil")
        return gameRoomInfo
    end
    gameRoomInfo = ret[1]
    return gameRoomInfo
end