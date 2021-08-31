package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		cnt := 0
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				cnt++
			} else {
				break
			}
		}
		fmt.Print(string(s[i]), cnt+1)
		i += cnt
	}
	fmt.Println()
}
