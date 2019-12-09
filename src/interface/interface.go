package main

import (
	"fmt"
	"time"
)

type I interface {
	M()
}
type T struct {
	S string
}
type Person struct {
	name string
	age  int
}

func (t *T) M() {
	fmt.Println(t.S)
}

// A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

// (As with fmt.Stringer, the fmt package looks for the error interface when printing values.)
type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}
func run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}
func main() {
	var i I
	i = &T{"Hello"}
	i.M()
	a := Person{"Rahul", 25}
	b := Person{"Karan", 26}
	fmt.Println(a, b)
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
