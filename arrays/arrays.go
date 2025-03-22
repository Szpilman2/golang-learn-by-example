package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	c := [...]int{2,3,4,5,6,7,8,9} //let compiler count number of elements.
	fmt.Println("dcl:", c)
	fmt.Println("len:",len(c))

	b = [...] int{100, 3:400, 500} //put 400 in the index 3 of array.
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i:=0; i<2;i++{
		for j:=0; j<3;j++{
			twoD[i][j] = i*i + j*j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int {
		{1,3,5},
		{2,4,6},
	}
	fmt.Println("2d: ", twoD)

	threeD := [2][2][3] int{
		{
			{1,2,3},
			{1,2,4},
		},
		{
			{4,5,6},
			{7,8,9},
		},
	}
	fmt.Println("3d: ", threeD)
}
