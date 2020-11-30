---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/10/12 15:00
---

-------------------------------------热更新----------------------------------------
--- 热更新文件列表
--- 记住： module 和 全局的函数都是可以reload 的，类的写法不行，要注意
-----------------------------------------------------------------------------





--------------------------------文件列表-- 为了方便加载和热更的文件名字统一， 那么建立一个表，统一使用这个表里面的文件列表----------------------------------------------

--数据库
RequireAndReloadDataBaseFiles = {}
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/redisConst")
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/redisGame")
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/redisPlayer")
--table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/redisSystem")
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/sqlStatistic")
--table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/sqlStatisticDaily")
--table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/sqlStatisticLogin")
--table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/sqlStatisticUser")
--table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/sqlStatisticSystemAction")
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/SQLServer/LoadGameServerInfo")
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/SQLServer/LoadUserInfo")
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/SQLServer/LoadUserChatInfo")
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/SQLServer/SaveUserDataInfo")
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/SQLServer/SaveDataBaseLogInfo")
table.insert(RequireAndReloadDataBaseFiles, "Script/DataBase/redisChat")
-- CSV
RequireAndReloadCSVFiles = {}
table.insert(RequireAndReloadCSVFiles, "Script/CSV/mgby_fish_sever")

table.insert(RequireAndReloadCSVFiles, "Script/CSV/mgby_vip")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Card")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Get_Gold_By_Time")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Gold_Mall")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Lottery")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Lottery_Person_Pool")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Lucky_Wheel")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Point")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Room")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Score_All_Pool")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Score_Person_Pool")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Sign_In_Reward")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Tiger")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/VIP")
--table.insert(RequireAndReloadCSVFiles, "Script/CSV/Wheel")

table.insert(RequireAndReloadCSVFiles, "Script/CSV/mgby_item")

-- 小游戏CSV
-- 小海兽
table.insert(RequireAndReloadCSVFiles, "Script/CSV/mgby_monster")



-- Protocol
RequireAndReloadProtocolFiles = {}
--table.insert(RequireAndReloadProtocolFiles, "Script/Protocol/build/CMD_Common_pb")
table.insert(RequireAndReloadProtocolFiles, "Script/Protocol/build/CMD_Game_pb")
table.insert(RequireAndReloadProtocolFiles, "Script/Protocol/build/CMD_GameServer_pb")
--table.insert(RequireAndReloadProtocolFiles, "Script/Protocol/build/CMD_LoginServer_pb")
--table.insert(RequireAndReloadProtocolFiles, "Script/Protocol/build/CMD_Monitor_pb")
--table.insert(RequireAndReloadProtocolFiles, "Script/Protocol/build/CMD_GlobalServer_pb")
--table.insert(RequireAndReloadProtocolFiles, "Script/Protocol/build/CMD_GlobalServer_Inner_pb")
--table.insert(RequireAndReloadProtocolFiles, "Script/Protocol/build/CMD_XHS_Game_pb")


-- Const
RequireAndReloadConstFiles = {}
table.insert(RequireAndReloadConstFiles, "Script/Const/Const")
table.insert(RequireAndReloadConstFiles, "Script/setting")
table.insert(RequireAndReloadConstFiles, "Script/Const/excel")
table.insert(RequireAndReloadConstFiles, "Script/Const/proto")
table.insert(RequireAndReloadConstFiles, "Script/Const/protocol_cs")
table.insert(RequireAndReloadConstFiles, "Script/Const/protocol_gs")
table.insert(RequireAndReloadConstFiles, "Script/Const/protocol_ls")
table.insert(RequireAndReloadConstFiles, "Script/Const/protocol_mo")
table.insert(RequireAndReloadConstFiles, "Script/Const/protocol_cp")
table.insert(RequireAndReloadConstFiles, "Script/Const/protocol_xhs")




-- Utils
RequireAndReloadUtilsFiles = {}
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/luaMySql")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/luaSqlServer")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/zMySql")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/zRandom")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/zRedis")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/zTable")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/zTime")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/Util")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/zCrypto")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/zBit32")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/zMongoDB")
table.insert(RequireAndReloadUtilsFiles,"Script/Utils/sStringFunction")

-- gameManager
RequireAndReloadManagerFiles = {}
table.insert(RequireAndReloadManagerFiles, "Script/GameCommonLogic/commonLogic")
table.insert(RequireAndReloadManagerFiles, "Script/GameManager/player")
table.insert(RequireAndReloadManagerFiles, "Script/GameManager/playerLogic")
table.insert(RequireAndReloadManagerFiles, "Script/GameManager/user")
table.insert(RequireAndReloadManagerFiles, "Script/GameManager/chatUser")
table.insert(RequireAndReloadManagerFiles, "Script/GameManager/games")
table.insert(RequireAndReloadManagerFiles, "Script/GameManager/gameManager")


