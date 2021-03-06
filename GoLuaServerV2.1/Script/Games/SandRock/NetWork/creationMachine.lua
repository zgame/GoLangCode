SandRockCreationMachineNet = {}

-- 组装机器
function SandRockCreationMachineNet.CreateMachine(serverId, userId, buf)
    local msg = ProtoGameSandRock.CreationMachine()
    msg:ParseFromString(buf)
    --print(msg.createId)
    local createId = msg.createId
    local itemList = {}
    local player = GameServer.GetPlayerByUID(userId)
    local MachineItemId = CSV_creationMachine.GetValue(createId, "ItemId")
    if MachineItemId == nil then
        ZLog.Logger("找不到MachineItemId" .. createId)
        return
    end
    local PartIdList = CSV_creationMachine.GetValue(createId, "PartIds")
    local MachineLevel = CSV_creationMachine.GetValue(createId, "MachineLevel")

    -- 遍历所有的零件
    -- 先检查数量
    for _, partId in ipairs(PartIdList) do
        local Material = CSV_creationMachinePart.GetValue(partId, "Material")
        local ItemId = Material[1]
        local ItemNum = Material[2]
        if Player.ItemHave(player, ItemId) < ItemNum then
            ZLog.Logger("道具数量不足" .. ItemId .. "   " .. ItemNum)
            return
        end
    end
    -- 再做处理
    for _, partId in ipairs(PartIdList) do
        local Material = CSV_creationMachinePart.GetValue(partId, "Material")
        local ItemId = Material[1]
        local ItemNum = Material[2]
        Player.ItemReduce(player, ItemId, ItemNum)
        itemList[tostring(ItemId)] = -ItemNum
    end

    -- 发放机器道具
    local item ={}
    item[tostring(MachineItemId)] = 1
    Player.ItemAdd(player, item)
    itemList[tostring(MachineItemId)] = 1
    -- 玩家的机器处理

    local sendCmd = SandRockSleepNet.SendItemList(player, itemList)
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_MACHINE, sendCmd, nil)

end
