package main

import "fmt"

func main(){
	/*
		See definintions on: .\module-1\main.go line 14
	*/
	fmt.Println("Loop1")
	loop1()
	fmt.Println("-----")
	fmt.Println("Loop2")
	loop2()
	fmt.Println("-----")
	fmt.Println("Loop3")
	loop3()
	fmt.Println("-----")
	fmt.Println("Arrays")
	array()
	fmt.Println("-----")
	fmt.Println("Slices")
	slice()
	fmt.Println("-----")
	fmt.Println("Maps")
	maps()
}

//for with validations if and switchcase
func loop1(){
	for i := 1; i <= 10; i++ {
		var result string
		if(i % 2 == 0){
			result = "even" 
		}else{
			result = "odd" 
		}
		switch i {
			case 5:
				result += " five"
			case 7:
				result += " seven"
			case 9:
				result += " nine"
		}
		fmt.Printf("%v (%s)\n", i, result)
	}
}


// Loop hierarchy
func loop2(){
	for hours:=0; hours<=12; hours++{
		fmt.Printf("Hour: %d ", hours)
		for minutes:=0;minutes<60; minutes++{
			fmt.Printf("Minute: %d\n", minutes)
		}
	}
}

// In golang, we dont have a 'while', otherwise we can use 'for' loop as 'while' 
func loop3(){
	i:= 0
	for i<=5 {
		fmt.Println("Looping:", i)
		i++
	}
}

func array(){
	var x [5]float64
	x[0] = 5.6
	x[1] = 3.2
	x[2] = 4.8
	x[3] = 2.1
	x[4] = 1.9

	var total float64 = 0

	/*
		Calculate the total
		the function 'len()' returns the length of the array
	*/
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	total = total / float64(len(x))

	fmt.Printf("Total: %.2f\n", total)
}

func slice(){
	// slice := make([]float64,5)
	arr:=[7]float64{1, 2, 3, 4, 5, 6, 7}
	slice:= arr[1:5] // slice from index 1 to 4 (5 is excluded)
	fmt.Println(arr, slice)	

	arr2 := []int{8, 9, 10}
	slice2 := append(arr2, 4, 5, 6) // append elements to the slice
	fmt.Println(arr2, slice2)
	
	arr3 := []int{15, 16, 17}
	slice3 := make([]int, 2)
	copy(slice3, arr3)
	fmt.Println(arr3, slice3)
}

func maps(){
	map1 := make(map[string]int)
	map1["key1"]= 10
	fmt.Println(map1["key1"])
	
	map2 := make(map[int]int)
	map2[1] = 20
	map2[2] = 30
	fmt.Println(map2[1], map2[2])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	fmt.Println(elements["Li"])

}