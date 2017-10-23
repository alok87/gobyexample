package main

import (
    "io"
    "net/http"
)

type HelloWorld struct {
    name string
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello " + h.name)
}

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hey jude, don't be sad, take a sad song and make it better")
}

func main() {
    // create a custom request multiplexer
    mux := http.NewServeMux()
    mux.Handle("/hello/alok", &HelloWorld{"Alok"})
    mux.Handle("/hello/vicky", &HelloWorld{"Vicky"})
    mux.HandleFunc("/", hello)
    http.ListenAndServe(":8080", mux)
}
