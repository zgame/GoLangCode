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
            element.relive = CSV_resourceTerrainType.GetValue(treeId, "RespawnDays")
            --element.trunkLife = true
            --element.stumpLife = true

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
                    element.relive = CSV_resourceTerrainType.GetValue(treeId, "RespawnDays")
                    --element.trunkLife = true
                    --element.stumpLife = true
                    table.insert(reliveList, element)
                else
                    element.relive = element.relive - 1
                end
            end

        end
    end
    --printTable(self.resourcePoint)
    return reliveList
end

-----------------------------------地形树 采集------------------------------------------
local function _treeDamage(element, damage)
    --local reliveList = {}
    local trunkKill = false
    local stumpKill = false

    if element.trunkHealth > 0 then             -- 先判断树干
        element.trunkHealth = element.trunkHealth - damage
        if element.trunkHealth <= 0 then
            element.trunkHealth = 0
            trunkKill = true               -- 树干死亡
        end
    else                                        -- 再判断树根
        if element.stumpHealth > 0 then
            element.stumpHealth = element.stumpHealth - damage
            if element.stumpHealth <= 0 then
                element.stumpHealth = 0
                stumpKill = true                -- 树根死亡
            end
        end
    end
   return element, trunkKill, stumpKill
end

function SandRockRoom:GetTerrainResource(userId, areaName, pointIndex, resourceType, toolId ,damage)
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
        if Player.SpGet(player ) <= ConstSandRock.TickCostSp then
            ZLog.Logger("没有体力踢树")
            return nil, nil
        end

        local exp = CSV_resourceTerrainType.GetValue(resourceType, "KickExp")
        Player.ExpAdd(player, exp)
        Player.SpAdd(player, -ConstSandRock.TickCostSp )-- 踢树消耗体力固定
        -- 踢树几率
        local KickDropChance = CSV_resourceTerrainType.GetValue(resourceType, "KickDropChance")
        if ZRandom.PercentRate(KickDropChance) == false then
            return nil,nil                  -- 没有踢中
        end

        --  踢树暴击
        local itemList
        local KickDropId = CSV_resourceTerrainType.GetValue(resourceType, "KickDropId")
        if KickDropId == 0 then
            print("踢该树木并不掉落东西")
            return nil,nil
        end

        local KickAllDropChance = CSV_resourceTerrainType.GetValue(resourceType, "KickAllDropChance")
        if ZRandom.PercentRate(KickAllDropChance)  then
            itemList = SandRockGeneratorItem.GetItems(KickDropId,true)              -- 踢树暴击
        else
            itemList = SandRockGeneratorItem.GetItems(KickDropId)                           -- 不暴击
        end

        -- 保存到背包
        Player.ItemAdd(player,itemList)

        return itemList, nil
    else
        -- 工具砍树，或者砸石头
        -- 体力判断

        local spCostList = CSV_itemFunctions.GetValue(toolId, "SpCost1")
        local exp = CSV_resourceTerrainType.GetValue(resourceType, "ChopTrunkExp")
        Player.ExpAdd(player, exp)
        for _, v in ipairs(spCostList) do
            Player.SpAdd(player, -v)
        end

        -- 树的伤害
        local reliveList,trunkKill, stumpKill = _treeDamage(point,damage)

        -- 获得物品
        local itemList
        if  trunkKill then
            local ChopTrunkDropId = CSV_resourceTerrainType.GetValue(resourceType, "ChopTrunkDropId")
            itemList = SandRockGeneratorItem.GetItems(ChopTrunkDropId)
        end
        if stumpKill then
            local ChopStumpDropId = CSV_resourceTerrainType.GetValue(resourceType, "ChopStumpDropId")
            itemList = SandRockGeneratorItem.GetItems(ChopStumpDropId)
        end


        -- 保存到背包
        Player.ItemAdd(player,itemList)

        -- 保存完毕
        return itemList, reliveList

    end
end

