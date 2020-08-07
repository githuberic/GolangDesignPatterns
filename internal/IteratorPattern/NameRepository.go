package IteratorPattern

type NameRepository struct {
}

func (s *NameRepository) GetIterator() Iterator {
	nameIterator := new(NameIterator)
	nameIterator.names = []string{"Robert", "John", "Julie", "Lora"}
	return nameIterator
}

type NameIterator struct {
	names []string
	index int
}

func (s *NameIterator) HasNext() bool {
	if s.index < len(s.names) {
		return true
	}
	return false
}

func (s *NameIterator) Next() SomeSlice {
	if s.HasNext() {
		n := s.index
		s.index++
		return s.names[n]
	}
	return nil
}
