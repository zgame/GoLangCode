---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/10/12 14:56
---

-- 网络发送函数
-- LuaCallGoNetWorkSend(string)
-- 这个函数是go实现的，lua直接调用LuaCallGoNetWorkSend即可
function NetWorkSend(string)
    luaCallGoNetWorkSend(string)
end


-- 网络接收函数
function GoCallLuaNetWorkReceive(data)
    Logger("lua收到了消息：",data)
    luaCallGoNetWorkSend("lua想发送消息")
end