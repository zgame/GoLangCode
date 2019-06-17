
-----------------------------------------------------------------
-- 常量的定义
-----------------------------------------------------------------
ConstPlayerNetworkWaitTime = 1000 * 1    -- 玩家断线等待的时间多长

-----------------------------------------------------------------
-- 游戏类型的常量定义
-----------------------------------------------------------------

GameHall = 0        --游戏大厅

GameTypeBY  = 1   -- 普通捕鱼
GameTypeBY30  = 2   -- 普通捕鱼30倍房间
GameTypeBY2 = 3   -- 精灵捕鱼
GameTypeBY3 = 4   -- ** 捕鱼

GameTypeTB10 = 100  -- 普通推币
GameTypeTB100 = 101  -- 普通推币100倍房
GameTypeTB1000 = 102  -- 普通推币1000倍房
GameTypeTB10000 = 103  -- 普通推币10000倍房


-----------------------------------------------------------------
-- 关于桌子的常量定义
-------------------------------------------------------------------
TABLE_CHAIR_NOBODY = -1   -- 桌子和椅子没人玩
BY_TABLE_MAX_PLAYER = 4		-- 捕鱼游戏桌子的最大人数
TB_TABLE_MAX_PLAYER = 500		-- 推币游戏桌子的最大人数


-----------------------------------------------------------------
-- 关于推币机的常量定义
-------------------------------------------------------------------




-- 奖券
COIN_TYPE_NORMAL = 0
COIN_TYPE_1 = 1
COIN_TYPE_2 = 2
COIN_TYPE_3 = 3
COIN_TYPE_4 = 4
COIN_TYPE_MAX = 5       -- 奖券的种类

-- 牌的类型
CT_SPADE	= 1--		//黑桃MDM_MB_SERVER_LIST
CT_HEART	= 2--		//红桃
CT_CLUB		= 3--		//梅花
CT_DIAMOND	= 4--		//方块


-- 翻牌的类型
CARD_TYPE_FAIL = 0      --			//没中奖
CARD_TYPE_DUIZI = 1      --		//对子
CARD_TYPE_DOUBLE_DUIZI = 2      --		//两个对子
CARD_TYPE_SANTIAO = 3      --			//三条
CARD_TYPE_SHUNZI = 4      --			//顺子
CARD_TYPE_TONGHUA = 5      --			//同花
CARD_TYPE_HULU = 6      --				//葫芦 3带一对
CARD_TYPE_TIEZHI = 7      --			//铁支 4带1
CARD_TYPE_TONGHUASHUN = 8      --		//同花顺
CARD_TYPE_WANG = 9      --				//王牌
CARD_TYPE_MAX = 10      --

-- 老虎机
TIGER_TYPE_APPLE = 0        --		//苹果
TIGER_TYPE_ORANGE = 1      --				//橙子
TIGER_TYPE_LEMON_DIMOND = 2      --				//柠檬 砖石
TIGER_TYPE_BELL = 3     --				//铃铛
TIGER_TYPE_WATERMELON = 4      --			//西瓜
TIGER_TYPE_STAR_CHERRY = 5      --				//星星 樱桃
TIGER_TYPE_7 = 6      --					//7
TIGER_TYPE_BAR = 7     --					//BAR
TIGER_TYPE_ROU = 8      --					//轮盘 特殊处理	这项的概率不受老虎机每次调整公式控制
TIGER_TYPE_MAX = 9      --

-- 老虎机翻倍的状态
TIGER_DOUBLE_MS_OPEN	= 0;
TIGER_DOUBLE_MS_CLOSE	= 1;
--TIGER_DOUBLE_MS_TIMEUP	= 2;
TIGER_DOUBLE_MS_INIT		= 3;
--TIGER_DOUBLE_MS_READY	= 4;


