
--------------------------------------------------------------------------------------
--- player 的数据是玩家的游戏中的数据
--- 这里直接定义成员变量， 但是不写成员函数，就不用reload
--- 这里要注意一点：  当热更新的时候， 所有的已经创建好的player对象是存在的， 结构也是老的结构， 如果你增加了字段，修改了字段，需要进行reload的单独数据处理
--------------------------------------------------------------------------------------

Player = Class:extend()
function Player:New(user)

    self.user = user                    -- user数据
    self.gameId = 0                   -- 游戏类型

    self.roomId = Const.ROOM_CHAIR_NOBODY   -- 房间id
    self.chairId = Const.ROOM_CHAIR_NOBODY   -- 椅子id

    self.scene = nil            -- 玩家所在场景,  同步的时候用来判断的

    self.netWorkState = true            -- 网络状态正常
    self.netWorkCloseTimer = 0         -- 等待玩家断线重连的时间倒计时

end

function Player:Reload(c)

    -- 如果热更新有改动成员变量的定义的话， 下面需要进行成员变量的处理
    -- 比如 1 增加了字段， 那么你需要将老数据进行， 新字段的初始化
    -- 比如 2 删除了字段， 那么你需要将老数据进行， 老字段=nil
    -- 比如 3 修改了字段， 那么你需要将老数据进行， 老字段=nil， 新字段初始化或者进行赋值处理

end