package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func rcvJuliusResp(_ch chan string) {
	for word := range _ch {
		fmt.Println(word)
	}
}

func main() {
	conn, _ := net.Dial("tcp", "localhost:10500")

	ch := make(chan string)

	go rcvJuliusResp(ch)

	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')
		// fmt.Println(message)
		isContain := strings.Contains(message, "WORD")

		if isContain == true {
			word := strings.Split(message, "\"")
			ch <- word[1]
		}
	}
}
