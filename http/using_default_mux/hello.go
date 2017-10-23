package main

import (
    "fmt"
    "net/http"
    "log"
)

type HelloWorld struct {
    name string
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello " + h.name)
}

func main() {
    http.Handle("/hello/alok", &HelloWorld{"Alok"})
    http.Handle("/hello/vicky", &HelloWorld{"Vicky"})

    // start web server
    log.Fatal(http.ListenAndServe(":8080", nil))
}
