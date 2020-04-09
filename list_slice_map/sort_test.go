package main

import (
	"math"
	"testing"
	"sort"
	"fmt"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}



// 正向排序
func TestSortIntSlice(t *testing.T) {
	data := ints
	a := sort.IntSlice(data[0:])
	//fmt.Printf("%v \n", a)

	// 排序
	sort.Sort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
	fmt.Printf("%v \n", a)
}

// 反向排序
func TestReverse(t *testing.T)  {
	data := ints
	a := sort.IntSlice(data[0:])
	//fmt.Printf("%v \n", a)

	// 翻转
	sort.Sort(sort.Reverse(a))
	fmt.Printf("%v \n", a)
}