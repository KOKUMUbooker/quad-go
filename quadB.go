package piscine

import "fmt"

// piscine.QuadB(5,3)
// /***\
// *   *
// \***/

// rows = 3
// column = 5

func QuadB(x,y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				fmt.Printf("/")
			} else if i == 1 && j == x {
				fmt.Printf("\\")
			} else if i == y && j == 1 {
				fmt.Printf("\\")
			} else if i == y && j == x {
				fmt.Printf("/")
			} else if i == 1 && j > 1 && j < x {
				fmt.Printf("*")
			} else if i == y && j > 1 && j < x {
				fmt.Printf("*")
			} else if j == 1 && i > 1 && i < y {
				fmt.Printf("*")
			} else if j == x && i > 1 && i < y {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}