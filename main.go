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

	//print output
	fmt.Println(fullName, "i am", age, "years old")
	fmt.Println("I live in", country)
	fmt.Println("My cgpa is ", cgpa)
}
