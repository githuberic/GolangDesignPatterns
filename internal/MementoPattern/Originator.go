package MementoPattern

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SaveStateToMemento() *Memento {
	memento := new(Memento)
	memento.Memento(o.state)
	return memento
}

func (o *Originator) GetStateFromMemento(memento *Memento) {
	o.state = memento.GetMemento()
}
