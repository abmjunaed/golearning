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


//https://tour.golang.org/moretypes/23
// exercise-maps.go
//sol 1
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m:=make(map[string]int)
	for _,v:= range(words) {
		elem, ok:=m[v]
		if ok{
			m[v]=elem+1
		} else {
			m[v]=1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

//Sol 2
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m:=make(map[string]int)
	for _,v:= range(strings.Fields(s)) {// make a slice of the sentence accroding to the space
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

// fibonacci exercise with function closure
// https://tour.golang.org/moretypes/26
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i:=0
	f1:=0
	f2:=0
	var res int
	return func() int{
		if i==0 {
			i++
			return 0
		}else if(i==1){
			i++
			f1=1
			f2=0
			return 1
		}
		res=f1+f2
		f1,f2=f2,res
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
