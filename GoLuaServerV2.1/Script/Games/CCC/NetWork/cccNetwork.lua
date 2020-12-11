---
---  CCC 的网络分发
---

CCCNetWork = {}

function CCCNetWork.Receive(serverId, userId, mainSgId, subMsgId, data, token)
    if subMsgId == CMD_CCC.SUB_LOGON then
        CCCNetWorkLogin.SevLoginGSGuest(serverId, data)
    end
end