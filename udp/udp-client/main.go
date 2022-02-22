package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	host := flag.String("host", "127.0.0.1", "specify the host for the UDP server")
	port := flag.Int("port", 51301, "specify the port for the UDP server")
	flag.Parse()

	url := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("send to: ", url)
	conn, err := net.Dial("udp4", url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	i := 0
	for {
		m := fmt.Sprintf("Hello(%d)", i)
		if _, err := fmt.Fprint(conn, m); err != nil {
			log.Fatal(err)
		}
		log.Println("send message:", m)
		i++
		time.Sleep(2 * time.Second)
	}
}
