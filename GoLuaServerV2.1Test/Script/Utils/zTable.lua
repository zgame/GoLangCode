

zTable={}

-- 获取table有多少个元素
function zTable.GetTableLen(table)
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