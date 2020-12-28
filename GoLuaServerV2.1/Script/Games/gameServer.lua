------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------
--- 游戏管理器， 管理很多游戏， 每个游戏再去管理自己的房间和玩家
--- 新添加游戏步骤：1, 定义类型  2 AddGame  3 在games.lua里面定义根据不同的游戏定义不同的房间 4 创建一个新游戏的房间.lua文件
------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------

GameServer = {}


-----------------------------------游戏服务器入口点-------------------------------------
--增加一个游戏， 指定这个游戏的类型， 并且创建一个房间，并启动房间逻辑
local function addGame(name, gameId)
    if GameServer.GetGameByID(gameId) ~= nil then
        ZLog.Logger("游戏类型[" .. gameId .. "已经添加过了，不用重复添加")
        return
    end

    local game = Game(name, gameId)
    -- 加入到游戏总列表中
    GameServer.SetAllGamesList(gameId, game)

    --Logger("--------------AddGame--------------------------")
    Game.CreateRoom(game, gameId)
    --game.GameScore = gameScore
end

-- 游戏服务器入口点
function GameServer.Start()
    --Logger("--------------------注册中心服----------------------------")
    --ServerMainServer = LuaNetWorkConnectOtherServer(Setting.ConstMainCenterServer)  -- 申请连接协调服务器，并 把serverId保存下来， 以后发送消息用
    --print("协调服 serverId ",ServerMainServer)

    print("-------------------  添加主循环  ------------------------------")
    ZTimer.SetNewTimer("GameServer", "RunGamesRooms", 2000, GameServer.RunGamesRooms)

    print("-------------------  添加游戏  ------------------------------")
    addGame("沙石镇", Const.GameTypeCCC)

    --SetNewTimer("NewTimerBy2Second",2 * 1000)      -- lua 自己设定计时器
    --SetNewClockTimer("NewTimerByAfternoon4", 16)    -- lua 自己设定的固定时间定时器，  16:00

    --Logger("--------------------注册协调服----------------------------")
    --ServerIDofCorrespondServer = LuaNetWorkConnectOtherServer(ConstServerAddressCorrespondServer)  -- 申请连接协调服务器，并 把serverId保存下来， 以后发送消息用
    --print("协调服 serverId ",ServerIDofCorrespondServer)
    --Logger("--------------------注册日志服----------------------------")

    --ServerIDofLogServer = LuaNetWorkConnectOtherServer(ConstServerAddressLogServer)                 -- 申请连接服务器，并 把serverId保存下来， 以后发送消息用
    --print("日志服 serverId ",ServerIDofLogServer)

    --Logger("--------------StartGamesServers  End--------------------------")
    --Logger("ServerIP_Port:"..ServerIP_Port)

end

-----------------------------------玩家列表管理-------------------------------------
local RedisDirAllPlayersUUID = "CCC:AllUer_UUID:"                         -- 所有玩家UUID
local function GetAllPlayersUUID(num)
    --return RedisAddNumber(RedisDirAllPlayersUUID.."BY_UUID" ,"BY_UUID",num)
    local dir = RedisDirAllPlayersUUID .. "CCC_UUID"
    local key = "CCC_UUID"
    local redis_lua_str = [[
    local r = redis.call('hget',"%s","%s")
       if r ~= false then
            r = r + %d
       else
            r = 1000000001
       end
    redis.call('hset',"%s","%s", r)
    return r
    ]]
    redis_lua_str = string.format(redis_lua_str, dir, key, num, dir, key)
    return Redis.RunLuaScript(redis_lua_str, "RedisMultiProcessGetAllPlayersUUID")
end

--有一个新的玩家注册了，那么给他分配一个UID
function GameServer.GetLastUserID()
    local r = 1     -- math.random(1, 4)        --返回[1,4]的随机整数
    local uuid = GetAllPlayersUUID(r)     -- 分布式申请UUID
    --Logger("给玩家分配新uid  ALLUserUUID "..uuid)
    return uuid
end


-- 根据user uid 返回user的句柄
function GameServer.GetPlayerByUID(uId)
    return GlobalVar.AllPlayerList[tostring(uId)]
end

function GameServer.SetAllPlayerList(userId, value)
    GlobalVar.AllPlayerList[tostring(userId)] = value
    if value == nil then
        GlobalVar.AllPlayerListNumber = GlobalVar.AllPlayerListNumber - 1   -- 玩家人数减少
    else
        GlobalVar.AllPlayerListNumber = GlobalVar.AllPlayerListNumber + 1   -- 玩家人数增加
    end
    ZLog.Logger("在线玩家数量" .. tostring(GlobalVar.AllPlayerListNumber))
end


-----------------------------------游戏列表管理-------------------------------------
function GameServer.GetRoomClass(gameId)
    local switch = {}
    switch[Const.GameTypeCCC] = CCCRoom
    return  switch[gameId]
end


--通过gameID获取是哪个游戏
function GameServer.GetGameByID(gameId)
    return GlobalVar.AllGamesList[tostring(gameId)]
end

function GameServer.SetAllGamesList(gameId, value)
    GlobalVar.AllGamesList[tostring(gameId)] = value
end

function GameServer.GetGameByUserId(userId)
    local player = GameServer.GetPlayerByUID(userId)
    if player ~= nil then
        local game = GameServer.GetGameByID(player.gameId)
        if game ~= nil then
            return game
        else
            ZLog.Logger("GetGameByID 游戏为空" .. player.gameId)
        end
    else
        ZLog.Logger("GetPlayerByUID 玩家为空" .. userId)
    end
    return nil
end
-----------------------------------房间-------------------------------------

function GameServer.GetRoomByUserId(userId)
    local player = GameServer.GetPlayerByUID(userId)
    if player ~= nil then
        local game = GameServer.GetGameByID(player.gameId)
        if game ~= nil then
            local room = Game.GetRoomById(game, player.roomId)
            if room ~= nil then
                return room
            else
                ZLog.Logger("GetRoomById room为空" .. player.roomId)
            end
        else
            ZLog.Logger("GetGameByID 游戏为空" .. player.gameId)
        end
    else
        ZLog.Logger("GetPlayerByUID 玩家为空" .. userId)
    end
    return nil
end

function GameServer.Login(gameId,player)
    local game = GameServer.GetGameByID(gameId)
    if game == nil then
        ZLog.Logger("没有找到游戏类型".. gameId)
        return false
    end
    player.gameId = gameId
    Game.PlayerLoginGame(game,player)
    return true
end

-----------------------------------定时器run-------------------------------------
-- 遍历所有的列表，然后依次run,  改为服务器自己创建定时器处理
function GameServer.RunGamesRooms()
    for _, game in pairs(GlobalVar.AllGamesList) do
        for _, room in pairs(game.allRoomList) do
            room:RunRoom() -- 执行注册的函数，table run
        end
    end
end
