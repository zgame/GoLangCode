---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/12/11 15:36
---

----------------------------保存服务器状态信息，是按照时间保存的，可以看历史记录-----------------------------
function SqlSaveServerState(state)
    local time= GetOsDateNow()
    local sql = string.format("insert into server_state (server_ip,time,table_num,player_num,rece_num,send_num, write_chan, head_err ) values ('%s','%s', %d,%d,%d,%d,%d,%d)",
            ServerIP_Port,time,state.TableNum,state.PlayerNum,state.ReceiveNum,state.SendNum,ServerSendWriteChannelNum,ServerDataHeadErrorNum)
    --print(sql)
    MysqlExec(sql)
end


----------------------------保存桌子状态信息，当前的运行信息，桌子销毁就删掉-----------------------------
function SqlSaveGameState(gameType,tableId, state)
    --RedisSaveString(RedisDirGameState..ServerIP_Port..":GameID_"..gameType..":TableId"..tableId, tableId, ZJson.encode(state))

    local select = string.format("select * from game_state where server_ip = '%s' and game_id = %d and table_id = %d",ServerIP_Port,gameType,tableId )
    local insert = string.format("insert into game_state (server_ip,game_id,table_id,fish_num,bullet_num,seat_array) values ('%s',%s, %d,%d,%d,%d)",ServerIP_Port,gameType,tableId ,state.FishNum,state.BulletNum,state.SeatArray)
    local update = string.format("update  game_state   set  fish_num =%d ,bullet_num =%d,seat_array=%d where server_ip = '%s' and game_id = %d and table_id = %d",state.FishNum,state.BulletNum,state.SeatArray,ServerIP_Port,gameType,tableId )
    --print(select)
    --print(insert)
    --print(update)

    local re = MysqlQuery(select)
    if #re == 0 then
        --print("insert")
        MysqlExec(insert)
    else
        --print("update")
        MysqlExec(update)
    end


    --insert into vclb_mm_inventory (`ID_`, `STOCK_ID_`, `ITEM_ID_`, `AMOUNT_`)
    --values ('489734716803514367', '仓库一', '水杯', 44)
    --ON DUPLICATE KEY UPDATE `AMOUNT_` = `AMOUNT_` + 44;

end
----------------------------删掉桌子状态信息-----------------------------
function SqlDelGameState(gameType,tableId)         -- 清理掉桌子的运行状态
    local sql = string.format("delete from game_state where server_ip = '%s' and game_id = %d and table_id = %d ",ServerIP_Port,gameType,tableId)
    --print(sql)
    MysqlExec(sql)
end
