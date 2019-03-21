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
	Mana   int
	Gold   int
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server running...")
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	dec := gob.NewDecoder(conn)
	h := &Hero{}
	dec.Decode(h)

	fmt.Println("Hero ID : ", h.ID, " Moving at x: ", h.XCoord, " y: ", h.YCoord, " Mana : ", h.Mana, " Gold : ", h.Gold)

	handleConnection(h, conn)
	conn.Close()
}

func handleConnection(h *Hero, conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(h)
}
