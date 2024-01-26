package main

import "fmt"

var ans int
var a, s [9]int

func ck(x, y, z int) bool {
	if s[x] < s[y] {
		x, y = y, x
	}
	if s[x] < s[z] {
		x, z = z, x
	}
	return !(a[y] == a[z] && a[x]^a[y] != 0)
}

func dfs(step int) {
	if step > 9 {
		if ck(0, 1, 2) && ck(3, 4, 5) && ck(6, 7, 8) && ck(0, 3, 6) && ck(1, 4, 7) && ck(2, 5, 8) && ck(0, 4, 8) && ck(2, 4, 6) {
			ans++
		}
		return
	}
	for i := 0; i < 9; i++ {
		if s[i] == 0 {
			s[i] = step
			dfs(step + 1)
			s[i] = 0
		}
	}
}

func main() {
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&a[3], &a[4], &a[5])
	fmt.Scan(&a[6], &a[7], &a[8])
	dfs(1)
	fmt.Printf("%.10f", float64(ans)/362880.0)
}
