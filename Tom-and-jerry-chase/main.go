package main

import "fmt"

func main() {
    var T int
    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        var X, Y int
        fmt.Scan(&X, &Y)

        if X <= Y {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
 