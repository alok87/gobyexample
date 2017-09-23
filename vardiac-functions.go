package main

import "fmt"

func vardiac(nums ...int) int {
    sum :=0
    for _, i := range nums {
        sum += i
    }
    return sum
}

func main() {
    fmt.Println(vardiac(1, 2, 3))
}
