package main

import "fmt"

func main(){
	types()
	arrays()
	params(10)
}

func types(){
	// It is possible to create variables with no type defined
	var name string = "Houdini"
	var age int = 20
	var version float32 = 1.0
	fmt.Println("Name:", name,",", "Age:", age, ",", "Version:", version,)
}

func arrays(){
	// One data type with fixed size.
	var arr [7]int
	arr[4] = 35
	fmt.Println("Array:", arr)
}

func slice(){
// A part of an array, unique data and dynamic size 
}

func maP(){
// A map is a collection of key-value pairs, a.K.a: hash tables, associative arrays, dictionaries.
}

func params(p1 int){
	fmt.Println("Parameter:", p1)
}

func structure(){
	// Structs are used to create complex data types.
	type Person struct {
		stName string
		ndName string
		Age int
	}
	p1 := Person{"Houdini","Silta", 20} 

	fmt.Println("Pessoa:", p1.stName, p1.ndName, p1.Age)
}

func interfaces(){
	type CrazyString interface {
		String() string
	}
}

