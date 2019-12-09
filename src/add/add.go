package main

import "fmt"

// Outside a function, every statement begins with a keyword (var, func, and so on)
// and so the := construct is not available.
// func add(a,b int) int could also be used where a nad b both are int
// The expression T(v) converts the value v to the type T.
// Constants cannot be declared using the := syntax.
var i int = 5
var f float64 = float64(i)

func add(a int, b int) int {
	return a + b
}
func swap(a, b string) (string, string) {
	return b, a
}
func main() {
	fmt.Println(add(42, 3))
	var a string = "hello"
	var b string = "world"
	a, b = swap(a, b)
	fmt.Println(a, b)
	fmt.Printf("\n f is of type %T : ", f)
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
}
