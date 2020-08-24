package main 

import (
	"fmt"
)

func add(x,  y int) int{
	return x + y
}

func main(){

	fmt.Println("Addition of 3 and 5 is", add(3, 5))
}
