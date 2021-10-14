package main

import (
	"fmt"
	"time"
)

func main() {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		Dob       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

	var dilbert Employee
	dilbert.Salary -= 5000
	dilbert.Position = "engineer"
	position := &dilbert.Position
	*position = "Senior " + *position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " addtional"

	fmt.Println(dilbert)
}
