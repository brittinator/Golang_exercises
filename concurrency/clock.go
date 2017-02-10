/* Clock implements a clock server with go routines,
which enables it to serve multiple requests to it.

Usage:
go run clock.go
nc localhost 3000 (try this on multiple tabs)
*/

package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	// infinite loop
	for {
		_, err := io.WriteString(c, time.Now().Format("15:03:05\n"))
		if err != nil {
			// this means the client disconnected
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // connection aborted
			continue
		}
		go handleConn(conn)
	}
}
