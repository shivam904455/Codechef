package main

import "fmt"

func main() {
    var T int
    fmt.Scanln(&T) // Read the number of test cases

    for i := 0; i < T; i++ {
        var X, Y int
        fmt.Scanln(&X, &Y) // Read X and Y for each test case

        diff := X - Y // Calculate the minimum investment required
        fmt.Println(diff) // Print the result
    }
}
