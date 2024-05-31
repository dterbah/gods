package set

import "github.com/dterbah/gods/list"

/*
Interface used to define the available methods
for a Set of data
*/
type BasicSet[T any] interface {
	list.List[T]
}
