package main

import "fmt"

func main() {
    var a [5]int
    fmt.Println(a)

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println(b)
    fmt.Println(b[0])
    b[1] = 99 
    fmt.Println(b)

    var twoD [3][2]int
    for i:=0; i<3; i++ {
        for j:=0; j<2; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println(twoD)
    
}
