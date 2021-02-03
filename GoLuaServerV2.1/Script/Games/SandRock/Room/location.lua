
-------------------------位置---------------------------------
function SandRockRoom:LocationPlayerSet(uId, msg)
    if uId==nil and msg ==nil then      -- 如果都是空的， 那么就清空
        self.locationList ={}
        return
    end
    self.locationList[tostring(uId)] = msg  -- 不是空的就添加
end
function SandRockRoom:GetPlayerLocation(uId)
    if uId == nil then                      -- 输入空，返回所有
        return self.locationList
    else
        return self.locationList[tostring(uId)]     -- 不空返回单条
    end
end

-- 同步其他玩家位置和状态
-- 这个地方为了节省cpu和内存，我就统一形成一次发送数据， 每个玩家都一样的发送，不然我要针对每个玩家单独处理数据，要费一些
function SandRockRoom:LocationOther()
    if self.userSeatArrayNumber == 1 then  -- 如果房间只有一个人，那么就不用同步了
        return
    end

    --print("************************同步所有玩家位置*****************")
    local sendCmd = ProtoGameSandRock.PlayerLocation()
    local lens = 0
    for _, value in pairs(self.locationList)do
        local location = sendCmd.location:add()
        location = SandRockLocationNet.Copy(value, location)
        lens = lens + 1
    end
    if lens == 0 then
        return  --没有消息就不发
    end
    sendCmd.time = 22
    --print("------------------------------------------同步位置和动作------------------------------".. os.time())
    --print(sendCmd)

    self:SendMsgToAllUsers(CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_OTHER_LOCATION, sendCmd)
    self:LocationPlayerSet(nil,nil)      -- 清空
end
