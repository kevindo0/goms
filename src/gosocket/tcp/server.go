package main

import (
	"fmt"
	"net"
	"log"
	"time"
)

func Serial(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(conn.RemoteAddr().String())

		//通过conn的wirte方法将这些数据返回给客户端。
		conn.Write([]byte("hello world\n"))
		conn.Write([]byte("hello golang\n"))
		conn.Close()
	}
}

func Handle_conn(conn net.Conn) {
	fmt.Println(conn.RemoteAddr().String())
	time.Sleep(5*time.Second)
	conn.Write([]byte(time.Now().Local().String()))
	conn.Close()
}

func Parallel(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go Handle_conn(conn)
	}
}

func main() {
	fmt.Println("Server Start...")
	addr := ":10001"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	// Serial(listener)
	Parallel(listener)
}