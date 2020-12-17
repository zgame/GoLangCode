package Lua

import (
	"GoLuaServerV2.1/Core/NetWork"
	"GoLuaServerV2.1/Core/Utils/zLog"
)

func ReadDataPackage( buf []byte, Uid int) (int,int,int,string) {
	//fmt.Printf("buf......%x",buf)
	//-----------------------------头部数据不完整----------------------------
	if len(buf)< NetWork.TCPHeaderSize {
		//str:= fmt.Sprintf("%d数据包头部数据不全 : %x \n",a.UserId,buf)
		return 0 , 0,0,""
	}
	//-----------------------------解析头部信息----------------------------
	headFlag,msgId, subMsgId, bufferSize, _ , msgSize := NetWork.DealRecvTcpHeaderData(buf)
	BufAllSize := NetWork.TCPHeaderSize + int(bufferSize)+ int(msgSize) + 1 // 整个数据包长度，末尾有标示位

	////-----------------------------头部信息错误----------------------------
	if headFlag != uint8(254){
		zLog.PrintfLogger("%d 数据包头部标识不正确 %x",Uid, buf)
		return -1 	, 0,0,""		// 数据包格式校验不正确
	}

	//-----------------------------错误提示----------------------------

	//-----------------------------proto buffer 内容不完整----------------------------
	if len(buf) < BufAllSize{
		//str:= fmt.Sprintf("%d数据包格式不正确buflen%d,bufferSize%d,%x  \n", a.UserId,len(buf),int(bufferSize),buf)
		StaticDataPackageProtoDataLess++
		return  0, 0,0,"" //int(bufferSize) + offset
	}

	//// ------------------------数据包尾部的判断----------------------
	endData := NetWork.DealRecvTcpEndData(buf[BufAllSize -1 :BufAllSize])
	if endData!= uint8(NetWork.TCPEnd){ // EE
		zLog.PrintfLogger("%d数据包尾部判断不正确 %x ",Uid, buf)
		return -1, 0,0,""
	}

	//-----------------------------数据包重复----------------------------

	//-----------------------------取出proto buffer的内容----------------------------

	finalBuffer := buf[NetWork.TCPHeaderSize : NetWork.TCPHeaderSize+ int(bufferSize)]

	//a.TokenId = int(tokenId) // 记录当前最后接收的数据包编号，防止重复
	//a.TokenTime = ztimer.GetOsTimeMillisecond()


	return BufAllSize,int(msgId), int(subMsgId), string(finalBuffer)

}