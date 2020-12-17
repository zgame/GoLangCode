



-----------------------------------客户端开启-------------------------------------
-- 服务器开始创建各个游戏，这里的游戏都是多人的游戏， 如果是单人游戏，玩家自己创建即可
function GoCallLuaStartGamesServers(serverId)
    --print("客户端开启")
    -- 开始连接服务器
    LoginServer.SendLogin(serverId)
end


-- 根据user uid 返回user的句柄
function GetPlayerByUID(uid)
    return AllPlayerList[tostring(uid)]     --- 这里一定要注意， goperlua在用数字作为key的时候会默认为数组，内存消耗惊人， 所以要用string
end
function SetAllPlayerList(userId,value)
    AllPlayerList[tostring(userId)] = value
end

-- go通知lua玩家掉线了
function GoCallLuaPlayerNetworkBroken(uid)
    --Logger("go 通知："..uid .. "  掉线了")
    local player = GetPlayerByUID(uid)

    print("玩家被t了"..player.User.UserId)

end
