---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.
--- DateTime: 2019/10/29 11:14
---这里放通用的关于string的接口

---------------将字符串以某个自定的标记，拆分成字符串数组---------------

function SplitString(sBaseString,sSpeString,tContString)
    tContString = {};
    --sBaseString = string.format("%s%s",sBaseString,"SplitStringEnd")
    --sBaseString = string.gsub(sBaseString,)
    local dwStartPos = 1;
    local dwNextPos = 1;
    local dwSubEndPos = 1;
    --local sSubString = "";
    local nLength = string.len(sBaseString);
    local nSpeLength = string.len(sSpeString);
    --print("SplitString:sBaseString="..sBaseString..",sSpeString="..sSpeString..",nLength="..nLength..",nSpeLength="..nSpeLength)
    --如果标记位为空字符串，则返回sBaseString
    if nSpeLength < 1  then
        table.insert(tContString,sBaseString);
        return tContString;
    end
    --如果基础字符串长度小于等于1，分情况返回{}or{sBaseString}
    if nLength <= dwStartPos and string.find(sBaseString,sSpeString,1,true)  then---sBaseString=sSpeString
    return tContString;
    elseif nLength <= dwStartPos and not string.find(sBaseString,sSpeString,1,true)  then---sBaseString~=sSpeString
    table.insert(tContString,sBaseString);
        return tContString;
    end
    while (dwStartPos < nLength) do
        dwNextPos,dwSubEndPos = string.find(sBaseString,sSpeString,dwStartPos,true);
        --print("dwNextPos="..(dwNextPos or 0),"dwSubEndPos="..(dwSubEndPos or 0));
        if dwNextPos == nil then
            table.insert(tContString,string.sub(sBaseString,dwStartPos));
            dwStartPos = nLength;
        elseif dwNextPos > dwStartPos then
            local sSubContString = string.sub(sBaseString,dwStartPos,dwNextPos - 1);
            --print("sSubContString="..sSubContString)
            table.insert(tContString,sSubContString);
            dwStartPos = dwNextPos + nSpeLength;
        end
    end
    --printTable(tContString,0,"tContString")
    return tContString;
end


--- 功能: 按照正则表达式规则拆分string字符串

function StringSplitToStringArrayBYZZ( str,reps )
    local resultStrList = {}
    string.gsub(str,'[^'..reps..']+',function ( w )
        table.insert(resultStrList,w)
    end)
    return resultStrList
end

--- 功能: 拆分string字符串

function StringSplitToStringArray(str, delim, maxNb)
    if string.find(str, delim, 1, true) == nil then
        return { str }
    end

    if maxNb == nil or maxNb < 1 then
        maxNb = 0
    end

    local result = {}
    local pat = "(.-)" .. delim .. "()"
    local nb = 0
    local lastPos
    for part, pos in string.gmatch(str, pat) do
        nb = nb + 1
        result[nb] = part
        lastPos = pos
        if nb == maxNb then break end
    end

    if nb ~= maxNb then
        for i = nb, maxNb do
            result[nb + 1] = string.sub(str, lastPos)
        end
    end

    return result
end

--- 功能: 拆分数字字符串
function StringSplitToNumberArray(str, delim, maxNb)
    local result = {}
    if str == nil or string.len(str) == 0 or string.gmatch(str, delim) == nil then
        -- 拆分数量存在,则返回maxNb大小默认值为0的数组
        if maxNb then
            for i = 1, maxNb do
                result[i] = 0
            end
        end
        return result
    end

    if maxNb == nil or maxNb < 1 then
        maxNb = 0
    end

    local result = {}
    local pat = "(.-)" .. delim .. "()"
    local nb = 0
    local lastPos
    for part, pos in string.gmatch(str, pat) do
        nb = nb + 1
        result[nb] = tonumber(part)
        lastPos = pos
        if nb == maxNb then break end
    end


    if nb ~= maxNb then
        for i = nb, maxNb do
            result[i] = 0
        end
    end

    return result
end
