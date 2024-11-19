package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) circumference() float64 {
	return c.Radius * 2 * math.Pi
}
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}


func main()  {
	newCircle := Circle{
		Radius: 5,
	}
	
	fmt.Println("The circumference of the cirle is", math.Floor(newCircle.circumference()))
	fmt.Println("The area of the cirle is", math.Floor(newCircle.area()))
}


