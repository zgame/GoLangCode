package main

import (
	"github.com/yuin/gopher-lua"
	"strconv"
	"fmt"
)

const LuaInt64Max  = 18446744073709551615

func luaopen_pb(L *lua.LState) int {
	//mt:=L.NewTypeMetatable("zzz")

	//L.Push(lua.LNumber(-1))
	//L.SetField(mt, "__index", L.Get(lua.GlobalsIndex))
	//L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), personMethods))

	L.Register("__tostring", iostring_str)
	L.Register("__len", iostring_len)
	L.Register("write", iostring_write)
	L.Register("sub", iostring_sub)
	L.Register("clear", iostring_clear)

	//lua_setfield(L, -2, "__index");
	//luaL_register(L, NULL, _c_iostring_m);
	//luaL_register(L, "pb", _pb);

	//L.Register("", _c_iostring_m)
	//L.Register("pb", _pb)


	//mt:= L.NewTypeMetatable("zsw")
	//L.SetGlobal("pb", mt)							// 设定全局mudule
	////L.SetField(mt, "new", L.NewFunction(zprint))		// 绑定new函数
	//
	////mt.RawSetString("__index",mt)						// 设定__index
	//L.SetFuncs(mt, _pb)								// 设定metaTable的函数列表

	L.PreloadModule("pb", pbLoader)
	return 1
}

// 加载自己的module pb
func pbLoader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), _pb)
	// register other stuff
	L.SetField(mod, "name", lua.LString("value"))
	// returns the module
	L.Push(mod)
	return 1
}


var _pb = map[string]lua.LGFunction{
	//static const struct luaL_Reg _pb[] = {
	"varint_encoder": varint_encoder,
	"signed_varint_encoder": signed_varint_encoder,
	"read_tag": read_tag,
	"struct_pack": struct_pack,
	"struct_unpack": struct_unpack,
	"varint_decoder": varint_decoder,
	"signed_varint_decoder": signed_varint_decoder,
	"zig_zag_decode32": zig_zag_decode32,
	"zig_zag_encode32": zig_zag_encode32,
	"zig_zag_decode64": zig_zag_decode64,
	"zig_zag_encode64": zig_zag_encode64,
	"new_iostring": iostring_new,
	"ZswLuaShowBytesToString":ZswLuaShowBytesToString,
}




//
//type  IOString struct{
//	size uint64
//	char buf[IOSTRING_BUF_LEN]
//}
//



//
//
//typedef struct{
//size_t size;
//char buf[IOSTRING_BUF_LEN];
//} IOString;
//

//-----------------------------------------int encode-------------------------------------
func luaL_addchar(b []byte, c uint64)  []byte{
	//str := strconv.FormatInt(int64(c),10)
	//str := string(byte(c))
	return append(b,byte(c))
}


func pack_varint(str string, value uint64 ) string{
	b := []byte(str)

	if value >= 0x80 {
		b = luaL_addchar(b, value|0x80)
		value >>= 7
		if value >= 0x80 {
			b = luaL_addchar(b, value|0x80)
			value >>= 7
			if value >= 0x80 {
				b = luaL_addchar(b, value|0x80)
				value >>= 7
				if value >= 0x80 {
					b = luaL_addchar(b, value|0x80)
					value >>= 7
					if value >= 0x80 {
						b = luaL_addchar(b, value|0x80)
						value >>= 7
						if value >= 0x80 {
							b = luaL_addchar(b, value|0x80)
							value >>= 7
							if value >= 0x80 {
								b = luaL_addchar(b, value|0x80)
								value >>= 7
								if value >= 0x80 {
									b = luaL_addchar(b, value|0x80)
									value >>= 7
									if value >= 0x80 {
										b = luaL_addchar(b, value|0x80)
										value >>= 7
									}
								}
							}
						}
					}
				}
			}
		}
	}
	b = luaL_addchar(b, value)
	return string(b)
}

func varint_encoder(L *lua.LState) int {
	//lua_Number 	l_value = luaL_checknumber(L, 2);
	//uint64_t 	value = (uint64_t)	l_value
	//luaL_Buffer b;
	//luaL_buffinit(L, &b);
	//pack_varint(b, value);
	//lua_settop(L, 1);
	//luaL_pushresult(&b);
	//lua_call(L, 1, 0);


	println("pb.go   ------------       varint_encoder:")
	l_value := L.ToNumber(2)
	value := uint64(l_value)
	b := pack_varint("", value)			// 把数字变成string

	//L.Push(lua.LString(b))
	//L.Call(1,0)			// string作为参数，调用参数1作为函数名
	l_func := L.ToFunction(1)
	if err := L.CallByParam(lua.P{
		Fn:      l_func,
		NRet:    0,
		Protect: true,
	}, lua.LString(b)); err != nil {
		println("signed_varint_encoder error:", err.Error())
	}
	return 0
}

