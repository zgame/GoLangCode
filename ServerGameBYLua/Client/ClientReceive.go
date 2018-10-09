package Client

// -------------------------------------------------------------------
// 接收和发送的业务逻辑
// 注意： 客户端发过来带token  ，  服务器发过去不要带token
// -------------------------------------------------------------------


import (
	"fmt"
	. "../Const"
	"../NetWork"
)

// 处理单个包内容
func (client *Client)handlerRead(buf []byte) int {
	//var err error
	//fmt.Printf("Receive buf: %x",buf)
	//fmt.Println(" ")
	msgId, subMsgId, bufferSize,ver := NetWork.DealRecvTcpDeaderData(buf)

	offset := 10

	//fmt.Println("len(buf)",len(buf))
	//fmt.Println("offset",offset)
	//fmt.Println("bufferSize",bufferSize)

	if len(buf) < offset + int(bufferSize){
		fmt.Println("出现数据包异常")
		return  int(bufferSize) + offset
	}

	if ver > 0{
		offset = 12		// version == 1 的时候， 加了一个token
	}
	finalBuffer := buf[offset:offset + int(bufferSize)]
	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	//fmt.Println("")

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
			client.SevLoginGSGuest(finalBuffer)

		}

		//# -----------------游戏场景 msg -----------------
	}else if msgId == MDM_GF_FRAME {
		if subMsgId == SUB_GF_GAME_OPTION {
			fmt.Println("**************游客进入大厅申请******************* ")
			client.SevEnterScence(buf)

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
			client.handleUserFire(finalBuffer) // 客户端开火
		}else if subMsgId == SUB_C_CATCH_FISH {
			client.handleCatchFish(finalBuffer) //	客户端抓鱼
		}

	}

	return int(bufferSize) + offset
}
