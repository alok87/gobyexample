package main

import "fmt"

type product struct {
    name string
    repo string
}

func (p *product) setName(s string) {
    p.name = s
}

func (p * product) setRepo(s string) {
    p.repo = s
}

func main() {
    p := product{}
    p.setName("product1")
    p.setRepo("repo1")
    fmt.Println(p)

    np := &p
    fmt.Println(np.name)
    fmt.Println(np.repo)
}
