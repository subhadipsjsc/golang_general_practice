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
	fmt.Printf("type of s is %T\n", s)
	fmt.Println("volume of s of interface type Shape is", s.Area())
	fmt.Printf("type of o is %T\n", o)
	fmt.Println("area of o of interface type Object is", o.Volume())
}

/*

	In the above program, we created Shape interface with Area method and
	Object interface with Volume method. Since struct type Cube implements
	both these method, it implements both these interfaces. Hence we can
	assign a value of struct type Cube to the variable of type Shape or Object.

	We expect s to have a dynamic value of c and o to also have a dynamic value of c.
	We used Area method on s of type Shape interface because it defines Area method and
	Volume method on o of type Object interface because it defines Volume method.
	But what will happen if we used Volume method on s and Area method on o.

*/
