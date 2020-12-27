
CCCRoom = Class:extend()
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
function CCCRoom:PlayerSeat(chairId, player)
    self.userSeatArray[chairId] = player
    self.userSeatArrayNumber = self.userSeatArrayNumber + 1   -- 房间上玩家数量增加
    player.roomId = self.roomId
    player.chairId = chairId

    GameServer.SetAllPlayerList(Player.UId(player), player)  --创建好之后加入玩家总列表
    self:SendLoginToOthers(player)
    return player
end

--玩家离开椅子
function CCCRoom:PlayerStandUp(uId)
    local player = GameServer.GetPlayerByUID(uId)
    --ZLog.Logger(uId .. "离开房间" .. player.roomId .. "椅子" .. player.chairId .. "self.roomId" .. self.gameId)
    -- 保存玩家基础数据
    --SaveUserBaseData(player.User)

    GameServer.SetAllPlayerList(Player.UId(player), nil)         -- 清理掉游戏管理的玩家总列表
    self.userSeatArray[player.chairId] = nil                -- 清理掉房间的玩家列表
    self.userSeatArrayNumber = self.userSeatArrayNumber - 1  -- 房间上玩家数量减少
    self:SendLogoutToOthers(player)
    player.roomId = Const.ROOM_CHAIR_NOBODY
    player.chairId = Const.ROOM_CHAIR_NOBODY

    --如果是空房间的话，清理一下房间
    if self:CheckTableEmpty() then
        self:ClearTable()
        Game.ReleaseRoom(self.gameId,self.roomId)    --回收房间
    end
end


----------------------- 消息 ---------------------------------

-- 发消息给同房间的其他玩家，告诉他们你登录了
function CCCRoom:SendLoginToOthers(player)
    local userId =  Player.UId(player)
    print("玩家登录", userId, "房间",self.roomId,"椅子",player.chairId)
    local sendCmd = ProtoGameCCC.UserList()
    local uu = sendCmd.user:add()
    Player.Copy(player,uu)
    self:SendMsgToOtherUsers(CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_OTHER_LOGON,sendCmd,userId)
end

-- 发消息给同房间的其他玩家，告诉他们你登出了
function CCCRoom:SendLogoutToOthers(player)
    local userId =  Player.UId(player)
    print("玩家登出", userId, "房间",self.roomId,"椅子",player.chairId)
    local sendCmd = ProtoGameCCC.OtherLeaveRoom()
    sendCmd.userId = userId
    self:SendMsgToOtherUsers(CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_OTHER_LOGOUT,sendCmd,userId)
end

----------------------- 同步消息 ---------------------------------
--给桌上的所有玩家同步消息
function CCCRoom:SendMsgToAllUsers(mainCmd, subCmd, sendCmd)
    self:SendMsgToOtherUsers(mainCmd, subCmd, sendCmd,nil)
end

--给桌上的其他玩家同步消息
function CCCRoom:SendMsgToOtherUsers(mainCmd, subCmd,sendCmd,userId)
    for _, player in pairs(self.userSeatArray) do
        local uId = Player.UId(player)
        if player ~= nil and userId ~= uId then
            local result = NetWork.SendToUser(uId, mainCmd, subCmd, sendCmd, nil, 0)       -- 注意，这里因为是群发，所以token标记是0，就是不需要
            if not result then
                -- 发送失败了，玩家网络中断了
                self:PlayerStandUp(uId)
            end
        end
    end
end
