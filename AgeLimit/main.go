package main

import "fmt"

// func main() {
// 	t := 0
// 	fmt.Scanf("%d\n", &t)

// 	for i := 1; i <= t ; i++ {
// 		x, y, a := 0, 0, 0
// 		fmt.Scanf("%d\n", &x)
// 		fmt.Scanf("%d\n", &y)
// 		fmt.Scanf("%d\n", &a)

// 		if a >= x && a < a {
// 			fmt.Println("YES")
// 		} else {
// 			fmt.Println("NO")
// 		}
// 	}
// } 


func main(){
	// your code goes here
	var n int
	fmt.Scanln(&n)
	
	for i := 0; i < n; i++ {
	    var a, b, c int
	    fmt.Scanln(&a, &b, &c)
	    
	    if c >= a && c < b {
	        fmt.Println("YES")
	    } else {
	        fmt.Println("NO")
	    }
	}
}