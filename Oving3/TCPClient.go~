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
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	
	//serverAddr := "129.241.187.161:34933"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "129.241.187.161:33546")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	checkError(err)
	data := make([]byte, 512)
	_, err = conn.Read(data)
	dataIn := string(data[:])
	fmt.Println(dataIn)

	b := make([]byte, 1024)
	conn.Write([]byte("hallasldasdasdasdasdasdadssd\x00"))
	conn.Read(b)
	fmt.Printf("%s", b)
	conn.Write([]byte("en ny beskjed\x00"))
	conn.Read(b)
	fmt.Printf("%s", b)

	checkError(err)

	//
	tcpServer, err := net.ResolveTCPAddr("tcp", "129.241.187.140:33443")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpServer)
	conn.Write([]byte("connect to: 129.241.187.140:33443"))
	conn.Read(b)
	fmt.Println("%s", b)
	

	go func(){
		for {
			connn,err := listener.Accept()
			if err != nil {
				continue
			}
			connn.Write([]byte("ahllow hal\x00"))
			connn.Read(b)
			fmt.Printf("%s", b)
			
		}
	}()
	
	


	os.Exit(0)
	
}
