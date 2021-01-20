local creationMachine = {
["49000011"] = {  MachineLevel = 1,  ItemId = 13000013,  PartIds = {98,99,100},  },
["49000080"] = {  MachineLevel = 1,  ItemId = 14000035,  PartIds = {244,245,246},  },
["49000081"] = {  MachineLevel = 1,  ItemId = 14000036,  PartIds = {},  },
["49000082"] = {  MachineLevel = 1,  ItemId = 14000037,  PartIds = {249,250,251,252},  },
["49000001"] = {  MachineLevel = 1,  ItemId = 13100001,  PartIds = {1,2,233},  },
["49000074"] = {  MachineLevel = 1,  ItemId = 14000033,  PartIds = {223,224,225,226,247},  },
["49000003"] = {  MachineLevel = 1,  ItemId = 13100005,  PartIds = {6,7,9,208},  },
["49000071"] = {  MachineLevel = 1,  ItemId = 14000031,  PartIds = {198,199,200,201},  },
["49000070"] = {  MachineLevel = 1,  ItemId = 14000030,  PartIds = {195,196,197},  },
["49000078"] = {  MachineLevel = 1,  ItemId = 14000034,  PartIds = {234,235,236},  },
["49000012"] = {  MachineLevel = 1,  ItemId = 13000007,  PartIds = {15,16,17},  },
["49000013"] = {  MachineLevel = 1,  ItemId = 13100034,  PartIds = {3,4,5,215},  },
["49000062"] = {  MachineLevel = 1,  ItemId = 14000024,  PartIds = {166,167,168,169},  },
["49000006"] = {  MachineLevel = 1,  ItemId = 14000003,  PartIds = {72,73,103},  },
["49000007"] = {  MachineLevel = 1,  ItemId = 14000004,  PartIds = {30,75,90,102},  },
["49000068"] = {  MachineLevel = 1,  ItemId = 14000028,  PartIds = {187,188,190},  },
["49000008"] = {  MachineLevel = 1,  ItemId = 13100009,  PartIds = {58,59},  },
["49000073"] = {  MachineLevel = 1,  ItemId = 14000020,  PartIds = {218,219,220,221},  },
["49000072"] = {  MachineLevel = 1,  ItemId = 14000032,  PartIds = {202,203,204},  },
["49000047"] = {  MachineLevel = 1,  ItemId = 14000021,  PartIds = {79,80,81,101,217},  },
["49000067"] = {  MachineLevel = 2,  ItemId = 13100042,  PartIds = {183,184,185,186},  },
["49000066"] = {  MachineLevel = 2,  ItemId = 13100041,  PartIds = {180,181,182,237},  },
["49000046"] = {  MachineLevel = 1,  ItemId = 14000022,  PartIds = {82,83,84,85},  },
["49000015"] = {  MachineLevel = 1,  ItemId = 13100016,  PartIds = {92,93,94,238},  },
["49000005"] = {  MachineLevel = 1,  ItemId = 13100022,  PartIds = {23,24,25,239},  },
["49000002"] = {  MachineLevel = 1,  ItemId = 13100002,  PartIds = {8,14,18,19},  },
["49000000"] = {  MachineLevel = 1,  ItemId = 14000006,  PartIds = {22,28,32},  },
["49000014"] = {  MachineLevel = 1,  ItemId = 13100012,  PartIds = {95,96,97},  },
["49000010"] = {  MachineLevel = 1,  ItemId = 13100006,  PartIds = {10,11,12,13},  },
["49000039"] = {  MachineLevel = 1,  ItemId = 13100019,  PartIds = {107,108,109,240},  },
} 
 CSV_creationMachine = {}

function CSV_creationMachine.GetValue(index, key)
	index = tostring(index)
	key = tostring(key)
    if creationMachine[index] == nil then
        ZLog.Logger("Excel 获取表:creationMachine  主键:".. index .." key:".. key.."出错!")
        return nil
    end
    if creationMachine[index][key] == nil then
        ZLog.Logger("Excel 获取表: creationMachine  主键:".. index .." key:".. key.."出错!")
        return nil
    end

    return creationMachine[index][key]
end


function CSV_creationMachine.Get()
    return creationMachine
end

function CSV_creationMachine.GetAllKeys()
    local keys = {}
    for k in pairs(creationMachine) do
        keys[#keys + 1] = k
    end
    if #keys == 0 then
        ZLog.Logger(string.format("Excel 获取表: 所有主键出错"))
        return nil
    end
    return keys
end