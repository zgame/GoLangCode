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


    switch[subMsgId](serverId, userId, buf)




end
