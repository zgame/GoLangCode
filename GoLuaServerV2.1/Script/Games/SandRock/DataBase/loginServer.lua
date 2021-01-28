SandRockLoginDB = {}

---------------- openId 和 uId 对应关系---------------
function SandRockLoginDB.OpenIdInsert(openId, uId)
    local startTime = ZTime.GetOsTimeMillisecond()
    Redis.SaveString(RedisDirAllPlayersOpenID,openId, uId)
    local endTime = ZTime.GetOsTimeMillisecond()
    print("redis db save UId cost:".. endTime - startTime)
end

function SandRockLoginDB.UId(openId)
    local startTime = ZTime.GetOsTimeMillisecond()
    local userId = Redis.GetString(RedisDirAllPlayersOpenID,openId)
    local endTime = ZTime.GetOsTimeMillisecond()
    print("redis db get UId cost:".. endTime - startTime .. ":" .. userId .. ";")
    return userId
end

----------------- User 数据---------------------------------
function SandRockLoginDB.User(userId)
    if userId == nil or userId == "" then
        return nil
    end
    local t = {}
    t.userId = tonumber(userId)

    local re =  MongoDB.Find('User',t)
    return re
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