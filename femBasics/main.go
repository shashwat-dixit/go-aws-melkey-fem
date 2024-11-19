package main

import (
	"femBasics/imports"
	"fmt"
)


func main() {
	fmt.Println("Hello, World")

	newTicket := imports.Ticket{
		ID: 123,
		Event: "Frontend Masters",
	}
	
	newTicket.PrintEvent()
	
}