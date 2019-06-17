----------------------------------------------------------------------------------------------------------------
--- 跑马灯逻辑
----------------------------------------------------------------------------------------------------------------



--    {
-- [1] = "{
-- "id ": 0,
-- "server ": "22 ",
-- "tip ": "2222 ",
-- "start_time ": "2019 - 05 - 15 00:00:00 ",
-- "interval ": "222 ",
-- "loop ": "222 "
--}"
--    }

-- 检查跑马灯是否到时间了
function SystemCheckTip()
    --print("SystemCheckTip-----------------")
    local list = RedisGetSystemTip()        -- 获得跑马灯列表， 这里获得的是一个json
    --print("list          "..list..type(list))
    if list == "null" then
        return      -- redis里面没有数据
    end

    local tip_list = ZJson.decode(list)     -- 解开json，里面是一个列表
    for _, tip_json in pairs(tip_list) do
        local tip_table = ZJson.decode(tip_json)        -- 每一个元素是一个json字符串， 这里是redis保存的原始格式
        --print(tip_table.id)
        --print("跑马灯 "..tip_table.start_time)

        local tip_time = GetTimeFromString(tip_table.start_time)    -- 把字符串时间转换成时间戳
        local now = os.time()
        --print("now ",now)
        --print("tip_time ",tip_time)
        --print("day ", os.date("%Y-%m-%d",tip_time))

        if tip_time ~= nil then
            local time_diff =   now - tip_time
            --print("time_diff ",time_diff)
            if time_diff > 0 then
                -- 跑马灯时间到了
                if time_diff >  tonumber(tip_table.interval * tip_table.loop) then
                    -- 超过了最大的重复次数 ，那么删除这个记录
                    print("超过了最大重复次数，可以删除了")
                    -- print(tip_json)
                     RedisDelSystemTip(tip_json)
                else
                    -- 没超过，那么判断次数和间隔时间
                    if time_diff % tonumber(tip_table.interval) <= 5 then
                        print(" 到发送跑马灯时间了 "..GetOsDateNow())
                        -- 发送跑马灯给桌子上每个玩家
                        SystemTipSendToAllPlayers(tip_table.tip,0)

                    end
                end
            else
                -- 时间没到

            end
        end
    end
end

-- 发送固定字符串定义编号给所有客户端
function SystemTipLocalStringIndexSendToAllPlayers(local_string_index)

    -- 这里有bug， 不知道为啥
    --SystemTipSendToAllPlayers("", local_string_index )
end




-- 发送跑马灯消息给所有客户端
function SystemTipSendToAllPlayers(txt, type)
    --for _, game in pairs(AllGamesList) do
    --    for _, gameTable in pairs(game.AllTableList) do
    --
    --    end
    --end
    local sendCmd = CMD_GameServer_pb.CMD_CR_Tip()
    sendCmd.type = type
    sendCmd.level = 1
    sendCmd.delay_time = 1
    sendCmd.text = txt

    for _,player in pairs(AllPlayerList) do
        --print("player.User.UserId"..player.User.UserId)
        LuaNetWorkSendToUser(player.User.UserId, MDM_GF_GAME_TB, SUB_S_TIPS, sendCmd, nil, nil)
    end
end

