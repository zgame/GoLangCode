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
		//Zswnameb:-2147483648,
		//Address:true,
		Ad:1,
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

