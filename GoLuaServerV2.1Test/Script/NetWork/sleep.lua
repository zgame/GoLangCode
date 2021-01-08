Sleep ={}


function Sleep.Sleep(serverId)
    print("申请睡觉")
    local sendCmd = ProtoGameSandRock.Sleep()
    Network.Send(serverId, CMD_MAIN.MDM_GAME_SAND_ROCK, CMD_SAND_ROCK.SUB_SLEEP,sendCmd,nil)
end


function Sleep.UpdateResource(serverId,userId, buf)
    print("资源刷新列表")
    local msg = ProtoGameSandRock.ResourceUpdate()
    msg:ParseFromString(buf)
    print(msg)

end

