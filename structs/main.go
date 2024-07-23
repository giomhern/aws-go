package main

import "fmt"

type Person struct {
	Name string
	Age int // if not provided in initialization -> defaults to 0 according to GO
}


func main(){
	myPerson := Person {
		Name: "Gio", 
		Age: 26,
	}

	myPerson.Name = "Mark"

	fmt.Println(myPerson)
	fmt.Printf("This is my persion: %+v", myPerson)
}