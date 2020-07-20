package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 打印主题列表， broker列表
func KafKaInfo(address []string)  {
	fmt.Println("--------------kafka info--------------")
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_0_0
	client, err := sarama.NewClient(address, config)
	if err != nil {
		panic("client create error")
	}
	defer client.Close()
	//获取主题的名称集合
	topics, err := client.Topics()
	if err != nil {
		panic("get topics err")
	}
	for _, e := range topics {
		fmt.Println(e)
	}
	//获取broker集合
	brokers := client.Brokers()
	//输出每个机器的地址
	for _, broker := range brokers {
		fmt.Println(broker.Addr())
	}
	fmt.Println("-------------info end-------------")

}
