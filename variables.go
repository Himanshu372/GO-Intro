package main 

import "fmt"

// Variable at package level
var atpackagelevel bool 

func main(){
	// Variable at function level
	var atfunclevel bool 
	fmt.Println(atpackagelevel, atfunclevel)
}
