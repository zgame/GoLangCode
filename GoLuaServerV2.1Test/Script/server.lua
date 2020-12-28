
-------------------------------------说明----------------------------------------
---本文件是服务器启动的时候加载使用
-----------------------------------------------------------------------------


-------------------------------------CSV----------------------------------------
package.path = "Script/CSV/?.lua;"..package.path
--data_csv1 = require("mgby_fish_sever")
--print(data_csv1[134].type)
--print(data_csv1[134].min_force_killed_bullet)
--Logger("csv ok")

-------------------------------------protocol buffer----------------------------------------
package.path = "Script/Protocol/build/?.lua;"..package.path
package.path = "Script/Protocol/protobuf/?.lua;"..package.path
--require("protocol_test")
--Logger("protocol buffer ok")

-------------------------------------Const----------------------------------------

package.path = "Script/NetWork/?.lua;"..package.path
package.path = "Script/GameCommonLogic/?.lua;"..package.path
package.path = "Script/HotReload/?.lua;"..package.path
package.path = "Script/Const/?.lua;"..package.path
package.path = "Script/Utils/?.lua;"..package.path
--package.path = "Script/GlobalVar/?.lua;"..package.path
package.path = "Script/GameManager/?.lua;"..package.path

require("Const")
require("Excel")
require("proto")
require("constCmd")
require("constCmdGame")
require("constCmdServer")


-------------------------------------Const----------------------------------------
--require("commonLogic")
--require("hotReload")
require("zTime")
require("zTimer")
require("zRandom")
require("dumpTable")
require("zTable")


-------------------------------------NetWork----------------------------------------


require("loginServer")
require("location")
require("gameSandRock")
require("network")

-------------------------------------GameManager----------------------------------------

require("globalVar")
require("player")
require("user")


require("gameManager")

ZJson = require("Json")