package main

import "fmt"

func main() {
	var x, y, z [10][10]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var a int
			fmt.Scan(&a)
			x[i][a] = 1
			y[j][a] = 1
			z[i/3*3+j/3][a] = 1
		}
	}
	for i := 0; i < 9; i++ {
		for j := 1; j <= 9; j++ {
			if !(x[i][j] != 0 && y[i][j] != 0 && z[i][j] != 0) {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
