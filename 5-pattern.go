// ABCDE
// ABCDE
// ABCDE
// ABCDE
// ABCDE

package main

import (
	"fmt"
)

func main() {

	for i := 'A'; i <= 'E'; i++ {
		for j := 'A'; j <= 'E'; j++ {
			fmt.Printf("%c", j)
		}
		fmt.Print("\n")
	}
}
