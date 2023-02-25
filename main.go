package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}

	for i := 1; i < 65536; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			address := fmt.Sprintf("127.0.0.1:%d", i)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				log.Printf("%s cloesd\n", address)
				return
			}
			conn.Close()
			log.Printf("%s opened\n", address)
		}(i)
	}

	wg.Wait()
	log.Printf("cost %f Seconds", time.Since(start).Abs().Seconds())
}
