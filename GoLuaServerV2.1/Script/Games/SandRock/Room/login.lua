
----------------------- 玩家操作 ---------------------------------

-- 发消息给同房间的其他玩家，告诉他们你登录了
local function sendLoginToOthers(room, player)
    local userId =  Player.UId(player)
    --print("玩家登录", userId, "房间", room.roomId,"椅子",player.chairId)
    local sendCmd = ProtoGameSandRock.UserList()
    local uu = sendCmd.user:add()
    Player.Copy(player,uu)
    SandRockRoom.SendMsgToOtherUsers(room,CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_OTHER_LOGON,sendCmd,userId)
end

-- 发消息给同房间的其他玩家，告诉他们你登出了
local function sendLogoutToOthers(room, player)
    local userId =  Player.UId(player)
    print("玩家登出", userId, "房间", room.roomId,"椅子",player.chairId)
    local sendCmd = ProtoGameSandRock.OtherLeaveRoom()
    sendCmd.userId = userId
    SandRockRoom.SendMsgToOtherUsers(room,CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_OTHER_LOGOUT,sendCmd,userId)
end

--玩家坐到椅子上
function SandRockRoom:PlayerSeat(chairId, player)
    self.userSeatArray[chairId] = player
    self.userSeatArrayNumber = ZTable.Len(self.userSeatArray)
    --local oldPlayer = GameServer.GetPlayerByUID(Player.UId(player)) -- 把之前的玩家数据取出来
    --if oldPlayer == nil then
    --    self.userSeatArrayNumber = self.userSeatArrayNumber + 1   -- 房间上玩家数量增加
    --end

    player.roomId = self.roomId
    player.chairId = chairId

    GameServer.SetAllPlayerList(Player.UId(player), player)  --创建好之后加入玩家总列表
    sendLoginToOthers(self,player)
    return player
end

--玩家离开椅子
function SandRockRoom:PlayerStandUp(uId)
    local player = GameServer.GetPlayerByUID(uId)
    --ZLog.Logger(uId .. "离开房间" .. player.roomId .. "椅子" .. player.chairId .. "self.roomId" .. self.gameId)
    -- 保存玩家基础数据
    --SaveUserBaseData(player.User)

    GameServer.SetAllPlayerList(Player.UId(player), nil)         -- 清理掉游戏管理的玩家总列表
    self.userSeatArray[player.chairId] = nil                -- 清理掉房间的玩家列表
    self.userSeatArrayNumber = self.userSeatArrayNumber - 1  -- 房间上玩家数量减少
    sendLogoutToOthers(self,player)
    player.roomId = Const.ROOM_CHAIR_NOBODY
    player.chairId = Const.ROOM_CHAIR_NOBODY

    --如果是空房间的话，清理一下房间
    if self:CheckTableEmpty() then
        self:ClearTable()
        Game.ReleaseRoom(self.gameId,self.roomId)    --回收房间
    end
end
