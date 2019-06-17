

TbTable = {}
function TbTable:New(tableId,gameTypeId)
    local c = {
        GameID = gameTypeId,
        TableID = tableId,
        TableMax = TB_TABLE_MAX_PLAYER, --桌子容纳玩家数量
        RoomScore = 0,  --房间分数

        -- 椅子
        UserSeatArray = {},  -- 座椅对应玩家uid的哈希表 ， key ： seatID (1,2,3,4)   ，value： player
        UserSeatArrayNumber = 0,  -- 桌子上有几个玩家， 记住，这里不能用#UserSeatArray, 因为有可能中间有椅子是空的，不连续的不能用#， 本质UserSeatArray是map ；  也不能遍历， 慢


        -- 房间库存
        PoolAllScore = 0 ,  -- 所有玩家的库存，以后弄到数据库做跨服

        -- 大奖池
        JackpotAll = 0 ,    -- 大奖池
        JackpotLastPlayerName = '',  -- 上次中奖玩家昵称
        JackpotLastGetScore = 0,  -- 上次中奖玩家获得分数

        -- 老虎机翻倍
        TigerDoublePool = 0 ,   -- 老虎机翻倍池子
        TigerDoubleType = 0 ,   -- 老虎机翻倍的类型
        TigerDoubleRate = 0 ,   -- 老虎机翻倍的倍率，随机出来的
        TigerDoubleTimeStart = 0 ,  -- 老虎机翻倍开始的时间
        TigerDoubleTimeDuring = 0 ,  -- 老虎机翻倍持续的时间,秒
        TigerDoubleState = TIGER_DOUBLE_MS_CLOSE ,      -- 翻倍状态


        -- 返奖率的计算
        StatisticAllScoreGenerate =  0 ,  -- 总产出分数 ， 总玩家获得分数
        StatisticAllScoreCost =  0 ,    -- 总消耗分数 ，总玩家投入分数


        -- 小奖池库存
        PointsScore = 0 ,       -- 小奖池库存
        PointsState = QD_INIT,   -- 小奖池状态
        PointsTopList = {}  ,     --  玩家的前几名   结构是一个数组，里面每个元素是键值对 uid: **  ,point: **
        PointsAllPlayer = {},       -- 所有参与玩家的数据保存,  key uid ， 注意是string    value  point
        PointsPlayerNumber = 0 ,        -- 所有参与玩家的数量
        PointsStateStartTime = 0,            -- 状态的开始时间，用来计算
        PointsStateRemainTime = 0 ,         -- 状态的持续时间


        --- 库存
        --m_llRewardMax = 0, --奖励最大值
        --m_llPrizeValue = 0, --金币派奖峰值
        --m_dwPrizeCount = 0, --派奖次数
        --m_llPrizeAllScore = 0, --派奖总分数
        --m_dwPrizeAllPlayer = 0, --派奖总人数
        --m_llStockScore = 0, --库存数目


        --m_dRecycleBigPool = 0, --回收值
        --m_llLineUpOffSet = 0, --上线偏移
        --
        --m_llScoreLineDown = 0, --下线
        --m_llScoreLineUp = 0, --上线
        --m_lPerRevenue = 0, --税率/万比
        --m_llRevenue = 0, --已收税
        --m_llPutMoneyMax = 0, --最大吐钱
        --m_llSysScore = 0, --系统变化分数
        --m_lVarScoreOnce = 0, --单局变化

        --- 房间




        --m_wTax = 0, --税率
        --LONGLONG	m_llBigJackpot;	--大奖池初始值
        --LONGLONG	m_llSmallJackPot; --小奖池初始值
        --m_wInventorySwitch = 0, --库存开关
        --m_wLotterySwitch = 0, --奖券开关
        --FLOAT		m_fSlamAddValue;	--大奖池增加值
        --m_fSlotAffValue = 0, --老虎机奖池增加值
        --m_IMultiSlotReadyTimeRandUpLimit = 0, --翻倍老虎机准备时间随机上限
        --m_llInventoryInitial = 0, --初始库存
        --m_llInventoryUp = 0, --库存上限
        --m_llInventoryDown = 0, --库存下限


        --m_fRoomAdjust = 0, --房间调节初值
        --m_fRoomParam = 0, --房间参数
        --m_iAutoCoinSpeed = 0, --自动投币速度


        --m_fSlotAdjust = 0, --老虎机调节值
        --m_wSlotWin = 0, --老虎机中间初始几率
        --m_wSlotStandard = 0, --老虎机标准值
        --m_fSlotRange = 0, --老虎机幅度调节
        --m_fSlotAdjustUp = 0, --老虎机库存上限后调节值起始值
        --m_fSlotNobleRange = 0, --老虎机贵族值扣减系数
        --m_fSlotWinRange = 0, --老虎机贵族值影响最终中奖率的调节系数
        --m_iSlotLowRou = 0, --中轮盘的下限几率
        --m_iSlotTopRou = 0, --中轮盘的上限几率
        --m_iNobleSlotTopRou = 0, --中轮盘的上限几率---贵族值
        --m_iMultiSlotRemainTime = 0, --翻倍时间
        --
        --m_fPokerAdjust = 0, --翻牌中奖调节
        --m_wPokerWin = 0, --翻牌中奖初始几率
        --m_wPokerStandard = 0, --翻牌标准值
        --m_fPokerRange = 0, --翻牌幅度调节
        --m_fPokerAdjustUp = 0, --翻牌库存上限后调节值起始值
        --m_fPokerNobleRange = 0, --翻牌贵族值扣减系数
        --m_fPokerWinRange = 0, --翻牌贵族值影响最终中奖率的调节系数
        --m_fPokerBackTime = 0, --扑克牌收回时间
        --m_fBloodRewardScale = 0, --热血模式奖励翻倍
        --
        --m_wPokerRedTime = 0, --热血模式事件 毫秒
        --m_wPokerRedCondition = 0, --热血进入下回合初始条件
        --m_wPokerRedRound = 0, --热血回合调节值
        --m_fPokerRedAdjust = 0, --热血时间调节值
        --m_wPokerDoor = 0, --热血开门初始几率
        --m_fPokerDoorAdjust = 0, --热血开门调节
        --
        --m_wRouletteEnergy = 0, --轮盘能量满值
        --m_wRouletteStandard = 0, --轮盘标准值
        --m_fRouletteRange = 0, --轮盘幅度调节
        --m_wRouletteRedRate = 0, --轮盘猜红黑初始几率
        --m_fRouletteRedAdjust = 0, --轮盘猜红黑调节值
        --m_iCastCoinEnergy = 0, --投币能量
        ----m_iObtainCoinEnergy = 0, --推币能量
        --m_fRouletteNobleRange = 0, --轮盘贵族值扣减系数
        --m_fRouletteWinRange = 0, --轮盘贵族值影响最终中奖率的调节系数
        --m_iRouDouProNoble_Up = 0, --贵族值影响猜双倍几率上限
        --m_iRouDouProHonour_Up = 0, --荣誉值影响猜双倍几率上限

        --m_dwTigerCD = 0, --老虎机触发cd，ms单位
        --m_dwPokerCD = 0, --卡牌触发cd，ms单位
        --m_dwTigerResetCoinNum = 0, --老虎机触发限制重置投币数目
        --m_dwTigerTrigerLimitNum = 0, --老虎机触发超过此限制则当次开奖几率为0
        --m_dwPokerResetCoinNum = 0, --卡牌触发限制重置投币数目
        --m_dwPokerTrigerLimitNum = 0, --卡牌触发超过此限制则当次开奖几率为0


        --m_iDrawCoefficent = 0,		--抽奖值增加因数
        --m_fBigPoolAddNumPerCoin= 0,		--大奖池增加值


        -- 服务器开关
        cbCloseLottery = 0,  -- 奖券的开关 , 0 是开启奖券， 1是关闭奖券

        -- 用来做服务器记录的
        LastRunTime = 0 ,   -- 循环周期时间

        -- 统计相关
        PlayerWinTopList = {}  ,     --  玩家赢钱的前几名   结构是一个数组，里面每个元素是键值对 uid: **  ,score: **
        PlayerLostTopList = {}       --  玩家输钱的前几名   结构是一个数组，里面每个元素是键值对 uid: **  ,score: **

    }
    setmetatable(c, self)
    self.__index = self
    return c
