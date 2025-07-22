package main

import ("fmt")

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
