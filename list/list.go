//go:generate mockgen -package list -source=list.go -destination list_mock.go

package list

// List defines the behaviour of a data structure able to provide insert, remove and length operations.
type List interface {
	Insert(int)
	Remove(int) error
	Length() int
}
