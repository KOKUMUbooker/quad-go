package piscine

import "fmt"

// eg piscine.QuadD(5,3)
func QuadD(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 { // top left corner
				fmt.Printf("A")
			} else if i == 1 && j == x { // top right corner
				fmt.Printf("C")
			} else if i == y && j == 1 { // bottom left corner
				fmt.Printf("A")
			} else if i == y && j == x { // bottom right corner
				fmt.Printf("C")
			} else if (i > 1 && i < y) && (j == 1 || j == x) { // left & right non corner borders
				fmt.Printf("B")
			} else if (i == 1 || i == y) && (j > 1 && j < x) { // top & bottom non corner borders
				fmt.Printf("B")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
