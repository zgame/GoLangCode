//------------------------------------------------------------------------------------------
// mysql  的数据库查询
//------------------------------------------------------------------------------------------
package mysql

import (
	"reflect"
	"github.com/junhsieh/goexamples/fieldbinding/fieldbinding"
	util "github.com/tengattack/gluasql/util"
	"github.com/yuin/gopher-lua"
	"fmt"
)

func clientQueryMethod(L *lua.LState) int {
	client := checkClient(L)
	query := L.ToString(2)

	if client.DB == nil {
		client.DB = GlobalDB
		//fmt.Println("client.DB == nil  !!!!!!!!!!")
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

		tbRow := util.ToTableFromMap(L, reflect.ValueOf(fb.GetFieldArr()))
		tb.Append(tbRow)
	}

	L.Push(tb)
	return 1
}


func clientExecMethod(L *lua.LState) int {
	client := checkClient(L)
	exec := L.ToString(2)

	if client.DB == nil {
		client.DB = GlobalDB
		//fmt.Println("client.DB == nil  !!!!!!!!!!")
		//return 0
	}

	if exec == "" {
		L.ArgError(2, "exec string required")
		return 0
	}

	result, err := client.DB.Exec(exec)
	if err != nil {
		fmt.Println("mysql 数据库 exec error",err.Error())
		fmt.Println("result",result)
		return 0
	}
	return 0
}
