package main

import "fmt"

func main() {
    products := make(map[string]map[string]string)
    kuber := make(map[string]string)
    kuber["repo"] = "github.com/practo/kuber"
    kuber["admins"] = "alok.singh@practo.com"
    products["kuber"] = kuber
    fmt.Println(products)
    fmt.Println(products["kuber"]["admins"])

    fmt.Println(len(products), len(kuber))
    delete(kuber, "repo")
    fmt.Println(kuber)

    var m = map[int]string{2: "two"}
    fmt.Println(m)
}
