package Games

import "../Core/ZServer"

// 回调函数注册
func NetWorkFuncRegister()  {
	ZServer.NetWorkReceive = NetWorkReceive		// 绑定接收网络消息处理函数回调地址
	ZServer.NetWorkInit = NetWorkInit			// 绑定网络连接成功初始化
	ZServer.NetworkBroken = NetWorkBroken			// 网络中断回调
}


//---------------------------------------消息分发--------------------------------------

// 分发处理消息
func NetWorkReceive(serverId int,userId int, msgId int, subMsgId int, data  []byte, token int) {


}

//  网络连接成功初始化
func NetWorkInit(serverId int)  {

}

//  网络中断回调
func NetWorkBroken(serverId int,userId int)  {

}
