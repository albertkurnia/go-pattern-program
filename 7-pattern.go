// EEEEE
// DDDDD
// CCCCC
// BBBBB
// AAAAA

package main

import (
	"fmt"
)

func main() {

	for i := 'E'; i >= 'A'; i-- {
		for j := 'E'; j >= 'A'; j-- {
			fmt.Printf("%c", i)
		}
		fmt.Print("\n")
	}
}
