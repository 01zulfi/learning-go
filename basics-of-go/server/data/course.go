package data

import "fmt"

type Duration float32 // in hours

type Course struct {
	Id       int
	Name     string
	Slug     string
	Duration Duration
}

func (c Course) String() string {
	return fmt.Sprintf("Course: %d %s %s %f", c.Id, c.Name, c.Slug, c.Duration)
}

func (c Course) SignUp() bool {
	return true
}
