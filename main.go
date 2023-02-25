package main

import (
	"fmt"
	"log"
	"net"
	"sort"
	"time"
)

func worker(ports chan int, port chan int) {
	for p := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			port <- -p
			continue
		}
		conn.Close()
		port <- p
	}
}

func main() {
	start := time.Now()

	ports := make(chan int, 1024*2*2)
	result := make(chan int)
	const COUNT = 65535

	go func() {
		for i := 1; i <= COUNT; i++ {
			ports <- i
		}
	}()

	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}

	var open_ports []int
	var closed_ports []int
	for i := 0; i < COUNT; i++ {
		port := <-result
		if port <= 0 {
			closed_ports = append(closed_ports, -port)
		} else {
			open_ports = append(open_ports, port)
		}
	}

	sort.Ints(open_ports)
	sort.Ints(closed_ports)

	log.Printf("closed ports %v", closed_ports)
	log.Printf("open ports %v", open_ports)

	close(ports)
	close(result)

	log.Printf("cost %f Seconds", time.Since(start).Abs().Seconds())
}
