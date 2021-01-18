----------------------------------------------------------------
--- 地形树和石头资源刷新
----------------------------------------------------------------

-- 生成地形资源树和石头的结构
function SandRockRoom:ResourceTerrainInit()
    for areaName, _ in pairs(CSV_resourceTerrainArea.Get()) do
        --print(areaName)
        self.resourceTerrain[areaName] = {}
        local TreeIndexList = CSV_resourceTerrainArea.GetValue(areaName, "TreeIndex")
        local TreeIDList = CSV_resourceTerrainArea.GetValue(areaName, "TreeID")
        for index, treeId in pairs(TreeIDList) do
            local element = {}
            element.areaName = areaName
            element.areaPoint = TreeIndexList[index]        -- 第几个位置
            element.resourceType = treeId
            element.trunkHealth, element.stumpHealth = SandRockResourceTerrain.GetHp(treeId)
            element.kickCountLimit = CSV_resourceTerrainType.GetValue(treeId, "KickCountLimit")

            self.resourceTerrain[areaName][TreeIndexList[index]] = element
        end
    end
    --printTable(self.resourceTerrain)
end



-----------------------------------地形树 刷新------------------------------------------
-- 资源点刷新
function SandRockRoom:ResourceTerrainUpdate()
    local reliveList = {}
    -- 判断重生的时间
    for areaName, pointList in pairs(self.resourceTerrain) do
        for index, element in pairs(pointList) do
            local treeId = element.resourceType
            -- 刷新踢树上限
            element.kickCountLimit = CSV_resourceTerrainType.GetValue(treeId, "KickCountLimit")

            -- 如果死亡，刷新
            if element.trunkHealth + element.stumpHealth <= 0 then
                if element.relive <= 0 then
                    -- 重生
                    element.trunkHealth, element.stumpHealth = SandRockResourceTerrain.GetHp(treeId)
                    table.insert(reliveList, element)
                else
                    element.relive = element.relive - 1
                end
            end
            --local  RespawnDays = CSV_resourceTerrainType.GetValue(treeId, "RespawnDays")
        end
    end
    --printTable(self.resourcePoint)
    return reliveList
end

-----------------------------------地形树 采集------------------------------------------
local function _treeDamage(element, damage)

end

function SandRockRoom:GetTerrainResource(userId, areaName, pointIndex, resourceType, toolId)
    if self.resourceTerrain[areaName] == nil then
        ZLog.Logger("GetTerrainResource  areaName 生成区域出错 " .. areaName)
        return
    end

    local point = self.resourceTerrain[areaName][pointIndex]
    if point == nil then
        ZLog.Logger("GetTerrainResource pointIndex 生成点index出错" .. pointIndex)
        return
    end
    if point.resourceType ~= resourceType then
        ZLog.Logger("GetTerrainResource resourceType 生成资源类型出错" .. resourceType)
    end
    local player = GameServer.GetPlayerByUID(userId)
    -- 地形树采集

    if toolId == 0 then
        -- 踢树
        local CantKick = CSV_resourceTerrainType.GetValue(resourceType, "Tags")
        if CantKick == "" then
            ZLog.Logger("这颗树不能踢")
            return nil, nil
        end
        local spCost = ConstSandRock.TickCostSp   -- 踢树消耗体力固定
        local exp = CSV_resourceTerrainType.GetValue(resourceType, "KickExp")
        Player.ExpAdd(player, exp)
        Player.SpAdd(player, -spCost)
        -- 获得物品
        local KickDropId = CSV_resourceTerrainType.GetValue(resourceType, "KickDropId")
        local KickDropChance = CSV_resourceTerrainType.GetValue(resourceType, "KickDropChance")
        if ZRandom.PercentRate(KickDropChance) == false then
            return nil,nil                  -- 没有掉落
        end
        local KickAllDropChance = CSV_resourceTerrainType.GetValue(resourceType, "KickAllDropChance")
        if ZRandom.PercentRate(KickAllDropChance)  then
              print("踢树暴击， 但是不知道怎么掉")                                               -- 踢树暴击
        end
        local itemList = SandRockGeneratorItem.GetItems(KickDropId)

        -- 保存到背包

        return itemList, nil
    else
        -- 工具砍树，或者砸石头

        local spCostList = CSV_itemFunctions.GetValue(toolId, "SpCost1")
        local exp = CSV_resourceTerrainType.GetValue(resourceType, "ChopTrunkExp")
        Player.ExpAdd(player, exp)
        for _, v in ipairs(spCostList) do
            Player.SpAdd(player, -v)
        end


        -- 树的伤害
        self.resourcePoint[areaName][pointIndex] = nil
        -- 获得物品
        local ChopTrunkDropId = CSV_resourceTerrainType.GetValue(resourceType, "ChopTrunkDropId")
        local ChopStumpDropId = CSV_resourceTerrainType.GetValue(resourceType, "ChopStumpDropId")
        local itemList = SandRockGeneratorItem.GetItems(ChopTrunkDropId)

        -- 保存到背包



        -- 把树的变化更新一下
        local reliveList = {}

        -- 保存完毕
        return itemList, reliveList

    end
end

