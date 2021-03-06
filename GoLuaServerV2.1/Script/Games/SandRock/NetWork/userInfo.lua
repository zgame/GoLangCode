SandRockUserInfoNet ={}

function SandRockUserInfoNet.PlayerInfo(serverId, userId, buf)
    local msg = ProtoGameSandRock.PlayerInfo()
    msg:ParseFromString(buf)

    local nickName = msg.nickName
    if nickName == "" then
        return
    end
    local gender = msg.gender

    -- 后面要增加敏感词的判断, 大部分放到客户端显示的时候

    local player = GameServer.GetPlayerByUID(userId)
    Player.NickNameSet(player, nickName)
    player.user.gender = gender
    SandRockUserDB.InfoUpdate(userId, nickName,gender)     -- 保存到db

    local sendCmd = ProtoGameSandRock.PlayerInfo()
    sendCmd.nickName = Player.NickNameGet(player)
    sendCmd.userId =  Player.UId(player)
    sendCmd.gender = player.user.gender

    local room = GameServer.GetRoomByUserId(userId)
    SandRockRoom.SendMsgToAllUsers(room, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_PLAYER_INFO, sendCmd)
end