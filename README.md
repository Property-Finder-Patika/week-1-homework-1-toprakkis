# week-1-homework-1

- Going all over what has been studied including code examples in the class. Do not forget to commit the initial and final work.

- Please read Ch02 and Ch03 of [GoPL book](https://drive.google.com/file/d/1kvsEfCuOYecBrfy12tTI1kDDbC4e4AVy/view?usp=sharing) and run its code. Do not forget to commit the initial and final work.



My HW1 comments for Week 1

                                                          --------> CH.2 Declarations <--------

// In this code a constant is defined to denote temperature in Fahrenheit and converted to Celcius with basic math operations.
package main\n
import "fmt"
const boilingF = 212.0
func main() {
var f = boilingF
var c = (f - 32) * 5 / 9
fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

}
// In this second example, both freezing and boiling point of water in sea level conditions defined in Fahrenheit and converted to Celcius similar to first example.
package main
import "fmt"
func main() {
const freezingF, boilingF = 32.0, 212.0
fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF)) // "212°F = 100°C"
}
func fToC(f float64) float64 {
return (f - 32) * 5 / 9
}

 After that possible variable definiton methods are discussed and pointers are explained in depth.
 
 // In here, a Flag variable is defined to get 1 boolean and 2 string value.
package main
import (
"flag"
"fmt"
"strings"
)
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
func main() {
flag.Parse()
fmt.Print(strings.Join(flag.Args(), *sep))
if !*n {
fmt.Println()
}
}


                                                          --------> CH.3 Data Types <--------

