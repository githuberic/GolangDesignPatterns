package MementoPattern

type CareTaker struct {
	mementos []*Memento
}

func (c *CareTaker) Add(memento *Memento) {
	c.mementos = append(c.mementos, memento)
}

func (c *CareTaker) Get(index int) *Memento {
	if len(c.mementos) >= index-1 {
		return c.mementos[index]
	}
	return nil
}
