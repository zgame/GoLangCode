-----------------------------------------------------------------------------------------------
--- 资源刷新点的数据 ，因为效率问题， 所以初始化的时候， 进行一些数据处理， 方便以后调用
-----------------------------------------------------------------------------------------------


-- 资源生成组
SandRockResourceGenerator = {}


-- 根据权重随机生成类型
local function _setType(areaName)
    SandRockResourceGenerator[areaName] ={}
    local area = SandRockResourceGenerator[areaName]
    local resourceType = CSV_resourceArea.GetValue(areaName, "Resource")
    local weight = CSV_resourceArea.GetValue(areaName, "Weight")

    local resourceList = ZString.Split(resourceType,",")
    if #resourceList == 1 then
        area.resourceType = resourceType        -- 赋值数值
        return
    else
        area.resourceList = resourceList        -- 赋值list
    end
    local weightList = ZString.Split(weight,",")
    area.weightList = weightList

end


-- 为了效率，初始化的时候生成一下数据
function SandRockResourceGenerator.Init()
    for  areaName,_ in pairs(CSV_resourceArea.Get()) do
        _setType(areaName)
    end
end


-- 根据权重随机生成类型
function SandRockResourceGenerator.GetType(areaName)
    local area = SandRockResourceGenerator[areaName]

    -- 如果没有数组， 那么就直接返回数值
    if area.resourceList == nil then
        return area.resourceType
    end

    local index = ZRandom.GetList(area.weightList)
    return  ZString.Trim(area.resourceList[index])
end
