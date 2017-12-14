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

	for i := a; i >= 1; i-- {
		for j := a - 1; j >= i; j-- {
			fmt.Print(" ")
		}

		for k := i; k >= 1; k-- {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
}
