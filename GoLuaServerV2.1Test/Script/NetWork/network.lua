----------------------------------------------------------------------
---发送消息
----------------------------------------------------------------------

Network = {}

------- 玩家自己的网络发送函数
function Network.Send(myServerId, msgId, subMsgId, sendCmd, err, udp)
    --return LuaNetWorkSendToUser(0,serverId,msgId,subMsgId,sendCmd,err)      -- userId 如果是0的话， 就是给玩家自己回消息 ，这是在go那边定义的
    local buffer = ""
    if sendCmd ~= nil then
        buffer = sendCmd:SerializeToString()
    end

    if err == nil then
        err = ""
    end
    if udp == nil then
        return luaCallGoNetWorkSend(0, myServerId, msgId, subMsgId, buffer, err)       -- 返回结果 true 发送成功  false 发送失败
    else
        return luaCallGoNetWorkSendUdp(0, myServerId, msgId, subMsgId, buffer, err)       -- 返回结果 true 发送成功  false 发送失败
    end
end


-- 发送消息给其他玩家
function Network.SendToUser(userId, msgId, subMsgId, sendCmd, err)
    local buffer = ""
    if sendCmd ~= nil then
        buffer = sendCmd:SerializeToString()
    end

    if err == nil then
        err = ""
    end
    return luaCallGoNetWorkSend(userId, 0, msgId, subMsgId, buffer, err)       -- 返回结果 true 发送成功  false 发送失败
end


----------------------------------------------------------------------
---接收消息
----------------------------------------------------------------------
-- 网络接收函数
function Network.Receive(serverId, userId, msgId, subMsgId, data)
    --Logger("lua收到了消息："..msgId)
    --Logger("lua收到了消息："..subMsgId)
    --Logger("lua收到了消息："..data)
    Network.Msg(serverId, userId, msgId, subMsgId, data)

end
-- 网络接收函数
function Network.UdpReceive(serverId, msgId, subMsgId, data)
    local serverAddress = serverId
    Network.Receive(serverAddress, 0, msgId, subMsgId, data)

end

function Network.Broken(uId)

end


-- 根据命令进行分支处理
function Network.Msg(serverId, userId, msgId, subMsgId, buf)
    --print("msgId", msgId, "subMsgId", subMsgId)

    if msgId == CMD_MAIN.MDM_GAME_SAND_ROCK then
        local switch={}
        switch[CMD_SAND_ROCK.SUB_LOGON] = LoginServer.Login
        switch[CMD_SAND_ROCK.SUB_LOGOUT] = LoginServer.Logout
        switch[CMD_SAND_ROCK.SUB_ROOM_INFO] = LoginServer.GameInfo
        switch[CMD_SAND_ROCK.SUB_ROOM_LIST] = LoginServer.RoomList
        switch[CMD_SAND_ROCK.SUB_OTHER_LOGON] = LoginServer.OtherLogin
        switch[CMD_SAND_ROCK.SUB_OTHER_LOGOUT] = LoginServer.OtherLogout
        switch[CMD_SAND_ROCK.SUB_OTHER_LOCATION] = Location.OtherLocation
        switch[CMD_SAND_ROCK.SUB_RESOURCE_POINT] = Sleep.UpdateResource
        switch[CMD_SAND_ROCK.SUB_RESOURCE_GET] = Resource.Get

        switch[subMsgId](serverId,userId, buf)

        --if subMsgId == CMD_CCC.SUB_LOGON then
        --    LoginServer.LoginGameServer(serverId, data)
        --elseif subMsgId == CMD_CCC.SUB_LOGOUT then
        --    LoginServer.Logout(data)
        --end
    end
end


function Network.Print(func,buf)
    local msg = func()
    msg:ParseFromString(buf)
    print(msg)
end