package main

import "fmt"

func main() {
	getsum(1 ,2,3)
	getsum(400 ,400,500,200,)
getsum([]int{20,30,30,40,100}...)
}

func getsum(number ...int) {
fmt.Println(number ," = ")
fmt.Println(len(number))

	total := 0
	for _, val := range number {
		total += val
	}
	fmt.Println(total)
}
