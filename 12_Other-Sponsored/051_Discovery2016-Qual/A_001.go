package main

import "fmt"

func main() {
	s := "DiscoPresentsDiscoveryChannelProgrammingContest2016"
	var n int
	fmt.Scan(&n)
	for i := 1; i <= 51; i++ {
		fmt.Print(string(s[i-1]))
		if i%n == 0 {
			fmt.Println()
		}
	}
	if 51%n != 0 {
		fmt.Println()
	}
}
