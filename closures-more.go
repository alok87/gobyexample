package main

import "fmt"

func outer(s string) {
    text := "my name is " + s
    inner := func() {
        fmt.Println(text)
    }

    inner()
}


func outer2(s string) func() {
    text := "my name2 is " + s
    return func() {
        fmt.Println(text)
    }
}

func closure(s string) func() string {
    return func() string {
        s = s + "alok" + " "
        return s
    }
}

func main() {
    // 1
    outer("alok")

    // 2
    inner := outer2("vicky")
    inner()
    // or 2
    outer2("vicky")()

    //retains the values
    c := closure("start")
    fmt.Println(c())
    fmt.Println(c())
    fmt.Println(c())

    d := closure("end")
    fmt.Println(d())
    fmt.Println(d())
    fmt.Println(d())

}
