package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := "10.0.50.188:12000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	defer conn.Close()
	_, err = conn.Write([]byte("anything"))
    checkError(err)
	var buf [512]byte
    n, err := conn.Read(buf[0:])
    checkError(err)
    fmt.Println(string(buf[0:n]))
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
		fmt.Println("Error ", err.Error())
		os.Exit(1)
	}
}