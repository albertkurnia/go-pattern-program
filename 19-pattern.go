//     1
//    333
//   55555
//  7777777
// 999999999

package main

import (
	"fmt"
)

func main() {

	min_stars := 1

	for i := 0; i < 5; i++ {
		for j := 4; j > i; j-- {
			fmt.Print(" ")
		}

		for k := 0; k < min_stars; k++ {
			fmt.Print(min_stars)
		}

		min_stars += 2
		fmt.Print("\n")
	}
}
