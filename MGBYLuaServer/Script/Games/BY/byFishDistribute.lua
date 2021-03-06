---
--- Generated by EmmyLua(https:--github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/5 14:23
---

FishDistribute = {}
function FishDistribute:New()
    c = {
        FishKindID = 0, -- 鱼的类型
        CreateTime = 0, --创建时间
        DistributeIntervalTime = 0, -- 下次生成鱼的时间是一个间隔，保存这个随机间隔

        -- 同时生成多条鱼，并且路径相同，路上有时间差
        BuildNumber = 0, -- 生成鱼的数量
        FirstPathID = 0, -- 鱼的路径
        CreateType = 0, -- 生成下条鱼的方式， 1表示一条路径，不做位置偏移， 2表示要做位置偏移

        NextCreateTime = 0, -- 下条鱼的时间起始时间
        NextInterBuildTime = 0, -- 下条鱼的生成时间间隔
    }
    setmetatable(c, self)
    self.__index = self
    return c
end

function FishDistribute:Reload(c)
    setmetatable(c, self)
    self.__index = self

    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end

----获取生成时间间隔
function FishDistribute:GetIntervalTime(kindId)
    kindId  = tostring(kindId)          -- 这里要注意，之所以用string是因为用int，消耗内存大
    -- 获取时间间隔
    return GetRandom(GetExcelFishValue(kindId,"distribute_interval_min"),GetExcelFishValue(kindId,"distribute_interval_max"))
end
----获取生成数量间隔
function FishDistribute:GetCount(kindId)
    kindId  = tostring(kindId)          -- 这里要注意，之所以用string是因为用int，消耗内存大
    -- 获取时间间隔
    return GetRandom(GetExcelFishValue(kindId,"count_min"),GetExcelFishValue(kindId,"count_max"))
end
----多条鱼生成时间间隔
function FishDistribute:GetCountFishTime(kindId)
    kindId  = tostring(kindId)          -- 这里要注意，之所以用string是因为用int，消耗内存大
    -- 获取时间间隔
    return GetRandom(GetExcelFishValue(kindId,"time_min"),GetExcelFishValue(kindId,"time_max"))
end

----获取路径的类型
function FishDistribute:GetPathType()
    return 1
end


-- 切记， 像这样的常量数组， 一定要定义成全局，  不要local然后返回， 会分配内存的，非常消耗性能
FishDistributeOffsetXY = {{0,0}, {-1,1}, {-0.5,1},{0,1},{0.5,1},{1,1},
                          {-1,0.5}, {-0.5,0.5},{0,0.5},{0.5,0.5},{1,0.5},{-1,1},{1,0},
                          {-1,-0.5}, {-0.5,-0.5},{0,-0.5},{0.5,-0.5},{1,-0.5},
                          {-1,-1}, {-0.5,-1},{0,-1},{0.5,-1},{1,-1}}
----获得路径位置偏移
function FishDistribute:GetOffsetXY()
    return FishDistributeOffsetXY[GetRandom(1,23)]
end


