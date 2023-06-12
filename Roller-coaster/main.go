package main

import "fmt"

func canGoOnRide(height, minHeight int) string {
	if height >= minHeight {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var T int
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		var height, minHeight int
		fmt.Scanln(&height, &minHeight)
		fmt.Println(canGoOnRide(height, minHeight))
	}
}
