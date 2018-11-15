---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/1 14:59
---
local FishServerExcel = require("mgby_fish_sever")
require("byBullet")
require("byFish")
require("byFishDistribute")
require("player")
require("user")

local CMD_Game_pb = require("CMD_Game_pb")

ByTable = {}

function ByTable:New(tableId,gameTypeId)
    c = {
        GameID = gameTypeId,
        TableID = tableId,
        TableMax = BY_TABLE_MAX_PLAYER, --桌子容纳玩家数量
        RoomScore = 0,  --房间分数

        UserSeatArray = {},  -- 座椅对应玩家uid的哈希表 ， key ： seatID (1,2,3,4)   ，value： player


        GenerateFishUid = 1, -- 生成鱼的uid
        GenerateBulletUid = 1, -- 生成子弹的uid

        FishArray = {},   -- 鱼的哈希表    uid, fish
        BulletArray = {},   -- 子弹的哈希表   id,  bullet


        DistributeArray = {},   -- 鱼的生成信息数据    key顺序生成1,2,3,4...  Distribute
        BossDistributeArray = {},   -- Boss鱼的生成信息数组 key顺序生成1,2,3,4...  Distribute

        LastRunTime = 0   -- 循环周期时间
    }
    setmetatable(c,self)
    self.__index = self
    return c
end

------------主循环-------------------
function ByTable:RunTable()
    --    luaCallGoCreateGoroutine("RunTable")
--    self:InitTable()        -- 可以进行初始化

    -- 开始桌子的主循环
    local RunTable = function()
        if self:CheckTableEmpty() then
            --            print("这是一个空桌子")
            LastRunTime = GetOsTimeMillisecond()
        else
            local now = GetOsTimeMillisecond()

            if self:GetFishNum() < MAX_Fish_NUMBER then
                self:RunDistributeInfo(table.RoomScore)
                self:RunBossDistributeInfo(table.RoomScore)
            end
            for k, bullet in pairs(self.BulletArray) do
                bullet:BulletRun(self)      -- 遍历所有子弹，并且run
            end
            for k, fish in pairs(self.FishArray) do
                fish:FishRun(self)              --遍历所有鱼，并且run
            end
            for k, player in pairs(self.UserSeatArray) do
                if player.NetWorkState == false then
                    if now - player.NetWorkCloseTimer > 1000 then
                        -- 玩家长时间断线，t掉吧
                        self:PlayerStandUp(player.ChairID,player)
                        print("长时间断线， t掉这个玩家",player.User.UserId)
                    end
                end
            end

            if now - LastRunTime > 3000 then
                print(self.TableID.."当前有多少条鱼", self:GetFishNum())
                print(self.TableID.."当前有多少子弹", self:GetBulletNum())
                print(self.TableID.."当前有多少玩家", GetTableLen(self.UserSeatArray))
                LastRunTime = GetOsTimeMillisecond()
            end
        end
    end
    FindGoRoutineAndRegisterTableRun(RunTable)    -- 注册开始一个新的协程
end

function ByTable:InitTable()
    if self:CheckTableEmpty() then
        -- 如果桌子是空的， 那么需要初始化一下
        self:InitDistributeInfo()
    end
end


--- 判断玩家是否捕到鱼的逻辑判断
function ByTable:LogicCatchFish(player, LockFishIdList, BulletId)
--    print("玩家申请捕鱼")

    local bullet = self:GetBullet(BulletId)
    if bullet == nil then
        LuaNetWorkSendToUser(player.User.UserId,MDM_GF_GAME, SUB_S_CATCH_FISH,nil,"子弹的uid不正确")
        return
    end

    local ALLCurrScore = 0   -- 获得的分数
    local AllFishes = {}     -- 抓获的鱼list


    --printTable(LockFishIdList)

    for k,v in pairs(LockFishIdList) do
