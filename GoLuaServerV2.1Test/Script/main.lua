---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/10/9 14:12
---

--print("start lua")

-------------------------------------Logger----------------------------------------
--package.path = "Script/Logger/?.lua;"..package.path
require("Script/Utils/logger")
ZLog.Logger("========Game Server Start ....========")

---------------------------------------protocol buffer----------------------------------------
--package.path = "Script/Protocol/build/?.lua;"..package.path
--package.path = "Script/Protocol/protobuf/?.lua;"..package.path
----require("protocol_test")
--Logger("protocol buffer ok")

package.path = "Script/?.lua;"..package.path
require("server")
