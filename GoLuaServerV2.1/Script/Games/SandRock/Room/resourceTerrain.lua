----------------------------------------------------------------
--- 地形树和石头资源刷新
----------------------------------------------------------------

-- 生成地形资源树和石头的结构
function SandRockRoom:ResourceTerrainInit()
    for areaName, _ in pairs(CSV_resourceTerrainArea.Get()) do
        --print(areaName)
        self.resourceTerrain[areaName] = {}
        self.resourceTerrainChange[areaName] = {}

        local TreeIDList = CSV_resourceTerrainArea.GetValue(areaName, "TreeID")
        local TreeScale = CSV_resourceTerrainArea.GetValue(areaName, "TreeScale")
        for index, treeId in pairs(TreeIDList) do
            if treeId > 0 then
                local element = {}
                element.areaName = areaName
                element.areaPoint = index        -- 第几个位置
                element.resourceType = treeId
                element.scale = SandRockResourceTerrain.GetScaleFactor(TreeScale[index], 0.25)                    -- 缩放比例
                element.trunkHealth, element.stumpHealth = SandRockResourceTerrain.GetHp(treeId, element.scale)
                element.trunkHealthMax = element.trunkHealth
                element.stumpHealthMax = element.stumpHealth
                element.kickCountLimit = CSV_resourceTerrainType.GetValue(treeId, "KickCountLimit")
                element.cantKick = SandRockResourceTerrain.CantKick(treeId)
                element.relive = CSV_resourceTerrainType.GetValue(treeId, "RespawnDays")

                self.resourceTerrain[areaName][index] = element
            end
        end
    end

    --printTable(self.resourceTerrain)

end

-----------------------------------地形树 同步------------------------------------------
function SandRockRoom:ResourceTerrainSync()
    local treeList = {}
    for _, pointList in pairs(self.resourceTerrainChange) do
        for _, element in pairs(pointList) do
            --local treeId = element.resourceType
            --local trunkHealth, stumpHealth = SandRockResourceTerrain.GetHp(treeId, element.scale)
            --if element.trunkHealth < element.trunkHealthMax then
                -- 说明这棵树被伤害过，要同步
                table.insert(treeList, element)
            --end
        end
    end
    return treeList
    --return self.resourceTerrainChange
end

-----------------------------------地形树 刷新------------------------------------------
-- 资源点刷新
function SandRockRoom:ResourceTerrainUpdate()
    local reliveList = {}
    -- 判断重生的时间
    for _, pointList in pairs(self.resourceTerrain) do
        for _, element in pairs(pointList) do
            if element.cantKick == false then
                -- 刷新踢树上限
                element.kickCountLimit = CSV_resourceTerrainType.GetValue(element.resourceType, "KickCountLimit")
            end

            -- 如果死亡，刷新
            if element.trunkHealth <=0 and element.stumpHealth <= 0 then
                if element.relive <= 0 then
                    -- 重生
                    element.trunkHealth = element.trunkHealthMax
                    element.stumpHealth = element.stumpHealthMax
                    element.relive = CSV_resourceTerrainType.GetValue(element.resourceType, "RespawnDays")
                    table.insert(reliveList, element)           -- 同步给客户端树重生

                    self.resourceTerrainChange[element.areaName][element.areaPoint] = nil           -- 变化同步的列表要去掉，因为树重置了之后没有变化
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
local function _treeDamage(room,element, damage)
    local reliveList = {}
    local trunkKill = false
    local stumpKill = false

    if element.trunkHealth > 0 then
        -- 先判断树干
        element.trunkHealth = element.trunkHealth - damage
        if element.trunkHealth <= 0 then
            element.trunkHealth = 0
            trunkKill = true               -- 树干死亡
        end
    else
        -- 再判断树根
        if element.stumpHealth > 0 then
            element.stumpHealth = element.stumpHealth - damage
            if element.stumpHealth <= 0 then
                element.stumpHealth = 0
                stumpKill = true                -- 树根死亡
            end
        end
    end
    table.insert(reliveList, element)

    -- 树有变化了，增加到变化列表里面去
    room.resourceTerrainChange[element.areaName][element.areaPoint] = element           -- 变化同步的列表要去掉，因为树重置了之后没有变化


    return reliveList, trunkKill, stumpKill
