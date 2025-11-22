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
	h := x // width
	w := y // height

	for i := 1; i <= w; i++ {
		for j := 1; j <= h; j++ {
			
			if (j == 1 || j == h) && !(i == 1 || i == w) { // Lines on y axis
				fmt.Printf("|")
			} else if j == 1 || j == h { // Corners
				fmt.Printf("o")
			} else if i == 1 || i == w { // lines on x axis
				fmt.Printf("-")
			} else {
				fmt.Printf(" ")
			}

		}
			fmt.Printf("\n")
	}
}