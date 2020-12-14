//------------------------------------------------------------------------------------------
// MongoDB 的数据库查询
//------------------------------------------------------------------------------------------
package mongoDB

import (
	"GoLuaServerV2.1/Utils/zLog"
	"GoLuaServerV2.1/Utils/zLua"
	"github.com/yuin/gopher-lua"
	"gopkg.in/mgo.v2/bson"
)

// 插入数据
func clientInsertMethod(L *lua.LState) int {
	client := checkClient(L)
	collection := L.ToString(2)

	s := client.MongoSession.Copy()
	defer s.Close()
	Collection := s.DB("").C(collection)
	//fmt.Println("----------------insert--------------")

	tb := zLua.LuaGetValue(L, 3)
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

	tb := zLua.LuaGetValue(L, 3)
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

	tb := zLua.LuaGetValue(L, 3)
	options, ok := tb.(map[string]interface{})
	if tb == nil || !ok {
		L.ArgError(3, "options excepted")
		zLog.PrintLogger("mongo db update参数错误")
		return 0
	}
	utb := zLua.LuaGetValue(L, 4)
	updateO, ok := utb.(map[string]interface{})
	if utb == nil || !ok {
		L.ArgError(4, "options excepted")
		zLog.PrintLogger("mongo db update参数错误")
		return 0
	}

	cmd := L.ToString(5)

	err := Collection.Update(options, bson.M{cmd: updateO})
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

	tb := zLua.LuaGetValue(L, 3)
	findTb, ok := tb.(map[string]interface{})
	if tb == nil || !ok {
		L.ArgError(3, "options excepted")
		return 0
	}

	result := make(map[string]interface{}, 0)
	err := Collection.Find(findTb).Select(bson.M{"_id":0}).One(&result)
	//err := Collection.Find(  findTb, bson.M{"title":1,"_id":0}).One(&result)
	if err != nil {
		return 0
	}
	//result["_id"] = nil
	//fmt.Println(result)

	//returnTb := gluasql_util.toTableFromMap(L, reflect.ValueOf(result))
	returnTb := zLua.LuaSetValue(L, result)
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

	tb := zLua.LuaGetValue(L, 3)
	findTb, ok := tb.(map[string]interface{})
	if tb == nil || !ok {
		L.ArgError(3, "options excepted")
		return 0
	}

	result := make([]map[string]interface{}, 0)
	err := Collection.Find(findTb).Select(bson.M{"_id":0}).Sort(sort).All(&result)
	if err != nil {
		return 0
	}

	//returnTb := toTableFromSlice(L, reflect.ValueOf(result))
	returnTb := zLua.LuaSetValue(L, result)
	L.Push(returnTb)

	return 1 // 执行成功
}
