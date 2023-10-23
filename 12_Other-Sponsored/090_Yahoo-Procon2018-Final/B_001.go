package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	var a [80][80]string
	for i := 0; i < 78; i++ {
		for j := 0; j < 78; j++ {
			a[i][j] = "#"
		}
	}
	for i := 0; i < 79; i++ {
		a[i][78] = "."
		a[78][i] = "."
	}
	for i := 0; i < 38; i, k = i+1, k/3 {
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if x^2 != 0 || y^1 != 0 {
					a[i*2+x][i*2+y] = "."
				}
			}
		}
		t := k % 3
		if t != 0 {
			for x := i * 2; x < 78; x++ {
				a[i*2][x] = "."
			}
		}
		if t > 1 {
			for x := i * 2; x < 78; x++ {
				a[x][i*2] = "."
			}
		}
	}
	fmt.Println(79, 79)
	for i := 0; i < 79; i++ {
		for j := 0; j < 79; j++ {
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}
}
