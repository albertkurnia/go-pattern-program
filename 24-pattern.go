//     *
//    * *
//   * * *
//  * * * *
// * * * * *
//  * * * *
//   * * *
//    * *
//     *

package main

import (
	"fmt"
)

func main() {

	a := 5

	for i := 1; i <= a; i++ {
		for j := a - 1; j >= i; j-- {
			fmt.Print(" ")
		}

		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}

	for b := a - 1; b >= 1; b-- {
		for c := a - 1; c >= b; c-- {
			fmt.Print(" ")
		}

		for d := b; d >= 1; d-- {
			fmt.Print("* ")
		}

		fmt.Print("\n")
	}
}
