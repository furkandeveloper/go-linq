package main

import (
	"fmt"
	"github.com/furkandeveloper/golinq/linq" // Burada kendi linq paketinin yolu, aynı modül içindeyse "linq" ya da "./linq" olabilir
)

// Person tipi
type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	people := []Person{
		{"Alice", 30, "Female"},
		{"Bob", 25, "Male"},
		{"Charlie", 35, "Male"},
		{"Diana", 22, "Female"},
		{"Eve", 29, "Female"},
		{"Frank", 30, "Male"},
		{"Grace", 35, "Female"},
	}

	// LINQ query başlat
	query := linq.From(people)

	// 1. WhereGroup - AND operatörüyle 30 yaşından büyük ve kadın olanlar
	filtered := query.WhereGroup(linq.PredicateGroup[Person]{
		Predicates: []func(Person) bool{
			func(p Person) bool { return p.Age > 30 },
			func(p Person) bool { return p.Gender == "Female" },
		},
		LogicalOperator: linq.And,
	})

	fmt.Println("30 yaşından büyük kadınlar:")
	for _, p := range filtered.ToSlice() {
		fmt.Printf("  %+v\n", p)
	}

	// 2. Select - sadece isimleri al
	names := linq.Select(filtered, func(p Person) string { return p.Name })

	fmt.Println("\nSeçilen isimler:")
	for _, name := range names.ToSlice() {
		fmt.Println(" -", name)
	}

	// 3. Max yaş
	maxAge := query.Max(func(p Person) int { return p.Age })
	fmt.Println("\nEn büyük yaş:", maxAge)

	// 4. GroupBy cinsiyete göre gruplama
	grouped := linq.GroupBy(query, func(p Person) string { return p.Gender })
	fmt.Println("\nCinsiyete göre gruplar:")
	for gender, group := range grouped {
		fmt.Printf("%s:\n", gender)
		for _, p := range group {
			fmt.Printf("  %+v\n", p)
		}
	}

	// 5. ToMap isim -> yaş map
	nameAgeMap := linq.ToMap(query, func(p Person) string { return p.Name }, func(p Person) int { return p.Age })
	fmt.Println("\nİsim -> Yaş map:")
	for name, age := range nameAgeMap {
		fmt.Printf("  %s: %d\n", name, age)
	}

	// 6. OrderBy - yaşa göre artan sıralama
	sortedByAge := query.OrderBy(func(a, b Person) bool { return a.Age < b.Age })
	fmt.Println("\nYaşa göre artan sıralama:")
	for _, p := range sortedByAge.ToSlice() {
		fmt.Printf("  %+v\n", p)
	}

	// 7. Distinct örneği - yaşlara göre distinct (eşitlik fonksiyonu)
	distinctByAge := query.Distinct(func(a, b Person) bool { return a.Age == b.Age })
	fmt.Println("\nDistinct yaşlar:")
	for _, p := range distinctByAge.ToSlice() {
		fmt.Printf("  %+v\n", p)
	}

	// 8. Any ve All örneği
	hasTeen := query.Any(func(p Person) bool { return p.Age >= 13 && p.Age < 20 })
	fmt.Println("\nGenç (13-19) var mı?:", hasTeen)

	allAdults := query.All(func(p Person) bool { return p.Age >= 18 })
	fmt.Println("Herkes yetişkin mi?:", allAdults)
}
