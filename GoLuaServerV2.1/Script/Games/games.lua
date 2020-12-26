----------------------------------------------------------------
-----------------------------game管理房间和玩家-----------------
----------------------------------------------------------------



Game = Class:extend()
function Game:New(name, gameId)

    self.name = name
    self.gameId = gameId
    self.switch = true              -- 游戏是否开启

    self.allRoomList = {}                  -- 所有房间列表       key tableUid  ,value table          --- 要注意， key 不能用数字，因为占用内存太大， goperlua的问题
    self.allRoomNumber = 0                  -- 所有该游戏的房间数量
    self.tableUUID = 1                   -- tableUid 从1开始

    --GoRunTableAllList = {},                   -- 房间的run函数在里面                --- 要注意， key 不能用数字，因为占用内存太大， goperlua的问题
    --GameScore           = 0 ,                   --  游戏倍率
    --GameRoomInfo        = {},                   -- 游戏房间信息(GameRommInfo数据库表相关信息)
    --}
    --setmetatable(c,self)
    --self.__index = self
    --return c
end

function Game:Reload(self)
    --setmetatable(c, self)
    --self.__index = self
    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end

-----------------------------管理房间---------------------------

-- 创建房间，并启动它
function Game: CreateRoom(gameId)
    local room
    if gameId == Const.GameTypeCCC then
        room = CCCRoom(self.tableUUID, gameId)
        -- 房间开始
        room:InitRoom()
        --elseif
    end
    if room == nil then
        ZLog.Logger("Create room error , gameType" .. gameId)
        return nil
    end

    ZLog.Logger("创建了一个新的房间,type:" .. gameId)

    --增加该房间到总列表中
    self.allRoomList[tostring(self.tableUUID)] = room
    self.allRoomNumber = self.allRoomNumber + 1
    self.tableUUID = self.tableUUID + 1     -- table uuid 自增

    return room

end

-- 根据房间uid 返回房间的句柄
function Game.GetRoomByUID(game, roomId)
    return game.allRoomList[tostring(roomId)]
end

-- 房间回收
function Game.ReleaseRoom(game, roomId)
    if roomId ~= 1 then
        game.allRoomList[tostring(roomId)] = nil
        game.allRoomNumber = game.allRoomNumber - 1
        --self.GoRunTableAllList[tostring(roomId)] = nil
        --SqlDelGameState(self.GameTypeID, roomId)   -- 把记录房间状态的redis删掉
        ZLog.Logger("清理掉房间" .. roomId)
    else
        -- 第一个房间是保留着的，只是清理一下
        --local state ={}
        --state["FishNum"] = 0
        --state["BulletNum"] = 0
        --state["SeatArray"] = 0
        --SqlSaveGameState(self.GameTypeID, roomId, state)       -- mysql房间状态修改一下
    end
    collectgarbage()        -- 强制gc
end


----- 然后注册TableRun
--function Game:FindGoRoutineAndRegisterTableRun(roomId,func)
--    self.GoRunTableAllList[tostring(roomId)] = func  --注册TableRun函数
--end

--
---------------------------管理玩家-------------------------------------
--
--local function seat(room, player, seatId)
--    room:PlayerSeat(seatId, player)              --让玩家坐下.
--    player.roomId = room.roomId
--    player.chairId = seatId
--    --self:SendYouLoginToOthers(player, room)-- 发消息给同房间的其他玩家，告诉他们你登录了
--    return player
--end
----- 有玩家登陆游戏
--function Game.PlayerLoginGame(self,oldPlayer)
--    local player = GameServer.GetPlayerByUID(Player.UId(oldPlayer)) -- 把之前的玩家数据取出来
--    -- 如果玩家是断线重连的
--    if player ~= nil then
--        --找到之前有玩家在线
--        if oldPlayer.gameId == player.gameId then
--            -- 同一个游戏， 并且玩家状态是等待断线重连
--            --player.NetWorkState = true                      -- 网络恢复正常
--            --player.NetWorkCloseTimer = 0
--            print("把断线重连的player返回去， 玩家本来就坐在这里，不用同步信息给其他玩家， 就是反应他傻了一会后继续游戏了")
--            return player
--        else
--            -- 不是同一个游戏，或者有玩家在里面玩呢
--            -- player会被替换掉，那么之前的连接也到t掉才可以
--
--            -- 这里以后增加，t掉玩家的连接的功能
--        end
--    end
--
--    -- 不是断线重连的就重新建一个玩家数据
--    --player = Player:New(oldPlayer.User)
--    --player.GameType = oldPlayer.GameType            -- 设定游戏类型
--    player = oldPlayer
--    GameServer.SetAllPlayerList(Player.UId(player), player)  --创建好之后加入玩家总列表
--
--    --然后找一个有空位的房间让玩家加入游戏
--    for k, room in pairs(self.allRoomList) do
--        local seatId = BaseRoom.GetEmptySeatInTable(room)
--        if seatId > 0 then
--            print("有空座位")
--            room:InitRoom()    -- 看看是不是空房间，如果是，需要初始化
--            return seat(room,player,seatId)
--        end
--    end
--
--    --没有空座位的房间了，创建一个
--    print("没有空座位的房间了，创建一个吧,  score".. self.gameId)
--    local gameId = self.allRoomList["1"].gameId
--    local room = Game.CreateRoom(self, gameId)
--    local seatId = BaseRoom.GetEmptySeatInTable(room)  --获取空椅位
--    return seat(room,player,seatId)
--
--end


----玩家登出
function Game.PlayerLogOutGame(self,player)
    ZLog.Logger("玩家登出 "..Player.UId(player).. "    房间 "..player.roomId)
    local room = Game.GetRoomByUID(self,player.roomId)
    if room ~= nil then
        room:PlayerStandUp(player.chairId, player)        -- 玩家离开房间
        ZLog.Logger("玩家"..Player.UId(player).."离开房间 "..player.roomId.."椅子"..player.chairId)
    else
        ZLog.Logger("玩家登出时候房间为空"..player.roomId)
    end
end
