package main 

import (
	"fmt"
	"time"
)

func main() {
	
	const format = "2011-09-27 22:46"
	t, _ := time.Parse(format, "2014-09-14 14:25")
	fmt.Println(t)
}
