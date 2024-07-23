package main

import (
	"aws-go/imports" // Adjust the import path based on your project structure
	"fmt"
)

func main(){
	newTicket := imports.Ticket{
		ID: 123, 
		Event: "Course walkthrough", 
	}
	fmt.Println(newTicket)
}