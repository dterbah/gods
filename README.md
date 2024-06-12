## Go Data Structure (AKA gods)

This library implements multiple data structures than can be used in many contexts. It provides lists, queues, stack and much more !
Get ready to store and manipulate your data 📊.

Here are a summary to access to some examples for each structures. Be free to access directly to the code in the folder `examples`.

![CI](https://github.com/dterbah/gods/actions/workflows/go-test.yml/badge.svg)
[![codecov](https://codecov.io/gh/dterbah/gods/branch/main/graph/badge.svg)](https://codecov.io/gh/dterbah/gods)

# Table of Contents

1. [Installation](#installation)
2. [List](#list)
   - [ArrayList](#arraylist)
   - [LinkedList](#linkedlist)
3. [Set](#Set)
4. [CircularBuffer](#circularbuffer)

# Installation

Here is the following command to install this library :

```bash
    go get github.com/dterbah/gods
```

# List

## ArrayList

```golang
import (
	comparator "github.com/dterbah/gods/utils"
    "github.com/dterbah/gods/list/arraylist"
)

list := arraylist.New[int](comparator.IntComparator)
list.Add(1, 3, 5, -10)
value, err := list.At(0) // Return 1
size := list.Size() // Return
isRemoved := list.RemoveAt(0) // Return true
containsHeight := list.Contains(1090) // Return false
indexOfThree := list.IndexOf(3) // Return 1
isEmpty := list.Ismpty() // Return false
isReplaced := list.ReplaceAt(100, 2) // Return true
reveresedList := list.Reverse() // Return [-10, 5, 3, 1]
list.some(func(element, index int) {
    return element == 2
}) // Return false
list.Sort() // Content of the list : [-10, 1, 3, 5]
list.SubList(1, 3) // Return [3, 5, -10]
list.Filter(func(element int) {
    return element > 4
}) // [5]
list.Some(func(element, index int) {
    return element > 4
}) // true
list.Sort() // [-10, 1, 3, 5]
}
```

## LinkedList

```golang
import (
	comparator "github.com/dterbah/gods/utils"
    "github.com/dterbah/gods/list/linkedlist"
)

list := linkedlist.New[int](comparator.IntComparator)
list.Add(1, 3, 5, -10)
value, err := list.At(0) // Return 1
size := list.Size() // Return
isRemoved := list.RemoveAt(0) // Return true
containsHeight := list.Contains(1090) // Return false
indexOfThree := list.IndexOf(3) // Return 1
isEmpty := list.Ismpty() // Return false
isReplaced := list.ReplaceAt(100, 2) // Return true
reveresedList := list.Reverse() // Return [-10, 5, 3, 1]
list.some(func(element, index int) {
    return element == 2
}) // Return false
list.Sort() // Content of the list : [-10, 1, 3, 5]
list.SubList(1, 3) // Return [3, 5, -10]
list.Filter(func(element int) {
    return element > 4
}) // [5]
list.Some(func(element, index int) {
    return element > 4
}) // true
list.Sort() // [-10, 1, 3, 5]
}
```

# Set

```golang
import (
	comparator "github.com/dterbah/gods/utils"
    "github.com/dterbah/gods/set/set"
)

set := set.New[int](comparator.IntComparator, 1, 2, 3)
set2 := set.New(comparator.IntComparator, 1, 2, 5, 6)
set.Add(4) // {1, 2, 3, 4}
set.Add(1) // {1, 2, 3, 4}
set.Size() // 4
set.At(0) // 1, nil
set.At(-1) // 0, err
set.Clear() // {}
set.Contains(1) // true
set.ContainsAll(set2) // false
set.Copy() // {1, 2, 3, 4}
set.IndexOf(4) // 3
set.Intersection(set2) // {1, 2}
set.Union(set2) // {3, 5, 6}
set.Diff(set2) // {3}
set.IsSubset(set2) // false
set.Remove(1) // {2, 3, 4}
}
```

## CircularBuffer

```golang
import (
    "github.com/dterbah/gods/buffer/circular"
)

buffer := circular.New[int](2)
buffer.Enqueue(1) // nil
buffer.Enqueue(2) // nil
buffer.Enqueue(3) // error

buffer.IsFull() // true
buffer.IsEmpty() // false

buffer.Dequeue() // 1
buffer.Dequeue() // 2
buffer.Dequeue() // 3
buffer.Dequeue() // (0, err)

```
