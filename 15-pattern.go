// ABCDE
// ABCD
// ABC
// AB
// A

package main

import (
	"fmt"
)

func main() {

	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf("%c", j+65)
		}
		fmt.Print("\n")
	}
}
