// world server入口文件
package main

//
import (
	"./core/logs"
	"./core/server"
	"./world/world"
)

// 程序入口
func main() {
	defer logs.PrintPanic()
	logs.InitLogs()

	server.Run(world.New())
}
