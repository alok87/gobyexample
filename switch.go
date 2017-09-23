package main

import "fmt"
import "time"

func main() {
    switch time.Now().Weekday() {
        case time.Saturday:
            fmt.Println("Njoy", time.Now().Weekday())
        case time.Sunday:
            fmt.Println("Make chicken", time.Now().Weekday())
        default:
            fmt.Println("Work hard on", time.Now().Weekday())
    }
}
