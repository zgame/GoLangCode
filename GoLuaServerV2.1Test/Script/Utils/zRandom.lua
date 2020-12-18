ZRandom={}

--- 获取随机数
function ZRandom.GetRandom(min,max)
    math.randomseed(ZTime.GetOsTimeMillisecond())
    return math.random(min,max)  -- [min,max]
end

---获取百分比方法， 比如10几率， 那么小于等于10，返回true
function ZRandom.PercentRate(rate)
    local rr = ZRandom.GetRandom(1, 100)
    if rr <= rate then
        return true
    else
        return false
    end
end
