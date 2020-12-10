-- 用法：
--Point = Class:extend()
--function Point:new()
--
--Rect = Point:extend()
--function Rect:new(x, y, width, height)
--  Rect.super.new(self, x, y)

-- 用法说明： 可以使用：做为成员函数，优点：标准的类的用法， 缺点：idea的很多功能不能使用，不能重构，不能定位，只能搜文本，容易出错
-- 用法说明： 可以使用. 做为模块函数，优点：idea的很多功能能使用，能重构，能定位，不容易出错  缺点： 写起来代码量多一些，继承关系中调用模块函数要指定才行

-- 推荐用法： 类定义使用：定义成员函数， 调用的时候采用 Class.Function(object,...) 这种方式来写


--local Class = {}
Class = {}
Class.__index = Class

function Class:new()
end

function Class:extend()
    local cls = {}
    for k, v in pairs(self) do
        if k:find("__") == 1 then
            cls[k] = v
        end
    end
    cls.__index = cls
    cls.super = self
    setmetatable(cls, self)
    return cls
end

function Class:implement(...)
    for _, cls in pairs({ ... }) do
        for k, v in pairs(cls) do
            if self[k] == nil and type(v) == "function" then
                self[k] = v
            end
        end
    end
end

function Class:is(T)
    local mt = getmetatable(self)
    while mt do
        if mt == T then
            return true
        end
        mt = getmetatable(mt)
    end
    return false
end

function Class:__tostring()
    return "Object"
end

function Class:__call(...)
    local obj = setmetatable({}, self)
    obj:new(...)
    return obj
end

--return Class