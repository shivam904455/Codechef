package main

import "fmt"

func newid() func () int {
	i := 5
	return func  ()int{
 i++
 return i
	}
	
}
func main() {
	result :=newid()
	fmt.Println("result =", result())
	fmt.Println("result =", result())
} 
