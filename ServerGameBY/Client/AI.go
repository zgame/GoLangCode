package Client

//
//// game AI
//func (client *Client)GameAI()  {
//	if !client.StartAI{
//		return
//	}
//	//fmt.Println("AI-----")
//	if time.Now().After(client.LastFireTick)   {
//		client.LastFireTick = time.Now().Add( time.Microsecond * 200)
//		//this.do_fire()
//	}
//
//	//# 处理过期的鱼
//	if time.Now().After(client.LastCheckDueTick) {
//		client.LastCheckDueTick = time.Now().Add(time.Second)
//		client.check_overdue_fish()
//	}
//
//}
//
//
//// 选择鱼
//func (client *Client)select_fish() {
//	//# 1s 选择一次
//
//	if time.Now().After(client.SelectTick) {
//		client.SelectTick = time.Now().Add(time.Second)
//
//		if client.FailedCnt < 10 && client.FishId != 0 {
//			return
//		}
//
//
//		if client.FailedCnt > 10 && client.FishId != 0 { // 删除鱼
//			for i,v:=range client.MyGameInfo.fishPool {
//				if v.uid == client.FishId {
//					//fmt.Println("开了10炮打不死",v.uid)
//					client.MyGameInfo.fishPool = append(client.MyGameInfo.fishPool[:i], client.MyGameInfo.fishPool[i+1:]...)
//				}
//			}
//		}
//
//		if len(client.MyGameInfo.fishPool) > 0 {
//			rand.Seed(time.Now().UnixNano())
//			x := rand.Intn(len(client.MyGameInfo.fishPool))
//			client.FishId = client.MyGameInfo.fishPool[x].uid
//
//			//fmt.Println("重新选择鱼id", this.Fish_id)
//		}
//	}
//
//}
//
//// 定期删除过期的鱼
//
//func (client *Client)check_overdue_fish(){
//	for i,v:= range client.MyGameInfo.fishPool {
//		if time.Now().After(v.tick.Add(time.Second*10)){
//			if client.FishId == v.uid{
//				client.FishId = 0
//			}
//			//fmt.Println("鱼过期了", v.uid)
//			//fmt.Println("删除定期的鱼", i, "--len" , len(this.MyGameInfo.fish_pool) )
//			//if len(this.MyGameInfo.fish_pool) <= 1{
//			//	this.MyGameInfo.fish_pool = make([]*FishObj, 0)
//			//}else{
//			client.MyGameInfo.fishPool = append(client.MyGameInfo.fishPool[:i], client.MyGameInfo.fishPool[i+1:]...)
//			//}
//			return
//		}
//	}
//
//}
//
