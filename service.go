package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
)

var logger service.Logger

type rdService struct {
	exit chan struct{}
}

func (r *rdService) run() error {

	logger.Info("RD-Service Start !!!")

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			bot()
		case <-r.exit:
			ticker.Stop()
			logger.Info("RD-Service Stop ...")
			return nil
		}
	}
}

func (r *rdService) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	r.exit = make(chan struct{})

	go r.run()
	return nil
}

func (r *rdService) Stop(s service.Service) error {
	close(r.exit)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "RD-Service",
		DisplayName: "RD-Service",
		Description: "This is remote dakoku service.",
	}

	// Create RD-Service service
	program := &rdService{}
	s, err := service.New(program, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Setup the logger
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal()
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			fmt.Printf("Failed (%s) : %s\n", os.Args[1], err)
		} else {
			fmt.Printf("Succeeded (%s)\n", os.Args[1])
		}
		return
	}

	// run in terminal
	s.Run()
}
