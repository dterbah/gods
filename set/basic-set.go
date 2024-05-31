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

	/*
		Create a new Set with all elements in the current set that
		are not present on the set passed in param
	*/
	Diff(other BasicSet[T]) BasicSet[T]

	/*
		Return true if the set passed in param is a subset of the current set, else false
	*/
	IsSubset(other BasicSet[T]) bool

	/*
		Compute the intersection between the current set and the one passed in parameter.
		The result is equivalent of A ∩ B
	*/
	Intersection(other BasicSet[T]) BasicSet[T]

	/*
	   Compute the union between the current set and the one passed in parameter.
	   The result is equivalent of A ∪ B
	*/
	Union(other BasicSet[T]) BasicSet[T]
}
