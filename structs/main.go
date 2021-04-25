package main

import "fmt"

type person struct{
	firstName string
	lastName string
}

func main(){
	pavan := person{"pavan", "baloju"}
	fmt.Println(pavan)

	pavan = person{firstName: "pavan", lastName: "baloju"}
	fmt.Println(pavan)
}