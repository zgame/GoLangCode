-- 道具生成组
SandRockItemGenerator = {}


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
    SandRockItemGenerator[groupId] = {}
    local generator = SandRockItemGenerator[groupId]

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

function SandRockItemGenerator.Init()
    local GeneratorList = CSV_generateGroup.GetAllKeys()
    for _, groupId in ipairs(GeneratorList) do
        _setGroup(groupId)
    end
    --_setGroup("20900008")
end


--  道具掉落 ，正常返回 hash  key是itemId， value是数量
function SandRockItemGenerator.GetItems(groupId)
    local groupList = {}

    local GenSceneType = CSV_generateGroup.GetValue(groupId, "GenSceneType")
    if GenSceneType == "Item" then
        -- 走道具掉落规则
        local generator = SandRockItemGenerator[tostring(groupId)]
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
        local itemId = CSV_generateItem.GetValue(tostring(groupId), "GenObjectId")
        -- 根据生成规则进行生成

        .... ...................................
        itemList[itemId] = num
    end


    --printTable(itemList)
    return itemList
    --return nil
end