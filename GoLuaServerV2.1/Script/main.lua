
-------------------------------------------------------------------------------------
--- 入口点程序 ， 普通程序员不要动
-------------------------------------------------------------------------------------

--print("start lua")

-------------------------------------Logger----------------------------------------
--package.path = "Script/Logger/?.lua;"..package.path
require("Script/Utils/print/zLog")
ZLog.Logger("Game Server Start ....")

-------------------------------------protocol buffer----------------------------------------
package.path = "Script/Protocol/build/?.lua;"..package.path
package.path = "Script/Protocol/protobuf/?.lua;"..package.path
--require("protocol_test")
--Logger("protocol buffer ok")

-------------------------------------Const----------------------------------------
-- 这部分不参与平时的热更新, 如果想更新，临时加入到reload里面即可

require("Script/Utils/byte/zCrypto")
require("Script/Utils/byte/zBit32")
require("Script/Utils/byte/zZip")
require("Script/Utils/database/dbMySql")
require("Script/Utils/database/dbSqlServer")
require("Script/Utils/database/dbMongoDB")
require("Script/Utils/database/dbRedis")
require("Script/Utils/string/json")
require("Script/Utils/string/zStrings")
require("Script/Utils/zClass")
require("Script/Utils/zNetwork")
require("Script/Utils/zRandom")
require("Script/Utils/zTime")
require("Script/Utils/zTimer")
require("Script/globalVar")

-------------------------------------Start----------------------------------------
-- 下面参与热更新，新增加lua文件， 写到reload里面去，不要在这里添加
require("Script/reload")

-- CSV
for _,fileName in ipairs(LuaFiles.CSV) do
    require(fileName)
end

-- Protocol
for _,fileName in ipairs(LuaFiles.Protocol) do
    require(fileName)
end

-- Const
for _,fileName in ipairs(LuaFiles.Const) do
    require(fileName)
end

-- Utils
for _,fileName in ipairs(LuaFiles.Utils) do
    require(fileName)
end

-- DataBase
for _,fileName in ipairs(LuaFiles.DataBase) do
    require(fileName)
end



-- Model
for _,fileName in ipairs(LuaFiles.Manager) do
    require(fileName)
end

-- Games
for _,fileName in ipairs(LuaFiles.Games) do
    require(fileName)
end

-- NetWork
for _,fileName in ipairs(LuaFiles.NetWork) do
    require(fileName)
end
