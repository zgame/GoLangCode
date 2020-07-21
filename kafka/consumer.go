package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
)

//Consumers have two modes of operation.

// 消费组2   Users who require access to individual partitions can use the partitioned mode which exposes access to partition-level consumers:
func consumerGroup2(address, topicName []string, groupId string)  {

	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions
	//config.

	consumer, err := cluster.NewConsumer(address, groupId, topicName, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// consume partitions
	for {
		select {
		case part, ok := <-consumer.Partitions():
			if !ok {
				return
			}
			// start a separate goroutine to consume messages
			go func(pc cluster.PartitionConsumer) {
				for msg := range pc.Messages() {
					fmt.Fprintf(os.Stdout, "Topic:%s/Partition:%d/Offset:%d\t Key:%s\t Value:%s\t GroupId:%s \n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value,groupId)
					consumer.MarkOffset(msg, "")	// mark message as processed
				}
			}(part)
		}
	}
}

// 消费组1  In the default multiplexed mode messages (and errors) of multiple topics and partitions are all passed to the single channel:
func consumerGroup1(address []string, topicName []string , group string) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = false
	// init consumer
	consumer, err := cluster.NewConsumer(address, group, topicName, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "Topic:%s/Partition:%d/Offset:%d\t Key:%s\t Value:%s\t GroupId:%s \n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value,group)
				consumer.MarkOffset(msg, "")	// mark message as processed
			}
		}
	}
}


// 普通接收者， 只能接收一个主题
func consumer(address []string,topicName string, partition int)  {
	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V2_3_0_0
	//新建一个消费者
	consumer, e := sarama.NewConsumer(address, config)
	if e != nil {
		panic("error get consumer"+e.Error())
	}
	defer consumer.Close()

	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err := consumer.ConsumePartition(topicName , int32(partition), sarama.OffsetNewest)    // 只接收最新消息
	//partitionConsumer, err := consumer.ConsumePartition(topicName, int32(partition), sarama.OffsetOldest)  // 从第一个开始一直到最后
	if err != nil {
		fmt.Println("error get partition consumer", err)
	}
	defer partitionConsumer.Close()
	//循环等待接受消息.
	for {
		select {
		//接收消息通道和错误通道的内容.
		case msg := <-partitionConsumer.Messages():
			fmt.Println("msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Println(err.Err)
		}
	}

}
