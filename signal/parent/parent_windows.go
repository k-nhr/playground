package main

import (
	"bufio"
	"golang.org/x/sys/windows"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main()  {
	cmd := exec.Command("../child/child.exe")
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
				if err := terminateProc(cmd.Process.Pid); err != nil {
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

func terminateProc(pid int) error {
	dll, err := windows.LoadDLL("kernel32.dll")
	if err != nil {
		return err
	}
	defer dll.Release()

	f, err := dll.FindProc("AttachConsole")
	if err != nil {
		return err
	}
	r1, _, err := f.Call(uintptr(pid))
	if r1 == 0 && err != syscall.ERROR_ACCESS_DENIED {
		return err
	}

	f, err = dll.FindProc("SetConsoleCtrlHandler")
	if err != nil {
		return err
	}
	r1, _, err = f.Call(0, 1)
	if r1 == 0 {
		return err
	}
	f, err = dll.FindProc("GenerateConsoleCtrlEvent")
	if err != nil {
		return err
	}
	r1, _, err = f.Call(windows.CTRL_BREAK_EVENT, uintptr(pid))
	if r1 == 0 {
		return err
	}
	return nil
}