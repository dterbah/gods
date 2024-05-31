## Go Data Structure (AKA gods)

This library implements multiple data structures than can be used in many contexts. It provides lists, queues, stack and much more !
Get ready to store and manipulate your data 📊.

Here are a summary to access to some examples for each structures. Be free to access directly to the code in the folder `examples`.

# Table of Contents

1. [Installation](#installation)
2. [List](#list)
   - [ArrayList](#arraylist)
3. [Set](#Set)
   - [Basic Set](#basicset)

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
isEmpty := list.EIsmpty() // Return false
isReplaced := list.ReplaceAt(100, 2) // Return true
reveresedList := list.Reverse() // Return [-10, 5, 3, 1]
list.some(func(element, index int) {
    return element == 2
}) // Return false
list.Sort() // Content of the list : [-10, 1, 3, 5]
list.SubList(1, 3) // Return [3, 5]
}
```

# Set

## BasicSet

```golang
import (
	comparator "github.com/dterbah/gods/utils"
    "github.com/dterbah/gods/set/set"
)

set := set.New[int](comparator.IntComparator, 1, 2, 3)
set.Add(4) // [1, 2, 3, 4]
set.Add(1) // [1, 2, 3, 4]
set.Size() // 4
set.At(0) // 1, nil
set.At(-1) // 0, err
}
```
