//     *
//    **
//   ***
//  ****
// *****

package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 5

	for i := 1; i <= b; i++ {
		for j := b - 1; j >= i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= a; k++ {
			fmt.Print("*")
		}
		a++
		fmt.Print("\n")
	}
}
