package main

import (
	"fmt"
	"math"
)

func cubeArea(edge int) int {
	return int(6 * math.Pow(float64(edge), 2))
}

func cubeVolume(edge int) int {
	return int(math.Pow(float64(edge), 3))
}

func arithmaticSelector(selector string) func(int) int {
	switch selector {
	case "area":
		return cubeArea
	case "volume":
		return cubeVolume
	default:
		return cubeArea
	}
}

func main() {
	area := arithmaticSelector("area")
	fmt.Println("area of cube with edge=2 is: ", area(2))

	vol := arithmaticSelector("volume")
	fmt.Println("volume of cube with edge=4 is: ", vol(2))

	defaultArea := arithmaticSelector("arithmatic")
	fmt.Println("area of cube with edge=3 is: ", defaultArea(3))
}
