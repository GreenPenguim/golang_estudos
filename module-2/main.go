package main

import "fmt"

func main(){
	loop1()
	fmt.Println("-----")
	loop2()
	fmt.Println("-----")
	loop3()
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
	i:=0
	for i<=5 {
		fmt.Println("Looping:", i)
		i++
	}
}