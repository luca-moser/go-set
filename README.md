# Go Set [![Build Status](https://travis-ci.org/luca-moser/set.svg?branch=master)](https://travis-ci.org/luca-moser/set)

A simple and thread-safe set in Go.

```
// create a new set
set := NewSet(123, "123", struct{name}{"alice"})

// length
set.Length() // 3

// check existance
set.Has(123) // true
set.Has(321) // false

// add entries
amountAdded := set.Add(503, 403, "apple", "orange") // 4 
set.Length() // 7

// remove entries
amountRemoved := set.Remove(122, "123") // 2

// clear existing data
set.Clear()
set.Length() // 0

// add a new set
set2 := NewSet("grapefruit", "orange")
set.Join(set2)
set.Length() // 2

set.Add(1,2,3,4,5)
set.Ints() // []int{1,2,3,4,5}
```