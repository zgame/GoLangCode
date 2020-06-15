package mssql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

// 连接数据库
func ConnectDB(userId string,password string,server string,database string) *sql.DB{

	//flag.Parse()

	dsn := "server=" + server + ";user id=" + userId + ";password=" + password + ";database=" + database
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return nil
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(200)
	db.SetConnMaxLifetime(   600 * time.Second )

	//defer db.CloseDB()
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return nil
	}

	//cmd := " INSERT into dbo.AaWhiteIPList(ip,comments) VALUES('11','222')"
	//cmd := " UPDATE AaWhiteIPList set ip = '22' where ip = '11'"
	//cmd := " select * from AaWhiteIPList"
	//cmd := " delete  from AaWhiteIPList"

	//err,rows := Exec(db, cmd)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("rows",rows)
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
	fmt.Println(database, " connect success")
	return db
}

// 关闭数据库
func CloseDB(db *sql.DB)  {
	db.Close()
}

// 执行
func Exec(db *sql.DB, cmd string) (error , sql.Result) {
	result, err := db.Exec(cmd)
	if err != nil {
		fmt.Println("Exec Error: ", err.Error())
		return err, nil
	}
	return err, result
}

// 查询
func Query(db *sql.DB, cmd string) (error ,*sql.Rows, int){
	rows, err := db.Query(cmd)
	if err != nil {
		fmt.Println("Query Error", err.Error())
		return err,nil,0
	}
	//defer rows.CloseDB()
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Query Error", err.Error())
		return err,nil,0
	}
	if cols == nil {
		fmt.Println("Query Error", err.Error())
		return nil,nil,0
	}
	//vals := make([]interface{}, len(cols))
	//for i := 0; i < len(cols); i++ {
	//	vals[i] = new(interface{})
	//	if i != 0 {
	//		fmt.Print("\t")
	//	}
	//	//fmt.Print(cols[i])		// 这里打印列名
	//}
	//fmt.Println()
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
	if rows.Err() != nil {
		return rows.Err(),rows,0
	}
	return nil,rows,len(cols)
}

// 关闭查询
func CloseQuery(rows *sql.Rows)  {
	rows.Close()
}

// 打印调试
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

// 返回值
func GetValue(pval *interface{}) interface{}{
	switch v := (*pval).(type) {
	case nil:
		return nil
	case bool:
		if v {
			return true
		} else {
			return false
		}
	case []byte:
		return string(v)
	case time.Time:
		return v.Format("2006-01-02 15:04:05.999")
	default:
		return v
	}
}
