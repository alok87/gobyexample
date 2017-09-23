package main

import "fmt"

func main() {
    for i:=0; i < 5; i ++ {
        fmt.Println("i=", i)
    }

    j := 0
    for j < 10 {
        fmt.Println("j", j)
        j += 5
    }
}
