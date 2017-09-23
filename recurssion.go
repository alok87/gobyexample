package main

import "fmt"


func fact(n int) int {
    if n != 1 {
        return n * fact(n-1)
    } else {
        return 1
    }
}

func main() {
    fmt.Println("Factorial of 5=", fact(5))
}
