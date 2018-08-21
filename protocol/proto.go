package main

import (
	"github.com/golang/protobuf/proto"
	t "./tutorial"
	"log"
	"fmt"
	"encoding/json"
)

func main() {
	test := &t.AddressBook {
		People:[]* t.Person{
			{
				Name:"zsw",
				Id:20,
				Email:"emaildd1",
				Phones:[]*t.Person_PhoneNumber{
					{
						Number:"112",
						Type:t.Person_PhoneType(t.Person_WORK),
					},
					{
						Number:"119",
						Type:t.Person_PhoneType(t.Person_HOME),
					},
				},
			},
			{
				Name:"zsw2",
				Id:2,
				Email:"emaildd2",
				Phones:[]*t.Person_PhoneNumber{
					{
						Number: "112",
						Type:t.Person_WORK,
					},
				},
			},
		},
	}


	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("%v",data)
	fmt.Println("")
	
	newTest := &t.AddressBook{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	data_j, err := json.MarshalIndent(newTest, "", " ")
	fmt.Printf("%s",data_j)
	fmt.Println("")

}

