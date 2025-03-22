package main

import (
	"fmt"
	"math"
)

func CubeAreaAndVolume(edge int) (int, int) {
	return int(6*math.Pow(float64(edge),2)), int(math.Pow(float64(edge),3))
}

func main() {
	edge := 2
	fmt.Println("Cube with edge = ", edge)
	area, volume := CubeAreaAndVolume(edge)
	fmt.Println("area = ", area, "volume = ", volume)
}