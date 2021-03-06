

ZTable = {}
-- 获取table有多少个元素
function ZTable.Len(table)
    local num = 0
    for k,v in pairs(table) do
        num = num + 1
    end
    return num
end

--
--function ToStringEx(value)
--    if type(value)=='table' then
--        return TableToStr(value)
--    elseif type(value)=='string' then
--        return "\'"..value.."\'"
--    else
--        return tostring(value)
--    end
--end
--
--
---- table 转换成 string
--function TableToStr(t)
--    if t == nil then return "" end
--    local retStr = "{"
--
--    local i = 1
--    for key,value in pairs(t) do
--        local signal = ","
--        if i==1 then
--            signal = ""
--        end
--
--        if key == i then
--            retStr = retStr ..signal..ToStringEx(value)
--        else
--            if type(key)=='number' or type(key) == 'string' then
--                retStr = retStr ..signal..'['..ToStringEx(key).."]="..ToStringEx(value)
--            else
--                if type(key)=='userdata' then
--                    retStr = retStr ..signal.."*s"..TableToStr(getmetatable(key)).."*e".."="..ToStringEx(value)
--                else
--                    retStr = retStr ..signal..key.."="..ToStringEx(value)
--                end
--            end
--        end
--
--        i = i+1
--    end
--
--    retStr = retStr .."}"
--    return retStr
--end
--
--
---- string 转成 table
--function StrToTable(str)
--    if str == nil or type(str) ~= "string" then
--        return
--    end
--
--    return loadstring("return " .. str)()
--end


--- 从权重数组中选择对应的索引下标
function TableWeightSelect(weightArray)
    local nCurWeight = 0
    local nTotalWeight = 0

    for i=1, #weightArray do
        nTotalWeight = nTotalWeight + weightArray[i]
    end

    local random = ZRandom.GetRandom(1,nTotalWeight)
    for i=1, #weightArray do
        nCurWeight = nCurWeight + weightArray[i]
        if nCurWeight >= random then
            return i
        end
    end
    return 0
end