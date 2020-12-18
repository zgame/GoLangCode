

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
    local sendCmd = ProtoGameCCC.GameLogin()

    sendCmd.machineId = build_mac_addr(serverId)
    Network.Send(serverId, CMD_MAIN.MDM_GAME_CCC, CMD_CCC.SUB_LOGON,sendCmd,nil,true)

    print("-----申请登录serverId:----------",serverId)
end

-- 服务器返回登录服务器成功，下发uid
function LoginServer.LoginGameServer(serverId, buf)
    local msg = ProtoGameCCC.GameLoginResult()
    msg:ParseFromString(buf)
    local uid = msg.user.userId

    print( msg)

    --luaCallGoResisterUID(uid,serverId)
    print("--------登录成功---UID:",uid)

    ZTimer.SetNewTimer("GameCCC", "RunGame", 500, GameCCC.RunGame)
end