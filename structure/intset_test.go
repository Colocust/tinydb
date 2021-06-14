package structure

import (
	"fmt"
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	is := NewIntSet()
	for i := 0; i <= 20; {
		is.Add(i)
		i += 2
	}
	for i := 1; i < 20; {
		is.Add(i)
		i += 2
	}
	fmt.Println(is.contents)
}

func TestIntSet_Remove(t *testing.T) {
	is := NewIntSet()
	for i := 0; i <= 20; {
		is.Add(i)
		i += 2
	}
	for i := 1; i < 20; {
		is.Add(i)
		i += 2
	}
	fmt.Println(is.contents)
	for i := 1; i < 20; {
		is.Remove(i)
		i += 2
	}
	fmt.Println(is.contents)
}

func TestIntSet_Get(t *testing.T) {
	is := NewIntSet()
	for i := 0; i <= 20; i++ {
		is.Add(i)
	}
	fmt.Println(is.contents)
	fmt.Println(is.Get(1))
}

func TestIntSet_Set(t *testing.T) {
	is := NewIntSet()
	for i := 0; i < 20; i++ {
		is.Add(i)
	}
	fmt.Println(is.contents)
	is.Set(19, 21)
	fmt.Println(is.contents)
}

func TestIntSet_Find(t *testing.T) {
	is := NewIntSet()
	for i := 0; i <= 20; i++ {
		is.Add(i)
	}
	for i := 0; i <= 25; i++ {
		fmt.Println(is.Find(i))
	}
}

func TestIntSet_Random(t *testing.T) {
	is := NewIntSet()
	for i := 0; i <= 20; i++ {
		is.Add(i)
	}
	for i := 0; i <= 20; i++ {
		fmt.Println(is.Random())
	}
}