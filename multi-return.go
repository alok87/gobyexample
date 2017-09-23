package main

import "fmt"

func add(a int, b int) (int, error) {
    return a + b, nil
}

func main() {
    sum, err := add(2, 3)
    fmt.Println("sum=", sum)
    fmt.Println(err)
}
