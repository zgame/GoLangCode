local resourceTerrainArea = {
["Terrain_2_2"] = {  TreeIndex = {0,4,5,6,7,8},  TreeID = {4,50,51,52,53,54},  },
["Terrain_2_3"] = {  TreeIndex = {0},  TreeID = {4},  },
["Terrain_2_4"] = {  TreeIndex = {3},  TreeID = {4},  },
["Terrain_2_5"] = {  TreeIndex = {3},  TreeID = {4},  },
["Terrain_3_2"] = {  TreeIndex = {0},  TreeID = {4},  },
["Terrain_3_3"] = {  TreeIndex = {0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19},  TreeID = {4,5,40,100,101,102,103,104,105,41,6,109,20,106,13,1,14,15,108,52},  },
["Terrain_3_4"] = {  TreeIndex = {0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30},  TreeID = {4,50,51,52,53,54,5,6,40,41,13,15,20,100,101,102,103,104,105,7,31,32,30,106,108,107,21,109,2,1,14},  },
["Terrain_3_5"] = {  TreeIndex = {3,4,5,6},  TreeID = {4,5,20,6},  },
["Terrain_4_2"] = {  TreeIndex = {3,4},  TreeID = {4,20},  },
["Terrain_4_3"] = {  TreeIndex = {3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19},  TreeID = {4,5,40,6,1,2,3,10,12,20,13,14,15,102,41,108,52},  },
["Terrain_4_4"] = {  TreeIndex = {0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27},  TreeID = {4,6,5,40,41,52,14,15,13,20,54,51,100,101,102,103,104,105,1,2,3,10,12,7,107,21,106,108},  },
["Terrain_4_5"] = {  TreeIndex = {3,4,5,6},  TreeID = {4,20,5,6},  },
["Terrain_5_2"] = {  TreeIndex = {3,4},  TreeID = {4,20},  },
["Terrain_5_3"] = {  TreeIndex = {3,4},  TreeID = {4,6},  },
["Terrain_5_4"] = {  TreeIndex = {3,4},  TreeID = {4,6},  },
["Terrain_5_5"] = {  TreeIndex = {3,4},  TreeID = {4,20},  },
["Terrain_5_6"] = {  TreeIndex = {0,1,2,3},  TreeID = {53,52,5,6},  },
} 
 CSV_resourceTerrainArea = {}

function CSV_resourceTerrainArea.GetValue(index, key)
	index = tostring(index)
	key = tostring(key)
    if resourceTerrainArea[index] == nil then
        ZLog.Logger("Excel 获取表:resourceTerrainArea  主键:".. index .." key:".. key.."出错!")
        return nil
    end
    if resourceTerrainArea[index][key] == nil then
        ZLog.Logger("Excel 获取表: resourceTerrainArea  主键:".. index .." key:".. key.."出错!")
        return nil
    end

    return resourceTerrainArea[index][key]
end


function CSV_resourceTerrainArea.Get()
    return resourceTerrainArea
end

function CSV_resourceTerrainArea.GetAllKeys()
    local keys = {}
    for k in pairs(resourceTerrainArea) do
        keys[#keys + 1] = k
    end
    if #keys == 0 then
        ZLog.Logger(string.format("Excel 获取表: 所有主键出错"))
        return nil
    end
    return keys
end