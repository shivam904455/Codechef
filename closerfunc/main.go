package main

import "fmt"

func newid() func() int {
	i := 5
	return func() int {
		i++
		return i
	}
}
func main() {
	Result := newid()
	fmt.Println("Result =", Result())
	fmt.Println("Result =", Result())

}
