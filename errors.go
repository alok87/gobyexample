package main

import (
    "fmt"
)

type ValueError struct {
    value int
}

func (v *ValueError) Error() string {
    return fmt.Sprintf("Error! value not allowed: %d", v.value)
}

func checkVal(n int) (bool, error) {
    if n == 5 {
        err := &ValueError{n}
        return false, err
    } else {
        return true, nil
    }
}

func main() {
    _, err := checkVal(5)
    if err != nil {
        fmt.Println("Failed", err)
    }
    fmt.Println("Passed")
    _, err2 := checkVal(6)
    if err2 != nil {
        fmt.Println("Failed", err)
    }
    fmt.Println("Passed")
}
