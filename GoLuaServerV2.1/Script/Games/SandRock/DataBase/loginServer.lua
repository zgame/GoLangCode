SandRockLoginDB = {}

---------------- openId 和 uId 对应关系---------------
function SandRockLoginDB.OpenIdInsert(openId, uId)
    local t={}
    t.openId = openId
    t.uerId = uId
    MongoDB.Insert('OpenId',t)
end

function SandRockLoginDB.UId(openId)
    local t = {}
    t.openId = openId
    local result = MongoDB.Find('OpenId',t)
    if result == nil then
        return nil
    else
        return result.uerId
    end
end

----------------- User 数据---------------------------------
function SandRockLoginDB.User(userId)
    if userId == nil then
        return nil
    end
    local t = {}
    t.userId = userId
    return MongoDB.Find('User',t)
end

function SandRockLoginDB.UserInsert(user)
    MongoDB.Insert('User',user)
end

function SandRockLoginDB.UserUpdate(userId, user)
    local t = {}
    t.userId = userId
    MongoDB.Update('User',t,user)
end


---------------------------------------------------------