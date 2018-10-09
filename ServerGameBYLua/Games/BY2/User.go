package BY2

//import "../BY"
import "../Common"
import "../../Const"

type User struct {
	Common.CommonUser
}

// -------------------------构造函数-------------------------
func (user *User) NewUser(uid int) Common.UserInterface{
	return &User{Common.CommonUser{UID:uid,TableID:Const.TABLE_CHAIR_NOBODY,ChairID:Const.TABLE_CHAIR_NOBODY}}
}

