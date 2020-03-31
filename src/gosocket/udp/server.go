package main

import (
	"fmt"
	"os"
	"net"
	"time"
)

func main() {
	fmt.Println("Udp server start...")
	service := ":12000"
	updAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", updAddr)
	checkError(err)
	defer conn.Close()
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	// fmt.Println(conn.RemoteAddr().String())
	var buf [512]byte
    _, addr, err := conn.ReadFromUDP(buf[0:])
    if err != nil {
        return
    }
    daytime := time.Now().String()
    conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())
		os.Exit(1)
	}
}