CreateItem ={}


function CreateItem.CreateItem(serverId,userId, buf)
    local msg = ProtoGameSandRock.ItemGet()
    msg:ParseFromString(buf)
    print(msg)
end

-- 创建道具
function CreateItem.SendCreationItem(serverId)
    local sendCmd = ProtoGameSandRock.CreationItem()
    sendCmd.createId = 40000408
    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_ITEM,sendCmd,nil)
end
-- 创建机器
function CreateItem.SendCreationMachine(serverId)
    local sendCmd = ProtoGameSandRock.CreationMachine()
    sendCmd.createId = 49000071
    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_MACHINE,sendCmd,nil)
end
-- 创建回收
function CreateItem.SendCreationRecycle(serverId)
    local sendCmd = ProtoGameSandRock.CreationRecycle()
    sendCmd.createId = 19340007
    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_RECYCLE,sendCmd,nil)
end
-- 创建烹饪
function CreateItem.SendCreationCooking(serverId)
    local sendCmd = ProtoGameSandRock.CreationCooking()
    sendCmd.createId = 1
    sendCmd.MaterialsID = 1
    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_CREATION_COOKING,sendCmd,nil)
end
