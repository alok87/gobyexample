package main

import "fmt"

func startAt(a int) func(int) int {
    return func(b int) int {
        return b + a 
    }
}

func intSeq() func() int {
    i := 0
    fmt.Println("i=", i)
    return func() int {
        i += 1
        fmt.Println("i+1=", i)
        return i
    }
}

func main() {
    closure_1 := startAt(5)
    closure_2 := startAt(6)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

    fmt.Println(closure_1(1))
    fmt.Println(closure_1(2))
    fmt.Println(closure_1(2))
    fmt.Println(closure_2(6))
}
