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
	"github.com/yourusername/linq"
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