func signed_varint_encoder(L *lua.LState) int {
	//lua_Number	l_value = luaL_checknumber(L, 2);
	//int64_t	value = (int64_t)	l_value;
	//luaL_Buffer	b;
	//luaL_buffinit(L, &b);
	//if (value < 0) {
	//	pack_varint(&b, *(uint64_t *)&value);
	//} else {
	//pack_varint(&b, value);
	//}
	//
	//lua_settop(L, 1);
	//luaL_pushresult(&b);
	//lua_call(L, 1, 0);

	l_value := L.ToNumber(2)
	fmt.Println("pb.go   ------------       signed_varint_encoder:", int64(l_value))

	l_func := L.ToFunction(1)
	value := int64(l_value)
	var b string
	if value < 0{
		b = pack_varint("", uint64(value))			// 把数字变成string
	}else{
		b = pack_varint("", uint64(value))			// 把数字变成string
	}
	fmt.Printf("pb.go   ------------       signed_varint_encoder     out :  %v \n", []byte(b))
	//L.Push(lua.LString(b))
	//fmt.Println("-------------------------",b,":",l_func.String())
	//L.Call(1,0)			// string作为参数，调用参数1作为函数名
	if err := L.CallByParam(lua.P{
		Fn:      l_func,
		NRet:    0,
		Protect: true,
	}, lua.LString(b)); err != nil {
		println("signed_varint_encoder error:", err.Error())
	}
	return 0
}

//------------------------------------------struct_pack------------------------------------------------------

func pack_fixed32(L *lua.LState, value lua.LNumber) int{
	//#ifdef	IS_LITTLE_ENDIAN
	//lua_pushlstring(L, (char *)	value, 4);

	fmt.Println("pack_fixed32----",value)
	str:= strconv.Itoa(int(value))
	L.Push(lua.LString(str))
	//# else
	//uint32_t	v = htole32(*(uint32_t *)	value);
	//lua_pushlstring(L, (char*)&v, 4);
	//#endif
	return 0
}

func pack_fixed64(L *lua.LState, value lua.LNumber) int {
	//#ifdef	IS_LITTLE_ENDIAN
	//lua_pushlstring(L, (char *)	value, 8);
	fmt.Println("pack_fixed64----",value)

	str:= strconv.FormatInt(int64(value),10)
	L.Push(lua.LString(str))
	//# else
	//uint64_t	v = htole64(*(uint64_t *)	value);
	//lua_pushlstring(L, (char*)&v, 8);
	//#endif
	return 0
}

func struct_pack(L *lua.LState) int {

	//function := L.ToFunction(1)             /* get argument */
	format := L.ToInt(2)             /* get argument */
	value := L.ToNumber(3)             /* get argument */


	fmt.Printf("pb.go   ----------------------------------------------------------      struct_pack:     format %d,    value %d       \n"   ,format,value)
	//pos:= L.ToInt(3)
	//L.Push(lua.LString(str)) /* push result */
	//buf:=buffer[pos:]

	//var ii lua.LNumber
	var out lua.LString

	//uint8_t	format = luaL_checkinteger(L, 2);
	//lua_Number	value = luaL_checknumber(L, 3);
	//lua_settop(L, 1);
	switch format {
	case 'i':
		{
			//ii := lua.LNumber(int(value))
			//pack_fixed32(L, lua.LNumber(ii))
			break
		}
	case 'q':
		{
			//int64_t			v = (int64_t)			value;
			//pack_fixed64(L, (uint8_t*)&v);
			//ii := lua.LNumber(int64(value))
			//pack_fixed64(L, lua.LNumber(ii))
			break
		}
	case 'f':
		{
			//float			v = (float)			value;
			//pack_fixed32(L, (uint8_t*)&v);
			fmt.Println("float")
			//ii = lua.LNumber(float32(value))
			bb := byte(float32(value))
			out = lua.LString(string(bb))
			//pack_fixed32(L, lua.LNumber(ii))
			break
		}
	case 'd':
		{
			//double			v = (double)			value;
			//pack_fixed64(L, (uint8_t*)&v);
			//ii := lua.LNumber(float64(value))
			//pack_fixed64(L, lua.LNumber(ii))
			break
		}
	case 'I':
		{
			//uint32_t			v = (uint32_t)			value;
			//pack_fixed32(L, (uint8_t*)&v);
			//ii := lua.LNumber(uint32(value))
			//pack_fixed32(L, lua.LNumber(ii))
			break
		}
	case 'Q':
		{
			//uint64_t			v = (uint64_t)			value;
			//pack_fixed64(L, (uint8_t*)&v);
			//ii := lua.LNumber(uint64(value))
			//pack_fixed64(L, lua.LNumber(ii))
			break
		}
	//default:
	//	ii := lua.LNumber(0)
	}
	//lua_call(L, 1, 0);
	//L.Call(1,0)
	l_func := L.ToFunction(1)			// lua 部分是 write(value)
	if err := L.CallByParam(lua.P{
		Fn:      l_func,
		NRet:    0,
		Protect: true,
	}, out); err != nil {
		println("signed_varint_encoder error:", err.Error())
	}
	return 0
}
//----------------------------------------int decode---------------------------------------------------------
func size_varint(buffer string, len int) uint64{
	pos := 0
	bytes := []byte(buffer)
	for	{
		if bytes[pos] & 0x80 == 0 {
			break
		}
		pos++
		if pos > len {
			return LuaInt64Max
		}
	}
	re:=uint64(pos + 1)
	//println("---------size_varint--------",re)
	return re
}

