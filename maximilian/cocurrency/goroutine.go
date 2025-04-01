package main

import (
	"fmt"
	"time"
)

func funcSlow(doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("slow execution...")
	doneChan <- true
}

func funcFast(doneChan chan bool) {
	fmt.Println("fast execution...")
	doneChan <- true
}

func main() {
	dones := make([]chan bool, 2) //channel : communication device.
	dones[0] = make(chan bool)
	dones[1] = make(chan bool)
	fmt.Println("main execution...")
	go funcSlow(dones[0])
	go funcFast(dones[1])
	for _, done := range dones {
		<- done //wait here to some data come out of channel.
	}
}
