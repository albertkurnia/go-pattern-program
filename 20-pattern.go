// *******
//  *****
//   ***
//    *

package main

import (
	"fmt"
)

func main() {

	max_stars := 7
	p_space := max_stars / 2
	p_height := max_stars - p_space

	for i := p_height; i >= 1; i-- {
		for j := p_space; j >= i; j-- {
			fmt.Print(" ")
		}

		for k := 1; k <= max_stars; k++ {
			fmt.Print("*")
		}

		max_stars -= 2
		fmt.Print("\n")
	}
}
