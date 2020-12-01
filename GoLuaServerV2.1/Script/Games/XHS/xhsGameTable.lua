---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by soonyo.zhengxh
--- DateTime: 2019/10/22 18:03
---

--------------------------------------------------------------------------------------
--- 主要处理本游戏桌子主游戏逻辑
--------------------------------------------------------------------------------------

--- 发送玩家游戏场景信息
--- @param player 玩家对象
function XhsTable:SendUserGameInfo(player)
    if player == nil then
        Logger("XhsTable:SendUserGameInfo player 为ni了")
        return
    end
    -- 玩家User对象信息
    local user = player.User
    if user == nil then
        Logger("XhsTable:SendUserGameInfo user 为ni了")
        return
    end
    local MaxXhsHP = GetExcelXhsValue(user.MonsterID, "HP") or 0
    local SendMsg = CMD_XHS_Game_pb.CMD_User_SeaMonsterInfo_S()
    -- 海怪信息
    SendMsg.user_monster.monster_id     = user.MonsterID
    SendMsg.user_monster.monster_hp     = user.MonsterHP
    SendMsg.user_monster.bullet_num     = user.MonsterBulletNum
    SendMsg.user_monster.left_times     = GetIntervalBetweenNowAndTomorrow()
    SendMsg.user_monster.monster_max_hp = MaxXhsHP
    SendMsg.user_monster.summon_times   = user.SummonTimes
    -- 物品信息(发送玩家的召唤石和拼图碎片信息)
    for i, v in pairs(user.SkillInfoArray) do
        local itemType = GetExcelItemValue(v.ItemID,"item_type")
        if (itemType == Enum_ItemType.eItemType_Summon or itemType == Enum_ItemType.eItemType_JsGsaw_Chip) then
            local tbItem = SendMsg.user_items:add()
            tbItem.item_id  = v.ItemID
            tbItem.used     = v.Used
            tbItem.total    = v.Total
        end
    end
    -- 拼图
    for i , v in pairs(user.PuzzleIDArray) do
        SendMsg.user_puzzle:append(v)
    end
    -- 碎片信息
    SendMsg.puzzid_start    = GetTimeFromString("2018-01-01 00:00:00")
    SendMsg.puzzid_end      = GetTimeFromString("2028-01-01 00:00:00")
    LuaNetWorkSendToUser(player.User.UserID, MDM_GF_GAME, SUB_S_USER_DHS_GAME_INFO,SendMsg,nil, nil)
end