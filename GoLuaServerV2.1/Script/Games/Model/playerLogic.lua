
--------------------------------------------------------------------------------------
--- 这个文件主要是做player对象的成员函数处理
--------------------------------------------------------------------------------------
--
----- 增加玩家Score值
--function Player:AddScore(addScore)
--    if addScore < 0 then return end
--    self.User.Score = self.User.Score + addScore
--end
----------------------------------------------------------------------------------------
----- 扣除玩家Score值
--function Player:DecScore(decScore)
--    if decScore < 0 then return end
--    if self.User.Score <= decScore then
--        self.User.Score = 0
--    else
--        self.User.Score = self.User.Score - decScore
--    end
--end

function Player:UId()
    return self.user.userId
end

-- 把用户信息整理准备发送
function Player:Copy(sendCmdUser)
    sendCmdUser.userId = self.user.userId
    sendCmdUser.openId = self.user.openId
    sendCmdUser.nickName = self.user.nickName
    sendCmdUser.level =self.user.level
    sendCmdUser.exp = self.user.exp
    sendCmdUser.faceId = self.user.faceId
    sendCmdUser.gender = self.user.gender
    sendCmdUser.roomId = self.roomId
    sendCmdUser.chairId = self.chairId
end
