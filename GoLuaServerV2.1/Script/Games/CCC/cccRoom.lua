
CCCRoom = BaseRoom:extend()
function CCCRoom:New(roomId, gameId)
    -- 重新赋值某些属性值
    CCCRoom.super.New(self)
    self.gameId = gameId
    self.roomId = roomId
    self.tableMax = Const.CCC_ROOM_MAX_PLAYER

    -- 椅子
    self.userSeatArray = {}        -- 座椅对应玩家uid的哈希表 ， key ： seatID (1,2,3,4)   ，value： player
    self.userSeatArrayNumber = 0         -- 房间上有几个玩家， 记住，这里不能用#UserSeatArray, 因为有可能中间有椅子是空的，不连续的不能用#， 本质UserSeatArray是map ；  也不能遍历， 慢

end

function CCCRoom:Reload(c)
    setmetatable(c, self)
    self.__index = self

    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end

----------------------- 房间操作 ---------------------------------
function CCCRoom:InitRoom()
    if self:CheckTableEmpty() then
        -- 如果房间是空的， 那么需要初始化一下
        --self:InitDistributeInfo()
    end
end
--清理房间
function CCCRoom:ClearTable()
    self.userSeatArray = {}     --  seatID    player
    self.userSeatArrayNumber = 0
end
--判断房间是有人，还是空房间
function CCCRoom:CheckTableEmpty()
    if self.userSeatArrayNumber > 0 then
        return false
    end
    return true -- 空房间
end

--获取房间的空座位, 返回座椅的编号，从0开始到tableMax， 如果返回-1说明满了-
function CCCRoom:GetEmptySeatInTable()
    for i = 1, self.tableMax do
        if self.userSeatArray[i] == nil then
            return i
        end
    end
    return -1
end


-- 房间的主循环
function CCCRoom:RunRoom()
    if self:CheckTableEmpty() then
        --print("这是一个空房间" .. self.GameID)

        -- 这部分是做一个内存的测试
        ---- create Global Map hash
        --for i=1,40 do
        --    GlobalMap[tostring(i)] = i
        --end
        --
        --for i=1,40 do
        --    GlobalMap[tostring(i)] = nil
        --end

        --table.foreach(GlobalMap,function(i,v)
        --     print(i,v)
        --end )

        --clear  Global Map hash element
        --for i,_ in pairs(GlobalMap)do
        --    GlobalMap[tostring(i)] = nil        -- memery leak
        --end
        --
        --

        ---- if you clear all data, it's ok, but,  if you clear some data,  memery leak
        ----GlobalMap = {}
        --collectgarbage()

        --self.LastRunTime = ZTime.GetOsTimeMillisecond()
    else
        local now = ZTime.GetOsTimeMillisecond()

        --if self:GetFishNum() < MAX_Fish_NUMBER then
        --    self:RunDistributeInfo(table.RoomScore)
        --    self:RunBossDistributeInfo(table.RoomScore)
        --end
        --for _, bullet in pairs(self.BulletArray) do
        --    if now > bullet.DeadTime then
        --        self:DelBullet(bullet.BulletUID)     -- 生存时间已经到了，销毁
        --    end
        --    --bullet:BulletRun(now,self)      -- 遍历所有子弹，并且run
        --end
        --for _, fish in pairs(self.FishArray) do
        --    if now > fish.DeadTime then
        --        --print("鱼生存时间到了",self.FishUID)
        --        self:DelFish(fish.FishUID)
        --    end
        --    --fish:FishRun(now,self)              --遍历所有鱼，并且run
        --end

        -- 记录房间的运行状态
        --if now - self.LastRunTime > 1000 * 60  then     -- 60秒记录一次
        --    local state = {}
        --    state["FishNum"] = self:GetFishNum()        --当前有多少条鱼
        --    state["BulletNum"] = self:GetBulletNum()    --当前有多少子弹
        --    state["SeatArray"] = self.UserSeatArrayNumber    --当前有多少玩家
        --    SqlSaveGameState(self.GameID, self.roomId, state)
        --    self.LastRunTime = now
        --    --print("记录房间的运行状态")
        --end
        -- 检查玩家的情况，如果玩家长期离线，那么t掉，没人就清空房间
        --for k, player in pairs(self.UserSeatArray) do
        --    if player.NetWorkState == false then
        --        if now - player.NetWorkCloseTimer > ConstPlayerNetworkWaitTime then
        --            -- 玩家长时间断线，t掉吧
        --            self:PlayerStandUp(player.ChairID,player)
        --            Logger("长时间断线， t掉这个玩家"..player.User.UserID)
        --        end
        --    end
        --end
    end
end

----------------------- 玩家操作 ---------------------------------

--玩家坐到椅子上
function CCCRoom:PlayerSeat(seatId, player)
    self.userSeatArray[seatId] = player
    self.userSeatArrayNumber = self.userSeatArrayNumber + 1   -- 房间上玩家数量增加
end

--玩家离开椅子
function CCCRoom:PlayerStandUp(seatId, player)
    ZLog.Logger(Player.UId(player) .. "离开房间" .. player.roomId .. "椅子" .. player.chairId .. "self.roomId" .. self.gameId)
    -- 保存玩家基础数据
    --SaveUserBaseData(player.User)

    GameServer.SetAllPlayerList(Player.UId(player), nil)         -- 清理掉游戏管理的玩家总列表
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


-------------------------管理玩家-------------------------------------

