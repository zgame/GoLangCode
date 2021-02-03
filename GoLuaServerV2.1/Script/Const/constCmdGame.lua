
CMD_SAND_ROCK = {
    SUB_LOGON = 1, -- 登陆命令GameLogin() , 登录结果GameLoginResult()
    SUB_LOGOUT = 2, -- 登出命令GameLogout() , 登出结果GameLogoutResult()
    SUB_OTHER_LOGON = 3, -- 其他玩家登陆命令UserList()
    SUB_OTHER_LOGOUT = 4, -- 其他玩家登出命令OtherLeaveRoom()
    SUB_ROOM_INFO = 5, -- 游戏信息同步GameInfo()
    SUB_ROOM_LIST = 6, -- 游戏中玩家列表UserList()
    SUB_LOCATION = 7, -- 上报位置信息PlayerLocation()
    SUB_OTHER_LOCATION = 8, -- 同步其他玩家位置PlayerLocation()
    SUB_SLEEP = 9 ,         -- 睡觉Sleep()
    SUB_RESOURCE_POINT = 10 ,   -- 刷新资源ResourceUpdate()
    SUB_RESOURCE_GET = 11 ,   -- 采集资源ResourceGet()  , 获得结果ItemGet()
    SUB_RESOURCE_TERRAIN = 12 ,   -- 刷新资源ResourceTerrainUpdate()
    SUB_RESOURCE_TERRAIN_GET = 13 ,   -- 采集资源ResourceTerrainGet()  , 获得结果ItemGet()

    SUB_CREATION_MACHINE = 14,     --  组装机器  CreationMachine() ,  获得结果ItemGet()
    SUB_CREATION_ITEM = 15,     --  合成道具  CreationItem() ,  获得结果ItemGet()
    SUB_CREATION_RECYCLE = 16,     --  回收道具  CreationRecycle() ,  获得结果ItemGet()
    SUB_CREATION_COOKING = 17,     --  回收道具  CreationCooking() ,  获得结果ItemGet()

    SUB_PLAYER_HOLD = 18,     --  手持同步  PlayerHold()


}

