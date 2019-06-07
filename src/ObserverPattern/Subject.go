package ObserverPattern

type Subject struct {
	observers []Observer
	state     int
}

func (s *Subject) GetState() int {
	return s.state
}

func (s *Subject) SetState(state int) {
	s.state = state
	s.NotifyAllObservers()
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) NotifyAllObservers() {
	for _, observer := range s.observers {
		observer.Update()
	}
}
