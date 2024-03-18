package main

import (
	"fem-go/fem/data"
	"fmt"
)

func main() {
	max := data.Instructor{Id: 3, FirstName: "Maximiliano", LastName: "Firtman"}

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: max}

	fmt.Printf("%v", goCourse)

	swift := data.NewWorkshop("Switch for iOS", max)

	fmt.Printf("%v", swift)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swift

	for _, course := range courses {
		fmt.Println(course)
	}
}
