
SandRockUserDB = {}

function SandRockUserDB.NickNameUpdate(userId, nickName)
    local c = {}
    c.userId = userId

    local u = {}
    u.userId = userId
    u.nickName = nickName
    MongoDB.Update('User', c,u)
end
