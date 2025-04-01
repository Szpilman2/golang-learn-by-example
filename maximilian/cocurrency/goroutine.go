package main

import (
	"fmt"
	"time"
)

func funcSlow() {
	fmt.Println("slow execution...")
	time.Sleep(3*time.Second)
}

func funcFast() {
	fmt.Println("fast execution...")
}

func main() {
	fmt.Println("main execution...")
	funcSlow()
	funcFast()
}