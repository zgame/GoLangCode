SandRockResourcePoint = {}


-- 同步资源列表
function SandRockResourcePoint.SendPointList(userId)
    local sendCmd = ProtoGameSandRock.ResourceUpdate()
    local room = GameServer.GetRoomByUserId(userId)
    if room == nil then
        return
    end

    --printTable(room.resourcePoint)
    for areaName,list in pairs(room.resourcePoint) do
            for ii,vv in ipairs(list) do
                local points = sendCmd.points:add()
                points.areaName = areaName
                points.areaPoint = vv.areaPoint
                points.resourceType = vv.resourceType
            end
    end
    --print(sendCmd)
    NetWork.SendToUser(userId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_RESOURCE_POINT, sendCmd, nil)


end

-- 采集资源
function SandRockResourcePoint.GetResource(serverId, userId, buf)
    local msg = ProtoGameSandRock.PlayerLocation()
    msg:ParseFromString(buf)
    print(msg)

end
