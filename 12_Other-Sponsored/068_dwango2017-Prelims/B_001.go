package main

import "fmt"

func main() {
	var st string
	fmt.Scan(&st)
	var f [233333][3]int
	ans := 0
	for i := 0; i < len(st); i++ {
		if st[i] == '2' || st[i] == '?' {
			f[i+1][1] = f[i][2] + 1
		}
		if st[i] == '5' || st[i] == '?' {
			if f[i][1] > 0 {
				f[i+1][2] = f[i][1] + 1
			}
		}
		ans = max(ans, f[i+1][2])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
