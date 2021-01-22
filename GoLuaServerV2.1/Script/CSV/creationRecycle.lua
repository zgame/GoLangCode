local creationRecycle = {
["19340001"] = {  DropSceneList = 20500076,  DropCountPercent = 0.4,  RecycleTimeMinute = 120,  MachineLevel = 1,  OrderId = 0,  Exp = 4,  },
["19340002"] = {  DropSceneList = 20500077,  DropCountPercent = 0.4,  RecycleTimeMinute = 150,  MachineLevel = 1,  OrderId = 0,  Exp = 4,  },
["19340003"] = {  DropSceneList = 20500078,  DropCountPercent = 0.4,  RecycleTimeMinute = 120,  MachineLevel = 1,  OrderId = 0,  Exp = 4,  },
["19340004"] = {  DropSceneList = 20500079,  DropCountPercent = 0.4,  RecycleTimeMinute = 180,  MachineLevel = 1,  OrderId = 0,  Exp = 10,  },
["19340005"] = {  DropSceneList = 20500080,  DropCountPercent = 0.4,  RecycleTimeMinute = 150,  MachineLevel = 1,  OrderId = 0,  Exp = 6,  },
["19340006"] = {  DropSceneList = 20500081,  DropCountPercent = 0.4,  RecycleTimeMinute = 150,  MachineLevel = 1,  OrderId = 0,  Exp = 8,  },
["19340007"] = {  DropSceneList = 20500097,  DropCountPercent = 0.4,  RecycleTimeMinute = 210,  MachineLevel = 1,  OrderId = 9999,  Exp = 6,  },
["19340008"] = {  DropSceneList = 20500098,  DropCountPercent = 0.4,  RecycleTimeMinute = 270,  MachineLevel = 1,  OrderId = 9999,  Exp = 6,  },
["19340009"] = {  DropSceneList = 20500099,  DropCountPercent = 0.4,  RecycleTimeMinute = 240,  MachineLevel = 1,  OrderId = 9999,  Exp = 6,  },
["19340010"] = {  DropSceneList = 20500100,  DropCountPercent = 0.4,  RecycleTimeMinute = 180,  MachineLevel = 1,  OrderId = 9999,  Exp = 6,  },
["19340011"] = {  DropSceneList = 20500101,  DropCountPercent = 0.4,  RecycleTimeMinute = 180,  MachineLevel = 1,  OrderId = 9999,  Exp = 6,  },
["19340012"] = {  DropSceneList = 20500102,  DropCountPercent = 0.4,  RecycleTimeMinute = 240,  MachineLevel = 1,  OrderId = 9999,  Exp = 6,  },
["19340013"] = {  DropSceneList = 20500122,  DropCountPercent = 0.4,  RecycleTimeMinute = 60,  MachineLevel = 1,  OrderId = 9999,  Exp = 3,  },
} 
 CSV_creationRecycle = {}

function CSV_creationRecycle.GetValue(index, key)
	index = tostring(index)
	key = tostring(key)
    if creationRecycle[index] == nil then
        ZLog.Logger("Excel 获取表:creationRecycle  主键:".. index .." key:".. key.."出错!")
        return nil
    end
    if creationRecycle[index][key] == nil then
        ZLog.Logger("Excel 获取表: creationRecycle  主键:".. index .." key:".. key.."出错!")
        return nil
    end

    return creationRecycle[index][key]
end


function CSV_creationRecycle.Get()
    return creationRecycle
end

function CSV_creationRecycle.GetAllKeys()
    local keys = {}
    for k in pairs(creationRecycle) do
        keys[#keys + 1] = k
    end
    if #keys == 0 then
        ZLog.Logger(string.format("Excel 获取表: 所有主键出错"))
        return nil
    end
    return keys
end