---创建者:zjy
---时间: 2019/10/23 9:54
--- 连接信息，网狐的机构体 TagBindParameter

---连接信息
---@class LinkInfo
ChatLinkInfo = {

}

---返回连接信息
---@return LinkInfo
function ChatLinkInfo:New()
    local t =  {
        ServerID = 0,
        ActiveTime = 0,
        LinkKind = 0,           --这里区分是客户端还是登录服的连接
    }
    setmetatable(t, self)
    self.__index = self
    return t
end

--- 重新加载类
function ChatLinkInfo:Reload(o)
    --- 重新刷一次元表,以便调用新定义方法，更新老方法
    G_SetMetaTable(o,self)
    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end