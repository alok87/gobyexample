package main

import "fmt"

func main() {

    slice := make([]string, 3)
    slice[0] = "a"
    slice[1] = "b"
    slice[2] = "c"
    fmt.Println("slice=", slice)
    fmt.Println("len(slice)=", len(slice))

    newSlice := make([]string, len(slice))
    copy(newSlice, slice)
    fmt.Println("newSlice=", newSlice)
    newSlice = append(newSlice, "d")
    fmt.Println("newSlice=", newSlice)

    twoD := make([][]int, 3)
    for i:=0; i<3; i++ {
        in := i + 1
        twoD[i] = make([]int, in)
        for j:=0; j<in; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println(twoD)
}
