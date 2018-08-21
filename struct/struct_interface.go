package main

import (
	"fmt"
)
//------------------------People-------------------------------
type People struct {
	id int
	name string
}

func (p *People)New(name string) Tinterface{
	return &People{1,name}
}

func (p *People) pPrintName()  {
	fmt.Println("名字：", p.name)
}

//------------------------Student-------------------------------
type Student struct {
	People
	Class int   // 班级
}

func (s * Student)newStudent(name string) *Student {
	return &Student{People{1,name},2}
}

//  该函数为重载， 如果不定义的话， 那么就用父类的方法
func (p *Student) pPrintName()  {
	fmt.Println("学生名字：", p.name)
}

//-----------------------Worker---------------------------
type Worker struct {
	People
	company  string  //公司
}
func (p *Worker)New(name string) Tinterface{
	return &Worker{People{1,name},""}
}
//  该函数为重载， 如果不定义的话， 那么就用父类的方法

//func (p *Worker) pPrintName()  {
//	fmt.Println("工人名字：", p.name)
//}

//-----------------------interface---------------------------
type Tinterface interface {
	New(name string) Tinterface
	pPrintName()
}






//-----------------------main---------------------------
func main()  {

	//------------------------Student-------------------------------
	var s1  Student
	s1.name = "s1"
	s1.id = 1
	s1.Class = 2
	s1.pPrintName()


	var s2 * Student
	s2 = s2.newStudent("s2")
	s2.pPrintName()

	s3 := new(Student)
	s3.name = "s3"
	s3.id = 4
	s3.pPrintName()


	var s4 * Student
	s4 = s4.newStudent("s4")
	Add(s4)


	var s5 * Student
	s5 = s5.newStudent("s5")
	Add(s5)

	//var s6 *Student
	//Add2(s6,"s6")

	//-----------------------------Worker--------------------------
	var w1 Worker
	w1.name = "w1"
	w1.id= 2
	w1.company = "w1 company"
	w1.pPrintName()

	var w2 *Worker
	w := w2.New("w2")
	w.pPrintName()

	//----------------------------People---------------------------
	var p1 People
	p1.name = "p1"
	Add(&p1)


	var p2 *People
	p := p2.New("p2")
	p.pPrintName()


	var p3 *People
	Add2(p3,"p3")

	//-------------------------------------------------------
}

func Add(ss Tinterface)  {
	ss.pPrintName()

}

func Add2(ss Tinterface, name string)  {
	p := ss.New(name)
	p.pPrintName()
}



