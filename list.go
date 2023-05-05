package lists

import "fmt"

type List []any

func NewList() *List {
	newList := make(List, 0)
	return &newList
}

func (l *List) Get(index int) (any, error) {
	if index < 0 || index >= len(*l) {
		return nil, fmt.Errorf("index must be greater or equal than 0 and less than the length of the list")
	}
	return (*l)[index], nil
}

func (l *List) Insert(index int, item any) error {
	if index < 0 || index > len(*l) {
		return fmt.Errorf("index must be greater or equal than 0 and less or equal than the length of the list")
	}

	newL := make([]any, 0)
	if index == len(*l) {
		newL = append((*l)[:], item)
	} else {
		newL = append((*l)[:index], item)
		newL = append(newL, (*l)[index:]...)
	}

	*l = make([]any, len(newL))
	*l = newL
	return nil
}

func (l *List) Remove(index int) (any, error) {
	if index < 0 || index >= len(*l) {
		return nil, fmt.Errorf("index must be greater or equal than 0 and less than the length of the list")
	}

	cropped := (*l)[index]

	newL := make([]any, 0)
	if index == len(*l)-1 {
		newL = (*l)[:index]
	} else {
		newL = append((*l)[:index], (*l)[index+1:]...)
	}

	*l = make([]any, len(newL))
	*l = newL
	return cropped, nil
}

func (l *List) Clear() {
	*l = make(List, 0)
}

func (l *List) Contains(item any) bool {
	_, err := l.FirstIndex(item)
	if err != nil {
		return false
	}
	return true
}

func (l *List) Indexes(item any) ([]int, error) {
	indexes := make([]int, 0)
	for i := 0; i < len(*l); i++ {
		if (*l)[i] == item {
			indexes = append(indexes, i)
		}
	}

	if len(indexes) == 0 {
		return nil, fmt.Errorf("item does not exist in the list")
	}
	return indexes, nil
}

func (l *List) FirstIndex(item any) (int, error) {
	firstIndex := -1
	for i := 0; i < len(*l); i++ {
		if (*l)[i] == item {
			firstIndex = i
			break
		}
	}

	if firstIndex == -1 {
		return -1, fmt.Errorf("item does not exist in the list")
	}
	return firstIndex, nil
}

func (l *List) LastIndex(item any) (int, error) {
	lastIndex := -1
	for i := len(*l) - 1; i >= 0; i-- {
		if (*l)[i] == item {
			lastIndex = i
			break
		}
	}

	if lastIndex == -1 {
		return -1, fmt.Errorf("item does not exist in the list")
	}
	return lastIndex, nil
}

func (l *List) Length() int {
	return len(*l)
}
