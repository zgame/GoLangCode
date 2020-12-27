
CMD_CCC = {
    SUB_LOGON = 1, -- 登陆命令GameLogin() , 登录结果GameLoginResult()
    SUB_LOGOUT = 2, -- 登出命令GameLogout() , 登出结果GameLogoutResult()
    SUB_OTHER_LOGON = 3, -- 其他玩家登陆命令UserList()
    SUB_OTHER_LOGOUT = 4, -- 其他玩家登出命令OtherLeaveRoom()
    SUB_ROOM_INFO = 5, -- 游戏信息同步GameInfo()
    SUB_ROOM_LIST = 6, -- 游戏中玩家列表UserList()
    SUB_LOCATION = 7, -- 上报位置信息PlayerLocation(),   同步其他玩家位置PlayerLocation()


}