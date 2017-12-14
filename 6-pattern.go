// 55555
// 44444
// 33333
// 22222
// 11111

package main

import (
	"fmt"
)

func main() {

	for i := 5; i >= 1; i-- {
		for j := 1; j <= 5; j++ {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
}
