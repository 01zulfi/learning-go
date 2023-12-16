package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func (i Instructor) Print() string {
	return fmt.Sprintf("Instructor: %d %s %s %d", i.Id, i.FirstName, i.LastName, i.Score)
}

func NewInstructor(firstName string, lastName string) Instructor {
	return Instructor{FirstName: firstName, LastName: lastName}
}
