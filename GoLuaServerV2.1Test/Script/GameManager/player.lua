
--------------------------------------------------------------------------------------
---player的数据是玩家的游戏中的数据，临时的，退出不保存的数据
--------------------------------------------------------------------------------------

Player = {}
function Player:New(user)
    local c = {
        User = user,  -- user数据

        TableID = Const.ROOM_CHAIR_NOBODY ,  -- 桌子id
        ChairID = Const.ROOM_CHAIR_NOBODY,   -- 椅子id

        GameType = 0 ,     -- 游戏类型

    }
    setmetatable(c, self)
    self.__index = self
    return c
end

