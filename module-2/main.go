package main

import "fmt"

func main(){
	
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

