package Games

import (
	"../Core/ZServer"
	"fmt"
	. "../Const"
	"./BY"
	//"../Core/Utils/zLog"
	"./Common"
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
	var table Common.TableInterface
	var game  *Games
	var player *Common.Player
	if userId > 0 {
		player = GetPlayerByUID(userId)
		game = GetGameByID(player.GetGameID())
		if game != nil {
			table = game.GetTableByUID(player.GetTableID())
		}
	}
	//# -----------------login server msg-----------------
	if msgId == MDM_MB_LOGON {
		if subMsgId == SUB_MB_GUESTLOGIN {
			fmt.Println("**************游客登录服申请******************* ")
			//client.SevLoginGuest(finalBuffer)
		}
		//# -----------------login game server msg-----------------
	} else if msgId == MDM_GR_LOGON {
		if subMsgId == SUB_GR_LOGON_USERID {
			fmt.Println("**************游客登录游戏服申请******************* ")
			BY.HandleLoginGameServerGuest(finalBuffer)
		}

		//# -----------------游戏场景 msg -----------------
	}else if msgId == MDM_GF_FRAME {
		if subMsgId == SUB_GF_GAME_OPTION {
			fmt.Println("**************游客进入大厅申请******************* ")
			BY.HandelEnterScence(finalBuffer)

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

		if subMsgId == SUB_C_USER_FIRE {
			BY.HandleUserFire(player, game, table ,finalBuffer) // 客户端开火
		}else if subMsgId == SUB_C_CATCH_FISH {
			BY.HandleCatchFish(player, game, table ,finalBuffer) //	客户端抓鱼
		}
	}
}

//  网络连接成功初始化
func NetWorkInit(serverId int)  {

}

//  网络中断回调
func NetWorkBroken(serverId int,userId int)  {

}
