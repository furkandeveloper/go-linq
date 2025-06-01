# linq

[![Go Reference](https://pkg.go.dev/badge/github.com/furkandeveloper/go-linq.svg)](https://pkg.go.dev/github.com/furkandeveloper/go-linq)
[![Go Report Card](https://goreportcard.com/badge/github.com/furkandeveloper/go-linq)](https://goreportcard.com/report/github.com/furkandeveloper/go-linq)
[![License](https://img.shields.io/github/license/furkandeveloper/go-linq)](LICENSE)

`linq` is a lightweight, idiomatic Go library that provides **LINQ-like querying capabilities** on slices and collections. Inspired by LINQ in C#, this library brings expressive, chainable, and functional querying features to Go, simplifying complex collection operations.

---

## Features

- Query slices with fluent, chainable methods
- Filtering with single predicates or predicate groups (AND / OR)
- Projection (`Select`) to transform collections
- Ordering (ascending and descending)
- Set operations: `Distinct`, `Union`, `Intersect`, `Except`
- Aggregation: `Sum`, `Min`, `Max`, `Aggregate`
- Element access: `First`, `ElementAt`, `DefaultIfEmpty`
- Pagination helpers: `Skip`, `Take`
- Reverse collections easily
- Grouping by keys with `GroupBy`
- Conversion utilities: `ToSlice`, `ToMap`
- Works with any type (via generics, Go 1.18+)

---

## Installation

```bash
go get github.com/furkandeveloper/go-linq
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/furkandeveloper/go-linq/linq"
)

func main() {
	data := []int{5, 3, 8, 6, 2}

	// Create a query from slice
	query := linq.From(data)

	// Filter even numbers greater than 3
	evenGreaterThanThree := query.
		Where(func(n int) bool { return n%2 == 0 && n > 3 }).
		ToSlice()

	fmt.Println(evenGreaterThanThree) // Output: [8 6]

	// Sum of all numbers
	total := query.Sum(func(n int) int { return n })
	fmt.Println(total) // Output: 24

	// Order descending and take top 3
	top3 := query.OrderByDescending(func(a, b int) bool { return a < b }).
		Take(3).
		ToSlice()

	fmt.Println(top3) // Output: [8 6 5]
}
```

# API Overview

## Creating Queries
- `From(slice []T) Query[T]`
Initializes a query from a slice.

# Query Functions

## Filtering
- **Where(predicate func(T) bool) Query[T]**  
  Filters items matching the predicate.

- **WhereGroup(group PredicateGroup[T]) Query[T]**  
  Filters with multiple predicates combined by AND or OR.

## Projection
- **Select(source Query[T], selector func(T) U) Query[U]**  
  Projects each element into a new form.

## Ordering
- **OrderBy(less func(T, T) bool) Query[T]**  
  Sorts elements in ascending order.

- **OrderByDescending(less func(T, T) bool) Query[T]**  
  Sorts elements in descending order.

## Set Operations
- **Distinct(equal func(T, T) bool) Query[T]**  
  Removes duplicates.

- **Union(other []T, equal func(T, T) bool) Query[T]**  
  Combines two sequences without duplicates.

- **Intersect(other []T, equal func(T, T) bool) Query[T]**  
  Returns elements present in both sequences.

- **Except(other []T, equal func(T, T) bool) Query[T]**  
  Returns elements in the first sequence but not in the second.

## Aggregation
- **Sum(selector func(T) int) int**  
  Computes the sum.

- **Min(selector func(T) int) int**  
  Finds the minimum value.

- **Max(selector func(T) int) int**  
  Finds the maximum value.

- **Aggregate(seed U, func(U, T) U) U**  
  Accumulates values.

## Element Access
- **First(predicate func(T) bool) (T, bool)**  
  Returns the first element matching the predicate.

- **ElementAt(index int) (T, bool)**  
  Returns the element at the specified index.

- **DefaultIfEmpty(defaultValue T) Query[T]**  
  Returns a default value if the sequence is empty.

## Pagination and Others
- **Skip(count int) Query[T]**  
  Skips the first N elements.

- **Take(count int) Query[T]**  
  Takes the first N elements.

- **Reverse() Query[T]**  
  Reverses the sequence.

## Grouping & Conversion
- **GroupBy(source Query[T], keySelector func(T) K) map[K][]T**  
  Groups elements by key.

- **ToSlice() []T**  
  Converts the query to a slice.

- **ToMap(source Query[T], keySelector func(T) K, valueSelector func(T) V) map[K]V**  
  Converts the query to a map.


