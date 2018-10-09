package BY

import "../Common"

type Bullet struct {
	Common.CommonBullet

}
func (bullet *Bullet) NewBullet(uid int)  Common.BulletInterface{
	return &Bullet{Common.CommonBullet{BulletUID:uid}}
}