local function seat(room, player, seatId)
    room:PlayerSeat(seatId, player)              --让玩家坐下.
    player.roomId = room.roomId
    player.chairId = seatId
    --self:SendYouLoginToOthers(player, room)-- 发消息给同房间的其他玩家，告诉他们你登录了
    return player
end
--- 有玩家登陆游戏
function Game.PlayerLoginGame(self,oldPlayer)
    local player = GameServer.GetPlayerByUID(Player.UId(oldPlayer)) -- 把之前的玩家数据取出来
    -- 如果玩家是断线重连的
    if player ~= nil then
        --找到之前有玩家在线
        if oldPlayer.gameId == player.gameId then
            -- 同一个游戏， 并且玩家状态是等待断线重连
            --player.NetWorkState = true                      -- 网络恢复正常
            --player.NetWorkCloseTimer = 0
            print("把断线重连的player返回去， 玩家本来就坐在这里，不用同步信息给其他玩家， 就是反应他傻了一会后继续游戏了")
            return player
        else
            -- 不是同一个游戏，或者有玩家在里面玩呢
            -- player会被替换掉，那么之前的连接也到t掉才可以

            -- 这里以后增加，t掉玩家的连接的功能
        end
    end

    -- 不是断线重连的就重新建一个玩家数据
    --player = Player:New(oldPlayer.User)
    --player.GameType = oldPlayer.GameType            -- 设定游戏类型
    player = oldPlayer
    GameServer.SetAllPlayerList(Player.UId(player), player)  --创建好之后加入玩家总列表

    --然后找一个有空位的房间让玩家加入游戏
    for k, room in pairs(self.allRoomList) do
        local seatId = BaseRoom.GetEmptySeatInTable(room)
        if seatId > 0 then
            print("有空座位")
            room:InitRoom()    -- 看看是不是空房间，如果是，需要初始化
            return seat(room,player,seatId)
        end
    end

    --没有空座位的房间了，创建一个
    print("没有空座位的房间了，创建一个吧,  score".. self.gameId)
    local gameId = self.allRoomList["1"].gameId
    local room = Game.CreateRoom(self, gameId)
    local seatId = BaseRoom.GetEmptySeatInTable(room)  --获取空椅位
    return seat(room,player,seatId)

end


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




----------------------- 消息 ---------------------------------
----玩家登陆的时候,发送场景其他消息
function CCCRoom:SendTableSceneInfo(player)
    if player == nil then
        ZLog.Logger("ByTable:SendTableSceneInfo player 对象nil")
        return
    end
    --1.发送场景Enter_scene信息
    self:SendEnterSceneInfo(Player.UId(player))
    --2.发送场景中鱼信息
    --self:SendSceneFishes(player.User.userId)
end


--- 发消息给同房间的其他玩家，告诉他们你登录了
function CCCRoom:SendYouLoginToOthers(player, table)
    --    print("玩家",player.User.UserID, "房间",table.roomId,"椅子",player.ChairID)

    --local CMD_Game_pb = require("CMD_Game_pb")
    --local sendCmd = CMD_Game_pb.CMD_S_OTHER_ENTER_SCENE()
    --sendCmd.user_info.user_id = player.User.UserID
    --sendCmd.user_info.chair_id = player.ChairID
    --sendCmd.user_info.table_id = player.roomId
    --table:SendMsgToOtherUsers(player.User.UserID, sendCmd, MDM_GF_GAME, SUB_S_OTHER_ENTER_SCENE)
end
--- 同步场景信息
function CCCRoom:SendEnterSceneInfo(UserId)
    local sendCmd = ProtoGameCCC.OtherEnterRoom()
    sendCmd.scene_id = self.gameId
    sendCmd.table_id = self.roomId
    for index, player in pairs(self.userSeatArray) do
        -- 从房间传递过来的其他玩家信息，原来坐着的玩家信息
        if player ~= nil then
            local uu = sendCmd.user:add()
            uu.user_id = Player.UId(player)
            uu.chair_id = index
        end
    end

    NetWork.SendToUser(UserId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_OTHER_LOGON, sendCmd, nil, nil) --进入房间
end
----玩家登陆的时候， 同步给玩家场景中目前鱼群的信息
--function CCCTable:SendSceneFishes(UserId)
--    --    print("鱼数量"..self.FishArrayNumber)
--    local sendCmd = CMD_Game_pb.CMD_S_SCENE_FISH()
--    local cmd
--    for _,fish in pairs(self.FishArray) do
--        cmd = sendCmd.scene_fishs:add()
--        cmd.uid = fish.FishUID
--        cmd.kind_id = fish.FishKindID
--    end
--    LuaNetWorkSendToUser(UserId, MDM_GF_GAME, SUB_S_SCENE_FISH, sendCmd, nil, nil)
--
--end

----------------------- 同步消息 ---------------------------------
--给桌上的所有玩家同步消息
function CCCRoom:SendMsgToAllUsers(mainCmd, subCmd, sendCmd)
    for _, player in pairs(self.userSeatArray) do
        if player ~= nil and player.netWorkState then
            local result = NetWork.SendToUser(Player.UId(player), mainCmd, subCmd, sendCmd, nil, 0)       -- 注意，这里因为是群发，所以token标记是0，就是不需要
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
function CCCRoom:SendMsgToOtherUsers(userId, sendCmd, mainCmd, subCmd)
    for _, player in pairs(self.userSeatArray) do
        if player ~= nil and userId ~= Player.UId(player) and player.netWorkState then
            NetWork.SendToUser(Player.UId(player), mainCmd, subCmd, sendCmd, nil, 0)       -- 注意，这里因为是群发，所以token标记是0，就是不需要
        end
    end
end
