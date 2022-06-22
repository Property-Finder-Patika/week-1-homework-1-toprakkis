// Greeting go file
package main
import (
"fmt"
)
type Person struct {
name string
}
func (p Person) greet() string {
return "Selam " + g.name + " :)"
}
func main() {
greeter := Person{"Toprak Kış"}
var greeting = greeter.greet()
fmt.Printf("%s\n", greeting)
}