-- 轮盘
ROULETTE_TYPE_STAR	= 0 -- //星星
ROULETTE_TYPE_1	= 1 -- //1 - 7
ROULETTE_TYPE_2	= 2 --
ROULETTE_TYPE_3	= 3 --
ROULETTE_TYPE_4	= 4 --
ROULETTE_TYPE_5	= 5 --
ROULETTE_TYPE_6	= 6 --
ROULETTE_TYPE_7	= 7 --
ROULETTE_TYPE_MAX	= 8 --
ROULETTE_TYPE_WRONG = 998   --		//错误


-- 轮盘双倍
ROULETTE_COLOR_BLACK = 0
ROULETTE_COLOR_RED = 1
ROULETTE_COLOR_CANCEL = 997	    --//不猜双倍
ROULETTE_COLOR_WRONG = 998		--//错误


-- 大金币类型
BIG_COIN_TYPE_NONE = 0      --  	//无
BIG_COIN_TYPE_APPLE = 1      --		//苹果
BIG_COIN_TYPE_ORANGE = 2      --			//橙子
BIG_COIN_TYPE_LEMON_DIMOND = 3      --			//柠檬 钻石
BIG_COIN_TYPE_BELL = 4      --			//铃铛
BIG_COIN_TYPE_WATERMELON = 5      --		//西瓜
BIG_COIN_TYPE_STAR_CHERRY = 6      --			//星星 樱桃
BIG_COIN_TYPE_777 = 7      --				//7
BIG_COIN_TYPE_BAR = 8      --				//BAR
BIG_COIN_TYPE_0 = 9      -- //星星
BIG_COIN_TYPE_1 = 10      -- //1 - 7
BIG_COIN_TYPE_2 = 11      --
BIG_COIN_TYPE_3 = 12      --
BIG_COIN_TYPE_4 = 13      --
BIG_COIN_TYPE_5 = 14      --
BIG_COIN_TYPE_6 = 15      --
BIG_COIN_TYPE_7 = 16      --
BIG_COIN_TYPE_BLOOD = 17      --
BIG_COIN_TYPE_MAX = 18      --

-- 热血模式状态
BMS_NULL	 = 0    --	//正常游戏
BMS_READY	 = 1	--  //准备进入热血模式
BMS_IN		 = 2	    --//热血模式中
BMS_HOLD	 = 3	    --//挂起状态 可以挤开
BMS_CONTINUE_READY = 4	        --//准备继续热血模式
BMS_CONTINUE_IN = 5     --	//继续

-- 热血模式客户端申请
BM_START	= 1	--//初始进入
BM_CONTINUE	= 2	--//继续 (挤开功能)
-- 热血模式的协议类型
BME_FAILED			= 0     --	//失败
BME_START_SUC		= 1	--//开始成功
BME_CAN_CONTINUE	= 2	--//可以继续
BME_CONTINUE_SUC	= 3	--//继续成功


-- 小奖池积分赛
QD_OPEN		= 0 	--//开启
QD_SYNC		= 1 	--//同步
QD_END		= 2	    --//结束
QD_USER		= 3 	--//个人
QD_TOP  	= 4 	--//榜首
QD_READY		= 5 	--//准备
QD_INIT		= 6 	--//间隔之后，可以开始了


-----------------------------------------------------------------
-- 关于邮件的常量定义
-------------------------------------------------------------------
MailTypeNormal = 0      -- 正常，还没有收取物品
MailTypeReceived = 1    -- 已经收取物品
MailTypeDeleted = 2     -- 已删除 状态，  每次再遍历系统邮件的时候，就不会再重复收取，







-----------------------------------------------------------------
-- 关于子弹的常量定义
-------------------------------------------------------------------
MAX_BULLET_NUMBER = 15 		--玩家最大子弹数量
MAX_Fish_NUMBER = 30 		--最大鱼数量






-----------------------------------------------------------------
-- 关于鱼的常量定义
-------------------------------------------------------------------

FT_SMALL = 1 	--小鱼
FT_MIDSIZE = 2 	--中型鱼
FT_BIG = 3 	--大型鱼
FT_BOSS = 4 	--boos鱼



