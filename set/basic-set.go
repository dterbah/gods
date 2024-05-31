package set

import (
	"github.com/dterbah/gods/collection"
	"github.com/dterbah/gods/iterable"
)

/*
Interface used to define the available methods
for a Set of data
*/
type BasicSet[T any] interface {
	collection.Collection[T]
	iterable.Iterable[T]

	/*
		Copy the current set to another. It is a shallow copy.
	*/
	Copy() BasicSet[T]

	Intersection(other BasicSet[T]) BasicSet[T]

	Union(other BasicSet[T]) BasicSet[T]
}
