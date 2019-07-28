package main

import "fmt"

type Rock struct {
	Mass   int
	Volume int
}

func (r Rock) Density() int {
	return r.Mass / r.Volume
}

type Geode struct {
}

func (g Geode) Density() int {
	return 100
}

type Dense interface {
	Density() int
}

func IsItDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

func main() {
	r := Rock{10, 1}
	g := Geode{}
	// Returns true because Geode's Density method always
	// returns 100
	fmt.Println(IsItDenser(g, r))
}
