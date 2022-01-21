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

	udpAddr := &net.UDPAddr{
		//IP:   net.ParseIP("localhost"),
		IP:   net.ParseIP("0.0.0.0"),
		Port: *p,
	}
	log.Println(udpAddr.String())

	listener, err := net.ListenUDP("udp4", udpAddr)
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
