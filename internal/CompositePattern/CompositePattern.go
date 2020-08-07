package CompositePattern

import "fmt"

type Employee struct {
	name         string
	dept         string
	salary       int
	subordinates []*Employee
}

func GetEmployee(name string, dept string, sal int) *Employee {
	employee := new(Employee)
	employee.name = name
	employee.dept = dept
	employee.salary = sal
	return employee
}

func (e *Employee) Add(ee *Employee) {
	e.subordinates = append(e.subordinates, ee)
}

func (e *Employee) Remove(ee *Employee) {
	target := e.subordinates[:0]
	for _, item := range e.subordinates {
		if item == ee {
			target = append(target, item)
		}
	}
	e.subordinates = target
}

func (e *Employee) GetSubordinates() []*Employee {
	return e.subordinates
}

func (e *Employee) ToString() {
	fmt.Println("name:", e.name, "dept:", e.dept, "salary:", e.salary)
}

func (e *Employee) PrintSubordinates() {
	fmt.Println("=============")
	e.ToString()
	for _, headEmployee := range e.GetSubordinates() {
		headEmployee.ToString()
		for _, employee := range headEmployee.GetSubordinates() {
			employee.ToString()
		}
	}
	fmt.Println("=============")
}
