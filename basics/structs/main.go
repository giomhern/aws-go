package main

import "fmt"

type Person struct {
	Name string
	Age int // if not provided in initialization -> defaults to 0 according to GO
}

func NewPerson(name string, age int) *Person {
	return &Person {
		Name: name, 
		Age: age,
	}
}


func (p *Person) changeName(newName string) { // passing in the actual reference of the Person struct 
	p.Name = newName
}

func main(){

	myPerson := NewPerson("gio", 19)
	myPerson.changeName("mark")

	a := 7
	b := &a
	*b = 9
	fmt.Println(b)
	fmt.Println(a)

	mySlice := []int{
		1, 2, 3,
	}

	for index, _ := range mySlice {
		mySlice[index]++
	}


	fmt.Println(mySlice)
	fmt.Printf("This is my person: %+v \n", myPerson)
}