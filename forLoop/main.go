package main

import (
	"fmt"
)

func main() {

	// simple for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------------------")

	// forloop with 2 variables incremented simultaneously
	for i, j := 0, 1; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println("------------------------")

	// simple for loop , with initializer is stated seperately
	// the value of i canbe used outside forloop , as it was initialized before forloop
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------------------")

	// simple for loop , acts as an Do-While loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	fmt.Println("------------------------")

	// forloop with break statement
	k := 0
	for {
		fmt.Println(k)
		k++
		if k > 5 {
			break
		}
	}
	fmt.Println("------------------------")

	// simple for loop , with continue method
	//this continue will escape the current loop , and forwared for the next loop
	for a := 0; a < 5; a++ {
		if a%2 == 0 {
			continue // so all the even numbers will be excaped
		}
		fmt.Println(a)
	}
	fmt.Println("------------------------")

	// if there is nested loop , we can use lable: Loop to break
	//the inner loops to the label.

Loop:
	for b := 1; b < 5; b++ {
		for c := 1; c < 5; c++ {
			fmt.Println("Result is :", b*c)
			if b*c > 6 {
				break Loop
			}

		}
	}
	fmt.Println("------------------------")

	//foreach loop for slices , arrays ,maps  with key and value
	s := []int{2, 4, 6, 8, 10}
	for k, v := range s {
		fmt.Println("Key is: ", k, " & value is : ", v)
	}
	// if you donot want key , only value , then make first parameter _
	for _, val := range s {
		fmt.Println("value is : ", val)
	}

	fmt.Println("------------------------")

	//foreach loop for strings
	str := "abcdefgh"
	for k1, v1 := range str {
		fmt.Println("Key is: ", k1, " & value is : ", string(v1), " & unicode value :", v1)
	}

}
