package Games

import (
	"../Core/ZServer"
	"fmt"
	. "../Const"
	"./BY"
	"./CommonLogin"
	//"../Core/Utils/zLog"
	//"./CommonLogin"
	"./Model/PlayerModel"
	"../Core/GameCore"

)

// 回调函数注册
func NetWorkFuncRegister()  {
	ZServer.NetWorkReceive = NetWorkReceive		// 绑定接收网络消息处理函数回调地址
	ZServer.NetWorkInit = NetWorkInit			// 绑定网络连接成功初始化
	ZServer.NetworkBroken = NetWorkBroken			// 网络中断回调
}


//---------------------------------------消息分发--------------------------------------

// 分发处理消息
func NetWorkReceive(serverId int,userId int, msgId int, subMsgId int, finalBuffer  []byte, token int) {
	var table GameCore.TableInterface
	var game  *GameCore.Games
	var player *PlayerModel.Player

	//fmt.Printf("server id   : %d    \n", serverId)
	//fmt.Printf("user id    : %d    \n", userId)
	if userId > 0 {
		player = GameCore.GetPlayerByUID(userId)
		//fmt.Printf("player 's game id    : %d    \n", player.GetGameID())
		game = GameCore.GetGameByID(player.GetGameID())
		//fmt.Printf(" player Id    :%d    gameid    :%d       \n" , player.GetUID(), game.GameID)
		if game != nil && player.GetTableID() > 0 {
			table = game.GetTableByUID(player.GetTableID())
			//fmt.Printf("  table  id      : %d    \n" ,  table.GetTableUID())
		}
	}

	//# -----------------login server msg-----------------
	if msgId == MDM_MB_LOGON {
		if subMsgId == SUB_MB_GUESTLOGIN {
			fmt.Println("**************游客  登录服   申请******************* ")
			//client.SevLoginGuest(finalBuffer)
		}
		//# -----------------login game server msg-----------------
	} else if msgId == MDM_GR_LOGON {
		if subMsgId == SUB_GR_LOGON_USERID {
			fmt.Println("**************游客  游戏服  申请******************* ")
			CommonLogin.HandleLoginGameServerGuest(serverId ,finalBuffer)
		}

		//# -----------------游戏场景 msg -----------------
	}else if msgId == MDM_GF_FRAME {
		if subMsgId == SUB_GF_GAME_OPTION {
			fmt.Println("**************进入游戏场景申请******************* ")
			HandelEnterScene(player,game,  finalBuffer)

		}
		//# -----------------场景内 msg------------------
	}else if msgId == MDM_GF_GAME {
		//if sub_msg_id == SUB_S_ENTER_SCENE {
		//	c.handleEnterScence(finalBuffer,int(bufferSize))
		//
		//	// 送一些金币
		//	//fmt.Println("发送gm命令，送金币")
		//
		//	c.SendGmCmd("@设置金币 10000000")
		//	//c.do_fire()
		//	c.StartAI = true
		//
		//
		//} else
		//if sub_msg_id == SUB_S_OTHER_ENTER_SCENE {
		//	client.handleOtherEnterScence(finalBuffer,int(bufferSize))			//进入场景,接收鱼数据
		//}else
		//if subMsgId == SUB_S_SCENE_FISH {
		//	client.handleSceneFish(finalBuffer,int(bufferSize))			//# 新生成鱼
		//}else if sub_msg_id == SUB_S_DISTRIBUTE_FISH {
		//	client.handleNewFish(finalBuffer, int(bufferSize)) //# 新生成鱼
		//}
		//------------------捕鱼----------------------
		BY.NetWorkReceive(serverId ,userId , msgId , subMsgId , finalBuffer  ,table, game ,player)

		//if subMsgId == SUB_C_USER_FIRE {
		//	BY.HandleUserFire(player,  BY.GetBYTableHandle(table) ,finalBuffer) // 客户端开火
		//}else if subMsgId == SUB_C_CATCH_FISH {
		//	BY.HandleCatchFish(player,  BY.GetBYTableHandle(table) ,finalBuffer) //	客户端抓鱼
		//}
	}
}

//  网络连接成功初始化
func NetWorkInit(serverId int)  {

}

//  网络中断回调
func NetWorkBroken(serverId int,userId int)  {

}

//
//// 发送网络消息
//func NetWorkSend(serverId int , uid int, data proto.Message,mainCmd int, subCmd int)  {
//	if uid > 0 {
//		ZServer.NetWorkSendByUid(uid, data, mainCmd, subCmd )
//	}else{
//		ZServer.NetWorkSendByServerId(serverId,data, mainCmd,subCmd)
//	}
//
//}
