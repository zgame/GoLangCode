
SandRockUserDB = {}

function SandRockUserDB.InfoUpdate(userId, nickName, gender)
    local c = {}
    c.userId = userId

    local u = {}
    u.userId = userId
    if nickName ~= nil then
        u.nickName = nickName
    end
    if gender ~= nil then
        u.gender = gender
    end
    MongoDB.Update('User', c,u)
end
