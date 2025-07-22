package main

import "fmt"

/*
	Create a program that uses the operator '%' and 'for'

	**Part 1**
	Show numbers from 1 to 100 and are divisible by 3.
*/
func main(){
	for i:=1; i<=100;i++{
		if(i % 3 == 0){
			fmt.Printf("%d is divisible by 3\n", i)
		}
	}
}