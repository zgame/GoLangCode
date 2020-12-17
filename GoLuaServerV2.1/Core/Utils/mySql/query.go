//------------------------------------------------------------------------------------------
// mySql  的数据库查询
//------------------------------------------------------------------------------------------
package mySql

import (
	"GoLuaServerV2.1/Core/Utils/zLog"
	"GoLuaServerV2.1/Core/Utils/zLua"
	"fmt"
	"github.com/junhsieh/goexamples/fieldbinding/fieldbinding"
	"github.com/yuin/gopher-lua"
)

func clientQueryMethod(L *lua.LState) int {
	client := checkClient(L)
	query := L.ToString(2)

	if client.DB == nil {
		zLog.PrintLogger("mysql client.DB == nil  !!!!!!!!!!")
		//return 0
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

	fb := fieldbinding.NewFieldBinding()
	cols, err := rows.Columns()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	fb.PutFields(cols)

	tb := L.NewTable()
	for rows.Next() {
		if err := rows.Scan(fb.GetFieldPtrArr()...); err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		//tbRow := util.toTableFromMap(L, reflect.ValueOf(fb.GetFieldArr()))
		tbRow := zLua.LuaSetValue(L, fb.GetFieldArr())
		tb.Append(tbRow)
	}

	L.Push(tb)
	return 1
}


func clientExecMethod(L *lua.LState) int {
	client := checkClient(L)
	exec := L.ToString(2)

	if client.DB == nil {
		zLog.PrintLogger("mysql client.DB == nil  !!!!!!!!!!")
		L.Push(lua.LString("mysql  client.DB == nil"))
		return 1
		//return 0
	}

	if exec == "" {
		L.ArgError(2, "exec string required")
		return 0
	}

	go func() {
		_, err := client.DB.Exec(exec)
		if err != nil {
			fmt.Println("mySql 数据库 exec error", err.Error())
			//fmt.Println("result",result)
			//return 0
		}
	}()
	return 0
}
