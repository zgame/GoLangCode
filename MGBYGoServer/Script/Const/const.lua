-----------------------------------------------------------------
-- 常量的定义
-----------------------------------------------------------------
ConstPlayerNetworkWaitTime = 1000 * 1    -- 玩家断线等待的时间多长

-----------------------------------------------------------------
-- 游戏类型的常量定义
-----------------------------------------------------------------
GameTypeBY = 1   -- 普通捕鱼
GameTypeBY30 = 2 -- 普通捕鱼30倍房间
GameTypeBY2 = 3   -- 精灵捕鱼
GameTypeBY3 = 4   -- ** 捕鱼

-----------------------------------------------------------------
-- 小游戏
GameTypeXHS = 5     -- 小海兽

-----------------------------------------------------------------
-- 聊天服务
GameTypeChat = 100  -- 聊天服务
-----------------------------------------------------------------
-- 关于桌子的常量定义
-------------------------------------------------------------------

TABLE_CHAIR_NOBODY = -1         -- 桌子和椅子没人玩
BY_TABLE_MAX_PLAYER = 4         -- 捕鱼游戏桌子的最大人数
XHS_TABLE_MAX_PLAYER = 1000     -- 小海兽游戏桌子的最大人数
CHAT_TABLE_MAX_PLAYER = 10000   -- 聊天服务桌子的最大人数






-----------------------------------------------------------------
-- 关于子弹的常量定义
-------------------------------------------------------------------
MAX_BULLET_NUMBER = 15        --玩家最大子弹数量
MAX_Fish_NUMBER = 30        --最大鱼数量


-----------------------------------------------------------------
-- 关于鱼的常量定义
-------------------------------------------------------------------

FT_SMALL = 1    --小鱼
FT_MIDSIZE = 2    --中型鱼
FT_BIG = 3    --大型鱼
FT_BOSS = 4    --boos鱼


----------------------------------------------------
--聊天服redis缓存数据的目录
-------------------------------------------
ByChatServerRedisDir = {
    UserGameIDInfo = "byChatServerUserInfo:", --缓存玩家的基本信息，key=gameid，value=包含玩家基本的信息viplv、faceid等等信息
    UserNickNameInfo = "byChatUserNickInfo:", --缓存玩家的基本信息，key=nickname，value=gameid
    UserUIDInfo = "byChatServerUserUIDInfo:", --缓存玩家的基本信息，key=userid，value=gameid
    UserIdInfo = "byChatServerUserIdInfo:", --缓存玩家的基本信息,feild=userid，value=包含玩家基本的信息viplv、faceid等等信息
}


-----------------------------------------------------------------
-- 关于聊天服的聊天、好友相关的常量定义
-------------------------------------------------------------------

--这里是消息的状态
FriendMsgListState = {
    Normal = 0, --正常
    Update = 1, --更新
    Locked = 2, --锁定
}

MESSAGE_SPARATE = "|+#"              --"|+#"
MESSAGE_FIELD_SPARATE = "#-|"
MAX_FRIEND_APPLY_NUM = 20 ---同步前端的申请好友信息的条数最大值

-- 查询类型
ChatServerSearchType = {
    GameID = 0, -- 游戏ID
    Nick = 1    -- 昵称
}
-- token生存时间（秒）,1天
TOKEN_LIVE_SEC = 86400

-- 连接聊天服类型
Enum_ChatServerLinkKind = {
    None = 0,
    Server = 1, -- 服务器
    Player = 2, -- 玩家
}


--每秒1000毫秒
TIME_MILLISPS = 1000
TIME_1MIN = 60 * TIME_MILLISPS

-- 查询类型
Enum_DBSearchType = {
    UserID = 0, -- 用户ID
    GameID = 1, -- 游戏ID
    Nick = 2, -- 昵称
}

--- 最大连接数
ConnectMax = 10000
--- 外部端口号
ServicePort = 8410
--- 聊天服内部连接地址
InnerServiceAddr = '192.168.101.173'
--- 聊天服外部连接地址
OuterServiceAddr = '192.168.101.173'
--- 好友关系及聊天缓存过期时间（秒）
FriendAndCacheTimeout = 60
---  从数据库加载N天之内的未读消息
LoadMessageInDays = 7
--- 与好友聊天的间隔时间（秒）
ChatWithFriendInterval = 5
---搜索用户的间隔时间（秒）
SearchInverval = 1
--- 聊天内容字符数量
ChatMsgLength = 40
---  同IP最大连接数
MaxLinkPerIP = 30
---  1秒内1连接最大发包数量
MaxPacketPerSecond = 1000

---申请列表数量
ApplyFriendListNum = 10
--- 与好友聊天的间隔时间（秒）

--- 是否开启调试 在发布的时候设置为false
IS_Debug = true

--- 申请协议返回结果 probuf的枚举使用整数
Enum_ReplyResult = {
    Successful = 0, --- 成功
    Failed = 1, --- 失败
    ---  聊天相关
    NotFoundUser = 2, --- 没有搜到用户
    ApplyRepeatedly = 3, ---重复申请
    ApplyUpperLimit = 4, --- 申请次数达到上限
    FriendUpperLimit = 5, ---好友数量达到上限
    FriendExisted = 6, --- 已经是好友了
    ContainConfineChar = 7, --- 包含限制字符
    OtherFriendFull = 8, --- 对方好友已满
}

--- 申请好友操作类型
Enum_ApplyOpt = {
    Apply = 1, --- 申请
    Cancel = 2, --- 取消
    Agree = 3, --- 同意
    refuse = 4, --- 拒绝
}

-----------------------------------------------------------------
--- 道具类型
-----------------------------------------------------------------
Enum_ItemType = {
    eItemType_Skill = 1, -- 1：技能道具
    eItemType_Other = 2, -- 2：其他道具
    eItemType_Material = 3, -- 3：材料
    eItemType_Bullet_Skin = 4, -- 4：子弹皮肤
    eItemType_Game_Ticket = 5, -- 5：参赛券
    eItemType_Compose_Material = 6, -- 6：合成材料
    eItemType_Activity = 10, -- 10:活动道具
    eItemType_Summon = 11, -- 11.召唤石
    eItemType_JsGsaw_Chip = 12, -- 12.拼图碎片
    eItemType_Cannon = 13, -- 13.炮台
    eItemType_Recharge_Volume = 14, -- 14.充值满减券
    eItemType_Arean_Star = 15, -- 15.竞技之星
}

-----------------------------------------------------------------
--- 登录服向聊天服推送用户信息类型
-----------------------------------------------------------------
Enum_UserInfoType = {
    UIT_FACEID = 0; ---头像ID
    UIT_VIPLV = 1; ---VIP等级
}



