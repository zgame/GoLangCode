package main

import (
	"fmt"
	"reflect"
)

type people struct {
	id int
	name string
}
func (p *people) PrintName()  {
	fmt.Println("名字：", p.name)
}



type student struct {
	people
	Class int   // 班级
}

func (s * student)newStudent() *student {
	return &student{people{1,"s2"},2}
}
func (p *student) PrintName()  {
	fmt.Println("学生名字：", p.name)
}



type worker struct {
	people
	company  string  //公司
}

//func (p *worker) PrintName()  {
//	fmt.Println("名字：", p.name)
//}



func main()  {
	var s1  student
	s1.name = "s1"
	s1.id = 1
	s1.Class = 2
	s1.PrintName()

	var w1 worker
	w1.name = "w1"
	w1.id= 2
	w1.company = "w1 company"
	w1.PrintName()

	var s2 * student
	s2 = s2.newStudent()
	s2.PrintName()

	s3 := new(student)
	s3.name = "s3"
	s3.id = 4
	s3.PrintName()


	var s4 * student
	s4 = s4.newStudent()
	add(s4.people)


	var s5 * student
	s5 = s5.newStudent()
	add(s5)
}

func add(ss interface{})  {
	v := reflect.ValueOf(ss)
	fmt.Println("", v.Type())
	if v.Type() == reflect.ValueOf(people{}).Type(){
		fmt.Println("1111111111")
		tt := ss.(people)
		tt.PrintName()
	}else if v.Type() == reflect.ValueOf(&student{}).Type(){
		fmt.Println("2222222222")
		tt := ss.(*student)
		tt.PrintName()
	}
	//tt:= ss.(v.Type())
	//fmt.Println("",tt)
	//ss.PrintName()
}


