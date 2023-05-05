package lists

type Listable interface {
	Get(index int) (any, error)
	Insert(index int, item any) error
	Remove(index int) (any, error)
}
