package piscine

import "fmt"
//ABBBC
//B   B
//CBBBA
//w=x,H=Y,i(y),r&j(5),c
//piscine.QuadE(5,3)
func QuadE(x,y int) {
	for i:=1; i<=y; i++ {
		for j:=1; j<=x; j++ {
			if i==1 && j==1{
				fmt.Printf("A")
			} else if i==1 && j == x{
				fmt.Printf("C")
			} else if i==y && j==1{
				fmt.Printf("C")
			} else if i==y && j ==x {
				fmt.Printf("A")
			} else if i==1 && j >1 && j<x {
				fmt.Printf("B")
			} else if i== y && j>1 && j < x {
				fmt.Printf("B")
			} else if j ==1 && i > 1 && i < x {
				fmt.Printf("B")
			} else if j ==x && i>1 && i< x {
				fmt.Printf("B")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

