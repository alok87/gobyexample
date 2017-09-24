package main

import "fmt"

type Product struct {
    name string
    repo string
    container string
    environment string
}

func main() {
    k := Product{name: "myproduct", repo: "github.com/alok87/myrepo"}
    fmt.Println(k)
    r := &Product{name: "myprod1", repo: "github.com/alok87/prod1"}
    fmt.Println(r.name)
    r.name = "prod1"
    fmt.Println(r.name)
    fmt.Println(r)
}
