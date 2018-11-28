package main

import (
	"fmt"
	"flag"
)
//-name=zsw -married=false -age=22 
func main()  {
	name := flag.String("name", "", "")
	age := flag.Int("age", 0, "")
	married := flag.Bool("married", false, "")
	flag.Parse()
	fmt.Println("name: ",*name)
	fmt.Println("age: ",*age)
	fmt.Println("age: ",*married)

}
