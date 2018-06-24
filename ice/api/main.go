package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ice-judge/ICE/ice/api/server"
)

func main() {
	config := parseConfig()
	srv, err := server.New(config)
	if err != nil {
		panic(err)
	}

	srv.Open()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	srv.Stop()
}
