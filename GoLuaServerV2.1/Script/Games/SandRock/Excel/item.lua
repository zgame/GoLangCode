SandRockItem = {}

--  把几率计算成百分比
local function _setItem(itemId)
    local list = CSV_item.GetValue(itemId,"QualitySceneSlot")
    SandRockItem[itemId] = {}
    local all = 0
    for i,v in ipairs(list)do
        local rate = v + all
        all = all + v
        table.insert(SandRockItem[itemId],rate)
    end
end
-- 把几率计算好
function SandRockItem.Init()
    for itemId,_ in pairs(CSV_item.Get()) do
        _setItem(itemId)
    end
end

-- 获取道具的品质
function SandRockItem.GetQuality(itemId)
    local quality = ZRandom.GetList(SandRockItem[tostring(itemId)])
    return quality
end

-- 获取道具的堆叠
function SandRockItem.GetStack(itemId)
    local stack = CSV_item.GetValue(tostring(itemId),"StackNumber")
    return stack
end

-- 获取道具的购买价格
function SandRockItem.GetBuyPrice(itemId)
    local buy = CSV_item.GetValue(tostring(itemId),"BuyPrice")
    return buy
end

-- 获取道具的卖出价格
function SandRockItem.GetSellPrice(itemId)
    local sell = CSV_item.GetValue(tostring(itemId),"SellPrice")
    local sellChange = CSV_item.GetValue(tostring(itemId),"SellPriceChange")     -- 卖出道具价格浮动 ？
    return sell
end

-- 获取道具能否出售
function SandRockItem.CanSell(itemId)
    local ok = CSV_item.GetValue(tostring(itemId),"CantSold")
    return ok
end
-- 获取道具能否丢弃
function SandRockItem.CanDrop(itemId)
    local ok = CSV_item.GetValue(tostring(itemId),"CantDiscard")
    return ok
end