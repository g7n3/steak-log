package main

import (
	"github.com/cookedsteak/steak-log/log"
	"fmt"
	"net"
	"strings"
)

func main() {
	localAddress, _ := net.ResolveTCPAddr("tcp4", log.GetConfig("LocalAddress"))
	var tcpListener, err= net.ListenTCP("tcp", localAddress)
	if err != nil {
		fmt.Println("Listening error：", err)
		return
	}
	defer func() {
		tcpListener.Close()
	}()
	fmt.Println("Waiting for connections...")
	for {
		var conn, err2= tcpListener.AcceptTCP()
		if err2 != nil {
			fmt.Println("Can not accept connections：", err2)
			return
		}

		var remoteAddr= conn.RemoteAddr()
		fmt.Println("Connection accepted：", remoteAddr)
		fmt.Println("Data transferred...")

		go tcpPipe(conn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	var data = make([]byte, 1024)
	for {
		len, err := conn.Read(data)
		if len <= 0 || err != nil {
			fmt.Println("Connection lost")
			conn.Close()
			break
		}
		res := strings.Replace(fmt.Sprintf("%s", string(data)), "\n", "", -1)
		log.Handle(res)
		data = make([]byte, 1024)
	}
}