func unpack_varint(buffer string, len uint64) uint64{
	//uint64_t	value = buffer[0] & 0x7f
	//size_t	shift = 7;
	//size_t	pos = 0;
	//for (pos = 1; pos < len; ++pos)
	//{
	//value |= ((uint64_t)(buffer[pos] & 0x7f)) << shift;
	//shift += 7;
	//}
	//bb,_:= strconv.Atoi(buffer[0:1])
	bb:= []byte(buffer)

	value := uint64(bb[0] & 0x7f)
	shift := uint64(7)
	pos := uint64(0)

	fmt.Printf("pb.go   -----read-------unpack_varint  %v %d  \n" , bb  , len)

	for pos=1;pos< len; pos++{
		value |= ((uint64)(bb[pos] & 0x7f)) << shift
		shift += 7
	}
	fmt.Println("pb.go   -----read-------   unpack_varint  out      ",value)
	return value
}

func varint_decoder(L *lua.LState) int {
	//size_t	len;
	//const char *buffer = luaL_checklstring(L, 1, &len);
	//size_t	pos = luaL_checkinteger(L, 2);
	//buffer += pos;
	//len = size_varint(buffer, len);
	//if (len == -1) {
	//	luaL_error(L, "error data %s, len:%d", buffer, len);
	//} else {
	//	lua_pushnumber(L, (lua_Number)
	//	unpack_varint(buffer, len));
	//	lua_pushinteger(L, len+pos);
	//}

	fmt.Printf("pb.go   ------------       varint_decoder:")
	buffer := L.ToString(1)             /* get argument */
	pos := L.ToInt64(2)             /* get argument */
	buf:= buffer[pos:]

	tLen := size_varint(buf, len(buffer))
	if tLen == LuaInt64Max{
		println("error varint_decoder data %s, tLen:%d", buffer, tLen)
	} else {
		ii := unpack_varint(buf, tLen)
		L.Push(lua.LNumber(ii))
		L.Push(lua.LNumber(tLen +uint64(pos)))
	}
	return 2
}

func signed_varint_decoder(L *lua.LState) int {
	//size_t	len;
	//const char *buffer = luaL_checklstring(L, 1, &len);
	//size_t	pos = luaL_checkinteger(L, 2);
	//buffer += pos;
	//len = size_varint(buffer, len);
	//if (len == -1) {
	//	luaL_error(L, "error data %s, len:%d", buffer, len);
	//} else {
	//	lua_pushnumber(L, (lua_Number)(int64_t)		unpack_varint(buffer, len));
	//	lua_pushinteger(L, len+pos);
	//}


	buffer := L.ToString(1)             /* get argument */
	pos := L.ToInt64(2)             /* get argument */
	buf:= buffer[pos:]

	fmt.Printf("pb.go   ------read------       signed_varint_decoder:    %v      %d  \n",[]byte(buffer),pos)
	tLen := size_varint(buf, len(buffer))
	if tLen == LuaInt64Max{
		println("error signed_varint_decoder data %s, tLen:%d", buffer, tLen)
	} else {
		ii := int64(unpack_varint(buf, tLen))
		L.Push(lua.LNumber(ii))
		L.Push(lua.LNumber(tLen +uint64(pos)))

		fmt.Println("pb.go   ------read------       signed_varint_decoder   out        ",ii)
	}


	return 2
}
//-------------------------------------------------------------------------------------------------
func zig_zag_encode32(L *lua.LState) int {
	//int32_t	n = luaL_checkinteger(L, 1);
	//uint32_t	value = (n << 1) ^ (n >> 31);
	//lua_pushinteger(L, value);

	println("pb.go   ------------       zig_zag_encode32:")
	n := L.ToInt(1)             /* get argument */
	value := uint32((n << 1) ^ (n >> 31))
	L.Push(lua.LNumber(value)) /* push result */

	return 1
}