end

-- 砍树
function SandRockRoom:GetTerrainResource(userId, areaName, pointIndex, resourceType, toolId, damage)
    if damage <= 0 then
        ZLog.Logger(userId.. "砍树伤害不正确 " .. damage)
        return
    end

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
        return
    end
    local player = GameServer.GetPlayerByUID(userId)
    -- 地形树采集

    -- 工具砍树，或者砸石头
    -- 体力判断

    local spCostList = CSV_itemFunctions.GetValue(toolId, "SpCost1")
    local exp = CSV_resourceTerrainType.GetValue(resourceType, "ChopTrunkExp")
    Player.ExpAdd(player, exp)
    for _, v in ipairs(spCostList) do
        Player.SpAdd(player, -v)
    end

    -- 树的伤害
    local reliveList, trunkKill, stumpKill = _treeDamage(self, point, damage)
    --print("树的伤害")
    --printTable(reliveList)
    --print(trunkKill)
    --print(stumpKill)

    -- 获得物品
    local itemList
    if trunkKill then
        local ChopTrunkDropId = CSV_resourceTerrainType.GetValue(resourceType, "ChopTrunkDropId")
        itemList = SandRockGeneratorItem.GetItems(ChopTrunkDropId, point.scale)                          -- 砍树干
    end
    if stumpKill then
        local ChopStumpDropId = CSV_resourceTerrainType.GetValue(resourceType, "ChopStumpDropId")
        itemList = SandRockGeneratorItem.GetItems(ChopStumpDropId,  point.scale)                  -- 砍树根
    end


    -- 保存到背包
    Player.ItemAdd(player, itemList)


    -- 保存完毕
    return itemList, reliveList

end

-- 踢树
function SandRockRoom:GetTerrainKick(userId, areaName, pointIndex, resourceType)
    if self.resourceTerrain[areaName] == nil then
        ZLog.Logger("GetTerrainKick  areaName 生成区域出错 " .. areaName)
        return
    end

    local point = self.resourceTerrain[areaName][pointIndex]
    if point == nil then
        ZLog.Logger("GetTerrainKick pointIndex 生成点index出错" .. pointIndex)
        return
    end
    if point.resourceType ~= resourceType then
        ZLog.Logger("GetTerrainKick resourceType 生成资源类型出错" .. resourceType)
        return
    end
    local player = GameServer.GetPlayerByUID(userId)
    -- 地形树采集


    -- 踢树
    if SandRockResourceTerrain.CantKick(resourceType) == true  then
        ZLog.Logger("这颗树不能踢")
        return nil
    end
    --if Player.SpGet(player) <= ConstSandRock.TickCostSp then
    --    ZLog.Logger("没有体力踢树")
    --    return nil
    --end

    local exp = CSV_resourceTerrainType.GetValue(resourceType, "KickExp")
    Player.ExpAdd(player, exp)
    Player.SpAdd(player, -ConstSandRock.TickCostSp)-- 踢树消耗体力固定
    -- 踢树几率
    local KickDropChance = CSV_resourceTerrainType.GetValue(resourceType, "KickDropChance")
    if ZRandom.PercentRate(KickDropChance) == false then
        return nil                  -- 没有踢中
    end

    --  踢树暴击
    local itemList
    local KickDropId = CSV_resourceTerrainType.GetValue(resourceType, "KickDropId")
    if KickDropId == 0 then
        print("踢该树木并不掉落东西")
        return nil
    end

    local KickAllDropChance = CSV_resourceTerrainType.GetValue(resourceType, "KickAllDropChance")
    if ZRandom.PercentRate(KickAllDropChance) then
        itemList = SandRockGeneratorItem.GetItems(KickDropId, point.scale, true)              -- 踢树暴击
        --print("踢树暴击")
    else
        itemList = SandRockGeneratorItem.GetItems(KickDropId, point.scale)                           -- 踢树不暴击
        --print("普通踢树")
    end

    -- 保存到背包
    Player.ItemAdd(player, itemList)

    return itemList

end

