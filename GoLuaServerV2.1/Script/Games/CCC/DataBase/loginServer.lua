CCCLoginDB = {}

---------------- openId 和 uId 对应关系---------------
function CCCLoginDB.OpenIdInsert(openId,uId)
    local t={}
    t.openId = openId
    t.uId = uId
    MongoDB.Insert('OpenId',t)
end

function CCCLoginDB.UId(openId)
    local t = {}
    t.openId = openId
    local result = MongoDB.Find('OpenId',t)
    if result == nil then
        return nil
    else
        return result.uId
    end
end

----------------- User 数据---------------------------------
function CCCLoginDB.User(userId)
    if userId == nil then
        return nil
    end
    local t = {}
    t.userId = userId
    return MongoDB.Find('User',t)
end

function CCCLoginDB.UserInsert(user)
    MongoDB.Insert('User',user)
end

function CCCLoginDB.UserUpdate(userId,user)
    local t = {}
    t.userId = userId
    MongoDB.Update('User',t,user)
end


---------------------------------------------------------