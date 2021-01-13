
ZRandom={}
local ZRan = require("zRandom")        -- Json

--- 获取随机数
function ZRandom.GetRandom(min,max)
    math.randomseed(ZTime.GetOsTimeMillisecond())
    return math.random(min,max)  -- [min,max]
end

--随机 [0,num)
function ZRandom.GetInt(number)
    return ZRan.intN(number)
end

---获取百分比方法， 比如10几率， 那么小于等于10，返回true
function ZRandom.PercentRate(rate)
    return ZRan.percent(rate)
end

--- 获取浮点的随机数
function ZRandom.GetFloat(min,max,bit)
    local ranF ZRan.float(min,max)
    local nRet = tonumber(string.format('%.' .. bit .. 'f',ranF))
    --print("ranf " .. nRet)
    return nRet
end

-- 正态分布
function ZRandom.Normal()
    return ZRan.normal()
end


-- 指数分布
function ZRandom.Exp()
    return ZRan.exp()
end
-- 随机数值排列
function ZRandom.Perm(len)
    return ZRan.perm(len)
end

-- 通过权重获得元素
function ZRandom.GetList(weightList)
    local ran = ZRandom.GetRandom(1,100)
    for index,rate in ipairs(weightList) do
        if ran <= tonumber(rate) then
            return index
        end
    end
end