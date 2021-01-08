
function Player:ExpGet()
    return self.user.exp
end
function Player:ExpAdd(exp)
    if exp > 0 then
        self.user.exp = self.user.exp + exp
    end
end
function Player:HpGet()
    return self.user.HP
end
function Player:HpAdd(hp)
    self.user.HP = self.user.HP + hp
end
function Player:SpGet()
    return self.user.SP
end
function Player:SpAdd(hp)
    if self.user.SP == nil then
        self.user.SP = 0
    end
    self.user.SP = self.user.SP + hp
end
