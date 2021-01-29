----------------------------------------------------------------
-----------------------------game管理房间和玩家-----------------
----------------------------------------------------------------

Game = Class:extend()
function Game:New(name, gameId)

    self.name = name
    self.gameId = gameId
    self.switch = true              -- 游戏是否开启

    self.allRoomList = {}                  -- 所有房间列表       key tableUid  ,value table          --- 要注意， key 不能用数字，因为占用内存太大， go per lua的问题
    self.allRoomNumber = 0                  -- 所有该游戏的房间数量
end

function Game:Reload(c)
    setmetatable(c, self)
    self.__index = self
    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end

-----------------------------管理房间---------------------------
function Game:GetUUID()
    for i = 1, self.allRoomNumber + 10  do
        if Game.GetRoomById(self, i) == nil then
            return i
        end
    end
end

-- 创建房间，并启动它
function Game: CreateRoom(gameId)

    local roomClass = GameServer.GetRoomClass(gameId)
    local newRoomId = Game.GetUUID(self)
    local room = roomClass(newRoomId, gameId)
    room:InitRoom()

    if room == nil then
        ZLog.Logger("Create room error , gameType" .. gameId)
        return nil
    end

    ZLog.Logger("创建了一个新的房间,type:" .. gameId)

    --增加该房间到总列表中
    self.allRoomList[tostring(room.roomId)] = room
    self.allRoomNumber = self.allRoomNumber + 1

    return room

end

-- 根据房间uid 返回房间的句柄
function Game:GetRoomById(roomId)
    return self.allRoomList[tostring(roomId)]
end

-- 房间回收
function Game.ReleaseRoom(gameId,roomId)
    if roomId ~= 1 then
        local game = GameServer.GetGameByID(gameId)
        game.allRoomList[tostring(roomId)] = nil
        game.allRoomNumber = game.allRoomNumber - 1
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

---------------------------管理玩家-------------------------------------

--- 有玩家登陆游戏
function Game:PlayerLoginGame(oldPlayer)
    local player = GameServer.GetPlayerByUID(Player.UId(oldPlayer)) -- 把之前的玩家数据取出来
    -- 如果玩家是断线重连的
    if player ~= nil then
        --找到之前有玩家在线
        if oldPlayer.gameId == player.gameId then
            -- 同一个游戏， 并且玩家状态是等待断线重连
            --player.NetWorkState = true                      -- 网络恢复正常
            --player.NetWorkCloseTimer = 0
            print("把断线重连的player返回去， 玩家本来就坐在这里，不用同步信息给其他玩家， 就是反应他傻了一会后继续游戏了")
            printTable(player)
            print(player.roomId)
            print(player.chairId)
            return player
        else
            -- 不是同一个游戏，或者有玩家在里面玩呢
            -- player会被替换掉，那么之前的连接也到t掉才可以

            -- 这里以后增加，t掉玩家的连接的功能
        end
    end

    -- 不是断线重连的就重新建一个玩家数据
    player = oldPlayer
    --然后找一个有空位的房间让玩家加入游戏
    for _, room in pairs(self.allRoomList) do
        local chairId = room:GetEmptySeatInTable()
        if chairId > 0 then
            --print("有空座位")
            --room:InitRoom()    -- 看看是不是空房间，如果是，需要初始化
            return room:PlayerSeat(chairId,player)
        end
    end

    --没有空座位的房间了，创建一个
    print("没有空座位的房间了，创建一个吧,  gameId:".. self.gameId)
    local room = self:CreateRoom(self.gameId)
    local chairId = room:GetEmptySeatInTable()  --获取空椅位
    return room:PlayerSeat(chairId,player)

end
