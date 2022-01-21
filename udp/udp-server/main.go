package main

import (
	"flag"
	"log"
	"net"
)

const (
	maxPacketLength = 65535
)

func main() {
	p := flag.Int("port", 51301, "specify the port for the UDP server")
	flag.Parse()
	log.Printf("listen port: %d", *p)

	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: *p,
	}
	//listenAddress, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", "", *p))
	//if err != nil {
	//	log.Fatal(err)
	//}

	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, maxPacketLength)
	for {
		n, addr, err := listener.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			log.Printf("From: %v Reciving data: %s", addr.String(), string(buffer[:n]))
		}()
	}

}
