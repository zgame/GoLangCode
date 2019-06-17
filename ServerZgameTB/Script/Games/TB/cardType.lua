
-------------------------------------------------------------------------------
--- 生成翻牌的类型
-------------------------------------------------------------------------------


--随机没中奖的
function RandFail(sendCmd)

    local l_num = {}        -- 数字的数组
    table.insert( l_num, GetRandom(2,4))
    table.insert( l_num, GetRandom(5,6))
    table.insert( l_num,GetRandom(7,9))
    table.insert( l_num, GetRandom(10,12))
    table.insert( l_num, GetRandom(13,14))
    local random_num = ZRandomShuffle(l_num)        -- 将数组顺序打乱

    local l_type = GetRandom(1,4)
    for i=0,4 do
        local pCard = sendCmd.item_array:add()
        pCard.ctype = (l_type + i + GetRandom(1,3))%4 + 1
        pCard.cnum = random_num[i+1]
    end
    --print("没中奖")
    return sendCmd
end


--随机Wang
function RandWang(sendCmd)
    sendCmd = RandFail(sendCmd)
    local index = GetRandom(0, 3)
    sendCmd.item_array[index+1].cnum = 998      -- 插入一张王
    --print("王")
    return sendCmd
end

--随机同花顺
function RandTongHuaShunZi(sendCmd)
    local l_num = GetRandom( 2, 10)
    local l_type = GetRandom( 1, 4)

    for i=0,4 do
        local pCard = sendCmd.item_array:add()
        pCard.ctype = l_type        -- 花色都一样
        pCard.cnum = l_num + i      -- 数字按照顺序来
    end
    --print("同花顺")
    return sendCmd
end

--随机铁支
function RandCardTieZhi(sendCmd)
    local l_num = GetRandom( 2,14)
    local l_type = GetRandom( 1,4)

    for i=0,4 do
        local pCard = sendCmd.item_array:add()
        pCard.ctype = (l_type + i)%4 + 1        -- 不同的花色
        pCard.cnum = l_num              -- 一样的数字
    end

    local l_num2 = GetRandom(3,14)
    if l_num2 == l_num then
        -- 如果随机相等了，那么让它不相等
        l_num2 = l_num2 - 1
    end

    --下面将第五张牌变化一下
    sendCmd.item_array[5].cnum = l_num2
    sendCmd.item_array[5].ctype = (l_type + GetRandom( 1,4))  % 4 + 1
    --print("铁支")
    return sendCmd

end

--随机葫芦
function RandCardHuLu(sendCmd)
    local l_num = GetRandom( 2,14)
    local l_type = GetRandom( 1,4)

    for i=0,4 do
        local pCard = sendCmd.item_array:add()
        pCard.ctype = (l_type + i)%4 + 1
        pCard.cnum = l_num          -- 全部都一样
    end

    local l_num2 = GetRandom(3,14)
    if l_num2 == l_num then
        -- 如果随机相等了，那么让它不相等
        l_num2 = l_num2 - 1
    end
    sendCmd.item_array[4].cnum = l_num2     -- 让两张是另外的数字
    sendCmd.item_array[5].cnum = l_num2
    --print("葫芦")
    return sendCmd

end

--随机同花
function RandTongHua(sendCmd)

    local l_type = GetRandom(1,4)
    for i=0,4 do
        local pCard = sendCmd.item_array:add()
        pCard.ctype = l_type    -- 花色都一样
    end
    local l_num = GetRandom( 7,9)
    local l_num_list = {}
    local temp = {}     -- 临时数组
    temp[1] = GetRandom(1,3)
    temp[2] = GetRandom(1,2)
    temp[3] = GetRandom(1,3)
    temp[4] = GetRandom(1,2)
    -- 如果其他的都是连续的，那么让一个不连续， 防止成为顺子
    if temp[1] == 1 and temp[2] == 1 and temp[3] == 1 and temp[4] == 1 then
        temp[GetRandom(1,4)] = 2
    end
    -- 开始把数字加入到牌中
    table.insert(l_num_list, l_num)
    table.insert(l_num_list, l_num - temp[1])
    table.insert(l_num_list, l_num - temp[1] - temp[2])
    table.insert(l_num_list, l_num + temp[3])
    table.insert(l_num_list, l_num + temp[3] + temp[4])

    local random_num = ZRandomShuffle(l_num_list)        -- 将数组顺序打乱
    for i=1,5 do
        sendCmd.item_array[i].cnum = random_num[i]
    end
    --print("同花")
    return sendCmd

end

--随机顺子
function RandCardShunZi(sendCmd)
    local l_num = GetRandom( 2,10)
    local l_type = GetRandom( 1,4)
    for i=0,4 do
        local pCard = sendCmd.item_array:add()
        pCard.ctype = (l_type + i + GetRandom(1,3))%4 + 1
        pCard.cnum = l_num + i
    end
    --print("顺子")
    return sendCmd
end

--随机三条
function RandCardSanTiao(sendCmd)
    RandFail(sendCmd)
    -- 找到第一张牌
    local num = sendCmd.item_array[1].cnum
    local type = sendCmd.item_array[1].ctype
    -- 把第二，第三张牌设置成一样的数字
    sendCmd.item_array[2].cnum = num
    sendCmd.item_array[3].cnum = num
    -- 然后花色设置成不同的即可
    sendCmd.item_array[2].ctype = (type+1)%4+1
    sendCmd.item_array[3].ctype = (type+2)%4+1

    --print("三条")
    return sendCmd

end

--随机双对子
function RandCardDoubleDuiZi(sendCmd)
    RandFail(sendCmd)
    -- 找到第一张牌
    local num = sendCmd.item_array[1].cnum
    local type = sendCmd.item_array[1].ctype
    -- 找到第三张牌
    local num3 = sendCmd.item_array[3].cnum
    local type3 = sendCmd.item_array[3].ctype
    -- 把第二张牌设置成一样的数字， 不一样的花色
    sendCmd.item_array[2].cnum = num
    sendCmd.item_array[2].ctype = (type+1)%4+1
    -- 把第四张牌设置成一样的数字， 不一样的花色
    sendCmd.item_array[4].cnum = num3
    sendCmd.item_array[4].ctype = (type3+1)%4+1
    --print("双对子")
    return sendCmd


end
--随机对子
function RandCardDuiZi(sendCmd)
    RandFail(sendCmd)
    -- 找到第一张牌
    local num = sendCmd.item_array[1].cnum
    local type = sendCmd.item_array[1].ctype
    -- 把第二张牌设置成一样的数字， 不一样的花色
    sendCmd.item_array[2].cnum = num
    sendCmd.item_array[2].ctype = (type+1)%4+1
    --print("对子")
    return sendCmd
end

