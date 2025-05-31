package linq

import (
	"testing"
)

type testStruct struct {
	ID    int
	Name  string
	Value int
}

func equalTestStruct(a, b testStruct) bool {
	return a.ID == b.ID && a.Name == b.Name && a.Value == b.Value
}

func TestFromAndToSlice(t *testing.T) {
	data := []int{1, 2, 3}
	q := From(data)
	result := q.ToSlice()
	if len(result) != 3 || result[0] != 1 || result[2] != 3 {
		t.Errorf("From or ToSlice failed, got %v", result)
	}
}

func TestWhereAndWhereGroup(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	q := From(data)
	// single predicate
	even := q.Where(func(n int) bool { return n%2 == 0 }).ToSlice()
	if len(even) != 2 || even[0] != 2 || even[1] != 4 {
		t.Errorf("Where failed, got %v", even)
	}

	// predicate group with AND
	group := PredicateGroup[int]{
		Predicates: []func(int) bool{
			func(n int) bool { return n > 2 },
			func(n int) bool { return n < 5 },
		},
		LogicalOperator: And,
	}
	result := q.WhereGroup(group).ToSlice()
	expected := []int{3, 4}
	if len(result) != len(expected) {
		t.Fatalf("WhereGroup AND length mismatch, got %v", result)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("WhereGroup AND expected %d, got %d", expected[i], v)
		}
	}

	// predicate group with OR
	groupOr := PredicateGroup[int]{
		Predicates: []func(int) bool{
			func(n int) bool { return n == 1 },
			func(n int) bool { return n == 5 },
		},
		LogicalOperator: Or,
	}
	resultOr := q.WhereGroup(groupOr).ToSlice()
	expectedOr := []int{1, 5}
	if len(resultOr) != len(expectedOr) {
		t.Fatalf("WhereGroup OR length mismatch, got %v", resultOr)
	}
	for i, v := range resultOr {
		if v != expectedOr[i] {
			t.Errorf("WhereGroup OR expected %d, got %d", expectedOr[i], v)
		}
	}
}

func TestAnyAll(t *testing.T) {
	data := []int{2, 4, 6}
	q := From(data)
	if !q.Any(func(n int) bool { return n == 4 }) {
		t.Error("Any failed: expected true")
	}
	if q.Any(func(n int) bool { return n == 5 }) {
		t.Error("Any failed: expected false")
	}

	if !q.All(func(n int) bool { return n%2 == 0 }) {
		t.Error("All failed: expected true")
	}
	if q.All(func(n int) bool { return n > 4 }) {
		t.Error("All failed: expected false")
	}
}

func TestFirst(t *testing.T) {
	data := []int{10, 20, 30, 40}
	q := From(data)
	item, found := q.First(func(n int) bool { return n > 25 })
	if !found || item != 30 {
		t.Errorf("First failed, got %d, found: %v", item, found)
	}
	item, found = q.First(func(n int) bool { return n > 100 })
	if found {
		t.Errorf("First should not find any item, got %d", item)
	}
}

func TestCount(t *testing.T) {
	data := []int{1, 2, 3, 4}
	q := From(data)
	if q.Count() != 4 {
		t.Errorf("Count failed, expected 4 got %d", q.Count())
	}
}

func TestDistinct(t *testing.T) {
	data := []testStruct{
		{1, "a", 10},
		{2, "b", 20},
		{1, "a", 10},
		{3, "c", 30},
	}
	q := From(data)
	dist := q.Distinct(equalTestStruct).ToSlice()
	if len(dist) != 3 {
		t.Errorf("Distinct failed, expected 3 got %d", len(dist))
	}
}

func TestOrderByAndOrderByDescending(t *testing.T) {
	data := []int{5, 2, 4, 1, 3}
	q := From(data)
	asc := q.OrderBy(func(a, b int) bool { return a < b }).ToSlice()
	for i := 1; i < len(asc); i++ {
		if asc[i] < asc[i-1] {
			t.Errorf("OrderBy ascending failed, %v", asc)
		}
	}
	desc := q.OrderByDescending(func(a, b int) bool { return a < b }).ToSlice()
	for i := 1; i < len(desc); i++ {
		if desc[i] > desc[i-1] {
			t.Errorf("OrderByDescending failed, %v", desc)
		}
	}
}

