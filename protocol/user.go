package main

import (
	"github.com/golang/protobuf/proto"
	t "./tutorial"
	"log"
	"fmt"
	"encoding/json"
)

func main() {
	test := &t.Friend {
		Zswnameb:7999932339.999867,
		Address:false,
		Ad:1,
		Ad1:-9223372036854775800,
		Ad2:9223372036854775800,

	}


	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("%v",data)
	fmt.Println("")

	newTest := &t.Friend{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	data_j, err := json.MarshalIndent(newTest, "", " ")
	fmt.Printf("%s",data_j)
	fmt.Println("")

}

