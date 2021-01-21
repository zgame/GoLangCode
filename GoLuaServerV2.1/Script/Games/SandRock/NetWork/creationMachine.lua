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
    for i, partId in ipairs(PartIdList) do
        local Material = CSV_creationMachinePart.GetValue(partId, "Material")
        local ItemId = Material[1]
        local ItemNum = Material[2]
        if Player.ItemHave(player, ItemId) < ItemNum then
            ZLog.Logger("道具数量不足" .. ItemId .. "   " .. ItemNum)
            return
        end
    end
    -- 再做处理
    for i, partId in ipairs(PartIdList) do
        local Material = CSV_creationMachinePart.GetValue(partId, "Material")
        local ItemId = Material[1]
        local ItemNum = Material[2]
        Player.ItemReduce(player, ItemId, ItemNum)
        itemList[ItemId] = -ItemNum
    end

    -- 发放机器道具
    Player.ItemAdd(player, { MachineItemId = 1 })
    itemList[MachineItemId] = 1
    -- 玩家的机器处理

    local sendCmd = SandRockSleepNet.SendItemList(player, itemList)
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_MACHINE, sendCmd, nil)

end

-- 道具合成
function SandRockCreationMachineNet.CreateItem(serverId, userId, buf)
    local msg = ProtoGameSandRock.CreationItem()
    msg:ParseFromString(buf)
    --print(msg.createId)
    local createId = msg.createId
    local itemList = {}
    local player = GameServer.GetPlayerByUID(userId)
    local CreateItemId = CSV_creationItem.GetValue(createId, "ItemId")
    if CreateItemId == nil then
        ZLog.Logger("找不到CreateItemId" .. createId)
        return
    end

    local FromMachineType = CSV_creationItem.GetValue(createId, "FromMachineType")
    local CreateItemNum = CSV_creationItem.GetValue(createId, "ItemNum")
    local Materials = CSV_creationItem.GetValue(createId, "Materials")
    local MachineLevel = CSV_creationItem.GetValue(createId, "MachineLevel")
    local SyntheticsTime = CSV_creationItem.GetValue(createId, "SyntheticsTime")
    local OrderId = CSV_creationItem.GetValue(createId, "OrderId")
    local Exp = CSV_creationItem.GetValue(createId, "Exp")

    -- 根据机器的类型进行判断
    local switch ={}
    switch["Cutting"] = {}                      -- 切割
    switch["DryingStand"] = {}                  -- 干燥台
    switch["Furnace"] = {}                      -- 熔炉 加工石头
    switch["Jewelry"] = {}                      -- 珠宝
    switch["Stirrer"] = {}                      -- 搅拌机
    switch["SyntheticTable"] = {}               -- 合成台
    switch["Tailor"] = {}                       -- 裁缝
    switch["Equipments"] = {}                   -- 装备
    switch["Weapons"] = {}                      -- 武器
    switch[FromMachineType]()

    -- 判断机器的等级


    -- 判断机器的时间是否过期


    -- 先检查数量
    for i, part in ipairs(Materials) do
        local ItemId = part[1]
        local ItemNum = part[2]
        if Player.ItemHave(player, ItemId) < ItemNum then
            ZLog.Logger("道具数量不足" .. ItemId .. "   " .. ItemNum)
            return
        end
    end
    -- 再做处理
    for i, part in ipairs(Materials) do
        local ItemId = part[1]
        local ItemNum = part[2]
        Player.ItemReduce(player, ItemId, ItemNum)
        itemList[ItemId] = -ItemNum
    end
    Player.ItemAdd(player, { CreateItemId = CreateItemNum })
    Player.ExpAdd(player, Exp)

    itemList[CreateItemId] = CreateItemNum

    local sendCmd = SandRockSleepNet.SendItemList(player, itemList)
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_ITEM, sendCmd, nil)

end