package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	fmt.Println("ready")
	sig := <-sigCh
	fmt.Println("get signal")

	switch sig {
	case syscall.SIGTERM:
		fmt.Println("SIGTERM")
	case syscall.SIGINT:
		fmt.Println("SIGINT")
	case os.Interrupt:
		fmt.Println("Interrupt")
	}
	fmt.Println("finish")
}
