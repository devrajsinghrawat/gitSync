package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Jhon"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Jhon"
	// fmt.Println(alex)

	dev := person{
		firstName: "DevD",
		lastName:  "Singh",
		contact: contactInfo{
			email: "test@email.com",
			zip:   247001,
		},
	}
	// dev.print()

	dev.updateName("Raj")
	dev.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func (p person) updateName(fname string) {
// 	p.firstName = fname
// 	p.print()
// }

func (p *person) updateName(fname string) {
	p.firstName = fname
}
