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

	fmt.Println("Graceful shutdown program, it will run for 10 seconds.")
	fmt.Println("It can be completed with ctrl + C. Upon successful completion, it will display")
	fmt.Println("\"Goodbye world\". If you close the program forcibly the program will display")
	fmt.Println("\"Stopped by the user after x seconds\"")

	fmt.Println("Hello World")

	select {
	case <-timer:
		fmt.Println("Goodbye world")
	case <-sigChan:
		exit := time.Now().Sub(start).Seconds()
		fmt.Printf("Stopped by the user after %f seconds", exit)
	}
}
