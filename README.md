# week-1-homework-1

- Going all over what has been studied including code examples in the class. Do not forget to commit the initial and final work.

- Please read Ch02 and Ch03 of [GoPL book](https://drive.google.com/file/d/1kvsEfCuOYecBrfy12tTI1kDDbC4e4AVy/view?usp=sharing) and run its code. Do not forget to commit the initial and final work.



<br>My HW1 comments for Week 1<br>

                                                          --------> CH.2 Declarations <--------

<br>// <b>In this code a constant is defined to denote temperature in Fahrenheit and converted to Celcius with basic math operations.</b>
<br>package main<br>
<br>import "fmt"<br>
<br>const boilingF = 212.0<br>
<br>func main() {<br>
<br>var f = boilingF<br>
<br>var c = (f - 32) * 5 / 9<br>
<br>fmt.Printf("boiling point = %g°F or %g°C\n", f, c)<br>

<br>}<br>
<br>// In this second example, both freezing and boiling point of water in sea level conditions defined in Fahrenheit and converted to Celcius similar to first example.
<br>package main<br>
<br>import "fmt"<br>
<br>func main() {<br>
<br>const freezingF, boilingF = 32.0, 212.0<br>
<br>fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"<br>
<br>fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF)) // "212°F = 100°C"<br>
<br>}
<br>func fToC(f float64) float64 {<br>
<br>return (f - 32) * 5 / 9<br>
<br>}<br>

<br> After that possible variable definiton methods are discussed and pointers are explained in depth.<br>
 
 <br>// In here, a Flag variable is defined to get 1 boolean and 2 string value.<br>
<br>package main<br>
<br>import (<br>
<br>"flag"<br>
<br>"fmt"<br>
<br>"strings"<br>
<br>)
<br>var n = flag.Bool("n", false, "omit trailing newline")<br>
<br>var sep = flag.String("s", " ", "separator")<br>
<br>func main() {<br>
<br>flag.Parse()<br>
<br>fmt.Print(strings.Join(flag.Args(), *sep))<br>
<br>if !*n {<br>
<br>fmt.Println()<br>
<br>}<br>
<br>}<br>


                                                          --------> CH.3 Data Types <--------
                                                          
                                                          
  //CH.3 mainly dealt with data types and their declarations, for my HW, I will continue with in-class exercises. I changed some of the variables to play with code and understand it further.
  
  
<br>package main<br>
<br>import (<br>
<br>"fmt"<br>
<br>)<br>
<br>type Person struct {<br>
<br>name string<br>
<br>}<br>
<br>func (p Person) greet() string {<br>
<br>return "Selam " + g.name + " :)"<br>
<br>}<br>
<br>func main() {<br>
<br>greetPrinter(createGreetInTurkish, "Hatice")<br>
<br>greetPrinter(createGreetInEnglish, "Mary")<br>
<br>greetPrinter(convertToUppercase, "naber?")<br>
<br>greetCreator := createGreetInTurkish<br>
<br>greetPrinter(greetCreator, "Toprak")<br>
<br>func(name string) {<br>
<br>greeting := "Merhaba " + name + " :)"<br>
<br>fmt.Printf("%s\n", greeting)<br>
<br>}<br>
<br>("Yesim")<br>
<br>closure := func(name string) {<br>
<br>greeting := "Merhaba " + name + " :)"<br>
<br>fmt.Printf("%s\n", greeting)<br>
<br>}<br>
<br>closure("Fatma")<br>
<br>anotherGreetPrinter(closure, "Zeynep")<br>
<br>}<br>
<br>func createGreetInTurkish(name string) string {<br>
<br>return "Selam " + name + " :)"<br>
<br>}<br>
<br>func createGreetInEnglish(name string) string {<br>
<br>return "Hi " + name + " :)"<br>
<br>}<br>
<br>func convertToUppercase(arg string) string {<br>
<br>return strings.ToUpper(arg)<br>
<br>}<br>
<br>func greetPrinter(function func(it string) string, name string){<br>
<br>var greeting = function(name)<br>
<br>fmt.Printf("%s\n", greeting)<br>
<br>}<br>
<br>func anotherGreetPrinter(function func(it string), name string){<br>
<br>function(name)<br>
<br>}<br>
