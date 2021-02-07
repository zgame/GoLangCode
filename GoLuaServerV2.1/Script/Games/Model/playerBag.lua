-- 背包


--------------------------------道具---------------------------------------
function Player:ItemHave(itemId)
    return 100
end

-- 获得道具
function Player:ItemAdd(itemList)
    if itemList == nil then
        return
    end
    for itemId,itemNum in pairs(itemList) do

    end
end
-- 减少道具
function Player:ItemReduce(itemId,itemNum)
    if itemId == nil then
        return false
    end
    if Player:ItemHave(itemId)< itemNum  then
        return false
    end


end

-- 使用道具
function Player:ItemUse(item,itemNum)
    if item == nil then
        return
    end
end