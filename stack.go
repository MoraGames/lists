package lists

type Stack struct {
	l *List
}

func NewStack() *Stack {
	newStack := Stack{NewList()}
	return &newStack
}

func (s *Stack) Get() (any, error) {
	return s.l.Get(s.l.Length() - 1)
}

func (s *Stack) Insert(item any) error {
	return s.l.Insert(s.l.Length(), item)
}

func (s *Stack) Remove() (any, error) {
	return s.l.Remove(s.l.Length() - 1)
}

func (s *Stack) Clear() {
	s.l.Clear()
}

func (s *Stack) Contains(item any) bool {
	return s.l.Contains(item)
}

func (s *Stack) LastIndex(item any) (int, error) {
	return s.l.LastIndex(item)
}

func (s *Stack) Length() int {
	return s.l.Length()
}
