package VisitorPattern

import "fmt"

type ComputerPartVisitor interface {
	VisitComputer(computer *Computer)
	VisitMouse(mouse *Mouse)
	VisitKeyboard(keyboard *Keyboard)
	VisitMonitor(monitor *Monitor)
}

type ComputerPartDisplayVisitor struct {
}

func (c *ComputerPartDisplayVisitor) VisitComputer(computer *Computer) {
	fmt.Println("Displaying VisitComputer.")
}

func (c *ComputerPartDisplayVisitor) VisitMouse(mouse *Mouse) {
	fmt.Println("Displaying VisitMouse.")
}

func (c *ComputerPartDisplayVisitor) VisitKeyboard(keyboard *Keyboard) {
	fmt.Println("Displaying VisitKeyboard.")

}

func (c *ComputerPartDisplayVisitor) VisitMonitor(monitor *Monitor) {
	fmt.Println("Displaying VisitMonitor.")
}
