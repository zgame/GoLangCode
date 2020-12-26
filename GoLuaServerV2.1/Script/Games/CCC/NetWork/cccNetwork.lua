---
---  CCC 的网络分发
---

CCCNetwork = {}

function CCCNetwork.Receive(serverId, userId, mainSgId, subMsgId, data, token)
    local switch={}
    switch[CMD_CCC.SUB_LOGON] = CCCNetworkLogin.SevLoginGSGuest

    switch[subMsgId](serverId, userId, data)




end
