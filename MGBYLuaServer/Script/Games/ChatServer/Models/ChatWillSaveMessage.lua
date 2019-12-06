---创建者:zjy
---时间: 2019/10/22 19:10
---聊天消息緩存结构体


---@class ChatWillSaveMessage 緩存聊天消息
ChatWillSaveMessage = {
    --UserID = 0,            ---用戶id
    --FriendID = 0,          ---好友id
    --EmotionID = 0,         ---表情id
    --sOriginalContent = "", --- 原始內容
}

function ChatWillSaveMessage:New(uid,fid,msgsize,content)
   local t = {
        UserID = uid,            ---用戶id
        FriendID = fid,          ---好友id
        EmotionID = msgsize,         ---长度
        sOriginalContent = content , --- 原始內容
    }
    setmetatable(t, self)
    self.__index = self
    return t
end


function ChatWillSaveMessage:Reload(o)
    --- 重新刷一次元表,以便调用新定义方法，更新老方法
   G_SetMetaTable(o,self)
    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end