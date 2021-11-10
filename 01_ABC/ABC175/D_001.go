package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	p := make([]int, n)
	c := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
		p[i]--
	}
	for i := range c {
		fmt.Scan(&c[i])
	}

	maxi := -1 << 60
	for i := 0; i < n; i++ {
		score := c[i]
		t := make([]int, 0)
		t = append(t, score)
		for j := p[i]; j != i; j = p[j] {
			score += c[j]
			t = append(t, score)
		}
		len := len(t)
		for j := 0; j < len; j++ {
			rest := k - j - 1
			if rest < 0 {
				break
			}
			maxi = max(maxi, t[j]+(rest/len)*max(score, 0))
		}
	}

	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
