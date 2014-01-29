package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {

	//port 34933 for message size 1024, port 33546 for terminering med \x00
	service := "127.0.0.1:12345"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	conn, _ := listener.Accept()
	buf := make([]byte, 1024)
	conn.Read(buf)

	tcpAddr2, _ := net.ResolveTCPAddr("tcp4", string(buf))
	conn2, _ := net.DialTCP("tcp", nil, tcpAddr2)
	conn2.Write([]byte("Hello!\n"))
	//fmt.Printf("%s", buf)
	
	for {
		n, err1 := conn2.Read(buf[0:])
		if err1 != nil {
			return
		}
		_, err2 := conn2.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
	conn.Close()
	fmt.Println("end")
	os.Exit(0)	
}
