// *
// **
// ***
// ****
// ***
// **
// *

package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3

	for i := a; i >= -a; i-- {
		for j := a; j >= int(math.Abs(float64(i))); j-- {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
