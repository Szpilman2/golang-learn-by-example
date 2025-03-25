package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// I should Investigate this part more.
	nexInt := intSeq()
	fmt.Println(nexInt())

	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())

	newInts := intSeq()
	fmt.Println(newInts())
}