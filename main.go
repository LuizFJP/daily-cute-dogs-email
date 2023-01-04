package main

import (
	"daily-cute-dogs-email/backend/sender"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"daily-cute-dogs-email/backend/api"
)

func main() {
	go api.Start()
	go sender.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
	fmt.Println("ending of execution")
	os.Exit(0)
}
