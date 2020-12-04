//------------------------------------------------------------------------------------------
// MongoDB 的数据库查询
//------------------------------------------------------------------------------------------
package mongoDB

import (
	"GoLuaServerV2.1/Utils/zLog"
	"fmt"
	"github.com/tengattack/gluasql/util"
	"github.com/yuin/gopher-lua"
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
	"time"
)

// 插入数据
func clientInsertMethod(L *lua.LState) int {
	client := checkClient(L)
	collection := L.ToString(2)

	s := client.MongoSession.Copy()
	defer s.Close()
	Collection := s.DB("").C(collection)
	//fmt.Println("----------------insert--------------")

	tb := gluasql_util.GetValue(L, 3)
	options, ok := tb.(map[string]interface{})
	if tb == nil || !ok {
		L.ArgError(3, "options excepted")
		zLog.PrintLogger("mongo db  insert 参数错误")
		//L.Push(lua.LString("options excepted"))
		return 0

	}

	err := Collection.Insert(options)
	if err != nil {
		zLog.PrintLogger("mongo db insert error "+err.Error())
		L.Push(lua.LString(err.Error()))
		return 1
	}

	return 0 // 执行成功
}

// 删除数据
func clientDelMethod(L *lua.LState) int {
	client := checkClient(L)
	collection := L.ToString(2)

	s := client.MongoSession.Copy()
	defer s.Close()
	Collection := s.DB("").C(collection)
	//fmt.Println("----------------Del--------------")

	tb := gluasql_util.GetValue(L, 3)
	options, ok := tb.(map[string]interface{})
	if tb == nil || !ok {
		L.ArgError(3, "options excepted")
		zLog.PrintLogger("mongo db  del 参数错误")
		return 0
	}

	err := Collection.Remove(options)
	if err != nil {
		zLog.PrintLogger("mongo db del error "+err.Error())
		L.Push(lua.LString(err.Error()))
		return 1
	}

	return 0 // 执行成功
}

// 更新数据， 只更新options里面的部分数据
func clientUpdateMethod(L *lua.LState) int {
	client := checkClient(L)
	collection := L.ToString(2)

	s := client.MongoSession.Copy()
	defer s.Close()
	Collection := s.DB("").C(collection)
	//fmt.Println("----------------update--------------")

	tb := gluasql_util.GetValue(L, 3)
	options, ok := tb.(map[string]interface{})
	if tb == nil || !ok {
		L.ArgError(3, "options excepted")
		zLog.PrintLogger("mongo db update参数错误")
		return 0
	}
	utb := gluasql_util.GetValue(L, 4)
	updateO, ok := utb.(map[string]interface{})
	if utb == nil || !ok {
		L.ArgError(4, "options excepted")
		zLog.PrintLogger("mongo db update参数错误")
		return 0
	}

	err := Collection.Update(options, bson.M{"$set": updateO})
	if err != nil {
		zLog.PrintLogger("mongo db update error "+err.Error())
		L.Push(lua.LString(err.Error()))
		return 1
	}

	return 0 // 执行成功
}

// 查询数据
func clientFindMethod(L *lua.LState) int {
	client := checkClient(L)
	collection := L.ToString(2)

	s := client.MongoSession.Copy()
	defer s.Close()
	Collection := s.DB("").C(collection)
	//fmt.Println("----------------find--------------")

	tb := gluasql_util.GetValue(L, 3)
	findTb, ok := tb.(map[string]interface{})
	if tb == nil || !ok {
		L.ArgError(3, "options excepted")
		return 0
	}

	result := make(map[string]interface{}, 0)
	err := Collection.Find(findTb).One(&result)
	//err := Collection.Find(  findTb, bson.M{"title":1,"_id":0}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	result["_id"] = nil
	//fmt.Println(result)

	returnTb := gluasql_util.ToTableFromMap(L, reflect.ValueOf(result))
	L.Push(returnTb)

	return 1 // 执行成功
}

// 查询数据
func clientFindsMethod(L *lua.LState) int {
	client := checkClient(L)
	collection := L.ToString(2)
	sort := L.ToString(4)

	s := client.MongoSession.Copy()
	defer s.Close()
	Collection := s.DB("").C(collection)
	//fmt.Println("----------------finds--------------")

	tb := gluasql_util.GetValue(L, 3)
	findTb, ok := tb.(map[string]interface{})
	if tb == nil || !ok {
		L.ArgError(3, "options excepted")
		return 0
	}

	result := make([]map[string]interface{}, 0)
	err := Collection.Find(findTb).Sort(sort).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	returnTb := ToTableFromSlice(L, reflect.ValueOf(result))
	L.Push(returnTb)

	return 1 // 执行成功
}

//------------------------------------------------------------------------------------
// utils
//------------------------------------------------------------------------------------

// map to  lua table
func ToTableFromMap(l *lua.LState, v reflect.Value) lua.LValue {
	tb := &lua.LTable{}
	for _, k := range v.MapKeys(){
		key := ToArbitraryValue(l, k.Interface())
		if key!= lua.LString("_id") {
			tb.RawSet(key, ToArbitraryValue(l, v.MapIndex(k).Interface()))
		}
	}
	return tb
}
// slice to lua table
func ToTableFromSlice(l *lua.LState, v reflect.Value) lua.LValue {
	tb := &lua.LTable{}
	for j := 0; j < v.Len(); j++ {
		tb.RawSet(ToArbitraryValue(l, j+1), // because lua is 1-indexed
			ToArbitraryValue(l, v.Index(j).Interface()))
	}
	return tb
}

func ToArbitraryValue(l *lua.LState, i interface{}) lua.LValue {
	if i == nil {
		return lua.LNil
	}

	switch ii := i.(type) {
	case bool:
		return lua.LBool(ii)
	case int:
		return lua.LNumber(ii)
	case int8:
		return lua.LNumber(ii)
	case int16:
		return lua.LNumber(ii)
	case int32:
		return lua.LNumber(ii)
	case int64:
		return lua.LNumber(ii)
	case uint:
		return lua.LNumber(ii)
	case uint8:
		return lua.LNumber(ii)
	case uint16:
		return lua.LNumber(ii)
	case uint32:
		return lua.LNumber(ii)
	case uint64:
		return lua.LNumber(ii)
	case float64:
		return lua.LNumber(ii)
	case float32:
		return lua.LNumber(ii)
	case string:
		return lua.LString(ii)
	case []byte:
		return lua.LString(ii)
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Ptr:
			return ToArbitraryValue(l, v.Elem().Interface())

		//case reflect.Struct:
		//	return ToTableFromStruct(l, v)

		case reflect.Map:
			return ToTableFromMap(l, v)

		case reflect.Slice:
			return ToTableFromSlice(l, v)

		//case mgo.Index{}

		default:
			return lua.LString("")
		}
	}
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
