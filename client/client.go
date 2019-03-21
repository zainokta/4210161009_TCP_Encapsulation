package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
)

type Hero struct {
	ID     int
	XCoord int
	YCoord int
	Mana   int
	Gold   int
}

func main() {
	fmt.Printf("Connecting to %s...\n", "localhost:8000")

	conn, err := net.Dial("tcp", "localhost:8000")
	fmt.Print("Connected!\n")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	hero := Hero{ID: 2, XCoord: 70, YCoord: 254, Gold: 3000, Mana: 600}
	encoder := gob.NewEncoder(conn)
	encoder.Encode(hero)

	handleConnection(conn)
	conn.Close()
}

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	h := &Hero{}
	dec.Decode(h)
	fmt.Println("Hero ID : ", h.ID, "Moving at x: ", h.XCoord, " y: ", h.YCoord, " Mana : ", h.Mana, " Gold : ", h.Gold)
}
