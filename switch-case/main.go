package main

import "fmt"

func main() {
	A := true
	switch A {
	case false:
		{
			fmt.Println("no its not true")
		}
	case true:
		{
			fmt.Println("its true ")
		}
	default:
		{
			fmt.Println("non of the match")
		}
	}
}
