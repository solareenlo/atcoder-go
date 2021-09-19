package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	s = s + string(s[0]) + string(s[1])

	res := make([]int, 100005)
	k := 0
	for ; k < 4; k++ {
		res[0] = k & 1
		res[1] = k >> 1
		for i := 2; i < n+2; i++ {
			if s[i-1] == 'o' {
				res[i] = res[i-2] ^ 0 ^ res[i-1]
			} else {
				res[i] = res[i-2] ^ 1 ^ res[i-1]
			}
		}
		if res[n] == res[0] && res[n+1] == res[1] {
			for i := 0; i < n; i++ {
				if res[i] != 0 {
					fmt.Print("W")
				} else {
					fmt.Print("S")
				}
			}
			fmt.Println()
			return
		}
	}
	fmt.Println(-1)
}
