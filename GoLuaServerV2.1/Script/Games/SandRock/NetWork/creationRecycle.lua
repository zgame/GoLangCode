
SandRockCreationRecycleNet = {}

function SandRockCreationRecycleNet.CreateRecycle(serverId, userId, buf)
    local msg = ProtoGameSandRock.CreationRecycle()
    msg:ParseFromString(buf)
    --print(msg.createId)
    local createId = msg.createId
    local itemList = {}
    local player = GameServer.GetPlayerByUID(userId)
    local DropSceneList = CSV_creationRecycle.GetValue(createId, "DropSceneList")
    if DropSceneList == nil then
        ZLog.Logger("找不到CreateItemId" .. createId)
        return
    end

    local DropCountPercent = CSV_creationRecycle.GetValue(createId, "DropCountPercent")
    local RecycleTimeMinute = CSV_creationRecycle.GetValue(createId, "RecycleTimeMinute")
    local MachineLevel = CSV_creationRecycle.GetValue(createId, "MachineLevel")
    local OrderId = CSV_creationRecycle.GetValue(createId, "OrderId")
    local Exp = CSV_creationRecycle.GetValue(createId, "Exp")

    -- 判断机器的等级


    -- 判断机器的时间是否过期


    -- 先检查数量

    if Player.ItemHave(player, createId) < 1 then
        ZLog.Logger("道具数量不足" .. createId .. "   " .. 1)
        return
    end

    -- 再做处理
    itemList = SandRockGeneratorItem.GetItems(DropSceneList)        -- 道具生成组

    -- 道具减少
    Player.ItemReduce(player, createId, 1)

    -- 道具增加
    Player.ItemAdd(player, itemList)
    Player.ExpAdd(player, Exp)

    itemList[createId] = -1                     -- 减少的道具同步给客户端

    local sendCmd = SandRockSleepNet.SendItemList(player, itemList)
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_RECYCLE, sendCmd, nil)

end