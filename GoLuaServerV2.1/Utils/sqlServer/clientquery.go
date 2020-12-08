//------------------------------------------------------------------------------------------
// sql server 的数据库查询
//------------------------------------------------------------------------------------------
package sqlServer

import (
	"GoLuaServerV2.1/Utils"
	"GoLuaServerV2.1/Utils/zLog"
	"fmt"
	"github.com/yuin/gopher-lua"
	"reflect"
	"time"
)

// 查询
func clientQueryMethod(L *lua.LState) int {
	client := checkClient(L)
	query := L.ToString(2)

	//fmt.Println("sql ",query)

	if client.DB == nil {
		zLog.PrintLogger("SQL server  client.DB == nil  !!!!!!!!!!")
		return 0
	}

	if query == "" {
		L.ArgError(2, "query string required")
		return 0
	}

	rows, err := client.DB.Query(query)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	defer rows.Close()

	tb := L.NewTable()
	cols, err := rows.Columns()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	//fmt.Println("rows.Columns() " ,cols)
	//
	//fb := fieldbinding.NewFieldBinding()
	//fb.PutFields(cols)
	//fmt.Println("*********start*********")

	//for rows.Next() {
	//	if err := rows.Scan(fb.GetFieldPtrArr()...); err != nil {
	//		L.Push(lua.LNil)
	//		L.Push(lua.LString(err.Error()))
	//		return 2
	//	}
	//	fmt.Println("------------------------------------")
	//	//printValue(fb.GetFieldArr())
	//	fmt.Println("",fb.GetFieldArr())
	//	tbRow := util.toTableFromMap(L, reflect.ValueOf(fb.GetFieldArr()))
	//	fmt.Println("---------------end---------------------")
	//	tb.Append(tbRow)
	//}

	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println("error when rows scan  -  ", err.Error())
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		tbRow := &lua.LTable{}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("\t")
			}
			//printValue(vals[i].(*interface{}))
			tbRow.RawSet(Utils.LuaSetValue(L, cols[i]), Utils.LuaSetValue(L, vals[i]))
		}
		//fmt.Println("")
		tb.Append(tbRow)
	}

	L.Push(tb)
	return 1
}

// 不带返回的执行， 改成异步执行
func clientExecMethod(L *lua.LState) int {

	client := checkClient(L)
	exec := L.ToString(2)

	if client.DB == nil {
		zLog.PrintLogger("SQL server  client.DB == nil  !!!!!!!!!!")
		L.Push(lua.LString("SQL server  client.DB == nil"))
		return 1
	}

	if exec == "" {
		L.ArgError(2, "exec string required")
		//return 0
	}
	go func() {
		_, err := client.DB.Exec(exec)
		if err != nil {
			zLog.PrintLogger("sql server 数据库 exec error" + err.Error())
			//L.Push(lua.LString("SQL server  error:"+err.Error()))
			//return 1
		}
	}()
	return 0 // 执行成功
}

// 用来调试打印的
func printValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("true")
		} else {
			fmt.Print("false")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
	fmt.Print("\t", reflect.TypeOf(*pval))

}
