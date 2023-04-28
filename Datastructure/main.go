package main

import "fmt"

type Student struct {
	Name           string
	Class          int
	Rollnumber     int
	StudentAddress Address
}

type Address struct {
	Vill  string
	Dist  string
	Block string
}
type phone struct {
	model string
	Brand string
	IMEI  int
	Config
}
type Laptop struct {
	model        string
	Brand        string
	Serialnum    int
	Configration Config
}
type Config struct {
	RAM       int
	ROM       int
	Processar string
}

func main() {
	var Vishal Student
	Vishal.Name = "Vishal singh"
	Vishal.Class = 12
	Vishal.Rollnumber = 233
	Vishal.StudentAddress.Vill = "Ghambhir singhpur"
	Vishal.StudentAddress.Dist = "Varansi"
	Vishal.StudentAddress.Block = "Aurai"

	Shivam := Student{
		Name:       "Shivam singh",
		Class:      13,
		Rollnumber: 334,
		StudentAddress: Address{
			Vill:  "aurai",
			Dist:  "Bhadohi",
			Block: "Aurai",
		},
	}
	var VishalAddress *Student
	VishalAddress = &Vishal
	ShivamAddress := &Shivam
	fmt.Println(VishalAddress, ShivamAddress)

	var phone phone
	phone.model = "vivo x21"
	phone.Brand = "vivo"
	phone.IMEI = 133343757
	phone.Config.RAM = 2
	phone.Config.ROM = 128
	phone.Config.Processar = " exctra"

	laptop := Laptop{
		model:     "I5",
		Brand:     "HP",
		Serialnum: 12333,
		Configration: Config{
			RAM:       16,
			ROM:       1024,
			Processar: "intel",
		},
	}
	fmt.Println("student id ", Vishal)
	fmt.Println("student id ", Shivam)
	fmt.Println(" mobile", phone)
	fmt.Println("Laptop", laptop)
}
