package main

import (
	"fmt"
)

/*
	this rectangle will work as a class .
	we will then call an instance of rectangle inside main() fnction
	react1 := rectangle{5, 8}

	this is not actually instanciation , as go lang has no classes

	but , we make a new type of variable , called rectangle
	this rectangle will work as other type of variable types like : int , float , bool etc.

	and react1  will be a variable of type rectangle



*/

type rectangle struct {
	length float32
	width  float32
}

/*
	here this function Area() is a method of type rectangle .
	rectangle is not actually a class , but for simplicity , we can think , rectangle as a class.
	And Area() function will be a method of rectangle calls .

	that is why we put (r rectangle) before Area() function
	Now , the function will take r as a instance of rectangle class , and will calculate
	area = r.length x r.width

*/

func (r rectangle) Area() float32 {
	return (r.length * r.width)
}

func main() {
	react1 := rectangle{5, 8}
	fmt.Println(react1.Area())
}
