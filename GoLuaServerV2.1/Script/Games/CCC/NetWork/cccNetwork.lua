---
---  CCC 的网络分发
---

CCCNetWork = {}

function CCCNetWork.Receive(serverId, userId, mainSgId, subMsgId, data, token)
    if subMsgId == CMD_CCC.SUB_LOGON then
        CCCNetWorkLogin.SevLoginGSGuest(serverId, data)
    end
end
function CCCNetWork.ReceiveUdp(serverAddr, mainSgId, subMsgId, data)
    if subMsgId == CMD_CCC.SUB_LOGON then
        CCCNetWorkLogin.SevLoginGSGuest(serverAddr, data)
    end
end