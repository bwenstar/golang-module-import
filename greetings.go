package main

import (
    "fmt"

    hello "github.com/bwenstar/golang-module-test"
)

func main() {
    message := hello.Hello("Brendan")
    fmt.Println(message)
}
