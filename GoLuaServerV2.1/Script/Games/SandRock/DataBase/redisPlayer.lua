



----------------------------玩家的数据-----------------------------
-- function RedisSavePlayer(User)
--     Redis.SaveString(RedisDirAllPlayers..User.UserID,User.UserID, ZJson.encode(User))
-- end

-- function RedisGetPlayer(uid)
--     --return  ZJson.decode(RedisGetString(RedisDirAllPlayers..uid, uid))
--     local json =  ZJson.decode(Redis.GetString(RedisDirAllPlayers..uid, uid))
--     local user = User:New(uid,"")

--     --printTable(user)
--     -- 下面做一个遍历， 这个遍历的好处是 ，如果游戏更新，增加了字段， 不需要维护， 因为玩家登录的时候， 没有的字段会默认为初始值 ，再保存，新字段就有了
--     if json ~= nil then
--         for k,v in pairs(user) do
--             if json[k] ~= nil then
--                 user[k] = json[k]       -- 将数据库保存数据赋值
--             end
--             --if type(v) == "string" then
--             --    print(k.."    " .. v)
--             --else
--             --    print(k .."    "..type(v))
--             --end
--         end
--     end
--     return  user
-- end

----------------------------玩家登录 open id -----------------------------
-- function RedisSavePlayerLogin(openId,Uid)
--     Redis.SaveString(RedisDirAllPlayersLogin..openId,openId, Uid)
-- end

-- function RedisGetPlayerLogin(openId)
--     return  Redis.GetString(RedisDirAllPlayersLogin..openId, openId)     -- 返回Uid
-- end