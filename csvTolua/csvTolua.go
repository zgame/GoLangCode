package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("-------------说明------------------")
	fmt.Println("把所有要转换的csv文件扔到app同路径下，运行即可。")
	fmt.Println("-------------csv格式---------------")
	fmt.Println("第一行是列名字")
	fmt.Println("第二行是列类型，支持int,string,bool,float,double")
	fmt.Println("-------------END------------------")

	err:= RunCurrentPahtAllFile()
	if err!=nil{
		fmt.Println("读当前目录下文件出错")
		fmt.Println(err.Error())
	}
	fmt.Println("--------------------end----------------------")
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


// 把新数据结构写入文件中
func writeLuaFiles(filename string,data string) {
	f, err := os.Create(filename+".lua")
	if err != nil {
		log.Fatal(err)
	}
	_,err = f.WriteString(data)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// 获取当前路径下所有文件
func RunCurrentPahtAllFile() error {
	pathname:=getCurrentDirectory()
	//pathname="C:/Users/Administrator/Documents/GitHub/GoLangCode/csvTolua/"
	pathname="./"


	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if !fi.IsDir() {
			files:= strings.Split(fi.Name(), ".")
			fileType := files[1]
			if fileType == "lua"{
				err = os.Remove("./"+fi.Name())
				if err == nil {
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
			fileName := files[0]
			fileType := files[1]
			if fileType == "csv"{
				fmt.Println("开始转换",fi.Name())
				rows := readSample(fi.Name())
				str:= DealRowsData(rows)
				writeLuaFiles(fileName,str)
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


func getString(RowIndex int, ListIndex int, rows [][]string) string {
	var strOut string
	ListName := rows[0][ListIndex] 		// 列名
	ListType := rows[1][ListIndex] // 列类型
	ListType = strings.ToLower(ListType)
	// 增加一个判断,如果是#就忽略
	if ListName[0] == '#'{
		return ""
	}

	data := rows[RowIndex][ListIndex]


	if ListType == "string"{
		strOut = fmt.Sprintf(" %s = \"%s\", ",ListName,data)
	}else if ListType == "int"{
		strOut = fmt.Sprintf(" %s = %s, ",ListName,data)
	}else if ListType == "float"{
		strOut = fmt.Sprintf(" %s = %s, " ,ListName,data)
	}else if ListType == "double"{
		strOut = fmt.Sprintf(" %s = %s, ",ListName,data)
	}else if ListType == "bool"{
		data = strings.ToUpper(data)
		if data == "1" || data == "TRUE"{
			data = "true"
		}else{
			data = "false"
		}
		strOut = fmt.Sprintf(" %s = %s, ",ListName,data)
	}
	return strOut
}



// 处理row数据
func DealRowsData(rows [][]string)  string{
	fmt.Println("-----------------------开始处理数据--------------------------")
	width:= len(rows[0])
	//fmt.Println("width:",width)
	StrOut:= "module = {\n"
	//fmt.Println("",StrOut)
	
	//ListName := rows[0]		// 列名
	//ListType := rows[1]		// 列类型

	for i := range rows {
		if i < 4 {
			continue	// 略过前2行
		}

		line := "["+ rows[i][0]+"] = { "
		for j := 1; j < width; j++ {
			str := getString(i,j, rows)
			line = line + str
		}
		line = line + " },\n"
		//fmt.Println("", line)
		StrOut += line
	}

	StrOut += "} \n return module"

	//fmt.Println("",StrOut)
	return StrOut
}