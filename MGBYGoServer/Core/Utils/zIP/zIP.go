package zIP

//----------------------------------------------------------------------------
// 获取本机ip， 可以获取内网， 外网ip，  内网ip要注意，有两个
//----------------------------------------------------------------------------

import (
	"net/http"
	"os"
	"io"
	"io/ioutil"
	"net"
)

// 获取外网ip，通过一个网站获取的
func GetExternal() string{
	re := ""
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	body,err:= ioutil.ReadAll(resp.Body)
	re = string(body)

	return re
}


// 获取内网ip ，如果连接了外网会获取两个ip
//fmt.Println("",get_internal()[0])
//fmt.Println("",get_internal()[1])

func GetInternal(index int) string{
	re := make(map[int]string)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	tempIndex := 0
	for _ , a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				re[tempIndex] = ipnet.IP.String()
				tempIndex++
			}
		}
	}
	return re[index]
}
