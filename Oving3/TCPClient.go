package main

import (
	"fmt"
	"net"
	"os"
	//"io/ioutil"
)

//Avbryter programmet og printer error-melding fra den aktuelle funksjonen
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	for {	
		buf := make([]byte, 1024)
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}


func main() {
	
	service := "127.0.0.1:12000"
	serverAddr := "127.0.0.1:12345"
	//serverAddr := "129.241.187.161:34933"
	//port 34933 for message size 1024, port 33546 for terminering med \x00

	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)


	conn.Write([]byte(service))
	
	listenAddr, _ := net.ResolveTCPAddr("tcp4", service)
	listener, _ := net.ListenTCP("tcp", listenAddr)

	rcv := make([]byte, 1024)
	connRcv, _ := listener.Accept()
	connRcv.Read(rcv)

	fmt.Printf("%s", rcv)

	//b := make([]byte, 1024)
	connRcv.Write([]byte("Beskjed 1: 12345678910 TestTestTest\x00\n"))
	//connRcv.Read(b)
	//fmt.Printf("%s", b)

	for {
		buf := make([]byte, 1024)
		n, err := connRcv.Read(buf[0:])
		fmt.Printf("%s", buf)
		if err != nil {
			return
		}
		_, err2 := connRcv.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
	

	fmt.Println("end")
	os.Exit(0)
}
