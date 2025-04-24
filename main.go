package main

import (
	"fmt"
	"net"
	"io"
	"os"
)

func main() {
	// Create a new server at port 6379
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	fmt.Println("Redis clone is running on port 6379...")

	for {
		buf := make([]byte, 1024)

		// read message from client
		_, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("New client connected:", conn.RemoteAddr())

		// Handle client
        go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintln(conn, "Welcome to the Redis clone!")
}