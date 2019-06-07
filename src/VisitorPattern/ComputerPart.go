package VisitorPattern

type ComputerPart interface {
	Accept(computerPartVisitor ComputerPartVisitor)
}

type Keyboard struct {
}

func (s *Keyboard) Accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.VisitKeyboard(s)
}

type Monitor struct {
}

func (s *Monitor) Accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.VisitMonitor(s)
}

type Mouse struct {
}

func (s *Mouse) Accept(computerPartVisitor ComputerPartVisitor) {
	computerPartVisitor.VisitMouse(s)
}


