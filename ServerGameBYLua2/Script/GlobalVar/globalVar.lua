---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/21 16:20
---



ServerIP_Port = ""      -- 当前服务器的地址和端口号

--GoRoutineMax = 4      --桌子使用的go routine函数上限
GoRunTableAllList = {}   -- 所有的go routine列表，桌子的run函数在里面

AllGamesList = {}   -- gameType , game
AllPlayerList = {}   -- 所有玩家列表   key  userId , value player

ALLUserUUID = 0   -- 玩家uid的自增