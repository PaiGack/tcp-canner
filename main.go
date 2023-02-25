package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for i := 21; i < 100; i++ {
		address := fmt.Sprintf("127.0.0.1:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Printf("%s cloesd\n", address)
			continue
		}
		conn.Close()
		log.Printf("%s opened\n", address)
	}
}
