package main

import (
	"fmt"
	"os/exec"
	"os"
	"net"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {

	// Initiate clone of self here
	cmd := exec.Command("mate-terminal", "-x go run phoenix.go")
	err := cmd.Start()

	// Establishes broadcast/listen connections
	
	bAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:12000")
	lAddr, err := net.ResolveUDPAddr("udp", ":0")
	CheckError(err)
	
	bConn, err := net.DialUDP("udp", nil, bAddr)
	lConn, err := net.ListenUDP("udp", lAddr)
	CheckError(err)
	
	
	buf := []byte("test")
	bConn.WriteToUDP(buf, bAddr)
	buf2 := make([]byte, 1024)
	lConn.ReadFromUDP(buf2)
	
	

}
