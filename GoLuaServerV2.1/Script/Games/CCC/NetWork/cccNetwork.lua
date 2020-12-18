---
---  CCC 的网络分发
---

CCCNetwork = {}

function CCCNetwork.Receive(serverId, userId, mainSgId, subMsgId, data, token)
    if subMsgId == CMD_CCC.SUB_LOGON then
        CCCNetworkLogin.SevLoginGSGuest(serverId, data)
    end
end
