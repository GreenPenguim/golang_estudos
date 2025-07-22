package main

import "fmt"

/*
	Convert Kalvin to Celsius

	Hint: C = K - 273
*/

const k float64 = 323.44

func main(){
	c := k - 273
	fmt.Printf("Temperature in Kalvin: %g \nTemperature in Celcius: %g", k,c)
}