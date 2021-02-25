Sleep ={}


function Sleep.Sleep(serverId)
    print("申请睡觉")
    local sendCmd = ProtoGameSandRock.Sleep()
    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_SLEEP,sendCmd,nil)
end


function Sleep.UpdateResourcePick(serverId, userId, buf)
    --print("拾取资源刷新列表")
    local msg = ProtoGameSandRock.ResourceUpdate()
    msg:ParseFromString(buf)
    --print(msg)
    --print(msg.points[1].areaName)
    --print(msg.points[1].areaPoint)
    --print(msg.points[1].resourceType)

    -- 采集资源
    --ResourcePick.Action(serverId,msg.points[1].areaName,msg.points[1].areaPoint,msg.points[1].resourceType )
end


function Sleep.UpdateResourceTerrain(serverId, userId, buf)
    --print("地形资源刷新列表")
    local msg = ProtoGameSandRock.ResourceTerrainUpdate()
    msg:ParseFromString(buf)
    --print(msg)
    --print(msg.points[1].areaName)
    --print(msg.points[1].areaPoint)
    --print(msg.points[1].resourceType)

 end

