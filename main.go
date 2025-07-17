package main

import "fmt"

func main(){
	// It is possible to create variables with no type defined
	var name string = "Houdini"
	var age int = 20
	var version float32 = 1.0
	fmt.Println("Name:", name,",", "Age:", age, ",", "Version:", version,)
	
	// Array is one data type with fixed size.
	var arr [7]int
	arr[4] = 35
	fmt.Println("Array:", arr)
	// Slice is a part of an array, unique data and dynamic size 
	// A map is a collection of key-value pairs, a.K.a: hash tables, associative arrays, dictionaries.

	params(10)

	// Structs are used to create complex data types.
	type Person struct {
		stName string
		ndName string
		Age int
	}
	p1 := Person{"Houdini","Silta", 20} 

	fmt.Println("Pessoa:", p1.stName, p1.ndName, p1.Age)

	type CrazyString interface {
		String() string
	}

	temp()
	/*
		Short operator :=        <- looks like gopher :F

		Is used to declare and initialize a variable in one step.
		It is only used inside functions, not at the package level.

	*/

	tempShortOperator()


}

func params(p1 int){
	fmt.Println("Parameter:", p1)
}


const ebulitionF float64 = 212.00

func temp(){
	var tempF float64 = ebulitionF
	var tempC float64 = ((tempF - 32) * 5) / 9 

	fmt.Println("Temperature of water in Fahrenheit:", tempF)
	fmt.Println("Temperature of water in Celsius:", tempC)
}

func tempShortOperator(){
	tempF := ebulitionF
	tempC := ((tempF - 32) * 5) / 9 

	fmt.Println("Temperature of water in Fahrenheit:", tempF)
	fmt.Println("Temperature of water in Celsius:", tempC)
}