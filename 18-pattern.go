// 	   *
// 	  ***
// 	 *****
// 	*******
// *********

package main

import (
	"fmt"
)

func main() {

	min_stars := 1
	p_height := 5
	p_space := p_height - 1

	for i := 0; i < p_height; i++ {
		for j := p_space; j > i; j-- {
			fmt.Print(" ")
		}

		for k := 0; k < min_stars; k++ {
			fmt.Print("*")
		}

		min_stars += 2
		fmt.Print("\n")
	}
}
