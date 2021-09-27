package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lstName   string
	contactInfo
}

func main() {
	jim := person{
		firstName: "shikha",
		lstName:   "NA",
		contactInfo: contactInfo{
			email:   "tshikha634@gmail.com",
			zipCode: 122006,
		},
	}
	jim.updateName("sh")
	jim.print()

}

func (pointerToName *person) updateName(newfirstName string) {
	(*pointerToName).firstName = newfirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
