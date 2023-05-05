package lists

type Queue struct {
	l *List
}

func NewQueue() *Queue {
	newQueue := Queue{NewList()}
	return &newQueue
}

func (q *Queue) Get() (any, error) {
	return q.l.Get(0)
}

func (q *Queue) Insert(item any) error {
	return q.l.Insert(q.l.Length(), item)
}

func (q *Queue) Remove() (any, error) {
	return q.l.Remove(0)
}

func (q *Queue) Clear() {
	q.l.Clear()
}

func (q *Queue) Contains(item any) bool {
	return q.l.Contains(item)
}

func (q *Queue) FirstIndex(item any) (int, error) {
	return q.l.FirstIndex(item)
}

func (q *Queue) Length() int {
	return q.l.Length()
}
