package main

import (
	"os"
    "fmt"
    "time"
    "syscall"
	"os/signal"
)

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() (stopCh <-chan struct{}) {
    fmt.Println("Clsoing OnlyOneSingalHandler")
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
    fmt.Println("Made stop channel chan struct{}")
	c := make(chan os.Signal, 2)
    fmt.Println("Made os.Signal channel")
	signal.Notify(c, shutdownSignals...)
    fmt.Println("signal.Notify called")
	go func() {
        fmt.Println("Goroutine: waiting for signal_1")
        cx := <-c
        fmt.Println("Received signal_1", cx)
        fmt.Println("Closing channel", stop)
		close(stop)
        fmt.Println("Goroutine: waiting for signal_2")
        cy := <-c
        fmt.Println("Received signal_2, will exit", cy)
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

func main() {
	fmt.Println("Setting up signal handlers")
	stopCh := SetupSignalHandler()
    time.Sleep(time.Second * 100)
    fmt.Println(stopCh)
}
