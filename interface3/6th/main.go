package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube{3}
	var s Shape = c
	var o Object = c
	fmt.Println("area of s of interface type Shape is", s.Volume())
	fmt.Println("volume of o of interface type Object is", o.Area())
}

/*

	The program won’t compile because of the static type of s is Shape
	and of o is Object. To make it work, we need to somehow extract the
	underlying value of these interfaces. This can be done using type assertion.

*/
