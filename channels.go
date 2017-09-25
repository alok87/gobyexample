package main

import "fmt"

func main() {

    fmt.Println("Inside main thread")

    channel := make(chan string)

    go func () {
        channel <- "Hello from annonymous go routine"
    }()

    msg_received := <-channel
    fmt.Println("Receivied message in the main thread:", msg_received)
}
