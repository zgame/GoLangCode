
SandRockCreationCookingNet = {}

function SandRockCreationCookingNet.CreateCooking(serverId, userId, buf)
    local msg = ProtoGameSandRock.CreationCooking()
    msg:ParseFromString(buf)
    print(msg.createId)
    local createId = msg.createId
    local MaterialsID = msg.MaterialsID
    local itemList = {}
    local player = GameServer.GetPlayerByUID(userId)
    local CookingOutId = CSV_creationCooking.GetValue(createId, "CookingOutId")
    if CookingOutId == nil then
        ZLog.Logger("找不到CookingOutId" .. createId)
        return
    end
    local Materials = CSV_creationCookingPart.GetValue(MaterialsID, "Materials")
    if Materials == nil then
        ZLog.Logger("找不到MaterialsId" .. MaterialsID)
        return
    end

    local CostMinutes = CSV_creationCooking.GetValue(createId, "CostMinutes")
    local MaterialsIDList = CSV_creationCooking.GetValue(createId, "MaterialsID")
    local IsUsed = CSV_creationCooking.GetValue(createId, "IsUsed")


    local CookingType = CSV_creationCookingPart.GetValue(MaterialsID, "CookingType")
    local MaterialsList = CSV_creationCookingPart.GetValue(MaterialsID, "Materials")



    -- 先检查数量， 这里要小心重复的材料
    for _,itemId in ipairs(MaterialsList) do
        if Player.ItemHave(player, itemId) < 1 then
            ZLog.Logger("道具数量不足" .. itemId .. "   " .. 1)
            return
        end
    end

    -- 道具减少
    for _,itemId in ipairs(MaterialsList) do
        Player.ItemReduce(player, itemId, 1)
        itemList[itemId] = -1                     -- 减少的道具同步给客户端
    end

    -- 道具增加
    Player.ItemAdd(player, {CookingOutId = 1})
    --Player.ExpAdd(player, Exp)
    itemList[CookingOutId] = 1


    local sendCmd = SandRockSleepNet.SendItemList(player, itemList)
    NetWork.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_COOKING, sendCmd, nil)

end