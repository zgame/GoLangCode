
--------------------------------------------------------------------------------------
---User的数据是玩家的游戏中永久数据，需要保存的数据
--------------------------------------------------------------------------------------

User = {}
function User:New()
    local c = {
        FaceId = 0, -- # 头像id
        Gender = 0, --  # 性别
        UserId = 0, --  # 用户id
        GameId = 0, --  # 游戏id
        Exp = 0, --  # 经验
        Loveliness = 0, --  # 魅力
        Score = 0, --  # 分数
        NickName = "", --  # 昵称
        Level = 0, --  # 等级
        VipLevel = 0, --  # vip等级
        AccountLevel = 0, --  # 账号等级
        SiteLevel = 0, --  # 炮等级
        CurLevelExp = 0, --  # 当前等级经验
        NextLevelExp = 0, --  # 下一等级经验
        PayTotal = 0, --  # 充值总金额
        Diamond = 0, --  # 钻石数量
        OpenId = "",  -- # 玩家的渠道账号，或者mac地址
    }
    setmetatable(c, self)
    self.__index = self
    return c
end


