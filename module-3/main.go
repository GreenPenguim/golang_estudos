package main

import (
	"fmt"
)

/*
	Function -> a.K.a.: Method (linked to POO), Subroutine, Procedure

	Function is a block of code that can be called by name, and can receive parameters and return values.
*/

func main(){
	list := []float64{98, 93, 77, 82, 83}
	fmt.Println("Functions")
	fmt.Printf("Average: %f \n", average(list))
	fmt.Println("-----")

	fmt.Println("Clousure and Recursion")
	fmt.Printf("Clousure:\n")
	clousure()
	fmt.Println("::::")
	fmt.Printf("Recursion:\n")
	fmt.Println(factorial(9))
	fmt.Println("-----")

	fmt.Println("Defer, Panic and Recover")
	fmt.Printf("Defer:\n")
	deferExample()
	fmt.Printf("Panic and Recover:\n")
	panicAndRecover()
	
	fmt.Println("Pointer")
	pointer()
	fmt.Println("-----")
}

func average(list []float64) float64{
	total := 0.0
	for _, value := range list {
		total += value
	}
	return total / float64(len(list))
}

/*
	Hability to call a function inside another function, and the inner function can access the variables of the outer function.
*/
func clousure(){
	x := 0
	number := func () int{
		x++
		return x
	}
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
}

/*
	Hability to call a function inside itself, allowing for repeated execution of the same code with different parameters.
*/
func factorial(x int) int {
	if(x==0){
		return 1
	}
	return x * factorial(x-1)
}

/*
	Defer: scalone that allows you to delay the execution of a function until the surrounding function returns.
*/
func day1(){
	fmt.Println("Starting day 1")
}
func day2(){
	fmt.Println("Starting day 2")
}

func deferExample() {
	defer day1()
	day2()
}

/*
	Panic: a way to stop the normal execution of a program and start the panic process, which can be recovered with recover.

	Recover: a way to recover from a panic and continue the execution of the program.
*/

func panicAndRecover(){
	defer func(){
		x := recover()
		fmt.Println("Recovered from panic:", x)
	
	}()

	panic("Panic!!")
}

/*
	Pointer: a way to reference a variable's memory address, allowing for direct manipulation of the variable's value.

	* = dereference operator, used to access the value stored at the memory address pointed to by a pointer.
	
	& = address operator, used to get the memory address of a variable.
*/
func initialValue(xPtr *int){
	*xPtr = 0
}

func pointer(){
	x := 5
	initialValue(&x)
	fmt.Println("Initial value:", x)
}