package main

import "fmt"

type employee struct{
	name string
	sex string
	age int
	salary int
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name: name,
		sex: sex,
		age: age,
		salary: salary,
	}
}

func (e employee) getInfo() string {
 return fmt.Sprintf("Name: %s\nAge: %d\nSalary: %d\n", e.name, e.age, e.salary)
}

func (e *employee) setName(name string) {
	e.name = name
}

func main() {
	employee1 := newEmployee("Dima", "M", 24, 2500)
  employee2 := newEmployee("Vladimir", "M", 42, 1000)

	fmt.Printf("%s\n", employee1.getInfo())
	fmt.Printf("%s\n", employee2.getInfo())

	fmt.Printf("--- Updated type ---\n")

	employee2.setName("Vuasia")

	fmt.Printf("%s\n", employee2.getInfo())
}

