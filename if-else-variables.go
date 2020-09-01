package main 

import (
	"fmt"
	"os"
)


func main() {
	fmt.Println(os.Args)
	if arg := os.Args[1]; arg == "firstname" {
		fmt.Println("Himanshu")
	} else if arg == "fullname" {
		fmt.Println("Himanshu Jagtap")
	} else {
		fmt.Println("Please enter correct argument from 'firstname' or 'fullname'")
	}
}
