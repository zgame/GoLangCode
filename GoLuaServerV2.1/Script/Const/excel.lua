---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/12/6 10:53
---

------------------------------------------------------------------------------------------
--- csv操作相关:
--- 为什么通过这种方式调用属性，为了避免配置相关出错，相关属性不存在影响程序代码正常运营
--- 统一函数命名规则(GetExcel开头)
------------------------------------------------------------------------------------------








--- excel 获取数据内部接口
--- @param excel    csv对象
--- @param index    主键
--- @param key      属性名
--- @param name     表名
--- @return 该属性名的值
local function GetValue(excel,index,key,name)
    if excel[tostring(index)] == nil then
        Logger(string.format("Excel 获取表:%s 主键:%s 出错!", name, index))
        return nil
    end
    if excel[tostring(index)][tostring(key)] == nil then
        Logger(string.format("Excel 获取表:%s 主键:%s 列名:%s 出错!", name, index, key))
        return nil
    end

    return excel[tostring(index)][tostring(key)]
end
------------------------------------------------------------------------------------------
--- 获取配置表中的所有key
--- @param  excel 表对象
--- @param  name  表名
--- @return excel中的所有key的table
local function GetAllKeys(excel, name)
    local keys = {}
    for k in pairs(excel) do
        keys[#keys+1] = k
    end
    if #keys == 0 then
        Logger(string.format("Excel 获取表:%s 所有主键出错", name))
        return nil
    end
    return keys
end


-------------------------------------------------------------------------------------------------------
-------------------------------------------------------------------------------------------------------
--------------------------------------------mgby_fish_sever.lua----------------------------------------
--- 获取mgby_fish_server数据对外接口
--- @param index    主键
--- @param key      属性名
--- @return 表mgby_fish_server主键index所对应对象数据中key属性的值
function GetExcelFishValue(index, key)
    return ""
end
--- 获取 mgby_fish_server.csv中所有的key
--- @return 表mgby_fish_server所有key组成的table
function GetExcelFishAllKeys()
    return GetAllKeys(FishServerExcel, "mgby_fish_server.lua")
end
--------------------------------------------mgby_item.lua----------------------------------------------
local ItemExcel = mgby_item.mgby_item
--- 获取mgby_item.csv数据对外接口
--- @param index    主键
--- @param key      属性名
--- @return 表mgby_item主键index所对应对象数据中key属性的值
function GetExcelItemValue(index, key)
    return GetValue(ItemExcel,index, key, "mgby_item.lua")
end
--- 获取 mgby_item.csv中所有的key
--- @return 表mgby_item所有key组成的table
function GetExcelItemAllKeys()
    return GetAllKeys(ItemExcel, "mgby_item.lua")
end
--------------------------------------------mgby_monster.lua----------------------------------------------
--- 小海兽配置
local XhsExcel = mgby_monster.mgby_monster
--- 获取mgby_monster.csv数据对外接口
--- @param index    主键
--- @param key      属性名
--- @return 表mgby_monster主键index所对应对象数据中key属性的值
function GetExcelXhsValue(index, key)
    return GetValue(XhsExcel,index, key, "mgby_monster.lua")
end
--- 获取 mgby_monster.csv中所有的key
--- @return 表mgby_monster所有key组成的table
function GetExcelXhsAllKeys()
    return GetAllKeys(XhsExcel, "mgby_monster.lua")
end

--------------------------------------------mgby_vip.lua----------------------------------------------
--- vip配置表
local MgbyVipExcel = mgby_vip.mgby_vip
--- 获取mgby_vip.csv数据对外接口
--- @param index    主键
--- @param key      属性名
--- @return 表mgby_vip主键index所对应对象数据中key属性的值
function GetExcelMgbyVipValue(index, key)
    return GetValue(MgbyVipExcel,index, key, "mgby_vip.lua")
end
--- 获取 mgby_vip.csv中所有的key
--- @return 表mgby_vip所有key组成的table
function GetExcelMgbyVipAllKeys()
    return GetAllKeys(MgbyVipExcel, "mgby_vip.lua")
end