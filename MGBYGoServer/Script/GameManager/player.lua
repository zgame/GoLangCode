---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/2 16:37
---


--------------------------------------------------------------------------------------
--- player 的数据是玩家的游戏中的数据
--- 这里直接定义成员变量， 但是不写成员函数，就不用reload
--- 这里要注意一点：  当热更新的时候， 所有的已经创建好的player对象是存在的， 结构也是老的结构， 如果你增加了字段，修改了字段，需要进行reload的单独数据处理
--------------------------------------------------------------------------------------

Player = {}
function Player:New(user)
    c = {
        User = user,                    -- user数据
        ChatUser = {},                  -- chatUser数据

        TableID = TABLE_CHAIR_NOBODY ,  -- 桌子id
        ChairID = TABLE_CHAIR_NOBODY,   -- 椅子id

        IsRobot = false,                -- 是不是机器人
        ActivityBulletNum = 0,          -- 当前已经发射的子弹数量

        GameType = 0 ,                  -- 游戏类型

        NetWorkState = true,            -- 网络状态正常
        NetWorkCloseTimer = 0 ,         -- 等待玩家断线重连的时间倒计时

        --- 小海兽相关
        XhsStartBuffTimes = 0,          -- 触发暴击状态开始时间
        XhsIntervalBuffTimes = 0,       -- 暴击状态持续时间
        XhsLastSaveGameInfoTime = 0,    -- 上次存储游戏信息时间

        ChatWithFriendInterval = 0,     -- 与好友之间的聊天间隔
    }
    setmetatable(c, self)
    self.__index = self
    return c
end

function Player:Reload(c)
    setmetatable(c, self)
    self.__index = self

    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理

end