package main 

import "fmt"

// Variable initializer with type given
var i, j int = 1, 2

func main(){
	// Type inference from value of variable and unpacking of values
	var a, b, c = 10, true, 1.56
	fmt.Println("Declared type variables", i, j)
	fmt.Println("Non declared type variables", a, b, c)
}
