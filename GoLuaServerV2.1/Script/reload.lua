

-------------------------------------热更新----------------------------------------
--- 热更新文件列表
--- 记住： module 和 全局的函数都是可以reload 的，类的写法不行，要注意
-----------------------------------------------------------------------------

LuaFiles = {}
local f
--------------------------------文件列表-- 为了方便加载和热更的文件名字统一， 那么建立一个表，统一使用这个表里面的文件列表----------------------------------------------

--数据库
LuaFiles.DataBase = {}
f = LuaFiles.DataBase
table.insert(f, "Script/Games/SandRock/DataBase/redisConst")
table.insert(f, "Script/Games/SandRock/DataBase/redisGame")
table.insert(f, "Script/Games/SandRock/DataBase/redisPlayer")
table.insert(f, "Script/Games/SandRock/DataBase/sqlStatistic")
-- table.insert(f, "Script/Games/SandRock/DataBase/SQLServer/LoadGameServerInfo")
-- table.insert(f, "Script/Games/SandRock/DataBase/SQLServer/LoadUserInfo")
-- table.insert(f, "Script/Games/SandRock/DataBase/SQLServer/LoadUserChatInfo")
-- table.insert(f, "Script/Games/SandRock/DataBase/SQLServer/SaveUserDataInfo")
-- table.insert(f, "Script/Games/SandRock/DataBase/SQLServer/SaveDataBaseLogInfo")
-- table.insert(f, "Script/Games/SandRock/DataBase/redisChat")


-- CSV
LuaFiles.CSV = {}
f = LuaFiles.CSV
table.insert(f, "Script/CSV/mgby_vip")
table.insert(f, "Script/CSV/mgby_item")
table.insert(f, "Script/CSV/mgby_monster")



-- Protocol
LuaFiles.Protocol = {}
f = LuaFiles.Protocol
--table.insert(f, "Script/Protocol/build/Proto_Game_CCC_pb")
table.insert(f, "Script/Protocol/build/proto")
--table.insert(f, "Script/Games/proto")

-- Const
LuaFiles.Const = {}
f = LuaFiles.Const
table.insert(f, "Script/Const/const")
table.insert(f, "Script/Const/constCmd")
table.insert(f, "Script/Const/constCmdGame")
table.insert(f, "Script/Const/constCmdServer")
table.insert(f, "Script/setting")
table.insert(f, "Script/Const/excel")





-- Utils
LuaFiles.Utils = {}
f = LuaFiles.Utils
table.insert(f, "Script/Utils/zTable")
table.insert(f, "Script/Utils/Enum")
table.insert(f, "Script/Utils/string/splitString")


-- Model
LuaFiles.Manager = {}
f = LuaFiles.Manager
table.insert(f, "Script/serverStart")           --  服务器入口点
table.insert(f, "Script/Server/commonLogic")
table.insert(f, "Script/Server/centerServer")
table.insert(f, "Script/Games/Model/player")
table.insert(f, "Script/Games/Model/playerLogic")
table.insert(f, "Script/Games/Model/user")
table.insert(f, "Script/Games/gameServer") --  游戏服务器入口点
table.insert(f, "Script/Games/games")
--table.insert(f, "Script/Games/baseRoom")
table.insert(f, "Script/Games/gameNetwork")     -- 游戏网络分发入口

-- game
LuaFiles.Games = {}
f = LuaFiles.Games
--ccc
table.insert(f, "Script/Games/SandRock/sandRockRoom")
table.insert(f, "Script/Games/SandRock/DataBase/loginServer")



-- network
LuaFiles.NetWork = {}
f = LuaFiles.NetWork
table.insert(f, "Script/network")       -- 网络分发入口点n
-- table.insert(f, "Script/Games/SandRock/NetWork/gameEnter")
-- table.insert(f, "Script/Games/SandRock/NetWork/gameFire")
table.insert(f, "Script/Games/SandRock/NetWork/loginServer")
table.insert(f, "Script/Games/SandRock/NetWork/location")
table.insert(f, "Script/Games/SandRock/NetWork/network")






