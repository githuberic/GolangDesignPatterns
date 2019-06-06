package MementoPattern

type Memento struct {
	state string
}

func (m *Memento) Memento(state string) {
	m.state = state
}

func (m *Memento) GetMemento() string {
	return m.state
}
