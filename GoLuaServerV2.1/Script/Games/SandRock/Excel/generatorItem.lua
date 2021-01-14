-- 道具生成组
SandRockGeneratorItem = {}


-- 处理一下几率， 处理完之后，计算的时候方便
local function _setRateList(list)
    local total = 0
    for i,v in ipairs(list) do      -- 计算总和
        total = total + v
    end

    local newList = {}
    for i,v in ipairs(list) do
        local rate =  math.floor( v * 100 / total )         -- 计算几率
        table.insert(newList,rate)
    end

    for i=2, #newList do
        if newList[i] ~= nil then
            newList[i] = newList[i] + newList[i-1]              -- 把几率进行累加
        end
        if i==#newList then
            newList[i] = 100
        end
    end

    return newList
end


-- 初始化道具生成组的数据
local function _setGroup(groupId)
    SandRockGeneratorItem[groupId] = {}
    local generator = SandRockGeneratorItem[groupId]

    local GeneratorList = CSV_generateGroup.GetValue(groupId, "GeneratorList")
    local allDropList = ZString.Split(GeneratorList, "|")
    for i, mainDrop in ipairs(allDropList) do
        local sub = {}
        sub.rateList = {}
        local rateList = {}
        local subList = ZString.Split(mainDrop, ";")
        for i, subDrop in ipairs(subList) do
            local element = {}
            local subStr = ZString.Split(subDrop, ",")
            element.id = tonumber(subStr[1])
            --element.rate = tonumber(subStr[2])
            element.lucky = tonumber(subStr[3])                                         -- 幸运值以后再处理
            table.insert(sub, element)                  -- 把每个元素添加进去

            local rate = tonumber(subStr[2])
            table.insert(rateList,rate)                 -- 把几率单独弄成一个list
        end

        sub.rateList = _setRateList(rateList)
        table.insert(generator, sub)
    end

    --printTable(generator)
end

function SandRockGeneratorItem.Init()
    for  groupId,_ in pairs(CSV_generateGroup.Get()) do
        _setGroup(groupId)
    end
    --_setGroup("20900008")
end

-- 根据生成规则获取道具的数量
local function _getGroupItemNum(groupId)
    local GenDistType = CSV_generateItem.GetValue(groupId,"GenDistType")
    local GenDistParams = CSV_generateItem.GetValue(groupId,"GenDistParams")
    local GenDistParams2 = CSV_generateItem.GetValue(groupId,"GenDistParams2")
    if GenDistType == "Num" then        -- 固定数值
        local num = tonumber(GenDistParams)
        local a , b = math.modf(num);
        if ZRandom.PercentRate(b * 100) then
            a = a + 1           -- 如果小数点后面的几率达到了， 那么就增加一个
        end
        return a
    elseif  GenDistType == "Uniform" then      -- 随机区间
        return ZRandom.GetRandom(GenDistParams,GenDistParams2)
    elseif  GenDistType == "UniformFloat" then      -- 随机区间
        return ZRandom.GetFloat(GenDistParams,GenDistParams2,3)
    elseif  GenDistType == "Normal" then      -- 正态分布区间
        local normal = ZRandom.Normal()
        if normal > 1 or normal < -1 then       -- 如果落到了一倍方差之外，那么进行均匀的概率分布，策划设计如此
            normal = math.random()*2-1
        end
        return GenDistParams + GenDistParams2 * normal
    end
end


--  道具掉落 ，正常返回 hash  key是itemId， value是数量
function SandRockGeneratorItem.GetItems(groupId)
    local groupList = {}

    local GenSceneType = CSV_generateGroup.GetValue(groupId, "GenSceneType")
    if GenSceneType == "Item" then
        -- 走道具掉落规则
        local generator = SandRockGeneratorItem[tostring(groupId)]
        for index, allType in pairs(generator) do
            local subIndex = 1                    -- 如果只有一个元素，那么就是这个
            if #generator[index] > 1 then
                subIndex = ZRandom.GetList(generator[index].rateList)       -- 多个元素就随机一个
            end

            if groupList[generator[index][subIndex].id] == nil then
                groupList[generator[index][subIndex].id] = 1                 -- 如果没有，那么发一个
            else
                groupList[generator[index][subIndex].id] = groupList[generator[index][subIndex].id] + 1       -- 如果已经有了，那么数量增加
            end
        end
    end

    local itemList = {}
    for groupId,num in pairs(groupList) do
        groupId = tostring(groupId)
        local itemId = CSV_generateItem.GetValue(groupId, "GenObjectId")
        local itemNum  = _getGroupItemNum(groupId)
        itemList[itemId] = num * itemNum

        if num * itemNum == 0 then
            itemList[itemId] = nil   -- 这是因为表格里面配置有数量为0， 不掉， 那配置它干啥呢， 理解不了，闲的蛋疼 ，害的我代码都乱了
        end
    end


    --printTable(itemList)
    return itemList
    --return nil
end