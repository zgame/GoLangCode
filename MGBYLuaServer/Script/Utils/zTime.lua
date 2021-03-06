---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/6 10:56
---



--- 获取系统时间，毫秒级别 ,  注意 os.time() 是秒级别
--- @return 获取系统时间毫秒级别
function GetOsTimeMillisecond()
    return luaCallGoGetOsTimeMillisecond()
end

--- 获取时间来显示
function GetOsDateNow(time)
    if time == nil then
        return os.date("%Y-%m-%d %H:%M:%S",os.time())
    else
        return os.date("%Y-%m-%d %H:%M:%S",time)
    end
end


--- 获取本月来显示
function GetOsMonthNow()
    return os.date("%Y-%m",os.time())
end

--- 获取今天来显示
function GetOsDayNow()
    return os.date("%Y-%m-%d",os.time())
end

--- 获取今天来显示
--- @return 当天的时间字符串格式(20191030)
function GetOsDayNowEx()
    return os.date("%Y%m%d",os.time())
end

--- 获取现在的时间来显示
function GetOsDateTimeNow()
    return os.date("%H:%M:%S",os.time())
end


--- 获取两个时间的天数差
function GetTwoTimesDays(time1,time2)

    -- 这里做了一个操作， 把时间给去掉了， 这样单纯的计算天数的差值
    local day1 = os.date("%Y-%m-%d 00:00:00", time1)
    local day2 = os.date("%Y-%m-%d 00:00:00", time2)
    local time11 = GetTimeFromString(day1)
    local time22 = GetTimeFromString(day2)

    local time = math.abs(time11 - time22)
    return  math.floor(time / (60*60*24))
end


-- 获取跟当前时间的天数差
function GetTimesDaysByString(timeString)
    local today_last
    if timeString == "" then        -- 如果没有上次时间
        today_last = 1      -- 如果没有记录， 那么就认为不连续
    else
        today_last = GetTwoTimesDays(GetTimeFromString(timeString) , os.time())     -- 有签到记录，那么就用计算的值
    end
    return today_last       -- 如果是今天，返回0
end


--------------------------------字符串处理------------------------------------
--- 从日期字符串中获取时间
function GetAllTimeFormString(dateTimeString)
    local Y = string.sub(dateTimeString,1,4)
    local M = string.sub(dateTimeString,6,7)
    local D = string.sub(dateTimeString,9,10)
    local H = string.sub(dateTimeString,12,13)
    local MM = string.sub(dateTimeString,15,16)
    local SS = string.sub(dateTimeString,18,19)
    local time = os.time({year=tonumber(Y),month=tonumber(M),day=tonumber(D),hour=tonumber(H),min=tonumber(MM),sec=tonumber(SS)})
    local toMonth = string.format("%s-%s",Y,M)
    local toDay = string.format("%s-%s-%s",Y,M,D)

    --print("输入："..dateTimeString)
    --print("Y"..Y)
    --print("M"..M)
    --print("D"..D)
    --print("H"..H)
    --print("MM"..MM)
    --print("SS"..SS)
    --print("time"..time)
    --print(os.date("%Y-%m-%d %H:%M:%S",time))

    return Y,M,D,H,MM,SS,time,toMonth,toDay
end

-- 从日期字符串中获取天的数据
function GetDayFromString(dateTimeString)
    local Y,M,D,H,MM,SS,time,toMonth,toDay = GetAllTimeFormString(dateTimeString)
    return toDay
end

-- 从日期字符串中获取时间
function GetTimeFromString(dateTimeString)
    local Y,M,D,H,MM,SS,time,toMonth,toDay = GetAllTimeFormString(dateTimeString)
    return time
end

-- 从日期字符串中获取小时
function GetHourFromString(dateTimeString)
    local Y,M,D,H,MM,SS,time,toMonth,toDay = GetAllTimeFormString(dateTimeString)
    return tonumber(H)
end

--- 获取今天初的时间戳
--- @return 今天凌晨的时间戳 单位:秒
function GetTodayStartTime()
    local todayTimeStr = os.date("%Y-%m-%d 00:00:00", os.time())
    return GetTimeFromString(todayTimeStr)
end

--- 获取今天过多少秒的时间
--- @param tSec 秒数
--- @return 时间戳
function GetNextTimeFromTodayStartTime(tSec)
    return GetTodayStartTime() + tSec
end

--- 获取当前时间距离明天凌晨的时间间隔(单位:秒)
--- @return 到明天凌晨的时间间隔单位秒
function GetIntervalBetweenNowAndTomorrow()
    return GetNextTimeFromTodayStartTime(24*60*60) - os.time()
end
--------------------------------计时器系列-------------------------------------

--- 自己设定定时器 回调函数的名字 和 间隔秒数
function SetNewTimer(funcName,timer)
    -- funcName is string
    luaCallGoCreateNewTimer(funcName,timer)
end

--- 自己设定固定时间触发器  回调函数名字 和 触发的时间24小时制
function SetNewClockTimer(funcName,clock)
    luaCallGoCreateNewClockTimer(funcName,clock)
end