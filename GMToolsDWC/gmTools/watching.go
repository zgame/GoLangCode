package main

import (
	"fmt"
	"time"
	"github.com/go-ini/ini"
	"strconv"
	"github.com/mahonia"

)

func watchGameServer()  {

	defer func() {
		if e := recover(); e != nil {
			logerDump()
		}
	}()

	f, _ := ini.Load("Setting.ini")
	watch,_ := f.Section("Daemon").Key("switch").Bool()
	//gameDir := f.Section("Daemon").Key("dir").Value()

	centerId,_ := f.Section("Daemon").Key("center_id").Int()
	centerDir := f.Section("Daemon").Key("center_dir").Value()
	centerIp := f.Section("Daemon").Key("center_ssh_ip").Value()
	centerPort,_ := f.Section("Daemon").Key("center_ssh_port").Int()
	centerUser := f.Section("Daemon").Key("center_ssh_user").Value()
	centerPwd := f.Section("Daemon").Key("center_ssh_pwd").Value()


	serverList := getServerList()
	daemonList := make([]int,0)
	for _,v := range serverList{
		if v.Daemon ==1{
			// 有需要守护的进程
			fmt.Println("有需要守护的进程:",v.Serverid)
			daemonList = append(daemonList, v.Serverid)
		}
	}
	// 加入守护中心服
	daemonList = append(daemonList, centerId)

	fmt.Printf("守护的进程:%v", daemonList)
	fmt.Println("")
	fmt.Println("------------------------")

	if watch {
		for {
			fmt.Println("--------watching game server-----------------")
			time.Sleep(5000 * time.Millisecond)

			//fmt.Printf("%v", serverList)
			for _, v := range mwGlobal.model.items {
				for i := range daemonList {
					if v.ServerId == daemonList[i] {
						//有游戏进程需要守护
						if v.ServerState == 1 {
							// 关机了，需要重启
							// 游戏服务器重启
							for _, v2 := range serverList {
								if v2.Serverid == v.ServerId { // 找到ssh的地址和信息
									user := v2.Ssh_user
									psw := v2.Ssh_passwd
									host := v2.Ipaddr
									port := v2.Ssh_port
									serverId := v2.Serverid
									serverName := v2.Room_name

									//拉起守护进程
									// 老版本，读配置文件
									//cmd := "cmd.exe /c \"start " + gameDir + " /ServerID:" + strconv.Itoa(serverId) + " /ServerName:" + serverName
									// 新版本，读数据库
									cmd := "cmd.exe /c \"start " + v2.Daemon_address + " /ServerID:" + strconv.Itoa(serverId) + " /ServerName:" + serverName

									//转换一下到gb2312
									enc := mahonia.NewEncoder("gb2312")
									cmd = enc.ConvertString(cmd)

									fmt.Println("cmd: ", cmd)
									zswSSH(user, psw, host, port, cmd)
								}
							}
							// 中心服务器重启
							if v.ServerId == centerId{
								cmd := "cmd.exe /c \"start " + centerDir				// 中心服的地址不一样
								fmt.Println("cmd: ", cmd)
								zswSSH(centerUser, centerPwd, centerIp, centerPort, cmd)
							}
							loger("守护进程重启:" + strconv.Itoa(v.ServerId))

						} else {
							//fmt.Println("没关机，不用重启", v.ServerId)
						}
					}
				}

			}

		}
	}

}