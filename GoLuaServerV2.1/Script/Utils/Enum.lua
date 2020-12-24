---创建者:zjy
---时间: 2019/10/23 10:02
---工具类

---直接使用类方法
---@class Util
Util = {
}


---   事例
---   local next = Util.GetEnumIndexGenerater() 默认值 0开始每次加1
--    print("-------------------------next--------------------",next())
--    print("-------------------------next--------------------",next())
--    这一次加5
--    print("-------------------------next--------------------",next(5))
---   local next2 = Util.GetEnumIndexGenerater(1,2) 默认值 1开始每次加2
--    print("-------------------------next--------------------",next2())
--    print("-------------------------next--------------------",next2())
--   这一次加5
--    print("-------------------------next--------------------",next2(5))
---获取枚举生成器

function Util.GetEnumIndexGenerater(initid,step)
    local CurrentIndex = initid or 0
    step = step or 1
    -- 初始化要减个步长 让初次调用能获取初始下标
    CurrentIndex = CurrentIndex - step
    return function(instep)
        instep = instep or step
        CurrentIndex = CurrentIndex + instep
        return CurrentIndex
    end
end

