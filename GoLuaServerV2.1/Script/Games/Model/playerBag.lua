-- 背包


--------------------------------道具---------------------------------------
function Player:ItemHave(itemId)
    return 100
end

-- 获得道具的唯一id
function Player:GetSpeItemUId()
    self.user.itemUUId = self.user.itemUUId + 1
    return self.user.itemUUId
end
--
---- 获得特殊道具
--function Player:ItemSpeAdd(itemId)
--
--    local stack = SandRockItem.GetStack(itemId)
--    local slotSuccess = self:SlotEnough(itemId, stack, 1)
--    if slotSuccess then
--        -- 格子够的话，道具增加
--        self:SaveToPackage(itemId, stack, 1 , itemUId, nil)
--    else
--        -- 格子不够， 道具要临时保存一下
--        print("格子不够了， 道具已满")
--    end
--
--end



-- 获得道具
function Player:ItemAdd(itemList)
    --print("获得道具")
    --printTable(itemList)
    if itemList == nil then
        return
    end

    for itemId,itemNum in pairs(itemList) do
        local stack = SandRockItem.GetStack(itemId)
        local slotSuccess = self:SlotEnough(itemId, stack, itemNum)
        if slotSuccess then
            -- 格子够的话，道具增加
            self:SaveToPackage(itemId, stack, itemNum)
        else
            -- 格子不够， 道具要临时保存一下
            print("格子不够了， 道具已满")
        end
    end
end
-- 减少道具
function Player:ItemReduce(itemId,itemNum)
    if itemId == nil then
        return false
    end
    if Player:ItemHave(itemId) < itemNum  then
        return false
    end
end

-- 使用道具
function Player:ItemUse(item,itemNum)
    if item == nil then
        return
    end
end
--------------------------------------格子----------------------------------------------
-- 格子变化
function Player:SlotChange(change)
    if self.user.slotNow + change <= self.user.slotMax then
        -- 格子可以放下
        self.user.slotNow = self.user.slotNow + change
        return true
    else
        --格子放不下
        return false
    end
end

-- 格子的判断
function Player:SlotEnough(itemId, stack, itemNum)
    if  stack > 1 then
        -- 可堆叠, 有东西，并且没有满
        local numberNow = self:PackageItemNum(itemId)
        if numberNow > 0 then
            local mod = math.fmod(numberNow,stack)   -- 原来道具剩余堆叠数量
            if mod + itemNum <= stack then
                return  true                              -- 格子够用
            end
        end
    end

    -- 其他情况 ， 不可堆叠，没有道具， 有道具但是满堆叠了
    local slotSuccess = self:SlotChange(1)            -- 新开一个格子行不行

    return slotSuccess
end

--------------------------------------背包----------------------------------------------
-- 获得可以堆叠的道具的数量
function Player:PackageItemNum(itemId)
    if self.user.package == nil then
        self.user.package = {}
    end
    if self.user.package[tostring(itemId)] == nil then
    return 0
    end
    return self.user.package[tostring(itemId)]       -- 可堆叠的记录数量
    end

-- 获得不可以堆叠的道具的信息
function Player:PackageItemGet(itemId, itemUId)
    itemId = tostring(itemId)
    if self.user.package[itemId] == nil then
        return nil
    end
    return self.user.package[itemId][tostring(itemUId)]    -- 不可堆叠记录道具hash表
end

-- 保存到背包里面
function Player:SaveToPackage(itemId, stack, itemNum)
    itemId = tostring(itemId)
    if  stack > 1 then
        -- 可堆叠
        self.user.package[itemId] = self:PackageItemNum(itemId) + itemNum
    else
        if self.user.package[itemId] == nil then
            self.user.package[itemId] = {}
        end
        -- 不能堆叠需要item uid
        local itemUId = self:GetSpeItemUId()
        self.user.package[itemId][tostring(itemUId)] = {}
    end

    -- 保存到数据库
    SandRockUserDB.UserUpdate(self:UId(), self.user)
end