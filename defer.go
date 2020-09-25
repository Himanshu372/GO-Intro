package main

import (
	"fmt"
)


func main() {
	//defer fmt.Println("world")
	//fmt.Println("hello")
	//time.Sleep(5 * time.Second)

	// Defered calls are stacked, last in first out order
	fmt.Println("Start of Deferred calls")
	for i := 0; i < 10; i++ {
		 defer fmt.Printf("%d call\n", i)
	}
	fmt.Println("Done")
}