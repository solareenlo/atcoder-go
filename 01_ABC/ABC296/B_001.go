package main

import "fmt"

func main() {
	var s [9]string
	fmt.Scan(&s[8])
	fmt.Scan(&s[7])
	fmt.Scan(&s[6])
	fmt.Scan(&s[5])
	fmt.Scan(&s[4])
	fmt.Scan(&s[3])
	fmt.Scan(&s[2])
	fmt.Scan(&s[1])
	for i := 1; i <= 8; i++ {
		for j := 0; j < 8; j++ {
			if s[i][j] == '*' {
				fmt.Print(string('a'+j), i, "\n")
			}
		}
	}
}
