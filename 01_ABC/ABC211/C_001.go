package main

import "fmt"

func main() {
	t := " chokudai"
	var s string
	fmt.Scan(&s)
	f := make([]int, 9)
	f[0] = 1
	for i := 0; i < len(s); i++ {
		for j := 1; j < 9; j++ {
			if s[i] == t[j] {
				f[j] = f[j] + f[j-1]
				f[j] %= 1000000007
			}
		}
	}
	fmt.Println(f[8] % 1000000007)
}
