package main

import "fmt"

type Zst struct{
	zst1 string

}

type zst222 struct {
	zst2 string
	Zst
}

type zsw333 struct {
	
}

func (s * zsw333)init()  {

	
}

type zstinterface interface {
	init()
}

func (s *Zst)init(num string)  {
	s.zst1 = "test init zsw --- " + num
}


func main()  {
	var zzz *Zst
	zzz = &Zst{zst1:""}
	fmt.Println("zzz:",zzz.zst1)

	zzz.init("zz1")
	fmt.Println("zzz:",zzz.zst1)


	//zz2 := &zst222{zst2:"", Zst:zzz}
	zz2 := &zst222{zst2:""}

	zz2.init("zz2")

	fmt.Println("zzz:",zz2.zst1)


}
