package main

import (
	"fmt"
	"os/exec"
	"os"
	"net"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func CreateSlave(i int) {
   strI := strconv.Itoa(i)
   cmd := exec.Command("mate-terminal", "-x", "go", "run",  "phoenix.go", strI)
	err := cmd.Start()
	CheckError(err)
}

func ReceiveCount(i int, conn *net.UDPConn) (int, error) {
   b := make([]byte, 1)
   _, _, err := conn.ReadFromUDP(b)
   i = i + 1
   return i, err
}

func Count(i int, conn *net.UDPConn) int {
   fmt.Println(i)
   b := make([]byte, 1)
   conn.Write(b)
   i = i + 1
   return i
}

func main() {
	
   fmt.Println(len(os.Args))
	var master bool
	var number int
   
   if len(os.Args) == 2 {
      // clone      
      master = false
      fmt.Println("test1")
   } else {
      // original instance
      master = true
      number := 0
      fmt.Println("test2")
   }
	
	switch master {
	case false:
      // Establish listen-socket
      service := ":12001"
      addr, err := net.ResolveUDPAddr("udp", service)
      CheckError(err)
      
      conn, err := net.ListenUDP("udp", addr)
      conn.SetDeadline(time.Now())
      CheckError(err)
      
      // Receiving number broadcast and checking if master is alive
      for  {
        number, err := ReceiveCount(number, conn)
      }
      master = true
      fmt.Println("test6")
      fallthrough
	
	case true:
	   // Establish broadcast-socket on first run
      service := ":12000"
      addr, err := net.ResolveUDPAddr("udp", service)
      CheckError(err)

      conn, err := net.ListenUDP("udp", addr)
      CheckError(err)

      CreateSlave()
		   
      for {
         number = Count(number, conn) // oppdateres den globale?
      }
	}
	
	fmt.Println("end")
}
