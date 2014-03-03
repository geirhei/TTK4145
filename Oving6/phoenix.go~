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
   //fmt.Println(err)
   i = i + 1
   conn.SetReadDeadline(time.Now().Add(time.Second))
   return i, err
}

func Count(i int, conn *net.UDPConn, a *net.UDPAddr) int {
   fmt.Println(i)
   b := make([]byte, 1)
   conn.WriteToUDP(b, a)
   i = i + 1
   time.Sleep(500 * time.Millisecond)
   return i
}

func main() {
	
   //fmt.Println(len(os.Args))
	var master bool
	var number int
   
   if len(os.Args) == 2 {
      // clone      
      master = false
      fmt.Println("test1")
   } else {
      // original instance
      master = true
      number = 0
      fmt.Println("test2")
   }
	
	switch master {
	case false:
      // Establish listen-socket
      laddr, err := net.ResolveUDPAddr("udp", ":12001")
      CheckError(err)
      
      conn, err := net.ListenUDP("udp", laddr)
      conn.SetReadDeadline(time.Now().Add(time.Second))
      CheckError(err)
     
      fmt.Println("test3")
      // Receiving number broadcast and checking if master is alive
      for err == nil {
         number, err = ReceiveCount(number, conn)
      }
      master = true
      fmt.Println("test6")
      fallthrough
	
	case true:
	   // Establish broadcast-socket on first run
      laddr, err := net.ResolveUDPAddr("udp", ":12000")
      baddr, err := net.ResolveUDPAddr("udp", "localhost:12001")
      CheckError(err)

      conn, err := net.ListenUDP("udp", laddr)
      CheckError(err)

      CreateSlave(number)
		   
      for {
         number = Count(number, conn, baddr)
      }
	}
}
