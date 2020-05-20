package main

import (
	"./mssql"
	"./zLog"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"strings"
)

func DealUserList(idStart int) {
	var (
		userId     = "dbuser"
		password   = "CEDFE2CDA7DB84AC"
		server     = "172.16.140.89"
		logDBName1 = "BY_LOG_202001"
		logDBName2 = "BY_LOG_202002"
		testDBName = "testdb"
	)
	fmt.Println(" --------------开始连接数据库-------------- ")
	testDB := mssql.ConnectDB(userId, password, server, testDBName)
	logDB1 := mssql.ConnectDB(userId, password, server, logDBName1)
	logDB2 := mssql.ConnectDB(userId, password, server, logDBName2)

	//fmt.Println(" --------------开始查询玩家列表--------------")
	sqlU := fmt.Sprintf("select top(%d)id, UserID, TheDays  from testdb.dbo.a1_user_free_new_sortid_match   with(nolock) where id >= %d and UserID_2 is NULL", Group, idStart)
	_, rows, _ := mssql.Query(testDB, sqlU)

	for rows.Next() { // 循环遍历
		var userInfo UserList
		err := rows.Scan(&userInfo.id, &userInfo.uid, &userInfo.days) // 赋值到结构体中
		if err != nil {
			zLog.PrintfLogger(" 遍历玩家列表 id %d    , %s \n", userInfo.id, err)
			continue
		}

		zLog.PrintfLogger(" --------------开始处理id : %d--------------", userInfo.id)

		tableColumns := fmt.Sprintf("select top 1 UserID, InitLogonDate, LastLogonDate from testdb.dbo.a0_user_RecordLogon_lost_lx_sortid with(nolock) where TheDays=%d order by newid()", userInfo.days)
		zLog.PrintfLogger("tableColumns: %s ", tableColumns)
		_, rowsGetColumns, _ := mssql.Query(testDB, tableColumns)
		var UserID_2 int
		var InitLogonDate_2 string
		var LastLogonDate_2 string
		for rowsGetColumns.Next() { // 循环遍历
			err := rowsGetColumns.Scan(&UserID_2, &InitLogonDate_2, &LastLogonDate_2)
			if err != nil {
				fmt.Printf("  select id：%d     %s \n ", userInfo.id, err.Error())
				fmt.Println("----------------------------------------")
				fmt.Println("", tableColumns)
				fmt.Println("----------------------------------------")
			}
			//fmt.Println("", allKeys)
		}
		mssql.CloseQuery(rowsGetColumns)
		// ----------------------------开始执行update------------------------------
		//allKeys = GetTableKeys(RecordTimeDict[j])

		// 每个不同的处理方式
		//allKeysDeal := GetTableKeysDeal(RecordTimeDict[j], userInfo)

		// 统一的insert语句
		UpdateSql := fmt.Sprintf("update testdb.dbo.a1_user_free_new_sortid_match with(rowlock,updlock) set UserID_2=%d, InitLogonDate_2='%s', LastLogonDate_2='%s', TheDays_2=%d, MatchType=1, daynum=datediff(DAY,'%s',InitLogonDate) where id=%d", UserID_2, InitLogonDate_2[0:10], LastLogonDate_2[0:10], userInfo.days, InitLogonDate_2[0:10],  userInfo.id)
		//selectSql := fmt.Sprintf(" select  %s  from  %s.dbo.%s  WITH(NOLOCK)  where UserID= %d", allKeysDeal, dbName2, table2, userInfo.uid2)
		//sqlString := UpdateSql + selectSql
		zLog.PrintfLogger("sql: %s ", UpdateSql)
		err, _ = mssql.Exec(testDB, UpdateSql)
		if err != nil {
			zLog.PrintfLogger("update Exec Error %s ,sql: %s", err.Error(),UpdateSql)
		}
		//}
		//}

	}
	mssql.CloseQuery(rows)
	mssql.CloseDB(testDB)
	mssql.CloseDB(logDB1)
	mssql.CloseDB(logDB2)

	wg.Done()
}

// 获取当前的月份
func GetMonth(table1 string, logDB1 *sql.DB, logDB2 *sql.DB) (*sql.DB, string) {
	var dbNow *sql.DB
	var dbName string
	if strings.Contains(table1, "202001") {
		dbNow = logDB1
		dbName = "BY_LOG_202001"
	} else {
		dbNow = logDB2
		dbName = "BY_LOG_202002"
	}
	return dbNow, dbName
}
