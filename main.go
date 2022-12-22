package main

import (
	"daily-cute-dogs-email/sender"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go sender.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
	fmt.Println("ending of execution")
	os.Exit(0)
}
