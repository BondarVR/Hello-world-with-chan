package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gracefulShutdown()
}

func gracefulShutdown() {
	start := time.Now()
	timer := time.After(10 * time.Second)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	fmt.Println("Hello World")

	select {
	case <-timer:
		fmt.Println("Goodbye world")
	case <-sigChan:
		exit := time.Now().Sub(start).Seconds()
		fmt.Printf("Stopped by the user after %f seconds", exit)
	}
}
