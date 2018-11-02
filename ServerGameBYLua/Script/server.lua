---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/1 14:34
---


-------------------------------------Logger----------------------------------------
package.path = "Script/Logger/?.lua;"..package.path
require("logger")
Logger("=============================Game Server Start ....==============================")

-------------------------------------system----------------------------------------

package.path = "Script/NetWork/?.lua;"..package.path
package.path = "Script/GameCommonLogic/?.lua;"..package.path
package.path = "Script/HotReload/?.lua;"..package.path
require("network")
require("commonLogic")
require("hotReload")
require("Const")
Logger("system ok")

-------------------------------------protocol buffer----------------------------------------
package.path = "Script/Protocol/build/?.lua;"..package.path
package.path = "Script/Protocol/protobuf/?.lua;"..package.path
require("protocol_test")
Logger("protocol buffer ok")
-------------------------------------CSV----------------------------------------
package.path = "Script/CSV/?.lua;"..package.path
data_csv1 = require("mgby_fish_sever")
--print(data_csv1[134].type)
--print(data_csv1[134].min_force_killed_bullet)
Logger("csv ok")


-------------------------------------GameManager----------------------------------------
package.path = "Script/GameManager/?.lua;"..package.path
require("gameManager")



