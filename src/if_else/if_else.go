package main

import (
	"fmt"
	"math"
)

type tangle struct {
	base   int
	height int
}

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}
func printSlice(s string, x []int) {
	fmt.Printf("\n%s len = %d cap = %d %v", s, len(x), cap(x), x)
}

var a [10]int

// ===========Slice================
// The type []T is a slice with elements of type T.
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 9),
	)
	t := tangle{5, 3}
	fmt.Println("base = ", t.base)
	a[0] = 5
	a[1] = 6
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	names := [4]string{
		"Tom",
		"John",
		"Ron",
		"Son",
	}
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a)
	fmt.Println(b)
	a[1] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	sliceA := make([]int, 5)
	printSlice("a", sliceA)
	sliceB := make([]int, 0, 5)
	printSlice("b", sliceB)
	sliceC := sliceB[:2]
	printSlice("c", sliceC)
	// ==============Append slice=========
	var s []int
	printSlice("s", s)
	s = append(s, 0)
	printSlice("s", s)
	s = append(s, 1)
	printSlice("s", s)
	s = append(s, 2, 3, 4)
	printSlice("s", s)
	for i, v := range names {
		fmt.Println(i, v)
	}
}
