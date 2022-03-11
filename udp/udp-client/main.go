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
	slist := []string{
		",MasterController,39.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,30.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,22.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,15.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,7.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,0.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,-7.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,-15.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,-22.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,-30.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
		",MasterController,-39.0,e9:e6:37:1a:2e:c8,E7D61EA3-F8DD-49C8-8F2F-F2484C07ACB9,1,6,",
	}
	flist := []float32{
		-52.00,
		-49.00,
		-40.00,
		-36.00,
		-33.00,
		-30.00,
		-32.00,
		-36.00,
		-40.00,
		-46.00,
		-50.00,
	}
	for {
		m := slist[i%11]
		f := flist[i%11]
		if i%2 > 0 {
			f += 2
		} else {
			f -= 2
		}
		now := fmt.Sprintf("%s%s%.2f", time.Now().Format("2006/01/02 15:04:05"), m, f)
		if _, err := fmt.Fprint(conn, now); err != nil {
			log.Fatal(err)
		}
		log.Println("send message:", now)
		i++
		time.Sleep(1 * time.Second)
	}
}
