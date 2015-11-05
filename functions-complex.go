/*
 * Functions can:
 * - return multiple values
 * - can be used like any other type
 * - function literals (anonymous functions or closure)
 */


package main

import "fmt"

//User defined struct
type Salutation struct{
	name string
	greeting string
}

func Greet(s Salutation, do func (string)) {
	message, alternate := CreateMessage(s.name, s.greeting)
	do (message) 
	do (alternate)
}

func CreateMessage(name, greeting string) ( message string, alternate string) {
    message = greeting + " " + name + "!"
	alternate = "Hey!" + name
	return
}

//Variadic Functions
func  GreetVariatic(s Salutation, do func (...string)){
	message, alternate := CreateMessageVariatic(s.name, s.greeting, "yo")
	do(message, alternate)
}

func CreateMessageVariatic(name string, greeting ...string) ( message string, alternate string) {
	
	fmt.Println(len(greeting))

    message = greeting[1] + " " + name + "!"
	alternate = "Hey!" + name
	return
}


func Print(s string){
	fmt.Println("Print Function \n")
	fmt.Print(s)
}

func Println(s ...string){
	fmt.Println("Println Variadic Function")
	fmt.Println(s[0])
}


func main() {
	
	//Struct declaration
	var s = Salutation{ name: "Ronald", greeting: "Hello" }
	Greet(s, Print)
	GreetVariatic(s, Println)
}
