package linq

import (
	"sort"
)

// Query is the main LINQ structure.
type Query[T any] struct {
	source []T
}

// From creates a new Query from a slice.
func From[T any](source []T) Query[T] {
	return Query[T]{source: source}
}

// Logical operators for WhereGroup
type LogicalOperator int

const (
	And LogicalOperator = iota
	Or
)

// PredicateGroup holds predicates and logical operator between them
type PredicateGroup[T any] struct {
	Predicates      []func(T) bool
	LogicalOperator LogicalOperator
}

// Where applies multiple predicate functions with a logical operator.
func (q Query[T]) WhereGroup(group PredicateGroup[T]) Query[T] {
	var result []T
	for _, item := range q.source {
		include := false
		if group.LogicalOperator == And {
			include = true
			for _, pred := range group.Predicates {
				if !pred(item) {
					include = false
					break
				}
			}
		} else if group.LogicalOperator == Or {
			for _, pred := range group.Predicates {
				if pred(item) {
					include = true
					break
				}
			}
		}
		if include {
			result = append(result, item)
		}
	}
	return Query[T]{source: result}
}

// Where single predicate
func (q Query[T]) Where(predicate func(T) bool) Query[T] {
	return q.WhereGroup(PredicateGroup[T]{Predicates: []func(T) bool{predicate}, LogicalOperator: And})
}

// ToSlice converts the query result to a slice.
func (q Query[T]) ToSlice() []T {
	return q.source
}

// Any returns true if any item satisfies the predicate.
func (q Query[T]) Any(predicate func(T) bool) bool {
	for _, item := range q.source {
		if predicate(item) {
			return true
		}
	}
	return false
}

// All returns true if all items satisfy the predicate.
func (q Query[T]) All(predicate func(T) bool) bool {
	for _, item := range q.source {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// First returns the first item matching the predicate.
func (q Query[T]) First(predicate func(T) bool) (T, bool) {
	for _, item := range q.source {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// Count returns the number of items.
func (q Query[T]) Count() int {
	return len(q.source)
}

// Distinct returns distinct elements based on equality function.
func (q Query[T]) Distinct(equal func(a, b T) bool) Query[T] {
	var result []T
	for _, item := range q.source {
		duplicate := false
		for _, r := range result {
			if equal(item, r) {
				duplicate = true
				break
			}
		}
		if !duplicate {
			result = append(result, item)
		}
	}
	return Query[T]{source: result}
}

// OrderBy sorts items based on less function.
func (q Query[T]) OrderBy(less func(a, b T) bool) Query[T] {
	result := make([]T, len(q.source))
	copy(result, q.source)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return Query[T]{source: result}
}

// OrderByDescending sorts items in descending order.
func (q Query[T]) OrderByDescending(less func(a, b T) bool) Query[T] {
	return q.OrderBy(func(a, b T) bool {
		return !less(a, b)
	})
}

// Sum calculates the sum over a selector.
func (q Query[T]) Sum(selector func(T) int) int {
	sum := 0
	for _, item := range q.source {
		sum += selector(item)
	}
	return sum
}

// Min finds the minimum element based on selector.
func (q Query[T]) Min(selector func(T) int) int {
	if len(q.source) == 0 {
		return 0
	}
	min := selector(q.source[0])
	for _, item := range q.source[1:] {
		v := selector(item)
		if v < min {
			min = v
		}
	}
	return min
}

// Max finds the maximum element based on selector.
func (q Query[T]) Max(selector func(T) int) int {
	if len(q.source) == 0 {
		return 0
	}
	max := selector(q.source[0])
	for _, item := range q.source[1:] {
		v := selector(item)
		if v > max {
			max = v
		}
	}
	return max
}

// Skip skips the first n elements.
func (q Query[T]) Skip(n int) Query[T] {
	if n >= len(q.source) {
		return Query[T]{source: []T{}}
	}
	return Query[T]{source: q.source[n:]}
}

// Take takes the first n elements.
func (q Query[T]) Take(n int) Query[T] {
	if n > len(q.source) {
		n = len(q.source)
	}
	return Query[T]{source: q.source[:n]}
}

// Reverse reverses the slice.
func (q Query[T]) Reverse() Query[T] {
	result := make([]T, len(q.source))
	for i := range q.source {
		result[i] = q.source[len(q.source)-1-i]
	}
	return Query[T]{source: result}
}

// ElementAt returns the element at the specified index.
func (q Query[T]) ElementAt(index int) (T, bool) {
	if index < 0 || index >= len(q.source) {
		var zero T
		return zero, false
	}
	return q.source[index], true
}

// DefaultIfEmpty returns a default value if the slice is empty.
func (q Query[T]) DefaultIfEmpty(defaultValue T) Query[T] {
	if len(q.source) == 0 {
		return Query[T]{source: []T{defaultValue}}
	}
	return q
}

// Aggregate performs an aggregation over the collection.
func (q Query[T]) Aggregate(seed T, aggregator func(T, T) T) T {
	result := seed
	for _, item := range q.source {
		result = aggregator(result, item)
	}
	return result
}

// Union returns the union of two slices based on an equality function.
func (q Query[T]) Union(other []T, equal func(a, b T) bool) Query[T] {
	combined := append(q.source, other...)
	return From(combined).Distinct(equal)
}

// Intersect returns the intersection of two slices based on an equality function.
func (q Query[T]) Intersect(other []T, equal func(a, b T) bool) Query[T] {
	var result []T
	for _, item := range q.source {
		for _, o := range other {
			if equal(item, o) {
				result = append(result, item)
				break
			}
		}
	}
	return From(result).Distinct(equal)
}

// Except returns elements from the first slice that are not in the second.
func (q Query[T]) Except(other []T, equal func(a, b T) bool) Query[T] {
	var result []T
Outer:
	for _, item := range q.source {
		for _, o := range other {
			if equal(item, o) {
				continue Outer
			}
		}
		result = append(result, item)
	}
	return From(result)
}

// --------------------------------------
// Generic functions instead of methods
// --------------------------------------

// Select projects each element into a new form.
func Select[T any, R any](q Query[T], selector func(T) R) Query[R] {
	var result []R
	for _, item := range q.source {
		result = append(result, selector(item))
	}
	return Query[R]{source: result}
}

// GroupBy groups items by key.
func GroupBy[T any, K comparable](q Query[T], keySelector func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range q.source {
		key := keySelector(item)
		result[key] = append(result[key], item)
	}
	return result
}

// ToMap converts the slice to a map using key and value selectors.
func ToMap[T any, K comparable, V any](q Query[T], keySelector func(T) K, valueSelector func(T) V) map[K]V {
	result := make(map[K]V)
	for _, item := range q.source {
		result[keySelector(item)] = valueSelector(item)
	}
	return result
}