func zig_zag_decode32(L *lua.LState) int {
	//uint32_t	n = (uint32_t)	luaL_checkinteger(L, 1);
	//int32_t	value = (n >> 1) ^ - (int32_t)(n & 1);
	//lua_pushinteger(L, value);

	println("pb.go   ------------       zig_zag_decode32:")
	n := uint32(L.ToInt(1))             /* get argument */

	value := (int)(n >> 1) ^ - (int)(n & 1)
	L.Push(lua.LNumber(value)) /* push result */


	return 1
}

func zig_zag_encode64(L *lua.LState) int {
	//int64_t	n = (int64_t)	luaL_checknumber(L, 1);
	//uint64_t	value = (n << 1) ^ (n >> 63);
	//lua_pushinteger(L, value);

	println("pb.go   ------------       zig_zag_encode64:")
	n := L.ToInt64(1)             /* get argument */
	value := uint64((n << 1) ^ (n >> 63))
	L.Push(lua.LNumber(value)) /* push result */


	return 1
}

func zig_zag_decode64(L *lua.LState) int {
	//uint64_t	n = (uint64_t)	luaL_checknumber(L, 1);
	//int64_t	value = (n >> 1) ^ - (int64_t)(n & 1);
	//lua_pushinteger(L, value);

	println("pb.go   ------------       zig_zag_decode64:")
	n := uint64(L.ToInt64(1))
	value :=  (int64)(n >> 1) ^ - (int64)(n & 1)
	L.Push(lua.LNumber(value)) /* push result */
	return 1
}


func read_tag(L *lua.LState) int {
	//size_t	tLen;
	//const char *buffer = luaL_checklstring(L, 1, &tLen);
	//size_t	pos = luaL_checkinteger(L, 2);
	//buffer += pos;
	//tLen = size_varint(buffer, tLen);
	//if (tLen == -1) {
	//	luaL_error(L, "error data %s, tLen:%d", buffer, tLen);
	//} else {
	//	lua_pushlstring(L, buffer, tLen);
	//	lua_pushinteger(L, tLen+pos);
	//}

	buffer := L.ToString(1)
	pos := uint64(L.ToInt64(2))
	len1:= len(buffer)

	//bb:= []byte(buffer)
	buf:=buffer[pos:]
	tLen := size_varint(buf, len1)

	fmt.Println("pb.go   ----read--------       read_tag:",buffer, "      ", pos)

	if tLen == LuaInt64Max {
		println("error data %s, tLen:%d", buffer, tLen)
	} else {
		str:= buffer[:tLen]
		L.Push(lua.LString(str))
		L.Push(lua.LNumber(tLen +pos))

		fmt.Printf("pb.go   ----read--------       read_tag   out  %v\n",[]byte(str))
	}
	return 2
}
//-----------------------------------------------struct unpack--------------------------------------------------
func unpack_fixed32(buffer string,  cache uint32)int{
//#ifdef IS_LITTLE_ENDIAN
	iInt32, err :=strconv.Atoi(buffer)
	if err != nil {
		println("unpack_fixed32  error!", err.Error())
	}
return iInt32
//#else
//*(uint32_t*)cache = le32toh(*(uint32_t*)buffer);
//return cache;
//#endif
}

func unpack_fixed64(buffer string, cache uint64) int64 {
	//#ifdef IS_LITTLE_ENDIAN

	iInt64, err := strconv.ParseInt(buffer, 10, 64)
	if err != nil {
		println("unpack_fixed64  error!", err.Error())
	}
	return iInt64
	//#else
	//*(uint64_t*)cache = le64toh(*(uint64_t*)buffer);
	//return cache;
	//#endif
}


