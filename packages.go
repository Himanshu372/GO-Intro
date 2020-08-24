package main 

import (
	"fmt"
	"math/rand"
)

func main(){
	rand.Seed(4)
	fmt.Println("Favourite number is", rand.Int())
}
