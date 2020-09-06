package main

import (
	"fmt"
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, pre int) float64{
	output := math.Pow(10, float64(pre))
	return float64(round(num * output)) / output

}

func Sqrt(x float64) float64 {
	z := 1.0
	for z <= x {
		pre := toFixed(z, 3)
		z -= (z*z - x) / (2 * z)
		after := toFixed(z, 3)
		if pre == after {
			break
		}
		fmt.Println(z)
	}
	fmt.Println(z)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
