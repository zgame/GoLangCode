

-----------------------------------------------------------------
--- Lua 位运算
-----------------------------------------------------------------
local Bit32 = require("bit32")                -- 位运算

-- 与
function Bit32And(a,b)
   return  Bit32.band(a,b)
end


-- 非
function Bit32Not(a)
    return  Bit32.bnot(a)
end


-- 或
function Bit32Or(a,b)
    return  Bit32.bor(a,b)
end


-- 异或
function Bit32Xor(a,b)
    return Bit32.bxor(a,b)
end


-- 左移
function Bit32LShift(a,b)
    return Bit32.lshift(a,b)
end


-- 右移
function Bit32RShift(a,b)
    return Bit32.rshift(a,b)
end

