package main

import (
	"fmt"
	"maps"
	"reflect"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("type(v1):", reflect.TypeOf(v1))

	v3 := m["k3"] //k3 is not defined as a key in map so it's value will be zero.
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	_, prs := m["k2"]
	fmt.Println("prs:", prs) 

	n2 := map[string]int{"foo":1, "bar":2}
	fmt.Println("map", n2)

	n := map[string]int{"foo":1, "bar":2}
	if maps.Equal(n, n2){
		fmt.Println("n == n2")
	}
}