---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.
--- DateTime: 2019/4/21 13:30
---


---------------------------------------------------------------------------
--- 库存 ，影响中奖的几率
---------------------------------------------------------------------------


---------------------------玩家中奖，库存减少-----------------------------------
function PoolReducePersonAndAll(player, count)
    local gameTable = player:GetTable()
    -- 个人库存减少
    PoolReducePerson(player, count * gameTable.RoomScore)
    -- 全局库存减少
    PoolReduceAll(gameTable, count * gameTable.RoomScore)
end









-------------------------个人库存----------------------------------

-- 个人库存值
function PoolGetPerson(player)
    return player.User.ScorePoolPerson
end


-- 个人库存值的增加
function PoolAddPerson(player,add)
    player.User.ScorePoolPerson = player.User.ScorePoolPerson + add
end

-- 个人库存值的减少
function PoolReducePerson(player,reduce)
    player.User.ScorePoolPerson = player.User.ScorePoolPerson - reduce
end


-- 获取个人的库存的几率影响
function PoolPersonRate(player)
    local gameTable = player:GetTable()
    local personPoolRate =   PoolGetPerson(player) / gameTable.RoomScore      -- 个人比率
    local listMax = 10      -- 最多多少行

    -- 先判断上下限
    if  personPoolRate > GetExcelValue(TBScorePersonPoolExcel , listMax , "max" ) then
        Logger("个人库存比率超过最大上限了"..personPoolRate)
        return GetExcelValue(TBScorePersonPoolExcel , listMax , "rate" )
    end
    if  personPoolRate < GetExcelValue(TBScorePersonPoolExcel , 1 , "min" ) then
        Logger("个人库存比率超过下限了"..personPoolRate)
        return GetExcelValue(TBScorePersonPoolExcel , 1 , "rate" )
    end

    for i=1,10 do
        local min = GetExcelValue(TBScorePersonPoolExcel , i , "min" )
        local max = GetExcelValue(TBScorePersonPoolExcel , i , "max" )

        if personPoolRate < max and personPoolRate >= min then
            -- 在这个档位里面
            local rate = GetExcelValue(TBScorePersonPoolExcel , i , "rate" )
            return rate
        end
    end
end






-------------------------公共库存----------------------------------

-- 全局库存值
function PoolGetAll(gameTable)
    return gameTable.PoolAllScore
end


-- 全局库存值的增加
function PoolAddAll(gameTable,add)
    -- 有一部分进入全局库存
    local rate = GetExcelValue(TBRoomExcel, gameTable.RoomScore, "pool")
    local value = add * rate /10000
    gameTable.PoolAllScore = gameTable.PoolAllScore + value
end

--全局库存值的减少
function PoolReduceAll(gameTable,reduce)
    gameTable.PoolAllScore = gameTable.PoolAllScore - reduce
end


-- 获取全局的库存的几率影响
function PoolAllRate(player)
    local gameTable = player:GetTable()
    local poolMax = GetExcelValue(TBRoomExcel , gameTable.RoomScore   , "poolMax" )
    local allPoolRate =   PoolGetAll(gameTable) / poolMax                                   -- 全局比率

    local listMax = 10      -- 最多多少行

    -- 先判断上下限
    if  allPoolRate > GetExcelValue(TBScoreAllPoolExcel , listMax , "max" ) then
        Logger("全局比率超过最大上限了"..allPoolRate)
        return GetExcelValue(TBScoreAllPoolExcel , listMax , "rate" )
    end
    if  allPoolRate < GetExcelValue(TBScoreAllPoolExcel , 1 , "min" ) then
        Logger("全局比率超过下限了"..allPoolRate)
        return GetExcelValue(TBScoreAllPoolExcel , 1 , "rate" )
    end



    for i=1,listMax do
        local min = GetExcelValue(TBScoreAllPoolExcel , i , "min" )
        local max = GetExcelValue(TBScoreAllPoolExcel , i , "max" )
        if allPoolRate < max and allPoolRate >= min then
            -- 在这个档位里面
            local rate = GetExcelValue(TBScoreAllPoolExcel , i , "rate" )
            return rate
        end
    end

    print("没有找到全局库存的档位........................")
end



