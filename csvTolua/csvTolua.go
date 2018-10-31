package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("-------------说明------------------")
	fmt.Println("把所有要转换的csv文件扔到app同路径下，运行即可。")
	fmt.Println("put all csv files in this path, run!")
	fmt.Println("-------------END------------------")

	err:= RunCurrentPahtAllFile()
	if err!=nil{
		fmt.Println("读当前目录下文件出错")
		fmt.Println(err.Error())
	}
	//rows := readSample()
	//appendSum(rows)
	//writeLuaFiles(rows)
}

func readSample(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func appendSum(rows [][]string) {
	rows[0] = append(rows[0], "SUM")
	for i := 1; i < len(rows); i++ {
		rows[i] = append(rows[i], sum(rows[i]))
	}

	for i,v := range rows{
		fmt.Print(i,"_",v,"		|")
		fmt.Println("")


	}
}

func sum(row []string) string {
	sum := 0
	for _, s := range row {
		x, err := strconv.Atoi(s)
		if err != nil {
			return "NA"
		}
		sum += x
	}
	return strconv.Itoa(sum)
}

// 把新数据结构写入文件中
func writeLuaFiles(filename string,rows [][]string) {
	f, err := os.Create(filename+".lua")
	if err != nil {
		log.Fatal(err)
	}
	err = csv.NewWriter(f).WriteAll(rows)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// 获取当前路径下所有文件
func RunCurrentPahtAllFile() error {
	pathname:=getCurrentDirectory()
	pathname="C:/Users/Administrator/Documents/GitHub/GoLangCode/csvTolua/"


	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if !fi.IsDir() {
			files:= strings.Split(fi.Name(), ".")
			fileType := files[1]
			if fileType == "lua"{
				del := os.Remove("./"+fi.Name())
				if del != nil {
					fmt.Println("删除老的文件",fi.Name())
				}
			}
		}
	}
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			//RunCurrentPahtAllFile(pathname + fi.Name() + "\\")
		} else {
			files:= strings.Split(fi.Name(), ".")
			fileType := files[1]
			if fileType == "csv"{
				fmt.Println("开始转换",fi.Name())
				rows := readSample(fi.Name())
				DealRowsData(rows)
			}
		}
	}
	return err
}

// 获取当前目录的路径
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}


func getString(index int, rows [][]string) string {
	ListType := rows[1]		// 列类型

	if ListType == "string"{
		
	}
	return rows
}



// 处理row数据
func DealRowsData(rows [][]string)  {
	fmt.Println("-----------------------开始处理数据--------------------------")
	//width:= len(rows[0])
	//fmt.Println("width:",width)
	StrOut:= "module = {}\n	module = {\n"
	fmt.Println("",StrOut)
	
	ListName := rows[0]		// 列名
	ListType := rows[1]		// 列类型

	for _,v := range rows{
		line:="["


		fmt.Print(i,"_",v,"		|")
		fmt.Println("")
	}



	for i:=0;i<width;i++{
		
	}
	fmt.Println("",rows[0][1])
	fmt.Println("",rows[1][1])

	for i,v := range rows{
		fmt.Print(i,"_",v,"		|")
		fmt.Println("")
	}

}