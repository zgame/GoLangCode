
CCCRoom = BaseRoom:extend()
function CCCRoom:New(roomId, gameTypeId)
    -- 重新赋值某些属性值
    CCCRoom.super.New(self)
    self.GameID = gameTypeId
    self.roomId = roomId
    self.TableMax = Const.CCC_ROOM_MAX_PLAYER

    -- 椅子
    self.UserSeatArray = {}        -- 座椅对应玩家uid的哈希表 ， key ： seatID (1,2,3,4)   ，value： player
    self.UserSeatArrayNumber = 0         -- 房间上有几个玩家， 记住，这里不能用#UserSeatArray, 因为有可能中间有椅子是空的，不连续的不能用#， 本质UserSeatArray是map ；  也不能遍历， 慢

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
    if BaseRoom.CheckTableEmpty(self) then
        -- 如果房间是空的， 那么需要初始化一下
        --self:InitDistributeInfo()
    end
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

        self.LastRunTime = ZTime.GetOsTimeMillisecond()
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
----玩家坐到椅子上
function CCCRoom:PlayerSeat(seatID, player)
    BaseRoom.PlayerSeat(self,seatID,player)
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

----------------------- 消息 ---------------------------------
----玩家登陆的时候,发送场景其他消息
function CCCRoom:SendTableSceneInfo(player)
    if player == nil then
        ZLog.Logger("ByTable:SendTableSceneInfo player 对象nil")
        return
    end
    --1.发送场景Enter_scene信息
    self:SendEnterSceneInfo(player.User.userId)
    --2.发送场景中鱼信息
    --self:SendSceneFishes(player.User.userId)
end


--- 同步场景信息
function CCCRoom:SendEnterSceneInfo(UserId)
    local sendCmd = protoGameCcc_pb.OtherEnterRoom()
    sendCmd.scene_id = self.GameID
    sendCmd.table_id = self.roomId
    for index, player in pairs(self.UserSeatArray) do
        -- 从房间传递过来的其他玩家信息，原来坐着的玩家信息
        if player ~= nil then
            local uu = sendCmd.table_users:add()
            uu.user_id = player.User.UserID
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
