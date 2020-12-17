package zPbc

import (
	"encoding/binary"
	"fmt"
	"github.com/yuin/gopher-lua"
	"math"
)

const LuaInt64Max  = 18446744073709551615

func LuaOpenPb(L *lua.LState) int {
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
	"varint_encoder":          varint_encoder,
	"signed_varint_encoder":   signed_varint_encoder,
	"read_tag":                read_tag,
	"struct_pack":             struct_pack,
	"struct_unpack":           struct_unpack,
	"varint_decoder":          varint_decoder,
	"signed_varint_decoder":   signed_varint_decoder,
	"zig_zag_decode32":        zig_zag_decode32,
	"zig_zag_encode32":        zig_zag_encode32,
	"zig_zag_decode64":        zig_zag_decode64,
	"zig_zag_encode64":        zig_zag_encode64,
	"new_iostring":            iostring_new,
	"ZswLuaShowBytesToString": ZswLuaShowBytesToString,
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
	//println("pb.go   ------------       varint_encoder:")
	l_value := L.ToNumber(2)
	value := uint64(l_value)
	b := pack_varint("", value) // 把数字变成string

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
	l_value := L.ToNumber(2)
	//fmt.Println("pb.go   ------------       signed_varint_encoder:", int64(l_value))

	l_func := L.ToFunction(1)
	value := int64(l_value)
	var b string
	if value < 0{
		b = pack_varint("", uint64(value)) // 把数字变成string
	}else{
		b = pack_varint("", uint64(value)) // 把数字变成string
	}
	//fmt.Printf("pb.go   ------------       signed_varint_encoder     out :  %v \n", []byte(b))
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

func struct_pack(L *lua.LState) int {

	format := L.ToInt(2)             /* get argument */
	value := L.ToNumber(3)             /* get argument */

	//fmt.Printf("pb.go   ----------------------------------------------------------      struct_pack:     format %d,    value %d       \n"   ,format,value)

	var bb []byte
	switch format {
	case 'i':
		{
			bb = Int32ToBytes(int32(value))
			break
		}
	case 'q':
		{
			bb = Int64ToBytes(int64(value))
			break
		}
	case 'f':
		{
			bb = Float32ToByte(float32(value))
			break
		}
	case 'd':
		{
			bb = Float64ToByte(float64(value))
			break
		}
	case 'I':
		{
			ii := uint32(value)
			bb = Int32ToBytes(int32(ii))
			break
		}
	case 'Q':
		{
			ii := uint64(value)
			bb = Int64ToBytes(int64(ii))
			break
		}
	//default:
	//	ii := lua.LNumber(0)
	}


	out := lua.LString(string(bb))

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
	bb:= []byte(buffer)

	value := uint64(bb[0] & 0x7f)
	shift := uint64(7)
	pos := uint64(0)

	//fmt.Printf("pb.go   -----read-------unpack_varint  %v %d  \n" , bb  , len)

	for pos=1;pos< len; pos++{
		value |= ((uint64)(bb[pos] & 0x7f)) << shift
		shift += 7
	}
	//fmt.Println("pb.go   -----read-------   unpack_varint  out      ",value)
	return value
}

func varint_decoder(L *lua.LState) int {

	buffer := L.ToString(1)             /* get argument */
	pos := L.ToInt64(2)             /* get argument */
	buf:= buffer[pos:]

	//fmt.Printf("pb.go   ------read------      varint_decoder:    %v      %d  \n",[]byte(buffer),pos)

	tLen := size_varint(buf, len(buffer))
	if tLen == LuaInt64Max {
		println("error varint_decoder data %s, tLen:%d", buffer, tLen)
	} else {
		ii := unpack_varint(buf, tLen)
		L.Push(lua.LNumber(ii))
		L.Push(lua.LNumber(tLen +uint64(pos)))
		//fmt.Printf("pb.go   ------read------      varint_decoder:   ii %d   pos   %d  \n",ii,  tLen+uint64(pos))
	}
	return 2
}

func signed_varint_decoder(L *lua.LState) int {

	buffer := L.ToString(1)             /* get argument */
	pos := L.ToInt64(2)             /* get argument */
	buf:= buffer[pos:]

	//fmt.Printf("pb.go   ------read------       signed_varint_decoder:    %v      %d  \n",[]byte(buffer),pos)
	tLen := size_varint(buf, len(buffer))
	if tLen == LuaInt64Max {
		println("error signed_varint_decoder data %s, tLen:%d", buffer, tLen)
	} else {
		ii := int64(unpack_varint(buf, tLen))
		L.Push(lua.LNumber(ii))
		L.Push(lua.LNumber(tLen +uint64(pos)))

		//fmt.Printf("pb.go   ------read------      signed_varint_decoder:   ii %d   pos   %d  \n",ii,  tLen+uint64(pos))
	}


	return 2
}
//-------------------------------------------------------------------------------------------------
func zig_zag_encode32(L *lua.LState) int {
	//println("pb.go   ------------       zig_zag_encode32:")
	n := L.ToInt(1)             /* get argument */
	value := uint32((n << 1) ^ (n >> 31))
	L.Push(lua.LNumber(value)) /* push result */

	return 1
}

func zig_zag_decode32(L *lua.LState) int {
	//println("pb.go   ------------       zig_zag_decode32:")
	n := uint32(L.ToInt(1))             /* get argument */

	value := (int)(n >> 1) ^ - (int)(n & 1)
	L.Push(lua.LNumber(value)) /* push result */


	return 1
}

func zig_zag_encode64(L *lua.LState) int {
	//println("pb.go   ------------       zig_zag_encode64:")
	n := L.ToInt64(1)             /* get argument */
	value := uint64((n << 1) ^ (n >> 63))
	L.Push(lua.LNumber(value)) /* push result */


	return 1
}

func zig_zag_decode64(L *lua.LState) int {
	//println("pb.go   ------------       zig_zag_decode64:")
	n := uint64(L.ToInt64(1))
	value :=  (int64)(n >> 1) ^ - (int64)(n & 1)
	L.Push(lua.LNumber(value)) /* push result */
	return 1
}


func read_tag(L *lua.LState) int {

	buffer := L.ToString(1)
	pos := uint64(L.ToInt64(2))
	len1:= len(buffer)

	buf:=buffer[pos:]
	tLen := size_varint(buf, len1)

	//fmt.Println("pb.go   -----       read_tag:    pos  ", pos)

	if tLen == LuaInt64Max {
		println("error data %s, tLen:%d", buffer, tLen)
	} else {
		str:= buf[:tLen]
		L.Push(lua.LString(str))
		L.Push(lua.LNumber(tLen +pos))

		//fmt.Printf("pb.go   ----     read_tag   out  %v\n",[]byte(str))
	}
	return 2
}
//-----------------------------------------------struct unpack--------------------------------------------------


func struct_unpack(L *lua.LState) int {

	//fmt.Printf("pb.go   ------------      struct_unpack:  \n")
	format := L.ToInt(1)             /* get argument */
	buffer := L.ToString(2)             /* get argument */

	pos:= L.ToInt(3)
	buf:=buffer[pos:]


	bb:=[]byte(buf)
	switch format {
	case 'i':
		{
			ii := BytesToInt32(bb)
			L.Push(lua.LNumber(int32(ii)))
			break
		}
	case 'q':
		{
			ii := BytesToInt64(bb)
			L.Push(lua.LNumber(int64(ii)))
			break
		}
	case 'f':
		{
			ii := ByteToFloat32(bb)
			L.Push(lua.LNumber(float32(ii)))
			break
		}
	case 'd':
		{
			ii := ByteToFloat64(bb)
			L.Push(lua.LNumber(float64(ii)))
			break
		}
	case 'I':
		{
			ii := BytesToInt32(bb)
			L.Push(lua.LNumber(uint32(ii)))
			break
		}
	case 'Q':
		{
			ii := BytesToInt64(bb)
			L.Push(lua.LNumber(uint64(ii)))
			break
		}
		default:
			println("error ------struct_unpack   Unknown, format")
	}
	return 1
}

func iostring_new(L *lua.LState) int {
	println("pb.go   ------------       iostring_new:")
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
	return 0
}

func iostring_sub(L *lua.LState) int {
	println("pb.go   ------------       iostring_sub:")
	str := L.ToString(1)             /* get argument */
	begin := L.ToInt64(2)             /* get argument */
	end := L.ToInt64(3)             /* get argument */

	re:=str[begin:end]
	L.Push(lua.LString(re)) /* push result */
	return 1
}

func iostring_clear(L *lua.LState) int {
	str := L.ToString(1)             /* get argument */
	println("pb.go   ------------        iostring_clear:", str)
	return 0
}

//------------------------------------------------- zsw --------------------------------------------------


func ZswLuaShowBytesToString(L *lua.LState) int  {
	str := L.ToString(1)
	fmt.Printf("********************************ZswLuaShowBytesToString: %v \n", []byte(str))
	return 0
}


func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}


func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}
