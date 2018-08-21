package main

const (
	SUB_C_REQ_BANKER = 600 //# 申请上庄
	SUB_S_REQ_BANKER = 601
	//#  坐庄
	SUB_S_BANKER_REQ_RESULT  = 602 //# 申请结果
	SUB_S_BANKER_SITUSERINFO = 603 //# 坐庄信息
	SUB_S_ChipStart          = 604 //# 下注开始
	SUB_S_ChipStop           = 605 //# 下注结束

	SUB_C_UserChip         = 606 //# 玩家下注
	SUB_S_ChipSucc         = 607 //# 下注成功
	SUB_S_Control_UserChip = 608 //# 玩家下注

	SUB_C_ContinueChip = 609 //# 玩家续投
	SUB_S_ContinueChip = 610 //# 续投回应

	SUB_S_GameStart       = 611 //# 游戏开始
	SUB_S_GameEnd         = 612 //# 游戏结束
	SUB_S_Game_ResultInfo = 613 //# 庄家游戏结算信息

	SUB_C_Control_SetDst    = 614 //# 控制结果
	SUB_C_Control_SelectDst = 615 //# 选中目标玩家
	SUB_S_Control_DealRet   = 616 //# 处理结果

	SUB_C_ClearChip = 617 //# 清空投注
	SUB_S_ClearChip = 618 //# 清空投注回应

	SUB_C_UpdateSelfScore = 619 //# 更新自己分数
	SUB_S_UpdateSelfScore = 620 //# 更新自己分数

	SUB_S_Sence_Free             = 621 //# 空闲场景
	SUB_S_GameResultInfo         = 622 //# 普通玩家游戏结算信息
	SUB_S_BondLost               = 623 //# 奖池空了
	SUB_S_UpdataChip             = 624 //# 更新下注
	SUB_S_Control_Update         = 625 //# 控制数据更新
	SUB_S_SysMessage             = 626 //# 系统消息
	SUB_S_UpdateUserInfo         = 627 //# 玩家信息
	SUB_S_GameInfo               = 628 //# 游戏信息
	SUB_S_GameOpenLog            = 629 //# 游戏开奖纪录
	SUB_S_SelfResult             = 630 //# 自己计算信息
	SUB_S_BankerResult           = 631 //# 庄家结果信息
	SUB_S_Control_UserInfo       = 632 //# 玩家信息
	SUB_S_Control_StockInfo      = 633 //# 库存信息
	SUB_S_Control_UserLeave      = 634 //# 玩家离开
	SUB_S_UserChipNotify         = 635 //# 玩家下注数据广播给所有用户
	SUB_S_UserContinueChipNotify = 636 //# 玩家续投数据广播给所有用户
	SUB_S_AndroidUpdataChip      = 637 //# 通知机器人下注
	SUB_S_UserTotalChipNotify    = 638 //# 通知超级管理员总下注信息
	SUB_C_ONLINE_TOTAL           = 639 //# 请求服务器在线人数
	SUB_S_ONLINE_TOTAL           = 640 //# 服务器返回在线人数
	SUB_C_ROBBANKREQUEST         = 641 //# 客户端请求抢庄
	SUB_C_StartGame_Single       = 642 //# 单人游戏开始
)
