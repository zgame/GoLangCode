---
-- @function : 获取table的字符串格式内容，递归
-- @tab ： table
-- @ind ： 不 用 传 此 参 数 ， 递 归 用 （ 前 缀 格 式 （ 空 格 ） ）
-- @return : format string of the table
function dumpTab(tab, ind)
    if (tab == nil) then return "nil" end;
    local str = "{";
    if (ind == nil) then ind = "  "; end;
    --//each of table
    for k, v in pairs(tab) do
        --//key
        if (type(k) == "string") then
            k = tostring(k) .. " = ";
        else
            k = "[" .. tostring(k) .. "] = ";
        end; --//end if
        --//value
        local s = "";
        if (type(v) == "nil") then
            s = "nil";
        elseif (type(v) == "boolean") then
            if (v) then s = "true"; else s = "false"; end;
        elseif (type(v) == "number") then
            s = v;
        elseif (type(v) == "string") then
            s = "\"" .. v .. "\"";
        elseif (type(v) == "table") then
            s = dumpTab(v, ind .. "  ");
            s = string.sub(s, 1, #s - 1);
        elseif (type(v) == "function") then
            s = "function : " .. v;
        elseif (type(v) == "thread") then
            s = "thread : " .. tostring(v);
        elseif (type(v) == "userdata") then
            s = "userdata : " .. tostring(v);
        else
            s = "nuknow : " .. tostring(v);
        end; --//end if
        --//Contact
        str = str .. "\n" .. ind .. k .. s .. " ,";
    end --//end for
    --//return the format string
    local sss = string.sub(str, 1, #str - 1);
    if (#ind > 0) then ind = string.sub(ind, 1, #ind - 2) end;
    sss = sss .. "\n" .. ind .. "}\n";
    return sss; --string.sub(str,1,#str-1).."\n"..ind.."}\n";
end

; --//end function

--//网摘,直接打印到屏幕
function printTable(t, n)
    if "table" ~= type(t) then
        return 0;
    end
    n = n or 0;
    local str_space = "";
    for i = 1, n do
        str_space = str_space .. "  ";
    end
    print(str_space .. "{");
    for k, v in pairs(t) do
        local str_k_v
        if (type(k) == "string") then
            str_k_v = str_space .. "  " .. tostring(k) .. " = ";
        else
            str_k_v = str_space .. "  [" .. tostring(k) .. "] = ";
        end
        if "table" == type(v) then
            print(str_k_v);
            printTable(v, n + 1);
        else
            if (type(v) == "string") then
                str_k_v = str_k_v .. "\"" .. tostring(v) .. "\"";
            else
                str_k_v = str_k_v .. tostring(v);
            end
            print(str_k_v);
        end
    end
    print(str_space .. "}");
end

----//test code
--b = { [printTable] = 1, [dumpTab] = 33 }
--a = { inter = { ext = 1000, uuid = "432-43243-43124324", inter2 = { ext = 1000, uuid = "432-43243-43124324", [printTable] = 2 }, "anil", nil }, "fda", fdsaf, true };
--
--local sa = dumpTab(b);
--print(sa);
--
--sa = dumpTab(a);
--print(sa);
--
--printTable(a)
--printTable(b)

