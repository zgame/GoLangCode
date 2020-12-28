---
---  SandRock 的网络分发
---

SandRockNetwork = {}

function SandRockNetwork.Receive(serverId, userId, mainSgId, subMsgId, buf, token)
    local switch={}
    switch[CMD_SAND_ROCK.SUB_LOGON] = SandRockLogin.Login
    switch[CMD_SAND_ROCK.SUB_LOGOUT] = SandRockLogin.Logout
    switch[CMD_SAND_ROCK.SUB_LOCATION] = SandRockLocation.Location
    
    switch[subMsgId](serverId, userId, buf)




end
