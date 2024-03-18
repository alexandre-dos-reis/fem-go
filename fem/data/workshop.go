package data

import "time"

// some kind of "inheritance" :
type Workshop struct {
	Date time.Time
	Instructor
	Course
}

// Implicit interface implementation
func (w Workshop) SignUp() bool {
	return true
}

func NewWorkshop(name string, i Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = i
	return w
}
