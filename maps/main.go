package main

import (
	"fmt"
)

func main() {
	statePopulation := map[string]int{
		"UP":         12000000,
		"Bihar":      80000000,
		"MP":         90000000,
		"WestBengal": 80000000,
		"Delhi":      35000000,
	}

	fmt.Println(statePopulation["MP"])

	// add new item inthe map
	statePopulation["punjab"] = 70000000
	fmt.Println(statePopulation)

	// delete element from map
	delete(statePopulation, "punjab")
	fmt.Println(statePopulation)

	//even if an element is not in the MAP , if we search that elemet , result will be zero
	fmt.Println(statePopulation["punjab"])

	// to verify , if element is actually available , we use " comma Ok " syntax
	mapResult, ok := statePopulation["punjab"]
	fmt.Println(mapResult, ok)

	_, ok2 := statePopulation["punjab"] // if we only want to know , if he emenet is there or not
	fmt.Println(ok2)

	mapResult3, ok3 := statePopulation["MP"]
	fmt.Println(mapResult3, ok3)

	/*
		MAps are passed by reference , so they are mutable
		if the copied MPA value changes , the Original source MAP will
		also chenge
	*/

	anotherPopulation := statePopulation
	anotherPopulation["MP"] = 500
	fmt.Println(statePopulation["MP"])

}
