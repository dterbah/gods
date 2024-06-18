## Go Data Structure (AKA gods)

This library implements multiple data structures than can be used in many contexts. It provides lists, queues, stack and much more !
Get ready to store and manipulate your data 📊.

Here are a summary to access to some examples for each structures.

![CI](https://github.com/dterbah/gods/actions/workflows/go-test.yml/badge.svg)
[![codecov](https://codecov.io/gh/dterbah/gods/branch/main/graph/badge.svg)](https://codecov.io/gh/dterbah/gods)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=dterbah_gods&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=dterbah_gods)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=dterbah_gods&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=dterbah_gods)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=dterbah_gods&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=dterbah_gods)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=dterbah_gods&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=dterbah_gods)

# Table of Contents

1. [Installation](#installation)
2. [List](#list)
   - [ArrayList](#arraylist)
   - [LinkedList](#linkedlist)
3. [Set](#Set)
4. [CircularBuffer](#circularbuffer)
5. [Queue](#queue)
6. [Stack](#stack)
7. [BinaryTree](#binarytree)

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

## Queue

```golang
import (
    "github.com/dterbah/gods/queue"
    "github.com/dterbah/gods/list/arraylist"
    comparator "github.com/dterbah/gods/utils"
)

queue := queue.New[int](comparator.IntComparator)
queue.Enqueue(1) // [1]
queue.Dequeue() // 1, nil
queue.Dequeue() // 0, error !
queue.isEmpty() // True
queue.Enqueue(1, 2, 3)
queue.Contains(1) // true
queue.Copy() // [1, 2, 3]
queue.Size() // 3
queue.Peek() // 1, nil
queue.Clear() // []
list := arraylist.New(comparator.IntComparator, 6, 7, 8)
queue = queue.FromIterable(list, comparator.IntComparator) // [6, 7, 8]
```

## Stack

```golang
import (
    "github.com/dterbah/gods/stack"
    "github.com/dterbah/gods/list/arraylist"
    comparator "github.com/dterbah/gods/utils"
)

stack := stack.New[int](comparator.IntComparator)
stack.Push(2, 3) // [2, 3]
stack.Pop() // 3, nil
stack.Peek() // 2, nil
stack.Contains(2) // True
stack.IsEmpty() // False
stack.Size() // 1
stack.Copy() // [2]
stack.Clear() // []

list := arraylist.New(comparator.IntComparator, 6, 7, 8)
stack = stack.FromIterable(list, comparator.IntComparator) // [6, 7, 8]
```

## BinaryTree

```golang
import (
    "github.com/dterbah/gods/tree"
    comparator "github.com/dterbah/gods/utils"
)

tree := tree.New(comparator.IntComparator)
tree.Add(1, 2, 3) //
tree.has(1) // True
tree.Min() // 1
tree.Max() // 3
it := tree.Iterator()
for it.HasRight() {
    value, err := it.Right()
}

tree.Remove(1) // True
tree.Remove(-1) // False
```
