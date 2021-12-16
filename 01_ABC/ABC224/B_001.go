package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	a := [1003][1003]int{}
	for i := 1; i < h+1; i++ {
		for j := 1; j < w+1; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	for i := 2; i < h+1; i++ {
		for j := 2; j < w+1; j++ {
			if a[i][j]+a[i-1][j-1] > a[i-1][j]+a[i][j-1] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