--        print("抓鱼uid",v)
        local fish = self:GetFish(v)
        if fish ~= nil then

            -- 这里判断鱼是否可以被捕获
            local isCatchFish = false

            -- 以后增加击中鱼之后的计算
            -- ...

            isCatchFish = true

            if isCatchFish then
                fish.CurrScore = tonumber( fish:GetFishScore())   --鱼的分数
                ALLCurrScore = ALLCurrScore + fish.CurrScore

                AllFishes[fish.FishUID] = fish
                -- 删除鱼
                self:DelFish(fish.FishUID)
            end

        end
    end

    if GetTableLen(AllFishes) == 0 then
        LuaNetWorkSendToUser(player.User.UserId,MDM_GF_GAME, SUB_S_CATCH_FISH,nil,"要捕获的鱼id不正确或者已经被别人捕捉了")
        return
    end

    --删除子弹
    self:DelBullet(BulletId)
    -- 获得鱼的金币
    player.User.Score = player.User.Score + ALLCurrScore

    --print("获得金币",ALLCurrScore)
    --print("当前金币",player.User.Score)
    --printTable(AllFishes)


    -- 给所有玩家同步一下，这个玩家捕到鱼了
    local sendCmd = CMD_Game_pb.CMD_S_CATCH_FISH()

    for k,v in pairs(AllFishes) do
        local cmd = sendCmd.catch_fishs:add()
        cmd.fish_uid = v.FishUID
        cmd.fish_score = v.CurrScore
    end

    sendCmd.chair_id = player.ChairID
    sendCmd.bullet.bullet_id = bullet.BulletUID
    sendCmd.curr_score = player.User.Score

    self:SendMsgToAllUsers(MDM_GF_GAME, SUB_S_CATCH_FISH,sendCmd)


end






--------------------------------------------------------------------------------
--------------玩家----------------------------------------------------------
--------------------------------------------------------------------------------

-----判断桌子是有人，还是空桌子
function ByTable:CheckTableEmpty()
    if GetTableLen(self.UserSeatArray) >0 then
        return false
    end

    return true -- 空桌子
end

--获取桌子的所有玩家-
--function ByTable:GetUsersSeatInTable()
--    local userList = {}
--    for i=1,BY_TABLE_MAX_PLAYER do
--        if self.UserSeatArray[i] ~= nil then
--            -- 说明有人在座位上
--            table.insert(userList,self.UserSeatArray[i])
--        end
--    end
--    return userList     -- 元素是player对象
--end

-----获取桌子的空座位, 返回座椅的编号，从0开始到tableMax， 如果返回-1说明满了-
function ByTable:GetEmptySeatInTable()
    for i=1,BY_TABLE_MAX_PLAYER do
        if self.UserSeatArray[i] == nil then
            return i        -- 返回当前空着的座位号(1,2,3,4)
        end
    end
    return -1   -- 座位满了
end

----玩家坐到椅子上
function ByTable:PlayerSeat(seatID,player)
    self.UserSeatArray[seatID] = player
end
----玩家离开椅子
function ByTable:PlayerStandUp(seatID,player)
    local game = GetGameByID(player.GameType)
    game.AllUserList[player.User.UserId] = nil      -- 清理掉游戏管理的玩家总列表
    self.UserSeatArray[seatID] = nil                -- 清理掉桌子的玩家列表
    player.TableID = TABLE_CHAIR_NOBODY
    player.ChairID = TABLE_CHAIR_NOBODY

    -- 清理掉玩家所有子弹
    self:DelBullets(player.User.UserId)
    --如果是空桌子的话，清理一下桌子
    if self:CheckTableEmpty() then
        self:ClearTable()
    end
end

-----清理桌子
function ByTable:ClearTable()
    -- 清理一下生成鱼的结构
    self.DistributeArray = {}
    self.BossDistributeArray = {}

    -- 清理掉所有子弹和鱼群
    self:DelBullets(-1)
    self:DelFishes()

    self.BulletArray = {}
    self.FishArray = {}
    self.UserSeatArray = {}     --  seatID    player

