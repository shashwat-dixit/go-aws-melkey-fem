package main

import (
	"fmt"
	"slices"
)

func main()  {
	// array
	animals := [2]string{
		"dog",
		"cat",
	}
	fmt.Println(animals)

	// slice

	animalsAgain := [] string {
		"cat",
		"dog",
	}

	animalsAgain = slices.Delete(animalsAgain, 0, 1)
	animalsAgain = append(animalsAgain, "rat")

	for i := 0; i < len(animalsAgain); i++ {
		fmt.Printf("This is my animal %s\n", animalsAgain[i])
	}

	for index, value := range animalsAgain {
		fmt.Println(index, value)
	}

	fmt.Println(animalsAgain)
}