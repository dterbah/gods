package comparator

import "strings"

type Comparator[T any] func(a, b T) int

func IntComparator(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}

}

func Float32Comparator(a, b float32) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func Float64Comparator(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func StringComparator(a, b string) int {
	return strings.Compare(a, b)
}
