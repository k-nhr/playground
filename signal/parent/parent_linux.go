package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func main()  {
	cmd := exec.Command("../child/child")
	// パイプを作る
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("parent: StdoutPipe error: ", err.Error())
		os.Exit(1)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("parent: StderrPipe error: ", err.Error())
		os.Exit(1)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	log.Println("parent: child exec success: ", cmd.Process.Pid)

	streamReader := func(scanner *bufio.Scanner, outputChan chan string, doneChan chan bool) {
		defer close(outputChan)
		defer close(doneChan)
		for scanner.Scan() {
			outputChan <- scanner.Text()
		}
		doneChan <- true
	}

	// stdout, stderrをひろうgoroutineを起動
	stdoutScanner := bufio.NewScanner(stdout)
	stdoutOutputChan := make(chan string)
	stdoutDoneChan := make(chan bool)
	stderrScanner := bufio.NewScanner(stderr)
	stderrOutputChan := make(chan string)
	stderrDoneChan := make(chan bool)
	go streamReader(stdoutScanner, stdoutOutputChan, stdoutDoneChan)
	go streamReader(stderrScanner, stderrOutputChan, stderrDoneChan)

	// channel経由でデータを引っこ抜く
	stillGoing := true
	for stillGoing {
		select {
		case <-stdoutDoneChan:
			stillGoing = false
		case line := <-stdoutOutputChan:
			log.Println("parent:", line)
			switch line {
			case "ready":
				if err := cmd.Process.Signal(os.Interrupt); err != nil {
					log.Println("parent: signal error: ", err)
					os.Exit(1)
				}
			case "finish":
				state, err := cmd.Process.Wait()
				if err != nil {
					log.Println("parent: child wait error: ", err)
					os.Exit(1)
				}
				log.Println("parent: success: ", state)
				os.Exit(0)
			}
		case line := <-stderrOutputChan:
			log.Println("parent: ", line)
		}
	}
}
