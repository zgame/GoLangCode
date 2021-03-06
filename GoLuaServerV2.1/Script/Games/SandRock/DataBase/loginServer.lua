SandRockLoginDB = {}

---------------- openId 和 uId 对应关系---------------
function SandRockLoginDB.OpenIdInsert(openId, uId)
    --local startTime = ZTime.GetOsTimeMillisecond()
    local t={}
    t.openId = openId
    t.uerId = uId
    MongoDB.Insert('OpenId',t)

    --Redis.SaveString(RedisDirAllPlayersOpenID,openId, uId)
    --local endTime = ZTime.GetOsTimeMillisecond()
    --print("redis db save UId cost:".. endTime - startTime)
end

function SandRockLoginDB.UId(openId)
    --local startTime = ZTime.GetOsTimeMillisecond()
    local t = {}
    t.openId = openId
    local result = MongoDB.Find('OpenId',t)
    if result == nil then
        return nil
    else
        return result.uerId
    end
    --local userId = Redis.GetString(RedisDirAllPlayersOpenID,openId)
    --local endTime = ZTime.GetOsTimeMillisecond()
    --print("redis db get UId cost:".. endTime - startTime .. ":" .. userId .. ";")
    --return userId
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


---------------------------------------------------------