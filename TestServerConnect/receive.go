package main

import (
	"fmt"
	. "./const"
	"strconv"
)

// 处理单个包内容
func (this *Client)handlerRead(buf []byte) int {
	//var err error
	offset := 10

	if len(buf)< offset {		// 接受不全，那么缓存
		//log.PrintfLogger("数据包头部小于 10 : %x   ",buf)

		this.ReceiveBuf = buf
		//str:= fmt.Sprintf("%d数据包头部小于 10 : %x   ",this.Index,buf)
		//this.PrintLogger(str)
		return 0
	}
	// 读取头部信息
	msg_id, sub_msg_id, bufferSize, ver := dealRecvTcpDeaderData(buf)

	//BufAllSize := offset + int(bufferSize)    // 整个数据包长度

	if len(buf) < offset + int(bufferSize) {	// 接受不全，那么缓存
		this.ReceiveBuf = buf

		return  0
	}
	if ver > 0{
		offset = 12		// version == 1 的时候， 加了一个token
	}
	finalBuffer := buf[offset:offset + int(bufferSize)]
	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	//fmt.Println("")

	//# -----------------login server msg-----------------
	if msg_id == MDM_MB_LOGON {
		if sub_msg_id == SUB_MB_LOGON_SUCCESS {
			this.handleLoginSucess(finalBuffer,int(bufferSize))
		} else if sub_msg_id == SUB_MB_LOGON_FAILURE {
			this.handleLoginFailed(finalBuffer,int(bufferSize))
		} else if sub_msg_id == SUB_MB_LOGON_FINISH {
			//fmt.Println("登录成功 ------ ", this.Index)



		}
	} else if msg_id == MDM_MB_GIFT_PACK {
		if sub_msg_id == SUB_MB_L2C_GIFT_PRODUCT_INFO {
			//fmt.Println("礼包 ------")
		}
	} else if msg_id == MDM_MB_ACTIVITY {
		if sub_msg_id == SUB_MB_S2C_ACTIVITY {
			//fmt.Println("活动 ------")
		} else if sub_msg_id == SUB_MB_S2C_ACTIVITY_CELL_INFO_LIST {
			//fmt.Println("活动 ------")
		}
	} else if msg_id == MAIN_CHAT_CMD {
		if sub_msg_id == SUB_S_LOGIN {
			this.handleLoginCS(finalBuffer,int(bufferSize))

		}
	} else if msg_id == MDM_MB_VIP {
		if sub_msg_id == SUB_MB_S_VIP_INFO {
			//fmt.Println("vip ------")
		}
	} else if msg_id == MDM_MB_SERVER_LIST {
		if sub_msg_id == SUB_MB_LIST_SERVER {
			this.handleServerList(finalBuffer,int(bufferSize))

			//获取完游戏服务器列表，开始登录游戏服务器
			//fmt.Println("开始登录游戏服务器",this.Serverlist[0].server_addr,":", strconv.Itoa(this.Serverlist[0].server_port))
			addr := this.Serverlist[0].server_addr + ":" + strconv.Itoa(this.Serverlist[0].server_port)
			this.ConnectGameServer(addr)


		} else if sub_msg_id == SUB_MB_LIST_FINISH {
			//fmt.Println("房间信息完成 ------")
		}
	} else if msg_id == MDM_MB_USER_INFO {
		if sub_msg_id == SUB_MB_S_GET_CHAT_SERVER_INFO {
			this.handleCsInfo(finalBuffer,int(bufferSize))
		} else if sub_msg_id == SUB_MB_S_USER_MATERIAL_OBJECT {
			//fmt.Println("玩家信息完成 ------")
		} else if sub_msg_id == SUB_MB_S_REQUEST_ARENA {
			//fmt.Println("竞技场数据 ------")




		}

		//# -----------------login game server msg-----------------
	} else if msg_id == MDM_GR_LOGON {
		if sub_msg_id == SUB_GR_LOGON_FAILURE {
			this.handleLoginFailedGs(finalBuffer,int(bufferSize))
		} else if sub_msg_id == SUB_GR_LOGON_SUCCESS {
			this.handleLoginSucessGs(finalBuffer,int(bufferSize))
		} else if sub_msg_id == SUB_GR_LOGON_FINISH {
			fmt.Println("gs 登录完成 ------")


			//  开始进入场景
			//fmt.Println("开始发送进入场景")
			this.EnterScence()

		}

	}else if msg_id == MDM_GR_USER {
		if sub_msg_id == SUB_GR_S_USER_STANDUP {
			this.handleLoginOutGs(finalBuffer,int(bufferSize))
			//fmt.Println("玩家离开桌子 ------")
		}
		//# -----------------游戏场景 msg -----------------
	}else if msg_id == MDM_GF_FRAME {
		if sub_msg_id == SUB_GF_GAME_STATUS {
			this.handleGameStatus(buf,int(bufferSize))
		} else if sub_msg_id == SUB_GF_SYSTEM_MESSAGE {
			this.handleGameMessage(buf,int(bufferSize))
		}else if sub_msg_id == SUB_GF_USER_SKILL {
			this.handleUserSkill(buf,int(bufferSize))
		}
		//# -----------------场景内 msg------------------
	}else if msg_id == MDM_GF_GAME {
		if sub_msg_id == SUB_S_ENTER_SCENE {
			this.handleEnterScence(finalBuffer,int(bufferSize))

			// 送一些金币
			//fmt.Println("发送gm命令，送金币")
			this.SendGmCmd("@设置金币 10000000")

			//GMAddSkill( this.User.user_id , SkillID)
			//this.do_fire()
			this.StartAI = true


		} else if sub_msg_id == SUB_S_OTHER_ENTER_SCENE {
			this.handleOtherEnterScence(finalBuffer,int(bufferSize))			//进入场景,接收鱼数据
		}else if sub_msg_id == SUB_S_SCENE_FISH {
			this.handleSceneFish(finalBuffer,int(bufferSize))			//# 新生成鱼
		}else if sub_msg_id == SUB_S_DISTRIBUTE_FISH {
			this.handleNewFish(finalBuffer,int(bufferSize))			//# 新生成鱼

			//------------------捕鱼----------------------
		}else if sub_msg_id == SUB_S_USER_FIRE {
			this.handleUserFire(finalBuffer,int(bufferSize))			// 开火
		}else if sub_msg_id == SUB_S_USE_SKILL {
			this.handleUserUseSkill(finalBuffer,int(bufferSize))			// 技能
		}else if sub_msg_id == SUB_S_CATCH_FISH {
			this.handleCatchFish(finalBuffer,int(bufferSize))			//	抓鱼
		}else if sub_msg_id == SUB_S_START_ALMS {
			this.handleDrawAlm(finalBuffer,int(bufferSize))			//# alms
		}else if sub_msg_id == SUB_S_START_SKILL {
			this.handleUserStartUseSkill(finalBuffer)			//# 申请使用技能返回
		}

	}

	return int(bufferSize) + offset
}
