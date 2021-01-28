Setting = {
    -----------------------------------------------------------------
    -- 服务器地址的定义
    -----------------------------------------------------------------
    --ConstServerAddressLogServer = "192.168.0.218:8330"      -- 日志服务器地址
    --ConstServerAddressCorrespondServer = "127.0.0.1:8310"   -- 协调服务器地址
    --ConstServerAddressChatServer = "192.168.101.58"         -- 聊天服的地址

    ConstMainCenterServer = "10.96.8.121:9888"    ,           -- 主中心服的地址

    -----------------------------------------------------------------
    -- 数据库地址的定义
    -----------------------------------------------------------------
    ---------------------- redis -------------------
    --RedisAddress = "47.92.150.31:6379"  , -- redis 服务器地址远程
    RedisAddress = "10.96.8.209:6379"  , -- redis 服务器地址 局域网
    RedisPass = "LncDnQaR502NWaFdCVXMeKacglgnf3"     ,           -- redis 密码
    ---------------------- mongodb -------------------
    MongoAddress = "10.96.8.209:27017",
    MongoDatabase = "SandRock",
    MongoUser = "patheaDev",
    MongoPass = "LncDnQaR502NWaFdCVXMeKacglgnf3",
    ---------------------- mySql -------------------
    MySqlServerIP = "47.92.150.31",
    MySqlServerPort = "3306",
    MySqlDatabase = "sandrock",
    MySqlUid = "patheaDev",
    MySqlPwd = "LncDnQaR502NWaFdCVXMeKacglgnf3",
    ----------------------- sql server --------------------------
    --SqlServerIP = "192.168.0.207"
    --SqlServerDatabase = "DataBase"
    --SqlServerUid = "sa"
    --SqlServerPwd = "Aa123456"



    -------------------------------数据库初始结构-------------------------------
    -- mongo db init
    -- use SandRock
    -- db.createCollection("User")

    -- redis db init
    -- set Ping Pong

}