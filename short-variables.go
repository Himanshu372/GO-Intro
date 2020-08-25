package main

import "fmt"

func main(){
	// Short declaration 
	k := 3
	fmt.Println("Inside func main", k)
	// This type of declaration is only available inside functions and not in package
	a, b, c := true, false, "CS"
	fmt.Println("Short variable declarations", a, b, c)
}
