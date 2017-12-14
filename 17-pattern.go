// *****
//  ****
//   ***
//    **
//     *

package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 5

	for i := 1; i <= a; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= b; k++ {
			fmt.Print("*")
		}
		b--
		fmt.Print("\n")
	}
}
