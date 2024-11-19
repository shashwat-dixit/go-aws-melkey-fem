package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func NewPerson(name string, age int) Person  {
	return Person {
		Name: name,
		Age: age,
	}
}

func (p *Person) changeName(newName string)  {
	fmt.Println("name before", p.Name)
	p.Name = newName
	fmt.Println("Name after", p.Name)	
}

func main()  {
	myPerson := NewPerson("Shashwat", 23)

	a := 7
	b := &a
	fmt.Println(b)

	*b = 9
	fmt.Println( a, *b)

	myPerson.changeName("Frontend Masters")
	fmt.Printf("This is my person %+v\n", myPerson)

	mySlice := []int {
		1,
		2,
		3,
	}

	for index, _ := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)

}