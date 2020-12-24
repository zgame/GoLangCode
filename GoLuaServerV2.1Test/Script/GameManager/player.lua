
--------------------------------------------------------------------------------------
---player的数据是玩家的游戏中的数据，临时的，退出不保存的数据
--------------------------------------------------------------------------------------

Player = {}
function Player:New(user)
    c = {
        User = user,  -- user数据

        TableID = Const.ROOM_CHAIR_NOBODY ,  -- 桌子id
        ChairID = Const.ROOM_CHAIR_NOBODY,   -- 椅子id

        IsRobot = false,            -- 是不是机器人
        ActivityBulletNum = 0,   --当前已经发射的子弹数量

        GameType = 0 ,     -- 游戏类型

        NetWorkState = true,   -- 网络状态正常
        NetWorkCloseTimer = 0 ,   -- 等待玩家断线重连的时间倒计时
    }
    setmetatable(c, self)
    self.__index = self
    return c
end