end

function TbTable:Reload(c)
    setmetatable(c, self)
    self.__index = self

    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end

------------主循环-------------------
function TbTable:StartTable()
    self:InitTable()        -- 可以进行初始化
end


-- 桌子的主循环
function TbTable:RunTable()
    if self:CheckTableEmpty() then
        --print("这是一个空桌子")
        --print("FishServerExcel[101].type"..FishServerExcel["101"].type)
        self.LastRunTime = GetOsTimeMillisecond()
    else
        local now = GetOsTimeMillisecond()

        -- 记录桌子的运行状态
        if now - self.LastRunTime > 1000 * 6  then     -- 60秒记录一次
            local state ={}
            state["PoolAll"] = self.PoolAllScore
            state["Jackpot"] = self.JackpotAll
            state["SeatArray"] = self.UserSeatArrayNumber    --当前有多少玩家
            state["RewardRate"] =  self:GetRewardRate()
            SqlSaveGameState(self.GameID, self.TableID, state)
            self.LastRunTime = now
            --print("记录桌子的运行状态")

        end
        -- 检查玩家的情况，如果玩家长期离线，那么t掉，没人就清空桌子
        --for k, player in pairs(self.UserSeatArray) do
        --    if player.NetWorkState == false then
        --        if now - player.NetWorkCloseTimer > ConstPlayerNetworkWaitTime then
        --            -- 玩家长时间断线，t掉吧
        --            self:PlayerStandUp(player.ChairID,player)
        --            Logger("长时间断线， t掉这个玩家"..player.User.UserId)
        --        end
        --    end
        --end
        end
end




function TbTable:InitTable()
    if self:CheckTableEmpty() then
        -- 如果桌子是空的， 那么需要初始化
        for i=1,5 do
            self.PlayerWinTopList[i] = { uid =0 , score = 0}
            self.PlayerLostTopList[i] = { uid =0 , score = 0}
        end

        TigerDoublePoolStart(self)  -- 随机老虎机翻倍的参数
    end
