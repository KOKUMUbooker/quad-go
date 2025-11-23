package piscine

import "fmt"

// piscine.QuadC(3,2)
func QuadC(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && (j == 1 || j == x) { // Top right & left corners
				fmt.Printf("A")
			} else if i == y && (j == 1 || j == x) { // bottom left & right corners
				fmt.Printf("C")
			} else if (i > 1 && i < y) && (j == 1 || j == x) { // left & right non corner borders
				fmt.Printf("B")
			} else if (j > 1 && j < x) && (i == 1 || i == y) { // top & bottom non corner borders
				fmt.Printf("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
}
