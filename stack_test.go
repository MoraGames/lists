package lists

import (
	"testing"
)

func TestGet(t *testing.T) {
	s := NewStack()

	item1, err := s.Get()
	if err == nil {
		t.Fail()
	}
	t.Logf("Stack.Get() = %+v, %v\n", item1, err)

	s.Insert("Item#01")
	item2, err := s.Get()
	if err != nil {
		t.Fail()
	}
	t.Logf("Stack.Get() = %+v, %v\n", item2, err)
}