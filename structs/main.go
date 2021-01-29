package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.contactInfo.email = "alex@gmail.com"
	alex.contactInfo.zipCode = 92037
	fmt.Println(alex.contactInfo)

	var jim = person{
		firstName: "Jim",
		lastName:  "Henderson",
		contactInfo: contactInfo{
			email:   "jimbo@gmail.com",
			zipCode: 92038,
		},
	}
	fmt.Println(jim)
}
