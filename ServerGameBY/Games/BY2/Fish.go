package BY2

import "../Common"

type Fish struct {
	Common.CommonFish

}

func (fish *Fish) NewFish(uid int)  Common.FishInterface{
	return &Fish{Common.CommonFish{FishUID:uid}}
}
