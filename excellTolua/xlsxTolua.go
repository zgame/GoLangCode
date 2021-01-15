package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("-------------说明------------------")
	fmt.Println("把所有要转换的xlsx文件扔到app同路径下，运行即可。")
	fmt.Println("-------------xlsx格式---------------")
	fmt.Println("第一行是列名字")
	fmt.Println("第二行是列类型，支持int,string,bool,float,double")
	fmt.Println("-------------END------------------")

	err := RunCurrentPahtAllFile()
	if err != nil {
		fmt.Println("读当前目录下文件出错")
		fmt.Println(err.Error())
	}
	fmt.Println("--------------------end----------------------")
}

// 读取csv文件内容
func readSample(filename string) *xlsx.Sheet {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Println("读取文件出错" + filename)
	}

	//StrOut:=  fmt.Sprintf("module('%s')\n\n %s = {\n" ,filename,filename)			// 开头

	//for s, sheet := range xlFile.Sheets {
	//	if s == 0 {  // 只读第一个页签
	//for i, row := range sheet.Rows {
	//	if i < 4 {
	//		continue // 略过前4行
	//	}
	//
	//	line := "[\"" + row.Cells[0].String() + "\"] = { " // 这里写入每行的key， 要注意， 这里需要用字符串， 因为用字符串就是key， 不用就是数组，必须用key
	//	for j := 1; j < len(row.Cells); j++ {
	//		str := getString(i, j, row.Cells)
	//		line = line + str
	//	}
	//	line = line + " },\n"
	//	//fmt.Println("", line)
	//	StrOut += line
	//
	//	//for k, cell := range row.Cells {
	//	//	text := cell.String()
	//	//	fmt.Printf("%s\n", text)
	//	//}
	//}
	//	}
	//}

	//f, err := os.Open(filename)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//rows, err := csv.NewReader(f).ReadAll()
	//f.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}

	return xlFile.Sheets[0]
}

// 把新数据结构写入文件中
func writeLuaFiles(filename string, data string) {
	f, err := os.Create(filename + ".lua")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(data)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// 获取当前路径下所有文件
func RunCurrentPahtAllFile() error {
	pathname := getCurrentDirectory()
	//pathname="C:/Users/Administrator/Documents/GitHub/GoLangCode/csvTolua/"
	pathname = "./"

	rd, err := ioutil.ReadDir(pathname)
	// 删除老的文件
	for _, fi := range rd {
		if !fi.IsDir() {
			files := strings.Split(fi.Name(), ".")
			fileType := files[1]
			if fileType == "lua" {
				err = os.Remove("./" + fi.Name())
				if err == nil {
					fmt.Println("删除老的文件", fi.Name())
				}
			}
		}
	}
	//  开始转换xlsx文件
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			//RunCurrentPahtAllFile(pathname + fi.Name() + "\\")
		} else {
			files := strings.Split(fi.Name(), ".")
			fileName := files[0]
			fileType := files[1]
			if fileType == "xlsx" {
				fmt.Println("开始转换", fi.Name())
				sheet := readSample(fi.Name())
				//str:= readSample(fi.Name())

				str := DealRowsData(fileName, sheet.Rows)
				writeLuaFiles(fileName, str)
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

// 处理每行每列中的数据
func getString(RowIndex int, ListIndex int, rows []*xlsx.Row) string {
	var strOut string
	ListName := rows[0].Cells[ListIndex].String() // 列名
	ListType := rows[1].Cells[ListIndex].String() // 列类型
	ListType = strings.ToLower(ListType)
	// 增加一个判断,如果是#就忽略
	if ListName[0] == '#' {
		return ""
	}

	data := rows[RowIndex].Cells[ListIndex].String()
	if rows[RowIndex].Cells[0].String() == "" {
		fmt.Println("报错了， excel格式不正确， 包含了空的行")
	}

	if ListType == "string" {
		strOut = fmt.Sprintf(" %s = \"%s\", ", ListName, data)
	} else if ListType == "int" {
		if data == "" {
			data = "0"
		}
		strOut = fmt.Sprintf(" %s = %s, ", ListName, data)
	} else if ListType == "float" {
		if data == "" {
			data = "0"
		}
		strOut = fmt.Sprintf(" %s = %s, ", ListName, data)
	} else if ListType == "double" {
		if data == "" {
			data = "0"
		}
		strOut = fmt.Sprintf(" %s = %s, ", ListName, data)
	} else if ListType == "bool" {
		data = strings.ToUpper(data)
		if data == "1" || data == "TRUE" {
			data = "true"
		} else {
			data = "false"
		}
		strOut = fmt.Sprintf(" %s = %s, ", ListName, data)
	}else if ListType == "array" {
		strOut = fmt.Sprintf(" %s = {%s}, ", ListName, data)
	}
	return strOut
}

// 处理row数据
func DealRowsData(fileName string, rows []*xlsx.Row) string {
	fmt.Println("-----------------------开始处理数据--------------------------")

	width := len(rows[0].Cells)
	//fmt.Println("width:",width)
	StrOut := fmt.Sprintf("local %s = {\n", fileName) // 开头
	//fmt.Println("",StrOut)

	//ListName := rows[0]		// 列名
	//ListType := rows[1]		// 列类型

	for i := range rows {
		if i < 4 {
			continue // 略过前2行
		}
		fmt.Println("开始处理第 " + strconv.Itoa(i) + " 行")

		line := "[\"" + rows[i].Cells[0].String() + "\"] = { " // 这里写入每行的key， 要注意， 这里需要用字符串， 因为用字符串就是key， 不用就是数组，必须用key
		for j := 1; j < width; j++ {
			str := getString(i, j, rows)
			line = line + str
		}
		line = line + " },\n"
		//fmt.Println("", line)
		StrOut += line
	}

	StrOut += "} \n " // 结尾

	// 增加处理函数
	StrOut += fmt.Sprintf(`CSV_%s = {}

function CSV_%s.GetValue(index, key)
	index = tostring(index)
	key = tostring(key)
    if %s[index] == nil then
        ZLog.Logger("Excel 获取表:%s  主键:".. index .." key:".. key.."出错!")
        return nil
    end
    if %s[index][key] == nil then
        ZLog.Logger("Excel 获取表: %s  主键:".. index .." key:".. key.."出错!")
        return nil
    end

    return %s[index][key]
end


function CSV_%s.Get()
    return %s
end

function CSV_%s.GetAllKeys()
    local keys = {}
    for k in pairs(%s) do
        keys[#keys + 1] = k
    end
    if #keys == 0 then
        ZLog.Logger(string.format("Excel 获取表: 所有主键出错"))
        return nil
    end
    return keys
end`, fileName, fileName, fileName, fileName, fileName, fileName, fileName,fileName,fileName,fileName,fileName)

	//fmt.Println("",StrOut)
	return StrOut
}
