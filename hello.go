//Go programs are organized in packages. Main package is the entry point of program, also identifies an executable program.
package main

//fmt is a built-in go package. provides input/output functions.
import (
	"fmt"
	"strings"
)

//where program starts, as mentioned above.
func main() {
	fmt.Println("Hello, world!")

	/*
		variables
		-can be defined at package or function level.
		-at package level, the variable becomes visible across all files that use that package.
		-at func level, the variable is visible only within the function.
	*/
	const age = 28 //this method of declaring the variable makes Go determine the type. in this case, is int.
	fmt.Println(age)
	/*
	 can declare variable without value, but must specify the type.
	 if the variable is not assigned a value, it will default to whatever type is specified. i.e. int = 0, string = ""
	 var name string
	 name = "Ryan"
	*/
	var name string = "Ryan"
	fmt.Println(name)
	dogName, breed := "Rocky", "Jack Russell-Chihuahua mix" //short variable declaration mixed with multiple variable declarations on one line.
	fmt.Println(dogName)
	fmt.Println(breed)

	//note - unused variables will throw an error and the program will not be able to compile.

	fmt.Println(len(name))
	fmt.Println("------- strings -------")
	//can assign previous string to new variable. however, strings are immutable, so the string value cannot be updated. example below.
	first := "test"
	second := first

	first = "another test"
	fmt.Println(first)
	fmt.Println(second)

	//concatenating strings
	var word = first + " " + second
	fmt.Println(word)
	fmt.Println(strings.HasPrefix("test", "te"))

	lowercaseFullName := "ryan soto"
	uppercaseFullName := strings.ToUpper(lowercaseFullName)

	fmt.Println(lowercaseFullName)
	fmt.Println(uppercaseFullName)
}
