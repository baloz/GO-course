package main

import "fmt"

type person struct{
	firstName string
	lastName string
	contact contactInfo
}
type contactInfo struct{
	email string
	pincode int
}
func main(){
	jim := person{
		firstName: "jim",
		lastName: "anderson",
		contact: contactInfo{
			email: "jim@gmail.com",
			pincode: 50634,
		},
	}
	fmt.Printf("%+v\n", jim)
}