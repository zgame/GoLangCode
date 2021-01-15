-----------------------------------------------------------------------------------------------
--- 资源刷新点的数据，主要是树木和石头
-----------------------------------------------------------------------------------------------


-- 资源生成组
SandRockResourceTerrain = {}

-- 是否可以进行交互
function SandRockResourceTerrain.CantKick(resourceId)
    local CantKick = CSV_resourceTerrainType.GetValue(tostring(resourceId),'Tags')
    if CantKick == "" then
        return true
    end
    return false
end

function SandRockResourceTerrain.GetType()

end

-- 获取树的生命值
function SandRockResourceTerrain.GetHp(treeId)
    local TrunkHealth= CSV_resourceTerrainType.GetValue(treeId,"TrunkHealth")
    local StumpHealth= CSV_resourceTerrainType.GetValue(treeId,"StumpHealth")
    return TrunkHealth,StumpHealth
end