package piscine

import "fmt"

// eg piscine.QuadA(3,2)
func QuadA(x, y int) {
	if x < 0 || y < 0 {
		return
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && (j == 1 || j == x) { // Top corner - right & left
				fmt.Printf("o")
			} else if i == y && (j == 1 || j == x) { // Bottom corner - left & right
				fmt.Printf("o")
			} else if (j > 1 && j < x) && (i == 1 || i == y) { // Between the 2 corner 'o'
				fmt.Printf("-")
			} else if (i > 1 && i < y) && (j == 1 || j == x) { // Between the 1st & last rows and 1st or last column
				fmt.Printf("|")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
