package main

import "fmt"

func canAttendConcert(cost int) string {
	totalCost := cost * 4

	if totalCost <= 1000 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var cost int
		fmt.Scan(&cost)

		result := canAttendConcert(cost)
		fmt.Println(result)
	}
}
 