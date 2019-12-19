package Games

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"../ProtocolBuffer/CMD"
	"../Core/Utils/zLog"
	"../Core/GameCore"
	"./Model/PlayerModel"
)


//*******************************************************进入游戏场景申请******************************************************
func HandelEnterScene(player *PlayerModel.Player, game *GameCore.Games ,  buf []byte){
	fmt.Println("------------客户端申请进入游戏场景-------------")
	protocolBuffer := buf
	msg := &CMD.CMD_GF_GameOption{}
	err := proto.Unmarshal(protocolBuffer, msg)
	zLog.CheckError(err)

	fmt.Println("客户端申请进入游戏场景, GetClientVersion: ", msg.GetClientVersion())

	//-------------------------------------------逻辑------------------------------------------------------------
	tableUid := int32(game.PlayerLoginGame(player, 1 ))   // 玩家登陆游戏，分配桌子
	tableInterface := game.GetTableByUID(int(tableUid)) // 设置桌子句柄

	// 跳转到各自桌子执行
	//fmt.Println("-----跳转到各自桌子执行")
	tableInterface.EnterSceneSyncMsg(player)			// 跳转到各个桌子自己去处理同步信息


	//client.User = client.Games.GetUserByUID(int(client.PlayerModel.UserId))	// 设置玩家的句柄
	//client.User.SetConn(client.Conn)								// 设置玩家的连接句柄
	//client.User.SetIsRobot(false)							// 连接进来的都不是机器人

}

