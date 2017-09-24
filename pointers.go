package main

import "fmt"

func byval(val int) {
    val = 0
    fmt.Println(&val)
}

func byref(val *int) {
    *val = 0
    fmt.Println(&(*val))
}

func main() {
    v := 10
    fmt.Println("address of v", &v)
    byval(v)
    fmt.Println("after pass by value", v)
    byref(&v)
    fmt.Println("after pass by ref", v)
}
