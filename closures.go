package main

import "fmt"

func startAt(a int) func(int) int {
    return func(b int) int {
        return b * (a + 1)
    }
}

func main() {
    closure_1 := startAt(5)
    closure_2 := startAt(6)
    fmt.Println(closure_1(0))
    fmt.Println(closure_2(6))
}
