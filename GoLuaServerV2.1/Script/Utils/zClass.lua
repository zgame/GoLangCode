-- 用法：
--Point = Class:extend()
--function Point:new()
--
--Rect = Point:extend()
--function Rect:new(x, y, width, height)
--Rect.super.new(self, x, y)

--


local Class = {}
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

return Class