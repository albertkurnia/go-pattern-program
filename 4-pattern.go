// AAAAA
// BBBBB
// CCCCC
// DDDDD
// EEEEE

package main

import (
	"fmt"
)

func main() {

	for i := 'A'; i <= 'E'; i++ {
		for j := 'A'; j <= 'E'; j++ {
			fmt.Printf("%c", i)
		}
		fmt.Print("\n")
	}
}
