-----------------------------------------------------------------------------------------------
--- 资源刷新点的数据，主要是树木和石头
-----------------------------------------------------------------------------------------------


-- 资源生成组
SandRockResourceTerrain = {}

-- 是否可以进行交互
function SandRockResourceTerrain.CantKick(resourceId)
    local CantKick = CSV_resourceTerrainType.GetValue(tostring(resourceId),'Tags')
    if CantKick == "CantKick" then
        return true
    end
    return false
end


-- 获取树的生命值
function SandRockResourceTerrain.GetHp(treeId, scale)
    local TrunkHealth= CSV_resourceTerrainType.GetValue(treeId,"TrunkHealth")
    local StumpHealth= CSV_resourceTerrainType.GetValue(treeId,"StumpHealth")

    -- 进行数值缩放
    --TrunkHealth = SandRockResourceTerrain.GetScaleFactor(scale, TrunkHealth)
    --StumpHealth = SandRockResourceTerrain.GetScaleFactor(scale, StumpHealth)

    TrunkHealth = math.floor(TrunkHealth*scale)
    StumpHealth = math.floor(StumpHealth*scale)
    return TrunkHealth,StumpHealth
end


-- 根据缩放比例，进行修改
function SandRockResourceTerrain.GetScaleFactor(scale, value)
    local  halfScale = scale * 0.5
    return 1 + value * ((6 * halfScale * halfScale * halfScale * halfScale * halfScale -
            15 * halfScale * halfScale * halfScale * halfScale +
            10 * halfScale * halfScale * halfScale) * 2 - 1)
end


