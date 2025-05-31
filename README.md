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
