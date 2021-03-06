
------------------------------------------------------------------------------
--- 策划用来调整游戏的参数
------------------------------------------------------------------------------


---------------奖券部分------------------
LotteryRechargeGiveRate = 0.9   --  奖券返奖率
LotteryRechargeRMBRate = 1000   -- 奖券对应RMB的比率
LotteryPoolMax  = 100000     -- 奖券水位的标准值， 用来做为分母，计算水位的比值
LotteryPoolInit = 23000         -- 奖券池的初始值，新玩家赋值



---------------翻牌部分------------------

CardBloodModeOpenRate = 300    --  热血模式开启几率 0.03,
CardBingoRate = 900  --    其他中奖率，0.09 % ，  不中奖率，0.88%


---------------老虎机部分------------------

TigerWheelOpenRate = 300    --  轮盘开启几率 0.03,
TigerBingoRate = 900  --    其他中奖率，0.09 % ，  不中奖率，0.88%



---------------轮盘第二轮猜红黑部分------------------
WheelDoubleRate = { 3500 ,9500 }  -- 中奖的几率  ， 第二次猜 35% ，  第一次猜45%


------------------积分赛奖励比例----------------------------------
PointsRewardList = {30 ,20, 10, 4 }     --1,2,3-5, 6-10


------------------幸运轮盘消耗奖券----------------------------------
LuckyWheelCostLottery = 1000