package main

import "fmt"

func main() {
	a := [6][6]int{}
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Scan(&a[i][j])
			if a[i][j] == a[i-1][j] || a[i][j] == a[i][j-1] {
				fmt.Println("CONTINUE")
				return
			}
		}
	}
	fmt.Println("GAMEOVER")
}
