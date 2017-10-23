package main

import (
    "io"
    "net/http"
)

// custom handler for the routes, it implements Handler interface
type HelloWorldHandler struct {
    name string
}

func (h *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if fn, ok := mux[r.URL.String()]; ok {
        fn(w, r)
        return
    }
    io.WriteString(w, "Hello Default")
}

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hey jude, don't be sad, take a sad song and make it better")
}

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func main() {

    // create custom server, instead of defaultServer
    // FYI: default server is defined here: https://golang.org/src/net/http/server.go?s=87921:87976#L2870
    server := &http.Server{
        Addr: ":8080",
        Handler: &HelloWorldHandler{},
    }

    // create a custom request multiplexer, instead of defaultServeMux
    // previously: mux.Handle("/hello/alok", &HelloWorldHandler{"Alok"})
    // previously: mux.Handle("/hello/vicky", &HelloWorldHandler{"Vicky"})
    // previously: mux.HandleFunc("/", hello)
    mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
    mux["/"] = hello
    mux["/hello/alok"] = hello
    mux["/hello/vicky"] = hello

    // start the server
    server.ListenAndServe()
}
