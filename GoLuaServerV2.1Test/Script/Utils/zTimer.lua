--------------------------------计时器系列-------------------------------------

ZTimer ={}

--- 自己设定定时器 模块名 | 回调函数的名字 | 间隔 毫秒数
function ZTimer.SetNewTimer(module,funcName,timer,func)
    -- funcName is string
    luaCallGoCreateNewTimer(module,funcName,timer)
end

--- 自己设定固定时间触发器  模块名 | 回调函数名字 | 触发的时间24小时制
function ZTimer.SetNewClockTimer(module,funcName,clock,func)
    luaCallGoCreateNewClockTimer(module,funcName,clock)
end