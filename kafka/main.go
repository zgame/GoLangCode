package main

var Address = []string{"172.16.140.110:9092", "172.16.140.110:9093", "172.16.140.110:9094"}
var TopicName = []string{"Hello-zswc"}
var GroupName = "group1"

func main() {
	//SyncProducer(Address)
	//AsyncProducer(Address)


	//topic := []string{"Hello-zswc"}
	//广播式消费：消费者1
	//go clusterConsumer( Address, topic, "group-1")
	////广播式消费：消费者2
	//go clusterConsumer( Address, topic, "group-2")
	//KafKaInfo(Address)
	//consumer(Address,TopicName)
	//consumerGroup(Address,TopicName,"group1")
	//consumerGroup1(Address ,TopicName, GroupName)
	clusterConsumer(Address ,TopicName, GroupName)

	for {
		select {}
	}
}
