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

/*
	We can find out the underlying dynamic value of an
	interface using the syntax i.(Type) where i is an interface and
	Type is a type that implements the interface i. Go will check if
	dynamic type of i is identical to Type.
*/

func main() {
	c := Cube{3}
	var s Shape = c
	cu := s.(Cube)

	fmt.Printf("type of cu is %T\n", cu)
	fmt.Println("area of c of type Cube is", cu.Area())
	fmt.Println("volume of c of type Cube is", cu.Volume())
}

/*
	From the above program, we now have access to the underlying value of the
	interface s in variable cu which is a structure of type Cube. Now, we can use
	Area and Volume methods on cu.
*/
