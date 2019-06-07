package VisitorPattern

type Computer struct {
	parts []ComputerPart
}

func (c *Computer) Computer() *Computer {
	c.parts = []ComputerPart{new(Mouse), new(Keyboard), new(Monitor)}
	return c
}

func (c *Computer) Accept(computerPartVisitor ComputerPartVisitor) {
	for _, part := range c.parts {
		part.Accept(computerPartVisitor)
	}
	computerPartVisitor.VisitComputer(c)
}
