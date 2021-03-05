
SandRockUserDB = {}

-- 保存用户信息
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


-- 用户数据全保存
function SandRockUserDB.UserUpdate(userId, user)
    local t = {}
    t.userId = userId
    MongoDB.Update('User',t,user)
end


-- 保存用户背包
function SandRockUserDB.PackageUpdate(userId, package)
    local c = {}
    c.userId = userId

    local u = {}
    u.package = package
    u.userId = userId
    MongoDB.Update('User',c,u)
end
