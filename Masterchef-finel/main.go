package main

import "fmt"

func main() {
    var T int
    fmt.Scanln(&T) // Read the number of test cases

    for i := 0; i < T; i++ {
        var X int
        fmt.Scanln(&X) // Read X for each test case

        if X <= 10 {
            fmt.Println("YES") // Print YES if X is less than or equal to 10
        } else {
            fmt.Println("NO") // Print NO if X is greater than 10
        }
    }
}
