// 11111
// 22222
// 33333
// 44444
// 55555

package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
}
