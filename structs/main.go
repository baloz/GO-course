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
	jim.updateName("jimmy") // allowed instead of (&jim).updateName 
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newname string){
	(*p).firstName = newname;
}
