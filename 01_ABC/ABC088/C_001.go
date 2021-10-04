package main

import "fmt"

func main() {
	c := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&c[i][j])
		}
	}

	for i := 0; i < 3-1; i++ {
		for j := 0; j < 3-1; j++ {
			if c[i][j]-c[i][j+1] != c[i+1][j]-c[i+1][j+1] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
