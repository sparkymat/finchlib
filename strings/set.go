package strings

type Set struct {
	elements map[string]struct{}
}

func NewSet() *Set {
	s := Set{}
	s.elements = make(map[string]struct{})

	return &s
}

func (s *Set) Add(elements ...string) {
	for _, element := range elements {
		s.elements[element] = struct{}{}
	}
}

func (s *Set) Remove(element string) {
	delete(s.elements, element)
}

func (s Set) Includes(element string) bool {
	_, exists := s.elements[element]
	return exists
}

func (s Set) List() List {
	l := NewList()
	for element, _ := range s.elements {
		l.Add(element)
	}
	return *l
}
