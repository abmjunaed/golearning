// https://tour.golang.org/flowcontrol/8 
// exercise-loops-and-functions.go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z:=1.0
	diff:=z
	for i:=0;i<10;i++{
		z-=(z*z-x)/(2*z)
		fmt.Printf("%g %d \n",z, i)
		
		if math.Abs(diff-z)<1e-7 {
			break;
		}
		diff=z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("Diff with lib: %g", math.Sqrt(2)-Sqrt(2))
	fmt.Println(Sqrt(4))
	fmt.Println("Diff with lib: %g", math.Sqrt(4)-Sqrt(4))
}
