package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func main() {
	var s Shape = Rect{10, 3}
	fmt.Println(s)
}

/*
	In the above program, we removed Perimeter method.
	Above program won’t compile and the compiler will throw an error

	It should be obvious from the above error that in order to successfully
	implement an interface, you need to implement all methods declared by the interface.

*/
