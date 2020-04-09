package main

import (
	"testing"
	"container/list"
	"fmt"
)

func TestDoubleList(t *testing.T){
	dl:= new(list.List)

	dl.PushFront(1)
	dl.PushFront(2)
	ss := dl.PushBack("s")
	dl.InsertBefore("sf",ss)

	for e:= dl.Front(); e!=nil; e=e.Next(){
		fmt.Println("",e.Value)
	}


}
