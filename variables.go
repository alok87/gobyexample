package main

import "fmt"

func main() {

    var a int
    fmt.Println("without initialization", a)

    var b int = 5
    fmt.Println("int", b)

    var c = true
    fmt.Println("boolean", c)

    var d string = "hello string"
    fmt.Println("string", d)

    var e bool
    fmt.Println(e)

    x := "hello"
    fmt.Println("Interpret type", x)

}
