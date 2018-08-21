package Client

//-----------------------------------------------------------
//	代码结构说明
//----------------------------------------------------------



// 服务器架构  树形管理结构

/***
  GameServer  - 游戏服务器掌握多个客户端连接

				网络接收
				- Clients	- 客户端连接掌握玩家的信息和game管理器句柄
							- AllUserClientList - 每个连接掌握自己的控制游戏的句柄，方便自己控制
												- Player
												- game句柄
												- Table句柄
												- User句柄

				游戏管理，服务器逻辑发起
				- GameManager	- 游戏管理器掌握所有游戏
								- AllGamesList	- 单个游戏掌握所有桌子和玩家
												- AllUserList
												- AllTableList	- 桌子掌握所有子弹和鱼，还有坐下的玩家
															- FishArray
															- BulletArray
															- UserSeatArray	- 游戏中玩家数据和玩家总数据
																			- Player
																			- net.conn - 掌握网络的socket连接句柄





***/


// 句柄

/***
服务器有所有Client的句柄，AllClientsList， 方便查询
服务器有所有game的句柄，AllGamesList

Client 有当前Game；当前桌子；当前Player和user的句柄， 当前玩家收到消息后可以方便调用游戏逻辑

Game 有所有Table的句柄， AllTableList
Game 有所有user的句柄， AllUserList

Table 有所有坐下玩家的句柄，UserSeatArray
Table 有所有鱼的句柄，FishArray
Table 有所有子弹的句柄，BulletArray


***/








