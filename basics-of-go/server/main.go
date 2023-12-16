package main

import (
	"fmt"
	data "server/data"
)

func main() {
	me := data.Instructor{Id: 1, LastName: "Doe", Score: 100}
	you := data.NewInstructor("John", "Doe")

	goCourse := data.Course{Id: 1, Name: "Go", Slug: "go", Duration: 3.2}

	swiftWorkshop := data.NewWorkshop("Swift", me)

	println(me.Print())
	println(you.Print())
	println(goCourse.String())
	fmt.Printf("%v", swiftWorkshop)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWorkshop

	for _, course := range courses {
		fmt.Println(course)
	}
}
