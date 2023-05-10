package main

import "fmt"

func main() {
	shivam(1, 23, 3, 4, 3, 2, 2, 2, 2, 2, 2)

	shivam([]int{1, 2, 2, 1, 1, 3, 3}...)
	
}
 
func shivam(numbers ...int) {
	
	fmt.Println(numbers)
	
	fmt.Println("length", len(numbers))
	
	total := 0
	
	for _, val := range numbers {
	
	total += val
	
}

	fmt.Println(total)

}
