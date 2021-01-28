

--------------------------------------------客户端----------------------------------------
local function build_mac_addr(index)
    local str = "74-D4-36-AD-"
    local i1 = math.floor(index / 256)
    local i2 = math.floor(index % 256)

    local ss1 = string.format("%x",i1)
    if #ss1 == 1 then
        ss1 = "0" .. ss1
    end

    str = str .. ss1
    str = str .. "-"
    local ss2 = string.format("%x",i2)
    if #ss2 == 1 then
        ss2 = "0" .. ss2
    end
    str = str .. ss2

    return str
    --return "74-D4-36-AD-1E-EE"
end

--------------------------登录服务器--------------------------------------------

LoginServer={}

-- 发送登录服务器请求
function LoginServer.SendLogin(serverId)
    print("-----申请登录serverId:----------",serverId)
    local sendCmd = ProtoGameSandRock.GameLogin()
    sendCmd.machineId = build_mac_addr(serverId)

    --sendCmd.openId = "sandRockMac74-D4-36-AD-00-01"
    --sendCmd.openId = "sandRockMacffbd8b4d121a2156ec122bfcc4303a4440519b94Enter-ZSW"
    sendCmd.gameId = Const.GameTypeCCC
    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_LOGON,sendCmd,nil)
end

-- 服务器返回登录服务器成功，下发uid
function LoginServer.Login(serverId, uId,buf)
    print("--------登录成功---UID:",uId)
    local msg = ProtoGameSandRock.GameLoginResult()
    msg:ParseFromString(buf)
    local uid = msg.user.userId

    local user = User:New()
    user.UserId = uid
    local player = Player:New(user)
    GameClient.AddPlayer(serverId,player)

    --print( msg)
    --luaCallGoResisterUID(uid,serverId)
    ZTimer.SetNewTimer("GameSandRock", "RunGame", 2000, serverId, GameSandRock.RunGame)
end

function LoginServer.SendLogout(uId)
    print("--------申请玩家登出-----------")
    local sendCmd = ProtoGameSandRock.GameLogout()
    sendCmd.userId = uId
    Network.SendToUser(uId,CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_LOGOUT,sendCmd,nil)
end

function LoginServer.Logout(serverId, uId,buf)
    print("--------玩家登出-----------")
    Network.Print(ProtoGameSandRock.GameLogoutResult, buf)
end

function LoginServer.GameInfo(serverId, uId, buf)
    print("--------场景信息下发-----------")
    Network.Print(ProtoGameSandRock.GameInfo, buf)
end

function LoginServer.RoomList(serverId, uId, buf)
    print("--------场景玩家下发-----------")
    Network.Print(ProtoGameSandRock.UserList, buf)
end

function LoginServer.OtherLogin(serverId, uId, buf)
    print("--------其他玩家登录-----------")
    Network.Print(ProtoGameSandRock.UserList, buf)
end

function LoginServer.OtherLogout(serverId, uId, buf)
    print("--------其他玩家登出-----------")
    Network.Print(ProtoGameSandRock.OtherLeaveRoom, buf)
end

-----------------------------其他玩家登录----------------------------
