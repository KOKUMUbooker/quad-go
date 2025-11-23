package piscine

import "fmt"

// piscine.QuadA( x = 5, y = 3)
// o---o
// |   |
// o---o

// columns = 5
// rows = 3

func QuadA(x,y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				fmt.Printf("o")
			} else if i == 1 && j == x {
				fmt.Printf("o")
			} else if i == y && j == 1 {
				fmt.Printf("o")
			} else if i == y && j == x {
				fmt.Printf("o")
			} else if i == 1 && j > 1 && j < x {
				fmt.Printf("-")
			} else if i == y && j > 1 && j < x {
				fmt.Printf("-")
			} else if i > 1 && i < y && j == 1 {
				fmt.Printf("|")
			} else if  i > 1 && i < y && j == x {
				fmt.Printf("|")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Printf("\n")
	}
}