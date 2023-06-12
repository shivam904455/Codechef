package main

import "fmt"

func findSecondMaximum(a, b, c int) int {
	if (a >= b && a <= c) || (a >= c && a <= b) {
		return a
	} else if (b >= a && b <= c) || (b >= c && b <= a) {
		return b
	} else {
		return c
	}
}

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)

		secondMax := findSecondMaximum(a, b, c)
		fmt.Println(secondMax)
	}
}
