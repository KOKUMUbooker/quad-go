package piscine

import "fmt"
// piscine.QuadC(5,3)
//ABBBA
//B   B
//CBBBC

// column = 5
// rows = 3

func QuadC(x,y int) {
	for i := 1;i <= 3;i ++ {
		for j :=1 ;j <= x; j ++{
			if i==1 &&  j==1{
				fmt.Printf("A")
			} else if i == 1 && j== x {
				fmt.Printf("A")
			} else if i == y && j == 1 {
				fmt.Printf("C")
			} else if  i==y && j == x {
				fmt.Printf("C")
			}else if j==1 && i>1 && i < y {
				fmt.Printf("B") 
			}else if j==x && i > 1 && i< y{
				fmt.Printf("B")
			}else if i==1 && j>1 && j< x{
				fmt.Printf("B")
			}else if i==y && j >1 && j< x{
				fmt.Printf("B")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
}
