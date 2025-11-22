package piscine

import "fmt"

// B - QuadB(5,3) x = 5, y = 3
// /***\
// *   *
// \***/

func QuadB(x, y int) {
	// corners := "o"
	// xLine := "*"
	// yLine := "*"
	// space := " "

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (j == 1 || j == x) && !(i == 1 || i == y) { // Lines on y axis
				fmt.Printf("*")
			} else if j == 1 || j == x { // Corners

				if j == 1 && i == 1 { // Top left corner
					fmt.Printf("/")
				} else if j == x && i == 1 { // Top right corner
					fmt.Printf("\\")
				} else if j == 1 && i == y { // Bottom left corner
					fmt.Printf("\\")
				} else { // Bottom right corner
					fmt.Printf("/")
				}
			} else if i == 1 || i == y { // lines on x axis
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
