package strings

import (
	"errors"
	"fmt"
	"sort"

	gostrings "strings"
)

type List struct {
	elements []string
}

func NewList() *List {
	l := List{}
	l.elements = []string{}

	return &l
}

func ListFromPgStringArray(rawColumnValue string) (*List, error) {
	if !gostrings.HasPrefix(rawColumnValue, "{") || !gostrings.HasSuffix(rawColumnValue, "}") {
		return nil, errors.New("Invalid column value")
	}
	valueCSV := rawColumnValue[1 : len(rawColumnValue)-1]
	return ListFromCSV(valueCSV, true), nil
}

func ListFromSlice(slice []string) *List {
	l := NewList()
	l.elements = append(l.elements, slice...)

	return l
}

func ListFromCSV(csv string, trimSpace bool) *List {
	l := NewList()
	words := gostrings.Split(csv, ",")

	if trimSpace {
		for _, element := range words {
			l.elements = append(l.elements, gostrings.TrimSpace(element))
		}
	} else {
		l.elements = append(l.elements, words...)
	}

	return l
}

func (l List) Slice() []string {
	slicedList := []string{}
	slicedList = append(slicedList, l.elements...)
	return slicedList
}

func (l *List) Add(elements ...string) {
	l.elements = append(l.elements, elements...)
}

func (l List) Select(filterFunc func(e string) bool) List {
	filteredList := NewList()
	for _, element := range l.elements {
		if filterFunc(element) {
			filteredList.Add(element)
		}
	}

	return *filteredList
}

func (l List) Reject(filterFunc func(e string) bool) List {
	filteredList := NewList()
	for _, element := range l.elements {
		if !filterFunc(element) {
			filteredList.Add(element)
		}
	}

	return *filteredList
}

func (l *List) Delete(i int) error {
	if i < 0 || i > len(l.elements) {
		return errors.New("Invalid index")
	}

	l.elements = append(l.elements[:i], l.elements[i+1:]...)
	return nil
}

func (l *List) RemoveAll(element string) {
	filteredList := l.Select(func(e string) bool { return e != element })
	l.elements = filteredList.elements
}

func (l List) Includes(element string) bool {
	for _, listElement := range l.elements {
		if element == listElement {
			return true
		}
	}

	return false
}

func (l List) Length() int {
	return len(l.elements)
}

func (l List) Each(eachFunc func(element string)) {
	for _, element := range l.elements {
		eachFunc(element)
	}
}

func (l List) Map(mapFunc func(e string) string) List {
	mappedList := NewList()

	for _, element := range l.elements {
		mappedList.Add(mapFunc(element))
	}

	return *mappedList
}

func (l List) Join(separator string) string {
	return gostrings.Join(l.elements, separator)
}

func (l List) CSV() string {
	return l.Join(",")
}

func (l List) Set() Set {
	s := NewSet()
	s.Add(l.elements...)
	return *s
}

func (l List) Sort() List {
	sortedList := NewList()
	sortedList.Add(l.Slice()...)
	sort.Strings(sortedList.elements)
	return *sortedList
}

func (l List) PgStringArrayValue() string {
	return fmt.Sprintf("{%v}", l.CSV())
}
