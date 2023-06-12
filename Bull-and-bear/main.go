package main

import "fmt"

func calculateProfitOrLoss(x, y int) string {
	if x > y {
		return "LOSS"
	} else if x < y {
		return "PROFIT"
	} else {
		return "NEUTRAL"
	}
}

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		result := calculateProfitOrLoss(x, y)
		fmt.Println(result)
	}
}
