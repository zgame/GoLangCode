
-------------------------------------------------------------------------------
--- 翻牌部分逻辑， 只掉小金币， 从侧面出
-------------------------------------------------------------------------------



-- 获取翻牌几率
 function GetCardWinPro(player)
    --local gameTable = player:GetTable()
     -- 中奖初始几率
     local bloodWin = CardBloodModeOpenRate    --  轮盘中奖率 0.03,
     local cardWin = CardBingoRate  --    其他中奖率，0.09 % ，  不中奖率，0.88%

     -- 玩家的校验
     if  player.m_mapUserPokerTriggerTimes >= 30 then
         player:AddConfineUser()
         print("强制不中翻牌"..player.User.UserId)
         return nil
     end


     -- 库存的影响
     local fAdjustAll = PoolAllRate(player)

     -- 个人库存的影响
     local fAdjustPerson = PoolPersonRate(player)


     ---- 黑名单改变几率
     --local cbConfineLevel = player.User.ConfineLevel
     --if cbConfineLevel > 0 then
     --    local tiger_adjust =  GetExcelValue(TBConfineExcel, gameTable.RoomScore  , "tiger_adjust")
     --    fAdjust = fAdjust * tiger_adjust / 10000
     --end

     bloodWin = bloodWin * fAdjustAll * fAdjustPerson / 10000  / 10000
     --print("热血模式的中奖几率为：", bloodWin)
     cardWin = cardWin * fAdjustAll * fAdjustPerson / 10000  / 10000
     --print("翻牌的中奖几率为：", cardWin)

     local rand = GetRandom(1,10000)

     --if 50 < GetRandom(0,100) then
     --    rand = bloodWin --调试
     --end


     if rand <= bloodWin then
         return "BloodBingo"
     elseif rand <= cardWin then
         return  "CardBingo"
     end

     return  nil   -- 没有中奖

end



--计算中的那种类型的牌型
 function CalcCardWinPro(player)
    local gameTable = player:GetTable()

    local rand = GetRandom(1,10000)
     for cardType =1,8 do
         local rate = GetExcelValue(TBCardExcel, cardType, "rate")
         if rand <= rate then
             -- 就是这种类型了
             local multi = GetExcelValue(TBCardExcel, cardType, "multi")
             return cardType,multi
         end
     end

     print("没有找到翻牌的类型.................")

end
