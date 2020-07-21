package main

import "time"

var Address = []string{"172.16.140.110:9092", "172.16.140.110:9093", "172.16.140.110:9094"}
var TopicName = []string{"Hello-zswcp"}
var GroupName = "group1"

func main() {

	// 生产者
	//SyncProducer(Address)
	//go AsyncProducer(Address,TopicName[0])


	// 消费者
	//for i:=0;i<10;i++ {
	//	go consumer(Address, TopicName[0], i)
	//}
	// 消费组
	//go consumerGroup1(Address ,TopicName, "1")
	//go consumerGroup2(Address ,TopicName, "2")
	//consumerGroup2(Address ,TopicName, GroupName)
	consumerGroup1(Address ,TopicName, GroupName)
	//go consumerGroup1(Address ,TopicName, GroupName)

	for {
		time.Sleep(time.Minute)
		select {

		}
	}
}
