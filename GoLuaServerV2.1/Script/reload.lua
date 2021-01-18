

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

-- CSV
LuaFiles.CSV = {}
f = LuaFiles.CSV
table.insert(f, "Script/CSV/generateGroup")
table.insert(f, "Script/CSV/generateItem")
table.insert(f, "Script/CSV/item")
table.insert(f, "Script/CSV/itemFunctions")
table.insert(f, "Script/CSV/resourcePickArea")
table.insert(f, "Script/CSV/resourcePickType")
table.insert(f, "Script/CSV/resourceTerrainArea")
table.insert(f, "Script/CSV/resourceTerrainType")



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
table.insert(f, "Script/Const/constSandRock")
table.insert(f, "Script/Const/constCmd")
table.insert(f, "Script/Const/constCmdGame")
table.insert(f, "Script/Const/constCmdServer")
table.insert(f, "Script/setting")
--table.insert(f, "Script/Const/excel")


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
table.insert(f, "Script/Games/Model/playerBag")
table.insert(f, "Script/Games/Model/playerLogic")
table.insert(f, "Script/Games/Model/playerValue")
table.insert(f, "Script/Games/Model/user")
table.insert(f, "Script/Games/gameServer") --  游戏服务器入口点
table.insert(f, "Script/Games/games")
--table.insert(f, "Script/Games/baseRoom")
table.insert(f, "Script/Games/gameNetwork")     -- 游戏网络分发入口

-- game
LuaFiles.Games = {}
f = LuaFiles.Games
--ccc
table.insert(f, "Script/Games/SandRock/Room/sandRockRoom")
table.insert(f, "Script/Games/SandRock/Room/login")
table.insert(f, "Script/Games/SandRock/Room/location")
table.insert(f, "Script/Games/SandRock/Room/resourcePick")
table.insert(f, "Script/Games/SandRock/Room/resourceTerrain")
table.insert(f, "Script/Games/SandRock/Room/weather")

table.insert(f, "Script/Games/SandRock/Excel/generatorItem")
table.insert(f, "Script/Games/SandRock/Excel/item")
table.insert(f, "Script/Games/SandRock/Excel/resourcePick")
table.insert(f, "Script/Games/SandRock/Excel/resourceTerrain")

table.insert(f, "Script/Games/SandRock/DataBase/loginServer")


-- network
LuaFiles.NetWork = {}
f = LuaFiles.NetWork
table.insert(f, "Script/network")       -- 网络分发入口点n
-- table.insert(f, "Script/Games/SandRock/NetWork/gameEnter")
-- table.insert(f, "Script/Games/SandRock/NetWork/gameFire")
table.insert(f, "Script/Games/SandRock/NetWork/loginServer")
table.insert(f, "Script/Games/SandRock/NetWork/location")
table.insert(f, "Script/Games/SandRock/NetWork/sleep")
table.insert(f, "Script/Games/SandRock/NetWork/resourcePick")
table.insert(f, "Script/Games/SandRock/NetWork/resourceTerrain")
table.insert(f, "Script/Games/SandRock/NetWork/network")



-- 热更新全部的逻辑代码，需要自己控制， 切记玩家数据部分不要加进去，不然会重置玩家数据，如果你是保存型代码，容易导致玩家清档
function ReloadAll()
    --HotReload
    ReloadFile("Script/reload")
    --ZLog.Logger("------------------***** start reload all ****---------------------------")

    --for _,fileName in ipairs(LuaFiles.Games) do
    --    ReloadFile(fileName)
    --end
    --for _, game in pairs(GlobalVar.AllGamesList) do
    --    local roomClass = GameServer.GetRoomClass(game.gameId)
    --    for _, room in pairs(game.allRoomList) do
    --        --room:Reload(roomClass)
    --        roomClass:Reload(room)
    --    end
    --end
    --for _,fileName in ipairs(LuaFiles.NetWork) do
    --    ReloadFile(fileName)
    --end

    ---- 所有玩家的数据刷新（如果结构定义有修改的话）
    --for _,player in pairs(AllPlayerList) do
    --    User:Reload(player.User)
    --    Player:Reload(player)
    --end


    collectgarbage()
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