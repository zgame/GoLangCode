
-- 道具生成组
SandRockItemGenerator = {}

-- 初始化道具生成组的数据
local function _setGroup(groupId)
    SandRockItemGenerator[groupId] = {}
    local generator = SandRockItemGenerator[groupId]

    local GeneratorList = CSV_generateGroup.GetValue(groupId, "GeneratorList")
    --local GeneratorList = "29600315,100,0.0|29600072,200,1.0;29500000,300,0.0;29600605,150,1.0|29500000,400,0.0;29600606,300,1.0"
    local allDropList = ZString.Split(GeneratorList,"|")
    for i, mainDrop in ipairs(allDropList) do
        print(mainDrop)
        local subList = ZString.Split(mainDrop,";")
        for i,subDrop in ipairs(subList) do
            print(subDrop)
        end
    end

end

function SandRockItemGenerator.Init()
    local GeneratorList = CSV_generateGroup.GetAllKeys()
    --for _, groupId in ipairs(GeneratorList) do
    --    _setGroup(groupId)
    --end
    _setGroup("20900008")
end


--  道具掉落 ，正常返回 hash  key是itemId， value是数量
function SandRockItemGenerator.GetItems(groupId)
    groupId = 20900008
    local GenSceneType = CSV_generateGroup.GetValue(groupId, "GenSceneType")
    if GenSceneType == "Item" then
        -- 走道具掉落规则
        for k, v in pairs(SandRockItemGenerator[tostring(groupId)].GeneratorList) do

        end

    end









    local itemList ={}
    itemList[999] = 1
    return itemList
    --return nil
end