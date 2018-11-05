---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/1 14:59
---
local FishServerExcel = require("mgby_fish_sever")
require("byBullet")
require("byFish")

ByTable = {}


function ByTable:New(tableId,gameTypeId)
    c = {
        GameID = gameTypeId,
        TableID = tableId,
        TableMax = BY_TABLE_MAX_PLAYER, --桌子容纳玩家数量
        RoomScore = 0,  --房间分数

        UserSeatArray = {},  -- 座椅对应玩家uid的哈希表 ， key ： seatID (1,2,3,4)   ，value： player


        GenerateFishUid = 0, -- 生成鱼的uid
        GenerateBulletUid = 0, -- 生成子弹的uid

        FishArray = {},   -- 鱼的哈希表
        BulletArray = {},   -- 子弹的哈希表


        DistributeArray = {},   -- 鱼的生成信息数据
        BossDistributeArray = {},   -- Boss鱼的生成信息数组
    }
    setmetatable(c,self)
    self.__index = self
    return c
end

------------主循环-------------------
function ByTable:RunTable()
    --    luaCallGoCreateGoroutine("RunTable")
    self:InitTable()        -- 可以进行初始化

    -- 开始桌子的主循环
    local RunTable = function()
        print("RunTable")
        if self:CheckTableEmpty() then
            print("这是一个空桌子")
        else

            --self:RunDistributeInfo(table.GetRoomScore())
            --self:RunBossDistributeInfo(table.GetRoomScore())


            --for _,bullet :=range self.BulletArray{
            --bullet.BulletRun(table)				// 遍历所有子弹，并且run
            --}
            --for _,fish :=range self.FishArray{
            --fish.FishRun(table)					// 遍历所有鱼，并且run
            --}

        end


    end
    FindGoRoutineAndRegisterTableRun(RunTable)    -- 注册开始一个新的协程
end

function ByTable:InitTable()

end
--------------------------------------------------------------------------------
--------------玩家----------------------------------------------------------
--------------------------------------------------------------------------------

-----判断桌子是有人，还是空桌子
function ByTable:CheckTableEmpty()
    if #self.UserSeatArray>0 then
        return false
    end

    return true -- 空桌子
end

----获取桌子的所有玩家-
function ByTable:GetUsersSeatInTable()
    local userList = {}
    for i=1,BY_TABLE_MAX_PLAYER do
        if self.UserSeatArray[i] ~= nil then
            -- 说明有人在座位上
            table.insert(userList,self.UserSeatArray[i])
        end
    end
    return userList
end

-----获取桌子的空座位, 返回座椅的编号，从0开始到tableMax， 如果返回-1说明满了-
function ByTable:GetEmptySeatInTable()
    for i=1,BY_TABLE_MAX_PLAYER do
        if self.UserSeatArray[i] ~= nil then
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
function ByTable:PlayerStandUp(seatID,user)
    self.UserSeatArray[seatID] = nil
    -- 清理掉玩家所有子弹
    self:DelBullets(user.UserId)
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
    self.UserSeatArray = {}

end

--------------------------------------------------------------------------------
----------------------子弹------------------------------------------------------
--------------------------------------------------------------------------------
---
-----玩家发射一个新的子弹
function ByTable:FireBullet(player , lockFishId)
    local num = player.ActivityBulletNum
    if num > MAX_BULLET_NUMBER then
        print("子弹超过上限了")
    end
    local cost = self.RoomScore
    if player.User.Score < cost then
        print("玩家没钱了")
    end
    -- 创建新的子弹
    local bullet = Bullet:New(self.GenerateBulletUid)
    self.BulletArray[bullet.BulletUID] = bullet     --把bullet加入列表
    bullet.UserID = player.User.UserID              --子弹的主人
    bullet.lockFishID = lockFishId                  --锁定鱼
    player.ActivityBulletNum = player.ActivityBulletNum + 1  --玩家已激活子弹增加
    self.GenerateBulletUid  = self.GenerateBulletUid + 1        -- 生成子弹id，自增
end

-----击中一条鱼
function ByTable:HitFish(userId ,bulletId, fishId)
    -- 增加CD判断，不可以太频繁

    --删除子弹
    self:DelBullet(bulletId)

    -- 获得鱼的金币
    local fish = self.FishArray[fishId]
    if fish ~= nil then
        print("捕鱼成功")
    end

    -- 删除鱼
    self:DelFish(fishId)
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

---清空所有的鱼群
function ByTable:DelFishes()
    self.FishArray = {}
    self.GenerateFishUid = 0  --重置一下生成鱼uuid
end


----玩家登陆的时候， 同步给玩家场景中目前鱼群的信息
function ByTable:SendSceneFishes(user)

end

--- 给所有玩家同步新建的鱼的信息
function ByTable:SendNewFishes(fish)

end

----------------------------------------------------------------------------
-----------------------------消息同步-----------------------------------------
----------------------------------------------------------------------------

-----给桌上的所有玩家同步消息
function ByTable:SendMsgToAllUsers(sendCmd,mainCmd,subCmd)

end

----给桌上的其他玩家同步消息
function ByTable:SendMsgToOtherUsers(userId,sendCmd,mainCmd,subCmd)

end


----------------------------------------------------------------------------
-----------------------------生成鱼池-----------------------------------------
----------------------------------------------------------------------------
----初始化鱼池的生成组----------------------------
function ByTable:InitDistributeInfo(roomScore)
    local startId = roomScore * 100
    local endId = startId + 100

    for k,v in pairs(FishServerExcel) do
        

    end

end