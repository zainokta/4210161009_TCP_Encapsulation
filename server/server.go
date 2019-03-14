package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type Hero struct {
	ID     int
	XCoord int
	YCoord int
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server running...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	h := &Hero{}
	dec.Decode(h)
	fmt.Println("Hero ID : ", h.ID, "Moving at x: ", h.XCoord, " y: ", h.YCoord)
	conn.Close()
}
