package main

import (
	"./mssql"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"strconv"
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
	sqlU:= fmt.Sprintf( "select top(%d)* from testdb.dbo.a1_user_chongzhi_new_match_2   with(nolock) where id >= %d", Group,idStart)
	_, rows, _ := mssql.Query(testDB, sqlU)

	for rows.Next() { // 循环遍历
		var userInfo UserList
		err := rows.Scan(&userInfo.id, &userInfo.uid, &userInfo.initDate, &userInfo.lastDate, &userInfo.days, &userInfo.uid2, &userInfo.initDate2, &userInfo.lastDate2, &userInfo.days2, &userInfo.matchType, &userInfo.dayNum) // 赋值到结构体中
		if err != nil {
			fmt.Printf(" 遍历玩家列表 id %d    , %s \n" , userInfo.id,  err)
			continue
		}
		fmt.Println(" --------------开始处理id--------------", userInfo.id)

		// -----------------------------获取一行数据------------------------
		//fmt.Println("", userInfo.id)
		//fmt.Println("", userInfo.uid)
		//fmt.Println("", userInfo.initDate)
		//fmt.Println("", userInfo.lastDate)
		//fmt.Println("", userInfo.days)
		//fmt.Println("", userInfo.uid2)
		//fmt.Println("", userInfo.initDate2)
		//fmt.Println("", userInfo.lastDate2)
		//fmt.Println("", userInfo.days2)
		//fmt.Println("", userInfo.matchType)
		//fmt.Println("", userInfo.dayNum)

		dayList := getTimeList(userInfo.initDate, userInfo.days) // 玩家的日期列表
		dayList2 := getTimeList(userInfo.initDate2, userInfo.days2)
		//fmt.Println("day list: ", dayList[0])
		//fmt.Printf("%v \n ", dayList2)

		// -----------------------------用户所有天数--------------------------------
		for i := 0; i < userInfo.days; i++ { // 按照日期遍历
			day1 := dayList[i]
			day2 := dayList2[i]
			day1 = strings.Replace(day1, "-", "", -1)
			day2 = strings.Replace(day2, "-", "", -1)

			// -----------------------------用户所有表格--------------------------------
			for j := range RecordTimeDict { // 按照表遍历
				table1 := RecordTimeDict[j] + day1
				table2 := RecordTimeDict[j] + day2

				//fmt.Println("day1",day1)
				//fmt.Println("table1",table1)
				//fmt.Println("table2",table2)
				//var dbNow,dbNow2 *sql.DB
				//var dbName string
				dbNow, dbName := GetMonth(table1,  logDB1,  logDB2)
				dbNow2, dbName2 := GetMonth(table2,  logDB1,  logDB2)
				//fmt.Println("",dbNow)

				// --------------------------这里各个表的所有列名-----------------------
				tableColumns := fmt.Sprintf(`
				use %s
				select stuff((
					select
						',' + c.name
						from sys.tables t with(nolock)
						left join sys.columns c with(nolock) on t.object_id=c.object_id
						where t.object_id=OBJECT_ID('%s')
						order by c.column_id asc
						for xml path('')
						),1,1,'')  as columns_list;
						`, dbName2, table2)

				//fmt.Println(" table:"  ,tableColumns)
				_, rowsGetColumns, _ := mssql.Query(dbNow2, tableColumns)
				var resultGetColumns string
				for rowsGetColumns.Next() { // 循环遍历
					err := rowsGetColumns.Scan(&resultGetColumns)
					if err != nil {
						fmt.Printf("  获取列名 id：%d     %s \n ", userInfo.id, err.Error())
						fmt.Println("----------------------------------------")
						fmt.Println("", tableColumns)
						fmt.Println("----------------------------------------")
					}
					//fmt.Println("", resultGetColumns)
				}
				mssql.CloseQuery(rowsGetColumns)
				// ----------------------------------------------------------
				strTmp := strings.Replace(resultGetColumns, "UserID", strconv.Itoa(userInfo.uid), -1)
				tableRes1 := strings.Replace(strTmp, "RecordTime", fmt.Sprintf("dateadd(day,%d,RecordTime) as RecordTime", userInfo.dayNum), -1)
				insertSql := fmt.Sprintf("insert into %s (%s) ", dbName+".dbo."+table1 , resultGetColumns)
				sql :=  insertSql + " select " + tableRes1 + " from " + dbName2 + ".dbo." + table2 + " where UserID=" + strconv.Itoa(userInfo.uid2)

				//fmt.Println("sql:",sql)
				mssql.Exec(dbNow, sql)
			}
		}

	}
	mssql.CloseQuery(rows)
	mssql.CloseDB(testDB)
	mssql.CloseDB(logDB1)
	mssql.CloseDB(logDB2)

	wg.Done()
}

// 获取当前的月份
func GetMonth(table1 string,  logDB1 *sql.DB, logDB2 *sql.DB) (*sql.DB, string) {
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
