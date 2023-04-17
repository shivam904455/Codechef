package main

import "fmt"

// bh:= Batery helth
func OptimalCondition(bh int) string {
	if bh >= 80 {
		return "Yes"
	} else {
		return "no"
	}
}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var bh int
		fmt.Scan(&bh)
		fmt.Println(OptimalCondition(bh))
	}
}
