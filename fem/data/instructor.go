package data

import "fmt"

type Instructor struct {
	FirstName string
	LastName  string
	Id        int
	Score     int
}

func NewInstructor(firstName string, lastName string) Instructor {
	return Instructor{
		FirstName: firstName,
		LastName:  lastName,
		Id:        0,
		Score:     0,
	}
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)", i.LastName, i.FirstName, i.Score)
}
