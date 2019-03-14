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
}

func main() {
	fmt.Printf("Connecting to %s...\n", "localhost:8000")

	conn, err := net.Dial("tcp", "localhost:8000")
	fmt.Print("Connected!")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for {
		hero := Hero{ID: 2, XCoord: 70, YCoord: 254}
		encoder := gob.NewEncoder(conn)
		encoder.Encode(hero)

		conn.Close()
	}
}
