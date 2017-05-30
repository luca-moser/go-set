package set

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestInit(t *testing.T) {
	set := NewSet("apple", "orange", "lemon", "banana")
	if set.Length() != 4 {
		t.Fatalf("expected set length to be %d but was %d", 4, set.Length())
	}
}

func TestAdd(t *testing.T) {
	set := NewSet("apple", "orange", "lemon", "banana")

	// add some more values but it shouldn't actually increase the set's size
	added := set.Add("apple", "banana")

	if added != 0 {
		t.Fatalf("expected Add() to add %d entries but actual amount was %d", 0, added)
	}
	if set.Length() != 4 {
		t.Fatalf("expected set length to still be %d but was %d", 4, set.Length())
	}
}

func TestRemove(t *testing.T) {
	set := NewSet("apple", "orange", "lemon", "banana")

	removed := set.Remove("apple", "orange")
	if removed != 2 {
		t.Fatalf("expected Remove() to remove some %d entries but actual amount was %d", 2, removed)
	}

	removed = set.Remove("grapefruit")
	if removed != 0 {
		t.Fatalf("expected second Remove() to remove %d entries but actual amount was %d", 0, removed)
	}

	if set.Length() != 2 {
		t.Fatalf("expected set length to be %d after removal but was %d", 2, set.Length())
	}
}

func TestJoin(t *testing.T) {
	set := NewSet("apple", "orange", "lemon", "banana")
	set2 := NewSet("python", "java", "ruby", "golang")

	// join set2 into set1
	set.Join(set2)

	if set.Length() != 8 {
		t.Fatalf("expected set length to be %d after join but was %d, set data: %v", 8, set.Length(), set)
	}

	if !set.Has("ruby") {
		t.Fatalf("expected set to have entry 'ruby', set data: %v", set)
	}

	if !set.Has("golang") {
		t.Fatalf("expected set to have entry 'golang', set data: %v", set)
	}

}

func TestHas(t *testing.T) {
	set := NewSet("apple", "orange", "lemon", "banana")

	if set.Has("grapefuit") {
		t.Fatalf("expected set to not have the entry 'grapefruit'")
	}

	if !set.Has("lemon") {
		t.Fatalf("expected set to have the entry 'lemon'")
	}
}

func TestClear(t *testing.T) {
	set := NewSet("apple", "orange", "lemon", "banana")
	if set.Length() != 4 {
		t.Fatalf("expected set length to be 4 but was %d", set.Length())
	}
	set.Clear()

	if set.Length() != 0 {
		t.Fatalf("expected set length to be empty but was %d", set.Length())
	}
}

func TestMixed(t *testing.T) {
	set := NewSet(123, false, false, "orange", float64(3.04334), 123, struct{ name string }{"vanessa"})
	if set.Length() != 5 {
		t.Fatalf("expected set length to be 5 but was %d", set.Length())
	}
	set.Add(struct{ name string }{"alice"})

	if set.Length() != 6 {
		t.Fatalf("expected set length to be 6 but was %d", set.Length())
	}

	if !set.Has(struct{ name string }{"vanessa"}) {
		t.Fatalf("expected set to have the struct entry with name 'vanessa'")
	}

	if !set.Has(false) {
		t.Fatalf("expected set to have the entry 'false'")
	}
}

func TestCasted(t *testing.T) {
	set := NewSet(1, 2, 3, "a", "b", "c", "d", "1", "2", "3")
	set.Add(float32(1.23), float32(5.34), float64(4.34), float64(6.342), float64(23.23))

	if len(set.Strings()) != 7 {
		t.Fatalf("expected the set's string slice to have length 4 but was %d", len(set.Strings()))
	}

	if len(set.Ints()) != 3 {
		t.Fatalf("expected the set's int slice to have length 3 but was %d", len(set.Ints()))
	}

	if len(set.Floats32()) != 2 {
		t.Fatalf("expected the set's float32 slice to have length 2 but was %d", len(set.Floats32()))
	}

	if len(set.Floats64()) != 3 {
		t.Fatalf("expected the set's float64 slice to have length 3 but was %d", len(set.Floats64()))
	}

}

type person struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	persons := []person{
		{"alice", 21},
		{"bob", 30},
		{"luca", 22},
	}
	set := NewSet(persons[0], persons[1], persons[2])

	if set.Length() != 3 {
		t.Fatalf("expected the set length to be 3 but was %d", set.Length())
	}

	if !set.Has(persons[1]) {
		t.Fatalf("expected the set to contain the second entry of the slice", set.Length())
	}
}
