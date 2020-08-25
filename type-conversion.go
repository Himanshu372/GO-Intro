package main 

import "fmt"

func main(){
	i := 2
	f := float64(i)
	z := uint(f)
	fmt.Println("Int variable", i)
	fmt.Println("Converted to float", f)
	fmt.Println("Converted to unsigned int", z)

}
