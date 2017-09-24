package main

import "fmt"

type Animal interface {
    Speak() string
}

type Cat struct {
}

func (c *Cat) Speak() string {
    return "Meow"
}

type Dog struct {
}

func (d *Dog) Speak() string {
    return "Bhow"
}

func main() {
    animals := []Animal{&Cat{}, &Dog{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
