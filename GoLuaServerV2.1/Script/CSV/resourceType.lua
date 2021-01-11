local resourceType = {
["1"] = {  GeneratorGroup = 20900008,  Weather = "1,2,4",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["2"] = {  GeneratorGroup = 20900009,  Weather = "1,2,4",  DeleteWeather = 5,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["3"] = {  GeneratorGroup = 20900040,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["4"] = {  GeneratorGroup = 20900021,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["5"] = {  GeneratorGroup = 20900021,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["6"] = {  GeneratorGroup = 20900020,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["7"] = {  GeneratorGroup = 20900013,  Weather = "2",  DeleteWeather = 5,  LifeCycle = 2,  Exp = 10,  SpCost = 2,  },
["8"] = {  GeneratorGroup = 20900008,  Weather = "1,2,4",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["9"] = {  GeneratorGroup = 20990005,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["10"] = {  GeneratorGroup = 20900008,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["11"] = {  GeneratorGroup = 20960005,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["12"] = {  GeneratorGroup = 20960001,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["13"] = {  GeneratorGroup = 20960002,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["14"] = {  GeneratorGroup = 20960002,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["15"] = {  GeneratorGroup = 20960002,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["16"] = {  GeneratorGroup = 20960002,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["17"] = {  GeneratorGroup = 20960003,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["18"] = {  GeneratorGroup = 20960004,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["19"] = {  GeneratorGroup = 20960004,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["20"] = {  GeneratorGroup = 20960006,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["50"] = {  GeneratorGroup = 20500087,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["51"] = {  GeneratorGroup = 20900060,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["100"] = {  GeneratorGroup = 20940000,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["101"] = {  GeneratorGroup = 20910026,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["102"] = {  GeneratorGroup = 20940002,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["103"] = {  GeneratorGroup = 20940003,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["104"] = {  GeneratorGroup = 20940004,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["105"] = {  GeneratorGroup = 20940005,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["106"] = {  GeneratorGroup = 20940006,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["107"] = {  GeneratorGroup = 20940007,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["108"] = {  GeneratorGroup = 20940008,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["109"] = {  GeneratorGroup = 20940009,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["110"] = {  GeneratorGroup = 20940010,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["111"] = {  GeneratorGroup = 20940011,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["112"] = {  GeneratorGroup = 20940012,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["113"] = {  GeneratorGroup = 20940014,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["114"] = {  GeneratorGroup = 20940013,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["115"] = {  GeneratorGroup = 20940015,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["116"] = {  GeneratorGroup = 20940016,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["994"] = {  GeneratorGroup = 20960005,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 2,  Exp = 10,  SpCost = 2,  },
["995"] = {  GeneratorGroup = 20960002,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 2,  Exp = 10,  SpCost = 2,  },
["996"] = {  GeneratorGroup = 20960004,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 2,  Exp = 10,  SpCost = 2,  },
["997"] = {  GeneratorGroup = 20960005,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 2,  Exp = 10,  SpCost = 2,  },
["998"] = {  GeneratorGroup = 20960002,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 2,  Exp = 10,  SpCost = 2,  },
["999"] = {  GeneratorGroup = 20960004,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 2,  Exp = 10,  SpCost = 2,  },
["1000"] = {  GeneratorGroup = 20500088,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1001"] = {  GeneratorGroup = 20500089,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1002"] = {  GeneratorGroup = 20500090,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1003"] = {  GeneratorGroup = 20500091,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1004"] = {  GeneratorGroup = 20500092,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1005"] = {  GeneratorGroup = 20500109,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1006"] = {  GeneratorGroup = 20500110,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1007"] = {  GeneratorGroup = 20500103,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1008"] = {  GeneratorGroup = 20500104,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1009"] = {  GeneratorGroup = 20500105,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1010"] = {  GeneratorGroup = 20500106,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1011"] = {  GeneratorGroup = 20500108,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1012"] = {  GeneratorGroup = 20500111,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1013"] = {  GeneratorGroup = 20500112,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1014"] = {  GeneratorGroup = 20900021,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1015"] = {  GeneratorGroup = 20500113,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1016"] = {  GeneratorGroup = 20500113,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1017"] = {  GeneratorGroup = 20500088,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1018"] = {  GeneratorGroup = 20500113,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1019"] = {  GeneratorGroup = 20500089,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1020"] = {  GeneratorGroup = 20500090,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1021"] = {  GeneratorGroup = 20900020,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["1022"] = {  GeneratorGroup = 20900020,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["1023"] = {  GeneratorGroup = 20900020,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["1024"] = {  GeneratorGroup = 20500114,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1025"] = {  GeneratorGroup = 20500115,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["1026"] = {  GeneratorGroup = 20500115,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["1027"] = {  GeneratorGroup = 20500115,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["1028"] = {  GeneratorGroup = 20500116,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["1029"] = {  GeneratorGroup = 20500116,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 3,  Exp = 10,  SpCost = 2,  },
["1030"] = {  GeneratorGroup = 20500116,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1031"] = {  GeneratorGroup = 20500117,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1032"] = {  GeneratorGroup = 20500121,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 1,  Exp = 10,  SpCost = 2,  },
["1033"] = {  GeneratorGroup = 20500119,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 5,  Exp = 10,  SpCost = 2,  },
["1034"] = {  GeneratorGroup = 20500120,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
["1035"] = {  GeneratorGroup = 20500123,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = 1,  Exp = 50,  SpCost = 2,  },
["1036"] = {  GeneratorGroup = 20500124,  Weather = "-1",  DeleteWeather = -1,  LifeCycle = -1,  Exp = 10,  SpCost = 2,  },
} 
 CSV_resourceType = {}

function CSV_resourceType.GetValue(index, key)
	index = tostring(index)
	key = tostring(key)
    if resourceType[index] == nil then
        ZLog.Logger("Excel 获取表:resourceType  主键:".. index .." key:".. key.."出错!")
        return nil
    end
    if resourceType[index][key] == nil then
        ZLog.Logger("Excel 获取表: resourceType  主键:".. index .." key:".. key.."出错!")
        return nil
    end

    return resourceType[index][key]
end

function CSV_resourceType.GetAllKeys()
    local keys = {}
    for k in pairs(resourceType) do
        keys[#keys + 1] = k
    end
    if #keys == 0 then
        ZLog.Logger(string.format("Excel 获取表: 所有主键出错"))
        return nil
    end
    return keys
end