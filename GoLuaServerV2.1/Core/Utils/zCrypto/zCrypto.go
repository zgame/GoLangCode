package zCrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"strings"
)

//-------------------------------------------------------------------------------
//加密 解密 验证
//-------------------------------------------------------------------------------

var exports = map[string]lua.LGFunction{
	"md5":           md5Str,          //与
	"base64_encode": base64EncodeStr, //非
	"base64_decode": base64DecodeStr, //或
}

// ----------------------------------------------------------------------------

func cryptoLoader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)
	return 1
}

// ----------------------------------------------------------------------------

func LuaCryptoLoad(L *lua.LState) {
	L.PreloadModule("crypto", cryptoLoader)
}



// md5验证
func md5Str(L *lua.LState) int {
	src := L.CheckString(1)
	h := md5.New()
	h.Write([]byte(src)) // 需要验证的字符串为
	//fmt.Printf("MD5:               %s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	result :=hex.EncodeToString(h.Sum(nil))
	L.Push(lua.LString(result))

	return 1

}


//------------------------BASE64--------------------------------------

// base编码
func base64EncodeStr(L *lua.LState) int {
	src := L.CheckString(1)
	result := string(base64.StdEncoding.EncodeToString([]byte(src)))
	L.Push(lua.LString(result))

	return 1
}

// base解码
func base64DecodeStr(L *lua.LState) int {
	src := L.CheckString(1)
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		L.Push(lua.LString("decode error :"+err.Error()))
		return 1
	}
	L.Push(lua.LString(string(a)))
	return 1
}


//----------------------AES---------------------------------

//aesKey :=[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,	}// Appendix C.3.  AES-256
//aesStr := AESEncodeStr(orignSrc,string(aesKey))

var ivspec = []byte("0000000000000000")

func AESEncodeStr(src, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, ivspec)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return hex.EncodeToString(crypted)
}

func AESDecodeStr(crypt, key string) string {
	crypted, err := hex.DecodeString(strings.ToLower(crypt))
	if err != nil || len(crypted) == 0 {
		fmt.Println("plain content empty")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("key error1", err)
	}
	ecb := cipher.NewCBCDecrypter(block, ivspec)
	decrypted := make([]byte, len(crypted))
	ecb.CryptBlocks(decrypted, crypted)

	return string(PKCS5Trimming(decrypted))
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

