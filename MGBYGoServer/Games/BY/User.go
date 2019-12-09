package BY

import "../Common"
import "../../Const"


type User struct {
	Common.CommonPlayer
}

// -------------------------构造函数-------------------------
func (user *User) NewPlayer(uid int) Common.PlayerInterface {
	return &User{Common.CommonPlayer{UID:uid,TableID:Const.TABLE_CHAIR_NOBODY,ChairID:Const.TABLE_CHAIR_NOBODY}}
}
