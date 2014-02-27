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
	
	fmt.Println(len(os.Args))
	
	/*
	if len(os.Args) == 2 {
	   // If clone
	   addr, err := net.ResolveTCPAddr("tcp", ":1200")
	   CheckError(err)
	   
	   listener, err := net.ListenTCP("tcp", addr)
	   CheckError(err)
	   
	   conn, err := listener.Accept()
	   CheckError(err)
	   
	   master = false
	   backup = true
	   fmt.Println("test1")
	   
	} else {
	   // If original instance
   	addr, err := net.ResolveTCPAddr("tcp", ":1200")
	   CheckError(err)
	   
	   conn, err := net.DialTCP("tcp", nil, addr)
	   CheckError(err)
	   
	   master = true
	   backup = false
      fmt.Println("test2")
	}
   */
   
   var master, backup bool
   var conn *net.UDPConn
   addr, err := net.ResolveUDPAddr("udp", ":1200")
   CheckError(err)
   
   
   if len(os.Args) == 2 {
      // clone
      conn, err = net.DialUDP("udp", nil, addr)
      
      master = false
      fmt.Println("test1")
   } else {
      // original instance
      conn, err = net.ListenUDP("udp", addr)
      CheckError(err)
      
      conn.Write([]byte("test"))
      
      master = true
      fmt.Println("test2")
   }
	
	for {
		switch master {
		case true:
		   //conn.Write([]byte("test"))
			fmt.Println("test5")
		case false:
			cmd := exec.Command("mate-terminal", "-x", "go", "run",  "phoenix.go", "slave")
			err := cmd.Start()
			CheckError(err)

			fmt.Println("test6")
		}
		
		switch backup {
		case true:
		case false:
		}
	}
	
	fmt.Println("end")
}
