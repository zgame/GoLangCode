package main

import (
	"fmt"
	. "./const"
	"./NetWork"
	"./log"
)

// 处理单个包内容
func (this *Client)handlerRead(buf []byte) int {
	//var err error
	//fmt.Printf("Receive buf: %x\n",buf)


	//------------------------头部判断----------------------
	if len(buf)< NetWork.TCPHeaderSize {		// 接受不全，那么缓存
		log.PrintfLogger("数据包头部小于 10 : %x   \n",buf)
		GlobalMutex.Lock()
		StaticDataPackageHeadLess++
		GlobalMutex.Unlock()

		this.ReceiveBuf = buf
		//str:= fmt.Sprintf("%d数据包头部小于 10 : %x   ",this.Index,buf)
		//this.PrintLogger(str)
		return 0
	}

	// 读取头部信息
	headFlag, msg_id, sub_msg_id, bufferSize, _, msgSize := NetWork.DealRecvTcpHeaderData(buf)

	if headFlag != uint8(254){		// FE
		GlobalMutex.Lock()
		StaticDataPackageHeadFlagError++
		GlobalMutex.Unlock()
		log.WritefLogger("%d数据包头部判断不正确 %x ",this.Index, buf)
		//this.PrintLogger(str)
		return -1 			// 数据包格式校验不正确
	}

	//offset := NetWork.TCPHeaderSize
	BufAllSize := NetWork.TCPHeaderSize + int(bufferSize)+ int(msgSize) + 1    // 整个数据包长度，末尾有标示位

	//if msgSize > 200 {
	//	str:= fmt.Sprintf("%d错误消息%d",this.Index,msgSize)
	//	this.PrintLogger(str)
	//}

	//fmt.Println("len(buf)",len(buf))
	//fmt.Println("offset",offset)
	//fmt.Println("bufferSize",bufferSize)
	//------------------------proto buffer 判断----------------------
	if len(buf) < BufAllSize {	// 接受不全，那么缓存
		this.ReceiveBuf = buf
		//fmt.Printf("出现数据包异常buflen=%d,bufferSize=%d,  msgSize=%d  %x  \n",len(buf),int(bufferSize),int(msgSize)buf)
		//str:= fmt.Sprintf("%d出现数据包异常buflen=%d,bufferSize=%d,  msgSize=%d  %x  ",this.Index,len(buf),int(bufferSize),int(msgSize),buf)
		//this.PrintLogger(str)
		//this.PrintLogger("数据包不完整buflen=" +strconv.Itoa(len(buf))  +"bufferSize="+ strconv.Itoa(int(bufferSize)) +"msgSize="+strconv.Itoa(int(msgSize))+"buf="+string(this.ReceiveBuf))
		GlobalMutex.Lock()
		StaticDataPackageProtoDataLess++
		GlobalMutex.Unlock()
		return  0 //int(bufferSize) + offset + int(msgSize)
	}
	//if ver > 0{
	//	offset = 12		// version == 1 的时候， 加了一个token
	//}


	// ------------------------数据包尾部的判断----------------------
	endData := NetWork.DealRecvTcpEndData(buf[BufAllSize -1 :BufAllSize])
	if endData!= uint8(NetWork.TCPEnd){		// EE
		log.WritefLogger("%d数据包尾部判断不正确 %x ",this.Index, buf)
		return -1
	}
	// ------------------------错误提示的判断----------------------
	if msgSize >0 {
		//fmt.Println("有错误提示了")
		//msgBuffer := buf[offset + int(bufferSize):offset + int(bufferSize)+ int(msgSize)]
		//fmt.Println(string(msgBuffer))
		return BufAllSize
	}



	//----------------------------解析 proto buffer-----------------------------------
	finalBuffer := buf[NetWork.TCPHeaderSize:NetWork.TCPHeaderSize + int(bufferSize)]
	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	//fmt.Println("msg_id",msg_id,"sub_msg_id",sub_msg_id)

	//# -----------------login server msg-----------------
	if msg_id == MDM_MB_LOGON {
		if sub_msg_id == SUB_MB_LOGON_SUCCESS {
			this.handleLoginSucess(finalBuffer,int(bufferSize))
			fmt.Println("login登录成功 ------ ", this.Index)
		} else if sub_msg_id == SUB_MB_LOGON_FAILURE {
			this.handleLoginFailed(finalBuffer,int(bufferSize))
		} else if sub_msg_id == SUB_MB_LOGON_FINISH {
			fmt.Println("login登录服成功 ------ ", this.Index)
			//this.ConnectGameServer("")
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
			//addr := this.Serverlist[0].server_addr + ":" + strconv.Itoa(this.Serverlist[0].server_port)
			//this.ConnectGameServer(addr)


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
			fmt.Println("游戏服务器登录完成 ------")

			//  开始进入场景
			//fmt.Println("开始发送进入场景",this.User.user_id)
			this.EnterScence()

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
			//this.SendGmCmd("@设置金币 10000000")
			//this.do_fire()



			this.StartAI = true

		} else if sub_msg_id == SUB_S_OTHER_ENTER_SCENE {
			this.handleOtherEnterScence(finalBuffer,int(bufferSize))			//进入场景,接收鱼数据
		}else if sub_msg_id == SUB_S_SCENE_FISH {
			this.handleSceneFish(finalBuffer,int(bufferSize))			//# 新生成鱼
		}else if sub_msg_id == SUB_S_DISTRIBUTE_FISH {
			this.handleNewFish(finalBuffer,int(bufferSize))			//# 新生成鱼
			//fmt.Println("新生成鱼 ------ ", this.Index)

			//------------------捕鱼----------------------
		}else if sub_msg_id == SUB_S_USER_FIRE {
			this.handleUserFire(finalBuffer,int(bufferSize))			// 开火
			//fmt.Println("开火收到 ------ ", this.Index)
		}else if sub_msg_id == SUB_S_CATCH_FISH {
			this.handleCatchFish(finalBuffer,int(bufferSize))			//	抓鱼
			//fmt.Println("抓鱼收到 ------ ", this.Index)
		}else if sub_msg_id == SUB_S_START_ALMS {
			this.handleDrawAlm(finalBuffer,int(bufferSize))			//# alms
		}else if sub_msg_id == SUB_S_BOSS_COME {
			// 统计网络延迟用
			if this.ShowMsgSendTime {
				now := this.GetOsTime()
				end := now - this.SendMsgTime
				log.PrintfLogger("消息间隔时间：%d毫秒|  发子弹  %d  收子弹  %d  打鱼  %d  打到鱼  %d  头部不全 %d   数据包不全  %d  粘贴数量 %d  粘贴成功  %d  头部信息错误  %d  ",
					int(end),this.ShowMsgFire,this.ShowMsgReFire,this.ShowMsgCatchFish ,this.ShowMsgReCatchFish,StaticDataPackageHeadLess,StaticDataPackageProtoDataLess,StaticDataPackagePasteNum,StaticDataPackagePasteSuccess,StaticDataPackageHeadFlagError)
				if IsWebSocket{
					log.PrintfLogger(" 当前在线人数 socket 0 wsocket %d " ,wsclients[0].Number())
				}else{
					log.PrintfLogger(" 当前在线人数 socket %d wsocket 0  " ,clients[0].Number())
				}

				this.ShowMsgFire = 0
				this.ShowMsgReFire= 0
				this.ShowMsgCatchFish=0
				this.ShowMsgReCatchFish=0

				//GlobalMutex.Lock()
				//sendMsgNum = 0
				//receiveMsgNum =0
				//GlobalMutex.Unlock()
			}
		}
	}


	//GlobalMutex.Lock()
	this.ReceiveMsgNum++
	//GlobalMutex.Unlock()
	this.SuccessBuf = buf 	// 记录最后一次成功的buf

	return BufAllSize
}
