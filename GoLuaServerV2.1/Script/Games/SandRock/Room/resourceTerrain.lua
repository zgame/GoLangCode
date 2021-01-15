----------------------------------------------------------------
--- 地形树和石头资源刷新
----------------------------------------------------------------




-----------------------------------地形树 刷新------------------------------------------
-- 资源点刷新
function SandRockRoom:ResourceTerrainUpdate()
    -- 判断生命周期， 到期的给删除掉
    for areaName, pointList in pairs(self.resourcePoint) do
        --print(areaName)
        --printTable(pointList)
        for index, point in pairs(pointList) do
            if point.live <= 1 then
                pointList[index] = nil              -- 删掉生命周期已经到了的点
            else
                point.live = point.live - 1
            end
        end
    end
    -- 开始刷新新东西
    --local areaList = CSV_resourceArea.Get()
    for areaName,_ in pairs(CSV_resourcePickArea.Get()) do
        --print("areaName"..areaName)
        if self.resourcePoint[areaName] == nil then
            self.resourcePoint[areaName] = {}           -- 初始化生成点列表
        end

        --print("随机获取本次更新资源数量num ："..num)
        local num = SandRockResourcePick.GetNum(areaName)
        local number_now = ZTable.Len(self.resourcePoint[areaName])        -- 已经包含多少个点
        --print("number_now"..number_now)
        if num > number_now then
            for i = 1, num - number_now do
                --print('生成一个point, 下面是point的结构')
                local resourceTypeRandom = SandRockResourcePick.GetType(areaName)           -- 获取一个生成类型，根据权重
                if resourceTypeRandom == "0" then
                    break
                end
                local element = {}
                local areaPoint = _getEmpty(areaName, self.resourcePoint)   -- 获取一个空的位置
                element.resourceType = tonumber(resourceTypeRandom)
                element.live = CSV_resourcePickType.GetValue(resourceTypeRandom, "LifeCycle")
                --print("保存到房间的资源列表里面")
                self.resourcePoint[areaName][areaPoint] = element
                --printTable(self.resourcePoint[areaName])
            end
        end

    end

    --printTable(self.resourcePoint)
end

-----------------------------------地形树 采集------------------------------------------
function SandRockRoom:GetTerrainResource(userId, areaName, pointIndex, resourceType)
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
    local itemList = SandRockGeneratorItem.GetItems(generatorGroup)
    -- 保存到背包


    -- 保存完毕
    return itemList

end