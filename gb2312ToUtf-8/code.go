package main

import (
	"github.com/mahonia"
	"fmt"
)

func main()  {
	str := "这是中文，在大师傅士大夫"
	enc := mahonia.NewEncoder("gb2312")
	output := enc.ConvertString(str)
	fmt.Println("11111",output)
	

	enc2 := mahonia.NewDecoder("gb2312")
	output = enc2.ConvertString(output)
	fmt.Println("222222",output)




}
