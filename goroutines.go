package main

import "fmt"

func scorePlz(m string) {
    for i:=0; i<3; i++ {
        fmt.Println(m + " "  + "235-2")
    }
}

func main() {
    scorePlz("Ind vs Aus")
    go scorePlz("async Ind vs SL")
    go func () {
        scorePlz("async Ind vs NZ")
    }()
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
