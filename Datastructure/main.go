package main

import "fmt"

type Student struct {
	Name           string
	Class          int
	Rollno         int
	Section        string
	Subjet         string
	Studentaddress Address
}
type Address struct {
	Vill  string
	Post  string
	Dist  string
	State string
}

func main() {

	var Vishal Student
	Vishal.Name = "vishal singh"
	Vishal.Class = 12
	Vishal.Subjet = "Hindi"
	Vishal.Rollno = 332
	Vishal.Section = "B"
	Vishal.Studentaddress.Vill = "Sawara"
	Vishal.Studentaddress.Post = "Aurai"
	Vishal.Studentaddress.Dist = "Badohi"
	Vishal.Studentaddress.State = "utter pradesh"

	Shivam := Student{
		Name:    "shivam",
		Class:   12212,
		Subjet:  "hindi",
		Section: "c",
		Rollno:  3333,
		Studentaddress: Address{
			Vill:  "ramayan ",
			Post:  "auraiya",
			Dist:  "varansi",
			State: "utter",
		},
	}
	var Vishaladdress * Student
	Vishaladdress =&Vishal
	Shivamaddress:=&Shivam
	fmt.Println(Vishaladdress,Shivamaddress)

	// fmt.Println(Vishal)
	// fmt.Println(Shivam)

}
