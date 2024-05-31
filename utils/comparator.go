package comparator

import "strings"

type Comparator[T any] func(a, b T) bool

func IntComparator(a, b int) bool {
	return a == b
}

func Float32Comparator(a, b float32) bool {
	return a == b
}

func Float64Comparator(a, b float64) bool {
	return a == b
}

func StringComparator(a, b string) bool {
	return strings.Compare(a, b) == 0
}
