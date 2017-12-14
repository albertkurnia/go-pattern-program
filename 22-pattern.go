// 	  *
// 	 **
// 	***
// ****
// 	***
// 	 **
// 	  *

package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3

	for i := a; i >= -a; i-- {
		for j := 1; j <= int(math.Abs(float64(i))); j++ {
			fmt.Print(" ")
		}

		for k := a; k >= int(math.Abs(float64(i))); k-- {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
