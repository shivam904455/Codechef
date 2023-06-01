package main

import "fmt"

func main() {
	var T int
	fmt.Scanln(&T) // Read the number of test cases

	for i := 0; i < T; i++ {
		var X int
		fmt.Scanln(&X) // Read X for each test case

		if X == 6 {
			fmt.Println("YES") // Chef rolled a 6, so he can enter a new token
		} else {
			fmt.Println("NO") // Chef did not roll a 6, so he cannot enter a new token
		}
	}
}
