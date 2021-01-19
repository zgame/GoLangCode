
--------------------------------等级和经验---------------------------------------
function Player:LevelGet()
    return self.user.level
end
function Player:ExpGet()
    return self.user.exp
end

function Player:ExpAdd(exp)
    if exp > 0 then
        self.user.exp = self.user.exp + exp
    end
    -- 是否升级的判断，后面加
end

--------------------------------hp sp cp------------------------------
function Player:HpGet()
    return self.user.HP
end
function Player:HpAdd(hp)
    self.user.HP = self.user.HP + hp
    if self.user.HP < 0 then
        self.user.HP = 0
    end
end
function Player:SpGet()
    return self.user.SP
end
function Player:SpAdd(hp)
    self.user.SP = self.user.SP + hp
    if self.user.SP < 0 then
        self.user.SP = 0
    end
end
