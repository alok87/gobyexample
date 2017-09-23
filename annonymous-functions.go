package main

import "fmt"

func named_func(s string) {
    fmt.Println(s)
}

func return_anonymous_func() func(string) {
    return func(s string) {
        fmt.Println(s)
    }
}

func main() {

    //named function
    named_func("hello")

    //anonymous function
    func (s string) {
        fmt.Println(s)
    }("hello anonymous function")
    
    //gets anonymous funcition
    ann_func := return_anonymous_func()
    ann_func("hello ann function")

}
