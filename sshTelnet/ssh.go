package main

import (
	"log"
	"time"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
)

func main() {
	session, err := connect("Administrator", "****", "*****", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//session.Run("ls /; ls /abc")
	cmd := "cmd.exe /c \"zsw\""			// windows cmd

	session.Run(cmd)
	fmt.Println("cmd:====", cmd)
	fmt.Println("ok")

	out,err := session.Output(cmd)
	fmt.Println("out:=====", string(out))
	if err != nil {
		fmt.Println("Remote command did not exit cleanly:", err)
	}
	w := "this-is-stdout."
	g := string(out)
	if g != w {
		fmt.Printf("Remote command did not return expected string:")
		fmt.Printf("want %q", w)
		fmt.Printf("got  %q", g)
	}

	//session.Run("ls /; ls /abc")
}


func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}