-- game
RequireAndReloadGamesFiles = {}
table.insert(RequireAndReloadGamesFiles, "Script/Games/baseTable")
table.insert(RequireAndReloadGamesFiles, "Script/Games/BY/byBullet")
table.insert(RequireAndReloadGamesFiles, "Script/Games/BY/byFish")
table.insert(RequireAndReloadGamesFiles, "Script/Games/BY/byFishDistribute")
table.insert(RequireAndReloadGamesFiles, "Script/Games/BY/byTable")
table.insert(RequireAndReloadGamesFiles, "Script/Games/BY/byTableBulletFunc")
table.insert(RequireAndReloadGamesFiles, "Script/Games/BY/byTableFishFunc")
table.insert(RequireAndReloadGamesFiles, "Script/Games/BY/byTablePlayerFunc")
-- 小海兽
table.insert(RequireAndReloadGamesFiles, "Script/Games/XHS/xhsTable")
table.insert(RequireAndReloadGamesFiles, "Script/Games/XHS/xhsGameTable")
table.insert(RequireAndReloadGamesFiles, "Script/Games/XHS/xhsReceiveMsg")
table.insert(RequireAndReloadGamesFiles, "Script/Games/XHS/GameLogic/xhsUserUseSummonGem")
table.insert(RequireAndReloadGamesFiles, "Script/Games/XHS/GameLogic/xhsUserFire")
table.insert(RequireAndReloadGamesFiles, "Script/Games/XHS/GameLogic/xhsUserPuzzle")
table.insert(RequireAndReloadGamesFiles, "Script/Games/XHS/DB/xhsDB")
-- 聊天
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/chatTable")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/chatReceiveMsgFromLoginServer")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/chatReceiveMsgFromClient")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/GameLogic/chatLoginServerLogic")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/GameLogic/chatUserLogin")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/GameLogic/chatUserGetFriendList")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/GameLogic/chatUserSearchPlayer")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/GameLogic/chatUserApplyFriend")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/GameLogic/chatUserDeleteFriend")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/GameLogic/chatUserGetUnReadMessage")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/GameLogic/chatUserSendMessage")
table.insert(RequireAndReloadGamesFiles, "Script/Games/Chat/DB/chatDB")


-- network
RequireAndReloadNetWorkFiles = {}
table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/gameEnter")
table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/gameFire")
table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/loginServer")
table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/registerCorrespondServer")
--table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/mail")
--table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/mainHall")
--table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/monthCard")
table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/network")
--table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/playerInfo")
--table.insert(RequireAndReloadNetWorkFiles, "Script/NetWork/recharge")



-- chatServer
RequireAndReloadChatServerFiles = {}


--- 逻辑相关处理，接收到各服务器消息后进行处理
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogic")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicApply")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicClient")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicFriend")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicLogin")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicSocketLink")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicMessage")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/GameLogic/ChatServerLogicFriendsRelation")


--- netWork 网络消息处理
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/NetWorkReceive/netWorkClientReceive")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/NetWorkReceive/netWorkLoginServerReceive")

--- 其他相关
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ChatServerHotReload")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ChatServerInit")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ChatServerGlobal")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ChatServerTimers")


---modules
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatLinkInfo")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatServerUserInfo")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatWillSaveMessage")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/FriendRelation")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/FriendUnreadMessage")

table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatServerMessage")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Models/ChatUnreadMessageCache")


--- managers
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Managers/ChatUserManager")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/Managers/ChatFriendManager")

--- DB
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/DB/FriendDB")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/DB/ChatDB")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/DB/MessageDB")


--- 聊天服后台好友申请相关
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ApplyController/ChatApplyController")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ApplyController/ApplyControllerTimer")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ApplyController/ApplyControllerDB")
table.insert(RequireAndReloadChatServerFiles, "Script/Games/ChatServer/ApplyController/ChatServerApplyController")



-- 热更新全部的逻辑代码，需要自己控制， 切记玩家数据部分不要加进去，不然会重置玩家数据，如果你是保存型代码，容易导致玩家清档
function ReloadAll()
    --if true then
    --    return
    --end
    Logger("------------------start reload all---------------------------")


     --HotReload
    ReloadFile("Script/HotReload/hotReload")
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

    -- GameManager
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
                -- 遍历所有游戏，所有桌子， 所有鱼，所有子弹，所有生成鱼池， 因为这些都是类， 已经生成的对象需要刷新函数
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
    Logger("当前服务器人数:"..AllPlayerListNumber)
end



function ReloadFile(module_name)
    package.loaded[module_name] = nil
    require(module_name)
end


--- 将旧表替换成新表
function OldModuleReloadByNew(module_name)
    --- 先把旧的模块保存
    local oldmodule =  package.loaded[module_name] or {}
    --- 清空模块
    package.loaded[module_name] = nil
    --- 加载新文件
    require(module_name)
    --- 获取新文件
    local newModule = package.loaded[module_name]
    ---将新模块的内容附加,旧模块原来存在的key不清除
    for k,v in pairs(newModule) do
        oldmodule[k] = v
    end
    package.loaded[module_name] = oldmodule
    return  oldmodule
end