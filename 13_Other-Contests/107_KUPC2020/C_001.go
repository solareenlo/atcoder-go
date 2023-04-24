package main

import "fmt"

func main() {
	var C [13][13]int
	for i := 1; i < 13; i++ {
		C[i][0] = i
	}
	for j := 1; j < 13; j++ {
		C[0][j] = j + 13
	}
	for i := 1; i < 13; i++ {
		for j := 1; j < 13; j++ {
			C[i][j] = (C[0][j] * C[i][0]) % 26
		}
	}
	for j := 0; j < 13; j += 2 {
		C[0][j] = 0
	}
	fmt.Println(13)
	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			fmt.Print(string('a' + C[i][j]))
		}
		fmt.Println()
	}
}
