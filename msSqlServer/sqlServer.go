package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	var (
		userid   = "sa"
		password = "Gh5b645xSGBnPl"
		server   = "1---------1----4.5-------5.1---------3---------6.7--------5"
		database = "DataBase----------BY"
	)
	//flag.Parse()

	dsn := "server=" + server + ";user id=" + userid + ";password=" + password + ";database=" + database
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}

	//cmd := " INSERT into dbo.AaWhiteIPList(ip,comments) VALUES('11','222')"
	//cmd := " UPDATE AaWhiteIPList set ip = '22' where ip = '11'"
	cmd := " select * from AaWhiteIPList"
	//cmd := " delete  from AaWhiteIPList"

	err,rows := exec(db, cmd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("rows",rows)
	//for rows.Next() {
	//	err = rows.Scan(vals...)
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//	for i := 0; i < len(vals); i++ {
	//		if i != 0 {
	//			fmt.Print("\t")
	//		}
	//		printValue(vals[i].(*interface{}))
	//	}
	//	fmt.Println()
	//
	//}

}

func exec(db *sql.DB, cmd string) (error ,*sql.Rows){
	rows, err := db.Query(cmd)
	if err != nil {
		return err,nil
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return err,nil
	}
	if cols == nil {
		return nil,nil
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("\t")
		}
		//fmt.Print(cols[i])		// 这里打印列名
	}
	fmt.Println()
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("\t")
			}
			printValue(vals[i].(*interface{}))
		}
		fmt.Println()

	}
	if rows.Err() != nil {
		return rows.Err(),rows
	}
	return nil,rows
}

func printValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
}
