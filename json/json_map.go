package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type TestS struct {
	IP   string
	Name string
	ID   float32
}

func main() {
	var test1 TestS

	b := []byte(`{"IP": "127.0.0.1", "name": "SKY","ID":1.2}`)

	m := make(map[string]interface{})

	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
	err = json.Unmarshal(b, &test1)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}

	fmt.Println("m:", m)
	fmt.Println("m.IP:", reflect.TypeOf(m["IP"]))
	fmt.Println("m.ID:", reflect.TypeOf(m["ID"]))

	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
	fmt.Println("test1:", test1)
	fmt.Println("test1.IP:", test1.IP)

	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
	for k, v := range m {
		fmt.Println(k, ":", v)
	}

}
