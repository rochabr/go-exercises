package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// func main() {
// 	jim := person{
// 		firstName: "Jim",
// 		lastName:  "James",
// 		contactInfo: contactInfo{
// 			email:   "jim@gmail.com",
// 			zipCode: 90989,
// 		},
// 	}

// 	jim.updateName("James")
// 	jim.print()
// }

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func updateValue(n string) {
	n = "Alex"
}
