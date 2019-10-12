package main

import "fmt"

type Company struct {
	name      string
	Employees []*Employee
}

type Employee struct {
	id   uint
	name string
}

func main() {
	emp1 := &Employee{1, "jhon doe"}
	emp2 := &Employee{2, "patrick seen"}

	employees := []*Employee{emp1, emp2}

	company := Company{"big company", employees}

	printEmployees(&company)
	modifyEmployeesInCompany(&company)
	printEmployees(&company)
}

func modifyEmployeesInCompany(company *Company) {
	company.Employees[0].name = "new jhon doe"
	company.Employees[1].name = "new patrick seen"
}

func printEmployees(company *Company) {
	for _, e := range company.Employees {
		fmt.Println(e)
	}
}
