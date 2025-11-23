package piscine

import "fmt"

// piscine.QuadB(3,4)
func QuadB(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 { // Upper Left corner
				fmt.Printf("/")
			} else if i == 1 && j == x { // Upper Right corner
				fmt.Printf("\\")
			} else if i == y && j == 1 { // Bottom left corner
				fmt.Printf("\\")
			} else if i == y && j == x { // Bottom right corner
				fmt.Printf("/")
			} else if (j > 1 && j < x) && (i == 1 || i == y) { // Top and bottom, between corners
				fmt.Printf("*")
			} else if (i > 1 && j < y) && (j == 1 || j == x) { // Non corner left and right borders
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