---- chatServer
--RequireAndReloadChatServerFiles = {}
--
--
----- 逻辑相关处理，接收到各服务器消息后进行处理
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogic")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicApply")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicClient")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicFriend")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicLogin")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicSocketLink")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicMessage")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicFriendsRelation")
--
--
----- netWork 网络消息处理
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/NetWorkReceive/netWorkClientReceive")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/NetWorkReceive/netWorkLoginServerReceive")
--
----- 其他相关
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ChatServerHotReload")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ChatServerInit")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ChatServerGlobal")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ChatServerTimers")
--
--
-----modules
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatLinkInfo")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatServerUserInfo")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatWillSaveMessage")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/FriendRelation")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/FriendUnreadMessage")
--
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatServerMessage")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatUnreadMessageCache")
--
--
----- managers
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Managers/ChatUserManager")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Managers/ChatFriendManager")
--
----- database
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/database/FriendDB")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/database/ChatDB")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/database/MessageDB")
--
--
----- 聊天服后台好友申请相关
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ApplyController/ChatApplyController")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ApplyController/ApplyControllerTimer")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ApplyController/ApplyControllerDB")
--table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ApplyController/ChatServerApplyController")
--


-- 热更新全部的逻辑代码，需要自己控制， 切记玩家数据部分不要加进去，不然会重置玩家数据，如果你是保存型代码，容易导致玩家清档
function ReloadAll()
    --if true then
    --    return
    --end
    ZLog.Logger("------------------start reload all---------------------------")


    --HotReload
    --ReloadFile("Script/reload")
    --[[
    -- CSV
    for _,fileName in ipairs(RequireAndReloadCSVFiles) do
        ReloadFile(fileName)
    end

    -- Protocol
    --local cost = GetOsTimeMillisecond()
    for _,fileName in ipairs(RequireAndReloadProtocolFiles) do
        ReloadFile(fileName)
    end
    --print(string.format("Protocol reload end %d", GetOsTimeMillisecond() - cost))

    -- Const
    for _,fileName in ipairs(RequireAndReloadConstFiles) do
        ReloadFile(fileName)
    end

    -- Utils
    for _,fileName in ipairs(RequireAndReloadUtilsFiles) do
        ReloadFile(fileName)
    end

    -- DataBase
    for _,fileName in ipairs(RequireAndReloadDataBaseFiles) do
        ReloadFile(fileName)
    end


    -- Games
    for _,fileName in ipairs(RequireAndReloadGamesFiles) do
        ReloadFile(fileName)
    end

    -- Model
    for _,fileName in ipairs(RequireAndReloadManagerFiles) do
        ReloadFile(fileName)
    end


    -- GlobalVar    -- 注意事项： 因为这里有游戏的全局变量，所以不能reload

    -- NetWork
    for _,fileName in ipairs(RequireAndReloadNetWorkFiles) do
        ReloadFile(fileName)
    end


    --    -- 已经生成的对象需要刷新函数
    for _, game in pairs(AllGamesList) do
        Game:Reload(game)
        -- 捕鱼游戏
        if game.GameTypeID == GameTypeBY or game.GameTypeID == GameTypeBY30 then
            for _, table in pairs(game.AllTableList) do
                -- 遍历所有游戏，所有房间， 所有鱼，所有子弹，所有生成鱼池， 因为这些都是类， 已经生成的对象需要刷新函数
                ByTable:Reload(table)
                for _, fish in pairs(table.FishArray) do
                    Fish:Reload(fish)
                end
                for _, bullet in pairs(table.BulletArray) do
                    Bullet:Reload(bullet)
                end
                for _, distribute in pairs(table.DistributeArray) do
                    FishDistribute:Reload(distribute)
                end
                for _, bossDistribute in pairs(table.BossDistributeArray) do
                    FishDistribute:Reload(bossDistribute)
                end
            end
        end
        -- 小海兽
        if game.GameTypeID == GameTypeXHS then
            for _, table in pairs(game.AllTableList) do
                XhsTable:Reload(table)
            end
        end
        -- 其他游戏
    end

    -- 所有玩家的数据刷新（如果结构定义有修改的话）
    for _,player in pairs(AllPlayerList) do
        User:Reload(player.User)
        Player:Reload(player)
    end


    -- main
    --reloadFile("Script/server")

    --collectgarbage()
    ]]--

    --- 聊天服重新加載腳本
    -- ChatServerReload()
    --Logger("当前服务器人数:"..GlobalVar.AllPlayerListNumber)
end

function ReloadFile(module_name)
    package.loaded[module_name] = nil
    require(module_name)
end

--- 将旧表替换成新表
function OldModuleReloadByNew(module_name)
    --- 先把旧的模块保存
    local oldmodule = package.loaded[module_name] or {}
    --- 清空模块
    package.loaded[module_name] = nil
    --- 加载新文件
    require(module_name)
    --- 获取新文件
    local newModule = package.loaded[module_name]
    ---将新模块的内容附加,旧模块原来存在的key不清除
    for k, v in pairs(newModule) do
        oldmodule[k] = v
    end
    package.loaded[module_name] = oldmodule
    return oldmodule
end