package main

import (
    "fmt"
)

func callOut() int {
    fmt.Println("Outside is beinge executed")
    return 1
}

var test = callOut()

func init() {
    fmt.Println("Init3 is being executed")
}

func init() {
    fmt.Println("Init is being executed")
}

func init() {
    fmt.Println("Init2 is being executed")
}

func main() {
    fmt.Println("Do your thing !")
}
