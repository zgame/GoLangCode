---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/11/6 16:54
---

local CMD_GameServer_pb = require("CMD_GameServer_pb")

----游客登录游戏服申请
function SevLoginGSGuest(buf)
    local msg = CMD_GameServer_pb.CMD_GR_LogonUserID()
    msg:ParseFromString(buf)

    print("gamekind id: ".. msg.kind_id)
    print("user_id id: ".. msg.user_id)

end