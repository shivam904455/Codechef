package main

import "fmt"

func main() {
	for day:=0 ;day<=100 ;day++{
	Day := 5

	switch Day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("monday")
	case 3:
	fmt.Println("tuesday")
	case 4:
		fmt.Println("wednesday")
	case 5:
		fmt.Println("thursday")
	case 6:
		fmt.Println("friday")
	case 7:
		fmt.Println("saturday")
	default:
	
	}
	// fmt.Println(day)
	if Day==5{
		fmt.Println("true")
	}else {
		fmt.Println("False")
	}
	  
}
}