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
	fmt.Println("------- variables -------")
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

	fmt.Println("------- arrays -------")
	// "[...]" counts array items for you.
	// arrays cannot be resized, as you have to explicitly define length of array in Go
	// arrays can only contain values of the same types.
	var myArray = [...]string{"First", "Second", "Third"}
	fmt.Println(myArray)

	myArrayCopy := myArray
	myArray[2] = "Tuesday"

	println(myArray[2])
	println(myArrayCopy[2])

	fmt.Println("------- slices -------")
	//slices are a similar data structure to arrays, big difference is that slices can change in size.
	//to define a slice, it's the same format as arrays, just without the specified length.
	var days = []string{"Monday", "Tuesday", "Wednesday"}
	fmt.Println(days)

	months := []string{"Jan", "Feb", "Mar"}
	fmt.Println(months)

	//can create an empty slice with specified length with the make() func
	emptySlice := make([]string, 3)
	fmt.Println(emptySlice)

	//can add values to slice
	months = append(months, "Apr", "May", "Jun")
	fmt.Println(months)

	//copy(newSlice, originalSlice) duplicates slice, but does not share same memory
	//the examples below share the same memory, so modifying one slice updates the underlying array and causes the other slices generated to also be modified.
	firstCopy := [3]string{"one", "two", "three"}
	secondCopy := firstCopy[:]
	thirdCopy := firstCopy[:]

	fmt.Println(firstCopy)
	fmt.Println(secondCopy)

	firstCopy[0] = "one million"

	fmt.Println(firstCopy)
	fmt.Println(thirdCopy)

	//can specify capacity with third parameter in make func
	capacityExample := make([]string, 0, 10)
	fmt.Println(capacityExample)

	capacityExample = append(capacityExample, "one", "two", "three", "four", "five", "six")
	lastThreeWords := capacityExample[2:5]

	fmt.Println(lastThreeWords)

	fmt.Println("------- maps -------")
	//other languages refer to maps as a dictionary, hash map or associative array.
	//do not need to specify how many items a map can hold.
	agesMap := make(map[string]int)

	agesMap["ryan"] = 28

	fmt.Println(agesMap)

	delete(agesMap, "ryan")

	fmt.Println(agesMap)

	agesMapV2 := map[string]int{"ryan": 29, "rocky": 10, "ariana": 27}
	fmt.Println(agesMapV2)

	fmt.Println("------- loops -------")
	//go only has the for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	monthsMap := map[int]string{1: "January", 2: "Febuary", 3: "March", 4: "April"}

	//loop through monthsMap
	for month := 1; month < len(monthsMap); month++ {
		fmt.Println("The " + fmt.Sprintf("%d", month) + " is " + monthsMap[month])
	}

	//however, Go can simulate a while loop:
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	//can omit condition and use break to end loop:
	p := 20

	for p > i {
		fmt.Println(p)

		if p < 15 {
			break
		}

		p--
	}

	//using range in for loops:
	numbers := []int{1, 2, 3}

	for i, num := range numbers {
		fmt.Printf("%d: %d\n", i, num)
	}

	fmt.Println("------- conditionals -------")
	myAge := 28

	if myAge < 30 {
		fmt.Println("still in the clear")
	} else if myAge > 25 {
		fmt.Println("getting close...")
	} else {
		fmt.Println("well then...")
	}

	//switch statement:
	switch myAge {
	case 10:
		fmt.Println("10 years old")
	case 20:
		fmt.Println("20 years old")
	case 28:
		fmt.Println("28 years old??")
	}

	fmt.Println("------- operators -------")
	//very similar to javascript and other languages
	//however, can no longer prepend unary operators to numbers like JS
	// i.e. ++num

	fmt.Println("------- structs -------")
	//struct: type that contains one or more variables. we call them fields.
	//example:
	type Person struct {
		Name string //important to use uppercase for the fields. if not uppercase, those fields are now private to the package,
		Age  int    //and whenthe struct is passed to a function provided by another package, those fields cannot be accessed.
	}

	//initialize variable with that type after struct is defined:
	ryan := Person{"ryan", 28}

	fmt.Println(ryan)
	fmt.Println(ryan.Age)

	randomPerson := Person{Age: 100}
	fmt.Println(randomPerson)

	//with structs, you can use one struct inside of another
	type fullName struct {
		firstName string
		lastName  string
	}

	me := fullName{"ryan", "soto"}

	type Occupation struct {
		Name  fullName
		Title string
	}

	developer := Occupation{Name: me, Title: "Go Developer"}
	fmt.Println(developer)
}
