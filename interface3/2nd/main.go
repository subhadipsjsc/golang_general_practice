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

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var s Shape
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	/*
		So in the above program, we created our Shape interface and rectangle
		struct type Rect.
		Then we defined methods like Area and Perimeter with Rect receiver type.
		Hence Rect implemented those methods. Since these methods are defined by
		Shape interface,
		--  Rect implements Shape interface. ----
	*/
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectange s", s.Area())
	fmt.Println("s == r is", s == r)
}

/*
	We can confirm that by creating nil interface s and
	assigning struct of type Rect to it that.
	Since Rect implements Shape interface, this is perfectly valid.
	From the above result, we can see that, dynamic type of s is now Rect
	and dynamic value of s is the value of the struct Rect which is {5 4}.
	Dynamic because we can assign new struct to it of a different type which
	also implements the interface Shape.
*/
