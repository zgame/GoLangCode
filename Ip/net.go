package main

import (
	"io"
	"net"
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

//var get_ip = flag.String("get_ip", "", "external|internal")

func main() {
	//fmt.Println("Usage of ./getmyip --get_ip=(external|internal)")
	//flag.Parse()
	//if *get_ip == "external" {
	fmt.Println("",get_external())

	//}

	//if *get_ip == "internal" {
	fmt.Println("",get_internal(0))
	fmt.Println("",get_internal(1))

	//}

}

func get_external() string{
	re :=""
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("",err.Error())
	}
	re = string(body)

	os.Exit(0)
	return re
}

func get_internal(index int) string {
	re :=""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				re = ipnet.IP.String()
			}
		}
	}
	os.Exit(0)
	return re
}