package main

import "fmt"

// if the function has return you should write it as output.
// when the return value of function is void then you should not explicitly write function output.
func sum(nums ...int) {
	fmt.Print(nums, " -> ")
	total := 0
	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main()  {
	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4,5}
	sum(nums...) //converts array to parameters.
}