package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	// input validation
	if x <= 0 || y <= 0 {
		return
	}

	// i is current row number

	for i :=1 ; i <= y ; i++ {
		// j is current col number
		for j := 1 ; j <= x ; j++ {
			// holds char to print
			var charToPrint rune = ''

			// identify current position is a border
			isTopRow := i == 1
			isBottomRow := i == y
			isLeftCol := j == 1
			isRightCol := j == x

			// Character Logic
			if isTopRow && isLeftCol {
				charToPrint = 'A'
			} else if isTopRow && isRightCol {
				charToPrint = 'C'
			} else if isBottomRow && isLeftCol {
				charToPrint = 'A'
			} else if isBottomRow && isRightCol{
				charToPrint = 'C'
			} else if isTopRow || isBottomRow || isLeftCol || isRightCol {
				charToPrint = 'B'
			} else {
				charToPrint = ' '
			}

			z01.PrintRune(charToPrint)
		}

		z01.PrintRune('\n')
	}
	
