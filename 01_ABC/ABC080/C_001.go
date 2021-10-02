package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 10)
		for j := range f[i] {
			fmt.Scan(&f[i][j])
		}
	}

	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, 11)
		for j := range p[i] {
			fmt.Scan(&p[i][j])
		}
	}

	res := -int(1e18)
	for bit := 1; bit < 1<<10; bit++ {
		sum := 0
		for i := 0; i < n; i++ {
			cnt := 0
			for j := 0; j < 10; j++ {
				if (bit>>j)&1 == 1 && f[i][j] != 0 {
					cnt++
				}
			}
			sum += p[i][cnt]
		}
		res = max(res, sum)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
