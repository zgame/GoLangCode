--------------------------------------------------------------------------------------
--- 游戏房间的基类
--------------------------------------------------------------------------------------

-- 房间
BaseRoom = Class:extend()

function BaseRoom:New()

    self.gameId = 0                    -- 游戏类型ID
    self.roomId = 0                    -- 房间ID
    self.tableMax = 0                    -- 房间最大容纳玩家数量

    self.userSeatArray = {}                -- 椅子[座椅对应玩家uid的哈希表 ， key ： seatID (1,2,3,4)   ，value： player]
    self.userSeatArrayNumber = 0           -- 房间上有几个玩家， 记住，这里不能用#UserSeatArray, 因为有可能中间有椅子是空的，不连续的不能用#， 本质UserSeatArray是map ；  也不能遍历， 慢
    self.lastRunTime = 0                   -- 循环周期时间

end

-- 重载房间
function BaseRoom:Reload(o)
    --Logger("调用了BaseTable:Reload()")
    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end

-- 房间的主循环
function BaseRoom:RunRoom()
    --print("房间基类主循环")
end

----------------------- 房间操作 ---------------------------------
--清理房间
function BaseRoom:ClearTable()
    self.userSeatArray = {}     --  seatID    player
    self.userSeatArrayNumber = 0
end
--判断房间是有人，还是空房间
function BaseRoom.CheckTableEmpty(self)
    if self.userSeatArrayNumber > 0 then
        return false
    end
    return true -- 空房间
end

--获取房间的空座位, 返回座椅的编号，从0开始到tableMax， 如果返回-1说明满了-
function BaseRoom:GetEmptySeatInTable()
    for i = 1, self.tableMax do
        if self.userSeatArray[i] == nil then
            return i
        end
    end
    return -1
end


----------------------- 玩家 ---------------------------------
--玩家坐到椅子上
function BaseRoom:PlayerSeat(seatId, player)
    self.userSeatArray[seatId] = player
    self.userSeatArrayNumber = self.userSeatArrayNumber + 1   -- 房间上玩家数量增加
end

--玩家离开椅子
function BaseRoom:PlayerStandUp(seatId, player)
    ZLog.Logger(player:UId() .. "离开房间" .. player.roomId .. "椅子" .. player.chairId .. "self.roomId" .. self.gameId)
    -- 保存玩家基础数据
    --SaveUserBaseData(player.User)

    GameServer.SetAllPlayerList(player:UId(), nil)         -- 清理掉游戏管理的玩家总列表
    self.userSeatArray[seatId] = nil                -- 清理掉房间的玩家列表
    self.userSeatArrayNumber = self.userSeatArrayNumber - 1  -- 房间上玩家数量减少
    player.roomId = Const.ROOM_CHAIR_NOBODY
    player.chairId = Const.ROOM_CHAIR_NOBODY

    --如果是空房间的话，清理一下房间
    if self:CheckTableEmpty() then
        self:ClearTable()
        local game = GameServer.GetGameByID(self.gameId)
        Game.ReleaseRoom(game,self.roomId)    --回收房间
    end
end

----------------------- 同步消息 ---------------------------------
--给桌上的所有玩家同步消息
function BaseRoom:SendMsgToAllUsers(mainCmd, subCmd, sendCmd)
    for _, player in pairs(self.userSeatArray) do
        if player ~= nil and player.netWorkState then
            local result = NetWork.SendToUser(player:UId(), mainCmd, subCmd, sendCmd, nil, 0)       -- 注意，这里因为是群发，所以token标记是0，就是不需要
            if not result then
                -- 发送失败了，玩家网络中断了
                --player.NetWorkState = false
                --player.NetWorkCloseTimer = GetOsTimeMillisecond()
                self:PlayerStandUp(player.chairId, player)
            end
        end
    end
end

--给桌上的其他玩家同步消息
function BaseRoom:SendMsgToOtherUsers(userId, sendCmd, mainCmd, subCmd)
    for _, player in pairs(self.userSeatArray) do
        if player ~= nil and userId ~= player:UId() and player.netWorkState then
            NetWork.SendToUser(player:UId(), mainCmd, subCmd, sendCmd, nil, 0)       -- 注意，这里因为是群发，所以token标记是0，就是不需要
        end
    end
end
