package main

import (
	"fmt"
)

// var m map[string]Vertex
type person struct {
	name string
	age  int
}

// method adult()
func (v person) adult() string {
	if v.age > 18 {
		return "Adult"
	}
	return "not an adult"

}
func change(v *person) {
	v.age = v.age + 1
}
func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println(m["Answer"])
	delete(m, "Answer")
	fmt.Println(m["Answer"])
	m["Answer"] = 42
	v, ok := m["Answer"]
	fmt.Println(v, ok)
	sum := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(sum(4, 5))
	per := person{"himesh", 44}
	fmt.Println(per.adult())
	change(&per)
	fmt.Println(per.age)
}
