package main

import (
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"time"
	"fmt"
	"encoding/json"
)
const (
	Father int = 1
	Mother int = 2
	son int =3
	daughter int =4

)

type Common struct {
	Name string
	Id   int
}

type Person struct {
	Common
	School []string
	Type   int
}

type Home struct {
	Family []Person
	Common
}

type Company struct {
	Staff       []*Person
	Common
	LastUpdated *google_protobuf.Timestamp
}

func main(){
	var myCommany Company
	myCommany.Id = 1
	myCommany.Name ="commony1"

	var tStamp google_protobuf.Timestamp
	tStamp.Seconds = time.Now().UnixNano()
	tStamp.Nanos = int32(time.Now().Nanosecond())

	myCommany.LastUpdated = &tStamp


	//var people [2]Person
	people := make([]Person,2)
	people[0].Name = "people1"
	people[0].Id = 1
	people[0].School = []string{"school1","school2"}
	people[0].Type=Father
	people[1].Name = "people2"
	people[1].Id = 2
	people[1].School = []string{"school2","school4"}
	people[1].Type=Mother

	var people_3 Person
	people_3.Name = "people3"
	people_3.Id = 3
	people_3.School = []string{"school5","school6"}
	people_3.Type = son
	people = append(people, people_3)

	var people_4 * Person
	people_4 = new(Person)
	people_4.Id = 4
	people_4.Name = "people4"
	people_4.School = []string{"school5","school6"}
	people_4.Type = daughter
	people = append(people, *people_4)


	myCommany.Staff = []*Person{&people[0],&people[1],&people[2]}

	//fmt.Printf("%v",myCommany)
	fmt.Println("---------------------------------")
	data, _ := json.MarshalIndent(myCommany, "", " ")
	fmt.Printf("%s",data)


	var myHome Home
	myHome.Id = 1
	myHome.Name = "home1"
	myHome.Family = people[:4]

	//fmt.Printf("%v",myHome)
	fmt.Println("-------------------------")
	data, _ = json.MarshalIndent(myHome, "", " ")
	fmt.Printf("%s",data)


	test:= Home{
		Family:[]Person{
			people[0],people[1],
		},
		Common:Common{
			Id:2,
			Name:"home2",
		},
	}

	fmt.Println("-------------------------")
	data, _ = json.MarshalIndent(test, "", " ")
	fmt.Printf("%s",data)

	test_commany := Company{
		Staff:[]*Person{
			&people[0],
			&people[1],
			&people[2],
		},
		Common:Common{
			Id:2,
			Name:"commany2",
		},
		LastUpdated:&google_protobuf.Timestamp{
			Seconds:time.Now().UnixNano(),
			Nanos:int32(time.Now().UnixNano()),
		},
	}

	fmt.Println("-------------------------")
	data, _ = json.MarshalIndent(test_commany, "", " ")
	fmt.Printf("%s",data)




}