end

--------------------------------------------------------------------------------
----------------------子弹------------------------------------------------------
--------------------------------------------------------------------------------
---
-----玩家发射一个新的子弹
function ByTable:HandleUserFire(player , lockFishId)
    --print("玩家发射一个子弹")

    local num = player.ActivityBulletNum
--    if num > MAX_BULLET_NUMBER then
----        print("子弹超过上限了")
--        LuaNetWorkSendToUser(player.User.UserId,MDM_GF_GAME, SUB_S_USER_FIRE,nil,"子弹超过上限了")
--        return
--    end
    local cost = self.RoomScore
    if player.User.Score < cost then
        print("玩家没钱了")
        LuaNetWorkSendToUser(player.User.UserId,MDM_GF_GAME, SUB_S_USER_FIRE,nil,"玩家没钱了")
        return
    end
    -- 创建新的子弹
    local bullet = Bullet:New(self.GenerateBulletUid)
    self.BulletArray[bullet.BulletUID] = bullet     --把bullet加入列表
    bullet.UserID = player.User.UserID              --子弹的主人
    bullet.lockFishID = lockFishId                  --锁定鱼
    player.ActivityBulletNum = player.ActivityBulletNum + 1  --玩家已激活子弹增加
    self.GenerateBulletUid  = self.GenerateBulletUid + 1        -- 生成子弹id，自增


    -- 给所有玩家同步一下，这个玩家发子弹了
    local sendCmd = CMD_Game_pb.CMD_S_USER_FIRE()
    sendCmd.chair_id = player.ChairID
    sendCmd.bullet_id = bullet.BulletUID
    sendCmd.lock_fish_id = lockFishId
    --sendCmd.curr_score = player.User.Score

    self:SendMsgToAllUsers(MDM_GF_GAME, SUB_S_USER_FIRE,sendCmd)


end

function ByTable:GetBullet(bulletId)
    return self.BulletArray[bulletId]
end

----删除特定id的子弹
function ByTable:DelBullet(bulletId)
    local bullet = self.BulletArray[bulletId]
    if bullet ~= nil then
        self.BulletArray[bulletId] = nil
    end
    if #self.BulletArray == 0 then
        self.GenerateBulletUid = 0  --重置一下生成子弹uuid
    end
end

---- 删除所有子弹， 1 如果传入玩家uid，删除玩家的  ； 2  如果传入 -1 ，那么删除所有的
function ByTable:DelBullets(userId)
    if userId == -1 then
        self.BulletArray = {}
        return
    end
    for k,v in pairs(self.BulletArray) do
        if v.UserID == userId then
            self.BulletArray[k] = nil
        end
    end
    if #self.BulletArray == 0 then
        self.GenerateBulletUid = 0  --重置一下生成子弹uuid
    end
end
---- 有多少子弹
function ByTable:GetBulletNum()
    local re = GetTableLen(self.BulletArray)
    return re
end
----------------------------------------------------------------------------
------------------------------鱼-----------------------------------------
---------------------------------------------------------------------------

----新建一个新的鱼
function ByTable:CreateFish()
    local fish = Fish:New(self.GenerateFishUid)
    self.FishArray[fish.FishUID] = fish
    self.GenerateFishUid = self.GenerateFishUid  + 1
    return fish
end

--- 获取鱼的句柄
function ByTable:GetFish(fishId)
    return self.FishArray[fishId]
end

----删除特定uid的鱼
function ByTable:DelFish(fishId)

    local fish = self.FishArray[fishId]
    if fish ~=nil then
        self.FishArray[fishId] = nil
    end
    if #self.FishArray == 0 then
        self.GenerateFishUid = 0  --重置一下生成鱼uuid
    end

