package piscine

import "fmt"

// eg piscine.QuadE(5,3)
func QuadE(x, y int) {
	if x < 0 || y < 0 {
		return
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 { // Top Left corner
				fmt.Printf("A")
			} else if i == 1 && j == x { // Top right corner
				fmt.Printf("C")
			} else if i == y && j == 1 { // Bottom Left corner
				fmt.Printf("C")
			} else if i == y && j == x { // Bottom right corner
				fmt.Printf("A")
			} else if (i == 1 || i == y) && (j > 1 && j < x) { // Non-corner top & bottom borders
				fmt.Printf("B")
			} else if (i > 1 && i < x) && (j == 1 || j == x) { // Non-corner left & right borders
				fmt.Printf("B")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
