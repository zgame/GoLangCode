package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/lucas-clemente/quic-go"
)

func ClientMain() error {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	fmt.Println("*********1*******")
	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		fmt.Println("连接不上服务器", err.Error())
		return err
	}
	for {
	fmt.Println("********2********")
	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		fmt.Println("OpenStreamSync ", err.Error())
		return err
	}

	fmt.Println("********3********")
	fmt.Printf("Client: Sending '%s'\n", message)
	_, err = stream.Write([]byte(message))
	if err != nil {
		fmt.Println("write buffer ", err.Error())
		return err
	}

		//fmt.Println("*********4*******")
		//stream, err := session.AcceptStream(context.Background())
		//if err != nil {
		//	fmt.Println("OpenStreamSync ", err.Error())
		//	return err
		//}
		fmt.Println("*********5*******")
		buf := make([]byte, 1024)
		_, err = stream.Read(buf)
		if err != nil {
			fmt.Println("read buf ", err.Error())
			return err
		}
		fmt.Printf("Client: Got '%s'\n", buf)

	}
	return nil
}
