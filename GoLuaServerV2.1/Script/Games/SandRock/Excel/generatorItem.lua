-- 道具生成组
SandRockGeneratorItem = {}

--------------------------------------------初始化 道具生成组----------------------------------------------------
-- 处理一下几率变成总数是100， 每个是几率阶梯， 处理完之后，计算的时候方便
local function _setRateList(list)
    local total = 0
    for _,v in ipairs(list) do      -- 计算总和
        total = total + v
    end

    local newList = {}
    for _,v in ipairs(list) do
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
local function _setGroupInit(groupId)
    SandRockGeneratorItem[groupId] = {}
    local generator = SandRockGeneratorItem[groupId]

    local GeneratorList = CSV_generateGroup.GetValue(groupId, "GeneratorList")
    local allDropList = ZString.Split(GeneratorList, "|")
    for _, mainDrop in ipairs(allDropList) do
        local sub = {}
        sub.rateList = {}
        local rateList = {}
        local subList = ZString.Split(mainDrop, ";")
        for _, subDrop in ipairs(subList) do
            local element = {}
            local subStr = ZString.Split(subDrop, ",")
            element.generateId = subStr[1]
            --element.rate = tonumber(subStr[2])
            element.lucky = tonumber(subStr[3])                                         -- 幸运值以后再处理
            table.insert(sub, element)                  -- 把每个元素添加进去

            local rate = tonumber(subStr[2])
            table.insert(rateList,rate)                 -- 把几率单独弄成一个list
        end

        sub.rateList = _setRateList(rateList)
        table.insert(generator, sub)
    end


    --printTable(SandRockGeneratorItem[groupId])
end

function SandRockGeneratorItem.Init()
    for  groupId,_ in pairs(CSV_generateGroup.Get()) do
        _setGroupInit(groupId)
    end
    --_setGroupInit("20900008")
end


--------------------------------------------道具生成组----------------------------------------------------


-- 根据不同生成规则获取道具的数量
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
        local a = ZRandom.GetRandom(GenDistParams,GenDistParams2)
        return a
    elseif  GenDistType == "UniformFloat" then      -- 随机区间
        local a = ZRandom.GetFloat(GenDistParams,GenDistParams2,3)
        return a
    elseif  GenDistType == "Normal" then      -- 正态分布区间
        local normal = ZRandom.Normal()
        if normal > 1 or normal < -1 then       -- 如果落到了一倍方差之外，那么进行均匀的概率分布，策划设计如此
            normal = math.random()*2-1
        end
        local a = GenDistParams + GenDistParams2 * normal
        return a
    end

end


-- 道具数量
local function _setGroupNum(groupList,Index)
    --print("---------------_setGroupNum--------------------")
    if groupList[Index] == nil then
        groupList[Index] = 1                 -- 如果没有，那么发一个
    else
        groupList[Index] = groupList[Index] + 1       -- 如果已经有了，那么数量增加
    end
end

--  根据生成组，生成道具和道具的数量
--  道具掉落 ，正常返回 hash  key是itemId， value是数量
function SandRockGeneratorItem.GetItems(groupId ,scale, all)
    local groupList = {}
    if scale == nil then
        scale = 1               -- 这里要计算一下根据缩放进行的掉落放面的影响
    end

    local GenSceneType = CSV_generateGroup.GetValue(groupId, "GenSceneType")
    if GenSceneType == "Item" then
        -- 走道具掉落规则
        local generator = SandRockGeneratorItem[tostring(groupId)]                  -- 一组生成器，用|分割的是每样一个
        if generator == nil then
            ZLog.Logger("GetItems 生成组报错，" .. groupId)
            return nil
        end
        for index, _ in pairs(generator) do
            local subIndex = 1                    -- 用；分割的取其中一个
            if #generator[index] > 1 then
                subIndex = ZRandom.GetList(generator[index].rateList)       -- 多个元素就随机一个 ，这里是用；分割的取其中一个，除非都要
            end
            _setGroupNum(groupList,generator[index][subIndex].generateId)
            -- 这里是全部掉落，不再使用生成组的；只取一个的规则，而是全部都要，用于踢树暴击
            if all~=nil and all == true then
                for _,v in ipairs(generator[index]) do
                    _setGroupNum(groupList,v.generateId)
                end
            end
        end
    end

    local itemList = {}
    for groupId2,num in pairs(groupList) do
        --groupId = tostring(groupId)
        local itemId = CSV_generateItem.GetValue(groupId2, "GenObjectId")
        local itemNum  = math.floor( _getGroupItemNum(groupId2) * scale)
        itemList[itemId] = num * itemNum

        if num * itemNum == 0 then
            itemList[itemId] = nil   -- 这是因为表格里面配置有数量为0， 不掉， 那配置它干啥呢， 理解不了，闲的蛋疼 ，害的我代码都乱了
        end
    end

    --printTable(itemList)
    return itemList
    --return nil
end