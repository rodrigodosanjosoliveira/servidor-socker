package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func handler(conn net.Conn) {
	for {
		m, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed")
				conn.Close()
				return
			}
			fmt.Println("Error reading from connection", err)
			return
		}
		_, err = conn.Write([]byte(m))
		if err != nil {
			fmt.Println("Error writing to connection")
			return
		}
		fmt.Printf("%v %q\n", conn.RemoteAddr(), m)
	}
}

func main() {
	fmt.Println("Listening on port 8080")

	ln, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := ln.Accept()
		fmt.Println("Connection accepted")
		go handler(conn)
	}
}