end


--------------------------------------------------------------------------------
----------------------桌子逻辑关于 玩家 的处理----------------------------------
--------------------------------------------------------------------------------



-----判断桌子是有人，还是空桌子
function TbTable:CheckTableEmpty()
    if self.UserSeatArrayNumber >0 then
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
function TbTable:GetEmptySeatInTable()
    for i=1,self.TableMax do
        if self.UserSeatArray[i] == nil then
            return i        -- 返回当前空着的座位号(1,2,3,4)
        end
    end
    return -1   -- 座位满了
end

----玩家坐到椅子上
function TbTable:PlayerSeat(seatID,player)
    self.UserSeatArray[seatID] = player
    self.UserSeatArrayNumber = self.UserSeatArrayNumber + 1   -- 桌子上玩家数量增加


end
----玩家离开椅子
function TbTable:PlayerStandUp(seatID,player)
    Logger(player.User.UserId.."离开桌子"..player.TableID.."椅子"..player.ChairID)
    local game = GetGameByID(player.GameType)

    self.UserSeatArray[seatID] = nil                -- 清理掉桌子的玩家列表
    self.UserSeatArrayNumber = self.UserSeatArrayNumber - 1  -- 桌子上玩家数量减少
    player.TableID = TABLE_CHAIR_NOBODY
    player.ChairID = TABLE_CHAIR_NOBODY
    player.GameType = GameHall      -- 返回大厅

    -- 清理掉玩家所有子弹
    --self:DelBullets(player.User.UserId)
    --如果是空桌子的话，清理一下桌子
    if self:CheckTableEmpty() then
        self:ClearTable()
        game:ReleaseTableByUID(self.TableID)    --回收桌子
    end
end

-----清理桌子
function TbTable:ClearTable()
    -- 清理一下生成鱼的结构
    --self.DistributeArray = {}
    --self.BossDistributeArray = {}
    --
    ---- 清理掉所有子弹和鱼群
    --self:DelBullets(-1)
    --self:DelFishes()
    --
    --self.BulletArray = {}
    --self.FishArray = {}
    self.UserSeatArray = {}     --  seatID    player

end


---有新玩家加入桌子
function TbTable:NewPlayerJoin(player)
    local sendCmd = CMD_Game_pb.CMD_S_ENTER_SCENE()
    sendCmd.scene_id = player.GameType
    sendCmd.table_id = player.TableID
    for index, playerSeat in pairs(self.UserSeatArray) do       -- 从桌子传递过来的其他玩家信息，原来坐着的玩家信息
        if playerSeat ~= nil then
            local uu = sendCmd.table_users:add()
            uu.user_id = playerSeat.User.UserId
            uu.chair_id = index
        end
    end
    sendCmd.room_score = self.RoomScore
    LuaNetWorkSendToUser(player.User.UserId,MDM_GF_GAME, SUB_S_ENTER_SCENE, sendCmd, nil, nil) --进入房间
end

----------------------------------------------------------------------------
-----------------------------消息同步----------------------------------------
----------------------------------------------------------------------------

-----给桌上的所有玩家同步消息
function TbTable:SendMsgToAllUsers(mainCmd,subCmd,sendCmd)
    for _,player in pairs(self.UserSeatArray) do
        if player ~= nil and player.IsRobot == false and player.NetWorkState then
             --LuaNetWorkSendToUser(player.User.UserId,mainCmd,subCmd,sendCmd,nil)
            local result = LuaNetWorkSendToUser(player.User.UserId,mainCmd,subCmd,sendCmd,nil, 0)       -- 注意，这里因为是群发，所以token标记是0，就是不需要
            if not result then
                -- 发送失败了，玩家网络中断了
                --player.NetWorkState = false
                --player.NetWorkCloseTimer = GetOsTimeMillisecond()
                self:PlayerStandUp(player.ChairID,player)
            end
        end
    end
end

----给桌上的其他玩家同步消息
function TbTable:SendMsgToOtherUsers(userId,sendCmd,mainCmd,subCmd)
    for _,player in pairs(self.UserSeatArray) do
        if player ~= nil and player.IsRobot == false and userId ~= player.User.UserId and player.NetWorkState then
            LuaNetWorkSendToUser(player.User.UserId,mainCmd,subCmd,sendCmd,nil, 0)       -- 注意，这里因为是群发，所以token标记是0，就是不需要
        end
    end
end



----------------------------------------------------------------------------
-----------------------------返奖率计算----------------------------------------
----------------------------------------------------------------------------

function TbTable:GetRewardRate()
    if self.StatisticAllScoreCost > 0 then
        --print("self.StatisticAllScoreGenerate"..self.StatisticAllScoreGenerate)
        --print("self.StatisticAllScoreCost"..self.StatisticAllScoreCost)
        local str = string.format("%.2f",self.StatisticAllScoreGenerate / self.StatisticAllScoreCost *100)
        --print("返奖率   "..str.."%")
        return str.."%"
    else
        return "0%"
    end
end