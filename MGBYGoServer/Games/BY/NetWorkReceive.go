package BY

import (

	. "../../Const"
	"../Model/PlayerModel"
	"../../Core/GameCore"

)

func NetWorkReceive(serverId int,userId int, msgId int, subMsgId int, finalBuffer  []byte,table GameCore.TableInterface,game  *GameCore.Games ,player *PlayerModel.Player)  {
	byTable := GetBYTableHandle(table)			// 转换成 by table

	if subMsgId == SUB_C_USER_FIRE {
		HandleUserFire(player, byTable  ,finalBuffer) // 客户端开火
	}else if subMsgId == SUB_C_CATCH_FISH {
		HandleCatchFish(player,  byTable ,finalBuffer) //	客户端抓鱼
	}
}