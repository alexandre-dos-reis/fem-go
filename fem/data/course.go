package data

import "fmt"

type Duration float32

type Course struct {
	Name       string
	Slug       string
	Instructor Instructor
	Id         int
	Duration   Duration
	Legacy     bool
}

// Implicit interface implementation
func (c Course) SignUp() bool {
	return true
}

func (c Course) String() string {
	return fmt.Sprintf("--- %v --- (%v)\n", c.Name, c.Instructor.FirstName)
}