func struct_unpack(L *lua.LState) int {

	fmt.Printf("pb.go   ------------      struct_unpack:  \n")
	format := L.ToInt(1)             /* get argument */
	buffer := L.ToString(2)             /* get argument */

	pos:= L.ToInt(3)
	//L.Push(lua.LString(str)) /* push result */
	buf:=buffer[pos:]



	//uint8_t	format = luaL_checkinteger(L, 1);
	//size_t	len;
	//const uint8_t *buffer = (uint8_t *)luaL_checklstring(L, 2, &len);
	//size_t	pos = luaL_checkinteger(L, 3);
	//buffer += pos;
	//uint8_t	out[8];
	switch format {
	case 'i':
		{
			//lua_pushinteger(L, *(int32_t *)			unpack_fixed32(buffer, nil))
			ii := unpack_fixed32(buf, 0)
			L.Push(lua.LNumber(int32(ii)))
			break
		}
	case 'q':
		{
			ii := unpack_fixed64(buf, 0)
			L.Push(lua.LNumber(int64(ii)))
			//lua_pushnumber(L, (lua_Number)*(int64_t*)			unpack_fixed64(buffer, out));
			break
		}
	case 'f':
		{
			ii := unpack_fixed32(buf, 0)
			L.Push(lua.LNumber(float32(ii)))
			//lua_pushnumber(L, (lua_Number)*(float*)			unpack_fixed32(buffer, out));
			break
		}
	case 'd':
		{
			ii := unpack_fixed64(buf, 0)
			L.Push(lua.LNumber(float64(ii)))
			//lua_pushnumber(L, (lua_Number)*(double*)			unpack_fixed64(buffer, out));
			break
		}
	case 'I':
		{
			ii := unpack_fixed32(buf, 0)
			L.Push(lua.LNumber(uint32(ii)))
			//lua_pushnumber(L, *(uint32_t *)			unpack_fixed32(buffer, out));
			break
		}
	case 'Q':
		{
			ii := unpack_fixed64(buf, 0)
			L.Push(lua.LNumber(uint64(ii)))
			//lua_pushnumber(L, (lua_Number)*(uint64_t*)			unpack_fixed64(buffer, out));
			break
		}
		//default:
		//	luaL_error(L, "Unknown, format");
	}
	return 1
}

func iostring_new(L *lua.LState) int {
	println("pb.go   ------------       iostring_new:")
	//IOString * io = (IOString *)
	//lua_newuserdata(L, sizeof(IOString));
	//io- > size = 0;
	//luaL_getmetatable(L, IOSTRING_META);
	//lua_setmetatable(L, -2);
	return 0
}


//----------------------------------------------------string--------------------------------------------

// __tostring 方法
func iostring_str(L *lua.LState) int {
	println("pb.go   ------------       iostring_str:")
	str := L.ToString(1)             /* get argument */
	L.Push(lua.LString(str)) /* push result */
	//IOString * io = checkiostring(L)
	//lua_pushlstring(L, io- > buf, io- > size)
	return 1
}
// __len
func iostring_len(L *lua.LState) int {
	println("pb.go   ------------       iostring_len:")
	str := L.ToString(1)             /* get argument */
	L.Push(lua.LNumber(len(str))) /* push result */
	//IOString * io = checkiostring(L);
	//lua_pushinteger(L, io- > size);
	return 1
}

func iostring_write(L *lua.LState) int {
	str := L.ToString(1)             /* get argument */
	str2 := L.ToString(2)             /* get argument */
	println("pb.go   ------------       iostring_write:", str,"+",str2)
	//IOString * io = checkiostring(L);
	//size_t
	//size;
	//const char *str = luaL_checklstring(L, 2, &size);
	//if (io- > size+size > IOSTRING_BUF_LEN) {
	//	luaL_error(L, "Out of range");
	//}
	//memcpy(io- > buf+io- > size, str, size);
	//io- > size += size;
	return 0
}

func iostring_sub(L *lua.LState) int {
	println("pb.go   ------------       iostring_sub:")
	str := L.ToString(1)             /* get argument */
	begin := L.ToInt64(2)             /* get argument */
	end := L.ToInt64(3)             /* get argument */

	re:=str[begin:end]
	L.Push(lua.LString(re)) /* push result */
	//IOString * io = checkiostring(L);
	//size_t
	//begin = luaL_checkinteger(L, 2);
	//size_t
	//end = luaL_checkinteger(L, 3);
	//if (begin > end || end > io- > size) {
	//	luaL_error(L, "Out of range");
	//}
	//lua_pushlstring(L, io- > buf+begin-1, end-begin+1);
	return 1
}

func iostring_clear(L *lua.LState) int {
	str := L.ToString(1)             /* get argument */
	println("pb.go   ------------        iostring_clear:", str)
	//IOString * io = checkiostring(L);
	//io- > size = 0;
	return 0
}

func ZswLuaShowBytesToString(L *lua.LState) int  {
	str := L.ToString(1)

	fmt.Printf("*******************************************************************************ZswLuaShowBytesToString: %v \n", []byte(str))
	return 0
}