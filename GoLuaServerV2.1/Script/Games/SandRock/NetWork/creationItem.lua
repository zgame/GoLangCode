
SandRockCreationItemNet ={}

-- 道具合成
function SandRockCreationItemNet.CreateItem(serverId, userId, buf)
    local msg = ProtoGameSandRock.CreationItem()
    msg:ParseFromString(buf)
    --print(msg)
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
    --switch[FromMachineType]()

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
        itemList[tostring(ItemId)] = -ItemNum
    end

    local item ={}
    item[tostring(CreateItemId)] = CreateItemNum
    Player.ItemAdd(player, item)
    Player.ExpAdd(player, Exp)

    itemList[tostring(CreateItemId)] = CreateItemNum

    local sendCmd = SandRockSleepNet.SendItemList(player, itemList)
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_ITEM, sendCmd, nil)

end