end
----删除特定uid的鱼list
function ByTable:DelFishList(fishIdList)
    for k,v in pairs(fishIdList) do
        local fish = self.FishArray[v]
        if fish ~=nil then
            self.FishArray[fishId] = nil
        end
    end
    if #self.FishArray == 0 then
        self.GenerateFishUid = 0  --重置一下生成鱼uuid
    end
end



---清空所有的鱼群
function ByTable:DelFishes()
    self.FishArray = {}
    self.GenerateFishUid = 0  --重置一下生成鱼uuid
end


---- 有多少条鱼
function ByTable:GetFishNum()
    local re = GetTableLen(self.FishArray)
    return re
end


----玩家登陆的时候， 同步给玩家场景中目前鱼群的信息
function ByTable:SendSceneFishes(UserId)
--    print("鱼数量"..GetTableLen(self.FishArray))
    local sendCmd = CMD_Game_pb.CMD_S_SCENE_FISH()
    for k,fish in pairs(self.FishArray) do
        local cmd = sendCmd.scene_fishs:add()
        cmd.uid = fish.FishUID
        cmd.kind_id = fish.FishKindID
    end
    LuaNetWorkSendToUser(UserId, MDM_GF_GAME, SUB_S_SCENE_FISH, sendCmd, nil)

end

--- 给所有玩家同步新建的鱼的信息
function ByTable:SendNewFishes(fish)
    local sendCmd = CMD_Game_pb.CMD_S_DISTRIBUTE_FISH()
    local cmd = sendCmd.fishs:add()
    cmd.uid = fish.FishUID
    cmd.kind_id = fish.FishKindID

    self:SendMsgToAllUsers(MDM_GF_GAME, SUB_S_DISTRIBUTE_FISH, sendCmd)
end

----------------------------------------------------------------------------
-----------------------------消息同步-----------------------------------------
----------------------------------------------------------------------------

-----给桌上的所有玩家同步消息
function ByTable:SendMsgToAllUsers(mainCmd,subCmd,sendCmd)
    for k,player in pairs(self.UserSeatArray) do
        if player ~= nil and player.IsRobot == false and player.NetWorkState then
            local result = LuaNetWorkSendToUser(player.User.UserId,mainCmd,subCmd,sendCmd,nil)
            if not result then
                -- 发送失败了，玩家网络中断了
                player.NetWorkState = false
                player.NetWorkCloseTimer = GetOsTimeMillisecond()
            end
        end
    end
end

----给桌上的其他玩家同步消息
function ByTable:SendMsgToOtherUsers(userId,sendCmd,mainCmd,subCmd)
    for k,player in pairs(self.UserSeatArray) do
        if player ~= nil and player.IsRobot == false and userId ~= player.User.UserId and player.NetWorkState then
            LuaNetWorkSendToUser(player.User.UserId,mainCmd,subCmd,sendCmd,nil)
        end
    end
end


----------------------------------------------------------------------------
-----------------------------生成鱼池-----------------------------------------
----------------------------------------------------------------------------
----初始化鱼池的生成组----------------------------
function ByTable:InitDistributeInfo()
    local startId = self.RoomScore * 100
    local endId = startId + 100

    for fishKind, v in pairs(FishServerExcel) do
        if v.is_use == 1 and fishKind > startId and fishKind < endId then
            local distribute = FishDistribute:New()
            distribute.FishKindID = fishKind

            distribute.CreateTime = GetOsTimeMillisecond()               --生成时间
            distribute.DistributeIntervalTime = distribute:GetIntervalTime(fishKind)      --获取时间间隔

            if distribute.DistributeIntervalTime == 0 then
                distribute.DistributeIntervalTime = 1000            -- 这里有生成间隔是0的东西
            end

            local fishType = v.type
            if fishType == FT_BOSS then
                table.insert(self.BossDistributeArray, distribute)  --加入到Boss鱼生成列表中
            else
                table.insert(self.DistributeArray, distribute)      --加入到普通鱼生成列表中
            end
        end
    end
--    print("桌子初始化鱼池结束")
end

