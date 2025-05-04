package main

import "fmt"

func main() {

	/*
		Variable Naming Convention

		1. letter,digit,_.
		2. keywords are not allowed as variable.
		3. variable can not have space.
		4. variable name can not start with digit.
	*/

	// Variable Declaration Convention

	var fullName, country string
	var age int
	var cgpa float32

	//Variable Initialization Convention

	// fullName = "Atikur Rahman"
	//  country = "Bangladesh"
	//  age = 25
	//  cgpa = 3.15

	//get input
	fmt.Print("Enter your name")
	fmt.Scan(&fullName)

	fmt.Print("Enter your age")
	fmt.Scan(&age)

	fmt.Print("Enter your country")
	fmt.Scan(&country)

	// Using 8-bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3)

	// Using 16-bit signed int
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)

	var x int16 = 170
	var y int16 = 83
	//Addition
	fmt.Printf(" addition :  %d + %d = %d\n ", x, y, x+y)
	//Subtraction
	fmt.Printf("subtraction : %d - %d = %d\n", x, y, x-y)
	//Multiplication
	fmt.Printf(" multiplication : %d * %d = %d\n", x, y, x*y)
	//Division
	fmt.Printf(" division : %d / %d = %d\n", x, y, x/y)
	//Modulus
	fmt.Printf(" remainder : %d %% %d = %d\n", x, y, x%y)

	//print output
	fmt.Println(fullName, "i am", age, "years old")
	fmt.Println("I live in", country)
	fmt.Println("My cgpa is ", cgpa)
}
