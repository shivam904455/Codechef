package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		totalPrize := 10*x + 90*y
		fmt.Println(totalPrize)
	}
}
