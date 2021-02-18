UserInfo ={}


function UserInfo.GetUserInfo(serverId,userId, buf)
    print("收到用户信息的修改")
    local msg = ProtoGameSandRock.PlayerInfo()
    msg:ParseFromString(buf)
    print(msg)
end

function UserInfo.SendUserInfo(serverId)
    local sendCmd = ProtoGameSandRock.PlayerInfo()
    sendCmd.nickName = "修改之后的名字"
    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_PLAYER_INFO,sendCmd,nil)

end