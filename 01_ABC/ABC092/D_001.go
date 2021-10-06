package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	s := make([][]string, 100)
	for i := 0; i < 50; i++ {
		s[i] = make([]string, 100)
		for j := range s[i] {
			s[i][j] = "#"
		}
	}
	for i := 50; i < 100; i++ {
		s[i] = make([]string, 100)
		for j := range s[i] {
			s[i][j] = "."
		}
	}

	for i := 0; i < a-1; i++ {
		s[i/50*2][i%50*2] = "."
	}
	for i := 0; i < b-1; i++ {
		s[99-i/50*2][99-i%50*2] = "#"
	}

	fmt.Println(100, 100)
	for i := 0; i < 100; i++ {
		for j := range s[i] {
			fmt.Print(s[i][j])
		}
		fmt.Println()
	}
}
