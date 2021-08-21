package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	e := [12][12]bool{}
	for x, y, i := 0, 0, 0; i < m; i++ {
		fmt.Scan(&x, &y)
		e[x-1][y-1] = true
		e[y-1][x-1] = true
	}

	res := 1
	for bit := 0; bit < 1<<n; bit++ {
		flag := true
		for i := 0; i < n-1; i++ {
			if bit&(1<<i) == 0 {
				continue
			}
			for j := i + 1; j < n; j++ {
				if bit&(1<<j) == 0 {
					continue
				}
				if !e[i][j] {
					flag = false
					break
				}
			}
		}
		if flag {
			res = max(res, bitCount(bit))
		}
	}
	fmt.Println(res)
}

func bitCount(x int) (res int) {
	for x > 0 {
		res += x & 1
		x >>= 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
