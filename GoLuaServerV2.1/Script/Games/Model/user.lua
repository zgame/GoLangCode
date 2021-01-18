
--------------------------------------------------------------------------------------
--- User 的数据是玩家的游戏中永久数据，需要保存的数据
------ 这里直接定义成员变量， 但是不写成员函数， 这样不用reload
--- 一定确保字段名与数据库表中字段名相同
--------------------------------------------------------------------------------------

User = {}
function User.New(userId, openId, machineId)

    local c = {
        --- 玩家基础信息
        userId = userId, --   用户id
        openId = openId,
        machineId = machineId,
        gameId = 0, --   游戏id

        nickName = "玩家" .. tostring(userId), --   昵称
        gender = 0, --   性别
        faceId = 0, --   头像id
        --hairMain = 0, --  头发
        --hairFront = 0,
        --hairBack = 0,
        --hairColor1 = 0,
        --hairColor2 = 0,
        customPlayer = "",  -- 玩家自定义捏脸

        level = 1, --   等级
        exp = 0, --   经验
        HP = 100 ,
        HPMax = 100,
        SP = 100,
        SPMax = 100,

        payTotal = 0, --   充值总金额
        offLineTime = 0, --   离线时间

    }
    return c

end

function User:Reload(c)
    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理
end
