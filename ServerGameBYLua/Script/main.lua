---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/10/9 14:12
---

--print("start lua")

-------------------------------------Logger----------------------------------------
package.path = "Script/Logger/?.lua;"..package.path
require("logger")
Logger("=============================Game Server Start ....==============================")


package.path = "Script/?.lua;"..package.path
require("dumpTable")
require("server")
