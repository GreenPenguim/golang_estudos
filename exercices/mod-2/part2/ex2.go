package main

import "fmt"

/*
	Create a program that uses the operator '%' and 'for'

	**Part 2** Pin & Pan
	Pin -> numbers that are divisible by 3
	Pan -> numbers that are divisible by 5

	DON'T show the numbers that are divisible by both 3 and 5.
	Show the numbers from 1 to 100 and if they are Pin or/and Pan.
*/
func main(){
	for i:=1; i<=100;i++{
		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Printf("Pin Pan!\n")
			break
		case i % 3 == 0:
			fmt.Printf("Pin!\n")
		case i % 5 == 0:
			fmt.Printf("Pan!\n")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}