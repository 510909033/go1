package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "3333", "port")

func server() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host+":"+*port)

	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + *host + ":" + *port)

	fmt.Println(l.Addr())

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()

	var b_msg []byte
	b_msg = make([]byte, 1000)
	var len int
	var err error
	for {
		conn.SetWriteDeadline(time.Now().Add(time.Second * 5))

		len, err = conn.Read(b_msg)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			panic("exit")
		}
		err = err
		fmt.Println(len)
		fmt.Println(b_msg)
		//io.Copy(conn, conn)
		conn.Write([]byte("hello"))
	}

}