----循环鱼池的生成组
function ByTable:RunDistributeInfo(roomScore)
    local now = GetOsTimeMillisecond()
    --printTable(self.DistributeArray)
    for k,Distribute in pairs(self.DistributeArray) do
        local kindId = Distribute.FishKindID
        -- 到下一个生成时间了, 那么我们来生成鱼吧

        --print("now",now,"Distribute.CreateTime",Distribute.CreateTime, " Distribute.DistributeIntervalTime", Distribute.DistributeIntervalTime)

        if now >  Distribute.CreateTime + Distribute.DistributeIntervalTime then
            --print("生成时间到了")
            --print("now",now,"Distribute.CreateTime",Distribute.CreateTime, " Distribute.DistributeIntervalTime", Distribute.DistributeIntervalTime)
            local createType = 0   --鱼怎么走
            local buildNum = 0      -- 鱼生成数量
            local max = FishServerExcel[kindId].count_max
            if max > 1 then
                -- 多生成几条鱼
                buildNum =  Distribute:GetCount(kindId)
--                print("随机生成"..buildNum.."条鱼")
                if buildNum < 1 then
                    buildNum = 1
                else
                    createType = 1  --生成一条路径的
                    if buildNum >=5 or ZRandomPercentRate(50) then
                        createType = 2  --位置要做偏移
                    end
                end
                Distribute.NextCreateTime = now
                Distribute.NextInterBuildTime = Distribute:GetCountFishTime(kindId)
            end
            Distribute.BuildNumber = buildNum
            Distribute.CreateTime = now
            Distribute.DistributeIntervalTime = Distribute:GetIntervalTime(kindId)
            Distribute.CreateType = createType                  --创建类型
            Distribute.FirstPathID = Distribute:GetPathType()    -- 获取路径

            -- 创建鱼
            self:DistributeNewFish(Distribute,0,0)
        end

        -- 多条鱼的判断
        if Distribute.BuildNumber > 1 then
            if now > Distribute.NextCreateTime + Distribute.NextInterBuildTime then
--                print("生成多条鱼"..kindId)
                local offsetX = 0
                local offsetY = 0
                Distribute.NextCreateTime = now
                Distribute.BuildNumber = Distribute.BuildNumber - 1
                Distribute.NextInterBuildTime = Distribute:GetCountFishTime(kindId)
                if Distribute.CreateType == 2 then
                    -- 位置偏移
                    offsetX = Distribute:GetOffsetXY()[1]
                    offsetY = Distribute:GetOffsetXY()[2]
                end
                -- 创建鱼
                self:DistributeNewFish(Distribute,offsetX,offsetY)
            end
        end

    end
end

-----循环Boss鱼池的生成组
function ByTable:RunBossDistributeInfo(roomScore)
    local now = GetOsTimeMillisecond()
    for k,Distribute in pairs(self.BossDistributeArray) do
        --到下一个生成时间了, 那么我们来生成鱼吧
        if now > Distribute.CreateTime + Distribute.DistributeIntervalTime then
            local kindId = Distribute.FishKindID
            local buildNum = 1   -- 鱼生成数量
            Distribute.BuildNumber = buildNum
            Distribute.CreateTime = now
            Distribute.DistributeIntervalTime = Distribute:GetIntervalTime(kindId)
            Distribute.FirstPathID = Distribute:GetPathType()    -- 获取路径

            self:DistributeNewFish(Distribute,0,0)
        end
    end
end


---- 具体生成鱼
function ByTable:DistributeNewFish(Distribute,offsetX,offsetY)
--    print("生成一条鱼"..Distribute.FishKindID)
    local kindId = Distribute.FishKindID
    -- 创建鱼
    local fish = self:CreateFish()
    fish:CreateFish(kindId,offsetY,offsetY,Distribute.FirstPathID)

    -- 发送给所有客户端生成鱼的消息
    self:SendNewFishes(fish)

end