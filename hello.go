package main

import (
	"fmt"
	"os"
	//"runtime"
	"reflect"
)
/**
 *  Package level variables scope
 */
var (
	name = "Pluralsight"
	course = "Docker Deep Dive"
	module = 3.2
)


/**
 * [main description]
 * @return {[type]} [description]
 */
func main() {
  
	// a := 1.0
	// b := 2
	// ptr := &course
	
	user := os.Getenv("USER")
	fmt.Println("user ", user,"and is type of ", reflect.TypeOf(user))

	// fmt.Println("Hello from", runtime.GOOS)
	// fmt.Println("module", module,"but is type of ", reflect.TypeOf(module))
	// fmt.Println("Sum a + b ", int(a) + b)
	// fmt.Println("memory address of (course) is ", ptr, "and the value is ", *ptr)


	fmt.Println("(main) Course memory address is", &course) 
	changeCourse(&course)
	fmt.Println("(main) Course is ", course)
}

/**
 * [changeCourse description]
 * @param  {[type]} course string        [description]
 * @return {[type]}        [description]
 */
func changeCourse(course *string) string {
	
	fmt.Println("(changeCourse) passed value", *course)

	*course = "Docker Fundamentals"

	fmt.Println("(changeCourse) Course memory address is", course)
	fmt.Println("(changeCourse) New value", *course)

	return *course
}