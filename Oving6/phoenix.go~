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

func CreateSlave() {
   number = str(number)
   cmd := exec.Command("mate-terminal", "-x", "go", "run",  "phoenix.go", number)
	err := cmd.Start()
	CheckError(err)
}

func Broadcast(conn *net.UDPConn) {
   
}

func ReceiveBroadcast(conn *net.UDPConn) int {
   _, _, err := conn.ReadFromUDP
}

func Count(i int, conn *net.UDPConn) int {
   b := make([]byte, 1)
   conn.Write(b)
   return ++i
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
	case true:
	   // Establish broadcast-socket on first run
      service := ":12000"
      addr, err := net.ResolveUDPAddr("udp", service)
      CheckError(err)

      conn, err := net.ListenUDP("udp", addr)
      CheckError(err)

      CreateSlave()
		   
      for {
         number = Count() // oppdateres den globale?
      }
		fmt.Println("test5")
		
   case false:
      // Establish listen-socket
      service := ":12001"
      addr, err := net.ResolveUDPAddr("udp", service)
      CheckError(err)
      
      conn, err := net.ListenUDP("udp", addr)
      CheckError(err)
      
      // Receiving number broadcast and checking if master is alive
      for alive {
         number = ReceiveBroadcast(conn)
         alive = CheckAlive
      }

		fmt.Println("test6")
	}
	
	fmt.Println("end")
}