func TestSumMinMax(t *testing.T) {
	data := []int{1, 3, 5, 7, 9}
	q := From(data)
	sum := q.Sum(func(n int) int { return n })
	if sum != 25 {
		t.Errorf("Sum failed, expected 25 got %d", sum)
	}
	min := q.Min(func(n int) int { return n })
	if min != 1 {
		t.Errorf("Min failed, expected 1 got %d", min)
	}
	max := q.Max(func(n int) int { return n })
	if max != 9 {
		t.Errorf("Max failed, expected 9 got %d", max)
	}
}

func TestSkipTakeReverse(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	q := From(data)
	skip := q.Skip(2).ToSlice()
	if len(skip) != 3 || skip[0] != 3 {
		t.Errorf("Skip failed, got %v", skip)
	}
	take := q.Take(3).ToSlice()
	if len(take) != 3 || take[2] != 3 {
		t.Errorf("Take failed, got %v", take)
	}
	reversed := q.Reverse().ToSlice()
	for i := range reversed {
		if reversed[i] != data[len(data)-1-i] {
			t.Errorf("Reverse failed, got %v", reversed)
			break
		}
	}
}

func TestElementAtDefaultIfEmpty(t *testing.T) {
	data := []int{10, 20, 30}
	q := From(data)
	elem, ok := q.ElementAt(1)
	if !ok || elem != 20 {
		t.Errorf("ElementAt failed, got %d, ok=%v", elem, ok)
	}
	elem, ok = q.ElementAt(10)
	if ok {
		t.Errorf("ElementAt out of range should fail")
	}

	empty := From([]int{})
	defaulted := empty.DefaultIfEmpty(100).ToSlice()
	if len(defaulted) != 1 || defaulted[0] != 100 {
		t.Errorf("DefaultIfEmpty failed, got %v", defaulted)
	}
}

func TestAggregate(t *testing.T) {
	data := []int{1, 2, 3, 4}
	q := From(data)
	sum := q.Aggregate(0, func(acc, item int) int { return acc + item })
	if sum != 10 {
		t.Errorf("Aggregate sum failed, got %d", sum)
	}
}

func TestUnionIntersectExcept(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	equalInt := func(a, b int) bool { return a == b }

	q := From(a)
	union := q.Union(b, equalInt).ToSlice()
	expectedUnion := []int{1, 2, 3, 4, 5, 6}
	if len(union) != len(expectedUnion) {
		t.Errorf("Union length mismatch, got %v", union)
	}

	intersect := q.Intersect(b, equalInt).ToSlice()
	expectedIntersect := []int{3, 4}
	if len(intersect) != len(expectedIntersect) {
		t.Errorf("Intersect length mismatch, got %v", intersect)
	}

	except := q.Except(b, equalInt).ToSlice()
	expectedExcept := []int{1, 2}
	if len(except) != len(expectedExcept) {
		t.Errorf("Except length mismatch, got %v", except)
	}
}

func TestSelectGroupByToMap(t *testing.T) {
	data := []testStruct{
		{1, "a", 10},
		{2, "b", 20},
		{3, "a", 30},
	}
	q := From(data)

	selected := Select(q, func(t testStruct) string { return t.Name }).ToSlice()
	expectedSelect := []string{"a", "b", "a"}
	for i, v := range selected {
		if v != expectedSelect[i] {
			t.Errorf("Select failed at index %d, expected %s got %s", i, expectedSelect[i], v)
		}
	}

	grouped := GroupBy(q, func(t testStruct) string { return t.Name })
	if len(grouped) != 2 {
		t.Errorf("GroupBy failed, expected 2 groups got %d", len(grouped))
	}
	if len(grouped["a"]) != 2 || len(grouped["b"]) != 1 {
		t.Errorf("GroupBy groups length mismatch")
	}

	toMap := ToMap(q, func(t testStruct) int { return t.ID }, func(t testStruct) string { return t.Name })
	if len(toMap) != 3 {
		t.Errorf("ToMap failed, expected 3 items got %d", len(toMap))
	}
	if toMap[1] != "a" || toMap[2] != "b" || toMap[3] != "a" {
		t.Errorf("ToMap values incorrect")
	}
}
