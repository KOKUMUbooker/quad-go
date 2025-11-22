package piscine

import "fmt"

// A - QuadA(5,3)
// o---o
// |   |
// o---o

func QuadA(x,y int) {
	// corners := "o"
	// xLine := "-"
	// yLine := "|"
	// space := " "

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			
			if (j == 1 || j == x) && !(i == 1 || i == y) { // Lines on y axis
				fmt.Printf("|")
			} else if j == 1 || j == x { // Corners
				fmt.Printf("o")
			} else if i == 1 || i == y { // lines on x axis
				fmt.Printf("-")
			} else {
				fmt.Printf(" ")
			}

		}
			fmt.Printf("\n")
	}
}