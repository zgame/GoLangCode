----------------------------------------------------------------
--- 资源刷新， 每天刷新一次， 到生命周期就删掉  ， 不足上限数量就补充，但是不一定非要补充，是随机的
----------------------------------------------------------------


-- 获取没有被占用的资源点列表
local function _getEmpty(areaName,resourcePoint)
    local number_max = CSV_resourcePickArea.GetValue(areaName, 'Points')
    if number_max == 0 then
        return 1
    end
    local pointList = {}            --没有占用的列表
    for j = 1, number_max do
        table.insert(pointList, j)           -- 生成全数组
    end
    for pointIndex, _ in pairs(resourcePoint[areaName]) do
        table.remove(pointList, pointIndex)         -- 把已经占用的删掉
    end
    -- 剩下的就是没有被占用的了
    local ran = ZRandom.GetRandom(1, #pointList)
    return pointList[ran]
end

-- 创建一个采集资源点数据
function SandRockRoom:_createPoint(areaName)
    --print('生成一个point, 下面是point的结构')
    local resourceTypeRandom = SandRockResourcePick.GetType(areaName)           -- 获取一个生成类型，根据权重
    --print("调试看看生成类型".. resourceTypeRandom)
    if resourceTypeRandom == 0 then
        return
    end
    local element = {}
    local areaPoint = _getEmpty(areaName,self.resourcePoint)   -- 获取一个空的位置
    element.resourceType = tonumber(resourceTypeRandom)
    element.live = CSV_resourcePickType.GetValue(resourceTypeRandom, "LifeCycle")
    --print("保存到房间的资源列表里面")
    self.resourcePoint[areaName][areaPoint] = element
    return areaPoint,element
end


-- 初始化
function SandRockRoom:ResourcePickPointInit()
    for areaName,_ in pairs(CSV_resourcePickArea.Get()) do
        --print("areaName"..areaName)
        self.resourcePoint[areaName] = {}           -- 初始化生成点列表

        --print("随机获取本次更新资源数量num ："..num)
        local num = SandRockResourcePick.GetNum(areaName)
            for i = 1, num do
                self:_createPoint(areaName)
            end
    end
    --printTable(self.resourcePoint)
end




-----------------------------------刷新------------------------------------------
-- 资源点刷新
function SandRockRoom:ResourcePickPointUpdate()
    local updateList = {}
    -- 判断生命周期， 到期的给删除掉
    for areaName, pointList in pairs(self.resourcePoint) do
        --print(areaName)
        --printTable(pointList)
        for index, point in pairs(pointList) do
            if point.live ~= nil then
                if point.live <= 1 then
                    pointList[index] = nil              -- 删掉生命周期已经到了的点

                    local element = {}
                    element.areaName = areaName
                    element.areaPoint = index
                    element.resourceType = 0
                    table.insert(updateList,element)        -- 减少的资源点
                else
                    point.live = point.live - 1
                end
            end
        end
    end
    -- 开始刷新新东西
    --local areaList = CSV_resourceArea.Get()
    for areaName,_ in pairs(CSV_resourcePickArea.Get()) do
        --print("areaName"..areaName)
        --print("随机获取本次更新资源数量num ："..num)
        local num = SandRockResourcePick.GetNum(areaName)
        local number_now = ZTable.Len(self.resourcePoint[areaName])        -- 已经包含多少个点
        --print("number_now"..number_now)
        if num > number_now then
            for i = 1, num - number_now do
                local areaPoint,element = self:_createPoint(areaName)
                if areaPoint~=nil then
                    local element2 = {}
                    element2.areaName = areaName
                    element2.areaPoint = areaPoint
                    element2.resourceType = element.resourceType
                    table.insert(updateList,element2)           -- 增加的资源点
                end
            end
        end
    end

    --printTable(updateList)
    return updateList
end

-----------------------------------采集------------------------------------------
function SandRockRoom:GetPickResource(userId, areaName, pointIndex, resourceType)
    if self.resourcePoint[areaName] == nil then
        ZLog.Logger("GetResource  areaName 生成区域出错 " .. areaName)
        return
    end

    local point = self.resourcePoint[areaName][pointIndex]
    if point == nil then
        ZLog.Logger("GetResource pointIndex 生成点index出错" .. pointIndex)
        return
    end
    if point.resourceType ~= resourceType then
        ZLog.Logger("GetResource resourceType 生成资源类型出错" .. resourceType)
    end
    local player = GameServer.GetPlayerByUID(userId)
    -- 采集
    local spCost = CSV_resourcePickType.GetValue(resourceType, "SpCost")
    local exp = CSV_resourcePickType.GetValue(resourceType, "Exp")

    Player.ExpAdd(player, exp)
    Player.SpAdd(player, -spCost)

    -- 销毁采集点
    self.resourcePoint[areaName][pointIndex] = nil
    -- 获得物品
    local generatorGroup = CSV_resourcePickType.GetValue(resourceType, "GeneratorGroup")
    local itemList = SandRockGeneratorItem.GetItems(generatorGroup)             -- pick采集
    -- 保存到背包

    Player.ItemAdd(player,itemList)

    -- 保存完毕
    return itemList

end