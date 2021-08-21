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
		cnt := 0
		for i := 0; i < n; i++ {
			if bit&(1<<i) != 0 {
				ok := true
				for j := 0; j < n; j++ {
					if i != j && bit&(1<<j) != 0 {
						if !e[i][j] {
							ok = false
							break
						}
					}
				}
				if ok {
					cnt++
				}
			}
		}
		if res < cnt {
			res = cnt
		}
	}
	fmt.Println(res)
}
