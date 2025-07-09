// Secure.go
package main

import "fmt"

type Greeter struct {
    name string
}

func (g Greeter) Greet() {
    fmt.Printf("Hello from %s in Go!\n", g.name)
}

func main() {
    greeter := Greeter{name: "Secure"}
    greeter.Greet()
}
