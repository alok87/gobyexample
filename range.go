package main

import "fmt"

func main() {

    fmt.Println("Range over slices/arrays")
    arr := [5]int{1, 2, 3, 4, 5}
    slice := []int{5, 6, 7, 8 , 9}
    for i, av  := range arr {
        fmt.Println(i, av)
    }
    for _, sv := range slice {
        fmt.Println(sv)
    }

    fmt.Println("Range over maps")
    dict := make(map[string]int)
    dict["a"] = 1
    dict["z"] = 26
    for key := range dict {
        fmt.Println(key)
    }

    for key, value := range dict {
        fmt.Println(key, value)
    }
}
