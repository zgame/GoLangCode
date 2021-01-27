---
---  SandRock 的网络分发
---

SandRockNetwork = {}

function SandRockNetwork.Receive(serverId, userId, mainSgId, subMsgId, buf, token)
    local switch={}
    switch[CMD_SAND_ROCK.SUB_LOGON] = SandRockLoginNet.Login
    switch[CMD_SAND_ROCK.SUB_LOGOUT] = SandRockLoginNet.Logout
    switch[CMD_SAND_ROCK.SUB_LOCATION] = SandRockLocationNet.Location
    switch[CMD_SAND_ROCK.SUB_SLEEP] = SandRockSleepNet.Sleep
    switch[CMD_SAND_ROCK.SUB_RESOURCE_GET] = SandRockResourcePickNet.GetPickResource
    switch[CMD_SAND_ROCK.SUB_RESOURCE_TERRAIN_GET] = SandRockResourceTerrainNet.GetTerrainResource
    switch[CMD_SAND_ROCK.SUB_CREATION_MACHINE] = SandRockCreationMachineNet.CreateMachine
    switch[CMD_SAND_ROCK.SUB_CREATION_ITEM] = SandRockCreationItemNet.CreateItem
    switch[CMD_SAND_ROCK.SUB_CREATION_RECYCLE] = SandRockCreationRecycleNet.CreateRecycle
    switch[CMD_SAND_ROCK.SUB_CREATION_COOKING] = SandRockCreationCookingNet.CreateCooking


    --print("消息 ".. subMsgId.. " start ： "  .. ZTime.GetOsTimeMillisecond())
    local startTime = ZTime.GetOsTimeMillisecond()
    switch[subMsgId](serverId, userId, buf)
    --print("消息 ".. subMsgId.. "  end  ： "  .. ZTime.GetOsTimeMillisecond())
    local endTime = ZTime.GetOsTimeMillisecond()
    if endTime - startTime > 100 then
        print("消息 ".. subMsgId.. "  cost time  ： "  .. endTime - startTime)
    end



end
