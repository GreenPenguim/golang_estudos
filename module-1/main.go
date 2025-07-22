package main

import "fmt"

func main(){
	// It is possible to create variables with no type defined
	var name = "Houdini"
	var age = 20
	var version float32 = 1.0
	fmt.Println("Name:", name,",", "Age:", age, ",", "Version:", version)
	fmt.Printf("\n")

	/*
		Array is one data type with fixed size.

		Slice is a part of an array, unique data and dynamic size 
		
		A map is a collection of key-value pairs, a.K.a: hash tables, associative arrays, dictionaries.
	*/
	var arr [7]int
	arr[4] = 35
	fmt.Println("Array:", arr)
	fmt.Printf("\n")
	params(10)

	// Structs are used to create complex data types.
	type Person struct {
		stName string
		ndName string
		age int
	}
	p1 := Person{"Houdini","Silta", 20} 

	fmt.Println("Pessoa:", p1.stName, p1.ndName, p1.age)

	type CrazyString interface {
		String() string
	}

	temp()
	tempShortOperator()
}

func params(p1 int){
	fmt.Println("Parameter:", p1)
	fmt.Printf("\n")
}

const ebulitionF float64 = 212.00

func temp(){
	var tempF float64 = ebulitionF
	var tempC float64 = ((tempF - 32) * 5) / 9 

	fmt.Println("Temperature of water in Fahrenheit:", tempF)
	fmt.Println("Temperature of water in Celsius:", tempC)
	fmt.Printf("\n")
}

func tempShortOperator(){
	/*
		Short operator :=        <- looks like gopher :F

		Is used to declare and initialize a variable in one step.
		It is only used inside functions, not at the package level.

	*/

	tempF := ebulitionF
	tempC := ((tempF - 32) * 5) / 9 
	
	// Using the Printf
	// %g is used to format the float without unnecessary decimal places
	fmt.Println("Using the short operator :=")
	fmt.Printf("Temperature of water in Fahrenheit: %g \nTemperature of water in Celsius: %g\n\n", tempF, tempC)

	// %T reveals the type of the variable
	fmt.Printf("Temperature of water in Fahrenheit: %g(%T) \nTemperature of water in Celsius: %g(%T)\n", tempF,tempF, tempC,tempC)
}

func dataConversion(){
	var1 := 10
	var2 := "John Doe"

	var var3 float64 = float64(var1)
	fmt.Printf("var3 value: %g \n var2 value: %s", var3, var2)

}